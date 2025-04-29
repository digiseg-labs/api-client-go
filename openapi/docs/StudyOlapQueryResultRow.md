# StudyOlapQueryResultRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurements** | [**[]Measurement**](Measurement.md) | Measurements related to this object | 
**Dimensions** | Pointer to [**[]StudyOlapFilterItem**](StudyOlapFilterItem.md) |  | [optional] 

## Methods

### NewStudyOlapQueryResultRow

`func NewStudyOlapQueryResultRow(measurements []Measurement, ) *StudyOlapQueryResultRow`

NewStudyOlapQueryResultRow instantiates a new StudyOlapQueryResultRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyOlapQueryResultRowWithDefaults

`func NewStudyOlapQueryResultRowWithDefaults() *StudyOlapQueryResultRow`

NewStudyOlapQueryResultRowWithDefaults instantiates a new StudyOlapQueryResultRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasurements

`func (o *StudyOlapQueryResultRow) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *StudyOlapQueryResultRow) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *StudyOlapQueryResultRow) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.


### GetDimensions

`func (o *StudyOlapQueryResultRow) GetDimensions() []StudyOlapFilterItem`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *StudyOlapQueryResultRow) GetDimensionsOk() (*[]StudyOlapFilterItem, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *StudyOlapQueryResultRow) SetDimensions(v []StudyOlapFilterItem)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *StudyOlapQueryResultRow) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


