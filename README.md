# Digiseg API client for Go

This module provides a SDK for interacting with the Digiseg API.

Documentation for the API can be found on https://developer.digiseg.net/

## Installing and using the SDK

Install the module in your Go project:

```sh
go get github.com/digiseg-labs/api-client-go/openapi
```

Import the module and start using it:

```go
import (
	digiseg "github.com/digiseg-labs/api-client-go/openapi"
)

// create API client
api := digiseg.NewAPIClient(digiseg.NewConfiguration())

// look up audiences
ctx := context.WithValue(context.Background(), digiseg.ContextAPIKeys, map[string]digiseg.APIKey{
    "apiKeyHeaderAuth": {Key: apiKey},
})
req = api.AudiencesAPI.ResolveAudiencesOfSingle(ctx, ipAddress)
audiencesResponse, _, err := req.Execute()
```

## Examples

More code examples for how to use this module can be found in the `examples` directory.

## Development notes

### Code generation

The code in this repo is generated based on the OpenAPI spec of the APIs.

To regenerate, run:

```sh
make codegen
```
