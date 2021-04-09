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
	Coach      string    `json:"coach,omitempty"`
	CoachID    int       `json:"coach_id,omitempty"`
	Formation  string    `json:"formation,omitempty"`
	StartingXI []StartXI `json:"startXI,omitempty"`
}

//StartXI Starting Player Info
type StartXI struct {
	TeamID     int    `json:"team_id,omitemtpy"`
	PlayerID   int    `json:"player_id,omitempty"`
	PlayerName string `json:"player,omitempty"`
	Number     int    `json:"number,omitempty"`
	Position   string `json:"pos,omitempty"`
}

//GetLineUpForFixture Returns team line up for fixture
func (l *FixtureLineUpService) GetLineUpForFixture(context context.Context, fixtureID int) (*FixtureLineUpResult, *Response, error) {
	req, err := l.client.NewRequest("GET", "lineups/"+fmt.Sprint(fixtureID), nil)
	if err != nil {
		return nil, nil, err
	}

	var fixtureLineUpResult *FixtureLineUpResult
	res, err := l.client.Do(context, req, &fixtureLineUpResult)

	if err != nil {
		return nil, nil, err
	}

	fixtureLineUpResult.API.FixtureID = fixtureID
	return fixtureLineUpResult, res, nil
}

//GetFlatDataWithHeader Returns flat data with header
func (f *FixtureLineUpService) Convert(result *FixtureLineUpResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid league result.")
	}

	startingXI := [][]string{}

	if includeHead {
		startingXI = append(startingXI, f.getStartingXIHead())
	}

	for _, fixtureTeam := range result.API.LineUp {
		startingXI = append(startingXI, f.getStartingXIData(result.API.FixtureID, fixtureTeam))
	}

	return startingXI, nil
}

//getStartingXIData Returns flat array from an object
func (service *FixtureLineUpService) getStartingXIData(fixtureID int, team FixtureTeam) []string {
	var row []string
	var playerXI string
	row = append(row, strconv.Itoa(fixtureID))
	row = append(row, strconv.Itoa(team.CoachID))
	row = append(row, team.Coach)
	row = append(row, team.Formation)

	for i, player := range team.StartingXI {
		if i == 0 {
			row = append(row, strconv.Itoa(player.TeamID))
		}
		playerXI = playerXI + fmt.Sprintf("%d|%s|%d|%s-", player.PlayerID, player.PlayerName, player.Number, player.Position)
	}
	row = append(row, strings.TrimRight(playerXI, "-"))
	return row
}

//getStartingXIHead Returns the array of head fields
func (service *FixtureLineUpService) getStartingXIHead() []string {
	var row []string
	row = append(row, "FixtureID")
	row = append(row, "CoachID")
	row = append(row, "CoachName")
	row = append(row, "Formation")
	row = append(row, "TeamID")
	row = append(row, "PlayerID|Payer|Number|Pos-PlayerID|Payer|Number|Pos")
	return row
}
