# UserAccountMembershipUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]UserAccountRole**](UserAccountRole.md) | The roles to apply to the user within the account | 

## Methods

### NewUserAccountMembershipUpdate

`func NewUserAccountMembershipUpdate(roles []UserAccountRole, ) *UserAccountMembershipUpdate`

NewUserAccountMembershipUpdate instantiates a new UserAccountMembershipUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountMembershipUpdateWithDefaults

`func NewUserAccountMembershipUpdateWithDefaults() *UserAccountMembershipUpdate`

NewUserAccountMembershipUpdateWithDefaults instantiates a new UserAccountMembershipUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *UserAccountMembershipUpdate) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAccountMembershipUpdate) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAccountMembershipUpdate) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


