package main

import (
	"context"

	digiseg "github.com/digiseg-labs/api-client-go/openapi"
)

func GetAudiences(accessToken string, ipAddress string) (*digiseg.AudienceResponse, error) {
	ctx := context.WithValue(context.Background(), digiseg.ContextAccessToken, accessToken)

	cfg := digiseg.NewConfiguration()
	api := digiseg.NewAPIClient(cfg)

	req := api.AudiencesAPI.ResolveAudiencesOfSingle(ctx, ipAddress)

	audiencesResponse, _, err := req.Execute()
	return audiencesResponse, err
}
