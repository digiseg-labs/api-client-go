# AccountAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account. Deprecated in favor of the &#x60;owner&#x60; role of the user&#39;s account membership. | [optional] 
**BillingEmail** | Pointer to **string** | The email address to send billing information to. Requires &#x60;owner&#x60; role to change. | [optional] 
**BillingAddress** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 
**BillingTaxIds** | Pointer to [**[]TaxId**](TaxId.md) | A list of Tax IDs used by the account, for billing purposes. | [optional] [readonly] 
**BillingCurrency** | Pointer to [**SubscriptionPriceCurrency**](SubscriptionPriceCurrency.md) |  | [optional] [default to SUBSCRIPTIONPRICECURRENCY_EUR]
**BillingName** | Pointer to **string** | An optional official name to use for billing purposes. Requires &#x60;owner&#x60; role to change. | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAccountAux

`func NewAccountAux() *AccountAux`

NewAccountAux instantiates a new AccountAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAuxWithDefaults

`func NewAccountAuxWithDefaults() *AccountAux`

NewAccountAuxWithDefaults instantiates a new AccountAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *AccountAux) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountAux) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountAux) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountAux) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetBillingEmail

`func (o *AccountAux) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountAux) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *AccountAux) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *AccountAux) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *AccountAux) GetBillingAddress() PostalAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountAux) GetBillingAddressOk() (*PostalAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountAux) SetBillingAddress(v PostalAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountAux) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBillingTaxIds

`func (o *AccountAux) GetBillingTaxIds() []TaxId`

GetBillingTaxIds returns the BillingTaxIds field if non-nil, zero value otherwise.

### GetBillingTaxIdsOk

`func (o *AccountAux) GetBillingTaxIdsOk() (*[]TaxId, bool)`

GetBillingTaxIdsOk returns a tuple with the BillingTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTaxIds

`func (o *AccountAux) SetBillingTaxIds(v []TaxId)`

SetBillingTaxIds sets BillingTaxIds field to given value.

### HasBillingTaxIds

`func (o *AccountAux) HasBillingTaxIds() bool`

HasBillingTaxIds returns a boolean if a field has been set.

### GetBillingCurrency

`func (o *AccountAux) GetBillingCurrency() SubscriptionPriceCurrency`

GetBillingCurrency returns the BillingCurrency field if non-nil, zero value otherwise.

### GetBillingCurrencyOk

`func (o *AccountAux) GetBillingCurrencyOk() (*SubscriptionPriceCurrency, bool)`

GetBillingCurrencyOk returns a tuple with the BillingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCurrency

`func (o *AccountAux) SetBillingCurrency(v SubscriptionPriceCurrency)`

SetBillingCurrency sets BillingCurrency field to given value.

### HasBillingCurrency

`func (o *AccountAux) HasBillingCurrency() bool`

HasBillingCurrency returns a boolean if a field has been set.

### GetBillingName

`func (o *AccountAux) GetBillingName() string`

GetBillingName returns the BillingName field if non-nil, zero value otherwise.

### GetBillingNameOk

`func (o *AccountAux) GetBillingNameOk() (*string, bool)`

GetBillingNameOk returns a tuple with the BillingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingName

`func (o *AccountAux) SetBillingName(v string)`

SetBillingName sets BillingName field to given value.

### HasBillingName

`func (o *AccountAux) HasBillingName() bool`

HasBillingName returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *AccountAux) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *AccountAux) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *AccountAux) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *AccountAux) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


