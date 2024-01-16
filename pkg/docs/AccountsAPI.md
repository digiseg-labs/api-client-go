# \AccountsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserInAccount**](AccountsAPI.md#CreateUserInAccount) | **Post** /accounts/{account_id}/users | Create user
[**GetAccountById**](AccountsAPI.md#GetAccountById) | **Get** /accounts/{account_id} | Get account
[**GetUsersByAccountId**](AccountsAPI.md#GetUsersByAccountId) | **Get** /accounts/{account_id}/users | List users for account



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


## GetUsersByAccountId

> GetUsersByAccountId200Response GetUsersByAccountId(ctx, accountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

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
    resp, r, err := apiClient.AccountsAPI.GetUsersByAccountId(context.Background(), accountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetUsersByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersByAccountId`: GetUsersByAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetUsersByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 

### Return type

[**GetUsersByAccountId200Response**](GetUsersByAccountId200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

