# UserAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMemberships** | Pointer to [**[]UserAccountMembership**](UserAccountMembership.md) |  | [optional] [readonly] 
**IsSuperAdmin** | Pointer to **bool** | Determines if the user is a super admin of Digiseg API services | [optional] 
**PlatformRoles** | Pointer to [**[]UserPlatformRole**](UserPlatformRole.md) |  | [optional] 

## Methods

### NewUserAux

`func NewUserAux() *UserAux`

NewUserAux instantiates a new UserAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuxWithDefaults

`func NewUserAuxWithDefaults() *UserAux`

NewUserAuxWithDefaults instantiates a new UserAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMemberships

`func (o *UserAux) GetAccountMemberships() []UserAccountMembership`

GetAccountMemberships returns the AccountMemberships field if non-nil, zero value otherwise.

### GetAccountMembershipsOk

`func (o *UserAux) GetAccountMembershipsOk() (*[]UserAccountMembership, bool)`

GetAccountMembershipsOk returns a tuple with the AccountMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMemberships

`func (o *UserAux) SetAccountMemberships(v []UserAccountMembership)`

SetAccountMemberships sets AccountMemberships field to given value.

### HasAccountMemberships

`func (o *UserAux) HasAccountMemberships() bool`

HasAccountMemberships returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *UserAux) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *UserAux) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *UserAux) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *UserAux) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetPlatformRoles

`func (o *UserAux) GetPlatformRoles() []UserPlatformRole`

GetPlatformRoles returns the PlatformRoles field if non-nil, zero value otherwise.

### GetPlatformRolesOk

`func (o *UserAux) GetPlatformRolesOk() (*[]UserPlatformRole, bool)`

GetPlatformRolesOk returns a tuple with the PlatformRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformRoles

`func (o *UserAux) SetPlatformRoles(v []UserPlatformRole)`

SetPlatformRoles sets PlatformRoles field to given value.

### HasPlatformRoles

`func (o *UserAux) HasPlatformRoles() bool`

HasPlatformRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


