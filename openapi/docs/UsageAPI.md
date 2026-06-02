# \UsageAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDataDailyUsage**](UsageAPI.md#ListDataDailyUsage) | **Get** /usage/{data}/daily | List daily usage of Audience Data
[**ListDataMonthlyUsage**](UsageAPI.md#ListDataMonthlyUsage) | **Get** /usage/{data}/monthly | List monthly usage of Audience Data
[**ListDataRealtimeUsage**](UsageAPI.md#ListDataRealtimeUsage) | **Get** /usage/{data}/realtime | List realtime usage of Audience Data



## ListDataDailyUsage

> ListDataDailyUsage200Response ListDataDailyUsage(ctx, data).FilterDateFrom(filterDateFrom).FilterDateTo(filterDateTo).FilterAccountId(filterAccountId).Execute()

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
	data := "data_example" // string | 
	filterDateFrom := time.Now() // string | Date to query from (optional)
	filterDateTo := time.Now() // string | Date to query to (optional)
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListDataDailyUsage(context.Background(), data).FilterDateFrom(filterDateFrom).FilterDateTo(filterDateTo).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListDataDailyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataDailyUsage`: ListDataDailyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListDataDailyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataDailyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDateFrom** | **string** | Date to query from | 
 **filterDateTo** | **string** | Date to query to | 
 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListDataDailyUsage200Response**](ListDataDailyUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataMonthlyUsage

> ListDataMonthlyUsage200Response ListDataMonthlyUsage(ctx, data).FilterYear(filterYear).FilterAccountId(filterAccountId).Execute()

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
	data := "data_example" // string | 
	filterYear := int32(56) // int32 | Year to filter by (optional)
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListDataMonthlyUsage(context.Background(), data).FilterYear(filterYear).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListDataMonthlyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataMonthlyUsage`: ListDataMonthlyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListDataMonthlyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataMonthlyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterYear** | **int32** | Year to filter by | 
 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListDataMonthlyUsage200Response**](ListDataMonthlyUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataRealtimeUsage

> ListDataRealtimeUsage200Response ListDataRealtimeUsage(ctx, data).FilterAccountId(filterAccountId).Execute()

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
	data := "data_example" // string | 
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to query usage of specific account IDs (only available to super admins).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListDataRealtimeUsage(context.Background(), data).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListDataRealtimeUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataRealtimeUsage`: ListDataRealtimeUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListDataRealtimeUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataRealtimeUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterAccountId** | **string** | Optional parameter used to query usage of specific account IDs (only available to super admins).  | 

### Return type

[**ListDataRealtimeUsage200Response**](ListDataRealtimeUsage200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

