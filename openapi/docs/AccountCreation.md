# AccountCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account | [optional] 
**Owner** | Pointer to [**AccountOwnerCreation**](AccountOwnerCreation.md) |  | [optional] 

## Methods

### NewAccountCreation

`func NewAccountCreation() *AccountCreation`

NewAccountCreation instantiates a new AccountCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationWithDefaults

`func NewAccountCreationWithDefaults() *AccountCreation`

NewAccountCreationWithDefaults instantiates a new AccountCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCreation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *AccountCreation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountCreation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountCreation) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountCreation) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountCreation) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountCreation) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountCreation) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountCreation) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *AccountCreation) GetOwner() AccountOwnerCreation`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountCreation) GetOwnerOk() (*AccountOwnerCreation, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountCreation) SetOwner(v AccountOwnerCreation)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountCreation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


