package api

//TopScorerService gets the top scorer players list from api
type TopScorerService service

//TopScorerResult contains the parsed result from api response
type TopScorerResult struct {
	API struct {
		Results   int         `json:"results"`
		TopScorer []TopScorer `json:"topscorers"`
	} `json:"api"`
}

//TopScorer contains the top scorer player information
type TopScorer struct {
	PlayerID    int    `json:"player_id"`
	PlayerName  string `json:"player_name"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Position    string `json:"position"`
	Nationality string `json:"nationality,omitempty"`
	TeamID      int    `json:"team_id"`
	TeamName    string `json:"team_name"`
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
