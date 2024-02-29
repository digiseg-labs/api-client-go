# MeasurementClientMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the measurement client | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this client | [optional] [readonly] 
**LogoUrl** | Pointer to **string** | The URL of an image representing the measurement client | [optional] [readonly] 

## Methods

### NewMeasurementClientMutation

`func NewMeasurementClientMutation() *MeasurementClientMutation`

NewMeasurementClientMutation instantiates a new MeasurementClientMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementClientMutationWithDefaults

`func NewMeasurementClientMutationWithDefaults() *MeasurementClientMutation`

NewMeasurementClientMutationWithDefaults instantiates a new MeasurementClientMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MeasurementClientMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeasurementClientMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeasurementClientMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeasurementClientMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *MeasurementClientMutation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MeasurementClientMutation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MeasurementClientMutation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MeasurementClientMutation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLogoUrl

`func (o *MeasurementClientMutation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *MeasurementClientMutation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *MeasurementClientMutation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *MeasurementClientMutation) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


