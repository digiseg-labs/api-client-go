# AccountSubscriptionFullAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | [readonly] 
**Plan** | [**SubscriptionPlanFull**](SubscriptionPlanFull.md) |  | 
**LastPaidAt** | Pointer to **time.Time** | Timestamp of the last processed payment | [optional] 

## Methods

### NewAccountSubscriptionFullAux

`func NewAccountSubscriptionFullAux(accountId string, plan SubscriptionPlanFull, ) *AccountSubscriptionFullAux`

NewAccountSubscriptionFullAux instantiates a new AccountSubscriptionFullAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionFullAuxWithDefaults

`func NewAccountSubscriptionFullAuxWithDefaults() *AccountSubscriptionFullAux`

NewAccountSubscriptionFullAuxWithDefaults instantiates a new AccountSubscriptionFullAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountSubscriptionFullAux) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountSubscriptionFullAux) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountSubscriptionFullAux) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPlan

`func (o *AccountSubscriptionFullAux) GetPlan() SubscriptionPlanFull`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AccountSubscriptionFullAux) GetPlanOk() (*SubscriptionPlanFull, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AccountSubscriptionFullAux) SetPlan(v SubscriptionPlanFull)`

SetPlan sets Plan field to given value.


### GetLastPaidAt

`func (o *AccountSubscriptionFullAux) GetLastPaidAt() time.Time`

GetLastPaidAt returns the LastPaidAt field if non-nil, zero value otherwise.

### GetLastPaidAtOk

`func (o *AccountSubscriptionFullAux) GetLastPaidAtOk() (*time.Time, bool)`

GetLastPaidAtOk returns a tuple with the LastPaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaidAt

`func (o *AccountSubscriptionFullAux) SetLastPaidAt(v time.Time)`

SetLastPaidAt sets LastPaidAt field to given value.

### HasLastPaidAt

`func (o *AccountSubscriptionFullAux) HasLastPaidAt() bool`

HasLastPaidAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


