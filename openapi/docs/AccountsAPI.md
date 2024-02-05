# \AccountsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserInAccount**](AccountsAPI.md#CreateUserInAccount) | **Post** /accounts/{account_id}/users | Create user
[**GetAccountById**](AccountsAPI.md#GetAccountById) | **Get** /accounts/{account_id} | Get account
[**ListUsersByAccountId**](AccountsAPI.md#ListUsersByAccountId) | **Get** /accounts/{account_id}/users | List users for account
[**UpdateAccountById**](AccountsAPI.md#UpdateAccountById) | **Put** /accounts/{account_id} | Update account



## CreateUserInAccount

> CreateUserInAccount201Response CreateUserInAccount(ctx, accountId).UserMutation(userMutation).Execute()

Create user



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
	accountId := "accountId_example" // string | 
	userMutation := *openapiclient.NewUserMutation() // UserMutation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateUserInAccount(context.Background(), accountId).UserMutation(userMutation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateUserInAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserInAccount`: CreateUserInAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateUserInAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserInAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userMutation** | [**UserMutation**](UserMutation.md) |  | 

### Return type

[**CreateUserInAccount201Response**](CreateUserInAccount201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountById

> GetAccountById200Response GetAccountById(ctx, accountId).Execute()

Get account

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
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountById(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountById`: GetAccountById200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountById200Response**](GetAccountById200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersByAccountId

> ListUsersByAccountId200Response ListUsersByAccountId(ctx, accountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

List users for account

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
	accountId := "accountId_example" // string | 
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListUsersByAccountId(context.Background(), accountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListUsersByAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersByAccountId`: ListUsersByAccountId200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListUsersByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 

### Return type

[**ListUsersByAccountId200Response**](ListUsersByAccountId200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountById

> GetAccountById200Response UpdateAccountById(ctx, accountId).AccountMutation(accountMutation).Execute()

Update account

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
	accountId := "accountId_example" // string | 
	accountMutation := *openapiclient.NewAccountMutation() // AccountMutation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateAccountById(context.Background(), accountId).AccountMutation(accountMutation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountById`: GetAccountById200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountMutation** | [**AccountMutation**](AccountMutation.md) |  | 

### Return type

[**GetAccountById200Response**](GetAccountById200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

