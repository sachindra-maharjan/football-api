package api

import "time"

//PlayerService gets the player information from api call
type PlayerService service

//PlayerResult contins the parsed result of api response of players
type PlayerResult struct {
	API struct {
		Results int      `json:"results"`
		Players []Player `json:"players"`
	} `json:"api"`
}

//BirthDate type
type BirthDate time.Time

//Player contains information about league
type Player struct {
	Name         string    `json:"player_name"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	Number       int       `json:"number,omitempty"`
	Position     string    `json:"position"`
	Age          int       `json:"age"`
	BirthDate    BirthDate `json:"birth_date"`
	BirthPlace   string    `json:"birth_place"`
	BirthCountry string    `json:"birth_country"`
	Nationality  string    `json:"nationality"`
	Height       string    `json:"height"`
	Weight       string    `json:"weight"`
}
