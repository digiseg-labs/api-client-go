# SubscriptionOfferBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account that has the offer | 
**SubscriptionPlanId** | Pointer to **string** | The ID of the subscription plan that is offered (may have changed from the original, or may not exist yet) | [optional] 
**AcceptedBy** | Pointer to **string** | The ID of the user who accepted the offer | [optional] 
**AcceptedAt** | Pointer to **time.Time** | Date and time of the acceptance | [optional] 
**OriginalPlanId** | Pointer to **string** | The ID of the subscription plan that this offer was originally based on | [optional] 
**OfferedPrice** | [**SubscriptionPrice**](SubscriptionPrice.md) |  | 

## Methods

### NewSubscriptionOfferBase

`func NewSubscriptionOfferBase(accountId string, offeredPrice SubscriptionPrice, ) *SubscriptionOfferBase`

NewSubscriptionOfferBase instantiates a new SubscriptionOfferBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOfferBaseWithDefaults

`func NewSubscriptionOfferBaseWithDefaults() *SubscriptionOfferBase`

NewSubscriptionOfferBaseWithDefaults instantiates a new SubscriptionOfferBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SubscriptionOfferBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubscriptionOfferBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubscriptionOfferBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSubscriptionPlanId

`func (o *SubscriptionOfferBase) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *SubscriptionOfferBase) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *SubscriptionOfferBase) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *SubscriptionOfferBase) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### GetAcceptedBy

`func (o *SubscriptionOfferBase) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *SubscriptionOfferBase) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *SubscriptionOfferBase) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *SubscriptionOfferBase) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *SubscriptionOfferBase) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *SubscriptionOfferBase) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *SubscriptionOfferBase) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *SubscriptionOfferBase) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetOriginalPlanId

`func (o *SubscriptionOfferBase) GetOriginalPlanId() string`

GetOriginalPlanId returns the OriginalPlanId field if non-nil, zero value otherwise.

### GetOriginalPlanIdOk

`func (o *SubscriptionOfferBase) GetOriginalPlanIdOk() (*string, bool)`

GetOriginalPlanIdOk returns a tuple with the OriginalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPlanId

`func (o *SubscriptionOfferBase) SetOriginalPlanId(v string)`

SetOriginalPlanId sets OriginalPlanId field to given value.

### HasOriginalPlanId

`func (o *SubscriptionOfferBase) HasOriginalPlanId() bool`

HasOriginalPlanId returns a boolean if a field has been set.

### GetOfferedPrice

`func (o *SubscriptionOfferBase) GetOfferedPrice() SubscriptionPrice`

GetOfferedPrice returns the OfferedPrice field if non-nil, zero value otherwise.

### GetOfferedPriceOk

`func (o *SubscriptionOfferBase) GetOfferedPriceOk() (*SubscriptionPrice, bool)`

GetOfferedPriceOk returns a tuple with the OfferedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedPrice

`func (o *SubscriptionOfferBase) SetOfferedPrice(v SubscriptionPrice)`

SetOfferedPrice sets OfferedPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


