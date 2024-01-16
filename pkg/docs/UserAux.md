# UserAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | ID of the account that this user pertains to | [optional] 
**IsSuperAdmin** | Pointer to **bool** | Determines if the user is a super admin of Digiseg API services | [optional] 

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

### GetAccountId

`func (o *UserAux) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserAux) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserAux) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserAux) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


