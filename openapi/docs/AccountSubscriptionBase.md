# AccountSubscriptionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the subscription | [readonly] 
**StartDate** | **string** |  | 
**EndDate** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** | Is the subscription currently active or not? | [readonly] 
**ActualPrice** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**PaymentConfiguration** | Pointer to [**AccountSubscriptionPaymentConfiguration**](AccountSubscriptionPaymentConfiguration.md) |  | [optional] 

## Methods

### NewAccountSubscriptionBase

`func NewAccountSubscriptionBase(id string, startDate string, isActive bool, ) *AccountSubscriptionBase`

NewAccountSubscriptionBase instantiates a new AccountSubscriptionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSubscriptionBaseWithDefaults

`func NewAccountSubscriptionBaseWithDefaults() *AccountSubscriptionBase`

NewAccountSubscriptionBaseWithDefaults instantiates a new AccountSubscriptionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSubscriptionBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSubscriptionBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSubscriptionBase) SetId(v string)`

SetId sets Id field to given value.


### GetStartDate

`func (o *AccountSubscriptionBase) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AccountSubscriptionBase) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AccountSubscriptionBase) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AccountSubscriptionBase) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AccountSubscriptionBase) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AccountSubscriptionBase) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AccountSubscriptionBase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActive

`func (o *AccountSubscriptionBase) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountSubscriptionBase) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountSubscriptionBase) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetActualPrice

`func (o *AccountSubscriptionBase) GetActualPrice() SubscriptionPrice`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *AccountSubscriptionBase) GetActualPriceOk() (*SubscriptionPrice, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *AccountSubscriptionBase) SetActualPrice(v SubscriptionPrice)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *AccountSubscriptionBase) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetPaymentConfiguration

`func (o *AccountSubscriptionBase) GetPaymentConfiguration() AccountSubscriptionPaymentConfiguration`

GetPaymentConfiguration returns the PaymentConfiguration field if non-nil, zero value otherwise.

### GetPaymentConfigurationOk

`func (o *AccountSubscriptionBase) GetPaymentConfigurationOk() (*AccountSubscriptionPaymentConfiguration, bool)`

GetPaymentConfigurationOk returns a tuple with the PaymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfiguration

`func (o *AccountSubscriptionBase) SetPaymentConfiguration(v AccountSubscriptionPaymentConfiguration)`

SetPaymentConfiguration sets PaymentConfiguration field to given value.

### HasPaymentConfiguration

`func (o *AccountSubscriptionBase) HasPaymentConfiguration() bool`

HasPaymentConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


