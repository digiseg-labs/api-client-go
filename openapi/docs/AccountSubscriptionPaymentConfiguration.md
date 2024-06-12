# AccountSubscriptionPaymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Timing** | **string** | Is the subscription pre-paid or post-paid? | 
**StripeSubscriptionId** | Pointer to **string** |  | [optional] [readonly] 
**StripeSubscriptionItemId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAccountSubscriptionPaymentConfiguration

`func NewAccountSubscriptionPaymentConfiguration(platform string, timing string, ) *AccountSubscriptionPaymentConfiguration`

NewAccountSubscriptionPaymentConfiguration instantiates a new AccountSubscriptionPaymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionPaymentConfigurationWithDefaults

`func NewAccountSubscriptionPaymentConfigurationWithDefaults() *AccountSubscriptionPaymentConfiguration`

NewAccountSubscriptionPaymentConfigurationWithDefaults instantiates a new AccountSubscriptionPaymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AccountSubscriptionPaymentConfiguration) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AccountSubscriptionPaymentConfiguration) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AccountSubscriptionPaymentConfiguration) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetTiming

`func (o *AccountSubscriptionPaymentConfiguration) GetTiming() string`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *AccountSubscriptionPaymentConfiguration) GetTimingOk() (*string, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *AccountSubscriptionPaymentConfiguration) SetTiming(v string)`

SetTiming sets Timing field to given value.


### GetStripeSubscriptionId

`func (o *AccountSubscriptionPaymentConfiguration) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *AccountSubscriptionPaymentConfiguration) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *AccountSubscriptionPaymentConfiguration) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *AccountSubscriptionPaymentConfiguration) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStripeSubscriptionItemId

`func (o *AccountSubscriptionPaymentConfiguration) GetStripeSubscriptionItemId() string`

GetStripeSubscriptionItemId returns the StripeSubscriptionItemId field if non-nil, zero value otherwise.

### GetStripeSubscriptionItemIdOk

`func (o *AccountSubscriptionPaymentConfiguration) GetStripeSubscriptionItemIdOk() (*string, bool)`

GetStripeSubscriptionItemIdOk returns a tuple with the StripeSubscriptionItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionItemId

`func (o *AccountSubscriptionPaymentConfiguration) SetStripeSubscriptionItemId(v string)`

SetStripeSubscriptionItemId sets StripeSubscriptionItemId field to given value.

### HasStripeSubscriptionItemId

`func (o *AccountSubscriptionPaymentConfiguration) HasStripeSubscriptionItemId() bool`

HasStripeSubscriptionItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


