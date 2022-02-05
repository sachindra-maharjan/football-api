package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
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

	result, _, err := client.FixtureEventService.GetFixtureEvent(context.Background(), 2, 65)

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
	remaining := 2
	index := 0
	maxReqPerMin := 30

	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		testAuthKey(t, r, index, apiAuthKeys[index])
		w.Header().Add("x-ratelimit-requests-limit", strconv.Itoa(limit))
		w.Header().Add("x-ratelimit-requests-remaining", strconv.Itoa(remaining))
		fmt.Fprint(w, string(data))
	})

	requestCount := 0
	var startTime time.Time
	waitFlag := true
	for i := 0; i < 380; i++ {
		if waitFlag {
			startTime = time.Now()
		}

		requestCount++
		waitFlag = wait(startTime, requestCount, maxReqPerMin)
		if waitFlag {
			requestCount = 0
		}
		_, _, err = client.FixtureEventService.GetFixtureEvent(context.Background(), 2, 65)

		if err == nil {
			remaining--
			if remaining < 0 {
				remaining = 100
				index++
			}
		}
	}
}

func wait(startTime time.Time, reqCount int, maxReqPerMin int) bool {
	waitFlag := false
	if reqCount == maxReqPerMin {
		elapsed := time.Now().Sub(startTime)
		if elapsed.Milliseconds() <= 60*1000 {
			fmt.Printf("Request limit per minute exceeded.Waiting for %d s before new request.\n", 61)
			time.Sleep(time.Duration(61) * time.Second)
			waitFlag = true
		}
	}
	return waitFlag
}
