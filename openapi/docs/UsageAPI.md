# \UsageAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAudienceDataDailyUsage**](UsageAPI.md#ListAudienceDataDailyUsage) | **Get** /usage/audience_data/daily | List daily usage of Audience Data
[**ListAudienceDataMonthlyUsage**](UsageAPI.md#ListAudienceDataMonthlyUsage) | **Get** /usage/audience_data/monthly | List monthly usage of Audience Data
[**ListAudienceDataRealtimeUsage**](UsageAPI.md#ListAudienceDataRealtimeUsage) | **Get** /usage/audience_data/realtime | List realtime usage of Audience Data



## ListAudienceDataDailyUsage

> ListAudienceDataDailyUsage200Response ListAudienceDataDailyUsage(ctx).FilterDateFrom(filterDateFrom).FilterDateTo(filterDateTo).FilterAccountId(filterAccountId).Execute()

List daily usage of Audience Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	filterDateFrom := time.Now() // string | Date to query from (optional)
	filterDateTo := time.Now() // string | Date to query to (optional)
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListAudienceDataDailyUsage(context.Background()).FilterDateFrom(filterDateFrom).FilterDateTo(filterDateTo).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListAudienceDataDailyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAudienceDataDailyUsage`: ListAudienceDataDailyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListAudienceDataDailyUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceDataDailyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateFrom** | **string** | Date to query from | 
 **filterDateTo** | **string** | Date to query to | 
 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListAudienceDataDailyUsage200Response**](ListAudienceDataDailyUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceDataMonthlyUsage

> ListAudienceDataMonthlyUsage200Response ListAudienceDataMonthlyUsage(ctx).FilterYear(filterYear).FilterAccountId(filterAccountId).Execute()

List monthly usage of Audience Data



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
	filterYear := int32(56) // int32 | Year to filter by (optional)
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListAudienceDataMonthlyUsage(context.Background()).FilterYear(filterYear).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListAudienceDataMonthlyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAudienceDataMonthlyUsage`: ListAudienceDataMonthlyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListAudienceDataMonthlyUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceDataMonthlyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterYear** | **int32** | Year to filter by | 
 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListAudienceDataMonthlyUsage200Response**](ListAudienceDataMonthlyUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceDataRealtimeUsage

> ListAudienceDataRealtimeUsage200Response ListAudienceDataRealtimeUsage(ctx).FilterAccountId(filterAccountId).Execute()

List realtime usage of Audience Data



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
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListAudienceDataRealtimeUsage(context.Background()).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListAudienceDataRealtimeUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAudienceDataRealtimeUsage`: ListAudienceDataRealtimeUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListAudienceDataRealtimeUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceDataRealtimeUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListAudienceDataRealtimeUsage200Response**](ListAudienceDataRealtimeUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

