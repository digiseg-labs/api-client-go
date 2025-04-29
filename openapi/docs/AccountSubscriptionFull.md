# AccountSubscriptionFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the subscription | [readonly] 
**StartDate** | **string** | Start date (inclusive) of the subscription | 
**EndDate** | Pointer to **string** | End date (inclusive) of the subscription | [optional] 
**IsActive** | **bool** | Is the subscription currently active or not? | [readonly] 
**CancelledAt** | Pointer to **time.Time** | Time of cancelling the subscription, if it has been cancelled. Note that a cancelled subscription may still be active, if it has been prepaid for the current period.  | [optional] [readonly] 
**ActualPrice** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**PaymentConfiguration** | Pointer to [**AccountSubscriptionPaymentConfiguration**](AccountSubscriptionPaymentConfiguration.md) |  | [optional] 
**Plan** | [**SubscriptionPlanFull**](SubscriptionPlanFull.md) |  | 
**AccountId** | **string** | The ID of the account | [readonly] 
**LastPaidAt** | Pointer to **time.Time** | Timestamp of the last processed payment | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 

## Methods

### NewAccountSubscriptionFull

`func NewAccountSubscriptionFull(id string, startDate string, isActive bool, plan SubscriptionPlanFull, accountId string, ) *AccountSubscriptionFull`

NewAccountSubscriptionFull instantiates a new AccountSubscriptionFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionFullWithDefaults

`func NewAccountSubscriptionFullWithDefaults() *AccountSubscriptionFull`

NewAccountSubscriptionFullWithDefaults instantiates a new AccountSubscriptionFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSubscriptionFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSubscriptionFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSubscriptionFull) SetId(v string)`

SetId sets Id field to given value.


### GetStartDate

`func (o *AccountSubscriptionFull) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AccountSubscriptionFull) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AccountSubscriptionFull) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AccountSubscriptionFull) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AccountSubscriptionFull) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AccountSubscriptionFull) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AccountSubscriptionFull) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActive

`func (o *AccountSubscriptionFull) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountSubscriptionFull) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountSubscriptionFull) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCancelledAt

`func (o *AccountSubscriptionFull) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *AccountSubscriptionFull) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *AccountSubscriptionFull) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *AccountSubscriptionFull) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetActualPrice

`func (o *AccountSubscriptionFull) GetActualPrice() SubscriptionPrice`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *AccountSubscriptionFull) GetActualPriceOk() (*SubscriptionPrice, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *AccountSubscriptionFull) SetActualPrice(v SubscriptionPrice)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *AccountSubscriptionFull) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetPaymentConfiguration

`func (o *AccountSubscriptionFull) GetPaymentConfiguration() AccountSubscriptionPaymentConfiguration`

GetPaymentConfiguration returns the PaymentConfiguration field if non-nil, zero value otherwise.

### GetPaymentConfigurationOk

`func (o *AccountSubscriptionFull) GetPaymentConfigurationOk() (*AccountSubscriptionPaymentConfiguration, bool)`

GetPaymentConfigurationOk returns a tuple with the PaymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfiguration

`func (o *AccountSubscriptionFull) SetPaymentConfiguration(v AccountSubscriptionPaymentConfiguration)`

SetPaymentConfiguration sets PaymentConfiguration field to given value.

### HasPaymentConfiguration

`func (o *AccountSubscriptionFull) HasPaymentConfiguration() bool`

HasPaymentConfiguration returns a boolean if a field has been set.

### GetPlan

`func (o *AccountSubscriptionFull) GetPlan() SubscriptionPlanFull`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AccountSubscriptionFull) GetPlanOk() (*SubscriptionPlanFull, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AccountSubscriptionFull) SetPlan(v SubscriptionPlanFull)`

SetPlan sets Plan field to given value.


### GetAccountId

`func (o *AccountSubscriptionFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountSubscriptionFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountSubscriptionFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLastPaidAt

`func (o *AccountSubscriptionFull) GetLastPaidAt() time.Time`

GetLastPaidAt returns the LastPaidAt field if non-nil, zero value otherwise.

### GetLastPaidAtOk

`func (o *AccountSubscriptionFull) GetLastPaidAtOk() (*time.Time, bool)`

GetLastPaidAtOk returns a tuple with the LastPaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaidAt

`func (o *AccountSubscriptionFull) SetLastPaidAt(v time.Time)`

SetLastPaidAt sets LastPaidAt field to given value.

### HasLastPaidAt

`func (o *AccountSubscriptionFull) HasLastPaidAt() bool`

HasLastPaidAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountSubscriptionFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountSubscriptionFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountSubscriptionFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountSubscriptionFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AccountSubscriptionFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountSubscriptionFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountSubscriptionFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccountSubscriptionFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountSubscriptionFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountSubscriptionFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountSubscriptionFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountSubscriptionFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AccountSubscriptionFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AccountSubscriptionFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AccountSubscriptionFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AccountSubscriptionFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


