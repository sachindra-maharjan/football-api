package api

import (
	"context"
	"fmt"
)

//LeagueService gets league informtion from api call
type LeagueService service

//LeagueResult contins the parsed result of api response of leagues
type LeagueResult struct {
	API struct {
		Results int      `json:"results"`
		Leagues []League `json:"leagues"`
	} `json:"api"`
}

//League contains information about league
type League struct {
	LeagueID    int    `json:"league_id"`
	Name        string `json:"name"`
	LeagueType  string `json:"type"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Season      int    `json:"season"`
	SeasonStart string `json:"season_start"`
	SeasonEnd   string `json:"season_end"`
	LogoURL     string `json:"logo"`
	FlagURL     string `json:"flag"`
	Standings   int    `json:"standings"`
	IsCurrent   int    `json:"is_current"`
	Coverage    struct {
		Standings       bool `json:"standings"`
		FixtureCoverage struct {
			Events           bool `json:"events"`
			Lineups          bool `json:"lineups"`
			Statistics       bool `json:"statistics"`
			PlayerStatistics bool `json:"players_statistics"`
		} `json:"fixtures"`
		Players     bool `json:"players"`
		TopScorers  bool `json:"topScorers"`
		Predictions bool `json:"predictions"`
		Odds        bool `json:"odds"`
	} `json:"coverage"`
}

//ListAll gets the all available season by league ID
func (l *LeagueService) ListAll(ctx context.Context, leagueID int) (*LeagueResult, *Response, error) {
	req, err := l.client.NewRequest("GET", "leagues/seasonsAvailable/"+fmt.Sprint(leagueID), nil)
	if err != nil {
		return nil, nil, err
	}

	var leagueResult *LeagueResult
	res, err := l.client.Do(ctx, req, &leagueResult)
	if err != nil {
		return nil, nil, err
	}

	return leagueResult, res, err
}
