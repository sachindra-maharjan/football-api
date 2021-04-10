package main

import (
	"casino_royal/vault/api"
	"context"
	"fmt"
)

func main() {
	client := api.NewClient(nil, []string{})

	standingResult, _, err := client.StandingService.GetLeagueStandings(context.Background(), 2)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(standingResult)

}
