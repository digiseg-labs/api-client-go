# AccountSubscriptionItem

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
**Plan** | Pointer to [**SubscriptionPlanItem**](SubscriptionPlanItem.md) |  | [optional] 

## Methods

### NewAccountSubscriptionItem

`func NewAccountSubscriptionItem(id string, startDate string, isActive bool, ) *AccountSubscriptionItem`

NewAccountSubscriptionItem instantiates a new AccountSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionItemWithDefaults

`func NewAccountSubscriptionItemWithDefaults() *AccountSubscriptionItem`

NewAccountSubscriptionItemWithDefaults instantiates a new AccountSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSubscriptionItem) SetId(v string)`

SetId sets Id field to given value.


### GetStartDate

`func (o *AccountSubscriptionItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AccountSubscriptionItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AccountSubscriptionItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AccountSubscriptionItem) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AccountSubscriptionItem) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AccountSubscriptionItem) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AccountSubscriptionItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActive

`func (o *AccountSubscriptionItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountSubscriptionItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountSubscriptionItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCancelledAt

`func (o *AccountSubscriptionItem) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *AccountSubscriptionItem) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *AccountSubscriptionItem) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *AccountSubscriptionItem) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetActualPrice

`func (o *AccountSubscriptionItem) GetActualPrice() SubscriptionPrice`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *AccountSubscriptionItem) GetActualPriceOk() (*SubscriptionPrice, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *AccountSubscriptionItem) SetActualPrice(v SubscriptionPrice)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *AccountSubscriptionItem) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetPaymentConfiguration

`func (o *AccountSubscriptionItem) GetPaymentConfiguration() AccountSubscriptionPaymentConfiguration`

GetPaymentConfiguration returns the PaymentConfiguration field if non-nil, zero value otherwise.

### GetPaymentConfigurationOk

`func (o *AccountSubscriptionItem) GetPaymentConfigurationOk() (*AccountSubscriptionPaymentConfiguration, bool)`

GetPaymentConfigurationOk returns a tuple with the PaymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfiguration

`func (o *AccountSubscriptionItem) SetPaymentConfiguration(v AccountSubscriptionPaymentConfiguration)`

SetPaymentConfiguration sets PaymentConfiguration field to given value.

### HasPaymentConfiguration

`func (o *AccountSubscriptionItem) HasPaymentConfiguration() bool`

HasPaymentConfiguration returns a boolean if a field has been set.

### GetPlan

`func (o *AccountSubscriptionItem) GetPlan() SubscriptionPlanItem`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AccountSubscriptionItem) GetPlanOk() (*SubscriptionPlanItem, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AccountSubscriptionItem) SetPlan(v SubscriptionPlanItem)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AccountSubscriptionItem) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


