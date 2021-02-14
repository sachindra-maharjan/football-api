package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFixtureEvent_GetFixtureEvents(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/fixture_events.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}

	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	result, _, err := client.FixtureEventService.GetFixtureEvent(context.Background(), 65)

	if err != nil {
		t.Fatalf("StandingService_GetLeagueStanding returned error %v", err)
	}

	if result.API.Results != 12 {
		t.Errorf("StandingService_GetLeagueStanding returned error. want %+v got %+v", 12, result.API.Results)
	}

}
