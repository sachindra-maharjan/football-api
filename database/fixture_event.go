package database

import (
	"context"
	"fmt"
	"strconv"
)

type FixtureEventService dbservice

//FixtureEvent contains the fixture event infromation
type FixtureEvent struct {
	LeagueID       int    `firestore:"league_id"`
	FixtureID      int    `firestore:"fixture_id"`
	Elapsed        int    `firestore:"elapsed"`
	ElapsedPlus    int    `firestore:"elapsed_plus,omitempty"`
	TeamID         int    `firestore:"team_id"`
	TeamName       string `firestore:"teamName"`
	PlayerID       int    `firestore:"player_id"`
	Player         string `firestore:"player"`
	AssistPlayerID int    `firestore:"assist_id"`
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

	eventMap := map[int]Events{}
	var events Events

	for _, r := range records {
		f := FixtureEvent{}
		f.LeagueID = parseInt(r[0])
		f.FixtureID = parseInt(r[1])
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

		if _, ok := eventMap[f.FixtureID]; ok {
			events = eventMap[f.FixtureID]
		} else {
			events = Events{}
		}
		events.Event = append(events.Event, f)
		eventMap[f.FixtureID] = events
	}

	for k, v := range eventMap {

		leagueRef := s.client.fs.Collection("football-leagues").Doc(leagueName)
		docRef := leagueRef.
			Collection("leagues").
			Doc("leagueId_" + strconv.Itoa(v.Event[0].LeagueID)).
			Collection("fixtures").
			Doc("fixtureId_" + strconv.Itoa(k)).
			Collection("fixture_details").
			Doc("events")

		fmt.Printf("adding fixture event in %s\n", docRef.Path)

		batch.Set(docRef, v)
	}

	_, err := batch.Commit(ctx)

	if err != nil {
		fmt.Println("Error occurred when commiting batch.", err)
	}

	return nil
}
