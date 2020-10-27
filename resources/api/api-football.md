# API-FOOTBALL

## Api Documentation

https://www.api-football.com/documentation

## API Call Strategy

- available season by league id : https://api-football-v1.p.rapidapi.com/v2/leagues/league/{league_id} - size: 10

  - teams for league: https://api-football-v1.p.rapidapi.com/v2/teams/league/{league_id} - size: 20

    - coach for team: https://api-football-v1.p.rapidapi.com/v2/coachs/team/{team_id} - size: 1
    - transfers: https://api-football-v1.p.rapidapi.com/v2/transfers/team/{team_id} - size: 1

  - standings: https://api-football-v1.p.rapidapi.com/v2/leagueTable/{league_id} - size: 1
  - player for a team by season: https://api-football-v1.p.rapidapi.com/v2/players/squad/{team_id}/{season} - size: 25

    - player statistic for all available season: https://api-football-v1.p.rapidapi.com/v2/players/player/{player_id} - size: 1
    - sidelined: https://api-football-v1.p.rapidapi.com/v2/sidelined/player/{player_id} - size: 1

  - All fixture for a league by season: https://api-football-v1.p.rapidapi.com/v2/fixtures/league/{league_id}?timezone={timezone} - size: 380
    - Events for fixture: https://api-football-v1.p.rapidapi.com/v2/events/{fixture_id} - size: 1
    - Line Ups: https://api-football-v1.p.rapidapi.com/v2/lineups/157215 - size: 1
    - Statistics : https://api-football-v1.p.rapidapi.com/v2/statistics/fixture/{fixture_id} - size: 1
    - player statistic for fixture: https://api-football-v1.p.rapidapi.com/v2/players/fixture/{fixture_id} - size: 1
    - odds: https://api-football-v1.p.rapidapi.com/v2/odds/fixture/{fixture_id} - size: 1

Total calls: 19730 (10 _ ((20 + 1 + 1) + 1 + (25 _ 2) + 380 \* 5))

## Api Requests

### Request Headers

- x-rapidapi-host: api-football-v1.p.rapidapi.com
- x-rapidapi-key: \*\*\*

### Requests

#### Season By League ID

- Url: https://api-football-v1.p.rapidapi.com/v2/leagues/league/{league_id}
- Response

```json
{
  "api": {
    "results": 970,
    "leagues": [
      {
        "league_id": 1,
        "name": "World Cup",
        "type": "Cup",
        "country": "World",
        "country_code": null,
        "season": 2018,
        "season_start": "2018-06-14",
        "season_end": "2018-07-15",
        "logo": "https://media.api-sports.io/football/leagues/1.png",
        "flag": null,
        "standings": 1,
        "is_current": 1,
        "coverage": {
          "standings": true,
          "fixtures": {
            "events": true,
            "lineups": true,
            "statistics": true,
            "players_statistics": false
          },
          "players": true,
          "topScorers": true,
          "predictions": true,
          "odds": false
        }
      }
    ]
  }
}
```

#### Teams By League ID

https://api-football-v1.p.rapidapi.com/v2/teams/league/{league_id}

#### Coach By Team ID

https://api-football-v1.p.rapidapi.com/v2/coachs/team/{team_id}

#### Transfer By Team ID

https://api-football-v1.p.rapidapi.com/v2/transfers/team/{team_id}

#### Standings By League ID

https://api-football-v1.p.rapidapi.com/v2/leagueTable/{league_id}

#### Player By Team ID By League ID

https://api-football-v1.p.rapidapi.com/v2/players/squad/{team_id}/{season}

#### Fixtures By League ID By Season

https://api-football-v1.p.rapidapi.com/v2/fixtures/league/{league_id}?timezone={timezone}

#### Events By Fixture ID

https://api-football-v1.p.rapidapi.com/v2/events/{fixture_id}

#### Line Ups By Fixture ID

https://api-football-v1.p.rapidapi.com/v2/lineups/{fixture_id}

#### Statistics By Fixture ID

https://api-football-v1.p.rapidapi.com/v2/statistics/fixture/{fixture_id}

#### Odds By Fixture ID

https://api-football-v1.p.rapidapi.com/v2/odds/fixture/{fixture_id}
