package database

import (
	"context"
	"fmt"
)

type TopScorerService dbservice

//TopScorer contains the top scorer player information
type TopScorer struct {
	LeagueID    int    `firestore:"leagueId,omitempty"`
	PlayerID    int    `firestore:"playerId,omitempty"`
	PlayerName  string `firestore:"playerName,omitempty"`
	FirstName   string `firestore:"firstname,omitempty"`
	LastName    string `firestore:"lastname,omitempty"`
	Position    string `firestore:"position,omitempty"`
	Nationality string `firestore:"nationality,omitempty"`
	TeamID      int    `firestore:"teamId,omitempty"`
	TeamName    string `firestore:"teamName,omitempty"`
	Games       struct {
		Appearences   int `firestore:"appearences,omitempty"`
		MinutesPlayed int `firestore:"minutesPlayed,omitempty"`
	} `firestore:"games"`
	Goals struct {
		Total    int `firestore:"total,omitempty"`
		Assists  int `firestore:"assists,omitempty"`
		Conceded int `firestore:"conceded,omitempty"`
		Saves    int `firestore:"saves,omitempty"`
	} `firestore:"goals"`
	Shots struct {
		Total int `firestore:"total,omitempty"`
		On    int `firestore:"on,omitempty"`
	} `firestore:"shots"`
	Penalty struct {
		Won      int `firestore:"won,omitempty"`
		Commited int `firestore:"commited,omitempty"`
	} `firestore:"penalty"`
	Cards struct {
		Yellow       int `firestore:"yello,omitempty"`
		SecondYellow int `firestore:"secondYellow,omitempty"`
		Red          int `firestore:"red,omitempty"`
	}
}

func (s *TopScorerService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new top scorer data to firestore \n", len(records))

	batch := s.client.fs.Batch()

	for _, r := range records {

		t := TopScorer{}
		t.LeagueID = parseInt(r[0])
		t.PlayerID = parseInt(r[1])
		t.PlayerName = r[2]
		t.FirstName = r[3]
		t.LastName = r[4]
		t.Position = r[5]
		t.Nationality = r[6]
		t.TeamID = parseInt(r[7])
		t.TeamName = r[8]
		t.Games.Appearences = parseInt(r[9])
		t.Games.MinutesPlayed = parseInt(r[10])
		t.Goals.Total = parseInt(r[11])
		t.Goals.Assists = parseInt(r[12])
		t.Goals.Conceded = parseInt(r[13])
		t.Goals.Assists = parseInt(r[14])
		t.Shots.Total = parseInt(r[15])
		t.Shots.On = parseInt(r[16])
		t.Penalty.Won = parseInt(r[17])
		t.Penalty.Commited = parseInt(r[18])
		t.Cards.Yellow = parseInt(r[19])
		t.Cards.SecondYellow = parseInt(r[20])
		t.Cards.Red = parseInt(r[21])

		leagueRef := s.client.fs.Collection("football").Doc(leagueName)

		docRef := leagueRef.
			Collection("leagues").
			Doc("leagueId_" + r[0]).
			Collection("topScorers").
			Doc(DocWithIDAndName(r[1], r[2]))
		batch.Set(docRef, t)

	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}
