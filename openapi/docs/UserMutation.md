# UserMutation

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
**Password** | Pointer to **string** | Password of the user | [optional] 

## Methods

### NewUserMutation

`func NewUserMutation() *UserMutation`

NewUserMutation instantiates a new UserMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMutationWithDefaults

`func NewUserMutationWithDefaults() *UserMutation`

NewUserMutationWithDefaults instantiates a new UserMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserMutation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserMutation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserMutation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserMutation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *UserMutation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserMutation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserMutation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserMutation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoles

`func (o *UserMutation) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserMutation) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserMutation) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserMutation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserMutation) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserMutation) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserMutation) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserMutation) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetLoggedInAt

`func (o *UserMutation) GetLoggedInAt() time.Time`

GetLoggedInAt returns the LoggedInAt field if non-nil, zero value otherwise.

### GetLoggedInAtOk

`func (o *UserMutation) GetLoggedInAtOk() (*time.Time, bool)`

GetLoggedInAtOk returns a tuple with the LoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInAt

`func (o *UserMutation) SetLoggedInAt(v time.Time)`

SetLoggedInAt sets LoggedInAt field to given value.

### HasLoggedInAt

`func (o *UserMutation) HasLoggedInAt() bool`

HasLoggedInAt returns a boolean if a field has been set.

### GetAccountMemberships

`func (o *UserMutation) GetAccountMemberships() []UserAccountMembership`

GetAccountMemberships returns the AccountMemberships field if non-nil, zero value otherwise.

### GetAccountMembershipsOk

`func (o *UserMutation) GetAccountMembershipsOk() (*[]UserAccountMembership, bool)`

GetAccountMembershipsOk returns a tuple with the AccountMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMemberships

`func (o *UserMutation) SetAccountMemberships(v []UserAccountMembership)`

SetAccountMemberships sets AccountMemberships field to given value.

### HasAccountMemberships

`func (o *UserMutation) HasAccountMemberships() bool`

HasAccountMemberships returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *UserMutation) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *UserMutation) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *UserMutation) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *UserMutation) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetPassword

`func (o *UserMutation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserMutation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserMutation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserMutation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


