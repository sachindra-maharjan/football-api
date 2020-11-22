package main

import (
	"casino_royal/vault/api"
	"context"
	"fmt"
)

func main() {

	result, err := getAllLeague(2)

	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}

	//fmt.Printf("Result: %+v\n", result)
	for _, league := range result.API.Leagues {
		fmt.Printf("%+v", league)
	}

}

func getAllLeague(leagueID int) (*api.LeagueResult, error) {
	api := api.NewClient(nil)
	leagueResult, _, err := api.LeagueService.ListAll(context.Background(), 2)
	return leagueResult, err

}
