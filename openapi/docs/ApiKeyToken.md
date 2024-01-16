# ApiKeyToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The actual API key token to use with the &#x60;X-API-KEY&#x60; header to authenticate with the key | [optional] 

## Methods

### NewApiKeyToken

`func NewApiKeyToken() *ApiKeyToken`

NewApiKeyToken instantiates a new ApiKeyToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyTokenWithDefaults

`func NewApiKeyTokenWithDefaults() *ApiKeyToken`

NewApiKeyTokenWithDefaults instantiates a new ApiKeyToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ApiKeyToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiKeyToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiKeyToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiKeyToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


