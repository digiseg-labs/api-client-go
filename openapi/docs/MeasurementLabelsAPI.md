# \MeasurementLabelsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMeasurementLabels**](MeasurementLabelsAPI.md#ListMeasurementLabels) | **Get** /measurement/labels | List measurement labels



## ListMeasurementLabels

> ListMeasurementLabels200Response ListMeasurementLabels(ctx).FilterAccountId(filterAccountId).Execute()

List measurement labels



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
	filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query measurement labels by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeasurementLabelsAPI.ListMeasurementLabels(context.Background()).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementLabelsAPI.ListMeasurementLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeasurementLabels`: ListMeasurementLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `MeasurementLabelsAPI.ListMeasurementLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMeasurementLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAccountId** | **string** | Optional parameter used to query measurement labels by specific account IDs (only available to super admins). The value &#x60;*&#x60; is synonymous for \&quot;all accounts\&quot;.  | [default to &quot;The user&#39;s account ID&quot;]

### Return type

[**ListMeasurementLabels200Response**](ListMeasurementLabels200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

