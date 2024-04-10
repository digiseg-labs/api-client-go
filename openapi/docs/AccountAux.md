# AccountAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account. Deprecated in favor of the &#x60;owner&#x60; role of the user&#39;s account membership. | [optional] 
**BillingEmail** | Pointer to **string** | The email address to send billing information to. Requires &#x60;owner&#x60; role to change. | [optional] 
**BillingAddress** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


