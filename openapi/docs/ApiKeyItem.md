# ApiKeyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] 
**Name** | Pointer to **string** | Human readable name of the API key | [optional] 
**Status** | Pointer to [**ApiKeyStatus**](ApiKeyStatus.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the key will expire | [optional] 
**UserId** | Pointer to **string** | The ID of the API key&#39;s user.  | [optional] [readonly] 
**AccountId** | Pointer to **string** | The ID of account that the API key is associated with.  | [optional] 
**LastUsedAt** | Pointer to **time.Time** | The approximate last time that the API key was used to authenticate API requests | [optional] [readonly] 
**TokenPrefix** | Pointer to **string** | A prefix of the API key | [optional] [readonly] 

## Methods

### NewApiKeyItem

`func NewApiKeyItem() *ApiKeyItem`

NewApiKeyItem instantiates a new ApiKeyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyItemWithDefaults

`func NewApiKeyItemWithDefaults() *ApiKeyItem`

NewApiKeyItemWithDefaults instantiates a new ApiKeyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKeyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiKeyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiKeyItem) GetStatus() ApiKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiKeyItem) GetStatusOk() (*ApiKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiKeyItem) SetStatus(v ApiKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiKeyItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyItem) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *ApiKeyItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiKeyItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiKeyItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiKeyItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ApiKeyItem) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyItem) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyItem) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyItem) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKeyItem) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKeyItem) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKeyItem) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKeyItem) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


