package api

import (
	"context"
	"fmt"
	"strconv"
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

//Converts result into a flat data
func (service *FixtureService) Convert(result *FixtureResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid standing data")
	}

	var rows [][]string

	if includeHead {
		rows = append(rows, service.getHeader())
	}

	for _, fixture := range result.API.Fixtures {
		rows = append(rows, service.getData(fixture))
	}

	return rows, nil
}

func (service *FixtureService) getHeader() []string {
	var row []string
	row = append(row, "league_id")
	row = append(row, "fixture_id")
	row = append(row, "league_name")
	row = append(row, "league_country")
	row = append(row, "league_logo")
	row = append(row, "league_flag")
	row = append(row, "event_date")
	row = append(row, "event_timestamp")
	row = append(row, "first_half_start")
	row = append(row, "second_half_start")
	row = append(row, "round")
	row = append(row, "status")
	row = append(row, "status_short")
	row = append(row, "elapsed")
	row = append(row, "venue")
	row = append(row, "referee")
	row = append(row, "home_team_team_id")
	row = append(row, "home_team_team_name")
	row = append(row, "home_team_logo")
	row = append(row, "away_team_team_id")
	row = append(row, "away_team_team_name")
	row = append(row, "away_team_logo")
	row = append(row, "goals_home_team")
	row = append(row, "goals_away_team")
	row = append(row, "score_half_time")
	row = append(row, "score_full_time")
	row = append(row, "score_extra_time")
	row = append(row, "score_penalty")
	return row
}

func (service *FixtureService) getData(f Fixture) []string {
	var row []string
	row = append(row, strconv.Itoa(f.LeagueID))
	row = append(row, strconv.Itoa(f.FixtureID))
	row = append(row, f.League.Name)
	row = append(row, f.League.Country)
	row = append(row, f.League.Logo)
	row = append(row, f.League.Flag)
	row = append(row, f.EventDate.String())
	row = append(row, strconv.FormatInt(f.EventTimestamp, 10))
	row = append(row, strconv.Itoa(f.FirstHalfStart))
	row = append(row, strconv.Itoa(f.SecondHalfStart))
	row = append(row, f.Round)
	row = append(row, f.Status)
	row = append(row, f.StatusShort)
	row = append(row, strconv.Itoa(f.Elapsed))
	row = append(row, f.Venue)
	row = append(row, f.Referee)
	row = append(row, strconv.Itoa(f.HomeTeam.TeamID))
	row = append(row, f.HomeTeam.TeamName)
	row = append(row, f.HomeTeam.Logo)
	row = append(row, strconv.Itoa(f.AwayTeam.TeamID))
	row = append(row, f.AwayTeam.TeamName)
	row = append(row, f.AwayTeam.Logo)
	row = append(row, strconv.Itoa(f.GoalsHomeTeam))
	row = append(row, strconv.Itoa(f.GoalsAwayTeam))
	row = append(row, f.Score.HalfTime)
	row = append(row, f.Score.FullTime)
	row = append(row, f.Score.ExtraTime)
	row = append(row, f.Score.Penalty)
	return row
}
