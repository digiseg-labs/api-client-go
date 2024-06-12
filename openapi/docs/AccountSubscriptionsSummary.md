# AccountSubscriptionsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | [readonly] 
**Subscriptions** | [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | 
**FeatureSet** | [**PlanFeatureSet**](PlanFeatureSet.md) |  | 

## Methods

### NewAccountSubscriptionsSummary

`func NewAccountSubscriptionsSummary(accountId string, subscriptions []AccountSubscriptionItem, featureSet PlanFeatureSet, ) *AccountSubscriptionsSummary`

NewAccountSubscriptionsSummary instantiates a new AccountSubscriptionsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionsSummaryWithDefaults

`func NewAccountSubscriptionsSummaryWithDefaults() *AccountSubscriptionsSummary`

NewAccountSubscriptionsSummaryWithDefaults instantiates a new AccountSubscriptionsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountSubscriptionsSummary) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountSubscriptionsSummary) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountSubscriptionsSummary) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSubscriptions

`func (o *AccountSubscriptionsSummary) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AccountSubscriptionsSummary) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AccountSubscriptionsSummary) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.


### GetFeatureSet

`func (o *AccountSubscriptionsSummary) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *AccountSubscriptionsSummary) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *AccountSubscriptionsSummary) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


