# AccountUserAddition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user (used as username when authenticating with password) | 
**Name** | Pointer to **string** | Human readable name of the user | [optional] 
**Roles** | Pointer to [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user will have within the account | [optional] 

## Methods

### NewAccountUserAddition

`func NewAccountUserAddition(email string, ) *AccountUserAddition`

NewAccountUserAddition instantiates a new AccountUserAddition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUserAdditionWithDefaults

`func NewAccountUserAdditionWithDefaults() *AccountUserAddition`

NewAccountUserAdditionWithDefaults instantiates a new AccountUserAddition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountUserAddition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountUserAddition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountUserAddition) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *AccountUserAddition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUserAddition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUserAddition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountUserAddition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *AccountUserAddition) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountUserAddition) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountUserAddition) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountUserAddition) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


