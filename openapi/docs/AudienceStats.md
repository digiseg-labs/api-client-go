# AudienceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurements** | Pointer to [**[]Measurement**](Measurement.md) | Measurements related to this object | [optional] 
**Comparisons** | Pointer to [**[]Comparison**](Comparison.md) |  | [optional] 
**Code** | Pointer to **string** | The code of the audience | [optional] 

## Methods

### NewAudienceStats

`func NewAudienceStats() *AudienceStats`

NewAudienceStats instantiates a new AudienceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceStatsWithDefaults

`func NewAudienceStatsWithDefaults() *AudienceStats`

NewAudienceStatsWithDefaults instantiates a new AudienceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasurements

`func (o *AudienceStats) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *AudienceStats) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *AudienceStats) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *AudienceStats) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetComparisons

`func (o *AudienceStats) GetComparisons() []Comparison`

GetComparisons returns the Comparisons field if non-nil, zero value otherwise.

### GetComparisonsOk

`func (o *AudienceStats) GetComparisonsOk() (*[]Comparison, bool)`

GetComparisonsOk returns a tuple with the Comparisons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisons

`func (o *AudienceStats) SetComparisons(v []Comparison)`

SetComparisons sets Comparisons field to given value.

### HasComparisons

`func (o *AudienceStats) HasComparisons() bool`

HasComparisons returns a boolean if a field has been set.

### GetCode

`func (o *AudienceStats) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudienceStats) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudienceStats) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AudienceStats) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


