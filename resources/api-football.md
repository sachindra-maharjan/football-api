# API-FOOTBALL

### API Call Strategy

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
