package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFixtureLineUp_GetLineUpForFixture(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/fixture_lineups.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}

	mux.HandleFunc("/lineups/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	result, _, err := client.FixtureLineUpService.GetLineUpForFixture(context.Background(), 2, 65)

	if err != nil {
		t.Fatalf("StandingService_GetLeagueStanding returned error %v", err)
	}

	if result.API.Results != 2 {
		t.Errorf("StandingService_GetLeagueStanding returned error. want %+v got %+v", 2, result.API.Results)
	}

	fmt.Println(result.API.LineUp["Manchester United"])

}
