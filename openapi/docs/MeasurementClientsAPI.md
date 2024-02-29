# \MeasurementClientsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeasurementClient**](MeasurementClientsAPI.md#CreateMeasurementClient) | **Post** /measurement/clients | Create measurement client
[**DeleteClientLogo**](MeasurementClientsAPI.md#DeleteClientLogo) | **Delete** /measurement/clients/{client_id}/assets/logo | Delete measurement client logo
[**DeleteMeasurementClientById**](MeasurementClientsAPI.md#DeleteMeasurementClientById) | **Delete** /measurement/clients/{client_id} | Delete measurement client
[**GetClientLogo**](MeasurementClientsAPI.md#GetClientLogo) | **Get** /measurement/clients/{client_id}/assets/logo | Get measurement client logo
[**GetMeasurementClientById**](MeasurementClientsAPI.md#GetMeasurementClientById) | **Get** /measurement/clients/{client_id} | Get measurement client
[**ListMeasurementClients**](MeasurementClientsAPI.md#ListMeasurementClients) | **Get** /measurement/clients | List measurement clients
[**UpdateMeasurementClientById**](MeasurementClientsAPI.md#UpdateMeasurementClientById) | **Put** /measurement/clients/{client_id} | Update measurement client
[**UploadClientLogo**](MeasurementClientsAPI.md#UploadClientLogo) | **Put** /measurement/clients/{client_id}/assets/logo | Upload measurement client logo



## CreateMeasurementClient

> CreateMeasurementClient201Response CreateMeasurementClient(ctx).MeasurementClientMutation(measurementClientMutation).Execute()

Create measurement client



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
	measurementClientMutation := *openapiclient.NewMeasurementClientMutation() // MeasurementClientMutation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeasurementClientsAPI.CreateMeasurementClient(context.Background()).MeasurementClientMutation(measurementClientMutation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.CreateMeasurementClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMeasurementClient`: CreateMeasurementClient201Response
	fmt.Fprintf(os.Stdout, "Response from `MeasurementClientsAPI.CreateMeasurementClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMeasurementClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **measurementClientMutation** | [**MeasurementClientMutation**](MeasurementClientMutation.md) |  | 

### Return type

[**CreateMeasurementClient201Response**](CreateMeasurementClient201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientLogo

> DeleteClientLogo(ctx, clientId).Execute()

Delete measurement client logo



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
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeasurementClientsAPI.DeleteClientLogo(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.DeleteClientLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientLogoRequest struct via the builder pattern


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


## DeleteMeasurementClientById

> DeleteMeasurementClientById(ctx, clientId).Execute()

Delete measurement client



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
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeasurementClientsAPI.DeleteMeasurementClientById(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.DeleteMeasurementClientById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMeasurementClientByIdRequest struct via the builder pattern


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


## GetClientLogo

> GetClientLogo(ctx, clientId).Execute()

Get measurement client logo



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
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeasurementClientsAPI.GetClientLogo(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.GetClientLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientLogoRequest struct via the builder pattern


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


## GetMeasurementClientById

> CreateMeasurementClient201Response GetMeasurementClientById(ctx, clientId).Execute()

Get measurement client



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
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeasurementClientsAPI.GetMeasurementClientById(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.GetMeasurementClientById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeasurementClientById`: CreateMeasurementClient201Response
	fmt.Fprintf(os.Stdout, "Response from `MeasurementClientsAPI.GetMeasurementClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeasurementClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMeasurementClient201Response**](CreateMeasurementClient201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMeasurementClients

> ListMeasurementClients200Response ListMeasurementClients(ctx).FilterAccountId(filterAccountId).FilterNameContains(filterNameContains).PageSize(pageSize).PageAfter(pageAfter).Execute()

List measurement clients



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
	filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query measurement clients by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")
	filterNameContains := "acme" // string | Optional parameter used to search for clients where the name contains a substring (case insensitive) (optional)
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeasurementClientsAPI.ListMeasurementClients(context.Background()).FilterAccountId(filterAccountId).FilterNameContains(filterNameContains).PageSize(pageSize).PageAfter(pageAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.ListMeasurementClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeasurementClients`: ListMeasurementClients200Response
	fmt.Fprintf(os.Stdout, "Response from `MeasurementClientsAPI.ListMeasurementClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMeasurementClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAccountId** | **string** | Optional parameter used to query measurement clients by specific account IDs (only available to super admins). The value &#x60;*&#x60; is synonymous for \&quot;all accounts\&quot;.  | [default to &quot;The user&#39;s account ID&quot;]
 **filterNameContains** | **string** | Optional parameter used to search for clients where the name contains a substring (case insensitive) | 
 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 

### Return type

[**ListMeasurementClients200Response**](ListMeasurementClients200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeasurementClientById

> CreateMeasurementClient201Response UpdateMeasurementClientById(ctx, clientId).MeasurementClientMutation(measurementClientMutation).Execute()

Update measurement client



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
	clientId := "clientId_example" // string | 
	measurementClientMutation := *openapiclient.NewMeasurementClientMutation() // MeasurementClientMutation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeasurementClientsAPI.UpdateMeasurementClientById(context.Background(), clientId).MeasurementClientMutation(measurementClientMutation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.UpdateMeasurementClientById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeasurementClientById`: CreateMeasurementClient201Response
	fmt.Fprintf(os.Stdout, "Response from `MeasurementClientsAPI.UpdateMeasurementClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeasurementClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **measurementClientMutation** | [**MeasurementClientMutation**](MeasurementClientMutation.md) |  | 

### Return type

[**CreateMeasurementClient201Response**](CreateMeasurementClient201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadClientLogo

> UploadClientLogo(ctx, clientId).Body(body).Execute()

Upload measurement client logo



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
	clientId := "clientId_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeasurementClientsAPI.UploadClientLogo(context.Background(), clientId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeasurementClientsAPI.UploadClientLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientLogoRequest struct via the builder pattern


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

