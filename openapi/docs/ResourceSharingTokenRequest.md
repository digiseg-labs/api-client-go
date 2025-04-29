# ResourceSharingTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional ID to encode into the token to enable future identification of it | [optional] 
**ResourceUrls** | Pointer to **[]string** |  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account to create the resource sharing token for. The account ID will be encoded into the token and ensure that e.g. subscription limitations and account settings are carried over to the interactions provided with the token.  | [optional] 
**ExpiresIn** | Pointer to **int32** | The duration of time (in seconds) the resource sharing token is granted for | [optional] 

## Methods

### NewResourceSharingTokenRequest

`func NewResourceSharingTokenRequest() *ResourceSharingTokenRequest`

NewResourceSharingTokenRequest instantiates a new ResourceSharingTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSharingTokenRequestWithDefaults

`func NewResourceSharingTokenRequestWithDefaults() *ResourceSharingTokenRequest`

NewResourceSharingTokenRequestWithDefaults instantiates a new ResourceSharingTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceSharingTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSharingTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSharingTokenRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSharingTokenRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceUrls

`func (o *ResourceSharingTokenRequest) GetResourceUrls() []string`

GetResourceUrls returns the ResourceUrls field if non-nil, zero value otherwise.

### GetResourceUrlsOk

`func (o *ResourceSharingTokenRequest) GetResourceUrlsOk() (*[]string, bool)`

GetResourceUrlsOk returns a tuple with the ResourceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrls

`func (o *ResourceSharingTokenRequest) SetResourceUrls(v []string)`

SetResourceUrls sets ResourceUrls field to given value.

### HasResourceUrls

`func (o *ResourceSharingTokenRequest) HasResourceUrls() bool`

HasResourceUrls returns a boolean if a field has been set.

### GetAccountId

`func (o *ResourceSharingTokenRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceSharingTokenRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceSharingTokenRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResourceSharingTokenRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ResourceSharingTokenRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ResourceSharingTokenRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ResourceSharingTokenRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ResourceSharingTokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


