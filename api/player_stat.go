package api

//PlayerStatService gets the player statistic for fixtures from api
type PlayerStatService service

//PlayerStatResult contains the parsed result from api response
type PlayerStatResult struct {
	API struct {
		Results    int          `json:"results"`
		PlayerStat []PlayerStat `json:"players"`
	} `json:"api"`
}

//PlayerStat contains player statistics for a fixture
type PlayerStat struct {
	EventID       int    `json:"event_id"`
	UpdatedAt     int    `json:"updateAt"`
	PlayerID      int    `json:"player_id"`
	PlayerName    string `json:"player_name"`
	TeamID        int    `json:"team_id"`
	Number        int    `json:"number"`
	Position      int    `json:"position"`
	Rating        string `json:"rating,omitempty"`
	MinutesPlayed int    `json:"minutes_played"`
	Captain       string `json:"captain"`
	Substitute    string `json:"substitute"`
	Offsides      int    `json:"offsides,omitempty"`
	Shots         struct {
		Total int `json:"total"`
		On    int `json:"on"`
	} `json:"shots"`
	Goals struct {
		Total    int `json:"total"`
		Conceded int `json:"conceded"`
		Assists  int `json:"assists"`
		Saves    int `json:"saves"`
	} `json:"goals"`
	Passes struct {
		Total    int `json:"total"`
		Key      int `json:"key"`
		Accuracy int `json:"accuracy"`
	} `json:"passes"`
	Tackles struct {
		Total         int `json:"total"`
		Blocks        int `json:"blocks"`
		Interceptions int `json:"interceptions"`
	} `json:"tackles"`
	Duels struct {
		Total int `json:"total"`
		Won   int `json:"won"`
	} `json:"duels"`
	Dribbles struct {
		Attempts int `json:"attempts"`
		Success  int `json:"success"`
		Past     int `json:"past"`
	} `json:"dribbles"`
	Fouls struct {
		Drawn     int `json:"drawn"`
		Committed int `json:"committed"`
	} `json:"fouls"`
	Cards struct {
		Yellow string `json:"yellow"`
		Red    string `json:"red"`
	} `json:"cards"`
	Penalty struct {
		Won      int `json:"won"`
		Commited int `json:"commited"`
		Success  int `json:"success"`
		Missed   int `json:"missed"`
		Saved    int `json:"saved"`
	} `json:"penalty"`
}
