# SubscriptionPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Currency** | [**SubscriptionPriceCurrency**](SubscriptionPriceCurrency.md) |  | [default to EUR]
**Interval** | [**SubscriptionPriceInterval**](SubscriptionPriceInterval.md) |  | 

## Methods

### NewSubscriptionPrice

`func NewSubscriptionPrice(amount int32, currency SubscriptionPriceCurrency, interval SubscriptionPriceInterval, ) *SubscriptionPrice`

NewSubscriptionPrice instantiates a new SubscriptionPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPriceWithDefaults

`func NewSubscriptionPriceWithDefaults() *SubscriptionPrice`

NewSubscriptionPriceWithDefaults instantiates a new SubscriptionPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SubscriptionPrice) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionPrice) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionPrice) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *SubscriptionPrice) GetCurrency() SubscriptionPriceCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionPrice) GetCurrencyOk() (*SubscriptionPriceCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionPrice) SetCurrency(v SubscriptionPriceCurrency)`

SetCurrency sets Currency field to given value.


### GetInterval

`func (o *SubscriptionPrice) GetInterval() SubscriptionPriceInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SubscriptionPrice) GetIntervalOk() (*SubscriptionPriceInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SubscriptionPrice) SetInterval(v SubscriptionPriceInterval)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


