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
		Results int    `json:"results,omitempty"`
		Teams   []Team `json:"teams,omitempty"`
	} `json:"api"`
}

//Team contains information about a team
type Team struct {
	TeamID        int    `json:"team_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Code          string `json:"code,omitempty"`
	Country       string `json:"country,omitempty"`
	IsNational    bool   `json:"is_national,omitempty"`
	Founded       int    `json:"founded,omitempty"`
	VenueName     string `json:"veneu_name,omitempty"`
	VenueSurface  string `json:"venue_surface,omitempty"`
	VenueAddress  string `json:"venue_address,omitempty"`
	VenueCity     string `json:"venue_city,omitempty"`
	VenueCapacity int    `json:"venue_capacity,omitempty"`
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
