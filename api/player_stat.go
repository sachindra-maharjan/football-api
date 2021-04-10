package api

import (
	"context"
	"fmt"
	"strconv"
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

func (p *PlayerStatService) Convert(result *PlayerStatResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid league result.")
	}

	rows := [][]string{}

	if includeHead {
		rows = append(rows, p.getHead())
	}

	for _, playerStat := range result.API.PlayerStat {
		rows = append(rows, p.getPlayerStat(result.API.FixtureID, playerStat))
	}

	return rows, nil
}

func (p *PlayerStatService) getPlayerStat(fixtureId int, playerStat PlayerStat) []string {
	row := []string{}
	row = append(row, strconv.Itoa(playerStat.EventID))
	row = append(row, strconv.Itoa(playerStat.UpdatedAt))
	row = append(row, strconv.Itoa(playerStat.PlayerID))
	row = append(row, playerStat.PlayerName)
	row = append(row, strconv.Itoa(playerStat.TeamID))
	row = append(row, playerStat.TeamName)
	row = append(row, strconv.Itoa(playerStat.Number))
	row = append(row, playerStat.Position)
	row = append(row, playerStat.Rating)
	row = append(row, strconv.Itoa(playerStat.MinutesPlayed))
	row = append(row, playerStat.Captain)
	row = append(row, playerStat.Substitute)
	row = append(row, strconv.Itoa(playerStat.Offsides))
	row = append(row, strconv.Itoa(playerStat.Shots.Total))
	row = append(row, strconv.Itoa(playerStat.Shots.On))
	row = append(row, strconv.Itoa(playerStat.Goals.Total))
	row = append(row, strconv.Itoa(playerStat.Goals.Conceded))
	row = append(row, strconv.Itoa(playerStat.Goals.Assists))
	row = append(row, strconv.Itoa(playerStat.Goals.Saves))
	row = append(row, strconv.Itoa(playerStat.Passes.Total))
	row = append(row, strconv.Itoa(playerStat.Passes.Key))
	row = append(row, strconv.Itoa(playerStat.Passes.Accuracy))
	row = append(row, strconv.Itoa(playerStat.Tackles.Total))
	row = append(row, strconv.Itoa(playerStat.Tackles.Blocks))
	row = append(row, strconv.Itoa(playerStat.Tackles.Interceptions))
	row = append(row, strconv.Itoa(playerStat.Duels.Total))
	row = append(row, strconv.Itoa(playerStat.Duels.Won))
	row = append(row, strconv.Itoa(playerStat.Dribbles.Attempts))
	row = append(row, strconv.Itoa(playerStat.Dribbles.Success))
	row = append(row, strconv.Itoa(playerStat.Dribbles.Past))
	row = append(row, strconv.Itoa(playerStat.Fouls.Drawn))
	row = append(row, strconv.Itoa(playerStat.Fouls.Committed))
	row = append(row, strconv.Itoa(playerStat.Cards.Yellow))
	row = append(row, strconv.Itoa(playerStat.Cards.Red))
	row = append(row, strconv.Itoa(playerStat.Penalty.Won))
	row = append(row, strconv.Itoa(playerStat.Penalty.Commited))
	row = append(row, strconv.Itoa(playerStat.Penalty.Success))
	row = append(row, strconv.Itoa(playerStat.Penalty.Missed))
	row = append(row, strconv.Itoa(playerStat.Penalty.Saved))
	return row
}

func (p *PlayerStatService) getHead() []string {
	rows := []string{}
	rows = append(rows, "FixtureID")
	rows = append(rows, "UpdatedAt")
	rows = append(rows, "PlayerID")
	rows = append(rows, "PlayerName")
	rows = append(rows, "TeamID")
	rows = append(rows, "TeamName")
	rows = append(rows, "Number")
	rows = append(rows, "Position")
	rows = append(rows, "Rating")
	rows = append(rows, "MinutesPlayed")
	rows = append(rows, "Caption")
	rows = append(rows, "Substitute")
	rows = append(rows, "Offsides")
	rows = append(rows, "Shots.Total")
	rows = append(rows, "Shots.On")
	rows = append(rows, "Goals.Total")
	rows = append(rows, "Goals.Conceded")
	rows = append(rows, "Goals.Assists")
	rows = append(rows, "Goals.Assists")
	rows = append(rows, "Goals.Saves")
	rows = append(rows, "Passes.Total")
	rows = append(rows, "Passes.Key")
	rows = append(rows, "Passes.Accuracy")
	rows = append(rows, "Tackles.Total")
	rows = append(rows, "Tackles.Blocks")
	rows = append(rows, "Tackles.Interceptions")
	rows = append(rows, "Duels.Total")
	rows = append(rows, "Duels.Won")
	rows = append(rows, "Dribbles.Attempts")
	rows = append(rows, "Dribbles.Success")
	rows = append(rows, "Dribbles.Past")
	rows = append(rows, "Fouls.Drawn")
	rows = append(rows, "Fouls.Committed")
	rows = append(rows, "Cards.Yellow")
	rows = append(rows, "Cards.Red")
	rows = append(rows, "Penalty.Won")
	rows = append(rows, "Penalty.Commited")
	rows = append(rows, "Penalty.Success")
	rows = append(rows, "Penalty.Missed")
	rows = append(rows, "Penalty.Saved")
	return rows
}
