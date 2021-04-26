package database

import (
	"context"
	"fmt"
	"strings"
)

type LeagueService dbservice

type League struct {
	LeagueID    int    `firestore:"league_id,omitempty"`
	Name        string `firestore:"name,omitempty"`
	LeagueType  string `firestore:"type,omitempty"`
	Country     string `firestore:"country,omitempty"`
	CountryCode string `firestore:"country_code,omitempty"`
	Season      string `firestore:"season,omitempty"`
	SeasonStart string `firestore:"season_start,omitempty"`
	SeasonEnd   string `firestore:"season_end,omitempty"`
	LogoURL     string `firestore:"logo,omitempty"`
	FlagURL     string `firestore:"flag,omitempty"`
	Standings   bool   `firestore:"standings,omitempty"`
	IsCurrent   bool   `firestore:"is_current,omitempty"`
	Coverage    struct {
		Standings       bool `firestore:"standings,omitempty"`
		FixtureCoverage struct {
			Events           bool `firestore:"events,omitempty"`
			Lineups          bool `firestore:"lineups,omitempty"`
			Statistics       bool `firestore:"statistics,omitempty"`
			PlayerStatistics bool `firestore:"players_statistics,omitempty"`
		} `firestore:"fixtures,omitempty"`
		Players     bool `firestore:"players,omitempty"`
		TopScorers  bool `firestore:"topScorers,omitempty"`
		Predictions bool `firestore:"predictions,omitempty"`
		Odds        bool `firestore:"odds,omitempty"`
	} `firestore:"coverage,omitempty"`
}

func (s *LeagueService) Add(ctx context.Context, records [][]string) error {
	fmt.Printf("Adding %d new data to firestore \n", len(records))

	batch := s.client.fs.Batch()

	for _, r := range records {

		l := League{}
		l.LeagueID = parseInt(r[0])
		l.Name = r[1]
		l.LeagueType = r[2]
		l.CountryCode = r[3]
		l.Country = r[4]
		l.Season = r[5]
		l.SeasonStart = r[6]
		l.SeasonEnd = r[7]
		l.LogoURL = r[8]
		l.FlagURL = r[9]
		l.IsCurrent = parseBool(r[10])
		l.Coverage.FixtureCoverage.Events = parseBool(r[11])
		l.Coverage.FixtureCoverage.Lineups = parseBool(r[12])
		l.Coverage.FixtureCoverage.PlayerStatistics = parseBool(r[13])
		l.Coverage.FixtureCoverage.Statistics = parseBool(r[14])
		l.Coverage.Standings = parseBool(r[15])
		l.Coverage.Players = parseBool(r[16])
		l.Coverage.Predictions = parseBool(r[17])
		l.Coverage.Odds = parseBool(r[18])
		l.Coverage.TopScorers = parseBool(r[19])

		docId := strings.ToLower(strings.ReplaceAll(r[1], " ", "") + "#" + r[0])
		docRef := s.client.fs.Collection("leagues").Doc(docId)
		batch.Set(docRef, l)

	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return err
}
