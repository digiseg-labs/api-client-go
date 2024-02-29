# UserBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the user (used as username when authenticating with password) | [optional] 
**Name** | Pointer to **string** | Human readable name of the user | [optional] 
**AccountId** | Pointer to **string** | ID of the account that this user pertains to. If the user has multiple account memberships, this account ID will represent the primary account of the user.  | [optional] 
**Roles** | Pointer to [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user has within the account | [optional] 
**AvatarUrl** | Pointer to **string** | The URL to an avatar of the user | [optional] 

## Methods

### NewUserBase

`func NewUserBase() *UserBase`

NewUserBase instantiates a new UserBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBaseWithDefaults

`func NewUserBaseWithDefaults() *UserBase`

NewUserBaseWithDefaults instantiates a new UserBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserBase) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBase) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBase) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserBase) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *UserBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserBase) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoles

`func (o *UserBase) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserBase) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserBase) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserBase) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserBase) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserBase) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserBase) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserBase) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


