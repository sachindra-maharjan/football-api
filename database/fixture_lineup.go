package database

import (
	"context"
	"fmt"
)

type FixtureLineUpService dbservice

//FixtureTeam Fixture Team Info
type FixtureTeam struct {
	Coach      string                `firestore:"coach,omitempty"`
	CoachID    int                   `firestore:"coach_id,omitempty"`
	Formation  string                `firestore:"formation,omitempty"`
	StartingXI map[string]PlayerInfo `firestore:"startXI,omitempty"`
	Substitute map[string]PlayerInfo `firestore:"substitutes,omitempty"`
}

//Player Starting Player Info
type PlayerInfo struct {
	TeamID     int    `firestore:"team_id,omitemtpy"`
	PlayerID   int    `firestore:"player_id,omitempty"`
	PlayerName string `firestore:"player,omitempty"`
	Number     int    `firestore:"number,omitempty"`
	Position   string `firestore:"pos,omitempty"`
}

func (s *FixtureLineUpService) Add(ctx context.Context, leagueName string, records [][]string) error {

	fmt.Printf("Adding %d new fixture event data to firestore \n", len(records))

	batch := s.client.fs.Batch()

	// for _, r := range records {
	// 	ft := FixtureTeam{}

	// }

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}
