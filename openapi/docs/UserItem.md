# UserItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
**Email** | Pointer to **string** | The email of the user (used as username when authenticating with password) | [optional] 
**Name** | Pointer to **string** | Human readable name of the user | [optional] 
**AccountId** | Pointer to **string** | ID of the account that this user pertains to. If the user has multiple account memberships, this account ID will represent the primary account of the user.  | [optional] 
**Roles** | Pointer to [**[]UserAccountRole**](UserAccountRole.md) | The roles that the user has within the account | [optional] 
**AvatarUrl** | Pointer to **string** | The URL to an avatar of the user | [optional] 
**LoggedInAt** | Pointer to **time.Time** | The approximate last time that the user logged in | [optional] [readonly] 

## Methods

### NewUserItem

`func NewUserItem() *UserItem`

NewUserItem instantiates a new UserItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserItemWithDefaults

`func NewUserItemWithDefaults() *UserItem`

NewUserItemWithDefaults instantiates a new UserItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *UserItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *UserItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoles

`func (o *UserItem) GetRoles() []UserAccountRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserItem) GetRolesOk() (*[]UserAccountRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserItem) SetRoles(v []UserAccountRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserItem) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserItem) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserItem) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserItem) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserItem) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetLoggedInAt

`func (o *UserItem) GetLoggedInAt() time.Time`

GetLoggedInAt returns the LoggedInAt field if non-nil, zero value otherwise.

### GetLoggedInAtOk

`func (o *UserItem) GetLoggedInAtOk() (*time.Time, bool)`

GetLoggedInAtOk returns a tuple with the LoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInAt

`func (o *UserItem) SetLoggedInAt(v time.Time)`

SetLoggedInAt sets LoggedInAt field to given value.

### HasLoggedInAt

`func (o *UserItem) HasLoggedInAt() bool`

HasLoggedInAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


