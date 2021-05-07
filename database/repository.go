package database

import (
	"context"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
)

type dbservice struct {
	client *FSClient
}

type FSClient struct {
	fs *firestore.Client

	common                   dbservice
	LeagueService            *LeagueService
	FixtureService           *FixtureService
	StandingsService         *StandingsService
	FixtureEventService      *FixtureEventService
	FixtureLineUpService     *FixtureLineUpService
	FixturePlayerStatService *FixturePlayerStatService
	TopScorerService         *TopScorerService
	TeamService              *TeamService
}

func NewClient(ctx context.Context, projectId string) (*FSClient, error) {

	firestore, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		return nil, err
	}

	fsc := &FSClient{
		fs: firestore,
	}

	fsc.common.client = fsc
	fsc.LeagueService = (*LeagueService)(&fsc.common)
	fsc.FixtureService = (*FixtureService)(&fsc.common)
	fsc.StandingsService = (*StandingsService)(&fsc.common)
	fsc.FixtureEventService = (*FixtureEventService)(&fsc.common)
	fsc.FixtureLineUpService = (*FixtureLineUpService)(&fsc.common)
	fsc.FixturePlayerStatService = (*FixturePlayerStatService)(&fsc.common)
	fsc.TopScorerService = (*TopScorerService)(&fsc.common)
	fsc.TeamService = (*TeamService)(&fsc.common)

	return fsc, nil
}

func parseBool(val string) bool {
	flag, err := strconv.ParseBool(val)
	if err != nil {
		flag = false
	}
	return flag
}

func parseInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return result
}

func parseInt64(val string) int64 {
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func parseDate(format string, val string) time.Time {
	mydate, err := time.Parse(format, val)
	if err != nil {
		return time.Now().UTC()
	}
	return mydate.UTC()
}
