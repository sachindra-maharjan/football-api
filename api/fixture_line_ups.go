package api

import (
	"context"
	"fmt"
	"strconv"
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
func (f *FixtureLineUpService) Convert(result *FixtureLineUpResult, includeHead bool) ([][]string, [][]string, error) {
	if result == nil {
		return nil, nil, fmt.Errorf("invalid league result.")
	}

	formation := [][]string{}
	startingXI := [][]string{}
	if includeHead {
		formation = append(formation, f.getFormationHead())
		startingXI = append(startingXI, f.getStartingXIHead())
	}

	for _, fixtureTeam := range result.API.LineUp {
		formation = append(formation, f.getFormationData(result.API.FixtureID, fixtureTeam))
		startingXI = append(startingXI, f.getStartingXIData(result.API.FixtureID, fixtureTeam))
	}

	return formation, startingXI, nil
}

//getData Returns flat array from an object
func (service *FixtureLineUpService) getFormationData(fixtureID int, team FixtureTeam) []string {
	var row []string
	row = append(row, strconv.Itoa(fixtureID))
	row = append(row, strconv.Itoa(team.CoachID))
	row = append(row, team.Coach)
	row = append(row, team.Formation)
	return row
}

//getStartingXIData Returns flat array from an object
func (service *FixtureLineUpService) getStartingXIData(fixtureID int, team FixtureTeam) []string {
	var row []string
	row = append(row, strconv.Itoa(fixtureID))

	for _, player := range team.StartingXI {
		row = append(row, fmt.Sprintf("%d|%d%s|%d|%s", player.TeamID, player.PlayerID, player.PlayerName, player.Number, player.Position))
	}
	return row
}

//getFormationHead Returns the array of head fields
func (service *FixtureLineUpService) getFormationHead() []string {
	var row []string
	row = append(row, "FixtureID")
	row = append(row, "CoachID")
	row = append(row, "CoachName")
	row = append(row, "Formation")
	return row
}

//getStartingXIHead Returns the array of head fields
func (service *FixtureLineUpService) getStartingXIHead() []string {
	var row []string
	row = append(row, "FixtureID")
	row = append(row, "TeamID|PlayerID|Payer|Number|Pos")
	return row
}
