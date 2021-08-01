package database

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/firestore"
)

type FixtureEventService dbservice

type FixtureEvents struct {
	HomeTeam []FixtureEvent `firestore:"homeTeam,omitempty"`
	AwayTeam []FixtureEvent `firestore:"awayTeam,omitempty"`
}

//FixtureEvent contains the fixture event infromation
type FixtureEvent struct {
	Elapsed        int    `firestore:"elapsed"`
	ElapsedPlus    int    `firestore:"elapsed_plus,omitempty"`
	TeamID         int    `firestore:"teamId"`
	TeamName       string `firestore:"teamName"`
	PlayerID       int    `firestore:"playerId"`
	Player         string `firestore:"player"`
	AssistPlayerID int    `firestore:"assistId"`
	AssistedBy     string `firestore:"assist"`
	Type           string `firestore:"type"`
	Detail         string `firestore:"detail"`
	Comments       string `firestore:"comments,omitempty"`
}

type Events struct {
	Event []FixtureEvent `firestore:"events"`
}

func (s *FixtureEventService) Add(ctx context.Context, leagueName string, records [][]string) error {

	fmt.Printf("Adding %d new fixture event data to firestore \n", len(records))

	batch := s.client.fs.Batch()
	var events FixtureEvents
	var homeTeam map[int][]FixtureEvent
	var awayTeam map[int][]FixtureEvent
	var leagueId int
	var fixtureId int
	var currentFixture int
	var homeTeamId int
	var awayTeamId int

	for _, r := range records {
		leagueId = parseInt(r[0])
		fixtureId = parseInt(r[1])

		f := FixtureEvent{}
		f.Elapsed = parseInt(r[2])
		f.ElapsedPlus = parseInt(r[3])
		f.TeamID = parseInt(r[4])
		f.TeamName = r[5]
		f.PlayerID = parseInt(r[6])
		f.Player = r[7]
		f.AssistPlayerID = parseInt(r[8])
		f.AssistedBy = r[9]
		f.Type = r[10]
		f.Detail = r[11]
		f.Comments = r[12]

		if currentFixture != fixtureId {
			if currentFixture > 0 {
				events.HomeTeam = homeTeam[homeTeamId]
				events.AwayTeam = awayTeam[awayTeamId]
				s.write(leagueName, leagueId, currentFixture, batch, events)
			}
			currentFixture = fixtureId
			events = FixtureEvents{}
			homeTeamId = 0
			awayTeamId = 0
			homeTeam = make(map[int][]FixtureEvent)
			awayTeam = make(map[int][]FixtureEvent)
		}

		if homeTeamId == 0 && awayTeamId == 0 {
			isHomeTeam, err := s.isHomeTeam(ctx, leagueName, leagueId, fixtureId, f.TeamID)
			if err != nil {
				return err
			}

			if isHomeTeam {
				homeTeamId = f.TeamID
				homeTeam[f.TeamID] = []FixtureEvent{}
			} else {
				awayTeamId = f.TeamID
				awayTeam[f.TeamID] = []FixtureEvent{}
			}
		}

		if homeTeamId == 0 && awayTeamId != f.TeamID {
			homeTeamId = f.TeamID
			homeTeam[f.TeamID] = []FixtureEvent{}
		}

		if awayTeamId == 0 && homeTeamId != f.TeamID {
			awayTeamId = f.TeamID
			awayTeam[f.TeamID] = []FixtureEvent{}
		}

		if _, ok := homeTeam[f.TeamID]; ok {
			homeTeam[f.TeamID] = append(homeTeam[f.TeamID], f)
		} else {
			awayTeam[f.TeamID] = append(awayTeam[f.TeamID], f)
		}
	}

	if currentFixture > 0 {
		events.HomeTeam = homeTeam[homeTeamId]
		events.AwayTeam = awayTeam[awayTeamId]
		s.write(leagueName, leagueId, currentFixture, batch, events)
	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}

func (s *FixtureEventService) write(leagueName string, leagueId int,
	fixtureId int, batch *firestore.WriteBatch, events FixtureEvents) {
	leagueRef := s.client.fs.Collection("football").Doc(leagueName)
	docRef := leagueRef.Collection("leagues").
		Doc("leagueId_" + strconv.Itoa(leagueId)).
		Collection("fixtures").
		Doc("fixtureId_" + strconv.Itoa(fixtureId)).
		Collection("fixture_details").
		Doc("events")

	fmt.Printf("adding fixture event in %s\n", docRef.Path)

	batch.Set(docRef, events)
}

func (s *FixtureEventService) isHomeTeam(ctx context.Context, leagueName string,
	leagueId int, fixtureId int, teamId int) (bool, error) {
	leagueRef := s.client.fs.Collection("football").Doc(leagueName)
	docRef := leagueRef.Collection("leagues").
		Doc("leagueId_" + strconv.Itoa(leagueId)).
		Collection("fixtures").
		Doc("fixtureId_" + strconv.Itoa(fixtureId))
	docSnap, err := docRef.Get(ctx)

	if err != nil {
		return false, err
	}

	var fixture Fixture
	docSnap.DataTo(&fixture)
	return fixture.HomeTeam.TeamID == teamId, nil
}
