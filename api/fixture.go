package api

import "time"

//FixtureService gets the fixture information from api
type FixtureService service

//FixtureResult contains the parsed result from api response
type FixtureResult struct {
	API struct {
		Results  int       `json:"results"`
		Fixtures []Fixture `json:"fixtures"`
	} `json:"api"`
}

//Fixture contains the fixture infromation
type Fixture struct {
	FixtureID int `json:"fixture_id"`
	LeagueID  int `json:"league_id"`
	League    struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
	} `json:"league"`
	EventDate       time.Time `json:"event_date"`
	EventTimestamp  time.Time `json:"event_timestamp"`
	FirstHalfStart  int       `json:"firstHalfStart"`
	SecondHalfStart int       `json:"secondHalfStart"`
	Round           string    `json:"round"`
	Status          string    `json:"status"`
	StatusShort     string    `json:"statusShort"`
	Elapsed         int       `json:"elapsed"`
	Venue           string    `json:"venue"`
	Referee         string    `json:"referee,omitempty"`
	HomeTeam        team      `json:"homeTeam"`
	AwayTeam        team      `json:"awayTeam"`
	GoalsHomeTeam   int       `json:"goalsHomeTeam"`
	GoalsAwayTeam   int       `json:"goalsAwayTeam"`
	Score           struct {
		HalfTime  string `json:"halftime"`
		FullTime  string `json:"fulltime"`
		ExtraTime string `json:"extratime,omitempty"`
		Penalty   string `json:"penalty,omitempty"`
	}
}

type team struct {
	TeamID   int    `json:"team_id"`
	TeamName string `json:"team_name"`
	Logo     string `json:"logo"`
}
