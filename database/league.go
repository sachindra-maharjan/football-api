package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

type LeagueService dbservice

type League struct {
	LeagueID    int    `firestore:"leagueId,omitempty"`
	Name        string `firestore:"name,omitempty"`
	LeagueType  string `firestore:"type,omitempty"`
	Country     string `firestore:"country,omitempty"`
	CountryCode string `firestore:"countryCode,omitempty"`
	LogoURL     string `firestore:"logo,omitempty"`
}

type Season struct {
	LeagueID    int       `firestore:"leagueId,omitempty"`
	Season      string    `firestore:"season,omitempty"`
	SeasonStart time.Time `firestore:"seasonStart,omitempty"`
	SeasonEnd   time.Time `firestore:"seasonEnd,omitempty"`
	FlagURL     string    `firestore:"flag,omitempty"`
	Standings   bool      `firestore:"standings,omitempty"`
	IsCurrent   bool      `firestore:"isCurrent,omitempty"`
	Coverage    struct {
		Standings       bool `firestore:"standings,omitempty"`
		FixtureCoverage struct {
			Events           bool `firestore:"events,omitempty"`
			Lineups          bool `firestore:"lineups,omitempty"`
			Statistics       bool `firestore:"statistics,omitempty"`
			PlayerStatistics bool `firestore:"playersStatistics,omitempty"`
		} `firestore:"fixtures,omitempty"`
		Players     bool `firestore:"players,omitempty"`
		TopScorers  bool `firestore:"topScorers,omitempty"`
		Predictions bool `firestore:"predictions,omitempty"`
		Odds        bool `firestore:"odds,omitempty"`
	} `firestore:"coverage,omitempty"`
}

func (s *LeagueService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new data to firestore \n", len(records))

	batch := s.client.fs.Batch()
	var leagueRef *firestore.DocumentRef

	for i, r := range records {

		if i == 0 {
			leagueName := strings.ToLower(strings.ReplaceAll(r[1], " ", ""))
			leagueRef = s.client.fs.Collection("football").Doc(leagueName)
			l := League{}
			l.LeagueID = parseInt(r[0])
			l.Name = r[1]
			l.LeagueType = r[2]
			l.CountryCode = r[3]
			l.Country = r[4]
			l.LogoURL = r[8]
			batch.Set(leagueRef, l)
		}

		season := Season{}
		season.LeagueID = parseInt(r[0])
		season.Season = r[5]
		season.SeasonStart = parseDate("2006-01-02", r[6])
		season.SeasonEnd = parseDate("2006-01-02", r[7])
		season.FlagURL = r[9]
		season.IsCurrent = parseBool(r[10])
		season.Coverage.FixtureCoverage.Events = parseBool(r[11])
		season.Coverage.FixtureCoverage.Lineups = parseBool(r[12])
		season.Coverage.FixtureCoverage.PlayerStatistics = parseBool(r[13])
		season.Coverage.FixtureCoverage.Statistics = parseBool(r[14])
		season.Coverage.Standings = parseBool(r[15])
		season.Coverage.Players = parseBool(r[16])
		season.Coverage.Predictions = parseBool(r[17])
		season.Coverage.Odds = parseBool(r[18])
		season.Coverage.TopScorers = parseBool(r[19])

		seasonRef := leagueRef.Collection("leagues").Doc("leagueId_" + r[0])
		batch.Set(seasonRef, season)

	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return err
}
