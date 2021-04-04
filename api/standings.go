package api

import (
	"context"
	"fmt"
	"strconv"
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

//Converts result into a flat data
func (service *StandingService) Convert(result *StandingResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid standing data")
	}

	var rows [][]string
	for _, standings := range result.API.Standings {
		if includeHead {
			rows = append(rows, service.getHeader())
		}
		if len(standings) > 0 {
			rows = append(rows, service.getData(standings[0]))
		}
	}

	return rows, nil
}

func (service *StandingService) getHeader() []string {
	var row []string
	row = append(row, "Rank")
	row = append(row, "TeamId")
	row = append(row, "TeamName")
	row = append(row, "Logo")
	row = append(row, "Group")
	row = append(row, "Status")
	row = append(row, "From")
	row = append(row, "Description")
	row = append(row, "AllStat.MatchsPlayed")
	row = append(row, "AllStat.Win")
	row = append(row, "AllStat.Draw")
	row = append(row, "AllStat.Lose")
	row = append(row, "AllStat.GoalsFor")
	row = append(row, "AllStat.GoalsAgainst")
	row = append(row, "HomeStat.MatchsPlayed")
	row = append(row, "HomeStat.Win")
	row = append(row, "HomeStat.Draw")
	row = append(row, "HomeStat.Lose")
	row = append(row, "HomeStat.GoalsFor")
	row = append(row, "HomeStat.GoalsAgainst")
	row = append(row, "AwayStat.MatchsPlayed")
	row = append(row, "AwayStat.Win")
	row = append(row, "AwayStat.Draw")
	row = append(row, "AwayStat.Lose")
	row = append(row, "AwayStat.GoalsFor")
	row = append(row, "AwayStat.GoalsAgainst")
	return row
}

func (service *StandingService) getData(s Standings) []string {
	var row []string
	row = append(row, strconv.Itoa(s.Rank))
	row = append(row, strconv.Itoa(s.TeamID))
	row = append(row, s.TeamName)
	row = append(row, s.Logo)
	row = append(row, s.Group)
	row = append(row, s.Status)
	row = append(row, s.Form)
	row = append(row, strconv.Itoa(s.AllStat.MatchsPlayed))
	row = append(row, strconv.Itoa(s.AllStat.Win))
	row = append(row, strconv.Itoa(s.AllStat.Draw))
	row = append(row, strconv.Itoa(s.AllStat.Lose))
	row = append(row, strconv.Itoa(s.AllStat.GoalsFor))
	row = append(row, strconv.Itoa(s.AllStat.GoalsAgainst))
	row = append(row, strconv.Itoa(s.HomeStat.MatchsPlayed))
	row = append(row, strconv.Itoa(s.HomeStat.Win))
	row = append(row, strconv.Itoa(s.HomeStat.Draw))
	row = append(row, strconv.Itoa(s.HomeStat.Lose))
	row = append(row, strconv.Itoa(s.HomeStat.GoalsFor))
	row = append(row, strconv.Itoa(s.HomeStat.GoalsAgainst))
	row = append(row, strconv.Itoa(s.AwayStat.MatchsPlayed))
	row = append(row, strconv.Itoa(s.AwayStat.Win))
	row = append(row, strconv.Itoa(s.AwayStat.Draw))
	row = append(row, strconv.Itoa(s.AwayStat.Lose))
	row = append(row, strconv.Itoa(s.AwayStat.GoalsFor))
	row = append(row, strconv.Itoa(s.AwayStat.GoalsAgainst))
	return row
}
