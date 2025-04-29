# UserAccountMembershipIncludeAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountFull**](AccountFull.md) |  | [optional] 
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | [optional] [readonly] 

## Methods

### NewUserAccountMembershipIncludeAux

`func NewUserAccountMembershipIncludeAux() *UserAccountMembershipIncludeAux`

NewUserAccountMembershipIncludeAux instantiates a new UserAccountMembershipIncludeAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountMembershipIncludeAuxWithDefaults

`func NewUserAccountMembershipIncludeAuxWithDefaults() *UserAccountMembershipIncludeAux`

NewUserAccountMembershipIncludeAuxWithDefaults instantiates a new UserAccountMembershipIncludeAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserAccountMembershipIncludeAux) GetAccount() AccountFull`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserAccountMembershipIncludeAux) GetAccountOk() (*AccountFull, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserAccountMembershipIncludeAux) SetAccount(v AccountFull)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserAccountMembershipIncludeAux) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFeatureSet

`func (o *UserAccountMembershipIncludeAux) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *UserAccountMembershipIncludeAux) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *UserAccountMembershipIncludeAux) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *UserAccountMembershipIncludeAux) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetSubscriptions

`func (o *UserAccountMembershipIncludeAux) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UserAccountMembershipIncludeAux) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UserAccountMembershipIncludeAux) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UserAccountMembershipIncludeAux) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


