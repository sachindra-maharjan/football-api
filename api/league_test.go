package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLeague_ListAll(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/seasons.json")
	if err != nil {
		t.Fatalf("Json file could not be read.")
	}

	mux.HandleFunc("/leagues/seasonsAvailable/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	leagueResult, _, err := client.LeagueService.ListAll(context.Background(), 2)
	if err != nil {
		t.Errorf("LeagueService.ListAll returned error: %v", err)
	}

	if leagueResult.API.Results != 11 {
		t.Errorf("LeagueService.ListAll returned error %v", err)
	}
}
