package api

import (
	"context"
	"fmt"
	"strconv"
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
	LeagueID    int    `json:"league_id,omitempty"`
	Name        string `json:"name,omitempty"`
	LeagueType  string `json:"type,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Season      int    `json:"season,omitempty"`
	SeasonStart string `json:"season_start,omitempty"`
	SeasonEnd   string `json:"season_end,omitempty"`
	LogoURL     string `json:"logo,omitempty"`
	FlagURL     string `json:"flag,omitempty"`
	Standings   int    `json:"standings,omitempty"`
	IsCurrent   int    `json:"is_current,omitempty"`
	Coverage    struct {
		Standings       bool `json:"standings,omitempty"`
		FixtureCoverage struct {
			Events           bool `json:"events,omitempty"`
			Lineups          bool `json:"lineups,omitempty"`
			Statistics       bool `json:"statistics,omitempty"`
			PlayerStatistics bool `json:"players_statistics,omitempty"`
		} `json:"fixtures,omitempty"`
		Players     bool `json:"players,omitempty"`
		TopScorers  bool `json:"topScorers,omitempty"`
		Predictions bool `json:"predictions,omitempty"`
		Odds        bool `json:"odds,omitempty"`
	} `json:"coverage,omitempty"`
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

//GetFlatDataWithHeader Returns flat data with header
func (l *LeagueService) Convert(leagueResult *LeagueResult, includeHead bool) ([][]string, error) {
	if leagueResult == nil {
		return nil, fmt.Errorf("invalid league result.")
	}

	rows := [][]string{}
	if includeHead {
		rows = append(rows, l.getHead())
	}

	for _, league := range leagueResult.API.Leagues {
		rows = append(rows, l.getData(league))
	}

	return rows, nil
}

//GetFlat Returns flat array from an object
func (service *LeagueService) getData(league League) []string {
	var row []string
	row = append(row, strconv.Itoa(league.LeagueID))
	row = append(row, league.Name)
	row = append(row, league.LeagueType)
	row = append(row, league.CountryCode)
	row = append(row, league.Country)
	row = append(row, strconv.Itoa(league.Season))
	row = append(row, league.SeasonStart)
	row = append(row, league.SeasonEnd)
	row = append(row, league.LogoURL)
	row = append(row, league.FlagURL)
	row = append(row, strconv.FormatBool(league.IsCurrent > 0))
	row = append(row, strconv.FormatBool(league.Coverage.FixtureCoverage.Events))
	row = append(row, strconv.FormatBool(league.Coverage.FixtureCoverage.Lineups))
	row = append(row, strconv.FormatBool(league.Coverage.FixtureCoverage.Statistics))
	row = append(row, strconv.FormatBool(league.Coverage.FixtureCoverage.PlayerStatistics))
	row = append(row, strconv.FormatBool(league.Coverage.Standings))
	row = append(row, strconv.FormatBool(league.Coverage.Players))
	row = append(row, strconv.FormatBool(league.Coverage.Predictions))
	row = append(row, strconv.FormatBool(league.Coverage.Odds))
	row = append(row, strconv.FormatBool(league.Coverage.TopScorers))
	return row
}

//GetHead Returns the array of head fields
func (service *LeagueService) getHead() []string {
	var row []string
	row = append(row, "league_id")
	row = append(row, "name")
	row = append(row, "type")
	row = append(row, "country_code")
	row = append(row, "country")
	row = append(row, "season")
	row = append(row, "season_start")
	row = append(row, "season_end")
	row = append(row, "logo_url")
	row = append(row, "flag_url")
	row = append(row, "current_season")
	row = append(row, "events")
	row = append(row, "line_ups")
	row = append(row, "statistics")
	row = append(row, "player_statistics")
	row = append(row, "standings")
	row = append(row, "players")
	row = append(row, "predictions")
	row = append(row, "odds")
	row = append(row, "top_scorers")
	return row
}
