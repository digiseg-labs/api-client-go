# UserAccountMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | 
**AccountName** | Pointer to **string** | The name of the account | [optional] 
**Roles** | [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user has within the account | 
**Account** | Pointer to [**AccountFull**](AccountFull.md) |  | [optional] 
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | [optional] [readonly] 

## Methods

### NewUserAccountMembership

`func NewUserAccountMembership(accountId string, roles []UserAccountRole, ) *UserAccountMembership`

NewUserAccountMembership instantiates a new UserAccountMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountMembershipWithDefaults

`func NewUserAccountMembershipWithDefaults() *UserAccountMembership`

NewUserAccountMembershipWithDefaults instantiates a new UserAccountMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserAccountMembership) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserAccountMembership) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserAccountMembership) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *UserAccountMembership) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UserAccountMembership) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UserAccountMembership) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *UserAccountMembership) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetRoles

`func (o *UserAccountMembership) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAccountMembership) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAccountMembership) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.


### GetAccount

`func (o *UserAccountMembership) GetAccount() AccountFull`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserAccountMembership) GetAccountOk() (*AccountFull, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserAccountMembership) SetAccount(v AccountFull)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserAccountMembership) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFeatureSet

`func (o *UserAccountMembership) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *UserAccountMembership) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *UserAccountMembership) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *UserAccountMembership) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetSubscriptions

`func (o *UserAccountMembership) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UserAccountMembership) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UserAccountMembership) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UserAccountMembership) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


