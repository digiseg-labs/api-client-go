# RegistrationByIdResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the registration | [optional] 
**Request** | Pointer to [**RegistrationRequest**](RegistrationRequest.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Date and time of the registration expiry | [optional] 
**VerifiedAt** | Pointer to **time.Time** | Date and time of the verification, if verified | [optional] 
**VerificationCode** | Pointer to **string** | The code needed to verify this registration | [optional] 

## Methods

### NewRegistrationByIdResponseData

`func NewRegistrationByIdResponseData() *RegistrationByIdResponseData`

NewRegistrationByIdResponseData instantiates a new RegistrationByIdResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationByIdResponseDataWithDefaults

`func NewRegistrationByIdResponseDataWithDefaults() *RegistrationByIdResponseData`

NewRegistrationByIdResponseDataWithDefaults instantiates a new RegistrationByIdResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistrationByIdResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrationByIdResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrationByIdResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrationByIdResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequest

`func (o *RegistrationByIdResponseData) GetRequest() RegistrationRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *RegistrationByIdResponseData) GetRequestOk() (*RegistrationRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *RegistrationByIdResponseData) SetRequest(v RegistrationRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *RegistrationByIdResponseData) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetExpiresAt

`func (o *RegistrationByIdResponseData) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RegistrationByIdResponseData) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RegistrationByIdResponseData) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *RegistrationByIdResponseData) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *RegistrationByIdResponseData) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *RegistrationByIdResponseData) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *RegistrationByIdResponseData) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *RegistrationByIdResponseData) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetVerificationCode

`func (o *RegistrationByIdResponseData) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *RegistrationByIdResponseData) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *RegistrationByIdResponseData) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *RegistrationByIdResponseData) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


