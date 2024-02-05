# RegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | The name of the account to create | 
**OwnerEmail** | **string** | The email of the account owner | 
**OwnerName** | **string** | The name of the account owner | 

## Methods

### NewRegistrationRequest

`func NewRegistrationRequest(accountName string, ownerEmail string, ownerName string, ) *RegistrationRequest`

NewRegistrationRequest instantiates a new RegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationRequestWithDefaults

`func NewRegistrationRequestWithDefaults() *RegistrationRequest`

NewRegistrationRequestWithDefaults instantiates a new RegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *RegistrationRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *RegistrationRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *RegistrationRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetOwnerEmail

`func (o *RegistrationRequest) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *RegistrationRequest) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *RegistrationRequest) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.


### GetOwnerName

`func (o *RegistrationRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *RegistrationRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *RegistrationRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


