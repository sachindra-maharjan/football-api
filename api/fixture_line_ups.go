package api

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

//FixtureLineUpService gets the line ups of a fixture from api
type FixtureLineUpService service

//FixtureLineUpResult contains the parsed result from api response
type FixtureLineUpResult struct {
	API struct {
		Results   int                    `json:"results,omitempty"`
		LineUp    map[string]FixtureTeam `json:"lineUps,omitempty"`
		LeagueID  int                    `json:"leagueID,omitempty"`
		FixtureID int                    `json:"fixtureID,omitempty"`
	} `json:"api"`
}

//TeamInfo Team Info
type TeamInfo map[string]FixtureTeam

//LineUp Linup for fixture
type LineUp struct {
	Team TeamInfo
}

//FixtureTeam Fixture Team Info
type FixtureTeam struct {
	Coach      string       `json:"coach,omitempty"`
	CoachID    int          `json:"coach_id,omitempty"`
	Formation  string       `json:"formation,omitempty"`
	StartingXI []PlayerInfo `json:"startXI,omitempty"`
	Substitute []PlayerInfo `json:"substitutes,omitempty"`
}

//Player Starting Player Info
type PlayerInfo struct {
	TeamID     int    `json:"team_id,omitemtpy"`
	PlayerID   int    `json:"player_id,omitempty"`
	PlayerName string `json:"player,omitempty"`
	Number     int    `json:"number,omitempty"`
	Position   string `json:"pos,omitempty"`
}

//GetLineUpForFixture Returns team line up for fixture
func (l *FixtureLineUpService) GetLineUpForFixture(context context.Context, leagueID int, fixtureID int) (*FixtureLineUpResult, *Response, error) {
	req, err := l.client.NewRequest("GET", "lineups/"+fmt.Sprint(fixtureID), nil)
	if err != nil {
		return nil, nil, err
	}

	var fixtureLineUpResult *FixtureLineUpResult
	res, err := l.client.Do(context, req, &fixtureLineUpResult)

	if err != nil {
		return nil, nil, err
	}

	fixtureLineUpResult.API.LeagueID = leagueID
	fixtureLineUpResult.API.FixtureID = fixtureID
	return fixtureLineUpResult, res, nil
}

//GetFlatDataWithHeader Returns flat data with header
func (f *FixtureLineUpService) Convert(result *FixtureLineUpResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid league result.")
	}

	linup := [][]string{}

	if includeHead {
		linup = append(linup, f.getStartingXIHead())
	}

	for _, fixtureTeam := range result.API.LineUp {
		linup = append(linup, f.getLineup(result.API.LeagueID, result.API.FixtureID, fixtureTeam))
	}

	return linup, nil
}

//getStartingXIData Returns flat array from an object
func (service *FixtureLineUpService) getLineup(leagueID int, fixtureID int, team FixtureTeam) []string {
	var row []string
	row = append(row, strconv.Itoa(leagueID))
	row = append(row, strconv.Itoa(fixtureID))
	row = append(row, strconv.Itoa(team.CoachID))
	row = append(row, team.Coach)
	row = append(row, team.Formation)

	playerXI, teamID := getPlayers(team.StartingXI, row)
	row = append(row, strconv.Itoa(teamID))
	row = append(row, strings.TrimRight(playerXI, "|"))

	substitutes, _ := getPlayers(team.Substitute, row)
	row = append(row, strings.TrimRight(substitutes, "|"))

	return row
}

func getPlayers(players []PlayerInfo, row []string) (string, int) {
	var playerInfo string
	var teamID int
	for i, player := range players {
		if i == 0 {
			teamID = player.PlayerID
		}
		playerInfo = playerInfo + fmt.Sprintf("%d-%s-%d-%s|", player.PlayerID, player.PlayerName, player.Number, player.Position)
	}
	return playerInfo, teamID
}

//getStartingXIHead Returns the array of head fields
func (service *FixtureLineUpService) getStartingXIHead() []string {
	var row []string
	row = append(row, "league_id")
	row = append(row, "fixture_id")
	row = append(row, "coach_id")
	row = append(row, "coach_name")
	row = append(row, "formation")
	row = append(row, "team_id")
	row = append(row, "startingXI")
	row = append(row, "substitutes")
	return row
}
