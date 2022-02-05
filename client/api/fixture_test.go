package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFixture_GetFixturesByLeagueID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/fixtures.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}

	mux.HandleFunc("/fixtures/league/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	result, _, err := client.FixtureService.GetFixturesByLeagueID(context.Background(), 2)

	if err != nil {
		t.Fatalf("StandingService_GetLeagueStanding returned error %v", err)
	}

	if result.API.Results != 380 {
		t.Errorf("StandingService_GetLeagueStanding returned error. want %+v got %+v", 380, result.API.Results)
	}
}
