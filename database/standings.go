package database

import (
	"context"
	"fmt"
)

type StandingsService dbservice

//Standings contains league table stangings
type Standings struct {
	LeagueID    int    `firestore:"leagueId,omitempty"`
	Rank        int    `firestore:"rank,omitempty"`
	TeamID      int    `firestore:"teamId,omitempty"`
	TeamName    string `firestore:"teamName,omitempty"`
	Logo        string `firestore:"logo,omitempty"`
	Group       string `firestore:"group,omitempty"`
	Status      string `firestore:"status,omitempty"`
	Form        string `firestore:"forme,omitempty"`
	Description string `firestore:"description,omitempty"`
	AllStat     Stat   `firestore:"all,omitempty"`
	HomeStat    Stat   `firestore:"home,omitempty"`
	AwayStat    Stat   `firestore:"away,omitempty"`
	GoalsDiff   int    `firestore:"goalsDiff,omitempty"`
	Points      int    `firestore:"points,omitempty"`
	LastUpdated string `firestore:"lastUpdate,omitempty"`
}

//Stat contains  team statistics
type Stat struct {
	MatchsPlayed int `firestore:"matchesPlayed,omitempty"`
	Win          int `firestore:"win,omitempty"`
	Draw         int `firestore:"draw,omitempty"`
	Lose         int `firestore:"lose,omitempty"`
	GoalsFor     int `firestore:"goalsFor,omitempty"`
	GoalsAgainst int `firestore:"goalsAgainst,omitempty"`
}

func (service *StandingsService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new data to firestore \n", len(records))

	batch := service.client.fs.Batch()

	for _, r := range records {

		s := Standings{}
		s.LeagueID = parseInt(r[0])
		s.Rank = parseInt(r[1])
		s.TeamID = parseInt(r[2])
		s.TeamName = r[3]
		s.Logo = r[4]
		s.Group = r[5]
		s.Status = r[6]
		s.Form = r[7]
		s.Description = r[8]
		s.AllStat.MatchsPlayed = parseInt(r[9])
		s.AllStat.Win = parseInt(r[10])
		s.AllStat.Lose = parseInt(r[11])
		s.AllStat.Draw = parseInt(r[12])
		s.AllStat.GoalsFor = parseInt(r[13])
		s.AllStat.GoalsAgainst = parseInt(r[14])
		s.HomeStat.MatchsPlayed = parseInt(r[15])
		s.HomeStat.Win = parseInt(r[16])
		s.HomeStat.Lose = parseInt(r[17])
		s.HomeStat.Draw = parseInt(r[18])
		s.HomeStat.GoalsFor = parseInt(r[19])
		s.HomeStat.GoalsAgainst = parseInt(r[20])
		s.AwayStat.MatchsPlayed = parseInt(r[21])
		s.AwayStat.Win = parseInt(r[22])
		s.AwayStat.Lose = parseInt(r[23])
		s.AwayStat.Draw = parseInt(r[24])
		s.AwayStat.GoalsFor = parseInt(r[25])
		s.AwayStat.GoalsAgainst = parseInt(r[26])
		leagueRef := service.client.fs.Collection("football").Doc(leagueName)

		docRef := leagueRef.
			Collection("leagues").
			Doc("leagueId_" + r[0]).
			Collection("standings").
			Doc(DocWithIDAndName(r[2], r[3]))

		batch.Set(docRef, s)
	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return err
}
