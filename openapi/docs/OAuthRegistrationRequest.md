# OAuthRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | The name of the account to create | 
**OwnerEmail** | **string** | The email of the account owner | 
**OwnerName** | **string** | The name of the account owner | 
**ProviderUserId** | **string** | The ID of the user from the OAuth provider | 
**RegistrationToken** | **string** | Registration token provided after successful OAuth authentication | 

## Methods

### NewOAuthRegistrationRequest

`func NewOAuthRegistrationRequest(accountName string, ownerEmail string, ownerName string, providerUserId string, registrationToken string, ) *OAuthRegistrationRequest`

NewOAuthRegistrationRequest instantiates a new OAuthRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRegistrationRequestWithDefaults

`func NewOAuthRegistrationRequestWithDefaults() *OAuthRegistrationRequest`

NewOAuthRegistrationRequestWithDefaults instantiates a new OAuthRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *OAuthRegistrationRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *OAuthRegistrationRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *OAuthRegistrationRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetOwnerEmail

`func (o *OAuthRegistrationRequest) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *OAuthRegistrationRequest) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *OAuthRegistrationRequest) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.


### GetOwnerName

`func (o *OAuthRegistrationRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *OAuthRegistrationRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *OAuthRegistrationRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetProviderUserId

`func (o *OAuthRegistrationRequest) GetProviderUserId() string`

GetProviderUserId returns the ProviderUserId field if non-nil, zero value otherwise.

### GetProviderUserIdOk

`func (o *OAuthRegistrationRequest) GetProviderUserIdOk() (*string, bool)`

GetProviderUserIdOk returns a tuple with the ProviderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserId

`func (o *OAuthRegistrationRequest) SetProviderUserId(v string)`

SetProviderUserId sets ProviderUserId field to given value.


### GetRegistrationToken

`func (o *OAuthRegistrationRequest) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *OAuthRegistrationRequest) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *OAuthRegistrationRequest) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


