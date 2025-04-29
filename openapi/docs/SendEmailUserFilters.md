# SendEmailUserFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**FilterAccountId** | Pointer to **string** |  | [optional] 
**FilterPlatformRole** | Pointer to [**UserPlatformRole**](UserPlatformRole.md) |  | [optional] 
**FilterAccountRole** | Pointer to [**UserAccountRole**](UserAccountRole.md) |  | [optional] 

## Methods

### NewSendEmailUserFilters

`func NewSendEmailUserFilters(queryType string, ) *SendEmailUserFilters`

NewSendEmailUserFilters instantiates a new SendEmailUserFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailUserFiltersWithDefaults

`func NewSendEmailUserFiltersWithDefaults() *SendEmailUserFilters`

NewSendEmailUserFiltersWithDefaults instantiates a new SendEmailUserFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *SendEmailUserFilters) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *SendEmailUserFilters) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *SendEmailUserFilters) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetUserId

`func (o *SendEmailUserFilters) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SendEmailUserFilters) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SendEmailUserFilters) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SendEmailUserFilters) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFilterAccountId

`func (o *SendEmailUserFilters) GetFilterAccountId() string`

GetFilterAccountId returns the FilterAccountId field if non-nil, zero value otherwise.

### GetFilterAccountIdOk

`func (o *SendEmailUserFilters) GetFilterAccountIdOk() (*string, bool)`

GetFilterAccountIdOk returns a tuple with the FilterAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAccountId

`func (o *SendEmailUserFilters) SetFilterAccountId(v string)`

SetFilterAccountId sets FilterAccountId field to given value.

### HasFilterAccountId

`func (o *SendEmailUserFilters) HasFilterAccountId() bool`

HasFilterAccountId returns a boolean if a field has been set.

### GetFilterPlatformRole

`func (o *SendEmailUserFilters) GetFilterPlatformRole() UserPlatformRole`

GetFilterPlatformRole returns the FilterPlatformRole field if non-nil, zero value otherwise.

### GetFilterPlatformRoleOk

`func (o *SendEmailUserFilters) GetFilterPlatformRoleOk() (*UserPlatformRole, bool)`

GetFilterPlatformRoleOk returns a tuple with the FilterPlatformRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPlatformRole

`func (o *SendEmailUserFilters) SetFilterPlatformRole(v UserPlatformRole)`

SetFilterPlatformRole sets FilterPlatformRole field to given value.

### HasFilterPlatformRole

`func (o *SendEmailUserFilters) HasFilterPlatformRole() bool`

HasFilterPlatformRole returns a boolean if a field has been set.

### GetFilterAccountRole

`func (o *SendEmailUserFilters) GetFilterAccountRole() UserAccountRole`

GetFilterAccountRole returns the FilterAccountRole field if non-nil, zero value otherwise.

### GetFilterAccountRoleOk

`func (o *SendEmailUserFilters) GetFilterAccountRoleOk() (*UserAccountRole, bool)`

GetFilterAccountRoleOk returns a tuple with the FilterAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAccountRole

`func (o *SendEmailUserFilters) SetFilterAccountRole(v UserAccountRole)`

SetFilterAccountRole sets FilterAccountRole field to given value.

### HasFilterAccountRole

`func (o *SendEmailUserFilters) HasFilterAccountRole() bool`

HasFilterAccountRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


