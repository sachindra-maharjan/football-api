package client

import (
	"casino_royal/vault/api"
	"context"
)

type HttpClient struct {
	client *api.Client
}

func NewHttpClient() HttpClient {
	return HttpClient{
		client: &api.Client{},
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
