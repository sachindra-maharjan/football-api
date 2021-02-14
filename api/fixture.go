package api

import (
	"context"
	"fmt"
	"time"
)

//FixtureService gets the fixture information from api
type FixtureService service

//FixtureResult contains the parsed result from api response
type FixtureResult struct {
	API struct {
		Results  int       `json:"results,omitempty"`
		Fixtures []Fixture `json:"fixtures,omitempty"`
	} `json:"api"`
}

//Fixture contains the fixture infromation
type Fixture struct {
	FixtureID int `json:"fixture_id,omitempty"`
	LeagueID  int `json:"league_id,omitempty"`
	League    struct {
		Name    string `json:"name,omitempty"`
		Country string `json:"country,omitempty"`
		Logo    string `json:"logo,omitempty"`
		Flag    string `json:"flag,omitempty"`
	} `json:"league"`
	EventDate       time.Time `json:"event_date,omitempty"`
	EventTimestamp  int64     `json:"event_timestamp,omitempty"`
	FirstHalfStart  int       `json:"firstHalfStart,omitempty"`
	SecondHalfStart int       `json:"secondHalfStart,omitempty"`
	Round           string    `json:"round,omitempty"`
	Status          string    `json:"status,omitempty"`
	StatusShort     string    `json:"statusShort,omitempty"`
	Elapsed         int       `json:"elapsed,omitempty"`
	Venue           string    `json:"venue,omitempty"`
	Referee         string    `json:"referee,omitempty,omitempty"`
	HomeTeam        team      `json:"homeTeam,omitempty"`
	AwayTeam        team      `json:"awayTeam,omitempty"`
	GoalsHomeTeam   int       `json:"goalsHomeTeam,omitempty"`
	GoalsAwayTeam   int       `json:"goalsAwayTeam,omitempty"`
	Score           struct {
		HalfTime  string `json:"halftime,omitempty"`
		FullTime  string `json:"fulltime,omitempty"`
		ExtraTime string `json:"extratime,omitempty,omitempty"`
		Penalty   string `json:"penalty,omitempty,omitempty"`
	}
}

type team struct {
	TeamID   int    `json:"team_id,omitempty"`
	TeamName string `json:"team_name,omitempty"`
	Logo     string `json:"logo,omitempty"`
}

//GetFixturesByLeagueID Return all fixtures for league
func (f *FixtureService) GetFixturesByLeagueID(context context.Context, leagueID int) (*FixtureResult, *Response, error) {
	r, err := f.client.NewRequest("GET", "fixtures/league/"+fmt.Sprint(leagueID), nil)
	if err != nil {
		return nil, nil, err
	}

	var fixtureResult *FixtureResult
	resp, err := f.client.Do(context, r, &fixtureResult)
	if err != nil {
		return nil, nil, err
	}

	return fixtureResult, resp, nil

}
