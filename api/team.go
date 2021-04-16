package api

import (
	"context"
	"fmt"
	"strconv"
)

//TeamService gets team informtion from api call
type TeamService service

//TeamResult contins the parsed result of api response of team
type TeamResult struct {
	API struct {
		Results  int    `json:"results,omitempty"`
		Teams    []Team `json:"teams,omitempty"`
		LeagueID int    `json:"leagueID,omitempty"`
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
	teamResult.API.LeagueID = leagueID
	return teamResult, res, nil
}

//GetFlatDataWithHeader Returns flat data with header
func (l *TeamService) Convert(teamResult *TeamResult, includeHead bool) ([][]string, error) {
	if teamResult == nil {
		return nil, fmt.Errorf("invalid team result.")
	}

	rows := [][]string{}
	if includeHead {
		rows = append(rows, l.getHead())
	}

	for _, team := range teamResult.API.Teams {
		rows = append(rows, l.getData(team, teamResult.API.LeagueID))
	}

	return rows, nil
}

//GetHead Returns the array of head fields
func (service *TeamService) getHead() []string {
	var row []string
	row = append(row, "league_id")
	row = append(row, "team_id")
	row = append(row, "name")
	row = append(row, "code")
	row = append(row, "country")
	row = append(row, "is_national")
	row = append(row, "founded")
	row = append(row, "venuename")
	row = append(row, "venuesurface")
	row = append(row, "venueaddress")
	row = append(row, "venuecity")
	row = append(row, "venuecapacity")
	return row
}

//GetFlat Returns flat array from an object
func (service *TeamService) getData(t Team, leagueID int) []string {
	var row []string
	row = append(row, strconv.Itoa(leagueID))
	row = append(row, strconv.Itoa(t.TeamID))
	row = append(row, t.Name)
	row = append(row, t.Code)
	row = append(row, t.Country)
	row = append(row, strconv.FormatBool(t.IsNational))
	row = append(row, strconv.Itoa(t.Founded))
	row = append(row, t.VenueName)
	row = append(row, t.VenueSurface)
	row = append(row, t.VenueAddress)
	row = append(row, t.VenueCity)
	row = append(row, strconv.Itoa(t.VenueCapacity))
	return row
}
