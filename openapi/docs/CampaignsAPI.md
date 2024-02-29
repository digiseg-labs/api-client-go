# \CampaignsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /campaigns | Create campaign
[**DeleteCampaignBannerImage**](CampaignsAPI.md#DeleteCampaignBannerImage) | **Delete** /campaigns/{campaign_id}/assets/banner_image | Delete campaign banner image logo
[**DeleteCampaignById**](CampaignsAPI.md#DeleteCampaignById) | **Delete** /campaigns/{campaign_id} | Delete campaign
[**GetCampaignBannerImage**](CampaignsAPI.md#GetCampaignBannerImage) | **Get** /campaigns/{campaign_id}/assets/banner_image | Get campaign banner image
[**GetCampaignById**](CampaignsAPI.md#GetCampaignById) | **Get** /campaigns/{campaign_id} | Get campaign
[**ListCampaigns**](CampaignsAPI.md#ListCampaigns) | **Get** /campaigns | List campaigns
[**QueryCampaignAudienceStats**](CampaignsAPI.md#QueryCampaignAudienceStats) | **Get** /campaigns/{campaign_id}/stats/audiences | Audience statistics for campaign
[**QueryCampaignCountryStats**](CampaignsAPI.md#QueryCampaignCountryStats) | **Get** /campaigns/{campaign_id}/stats/countries | Country statistics for campaign
[**QueryCampaignFrequencyStats**](CampaignsAPI.md#QueryCampaignFrequencyStats) | **Get** /campaigns/{campaign_id}/stats/frequencies | Frequency statistics for campaign
[**QueryCampaignTimingStats**](CampaignsAPI.md#QueryCampaignTimingStats) | **Get** /campaigns/{campaign_id}/stats/timing | Timing statistics for campaign
[**UpdateCampaignById**](CampaignsAPI.md#UpdateCampaignById) | **Put** /campaigns/{campaign_id} | Update campaign
[**UploadCampaignBannerImage**](CampaignsAPI.md#UploadCampaignBannerImage) | **Put** /campaigns/{campaign_id}/assets/banner_image | Upload campaign banner image



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
	campaignCreation := *openapiclient.NewCampaignCreation("Acme Summer 2024 brand awareness", openapiclient.CampaignEventSet("impressions_only")) // CampaignCreation | 

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


## DeleteCampaignBannerImage

> DeleteCampaignBannerImage(ctx, campaignId).Execute()

Delete campaign banner image logo



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
	r, err := apiClient.CampaignsAPI.DeleteCampaignBannerImage(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaignBannerImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCampaignBannerImageRequest struct via the builder pattern


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


## GetCampaignBannerImage

> GetCampaignBannerImage(ctx, campaignId).Execute()

Get campaign banner image



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
	r, err := apiClient.CampaignsAPI.GetCampaignBannerImage(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignBannerImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetCampaignBannerImageRequest struct via the builder pattern


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

> ListCampaigns200Response ListCampaigns(ctx).Sort(sort).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

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
	sort := "-created_at" // string | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  (optional) (default to "created_at")
	filterLabel := "CTV" // string | Optional parameter used to filter by campaign label (optional)
	filterNameContains := "acme" // string | Optional parameter used to search for campaigns where the name contains a substring (case insensitive) (optional)
	filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query campaigns by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.ListCampaigns(context.Background()).Sort(sort).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
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
 **sort** | **string** | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  | [default to &quot;created_at&quot;]
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


## QueryCampaignAudienceStats

> QueryCampaignAudienceStats200Response QueryCampaignAudienceStats(ctx, campaignId).Execute()

Audience statistics for campaign



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
	resp, r, err := apiClient.CampaignsAPI.QueryCampaignAudienceStats(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.QueryCampaignAudienceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCampaignAudienceStats`: QueryCampaignAudienceStats200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.QueryCampaignAudienceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryCampaignAudienceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryCampaignAudienceStats200Response**](QueryCampaignAudienceStats200Response.md)

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


## QueryCampaignFrequencyStats

> QueryCampaignFrequencyStats200Response QueryCampaignFrequencyStats(ctx, campaignId).Execute()

Frequency statistics for campaign



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
	resp, r, err := apiClient.CampaignsAPI.QueryCampaignFrequencyStats(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.QueryCampaignFrequencyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCampaignFrequencyStats`: QueryCampaignFrequencyStats200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.QueryCampaignFrequencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryCampaignFrequencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryCampaignFrequencyStats200Response**](QueryCampaignFrequencyStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCampaignTimingStats

> QueryCampaignTimingStats200Response QueryCampaignTimingStats(ctx, campaignId).Execute()

Timing statistics for campaign



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
	resp, r, err := apiClient.CampaignsAPI.QueryCampaignTimingStats(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.QueryCampaignTimingStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCampaignTimingStats`: QueryCampaignTimingStats200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.QueryCampaignTimingStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryCampaignTimingStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryCampaignTimingStats200Response**](QueryCampaignTimingStats200Response.md)

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


## UploadCampaignBannerImage

> UploadCampaignBannerImage(ctx, campaignId).Body(body).Execute()

Upload campaign banner image



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
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.UploadCampaignBannerImage(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UploadCampaignBannerImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadCampaignBannerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: image/gif, image/png, image/jpeg, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

