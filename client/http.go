package client

import (
	"casino_royal/vault/api"
	"context"
	"fmt"
	"os"
	"strings"
)

type HttpClient struct {
	client *api.Client
}

func NewHttpClient() HttpClient {
	keys := os.Getenv("RAPID_API_KEYS")
	apiKeys = strings.Split(keys, ",")

	fmt.Println("Rapid API Keys: " + keys)

	return HttpClient{
		client: api.NewClient(nil, apiKeys),
	}
}

func (c HttpClient) League(leagueId int) (*api.LeagueResult, error) {
	leagueResult, _, err := c.client.LeagueService.ListAll(context.Background(), leagueId)

	if err != nil {
		return nil, err
	}

	return leagueResult, nil
}

func (c HttpClient) Healthy(host string) bool {
	return true
}
