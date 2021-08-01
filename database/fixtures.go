package database

import (
	"context"
	"fmt"
	"time"
)

type FixtureService dbservice

//Fixture contains the fixture infromation
type Fixture struct {
	FixtureID int `firestore:"fixtureId,omitempty"`
	LeagueID  int `firestore:"leagueId,omitempty"`
	League    struct {
		Name    string `firestore:"name,omitempty"`
		Country string `firestore:"country,omitempty"`
	} `firestore:"league"`
	EventDate       time.Time `firestore:"eventDate,omitempty"`
	EventTimestamp  int64     `firestore:"eventTimestamp,omitempty"`
	FirstHalfStart  int64     `firestore:"firstHalfStart,omitempty"`
	SecondHalfStart int64     `firestore:"secondHalfStart,omitempty"`
	Round           string    `firestore:"round,omitempty"`
	Status          string    `firestore:"status,omitempty"`
	StatusShort     string    `firestore:"statusShort,omitempty"`
	Elapsed         int       `firestore:"elapsed,omitempty"`
	Venue           string    `firestore:"venue,omitempty"`
	Referee         string    `firestore:"referee,omitempty,omitempty"`
	HomeTeam        team      `firestore:"homeTeam,omitempty"`
	AwayTeam        team      `firestore:"awayTeam,omitempty"`
	GoalsHomeTeam   int       `firestore:"goalsHomeTeam,omitempty"`
	GoalsAwayTeam   int       `firestore:"goalsAwayTeam,omitempty"`
	Score           struct {
		HalfTime  string `firestore:"halftime,omitempty"`
		FullTime  string `firestore:"fulltime,omitempty"`
		ExtraTime string `firestore:"extratime,omitempty,omitempty"`
		Penalty   string `firestore:"penalty,omitempty,omitempty"`
	}
}

type team struct {
	TeamID   int    `firestore:"teamId,omitempty"`
	TeamName string `firestore:"teamName,omitempty"`
	Logo     string `firestore:"logo,omitempty"`
}

func (s *FixtureService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new data to firestore \n", len(records))

	batch := s.client.fs.Batch()

	for _, r := range records {

		f := Fixture{}
		f.LeagueID = parseInt(r[0])
		f.FixtureID = parseInt(r[1])
		f.League.Name = r[2]
		f.League.Country = r[3]
		f.EventDate = parseDate(time.RFC3339, r[6])
		f.EventTimestamp = int64(parseInt(r[7]))
		f.FirstHalfStart = int64(parseInt(r[8]))
		f.SecondHalfStart = int64(parseInt(r[9]))
		f.Round = r[10]
		f.Status = r[11]
		f.StatusShort = r[12]
		f.Elapsed = parseInt(r[13])
		f.Venue = r[14]
		f.Referee = r[15]
		f.HomeTeam.TeamID = parseInt(r[16])
		f.HomeTeam.TeamName = r[17]
		f.HomeTeam.Logo = r[18]
		f.AwayTeam.TeamID = parseInt(r[19])
		f.AwayTeam.TeamName = r[20]
		f.AwayTeam.Logo = r[21]
		f.GoalsHomeTeam = parseInt(r[22])
		f.GoalsAwayTeam = parseInt(r[23])
		f.Score.HalfTime = r[24]
		f.Score.FullTime = r[25]
		f.Score.ExtraTime = r[26]
		f.Score.Penalty = r[27]

		leagueRef := s.client.fs.Collection("football").Doc(leagueName)

		docRef := leagueRef.
			Collection("leagues").
			Doc("leagueId_" + r[0]).
			Collection("fixtures").
			Doc("fixtureId_" + r[1])
		batch.Set(docRef, f)

	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return err
}
