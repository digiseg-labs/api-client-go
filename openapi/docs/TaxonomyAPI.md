# \TaxonomyAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudience**](TaxonomyAPI.md#GetAudience) | **Get** /taxonomy/audiences/{audience_code} | Get an audience by its code
[**GetCategory**](TaxonomyAPI.md#GetCategory) | **Get** /taxonomy/categories/{category_code} | Get a category by its code
[**ListAudiencePlatforms**](TaxonomyAPI.md#ListAudiencePlatforms) | **Get** /taxonomy/audience_platforms | List audience platforms
[**ListAudiences**](TaxonomyAPI.md#ListAudiences) | **Get** /taxonomy/audiences | List audiences
[**ListCountries**](TaxonomyAPI.md#ListCountries) | **Get** /taxonomy/countries | List countries



## GetAudience

> AudienceResponse1 GetAudience(ctx, audienceCode).Platform(platform).Country(country).Execute()

Get an audience by its code

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
	audienceCode := "053" // string | The code of the audience to retrieve
	platform := "adform" // string | A platform code to apply for platform-specific audience codes (optional)
	country := "US" // string | A country code to apply for platform-specific and country-specific audience codes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.GetAudience(context.Background(), audienceCode).Platform(platform).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.GetAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudience`: AudienceResponse1
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.GetAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceCode** | **string** | The code of the audience to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platform** | **string** | A platform code to apply for platform-specific audience codes | 
 **country** | **string** | A country code to apply for platform-specific and country-specific audience codes | 

### Return type

[**AudienceResponse1**](AudienceResponse1.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategory

> CategoryResponse GetCategory(ctx, categoryCode).Platform(platform).Country(country).Execute()

Get a category by its code

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
	categoryCode := "A" // string | The code of the category to retrieve
	platform := "adform" // string | A platform code to apply for platform-specific audience codes (optional)
	country := "US" // string | A country code to apply for platform-specific and country-specific audience codes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.GetCategory(context.Background(), categoryCode).Platform(platform).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.GetCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategory`: CategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.GetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryCode** | **string** | The code of the category to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platform** | **string** | A platform code to apply for platform-specific audience codes | 
 **country** | **string** | A country code to apply for platform-specific and country-specific audience codes | 

### Return type

[**CategoryResponse**](CategoryResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudiencePlatforms

> ListAudiencePlatforms200Response ListAudiencePlatforms(ctx).Execute()

List audience platforms

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.ListAudiencePlatforms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.ListAudiencePlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAudiencePlatforms`: ListAudiencePlatforms200Response
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.ListAudiencePlatforms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAudiencePlatformsRequest struct via the builder pattern


### Return type

[**ListAudiencePlatforms200Response**](ListAudiencePlatforms200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudiences

> ListAudiences200Response ListAudiences(ctx).Platform(platform).Country(country).PageSize(pageSize).PageAfter(pageAfter).Sort(sort).Filter(filter).Labels(labels).Execute()

List audiences



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
	platform := "adform" // string | A platform code to apply for platform-specific audience codes (optional)
	country := "US" // string | A country code to apply for platform-specific and country-specific audience codes (optional)
	pageSize := int32(50) // int32 | Number of items to return per page (optional)
	pageAfter := "100" // string | Cursor for pagination to get the next page of results (optional)
	sort := "desc" // string | Sort order for audiences (by code) and categories (by display name) - either ascending (default) or descending. (optional)
	filter := "filter_example" // string | Optional parameter used to search for categories and audiences where the name contains a substring (case insensitive) (optional)
	labels := "demographic,sport" // string | Optional parameter used to filter categories and audiences by labels. Multiple labels can be provided, separated by commas. When multiple labels are provided, only categories and audiences that have all of the specified labels will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.ListAudiences(context.Background()).Platform(platform).Country(country).PageSize(pageSize).PageAfter(pageAfter).Sort(sort).Filter(filter).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.ListAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAudiences`: ListAudiences200Response
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.ListAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | A platform code to apply for platform-specific audience codes | 
 **country** | **string** | A country code to apply for platform-specific and country-specific audience codes | 
 **pageSize** | **int32** | Number of items to return per page | 
 **pageAfter** | **string** | Cursor for pagination to get the next page of results | 
 **sort** | **string** | Sort order for audiences (by code) and categories (by display name) - either ascending (default) or descending. | 
 **filter** | **string** | Optional parameter used to search for categories and audiences where the name contains a substring (case insensitive) | 
 **labels** | **string** | Optional parameter used to filter categories and audiences by labels. Multiple labels can be provided, separated by commas. When multiple labels are provided, only categories and audiences that have all of the specified labels will be returned. | 

### Return type

[**ListAudiences200Response**](ListAudiences200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCountries

> ListCountries200Response ListCountries(ctx).Execute()

List countries

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.ListCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.ListCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCountries`: ListCountries200Response
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.ListCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCountriesRequest struct via the builder pattern


### Return type

[**ListCountries200Response**](ListCountries200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

