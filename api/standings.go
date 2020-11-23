package api

import "time"

//StandingService gets the Standings information for a league from api call
type StandingService service

//StandingResult contins the parsed result of api response of standings
type StandingResult struct {
	API struct {
		Results   int         `json:"results"`
		Standings []Standings `json:"standings"`
	} `json:"api"`
}

//Standings contains league table stangings
type Standings struct {
	Rank        int       `json:"rank"`
	TeamID      int       `json:"team_id"`
	TeamName    string    `json:"teamName"`
	Logo        string    `json:"logo"`
	Group       string    `json:"group"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
	AllStat     Stat      `json:"all"`
	HomeStat    Stat      `json:"home"`
	AwayStat    Stat      `json:"away"`
	GoalsDiff   int       `json:"goalsDiff"`
	Points      int       `json:"points"`
	LastUpdated time.Time `json:"lastUpdate"`
}

//Stat contains  team statistics
type Stat struct {
	MatchsPlayed int `json:"matchsPlayed"`
	Win          int `json:"win"`
	Draw         int `json:"draw"`
	Lose         int `json:"lose"`
	GoalsFor     int `json:"goalsFor"`
	GoalsAgainst int `json:"goalsAgainst"`
}
