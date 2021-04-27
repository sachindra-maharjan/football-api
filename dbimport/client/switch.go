package client

import (
	"casino_royal/vault/database"
	fileutil "casino_royal/vault/util"
	"context"
	"flag"
	"fmt"
	"os"
)

type Key struct {
	database, command string
}

type Switch struct {
	commands map[Key]func() func(string) error
}

func NewSwitch() (Switch, error) {
	s := Switch{}

	s.commands = map[Key]func() func(string) error{
		{database: "firestore", command: "league"}:    s.league,
		{database: "firestore", command: "fixtures"}:  s.fixtures,
		{database: "firestore", command: "standings"}: s.standings,
	}

	return s, nil
}

func (s Switch) Switch() error {
	database := os.Args[1]
	cmdName := os.Args[2]
	key := Key{
		database: database,
		command:  cmdName,
	}
	cmd, ok := s.commands[key]
	if !ok {
		return fmt.Errorf("invalid command '%s'\n", cmdName)
	}
	return cmd()(cmdName)
}

func wrapError(customMessage string, originalError error) error {
	return fmt.Errorf("%s : %v", customMessage, originalError)
}

func (s Switch) clientFlags(f *flag.FlagSet) (*string, *string, *string) {
	projectID, file, format := "", "", ""
	f.StringVar(&file, "file", "", "The source file.")
	f.StringVar(&format, "format", "", "The format of source file.")
	f.StringVar(&projectID, "projectId", "", "Google cloud project ID.")
	return &file, &format, &projectID
}

func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 4 && os.Args[3] == "--help" {
		return nil
	}

	fmt.Printf("%d %s %s\n", len(os.Args), os.Args[0], os.Args[1])

	if len(os.Args)-1 < minArgs {
		fmt.Printf(
			"incorrect use of %s\n%s %s --help\n",
			os.Args[1], os.Args[0], os.Args[1],
		)
		return fmt.Errorf(
			"%s expected at least %d arg(s), %d provided",
			os.Args[1], minArgs, len(os.Args)-1,
		)
	}
	return nil
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[3:])
	if err != nil {
		return wrapError("could not parse '"+cmd.Name()+"' command flags", err)
	}
	return nil
}

func (s Switch) Help() {
	var help string
	for key := range s.commands {
		help += key.database + "\t" + key.command + "\t --help\n"
	}
	fmt.Printf("Usuage of: %s:\n <database> <command> [<agrs>]\n%s", os.Args[0], help)
}

func (s Switch) readAll(file string) ([][]string, error) {
	reader, err := fileutil.FileReader(file)
	if err != nil {
		return nil, err
	}
	return reader.ReadAll()
}

func (s Switch) parseCommand(cmd string) (string, [][]string, error) {
	leagueCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
	file, _, projectID := s.clientFlags(leagueCmd)

	if err := s.checkArgs(3); err != nil {
		return "", nil, err
	}

	if err := s.parseCmd(leagueCmd); err != nil {
		return "", nil, err
	}

	if *projectID == "" {
		*projectID = os.Getenv("GCLOUD_PROJECT_ID")
		if *projectID == "" {
			return "", nil, fmt.Errorf("GCloud Project ID not found." +
				"Project ID can be set in envionment variable 'GCLOUD_PROJECT_ID' OR " +
				"passed as argument.")
		}
	}

	fmt.Printf("Using ProjectID %s \n", *projectID)

	fmt.Println("Reading records from file.")
	records, err := s.readAll(*file)
	fmt.Printf("Read successfully. Found %d lines.\n", len(records))

	if err != nil {
		return "", nil, err
	}
	return *projectID, records, nil
}

func (s Switch) league() func(string) error {
	return func(cmd string) error {
		projectID, records, err := s.parseCommand(cmd)

		if err != nil {
			return err
		}

		client, err := database.NewClient(context.Background(), projectID)
		if err != nil {
			return err
		}
		fmt.Println("Inserting to firestore...")
		client.LeagueService.Add(context.Background(), records[1:][0:])

		fmt.Println("Inserting to firestore complete.")

		return nil
	}
}

func (s Switch) fixtures() func(string) error {
	return func(cmd string) error {
		projectID, records, err := s.parseCommand(cmd)

		if err != nil {
			return err
		}

		client, err := database.NewClient(context.Background(), projectID)
		if err != nil {
			return err
		}
		fmt.Println("Inserting to firestore...")

		client.FixtureService.Add(context.Background(), records[1:][0:])

		fmt.Println("Inserting to firestore complete.")

		return nil
	}
}

func (s Switch) standings() func(string) error {
	return func(cmd string) error {
		projectID, records, err := s.parseCommand(cmd)

		if err != nil {
			return err
		}

		if len(records[0]) != 27 {
			return fmt.Errorf("Invalid file. Please provide the correct file containing standings data.")
		}

		client, err := database.NewClient(context.Background(), projectID)
		if err != nil {
			return err
		}
		fmt.Println("Inserting to firestore...")

		client.StandingsService.Add(context.Background(), records[1:][0:])

		fmt.Println("Inserting to firestore complete.")

		return nil
	}
}
