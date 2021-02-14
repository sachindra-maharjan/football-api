package api

import (
	"context"
	"fmt"
)

//FixtureEventService gets the events of fixtures from api
type FixtureEventService service

//FixtureEventResult contains the parsed result from api response
type FixtureEventResult struct {
	API struct {
		Results       int            `json:"results"`
		FixtureEvents []FixtureEvent `json:"events"`
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

	return fixtureEventResult, resp, nil
}
