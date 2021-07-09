package database

import (
	"context"
	"fmt"
	"strings"
)

type TeamService dbservice

type Team struct {
	TeamID        int    `firestore:"team_id,omitempty"`
	Name          string `firestore:"name,omitempty"`
	Code          string `firestore:"code,omitempty"`
	Country       string `firestore:"country,omitempty"`
	IsNational    bool   `firestore:"is_national,omitempty"`
	Founded       int    `firestore:"founded,omitempty"`
	VenueName     string `firestore:"veneu_name,omitempty"`
	VenueSurface  string `firestore:"venue_surface,omitempty"`
	VenueAddress  string `firestore:"venue_address,omitempty"`
	VenueCity     string `firestore:"venue_city,omitempty"`
	VenueCapacity int    `firestore:"venue_capacity,omitempty"`
}

func (s *TeamService) Add(ctx context.Context, leagueName string, records [][]string) error {
	fmt.Printf("Adding %d new fixture event data to firestore \n", len(records))
	batch := s.client.fs.Batch()

	for _, r := range records {
		team := Team{
			TeamID:        parseInt(r[1]),
			Name:          r[2],
			Code:          r[3],
			Country:       r[4],
			IsNational:    parseBool(r[5]),
			Founded:       parseInt(r[6]),
			VenueName:     r[7],
			VenueSurface:  r[8],
			VenueAddress:  r[9],
			VenueCity:     r[10],
			VenueCapacity: parseInt(r[11]),
		}
		leagueRef := s.client.fs.Collection("football-leagues").Doc(leagueName)
		docRef := leagueRef.
			Collection("leagues").
			Doc("leagueId_" + r[0]).
			Collection("teams").
			Doc(DocWithIDAndName(r[1], r[2]))
		fmt.Printf("importing lineup in %s \n ", docRef.Path)
		batch.Set(docRef, team)
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}

func DocWithIDAndName(id, name string) string {
	return id + "#" + strings.ToUpper(name)
}
