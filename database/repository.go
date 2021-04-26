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

	common         dbservice
	LeagueService  *LeagueService
	FixtureService *FixtureService
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

func parseDate(val string) time.Time {
	mydate, err := time.Parse("2018-08-10 19:00:00", val)
	if err != nil {
		return time.Now().UTC()
	}
	return mydate.UTC()
}
