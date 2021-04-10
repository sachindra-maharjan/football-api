package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
		testHeader(t, r, "x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		w.Header().Add("x-ratelimit-requests-limit", "100")
		w.Header().Add("x-ratelimit-requests-remaining", "100")
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

func TestFixtureEvent_GetAllFixtureEventsWithRateLimits(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/fixture_events.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}
	limit := 100
	remaining := 100
	index := 0

	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		testAuthKey(t, r, index, apiAuthKeys[index])
		w.Header().Add("x-ratelimit-requests-limit", strconv.Itoa(limit))
		w.Header().Add("x-ratelimit-requests-remaining", strconv.Itoa(remaining))
		fmt.Fprint(w, string(data))
	})

	for i := 0; i < 501; i++ {
		_, _, err = client.FixtureEventService.GetFixtureEvent(context.Background(), 65)
		remaining--
		if remaining < 0 {
			remaining = 100
			index++
		}
	}
}
