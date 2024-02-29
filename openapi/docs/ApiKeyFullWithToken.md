# ApiKeyFullWithToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The actual API key token to use with the &#x60;X-API-KEY&#x60; header to authenticate with the key | [optional] 
**Id** | Pointer to **string** | Unique ID for the object | [optional] 
**Name** | Pointer to **string** | Human readable name of the API key | [optional] 
**Status** | Pointer to [**ApiKeyStatus**](ApiKeyStatus.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the key will expire | [optional] 
**UserId** | Pointer to **string** | The ID of the API key&#39;s user.  | [optional] [readonly] 
**AccountId** | Pointer to **string** | The ID of account that the API key is associated with.  | [optional] 
**LastUsedAt** | Pointer to **time.Time** | The approximate last time that the API key was used to authenticate API requests | [optional] [readonly] 
**TokenPrefix** | Pointer to **string** | A prefix of the API key | [optional] [readonly] 
**Scopes** | Pointer to [**PermissionScopes**](PermissionScopes.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] 

## Methods

### NewApiKeyFullWithToken

`func NewApiKeyFullWithToken() *ApiKeyFullWithToken`

NewApiKeyFullWithToken instantiates a new ApiKeyFullWithToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyFullWithTokenWithDefaults

`func NewApiKeyFullWithTokenWithDefaults() *ApiKeyFullWithToken`

NewApiKeyFullWithTokenWithDefaults instantiates a new ApiKeyFullWithToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ApiKeyFullWithToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiKeyFullWithToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiKeyFullWithToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiKeyFullWithToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *ApiKeyFullWithToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyFullWithToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyFullWithToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyFullWithToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiKeyFullWithToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyFullWithToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyFullWithToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyFullWithToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiKeyFullWithToken) GetStatus() ApiKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiKeyFullWithToken) GetStatusOk() (*ApiKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiKeyFullWithToken) SetStatus(v ApiKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiKeyFullWithToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyFullWithToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyFullWithToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyFullWithToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyFullWithToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyFullWithToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyFullWithToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyFullWithToken) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyFullWithToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *ApiKeyFullWithToken) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiKeyFullWithToken) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiKeyFullWithToken) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiKeyFullWithToken) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ApiKeyFullWithToken) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyFullWithToken) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyFullWithToken) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyFullWithToken) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKeyFullWithToken) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKeyFullWithToken) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKeyFullWithToken) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKeyFullWithToken) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.

### GetScopes

`func (o *ApiKeyFullWithToken) GetScopes() PermissionScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiKeyFullWithToken) GetScopesOk() (*PermissionScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiKeyFullWithToken) SetScopes(v PermissionScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiKeyFullWithToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKeyFullWithToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyFullWithToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyFullWithToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKeyFullWithToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApiKeyFullWithToken) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiKeyFullWithToken) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiKeyFullWithToken) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApiKeyFullWithToken) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApiKeyFullWithToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiKeyFullWithToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiKeyFullWithToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApiKeyFullWithToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ApiKeyFullWithToken) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ApiKeyFullWithToken) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ApiKeyFullWithToken) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ApiKeyFullWithToken) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


