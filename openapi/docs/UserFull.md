# UserFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
**Email** | Pointer to **string** | The email of the user (used as username when authenticating with password) | [optional] 
**Name** | Pointer to **string** | Human readable name of the user | [optional] 
**AccountId** | Pointer to **string** | ID of the account that this user pertains to. If the user has multiple account memberships, this account ID will represent the primary account of the user.  | [optional] [readonly] 
**Roles** | Pointer to [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user has within the account | [optional] [readonly] 
**AvatarUrl** | Pointer to **string** | The URL to an avatar of the user | [optional] 
**LoggedInAt** | Pointer to **time.Time** | The approximate last time that the user logged in | [optional] [readonly] 
**AccountMemberships** | Pointer to [**[]UserAccountMembership**](UserAccountMembership.md) |  | [optional] [readonly] 
**IsSuperAdmin** | Pointer to **bool** | Determines if the user is a super admin of Digiseg API services | [optional] 
**PlatformRoles** | Pointer to [**[]UserPlatformRole**](UserPlatformRole.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 

## Methods

### NewUserFull

`func NewUserFull() *UserFull`

NewUserFull instantiates a new UserFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFullWithDefaults

`func NewUserFullWithDefaults() *UserFull`

NewUserFullWithDefaults instantiates a new UserFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *UserFull) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserFull) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserFull) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserFull) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *UserFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserFull) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoles

`func (o *UserFull) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserFull) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserFull) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserFull) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserFull) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserFull) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserFull) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserFull) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetLoggedInAt

`func (o *UserFull) GetLoggedInAt() time.Time`

GetLoggedInAt returns the LoggedInAt field if non-nil, zero value otherwise.

### GetLoggedInAtOk

`func (o *UserFull) GetLoggedInAtOk() (*time.Time, bool)`

GetLoggedInAtOk returns a tuple with the LoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInAt

`func (o *UserFull) SetLoggedInAt(v time.Time)`

SetLoggedInAt sets LoggedInAt field to given value.

### HasLoggedInAt

`func (o *UserFull) HasLoggedInAt() bool`

HasLoggedInAt returns a boolean if a field has been set.

### GetAccountMemberships

`func (o *UserFull) GetAccountMemberships() []UserAccountMembership`

GetAccountMemberships returns the AccountMemberships field if non-nil, zero value otherwise.

### GetAccountMembershipsOk

`func (o *UserFull) GetAccountMembershipsOk() (*[]UserAccountMembership, bool)`

GetAccountMembershipsOk returns a tuple with the AccountMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMemberships

`func (o *UserFull) SetAccountMemberships(v []UserAccountMembership)`

SetAccountMemberships sets AccountMemberships field to given value.

### HasAccountMemberships

`func (o *UserFull) HasAccountMemberships() bool`

HasAccountMemberships returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *UserFull) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *UserFull) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *UserFull) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *UserFull) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetPlatformRoles

`func (o *UserFull) GetPlatformRoles() []UserPlatformRole`

GetPlatformRoles returns the PlatformRoles field if non-nil, zero value otherwise.

### GetPlatformRolesOk

`func (o *UserFull) GetPlatformRolesOk() (*[]UserPlatformRole, bool)`

GetPlatformRolesOk returns a tuple with the PlatformRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformRoles

`func (o *UserFull) SetPlatformRoles(v []UserPlatformRole)`

SetPlatformRoles sets PlatformRoles field to given value.

### HasPlatformRoles

`func (o *UserFull) HasPlatformRoles() bool`

HasPlatformRoles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UserFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UserFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *UserFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *UserFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *UserFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *UserFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


