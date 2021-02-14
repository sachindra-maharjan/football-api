package api

import (
	"context"
	"fmt"
)

//FixtureLineUpService gets the line ups of a fixture from api
type FixtureLineUpService service

//FixtureLineUpResult contains the parsed result from api response
type FixtureLineUpResult struct {
	API struct {
		Results int                    `json:"results,omitempty"`
		LineUp  map[string]FixtureTeam `json:"lineUps,omitempty"`
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
		return nil, nil, nil
	}

	return fixtureLineUpResult, res, nil

}
