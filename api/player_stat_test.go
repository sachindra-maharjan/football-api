package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPlayerStat_GetPlayerStatByFixtureID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	data, err := ioutil.ReadFile("data/playerstat_by_fixture.json")
	if err != nil {
		t.Fatalf("Json file could not be read. Error: %v", err)
	}

	mux.HandleFunc("/players/fixture/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(data))
	})

	result, _, err := client.PlayerStatService.GetPlayerStatByFixtureID(context.Background(), 2, 169080)

	if err != nil {
		t.Fatalf("PlayerStatService_GetPlayerStatByFixtureID returned error %v", err)
	}

	if result.API.Results != 27 {
		t.Errorf("PlayerStatService_GetPlayerStatByFixtureID returned error. want %+v got %+v", 27, result.API.Results)
	}
}
