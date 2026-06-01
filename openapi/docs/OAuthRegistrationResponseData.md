# OAuthRegistrationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**ProviderUserId** | **string** |  | 
**AccessToken** | [**AccessTokenData**](AccessTokenData.md) |  | 

## Methods

### NewOAuthRegistrationResponseData

`func NewOAuthRegistrationResponseData(userId string, providerUserId string, accessToken AccessTokenData, ) *OAuthRegistrationResponseData`

NewOAuthRegistrationResponseData instantiates a new OAuthRegistrationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRegistrationResponseDataWithDefaults

`func NewOAuthRegistrationResponseDataWithDefaults() *OAuthRegistrationResponseData`

NewOAuthRegistrationResponseDataWithDefaults instantiates a new OAuthRegistrationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *OAuthRegistrationResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OAuthRegistrationResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OAuthRegistrationResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetProviderUserId

`func (o *OAuthRegistrationResponseData) GetProviderUserId() string`

GetProviderUserId returns the ProviderUserId field if non-nil, zero value otherwise.

### GetProviderUserIdOk

`func (o *OAuthRegistrationResponseData) GetProviderUserIdOk() (*string, bool)`

GetProviderUserIdOk returns a tuple with the ProviderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserId

`func (o *OAuthRegistrationResponseData) SetProviderUserId(v string)`

SetProviderUserId sets ProviderUserId field to given value.


### GetAccessToken

`func (o *OAuthRegistrationResponseData) GetAccessToken() AccessTokenData`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthRegistrationResponseData) GetAccessTokenOk() (*AccessTokenData, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthRegistrationResponseData) SetAccessToken(v AccessTokenData)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


