# AccountIncludeAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | [optional] [readonly] 

## Methods

### NewAccountIncludeAux

`func NewAccountIncludeAux() *AccountIncludeAux`

NewAccountIncludeAux instantiates a new AccountIncludeAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIncludeAuxWithDefaults

`func NewAccountIncludeAuxWithDefaults() *AccountIncludeAux`

NewAccountIncludeAuxWithDefaults instantiates a new AccountIncludeAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureSet

`func (o *AccountIncludeAux) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *AccountIncludeAux) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *AccountIncludeAux) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *AccountIncludeAux) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetSubscriptions

`func (o *AccountIncludeAux) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AccountIncludeAux) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AccountIncludeAux) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AccountIncludeAux) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


