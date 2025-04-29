# AccountCreationAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**AccountOwnerCreation**](AccountOwnerCreation.md) |  | [optional] 
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account. Deprecated in favor of the &#x60;owner&#x60; role of the user&#39;s account membership. | [optional] 
**BillingEmail** | Pointer to **string** | The email address to send billing information to. Requires &#x60;owner&#x60; role to change. | [optional] 
**BillingAddress** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 

## Methods

### NewAccountCreationAux

`func NewAccountCreationAux() *AccountCreationAux`

NewAccountCreationAux instantiates a new AccountCreationAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationAuxWithDefaults

`func NewAccountCreationAuxWithDefaults() *AccountCreationAux`

NewAccountCreationAuxWithDefaults instantiates a new AccountCreationAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AccountCreationAux) GetOwner() AccountOwnerCreation`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountCreationAux) GetOwnerOk() (*AccountOwnerCreation, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountCreationAux) SetOwner(v AccountOwnerCreation)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountCreationAux) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountCreationAux) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountCreationAux) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountCreationAux) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountCreationAux) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetBillingEmail

`func (o *AccountCreationAux) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountCreationAux) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *AccountCreationAux) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *AccountCreationAux) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *AccountCreationAux) GetBillingAddress() PostalAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountCreationAux) GetBillingAddressOk() (*PostalAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountCreationAux) SetBillingAddress(v PostalAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountCreationAux) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


