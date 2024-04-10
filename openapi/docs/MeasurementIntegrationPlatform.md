# MeasurementIntegrationPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An ID for the integration platform, if the integration platform is a known platform. Note that integration platform ID uniqueness is a responsibility of the client since this is simply an optional reference point to keep. Can be null/omitted if the platform name is a one-off value with just a string.  | [optional] 
**Name** | Pointer to **string** | The name of the integration platform. | [optional] 

## Methods

### NewMeasurementIntegrationPlatform

`func NewMeasurementIntegrationPlatform() *MeasurementIntegrationPlatform`

NewMeasurementIntegrationPlatform instantiates a new MeasurementIntegrationPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementIntegrationPlatformWithDefaults

`func NewMeasurementIntegrationPlatformWithDefaults() *MeasurementIntegrationPlatform`

NewMeasurementIntegrationPlatformWithDefaults instantiates a new MeasurementIntegrationPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeasurementIntegrationPlatform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeasurementIntegrationPlatform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeasurementIntegrationPlatform) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeasurementIntegrationPlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MeasurementIntegrationPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeasurementIntegrationPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeasurementIntegrationPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeasurementIntegrationPlatform) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


