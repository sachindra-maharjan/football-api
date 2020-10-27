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

- Url: https://api-football-v1.p.rapidapi.com/v2/teams/league/{league_id}
- Response

```json
{
  "api": {
    "results": 5,
    "teams": [
      {
        "team_id": 541,
        "name": "Real Madrid",
        "code": null,
        "logo": "https://media.api-sports.io/football/teams/541.png",
        "is_national": false,
        "country": "Spain",
        "founded": 1902,
        "venue_name": "Estadio Santiago Bernabéu",
        "venue_surface": "grass",
        "venue_address": "Avenida de Concha Espina 1, Chamartín",
        "venue_city": "Madrid",
        "venue_capacity": 85454
      }
    ]
  }
}
```

#### Coach By Team ID

- Url: https://api-football-v1.p.rapidapi.com/v2/coachs/team/{team_id}
- Reponse

```json
{
  "api": {
    "results": 1,
    "coachs": [
      {
        "id": 18,
        "name": "Unai Emery",
        "firstname": "Unai",
        "lastname": "Emery Etxegoien",
        "age": 48,
        "birth_date": "03/11/1971",
        "birth_place": "Hondarribia",
        "birth_country": "Spain",
        "nationality": "Spain",
        "height": null,
        "weight": null,
        "team": {
          "id": 42,
          "name": "Arsenal"
        },
        "career": [
          {
            "team": {
              "id": 42,
              "name": "Arsenal"
            },
            "start": "23/05/2018",
            "end": null
          },
          {
            "team": {
              "id": 85,
              "name": "PSG"
            },
            "start": "28/06/2016",
            "end": "23/05/2018"
          },
          {
            "team": {
              "id": 536,
              "name": "Sevilla"
            },
            "start": "15/01/2013",
            "end": "12/06/2016"
          },
          {
            "team": {
              "id": 558,
              "name": "Spartak Moskva"
            },
            "start": "10/06/2012",
            "end": "25/11/2012"
          },
          {
            "team": {
              "id": 532,
              "name": "Valencia"
            },
            "start": "01/07/2008",
            "end": "10/06/2012"
          },
          {
            "team": {
              "id": 723,
              "name": "Almería"
            },
            "start": "01/07/2006",
            "end": "30/06/2008"
          },
          {
            "team": {
              "id": null,
              "name": "Lorca Deportiva CF"
            },
            "start": "01/11/2004",
            "end": "30/06/2006"
          }
        ]
      }
    ]
  }
}
```

#### Transfer By Team ID

- Url: https://api-football-v1.p.rapidapi.com/v2/transfers/team/{team_id}
- Response

```json
{
  "api": {
    "results": 1,
    "transfers": [
      {
        "player_id": 19018,
        "player_name": "Timothy Evans Fosu-Mensah",
        "transfer_date": "2019-07-01",
        "type": null,
        "team_in": {
          "team_id": 33,
          "team_name": "Manchester United"
        },
        "team_out": {
          "team_id": 36,
          "team_name": "Fulham"
        },
        "lastUpdate": 1561726193
      }
    ]
  }
}
```

#### Standings By League ID

- Url: https://api-football-v1.p.rapidapi.com/v2/leagueTable/{league_id}
- Response

```json
{
  "api": {
    "results": 1,
    "standings": [
      [
        {
          "rank": 1,
          "team_id": 85,
          "teamName": "Paris Saint Germain",
          "logo": "https://media.api-sports.io/football/teams/85.png",
          "group": "Ligue 1",
          "forme": "DLWLL",
          "status": "same",
          "description": "Promotion - Champions League (Group Stage)",
          "all": {
            "matchsPlayed": 35,
            "win": 27,
            "draw": 4,
            "lose": 4,
            "goalsFor": 98,
            "goalsAgainst": 31
          },
          "home": {
            "matchsPlayed": 18,
            "win": 16,
            "draw": 2,
            "lose": 0,
            "goalsFor": 59,
            "goalsAgainst": 10
          },
          "away": {
            "matchsPlayed": 17,
            "win": 11,
            "draw": 2,
            "lose": 4,
            "goalsFor": 39,
            "goalsAgainst": 21
          },
          "goalsDiff": 67,
          "points": 85,
          "lastUpdate": "2019-05-04"
        }
      ]
    ]
  }
}
```

#### Player By Team ID By League ID

- Url: https://api-football-v1.p.rapidapi.com/v2/players/squad/{team_id}/{season}
- Response

```json
{
  "api": {
    "results": 33,
    "players": [
      {
        "player_id": 272,
        "player_name": "Adrien Rabiot",
        "firstname": "Adrien",
        "lastname": "Rabiot",
        "number": null,
        "position": "Midfielder",
        "age": 24,
        "birth_date": "03/04/1995",
        "birth_place": "Saint-Maurice",
        "birth_country": "France",
        "nationality": "France",
        "height": "188 cm",
        "weight": "71 kg"
      },
      {
        "player_id": 85062,
        "player_name": "Thiago Motta",
        "firstname": "Thiago",
        "lastname": "Motta",
        "number": null,
        "position": "Midfielder",
        "age": 36,
        "birth_date": "28/08/1982",
        "birth_place": "São Bernardo do Campo",
        "birth_country": "Brazil",
        "nationality": "Italy",
        "height": "187 cm",
        "weight": "83 kg"
      },
      {
        "player_id": 254,
        "player_name": "Gianluigi Buffon",
        "firstname": "Gianluigi",
        "lastname": "Buffon",
        "number": null,
        "position": "Goalkeeper",
        "age": 41,
        "birth_date": "28/01/1978",
        "birth_place": "Carrara",
        "birth_country": "Italy",
        "nationality": "Italy",
        "height": "192 cm",
        "weight": "92 kg"
      },
      {
        "player_id": 273,
        "player_name": "Marco Verratti",
        "firstname": "Marco",
        "lastname": "Verratti",
        "number": null,
        "position": "Midfielder",
        "age": 27,
        "birth_date": "05/11/1992",
        "birth_place": "Pescara",
        "birth_country": "Italy",
        "nationality": "Italy",
        "height": "165 cm",
        "weight": "60 kg"
      }
    ]
  }
}
```

#### Fixtures By League ID By Season

- Url: https://api-football-v1.p.rapidapi.com/v2/fixtures/league/{league_id}?timezone={timezone}
- Response

```json
{
  "api": {
    "results": 380,
    "fixtures": [
      {
        "fixture_id": 65,
        "league_id": 2,
        "league": {
          "name": "Premier League",
          "country": "England",
          "logo": "https://media.api-sports.io/football/leagues/2.png",
          "flag": "https://media.api-sports.io/flags/gb.svg"
        },
        "event_date": "2018-08-10T19:00:00+00:00",
        "event_timestamp": 1533927600,
        "firstHalfStart": 1533927600,
        "secondHalfStart": 1533931200,
        "round": "Regular Season - 1",
        "status": "Match Finished",
        "statusShort": "FT",
        "elapsed": 90,
        "venue": "Old Trafford (Manchester)",
        "referee": null,
        "homeTeam": {
          "team_id": 33,
          "team_name": "Manchester United",
          "logo": "https://media.api-sports.io/football/teams/33.png"
        },
        "awayTeam": {
          "team_id": 46,
          "team_name": "Leicester",
          "logo": "https://media.api-sports.io/football/teams/46.png"
        },
        "goalsHomeTeam": 2,
        "goalsAwayTeam": 1,
        "score": {
          "halftime": "1-0",
          "fulltime": "2-1",
          "extratime": null,
          "penalty": null
        }
      }
    ]
  }
}
```

#### Events By Fixture ID

- Url: https://api-football-v1.p.rapidapi.com/v2/events/{fixture_id}
- Response

```json
{
  "api": {
    "results": 18,
    "events": [
      {
        "elapsed": 25,
        "elapsed_plus": null,
        "team_id": 463,
        "teamName": "Aldosivi",
        "player_id": 6126,
        "player": "F. Andrada",
        "assist_id": null,
        "assist": null,
        "type": "Goal",
        "detail": "Normal Goal",
        "comments": null
      },
      {
        "elapsed": 44,
        "elapsed_plus": null,
        "team_id": 463,
        "teamName": "Aldosivi",
        "player_id": 6262,
        "player": "E. Iniguez",
        "assist_id": null,
        "assist": null,
        "type": "Card",
        "detail": "Yellow Card",
        "comments": null
      },
      {
        "elapsed": 46,
        "elapsed_plus": null,
        "team_id": 442,
        "teamName": "Defensa Y Justicia",
        "player_id": 5947,
        "player": "B. Merlini",
        "assist_id": 35695,
        "assist": "D. Rodriguez",
        "type": "subst",
        "detail": "D. Rodriguez",
        "comments": null
      }
    ]
  }
}
```

#### Line Ups By Fixture ID

- Url: https://api-football-v1.p.rapidapi.com/v2/lineups/{fixture_id}
- Response

```json
{
  "api": {
    "results": 2,
    "lineUps": {
      "Stade Brestois 29": {
        "coach_id": 34,
        "coach": "O. Dall’Oglio",
        "formation": "4-2-3-1",
        "startXI": [
          {
            "team_id": 106,
            "player_id": 20541,
            "player": "G. Larsonneur",
            "number": 1,
            "pos": "G"
          },
          {
            "team_id": 106,
            "player_id": 20684,
            "player": "D. Bain",
            "number": 17,
            "pos": "D"
          }
        ],
        "substitutes": [
          {
            "team_id": 106,
            "player_id": 22266,
            "player": "S. Grandsir",
            "number": 25,
            "pos": "M"
          },
          {
            "team_id": 106,
            "player_id": 21506,
            "player": "A. Mendy",
            "number": 15,
            "pos": "F"
          }
        ]
      },
      "Paris Saint Germain": {
        "coach_id": 40,
        "coach": "T. Tuchel",
        "formation": "4-2-3-1",
        "startXI": [
          {
            "team_id": 85,
            "player_id": 19014,
            "player": "Sergio Rico",
            "number": 16,
            "pos": "G"
          },
          {
            "team_id": 85,
            "player_id": 259,
            "player": "Thiago Silva",
            "number": 2,
            "pos": "D"
          }
        ],
        "substitutes": [
          {
            "team_id": 85,
            "player_id": 275,
            "player": "E. Choupo-Moting",
            "number": 17,
            "pos": "F"
          },
          {
            "team_id": 85,
            "player_id": 257,
            "player": "Marquinhos",
            "number": 5,
            "pos": "D"
          }
        ]
      }
    }
  }
}
```

#### Statistics By Fixture ID

- Url: https://api-football-v1.p.rapidapi.com/v2/statistics/fixture/{fixture_id}
- Response

```json
{
  "api": {
    "results": 16,
    "statistics": {
      "Shots on Goal": {
        "home": "5",
        "away": "2"
      },
      "Shots off Goal": {
        "home": "7",
        "away": "4"
      },
      "Total Shots": {
        "home": "17",
        "away": "8"
      },
      "Blocked Shots": {
        "home": "5",
        "away": "2"
      },
      "Shots insidebox": {
        "home": "12",
        "away": "5"
      },
      "Shots outsidebox": {
        "home": "5",
        "away": "3"
      },
      "Fouls": {
        "home": "15",
        "away": "14"
      },
      "Corner Kicks": {
        "home": "9",
        "away": "1"
      },
      "Offsides": {
        "home": "2",
        "away": "2"
      },
      "Ball Possession": {
        "home": "61%",
        "away": "39%"
      },
      "Yellow Cards": {
        "home": "2",
        "away": "3"
      },
      "Red Cards": {
        "home": "",
        "away": ""
      },
      "Goalkeeper Saves": {
        "home": "1",
        "away": "4"
      },
      "Total passes": {
        "home": "633",
        "away": "414"
      },
      "Passes accurate": {
        "home": "575",
        "away": "365"
      },
      "Passes %": {
        "home": "91%",
        "away": "88%"
      }
    }
  }
}
```

#### Odds By Fixture ID

- Url: https://api-football-v1.p.rapidapi.com/v2/odds/fixture/{fixture_id}
- Response

```json
{
    "api": {
        "results": 1,
        "paging": {
            "current": 1,
            "total": 1
        },
        "odds": [
            {
                "fixture": {
                    "league_id": 404,
                    "fixture_id": 108705,
                    "updateAt": 1557496046
                },
                "bookmakers": [
                    {
                        "bookmaker_id": 6,
                        "bookmaker_name": "bwin",
                        "bets": [
                            {
                                "label_id": 1,
                                "label_name": "Match Winner",
                                "values": [
                                    {
                                        "value": "Home",
                                        "odd": "2.20"
                                    },
                                    {
                                        "value": "Draw",
                                        "odd": "3.70"
                                    },
                                    {
                                        "value": "Away",
                                        "odd": "2.60"
                                    }
                                ]
                            },
                            {
                                "label_id": 8,
                                "label_name": "Both Teams To Score",
                                "values": [
                                    {
                                        "value": "Yes",
                                        "odd": "1.40"
                                    },
                                    {
                                        "value": "No",
                                        "odd": "2.75"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "bookmaker_id": 1,
                        "bookmaker_name": "10Bet",
                        "bets": [
                            {
                                "label_id": 1,
                                "label_name": "Match Winner",
                                "values": [
                                    {
                                        "value": "Home",
                                        "odd": "2.30"
                                    },
                                    {
                                        "value": "Draw",
                                        "odd": "3.60"
                                    },
                                    {
                                        "value": "Away",
                                        "odd": "2.50"
                                    }
                                ]
                            },
                            {
                                "label_id": 2,
                                "label_name": "Home/Away",
                                "values": [
                                    {
                                        "value": "Home",
                                        "odd": "1.77"
                                    },
                                    {
                                        "value": "Away",
                                        "odd": "1.83"
                                    }
                                ]
                            }
                        }
                    }
                ]
            }
        ]
    }
}
```
