# UserCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the user (used as username when authenticating with password) | [optional] 
**Name** | Pointer to **string** | Human readable name of the user | [optional] 
**AccountId** | Pointer to **string** | ID of the account that this user pertains to. If the user has multiple account memberships, this account ID will represent the primary account of the user.  | [optional] 
**Roles** | Pointer to [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user has within the account | [optional] 
**AvatarUrl** | Pointer to **string** | The URL to an avatar of the user | [optional] 
**LoggedInAt** | Pointer to **time.Time** | The approximate last time that the user logged in | [optional] [readonly] 
**AccountMemberships** | Pointer to [**[]UserAccountMembership**](UserAccountMembership.md) |  | [optional] [readonly] 
**IsSuperAdmin** | Pointer to **bool** | Determines if the user is a super admin of Digiseg API services | [optional] 
**PlatformRoles** | Pointer to [**[]UserPlatformRole**](UserPlatformRole.md) |  | [optional] 
**Password** | Pointer to **string** | Password of the user | [optional] 
**NotifyUser** | Pointer to **bool** | Whether or not to notify the user that they have been registered | [optional] [default to true]

## Methods

### NewUserCreation

`func NewUserCreation() *UserCreation`

NewUserCreation instantiates a new UserCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreationWithDefaults

`func NewUserCreationWithDefaults() *UserCreation`

NewUserCreationWithDefaults instantiates a new UserCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserCreation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserCreation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *UserCreation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserCreation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserCreation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserCreation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoles

`func (o *UserCreation) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreation) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreation) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCreation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserCreation) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserCreation) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserCreation) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserCreation) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetLoggedInAt

`func (o *UserCreation) GetLoggedInAt() time.Time`

GetLoggedInAt returns the LoggedInAt field if non-nil, zero value otherwise.

### GetLoggedInAtOk

`func (o *UserCreation) GetLoggedInAtOk() (*time.Time, bool)`

GetLoggedInAtOk returns a tuple with the LoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInAt

`func (o *UserCreation) SetLoggedInAt(v time.Time)`

SetLoggedInAt sets LoggedInAt field to given value.

### HasLoggedInAt

`func (o *UserCreation) HasLoggedInAt() bool`

HasLoggedInAt returns a boolean if a field has been set.

### GetAccountMemberships

`func (o *UserCreation) GetAccountMemberships() []UserAccountMembership`

GetAccountMemberships returns the AccountMemberships field if non-nil, zero value otherwise.

### GetAccountMembershipsOk

`func (o *UserCreation) GetAccountMembershipsOk() (*[]UserAccountMembership, bool)`

GetAccountMembershipsOk returns a tuple with the AccountMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMemberships

`func (o *UserCreation) SetAccountMemberships(v []UserAccountMembership)`

SetAccountMemberships sets AccountMemberships field to given value.

### HasAccountMemberships

`func (o *UserCreation) HasAccountMemberships() bool`

HasAccountMemberships returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *UserCreation) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *UserCreation) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *UserCreation) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *UserCreation) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetPlatformRoles

`func (o *UserCreation) GetPlatformRoles() []UserPlatformRole`

GetPlatformRoles returns the PlatformRoles field if non-nil, zero value otherwise.

### GetPlatformRolesOk

`func (o *UserCreation) GetPlatformRolesOk() (*[]UserPlatformRole, bool)`

GetPlatformRolesOk returns a tuple with the PlatformRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformRoles

`func (o *UserCreation) SetPlatformRoles(v []UserPlatformRole)`

SetPlatformRoles sets PlatformRoles field to given value.

### HasPlatformRoles

`func (o *UserCreation) HasPlatformRoles() bool`

HasPlatformRoles returns a boolean if a field has been set.

### GetPassword

`func (o *UserCreation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNotifyUser

`func (o *UserCreation) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *UserCreation) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *UserCreation) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *UserCreation) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


