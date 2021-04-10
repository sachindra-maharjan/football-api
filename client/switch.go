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

var (
	apiKeys = []string{}
)

type idsFlag []string
type authKeysFlag []string

func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

func (list *idsFlag) Set(id string) error {
	*list = append(*list, id)
	return nil
}

func (list authKeysFlag) String() string {
	return strings.Join(list, ",")
}

func (list *authKeysFlag) Set(authKey string) error {
	*list = append(*list, authKey)
	return nil
}

type ApiClient interface {
	League(leagueId int) (*api.LeagueResult, error)
	Healthy(host string) bool
}

func NewSwitch() Switch {
	httpClient := api.NewClient(nil, apiKeys)

	s := Switch{
		client: httpClient,
	}

	s.commands = map[string]func() func(string) error{
		"league":         s.league,
		"standings":      s.standings,
		"fixtures":       s.fixtures,
		"fixture-event":  s.fixtureEvent,
		"fixture-lineup": s.fixtureLineup,
		"player-stat":    s.playerStat,
		"top-scorer":     s.topScorer,
		"team":           s.team,
		"health":         s.health,
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
		authKeys := authKeysFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		fetchCmd.Var(&ids, "leagueId", "The league id of league to fetch data.")
		basepath, _ := s.clientFlags(fetchCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

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
		authKeys := authKeysFlag{}
		standingsCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		standingsCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		standingsCmd.Var(&ids, "leagueId", "The league id of league to fetch data.")
		basepath, _ := s.clientFlags(standingsCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(standingsCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

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

func (s Switch) fixtures() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}

		fixturesCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fixturesCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		fixturesCmd.Var(&ids, "leagueId", "The league id of league to fetch data.")
		basepath, _ := s.clientFlags(fixturesCmd)

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(fixturesCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			leagueId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			fixtureResult, _, err := s.client.FixtureService.GetFixturesByLeagueID(context.Background(), leagueId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched fixture data successfully. Total count:  %d \n", fixtureResult.API.Results)

			if basepath != nil {
				fixtureData, err := s.client.FixtureService.Convert(fixtureResult, true)
				if err != nil {
					wrapError("unable to write flat data", err)
				}
				finalPath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%d/%s", leagueId, "fixtures.csv"),
					true,
				)
				s.writeData(finalPath, fixtureData)
			}
		}

		return nil
	}
}

func (s Switch) team() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}
		teamCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		teamCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		teamCmd.Var(&ids, "leagueId", "The league id of league to fetch data.")
		basepath, _ := s.clientFlags(teamCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(teamCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			leagueId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			teamResult, _, err := s.client.TeamService.ListTeamsByLeagueID(context.Background(), leagueId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched team data successfully. Total count:  %d \n", teamResult.API.Results)

			if basepath != nil {
				teamData, err := s.client.TeamService.Convert(teamResult, true)
				if err != nil {
					wrapError("unable to write flat data", err)
				}
				finalPath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%d/%s", leagueId, "team.csv"),
					true,
				)
				s.writeData(finalPath, teamData)
			}
		}

		return nil
	}
}

func (s Switch) fixtureEvent() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}
		fixtureEventCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fixtureEventCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		fixtureEventCmd.Var(&ids, "fixtureId", "The fixture id of league to fetch data.")
		leagueId, basepath, _ := s.clientEventFlags(fixtureEventCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(fixtureEventCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			fixtureId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			fixtureEventResult, _, err := s.client.FixtureEventService.GetFixtureEvent(context.Background(), fixtureId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched fixture event data successfully. Total count:  %d \n", fixtureEventResult.API.Results)

			if basepath != nil {
				fixtureData, err := s.client.FixtureEventService.Convert(fixtureEventResult, true)
				if err != nil {
					wrapError("unable to write flat data", err)
				}
				finalPath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%s/%s", *leagueId, "fixture-event.csv"),
					false,
				)
				s.writeData(finalPath, fixtureData)
			}
		}
		return nil
	}
}

func (s Switch) fixtureLineup() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}
		fixtureEventCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fixtureEventCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		fixtureEventCmd.Var(&ids, "fixtureId", "The fixture id of league to fetch data.")
		leagueId, basepath, _ := s.clientEventFlags(fixtureEventCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(fixtureEventCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			fixtureId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			fixtureLineUpResult, _, err := s.client.FixtureLineUpService.GetLineUpForFixture(context.Background(), fixtureId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched fixture linup data successfully. Total count:  %d \n", fixtureLineUpResult.API.Results)

			if basepath != nil {
				filepath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%s/%s", *leagueId, "fixture-lineup.csv"),
					false,
				)

				lineup, err := s.client.FixtureLineUpService.Convert(fixtureLineUpResult,
					!s.fileExists(filepath))

				if err != nil {
					wrapError("unable to write flat data", err)
				}
				s.writeData(filepath, lineup)
			}
		}
		return nil
	}
}

func (s Switch) playerStat() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}
		playerStatCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		playerStatCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		playerStatCmd.Var(&ids, "fixtureId", "The fixture id of league to fetch data.")
		leagueId, basepath, _ := s.clientEventFlags(playerStatCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(playerStatCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			fixtureId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			playerStatResult, _, err := s.client.PlayerStatService.GetPlayerStatByFixtureID(context.Background(), fixtureId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched fixture linup data successfully. Total count:  %d \n", playerStatResult.API.Results)

			if basepath != nil {
				filepath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%s/%s", *leagueId, "player-fixture-stat.csv"),
					false,
				)

				lineup, err := s.client.PlayerStatService.Convert(playerStatResult,
					!s.fileExists(filepath))

				if err != nil {
					wrapError("unable to write flat data", err)
				}
				s.writeData(filepath, lineup)
			}
		}
		return nil
	}
}

func (s Switch) topScorer() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		authKeys := authKeysFlag{}
		topScorerCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		topScorerCmd.Var(&authKeys, "authKey", "The authentication key to access api.")
		topScorerCmd.Var(&ids, "leagueId", "The fixture id of league to fetch data.")
		basepath, _ := s.clientFlags(topScorerCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(topScorerCmd); err != nil {
			return err
		}

		apiKeys = strings.Split(authKeys[0], ",")

		for _, id := range strings.Split(ids[0], ",") {
			leagueId, err := strconv.Atoi(id)

			if err != nil {
				wrapError("unable to convert string to int", err)
			}

			topScorerResult, _, err := s.client.TopScorerService.List(context.Background(), leagueId)
			if err != nil {
				return wrapError("could not fetch data", err)
			}

			fmt.Printf("fetched top scorer data successfully. Total count:  %d \n", topScorerResult.API.Results)

			if basepath != nil {
				filepath := s.getFileDestination(*basepath,
					fmt.Sprintf("leagueID_%d/%s", leagueId, "topScorer.csv"),
					true,
				)

				topScorer, err := s.client.TopScorerService.Convert(topScorerResult, true)

				if err != nil {
					wrapError("unable to write flat data", err)
				}
				s.writeData(filepath, topScorer)
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

func (s Switch) clientEventFlags(f *flag.FlagSet) (*string, *string, *string) {
	leagueId := ""
	f.StringVar(&leagueId, "leagueId", "", "The id of the league.")
	basepath, format := s.clientFlags(f)
	return &leagueId, basepath, format
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
