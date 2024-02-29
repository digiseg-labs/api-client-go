# MeasurementClientFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
**Name** | Pointer to **string** | The name of the measurement client | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this client | [optional] [readonly] 
**LogoUrl** | Pointer to **string** | The URL of an image representing the measurement client | [optional] [readonly] 
**CreatedAt** | **time.Time** | Date and time of the object creation | 
**CreatedBy** | **string** | ID of the user who created the object | 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] 

## Methods

### NewMeasurementClientFull

`func NewMeasurementClientFull(id string, createdAt time.Time, createdBy string, ) *MeasurementClientFull`

NewMeasurementClientFull instantiates a new MeasurementClientFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementClientFullWithDefaults

`func NewMeasurementClientFullWithDefaults() *MeasurementClientFull`

NewMeasurementClientFullWithDefaults instantiates a new MeasurementClientFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeasurementClientFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeasurementClientFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeasurementClientFull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MeasurementClientFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeasurementClientFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeasurementClientFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeasurementClientFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *MeasurementClientFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MeasurementClientFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MeasurementClientFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MeasurementClientFull) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLogoUrl

`func (o *MeasurementClientFull) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *MeasurementClientFull) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *MeasurementClientFull) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *MeasurementClientFull) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MeasurementClientFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeasurementClientFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeasurementClientFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *MeasurementClientFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MeasurementClientFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MeasurementClientFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *MeasurementClientFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MeasurementClientFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MeasurementClientFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MeasurementClientFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *MeasurementClientFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *MeasurementClientFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *MeasurementClientFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *MeasurementClientFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


