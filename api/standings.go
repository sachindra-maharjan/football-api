package api

import (
	"context"
	"fmt"
)

//StandingService gets the Standings information for a league from api call
type StandingService service

//StandingResult contins the parsed result of api response of standings
type StandingResult struct {
	API struct {
		Results   int           `json:"results,omitempty"`
		Standings [][]Standings `json:"standings,omitempty"`
	} `json:"api"`
}

//Standings contains league table stangings
type Standings struct {
	Rank        int    `json:"rank,omitempty"`
	TeamID      int    `json:"team_id,omitempty"`
	TeamName    string `json:"teamName,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Group       string `json:"group,omitempty"`
	Status      string `json:"status,omitempty"`
	Form        string `json:"forme,omitempty"`
	Description string `json:"description,omitempty"`
	AllStat     Stat   `json:"all,omitempty"`
	HomeStat    Stat   `json:"home,omitempty"`
	AwayStat    Stat   `json:"away,omitempty"`
	GoalsDiff   int    `json:"goalsDiff,omitempty"`
	Points      int    `json:"points,omitempty"`
	LastUpdated string `json:"lastUpdate,omitempty"`
}

//Stat contains  team statistics
type Stat struct {
	MatchsPlayed int `json:"matchsPlayed,omitempty"`
	Win          int `json:"win,omitempty"`
	Draw         int `json:"draw,omitempty"`
	Lose         int `json:"lose,omitempty"`
	GoalsFor     int `json:"goalsFor,omitempty"`
	GoalsAgainst int `json:"goalsAgainst,omitempty"`
}

//GetLeagueStandings service retuns the current league standings
func (service *StandingService) GetLeagueStandings(ctx context.Context, leagueID int) (*StandingResult, *Response, error) {
	req, err := service.client.NewRequest("GET", "leagueTable/"+fmt.Sprint(leagueID), nil)
	if err != nil {
		return nil, nil, err
	}

	var standingResult *StandingResult

	res, err := service.client.Do(ctx, req, &standingResult)
	if err != nil {
		return nil, nil, err
	}

	return standingResult, res, err
}
