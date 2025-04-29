# AuthTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username (typically an email address) of the user to authenticate | 
**Otp** | Pointer to **string** | A one-time password provided to perform passwordless auth | [optional] 
**Password** | Pointer to **string** | The password for the given username | [optional] 
**AccountId** | Pointer to **string** | Optional account_id to authenticate for, if the user has multiple account memberships | [optional] 
**RefreshToken** | Pointer to **string** | A previously issued refresh token for the given username | [optional] 
**Scopes** | Pointer to [**PermissionScopes**](PermissionScopes.md) |  | [optional] 

## Methods

### NewAuthTokenRequest

`func NewAuthTokenRequest(username string, ) *AuthTokenRequest`

NewAuthTokenRequest instantiates a new AuthTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenRequestWithDefaults

`func NewAuthTokenRequestWithDefaults() *AuthTokenRequest`

NewAuthTokenRequestWithDefaults instantiates a new AuthTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AuthTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetOtp

`func (o *AuthTokenRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *AuthTokenRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *AuthTokenRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *AuthTokenRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPassword

`func (o *AuthTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthTokenRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccountId

`func (o *AuthTokenRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuthTokenRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuthTokenRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AuthTokenRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AuthTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AuthTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScopes

`func (o *AuthTokenRequest) GetScopes() PermissionScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthTokenRequest) GetScopesOk() (*PermissionScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthTokenRequest) SetScopes(v PermissionScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthTokenRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


