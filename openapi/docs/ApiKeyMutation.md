# ApiKeyMutation

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

### NewApiKeyMutation

`func NewApiKeyMutation() *ApiKeyMutation`

NewApiKeyMutation instantiates a new ApiKeyMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyMutationWithDefaults

`func NewApiKeyMutationWithDefaults() *ApiKeyMutation`

NewApiKeyMutationWithDefaults instantiates a new ApiKeyMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiKeyMutation) GetStatus() ApiKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiKeyMutation) GetStatusOk() (*ApiKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiKeyMutation) SetStatus(v ApiKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiKeyMutation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyMutation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyMutation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyMutation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyMutation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyMutation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyMutation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyMutation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyMutation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *ApiKeyMutation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiKeyMutation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiKeyMutation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiKeyMutation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ApiKeyMutation) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyMutation) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyMutation) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyMutation) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKeyMutation) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKeyMutation) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKeyMutation) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKeyMutation) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.

### GetScopes

`func (o *ApiKeyMutation) GetScopes() PermissionScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiKeyMutation) GetScopesOk() (*PermissionScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiKeyMutation) SetScopes(v PermissionScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiKeyMutation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


