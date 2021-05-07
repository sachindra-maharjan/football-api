package database

import (
	"context"
	"fmt"
)

type FixturePlayerStatService dbservice

type FixturePlayerStat struct {
	FixtureID int            `firestore:"fixture_id,omitempty"`
	UpdatedAt int64          `firestore:"updateAt,omitempty"`
	HomeTeam  TeamPlayerStat `firestore:"home_team,omitempty"`
	AwayTeam  TeamPlayerStat `firestore:"away_team,omitempty"`
}

type TeamPlayerStat struct {
	TeamID     int                   `firestore:"team_id,omitempty"`
	TeamName   string                `firestore:"team_name,omitempty"`
	Statistics map[string]PlayerStat `firestore:"statistics,omitempty"`
}

//PlayerStat contains player statistics for a fixture
type PlayerStat struct {
	PlayerID      int    `firestore:"player_id,omitempty"`
	PlayerName    string `firestore:"player_name,omitempty"`
	Number        int    `firestore:"number,omitempty"`
	Position      string `firestore:"position,omitempty"`
	Rating        string `firestore:"rating,omitempty"`
	MinutesPlayed int    `firestore:"minutes_played,omitempty"`
	Captain       string `firestore:"captain,omitempty"`
	Substitute    string `firestore:"substitute,omitempty"`
	Offsides      int    `firestore:"offsides,omitempty"`
	Shots         struct {
		Total int `firestore:"total,omitempty"`
		On    int `firestore:"on,omitempty"`
	} `firestore:"shots,omitempty"`
	Goals struct {
		Total    int `firestore:"total,omitempty"`
		Conceded int `firestore:"conceded,omitempty"`
		Assists  int `firestore:"assists,omitempty"`
		Saves    int `firestore:"saves,omitempty"`
	} `firestore:"goals,omitempty"`
	Passes struct {
		Total    int `firestore:"total,omitempty"`
		Key      int `firestore:"key,omitempty"`
		Accuracy int `firestore:"accuracy,omitempty"`
	} `firestore:"passes,omitempty"`
	Tackles struct {
		Total         int `firestore:"total,omitempty"`
		Blocks        int `firestore:"blocks,omitempty"`
		Interceptions int `firestore:"interceptions,omitempty"`
	} `firestore:"tackles,omitempty"`
	Duels struct {
		Total int `firestore:"total,omitempty"`
		Won   int `firestore:"won,omitempty"`
	} `firestore:"duels,omitempty"`
	Dribbles struct {
		Attempts int `firestore:"attempts,omitempty"`
		Success  int `firestore:"success,omitempty"`
		Past     int `firestore:"past,omitempty"`
	} `firestore:"dribbles,omitempty"`
	Fouls struct {
		Drawn     int `firestore:"drawn,omitempty"`
		Committed int `firestore:"committed,omitempty"`
	} `firestore:"fouls,omitempty"`
	Cards struct {
		Yellow int `firestore:"yellow,omitempty"`
		Red    int `firestore:"red,omitempty"`
	} `firestore:"cards,omitempty"`
	Penalty Penalty `firestore:"penalty,omitempty"`
}

type Penalty struct {
	Won      int `firestore:"won,omitempty"`
	Commited int `firestore:"commited,omitempty"`
	Success  int `firestore:"success,omitempty"`
	Missed   int `firestore:"missed,omitempty"`
	Saved    int `firestore:"saved,omitempty"`
}

func (s *FixturePlayerStatService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new fixture event data to firestore \n", len(records))
	batch := s.client.fs.Batch()

	currentFixture := 0
	homeTeam := 0
	fixturePlayerStat := FixturePlayerStat{}
	for _, r := range records {
		fixtureId := parseInt(r[1])
		teamId := parseInt(r[5])

		if currentFixture != fixtureId {
			if fixturePlayerStat.FixtureID > 0 {
				leagueRef := s.client.fs.Collection("football-leagues").Doc(leagueName)
				docRef := leagueRef.
					Collection("leagues").
					Doc("leagueId_" + r[0]).
					Collection("fixtures").
					Doc("fixtureId_" + r[1]).
					Collection("fixture_details").
					Doc("player-stat")
				fmt.Printf("importing lineup in %s \n ", docRef.Path)
				batch.Set(docRef, fixturePlayerStat)
			}

			currentFixture = fixtureId
			homeTeam = 0
			fixturePlayerStat = FixturePlayerStat{
				FixtureID: fixtureId,
				UpdatedAt: parseInt64(r[2]),
				HomeTeam:  TeamPlayerStat{},
				AwayTeam:  TeamPlayerStat{},
			}
		}

		if homeTeam == 0 || homeTeam == teamId {
			homeTeam = teamId
			fixturePlayerStat.HomeTeam.TeamID = parseInt(r[5])
			fixturePlayerStat.HomeTeam.TeamName = r[6]
			if fixturePlayerStat.HomeTeam.Statistics == nil {
				fixturePlayerStat.HomeTeam.Statistics = make(map[string]PlayerStat)
			}
			fixturePlayerStat.HomeTeam.Statistics[r[3]] = s.getPlayerStat(r)
		} else if homeTeam != teamId {
			fixturePlayerStat.AwayTeam.TeamID = parseInt(r[5])
			fixturePlayerStat.AwayTeam.TeamName = r[6]
			if fixturePlayerStat.AwayTeam.Statistics == nil {
				fixturePlayerStat.AwayTeam.Statistics = make(map[string]PlayerStat)
			}
			fixturePlayerStat.AwayTeam.Statistics[r[3]] = s.getPlayerStat(r)
		}

	}

	_, err := batch.Commit(ctx)
	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}

func (s *FixturePlayerStatService) getPlayerStat(record []string) PlayerStat {
	playerstat := PlayerStat{
		PlayerID:      parseInt(record[3]),
		PlayerName:    record[4],
		Number:        parseInt(record[7]),
		Position:      record[8],
		Rating:        record[9],
		MinutesPlayed: parseInt(record[10]),
		Captain:       record[11],
		Substitute:    record[12],
		Offsides:      parseInt(record[13]),
	}

	playerstat.Shots.Total = parseInt(record[14])
	playerstat.Shots.On = parseInt(record[15])
	playerstat.Goals.Total = parseInt(record[16])
	playerstat.Goals.Conceded = parseInt(record[17])
	playerstat.Goals.Assists = parseInt(record[18])
	playerstat.Goals.Saves = parseInt(record[19])
	playerstat.Passes.Total = parseInt(record[20])
	playerstat.Passes.Key = parseInt(record[21])
	playerstat.Passes.Accuracy = parseInt(record[22])
	playerstat.Tackles.Total = parseInt(record[23])
	playerstat.Tackles.Blocks = parseInt(record[24])
	playerstat.Tackles.Interceptions = parseInt(record[25])
	playerstat.Duels.Total = parseInt(record[26])
	playerstat.Duels.Won = parseInt(record[27])
	playerstat.Dribbles.Attempts = parseInt(record[28])
	playerstat.Dribbles.Success = parseInt(record[29])
	playerstat.Dribbles.Past = parseInt(record[30])
	playerstat.Fouls.Drawn = parseInt(record[31])
	playerstat.Fouls.Committed = parseInt(record[32])
	playerstat.Cards.Yellow = parseInt(record[33])
	playerstat.Cards.Red = parseInt(record[34])
	playerstat.Penalty.Won = parseInt(record[35])
	playerstat.Penalty.Commited = parseInt(record[36])
	playerstat.Penalty.Success = parseInt(record[37])
	playerstat.Penalty.Missed = parseInt(record[38])
	playerstat.Penalty.Saved = parseInt(record[39])
	return playerstat
}
