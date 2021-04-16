package api

import (
	"context"
	"fmt"
	"strconv"
)

//FixtureEventService gets the events of fixtures from api
type FixtureEventService service

//FixtureEventResult contains the parsed result from api response
type FixtureEventResult struct {
	API struct {
		Results       int            `json:"results"`
		FixtureEvents []FixtureEvent `json:"events"`
		FixtureID     int            `json:"fixtureID,omitempty"`
	} `json:"api"`
}

//FixtureEvent contains the fixture event infromation
type FixtureEvent struct {
	Elapsed        int    `json:"elapsed"`
	ElapsedPlus    int    `json:"elapsed_plus,omitempty"`
	TeamID         int    `json:"team_id"`
	TeamName       string `json:"teamName"`
	PlayerID       int    `json:"player_id"`
	Player         string `json:"player"`
	AssistPlayerID int    `json:"assist_id"`
	AssistedBy     string `json:"assist"`
	Type           string `json:"type"`
	Detail         string `json:"detail"`
	Comments       string `json:"comments,omitempty"`
}

//GetFixtureEvent Returns events of a fixture
func (fe *FixtureEventService) GetFixtureEvent(context context.Context, fixtureID int) (*FixtureEventResult, *Response, error) {
	req, err := fe.client.NewRequest("GET", "events/"+fmt.Sprint(fixtureID), nil)
	if err != nil {
		return nil, nil, err
	}

	var fixtureEventResult *FixtureEventResult
	resp, err := fe.client.Do(context, req, &fixtureEventResult)
	if err != nil {
		return nil, nil, err
	}

	fixtureEventResult.API.FixtureID = fixtureID
	return fixtureEventResult, resp, nil
}

//Converts result into a flat data
func (service *FixtureEventService) Convert(result *FixtureEventResult, includeHead bool) ([][]string, error) {
	if result == nil {
		return nil, fmt.Errorf("invalid standing data")
	}

	var rows [][]string

	if includeHead {
		rows = append(rows, service.getHeader())
	}

	for _, event := range result.API.FixtureEvents {
		rows = append(rows, service.getData(event, result.API.FixtureID))
	}

	return rows, nil
}

func (service *FixtureEventService) getHeader() []string {
	var row []string
	row = append(row, "fixture_id")
	row = append(row, "elapsed")
	row = append(row, "elapsed_plus")
	row = append(row, "team_id")
	row = append(row, "team_name")
	row = append(row, "player_id")
	row = append(row, "player")
	row = append(row, "assist_player_id")
	row = append(row, "assisted_by")
	row = append(row, "type")
	row = append(row, "detail")
	row = append(row, "comments")
	return row
}

func (service *FixtureEventService) getData(f FixtureEvent, fixtureID int) []string {
	var row []string
	row = append(row, strconv.Itoa(fixtureID))
	row = append(row, strconv.Itoa(f.Elapsed))
	row = append(row, strconv.Itoa(f.ElapsedPlus))
	row = append(row, strconv.Itoa(f.TeamID))
	row = append(row, f.TeamName)
	row = append(row, strconv.Itoa(f.TeamID))
	row = append(row, f.Player)
	row = append(row, strconv.Itoa(f.AssistPlayerID))
	row = append(row, f.AssistedBy)
	row = append(row, f.Type)
	row = append(row, f.Detail)
	row = append(row, f.Comments)
	return row
}
