# Football-API CLI

This is a CLI application to fetch football api data from football-api hosted in Rapid-Api.

### RapidAPI - Football API Url

### Build

```script
 make football-api
```

### Available Commands

- league
- standings
- team
- top-scorer
- fixtures
- fixture-event
- fixture-lineup
- player-stat
- help

## Environment Variable

All the commands require environment variable RAPID_API_KEYS. This variable can be single key or a list of comma separated keys.

```script
export $RAPID_API_KEYS=key1,key2,key3
```

## Commands

The basepath is mandatory for all requests. This is the root path where csv file will be saved. Each league data will be saved in different folder.

### league

One request per league

```script
./bin/football-api league -leagueId=2 -basepath=<absolute path>
```

### standings

One request per league

```script
./bin/football-api standings -leagueId=2 -basepath=<absolute path>
```

### team

One request per league

```script
./bin/football-api team -leagueId=2 -basepath=<absolute path>
```

### Top-Scorer

```script
./bin/football-api top-scorer -leagueId=2 -basepath=<absolute path>
```

### fixtures

One request per league. One additional file will be created which contains list of fixture ids. This list can be used as input for other fixture related commands.

```script
./bin/football-api fixtures -leagueId=2 -basepath=<absolute path>
```

### fixture-event

One request per fixture. A list of fixtures can be given as input. A new file will be created for the first execution. Then afterward, data will be appended in the same file for subsiquent requests.

```script
./bin/football-api fixture-event -leagueId=2 -basepath=<absolute path> -fixtureId=fixtureId1,fixtureId2,...
```

### fixture-lineup

One request per fixture. A list of fixtures can be given as input. A new file will be created for the first execution. Then afterward, data will be appended in the same file for subsiquent requests.

```script
./bin/football-api fixture-lineup -leagueId=2 -basepath=<absolute path> -fixtureId=fixtureId1,fixtureId2,...
```

### player-stat

One request per fixture. A list of fixtures can be given as input. A new file will be created for the first execution. Then afterward, data will be appended in the same file for subsiquent requests.

```script
./bin/football-api player-stat -leagueId=2 -basepath=<absolute path> -fixtureId=fixtureId1,fixtureId2,...
```

### Help

```script
./bin/football-api --help
```
