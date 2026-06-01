# OAuthCallbackResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserExists** | **bool** | Indicates whether the user already exists (true) or is a new user (false) | 
**RegistrationToken** | Pointer to [**OAuthRegistrationToken**](OAuthRegistrationToken.md) |  | [optional] 
**AccessToken** | Pointer to [**AccessTokenData**](AccessTokenData.md) |  | [optional] 
**ProviderUserId** | **string** | The ID of the user from the OAuth provider | 
**Username** | **string** | Username from the OAuth provider | 
**Email** | **string** | Email from the OAuth provider | 
**UserId** | Pointer to **string** | User ID (present when user already exists) | [optional] 

## Methods

### NewOAuthCallbackResponseData

`func NewOAuthCallbackResponseData(userExists bool, providerUserId string, username string, email string, ) *OAuthCallbackResponseData`

NewOAuthCallbackResponseData instantiates a new OAuthCallbackResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCallbackResponseDataWithDefaults

`func NewOAuthCallbackResponseDataWithDefaults() *OAuthCallbackResponseData`

NewOAuthCallbackResponseDataWithDefaults instantiates a new OAuthCallbackResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserExists

`func (o *OAuthCallbackResponseData) GetUserExists() bool`

GetUserExists returns the UserExists field if non-nil, zero value otherwise.

### GetUserExistsOk

`func (o *OAuthCallbackResponseData) GetUserExistsOk() (*bool, bool)`

GetUserExistsOk returns a tuple with the UserExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExists

`func (o *OAuthCallbackResponseData) SetUserExists(v bool)`

SetUserExists sets UserExists field to given value.


### GetRegistrationToken

`func (o *OAuthCallbackResponseData) GetRegistrationToken() OAuthRegistrationToken`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *OAuthCallbackResponseData) GetRegistrationTokenOk() (*OAuthRegistrationToken, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *OAuthCallbackResponseData) SetRegistrationToken(v OAuthRegistrationToken)`

SetRegistrationToken sets RegistrationToken field to given value.

### HasRegistrationToken

`func (o *OAuthCallbackResponseData) HasRegistrationToken() bool`

HasRegistrationToken returns a boolean if a field has been set.

### GetAccessToken

`func (o *OAuthCallbackResponseData) GetAccessToken() AccessTokenData`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthCallbackResponseData) GetAccessTokenOk() (*AccessTokenData, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthCallbackResponseData) SetAccessToken(v AccessTokenData)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OAuthCallbackResponseData) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetProviderUserId

`func (o *OAuthCallbackResponseData) GetProviderUserId() string`

GetProviderUserId returns the ProviderUserId field if non-nil, zero value otherwise.

### GetProviderUserIdOk

`func (o *OAuthCallbackResponseData) GetProviderUserIdOk() (*string, bool)`

GetProviderUserIdOk returns a tuple with the ProviderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserId

`func (o *OAuthCallbackResponseData) SetProviderUserId(v string)`

SetProviderUserId sets ProviderUserId field to given value.


### GetUsername

`func (o *OAuthCallbackResponseData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OAuthCallbackResponseData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OAuthCallbackResponseData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *OAuthCallbackResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OAuthCallbackResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OAuthCallbackResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUserId

`func (o *OAuthCallbackResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OAuthCallbackResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OAuthCallbackResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OAuthCallbackResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


