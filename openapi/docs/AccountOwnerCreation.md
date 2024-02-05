# AccountOwnerCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The account owner&#39;s name | 
**Email** | **string** | The account owner&#39;s email | 
**Password** | Pointer to **string** | The account owner&#39;s password | [optional] 

## Methods

### NewAccountOwnerCreation

`func NewAccountOwnerCreation(name string, email string, ) *AccountOwnerCreation`

NewAccountOwnerCreation instantiates a new AccountOwnerCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOwnerCreationWithDefaults

`func NewAccountOwnerCreationWithDefaults() *AccountOwnerCreation`

NewAccountOwnerCreationWithDefaults instantiates a new AccountOwnerCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountOwnerCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountOwnerCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountOwnerCreation) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *AccountOwnerCreation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountOwnerCreation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountOwnerCreation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AccountOwnerCreation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountOwnerCreation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountOwnerCreation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AccountOwnerCreation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


