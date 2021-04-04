package client

import (
	"casino_royal/vault/api"
	csvutil "casino_royal/vault/util"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type idsFlag []string

func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

func (list *idsFlag) Set(id string) error {
	*list = append(*list, id)
	return nil
}

type ApiClient interface {
	League(leagueId int) (*api.LeagueResult, error)
	Healthy(host string) bool
}

func NewSwitch() Switch {
	httpClient := api.NewClient(nil)

	s := Switch{
		client: httpClient,
	}

	s.commands = map[string]func() func(string) error{
		"league": s.league,
		"health": s.health,
	}

	return s
}

type Switch struct {
	client        *api.Client
	backendApiUri string
	commands      map[string]func() func(string) error
}

func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command '%s'\n", cmdName)
	}
	return cmd()(cmdName)
}

func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usuage of: %s:\n <command> [<agrs>]\n%s", os.Args[0], help)
}

func (s Switch) league() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&ids, "id", "The id of league to fetch data.")
		dest, _ := s.clientFlags(fetchCmd)

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		lastId := ids[len(ids)-1]
		id, err := strconv.Atoi(lastId)

		if err != nil {
			return err
		}

		leagueResult, _, err := s.client.LeagueService.ListAll(context.Background(), id)
		if err != nil {
			return wrapError("could not fetch data", err)
		}

		fmt.Printf("fetched league data successfully. Total count:  %d \n", leagueResult.API.Results)

		if dest != nil {
			dir, _ := filepath.Split(*dest + "/")
			filepath := dir + "/league.csv"

			leagueData, err := s.client.LeagueService.GetFlatDataWithHeader(leagueResult)
			if err != nil {
				return nil
			}

			if err = csvutil.Write(filepath, leagueData); err != nil {
				wrapError("unable to write result", err)
			}

			fmt.Printf("data written to file %s successfully\n", filepath)
		}

		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("check health of api")
		return nil
	}
}

func (s Switch) clientFlags(f *flag.FlagSet) (*string, *string) {
	filepath, format := "", ""
	f.StringVar(&format, "format", "", "The file format to save data.")
	f.StringVar(&filepath, "filepath", "", "The distination to save file.")
	return &filepath, &format
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("could not parse '"+cmd.Name()+"' command flags", err)
	}
	return nil
}

func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}

	if len(os.Args)-2 < minArgs {
		fmt.Printf(
			"incorrect use of %s\n%s %s --help\n",
			os.Args[1], os.Args[0], os.Args[1],
		)
		return fmt.Errorf(
			"%s expected ast least %d arg(s), %d provided",
			os.Args[1], minArgs, len(os.Args)-2,
		)
	}
	return nil
}
