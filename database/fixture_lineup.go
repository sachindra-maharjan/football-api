package database

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type FixtureLineUpService dbservice

type FixtureTeams struct {
	HomeTeam FixtureTeam `firestore:"home_team,omitempty"`
	AwayTeam FixtureTeam `firestore:"away_team,omitempty"`
}

//FixtureTeam Fixture Team Info
type FixtureTeam struct {
	Coach      string                `firestore:"coach,omitempty"`
	CoachID    int                   `firestore:"coach_id,omitempty"`
	Formation  string                `firestore:"formation,omitempty"`
	TeamID     int                   `firestore:"team_id,omitempty"`
	StartingXI map[string]PlayerInfo `firestore:"startingXI,omitempty"`
	Substitute map[string]PlayerInfo `firestore:"substitutes,omitempty"`
}

//Player Starting Player Info
type PlayerInfo struct {
	PlayerID   int    `firestore:"player_id,omitempty"`
	PlayerName string `firestore:"player,omitempty"`
	Number     int    `firestore:"number,omitempty"`
	Position   string `firestore:"pos,omitempty"`
}

func (s *FixtureLineUpService) Add(ctx context.Context, leagueName string, records [][]string) error {

	fmt.Printf("Adding %d new fixture event data to firestore \n", len(records))

	batch := s.client.fs.Batch()
	var team FixtureTeams
	for i, r := range records {
		if i%2 == 0 {
			team = FixtureTeams{}
			team.HomeTeam = getTeam(r)
		} else {
			team.AwayTeam = getTeam(r)
			leagueRef := s.client.fs.Collection("football-leagues").Doc(leagueName)
			docRef := leagueRef.
				Collection("leagues").
				Doc("leagueId_" + r[0]).
				Collection("fixtures").
				Doc("fixtureId_" + r[1]).
				Collection("fixture_details").
				Doc("lineups")
			fmt.Printf("importing lineup in %s \n ", docRef.Path)
			batch.Set(docRef, team)
		}
	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}

func getTeam(record []string) FixtureTeam {
	ft := FixtureTeam{}
	ft.CoachID = parseInt(record[2])
	ft.Coach = record[3]
	ft.Formation = record[4]
	ft.TeamID = parseInt(record[5])
	ft.StartingXI = parseLineup(record[6])
	ft.Substitute = parseLineup(record[7])
	return ft
}

func parseLineup(lineup string) map[string]PlayerInfo {
	allplayers := strings.Split(lineup, "|")
	players := make(map[string]PlayerInfo)

	if len(allplayers) < 0 {
		return players
	}

	for _, p := range allplayers {
		player := strings.Split(p, "-")
		if len(player) >= 4 {
			playerInfo := PlayerInfo{
				PlayerID:   parseInt(player[0]),
				PlayerName: player[1],
				Number:     parseInt(player[2]),
				Position:   player[3],
			}
			players[strconv.Itoa(playerInfo.PlayerID)] = playerInfo
		}
	}

	return players
}
