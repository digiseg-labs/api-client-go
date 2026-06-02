# OAuthRegistrationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationToken** | Pointer to **string** | Registration token to be used in the registration request with OAuth | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Expiration time of the registration token | [optional] 

## Methods

### NewOAuthRegistrationToken

`func NewOAuthRegistrationToken() *OAuthRegistrationToken`

NewOAuthRegistrationToken instantiates a new OAuthRegistrationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRegistrationTokenWithDefaults

`func NewOAuthRegistrationTokenWithDefaults() *OAuthRegistrationToken`

NewOAuthRegistrationTokenWithDefaults instantiates a new OAuthRegistrationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationToken

`func (o *OAuthRegistrationToken) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *OAuthRegistrationToken) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *OAuthRegistrationToken) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.

### HasRegistrationToken

`func (o *OAuthRegistrationToken) HasRegistrationToken() bool`

HasRegistrationToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OAuthRegistrationToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OAuthRegistrationToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OAuthRegistrationToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OAuthRegistrationToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


