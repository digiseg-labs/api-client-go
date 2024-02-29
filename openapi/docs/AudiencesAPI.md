# \AudiencesAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResolveAudiencesOfClient**](AudiencesAPI.md#ResolveAudiencesOfClient) | **Get** /audiences | Get audiences of the API client
[**ResolveAudiencesOfMultiple**](AudiencesAPI.md#ResolveAudiencesOfMultiple) | **Post** /audiences | Get audiences for multiple IP addresses
[**ResolveAudiencesOfSingle**](AudiencesAPI.md#ResolveAudiencesOfSingle) | **Get** /audiences/{user_ip} | Get audiences for a given IP address



## ResolveAudiencesOfClient

> AudienceResponse ResolveAudiencesOfClient(ctx).Include(include).Type_(type_).Execute()

Get audiences of the API client

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
	include := "include_example" // string | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * `core` represents the core audiences that are directly linked to household characteristics   * `composite` represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * `name` and `category` refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  (optional) (default to "core,name,category")
	type_ := "jsonp" // string | Optional parameter to set to `jsonp` if a JSONP response format is needed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiencesAPI.ResolveAudiencesOfClient(context.Background()).Include(include).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ResolveAudiencesOfClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAudiencesOfClient`: AudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ResolveAudiencesOfClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveAudiencesOfClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * &#x60;core&#x60; represents the core audiences that are directly linked to household characteristics   * &#x60;composite&#x60; represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * &#x60;name&#x60; and &#x60;category&#x60; refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  | [default to &quot;core,name,category&quot;]
 **type_** | **string** | Optional parameter to set to &#x60;jsonp&#x60; if a JSONP response format is needed. | 

### Return type

[**AudienceResponse**](AudienceResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAudiencesOfMultiple

> ResolveAudiencesOfMultipleResponse ResolveAudiencesOfMultiple(ctx).ResolveAudiencesOfMultipleRequest(resolveAudiencesOfMultipleRequest).Include(include).Execute()

Get audiences for multiple IP addresses

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
	resolveAudiencesOfMultipleRequest := *openapiclient.NewResolveAudiencesOfMultipleRequest([]openapiclient.ResolveAudiencesOfMultipleRequestItem{*openapiclient.NewResolveAudiencesOfMultipleRequestItem("1.1.1.1")}) // ResolveAudiencesOfMultipleRequest | 
	include := "include_example" // string | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * `core` represents the core audiences that are directly linked to household characteristics   * `composite` represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * `name` and `category` refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  (optional) (default to "core,name,category")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiencesAPI.ResolveAudiencesOfMultiple(context.Background()).ResolveAudiencesOfMultipleRequest(resolveAudiencesOfMultipleRequest).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ResolveAudiencesOfMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAudiencesOfMultiple`: ResolveAudiencesOfMultipleResponse
	fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ResolveAudiencesOfMultiple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveAudiencesOfMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resolveAudiencesOfMultipleRequest** | [**ResolveAudiencesOfMultipleRequest**](ResolveAudiencesOfMultipleRequest.md) |  | 
 **include** | **string** | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * &#x60;core&#x60; represents the core audiences that are directly linked to household characteristics   * &#x60;composite&#x60; represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * &#x60;name&#x60; and &#x60;category&#x60; refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  | [default to &quot;core,name,category&quot;]

### Return type

[**ResolveAudiencesOfMultipleResponse**](ResolveAudiencesOfMultipleResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAudiencesOfSingle

> AudienceResponse ResolveAudiencesOfSingle(ctx, userIp).Include(include).Execute()

Get audiences for a given IP address

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
	userIp := "1.1.1.1" // string | The IP address to look up.
	include := "include_example" // string | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * `core` represents the core audiences that are directly linked to household characteristics   * `composite` represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * `name` and `category` refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  (optional) (default to "core,name,category")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiencesAPI.ResolveAudiencesOfSingle(context.Background(), userIp).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ResolveAudiencesOfSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAudiencesOfSingle`: AudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ResolveAudiencesOfSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIp** | **string** | The IP address to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAudiencesOfSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Optional parameter used to specify which audience information to be returned. The value is comprised of comma-separated values, each indicating a set of audiences:    * &#x60;core&#x60; represents the core audiences that are directly linked to household characteristics   * &#x60;composite&#x60; represents the composite audiences, used to model likely behaviours or buying     needs associated with the household characteristics.   * &#x60;name&#x60; and &#x60;category&#x60; refer to the fields of the same names in the returned Audience     objects. There is a slight performance gain in leaving these out when they are not needed.  | [default to &quot;core,name,category&quot;]

### Return type

[**AudienceResponse**](AudienceResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

