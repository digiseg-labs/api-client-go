# MeasurementEventLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The base URI of the link to an event | [optional] 
**Parameters** | Pointer to [**map[string]MeasurementEventLinkParameterInfo**](MeasurementEventLinkParameterInfo.md) | Describes any parameters that can be added to the event link | [optional] 

## Methods

### NewMeasurementEventLink

`func NewMeasurementEventLink() *MeasurementEventLink`

NewMeasurementEventLink instantiates a new MeasurementEventLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementEventLinkWithDefaults

`func NewMeasurementEventLinkWithDefaults() *MeasurementEventLink`

NewMeasurementEventLinkWithDefaults instantiates a new MeasurementEventLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *MeasurementEventLink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MeasurementEventLink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MeasurementEventLink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *MeasurementEventLink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetParameters

`func (o *MeasurementEventLink) GetParameters() map[string]MeasurementEventLinkParameterInfo`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MeasurementEventLink) GetParametersOk() (*map[string]MeasurementEventLinkParameterInfo, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MeasurementEventLink) SetParameters(v map[string]MeasurementEventLinkParameterInfo)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MeasurementEventLink) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


