# \PopulationsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPopuplationByKey**](PopulationsAPI.md#GetPopuplationByKey) | **Get** /populations/{category_key}/{population_key} | Get Population by key
[**ListPopuplations**](PopulationsAPI.md#ListPopuplations) | **Get** /populations/{category_key} | List populations of category



## GetPopuplationByKey

> GetPopuplationByKey200Response GetPopuplationByKey(ctx, categoryKey, populationKey).Execute()

Get Population by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	categoryKey := "aip" // string | An identifier for the category
	populationKey := "dk" // string | An identifier for the population

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PopulationsAPI.GetPopuplationByKey(context.Background(), categoryKey, populationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PopulationsAPI.GetPopuplationByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPopuplationByKey`: GetPopuplationByKey200Response
	fmt.Fprintf(os.Stdout, "Response from `PopulationsAPI.GetPopuplationByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | An identifier for the category | 
**populationKey** | **string** | An identifier for the population | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPopuplationByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPopuplationByKey200Response**](GetPopuplationByKey200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPopuplations

> ListPopuplations200Response ListPopuplations(ctx, categoryKey).Execute()

List populations of category



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	categoryKey := "aip" // string | An identifier for the category

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PopulationsAPI.ListPopuplations(context.Background(), categoryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PopulationsAPI.ListPopuplations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPopuplations`: ListPopuplations200Response
	fmt.Fprintf(os.Stdout, "Response from `PopulationsAPI.ListPopuplations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | An identifier for the category | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPopuplationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPopuplations200Response**](ListPopuplations200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

