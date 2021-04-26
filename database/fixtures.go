package database

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type FixtureService dbservice

//Fixture contains the fixture infromation
type Fixture struct {
	FixtureID int `firestore:"fixture_id,omitempty"`
	LeagueID  int `firestore:"league_id,omitempty"`
	League    struct {
		Name    string `firestore:"name,omitempty"`
		Country string `firestore:"country,omitempty"`
	} `firestore:"league"`
	EventDate       time.Time `firestore:"event_date,omitempty"`
	EventTimestamp  int64     `firestore:"event_timestamp,omitempty"`
	FirstHalfStart  int64     `firestore:"first_half_start,omitempty"`
	SecondHalfStart int64     `firestore:"second_half_start,omitempty"`
	Round           string    `firestore:"round,omitempty"`
	Status          string    `firestore:"status,omitempty"`
	StatusShort     string    `firestore:"status_short,omitempty"`
	Elapsed         int       `firestore:"elapsed,omitempty"`
	Venue           string    `firestore:"venue,omitempty"`
	Referee         string    `firestore:"referee,omitempty,omitempty"`
	HomeTeam        team      `firestore:"home_team,omitempty"`
	AwayTeam        team      `firestore:"away_team,omitempty"`
	GoalsHomeTeam   int       `firestore:"goals_home_team,omitempty"`
	GoalsAwayTeam   int       `firestore:"goals_away_team,omitempty"`
	Score           struct {
		HalfTime  string `firestore:"halftime,omitempty"`
		FullTime  string `firestore:"fulltime,omitempty"`
		ExtraTime string `firestore:"extratime,omitempty,omitempty"`
		Penalty   string `firestore:"penalty,omitempty,omitempty"`
	}
}

type team struct {
	TeamID   int    `firestore:"team_id,omitempty"`
	TeamName string `firestore:"team_name,omitempty"`
	Logo     string `firestore:"logo,omitempty"`
}

func (s *FixtureService) Add(ctx context.Context, records [][]string) error {
	fmt.Printf("Adding %d new data to firestore \n", len(records))

	batch := s.client.fs.Batch()

	for _, r := range records {

		f := Fixture{}
		f.LeagueID = parseInt(r[0])
		f.FixtureID = parseInt(r[1])
		f.League.Name = r[2]
		f.League.Country = r[3]
		f.EventDate = parseDate(r[6])
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

		docId := strings.ToLower(strings.ReplaceAll(f.League.Name, " ", "") + "#" + r[0] + "#" + r[1])
		docRef := s.client.fs.Collection("fixtures").Doc(docId)
		batch.Set(docRef, f)

	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return err
}
