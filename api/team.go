package api

import (
	"context"
	"fmt"
)

//TeamService gets team informtion from api call
type TeamService service

//TeamResult contins the parsed result of api response of team
type TeamResult struct {
	API struct {
		Results int    `json:"results"`
		Teams   []Team `json:"teams"`
	} `json:"api"`
}

//Team contains information about a team
type Team struct {
	TeamID        int    `json:"team_id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Country       string `json:"country"`
	IsNational    bool   `json:"is_national"`
	Founded       int    `json:"founded"`
	VenueName     string `json:"veneu_name"`
	VenueSurface  string `json:"venue_surface"`
	VenueAddress  string `json:"venue_address"`
	VenueCity     string `json:"venue_city"`
	VenueCapacity int    `json:"venue_capacity"`
}

//ListTeamsByLeagueID gets all the teams in the league
func (t *TeamService) ListTeamsByLeagueID(ctx context.Context, leagueID int) (*TeamResult, *Response, error) {
	req, err := t.client.NewRequest("GET", "teams/league/"+fmt.Sprint(leagueID), nil)
	if err != nil {
		return nil, nil, err
	}

	var teamResult *TeamResult
	res, err := t.client.Do(ctx, req, &teamResult)
	if err != nil {
		return nil, nil, err
	}

	return teamResult, res, nil
}
