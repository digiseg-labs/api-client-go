# ApiKeyCreation

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
**Scopes** | Pointer to [**PermissionScopes**](PermissionScopes.md) |  | [optional] 

## Methods

### NewApiKeyCreation

`func NewApiKeyCreation() *ApiKeyCreation`

NewApiKeyCreation instantiates a new ApiKeyCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreationWithDefaults

`func NewApiKeyCreationWithDefaults() *ApiKeyCreation`

NewApiKeyCreationWithDefaults instantiates a new ApiKeyCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyCreation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyCreation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiKeyCreation) GetStatus() ApiKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiKeyCreation) GetStatusOk() (*ApiKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiKeyCreation) SetStatus(v ApiKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiKeyCreation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyCreation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyCreation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyCreation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyCreation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyCreation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyCreation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyCreation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyCreation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *ApiKeyCreation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiKeyCreation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiKeyCreation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiKeyCreation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ApiKeyCreation) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyCreation) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyCreation) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyCreation) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKeyCreation) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKeyCreation) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKeyCreation) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKeyCreation) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.

### GetScopes

`func (o *ApiKeyCreation) GetScopes() PermissionScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiKeyCreation) GetScopesOk() (*PermissionScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiKeyCreation) SetScopes(v PermissionScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiKeyCreation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


