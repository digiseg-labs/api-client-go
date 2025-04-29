# SubscriptionOfferFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
**AccountId** | **string** | The ID of the account that has the offer | 
**SubscriptionPlanId** | Pointer to **string** | The ID of the subscription plan that is offered (may have changed from the original, or may not exist yet) | [optional] 
**AcceptedBy** | Pointer to **string** | The ID of the user who accepted the offer | [optional] 
**AcceptedAt** | Pointer to **time.Time** | Date and time of the acceptance | [optional] 
**OriginalPlanId** | Pointer to **string** | The ID of the subscription plan that this offer was originally based on | [optional] 
**OfferedPrice** | [**SubscriptionPrice**](SubscriptionPrice.md) |  | 
**Plan** | Pointer to [**SubscriptionPlanFull**](SubscriptionPlanFull.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 

## Methods

### NewSubscriptionOfferFull

`func NewSubscriptionOfferFull(accountId string, offeredPrice SubscriptionPrice, ) *SubscriptionOfferFull`

NewSubscriptionOfferFull instantiates a new SubscriptionOfferFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOfferFullWithDefaults

`func NewSubscriptionOfferFullWithDefaults() *SubscriptionOfferFull`

NewSubscriptionOfferFullWithDefaults instantiates a new SubscriptionOfferFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionOfferFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionOfferFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionOfferFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionOfferFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *SubscriptionOfferFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubscriptionOfferFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubscriptionOfferFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSubscriptionPlanId

`func (o *SubscriptionOfferFull) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *SubscriptionOfferFull) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *SubscriptionOfferFull) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *SubscriptionOfferFull) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### GetAcceptedBy

`func (o *SubscriptionOfferFull) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *SubscriptionOfferFull) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *SubscriptionOfferFull) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *SubscriptionOfferFull) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *SubscriptionOfferFull) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *SubscriptionOfferFull) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *SubscriptionOfferFull) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *SubscriptionOfferFull) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetOriginalPlanId

`func (o *SubscriptionOfferFull) GetOriginalPlanId() string`

GetOriginalPlanId returns the OriginalPlanId field if non-nil, zero value otherwise.

### GetOriginalPlanIdOk

`func (o *SubscriptionOfferFull) GetOriginalPlanIdOk() (*string, bool)`

GetOriginalPlanIdOk returns a tuple with the OriginalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPlanId

`func (o *SubscriptionOfferFull) SetOriginalPlanId(v string)`

SetOriginalPlanId sets OriginalPlanId field to given value.

### HasOriginalPlanId

`func (o *SubscriptionOfferFull) HasOriginalPlanId() bool`

HasOriginalPlanId returns a boolean if a field has been set.

### GetOfferedPrice

`func (o *SubscriptionOfferFull) GetOfferedPrice() SubscriptionPrice`

GetOfferedPrice returns the OfferedPrice field if non-nil, zero value otherwise.

### GetOfferedPriceOk

`func (o *SubscriptionOfferFull) GetOfferedPriceOk() (*SubscriptionPrice, bool)`

GetOfferedPriceOk returns a tuple with the OfferedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedPrice

`func (o *SubscriptionOfferFull) SetOfferedPrice(v SubscriptionPrice)`

SetOfferedPrice sets OfferedPrice field to given value.


### GetPlan

`func (o *SubscriptionOfferFull) GetPlan() SubscriptionPlanFull`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionOfferFull) GetPlanOk() (*SubscriptionPlanFull, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionOfferFull) SetPlan(v SubscriptionPlanFull)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionOfferFull) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionOfferFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionOfferFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionOfferFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionOfferFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SubscriptionOfferFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubscriptionOfferFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubscriptionOfferFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SubscriptionOfferFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionOfferFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionOfferFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionOfferFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionOfferFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SubscriptionOfferFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SubscriptionOfferFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SubscriptionOfferFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SubscriptionOfferFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


