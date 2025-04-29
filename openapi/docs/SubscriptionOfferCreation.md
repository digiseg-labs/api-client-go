# SubscriptionOfferCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPlanId** | **string** | The ID of the subscription plan | 
**AccountId** | **string** | The ID of the account that has the offer | 
**OfferedPrice** | [**SubscriptionPrice**](SubscriptionPrice.md) |  | 

## Methods

### NewSubscriptionOfferCreation

`func NewSubscriptionOfferCreation(originalPlanId string, accountId string, offeredPrice SubscriptionPrice, ) *SubscriptionOfferCreation`

NewSubscriptionOfferCreation instantiates a new SubscriptionOfferCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOfferCreationWithDefaults

`func NewSubscriptionOfferCreationWithDefaults() *SubscriptionOfferCreation`

NewSubscriptionOfferCreationWithDefaults instantiates a new SubscriptionOfferCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPlanId

`func (o *SubscriptionOfferCreation) GetOriginalPlanId() string`

GetOriginalPlanId returns the OriginalPlanId field if non-nil, zero value otherwise.

### GetOriginalPlanIdOk

`func (o *SubscriptionOfferCreation) GetOriginalPlanIdOk() (*string, bool)`

GetOriginalPlanIdOk returns a tuple with the OriginalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPlanId

`func (o *SubscriptionOfferCreation) SetOriginalPlanId(v string)`

SetOriginalPlanId sets OriginalPlanId field to given value.


### GetAccountId

`func (o *SubscriptionOfferCreation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubscriptionOfferCreation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubscriptionOfferCreation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetOfferedPrice

`func (o *SubscriptionOfferCreation) GetOfferedPrice() SubscriptionPrice`

GetOfferedPrice returns the OfferedPrice field if non-nil, zero value otherwise.

### GetOfferedPriceOk

`func (o *SubscriptionOfferCreation) GetOfferedPriceOk() (*SubscriptionPrice, bool)`

GetOfferedPriceOk returns a tuple with the OfferedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedPrice

`func (o *SubscriptionOfferCreation) SetOfferedPrice(v SubscriptionPrice)`

SetOfferedPrice sets OfferedPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


