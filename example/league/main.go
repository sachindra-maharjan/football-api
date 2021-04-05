package main

import (
	"casino_royal/vault/api"
	csvutil "casino_royal/vault/util"
	"context"
	"fmt"
	"log"
)

func main() {

	data, err := getAllLeague(2)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}

	err = csvutil.Write("pl/2020/league.csv", data)
	if err != nil {
		log.Fatalf("csvutil.Write returned error %v", err)
	}
}

func getAllLeague(leagueID int) ([][]string, error) {
	api := api.NewClient(nil)
	leagueResult, _, err := api.LeagueService.ListAll(context.Background(), 2)

	if err != nil {
		return nil, err
	}

	result, err := api.LeagueService.Convert(leagueResult, true)
	if err != nil {
		return nil, err
	}
	return result, nil
}
