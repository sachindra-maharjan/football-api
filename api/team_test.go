package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTeamService_ListTeamsByLeagueID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/team.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}

	mux.HandleFunc("/teams/league/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	result, _, err := client.TeamService.ListTeamsByLeagueID(context.Background(), 2)

	if err != nil {
		t.Fatalf("StandingService_GetLeagueStanding returned error %v", err)
	}

	if result.API.Results != 20 {
		t.Errorf("StandingService_GetLeagueStanding returned error. want %+v got %+v", 20, result.API.Results)
	}
}
