# \SubscriptionsAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountSubscriptionById**](SubscriptionsAPI.md#GetAccountSubscriptionById) | **Get** /accounts/{account_id}/subscriptions/{subscription_id} | Get account subscription
[**GetAccountSubscriptions**](SubscriptionsAPI.md#GetAccountSubscriptions) | **Get** /accounts/{account_id}/subscriptions | Get account subscriptions summary
[**ListSubscriptionOffers**](SubscriptionsAPI.md#ListSubscriptionOffers) | **Get** /subscription_offers | List subscription offers
[**ListSubscriptionPlans**](SubscriptionsAPI.md#ListSubscriptionPlans) | **Get** /subscription_plans | List subscription plans



## GetAccountSubscriptionById

> GetAccountSubscriptionById200Response GetAccountSubscriptionById(ctx, accountId, subscriptionId).Execute()

Get account subscription



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
	subscriptionId := "subscriptionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetAccountSubscriptionById(context.Background(), accountId, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetAccountSubscriptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubscriptionById`: GetAccountSubscriptionById200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetAccountSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubscriptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAccountSubscriptionById200Response**](GetAccountSubscriptionById200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSubscriptions

> GetAccountSubscriptions200Response GetAccountSubscriptions(ctx, accountId).Execute()

Get account subscriptions summary



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
	resp, r, err := apiClient.SubscriptionsAPI.GetAccountSubscriptions(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetAccountSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubscriptions`: GetAccountSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetAccountSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountSubscriptions200Response**](GetAccountSubscriptions200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionOffers

> ListSubscriptionOffers200Response ListSubscriptionOffers(ctx).FilterAccountId(filterAccountId).Execute()

List subscription offers

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
	filterAccountId := "filterAccountId_example" // string | Optional parameter used to filter on the account ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.ListSubscriptionOffers(context.Background()).FilterAccountId(filterAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ListSubscriptionOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubscriptionOffers`: ListSubscriptionOffers200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ListSubscriptionOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAccountId** | **string** | Optional parameter used to filter on the account ID | 

### Return type

[**ListSubscriptionOffers200Response**](ListSubscriptionOffers200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionPlans

> ListSubscriptionPlans200Response ListSubscriptionPlans(ctx).PageSize(pageSize).PageAfter(pageAfter).FilterIsPublic(filterIsPublic).FilterProductType(filterProductType).FilterListPriceCurrency(filterListPriceCurrency).FilterListPriceInterval(filterListPriceInterval).Execute()

List subscription plans

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
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)
	filterIsPublic := true // bool | Optional parameter used to filter on the `is_public` field (optional)
	filterProductType := openapiclient.SubscriptionProductType("base") // SubscriptionProductType | Optional parameter used to filter on the `product_type` field (optional)
	filterListPriceCurrency := openapiclient.SubscriptionPriceCurrency("EUR") // SubscriptionPriceCurrency | Optional parameter used to filter on the list price's currency (optional) (default to "EUR")
	filterListPriceInterval := openapiclient.SubscriptionPriceInterval("daily") // SubscriptionPriceInterval | Optional parameter used to filter on the list price's interval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.ListSubscriptionPlans(context.Background()).PageSize(pageSize).PageAfter(pageAfter).FilterIsPublic(filterIsPublic).FilterProductType(filterProductType).FilterListPriceCurrency(filterListPriceCurrency).FilterListPriceInterval(filterListPriceInterval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ListSubscriptionPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubscriptionPlans`: ListSubscriptionPlans200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ListSubscriptionPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 
 **filterIsPublic** | **bool** | Optional parameter used to filter on the &#x60;is_public&#x60; field | 
 **filterProductType** | [**SubscriptionProductType**](SubscriptionProductType.md) | Optional parameter used to filter on the &#x60;product_type&#x60; field | 
 **filterListPriceCurrency** | [**SubscriptionPriceCurrency**](SubscriptionPriceCurrency.md) | Optional parameter used to filter on the list price&#39;s currency | [default to &quot;EUR&quot;]
 **filterListPriceInterval** | [**SubscriptionPriceInterval**](SubscriptionPriceInterval.md) | Optional parameter used to filter on the list price&#39;s interval | 

### Return type

[**ListSubscriptionPlans200Response**](ListSubscriptionPlans200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

