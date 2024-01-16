package main

import (
	"context"
	"os"

	digiseg "github.com/digiseg-labs/api-client-go/pkg"
)

func Authenticate() (string, error) {
	cfg := digiseg.NewConfiguration()
	api := digiseg.NewAPIClient(cfg)

	username := os.Getenv("DIGISEG_USERNAME")
	password := os.Getenv("DIGISEG_PASSWORD")

	req := api.AuthAPI.CreateAccessToken(context.Background()).AuthTokenRequest(digiseg.AuthTokenRequest{
		Username: username,
		Password: &password,
	})
	authResp, _, err := req.Execute()
	if err != nil {
		return "", err
	}
	return authResp.AccessToken, nil
}
