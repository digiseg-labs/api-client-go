package main

import (
	"context"

	digiseg "github.com/digiseg-labs/api-client-go/openapi"
)

func GetAudiences(apiKey string, ipAddress string) (*digiseg.AudienceResponse, error) {
	ctx := context.WithValue(context.Background(), digiseg.ContextAPIKeys, map[string]digiseg.APIKey{
		"apiKeyHeaderAuth": {Key: apiKey},
	})

	cfg := digiseg.NewConfiguration()
	api := digiseg.NewAPIClient(cfg)

	req := api.AudiencesAPI.ResolveAudiencesOfSingle(ctx, ipAddress)

	audiencesResponse, _, err := req.Execute()
	return audiencesResponse, err
}
