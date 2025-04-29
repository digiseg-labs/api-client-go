# StripeAccountSubscriptionCheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | The URL to take the user to, to set up payment | [optional] 

## Methods

### NewStripeAccountSubscriptionCheckoutSession

`func NewStripeAccountSubscriptionCheckoutSession() *StripeAccountSubscriptionCheckoutSession`

NewStripeAccountSubscriptionCheckoutSession instantiates a new StripeAccountSubscriptionCheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeAccountSubscriptionCheckoutSessionWithDefaults

`func NewStripeAccountSubscriptionCheckoutSessionWithDefaults() *StripeAccountSubscriptionCheckoutSession`

NewStripeAccountSubscriptionCheckoutSessionWithDefaults instantiates a new StripeAccountSubscriptionCheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *StripeAccountSubscriptionCheckoutSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *StripeAccountSubscriptionCheckoutSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *StripeAccountSubscriptionCheckoutSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *StripeAccountSubscriptionCheckoutSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *StripeAccountSubscriptionCheckoutSession) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *StripeAccountSubscriptionCheckoutSession) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *StripeAccountSubscriptionCheckoutSession) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *StripeAccountSubscriptionCheckoutSession) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


