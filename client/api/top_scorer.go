package api

import (
	"context"
	"fmt"
	"strconv"
)

//TopScorerService gets the top scorer players list from api
type TopScorerService service

//TopScorerResult contains the parsed result from api response
type TopScorerResult struct {
	API struct {
		Results   int         `json:"results,omitempty"`
		TopScorer []TopScorer `json:"topscorers,omitempty"`
		LeagueID  int         `json:"leagueId,omitempty"`
	} `json:"api"`
}

//TopScorer contains the top scorer player information
type TopScorer struct {
	PlayerID    int    `json:"player_id,omitempty"`
	PlayerName  string `json:"player_name,omitempty"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	Position    string `json:"position,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	TeamID      int    `json:"team_id,omitempty"`
	TeamName    string `json:"team_name,omitempty"`
	Games       struct {
		Appearences   int `json:"appearences,omitempty"`
		MinutesPlayed int `json:"minutes_played,omitempty"`
	} `json:"games"`
	Goals struct {
		Total    int `json:"total,omitempty"`
		Assists  int `json:"assists,omitempty"`
		Conceded int `json:"conceded,omitempty"`
		Saves    int `json:"saves,omitempty"`
	} `json:"goals"`
	Shots struct {
		Total int `json:"total,omitempty"`
		On    int `json:"on,omitempty"`
	} `json:"shots"`
	Penalty struct {
		Won      int `json:"won,omitempty"`
		Commited int `json:"commited,omitempty"`
	} `json:"penalty"`
	Cards struct {
		Yellow       int `json:"yello,omitempty"`
		SecondYellow int `json:"second_yellow,omitempty"`
		Red          int `json:"red,omitempty"`
	}
}

func (t *TopScorerService) List(context context.Context, leagueId int) (*TopScorerResult, *Response, error) {

	r, err := t.client.NewRequest("GET", "topscorers/"+fmt.Sprint(leagueId), nil)
	if err != nil {
		return nil, nil, err
	}

	var topScorerResult *TopScorerResult
	response, err := t.client.Do(context, r, &topScorerResult)

	if err != nil {
		return nil, nil, err
	}

	topScorerResult.API.LeagueID = leagueId
	return topScorerResult, response, nil

}

func (t *TopScorerService) Convert(result *TopScorerResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid league result.")
	}

	rows := [][]string{}

	if includeHead {
		rows = append(rows, t.head())
	}

	for _, topScorer := range result.API.TopScorer {
		rows = append(rows, t.topScorer(result.API.LeagueID, topScorer))
	}

	return rows, nil
}

func (t *TopScorerService) head() []string {
	row := []string{}
	row = append(row, "LeagueID")
	row = append(row, "PlayerID")
	row = append(row, "PlayerName")
	row = append(row, "FirstName")
	row = append(row, "LastName")
	row = append(row, "Position")
	row = append(row, "Nationality")
	row = append(row, "TeamID")
	row = append(row, "TeamName")
	row = append(row, "Games.Appearences")
	row = append(row, "Games.MinutesPlayed")
	row = append(row, "Goals.Total")
	row = append(row, "Goals.Assists")
	row = append(row, "Goals.Conceded")
	row = append(row, "Goals.Saves")
	row = append(row, "Shots.Total")
	row = append(row, "Shots.On")
	row = append(row, "Penalty.Won")
	row = append(row, "Penalty.Commited")
	row = append(row, "Cards.Yellow")
	row = append(row, "Cards.SecondYellow")
	row = append(row, "Cards.Red")
	return row

}

func (t *TopScorerService) topScorer(leagueId int, topScorer TopScorer) []string {
	row := []string{}
	row = append(row, strconv.Itoa(leagueId))
	row = append(row, strconv.Itoa(topScorer.PlayerID))
	row = append(row, topScorer.PlayerName)
	row = append(row, topScorer.FirstName)
	row = append(row, topScorer.LastName)
	row = append(row, topScorer.Position)
	row = append(row, topScorer.Nationality)
	row = append(row, strconv.Itoa(topScorer.TeamID))
	row = append(row, topScorer.TeamName)
	row = append(row, strconv.Itoa(topScorer.Games.Appearences))
	row = append(row, strconv.Itoa(topScorer.Games.MinutesPlayed))
	row = append(row, strconv.Itoa(topScorer.Goals.Total))
	row = append(row, strconv.Itoa(topScorer.Goals.Assists))
	row = append(row, strconv.Itoa(topScorer.Goals.Conceded))
	row = append(row, strconv.Itoa(topScorer.Goals.Saves))
	row = append(row, strconv.Itoa(topScorer.Shots.Total))
	row = append(row, strconv.Itoa(topScorer.Shots.On))
	row = append(row, strconv.Itoa(topScorer.Penalty.Won))
	row = append(row, strconv.Itoa(topScorer.Penalty.Commited))
	row = append(row, strconv.Itoa(topScorer.Cards.Yellow))
	row = append(row, strconv.Itoa(topScorer.Cards.SecondYellow))
	row = append(row, strconv.Itoa(topScorer.Cards.Red))
	return row
}
