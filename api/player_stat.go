package api

import (
	"context"
	"fmt"
)

//PlayerStatService gets the player statistic for fixtures from api
type PlayerStatService service

//PlayerStatResult contains the parsed result from api response
type PlayerStatResult struct {
	API struct {
		Results    int          `json:"results"`
		PlayerStat []PlayerStat `json:"players"`
		FixtureID  int          `json:"fixtureID,omitempty"`
	} `json:"api"`
}

//PlayerStat contains player statistics for a fixture
type PlayerStat struct {
	EventID       int    `json:"event_id,omitempty"`
	UpdatedAt     int    `json:"updateAt,omitempty"`
	PlayerID      int    `json:"player_id,omitempty"`
	PlayerName    string `json:"player_name,omitempty"`
	TeamID        int    `json:"team_id,omitempty"`
	TeamName      string `json:"team_name,omitempty"`
	Number        int    `json:"number,omitempty"`
	Position      string `json:"position,omitempty"`
	Rating        string `json:"rating,omitempty"`
	MinutesPlayed int    `json:"minutes_played,omitempty"`
	Captain       string `json:"captain,omitempty"`
	Substitute    string `json:"substitute,omitempty"`
	Offsides      int    `json:"offsides,omitempty"`
	Shots         struct {
		Total int `json:"total,omitempty"`
		On    int `json:"on,omitempty"`
	} `json:"shots"`
	Goals struct {
		Total    int `json:"total,omitempty"`
		Conceded int `json:"conceded,omitempty"`
		Assists  int `json:"assists,omitempty"`
		Saves    int `json:"saves,omitempty"`
	} `json:"goals"`
	Passes struct {
		Total    int `json:"total,omitempty"`
		Key      int `json:"key,omitempty"`
		Accuracy int `json:"accuracy,omitempty"`
	} `json:"passes,omitempty"`
	Tackles struct {
		Total         int `json:"total,omitempty"`
		Blocks        int `json:"blocks,omitempty"`
		Interceptions int `json:"interceptions,omitempty"`
	} `json:"tackles,omitempty"`
	Duels struct {
		Total int `json:"total,omitempty"`
		Won   int `json:"won,omitempty"`
	} `json:"duels,omitempty"`
	Dribbles struct {
		Attempts int `json:"attempts,omitempty"`
		Success  int `json:"success,omitempty"`
		Past     int `json:"past,omitempty"`
	} `json:"dribbles,omitempty"`
	Fouls struct {
		Drawn     int `json:"drawn,omitempty"`
		Committed int `json:"committed,omitempty"`
	} `json:"fouls,omitempty"`
	Cards struct {
		Yellow int `json:"yellow,omitempty"`
		Red    int `json:"red,omitempty"`
	} `json:"cards,omitempty"`
	Penalty struct {
		Won      int `json:"won,omitempty"`
		Commited int `json:"commited,omitempty"`
		Success  int `json:"success,omitempty"`
		Missed   int `json:"missed,omitempty"`
		Saved    int `json:"saved,omitempty"`
	} `json:"penalty,omitempty"`
}

func (p *PlayerStatService) GetPlayerStatByFixtureID(context context.Context, fixtureID int) (*PlayerStatResult, *Response, error) {
	r, err := p.client.NewRequest("GET", "players/fixture/"+fmt.Sprint(fixtureID), nil)
	if err != nil {
		return nil, nil, err
	}

	var playerStatResult *PlayerStatResult
	response, err := p.client.Do(context, r, &playerStatResult)

	if err != nil {
		return nil, nil, err
	}

	playerStatResult.API.FixtureID = fixtureID
	return playerStatResult, response, nil

}
