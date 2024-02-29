# ApiKeyBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the API key | [optional] 
**Status** | Pointer to [**ApiKeyStatus**](ApiKeyStatus.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the key will expire | [optional] 
**UserId** | Pointer to **string** | The ID of the API key&#39;s user.  | [optional] [readonly] 
**AccountId** | Pointer to **string** | The ID of account that the API key is associated with.  | [optional] 
**LastUsedAt** | Pointer to **time.Time** | The approximate last time that the API key was used to authenticate API requests | [optional] [readonly] 
**TokenPrefix** | Pointer to **string** | A prefix of the API key | [optional] [readonly] 

## Methods

### NewApiKeyBase

`func NewApiKeyBase() *ApiKeyBase`

NewApiKeyBase instantiates a new ApiKeyBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyBaseWithDefaults

`func NewApiKeyBaseWithDefaults() *ApiKeyBase`

NewApiKeyBaseWithDefaults instantiates a new ApiKeyBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiKeyBase) GetStatus() ApiKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiKeyBase) GetStatusOk() (*ApiKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiKeyBase) SetStatus(v ApiKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiKeyBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyBase) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyBase) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyBase) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyBase) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyBase) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyBase) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyBase) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyBase) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *ApiKeyBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiKeyBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiKeyBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiKeyBase) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ApiKeyBase) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyBase) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyBase) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyBase) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKeyBase) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKeyBase) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKeyBase) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKeyBase) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


