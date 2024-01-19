# \CampaignsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /campaigns | Create campaign
[**DeleteCampaignById**](CampaignsAPI.md#DeleteCampaignById) | **Delete** /campaigns/{campaign_id} | Delete campaign
[**GetCampaignById**](CampaignsAPI.md#GetCampaignById) | **Get** /campaigns/{campaign_id} | Get campaign
[**ListCampaigns**](CampaignsAPI.md#ListCampaigns) | **Get** /campaigns | List campaigns
[**QueryCampaignCountryStats**](CampaignsAPI.md#QueryCampaignCountryStats) | **Get** /campaigns/{campaign_id}/stats/countries | Country statistics for campaign
[**UpdateCampaignById**](CampaignsAPI.md#UpdateCampaignById) | **Put** /campaigns/{campaign_id} | Update campaign



## CreateCampaign

> CreateCampaign201Response CreateCampaign(ctx).CampaignCreation(campaignCreation).Execute()

Create campaign



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
    campaignCreation := *openapiclient.NewCampaignCreation("Acme Summer 2024 brand awareness") // CampaignCreation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsAPI.CreateCampaign(context.Background()).CampaignCreation(campaignCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: CreateCampaign201Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignCreation** | [**CampaignCreation**](CampaignCreation.md) |  | 

### Return type

[**CreateCampaign201Response**](CreateCampaign201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignById

> DeleteCampaignById(ctx, campaignId).Execute()

Delete campaign



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
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAPI.DeleteCampaignById(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignById

> CreateCampaign201Response GetCampaignById(ctx, campaignId).Execute()

Get campaign



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
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsAPI.GetCampaignById(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignById`: CreateCampaign201Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCampaign201Response**](CreateCampaign201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaigns

> ListCampaigns200Response ListCampaigns(ctx).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

List campaigns



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
    filterLabel := "CTV" // string | Optional parameter used to filter by campaign label (optional)
    filterNameContains := "acme" // string | Optional parameter used to search for campaigns where the name contains a substring (case insensitive) (optional)
    filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query campaigns by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")
    pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
    pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsAPI.ListCampaigns(context.Background()).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.ListCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCampaigns`: ListCampaigns200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.ListCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLabel** | **string** | Optional parameter used to filter by campaign label | 
 **filterNameContains** | **string** | Optional parameter used to search for campaigns where the name contains a substring (case insensitive) | 
 **filterAccountId** | **string** | Optional parameter used to query campaigns by specific account IDs (only available to super admins). The value &#x60;*&#x60; is synonymous for \&quot;all accounts\&quot;.  | [default to &quot;The user&#39;s account ID&quot;]
 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 

### Return type

[**ListCampaigns200Response**](ListCampaigns200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCampaignCountryStats

> QueryCampaignCountryStats200Response QueryCampaignCountryStats(ctx, campaignId).Execute()

Country statistics for campaign



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
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsAPI.QueryCampaignCountryStats(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.QueryCampaignCountryStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCampaignCountryStats`: QueryCampaignCountryStats200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.QueryCampaignCountryStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryCampaignCountryStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryCampaignCountryStats200Response**](QueryCampaignCountryStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignById

> CreateCampaign201Response UpdateCampaignById(ctx, campaignId).CampaignMutation(campaignMutation).Execute()

Update campaign



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
    campaignId := "campaignId_example" // string | 
    campaignMutation := *openapiclient.NewCampaignMutation() // CampaignMutation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsAPI.UpdateCampaignById(context.Background(), campaignId).CampaignMutation(campaignMutation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UpdateCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignById`: CreateCampaign201Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.UpdateCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignMutation** | [**CampaignMutation**](CampaignMutation.md) |  | 

### Return type

[**CreateCampaign201Response**](CreateCampaign201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

