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
		"league":    s.league,
		"standings": s.standings,
		"health":    s.health,
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
		basepath, _ := s.clientFlags(fetchCmd)

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		allIds := strings.Split(ids[0], ",")
		lastId := ids[len(allIds)-1]
		id, err := strconv.Atoi(lastId)

		if err != nil {
			return err
		}

		leagueResult, _, err := s.client.LeagueService.ListAll(context.Background(), id)
		if err != nil {
			return wrapError("could not fetch data", err)
		}

		fmt.Printf("fetched league data successfully. Total count:  %d \n", leagueResult.API.Results)

		if basepath != nil {
			leagueData, err := s.client.LeagueService.Convert(leagueResult, true)
			if err != nil {
				wrapError("unable to get data", err)
			}
			finalPath := s.getFileDestination(*basepath, "league.csv", true)
			s.writeData(finalPath, leagueData)
		}

		return nil
	}
}

func (s Switch) standings() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		standingsCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		standingsCmd.Var(&ids, "id", "The id of league to fetch data.")
		basepath, _ := s.clientFlags(standingsCmd)

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(standingsCmd); err != nil {
			return err
		}

		for _, id := range strings.Split(ids[0], ",") {
			leagueId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			standingResult, _, err := s.client.StandingService.GetLeagueStandings(context.Background(), leagueId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched standings data successfully. Total count:  %d \n", standingResult.API.Results)

			if basepath != nil {
				standingData, err := s.client.StandingService.Convert(standingResult, true)
				if err != nil {
					wrapError("unable to write flat data", err)
				}
				finalPath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%d/%s", leagueId, "standing.csv"),
					true,
				)
				s.writeData(finalPath, standingData)
			}
		}

		return nil
	}
}

func (s Switch) getFileDestination(basepath, filename string, delIfExists bool) string {
	dir, _ := filepath.Split(basepath + "/")
	finalDest := fmt.Sprintf("%s%s", dir, filename)

	if delIfExists && s.fileExists(finalDest) {
		os.Remove(finalDest)
	}
	return finalDest

}

func (s Switch) writeData(filepath string, data [][]string) error {
	if err := csvutil.Write(filepath, data); err != nil {
		wrapError("unable to write result", err)
	}
	fmt.Printf("data written to file %s successfully\n", filepath)
	return nil
}

func (s Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("check health of api")
		return nil
	}
}

func (s Switch) clientFlags(f *flag.FlagSet) (*string, *string) {
	basepath, format := "", ""
	f.StringVar(&format, "format", "", "The file format to save data.")
	f.StringVar(&basepath, "basepath", "", "The distination to save file.")
	return &basepath, &format
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

func (s Switch) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
