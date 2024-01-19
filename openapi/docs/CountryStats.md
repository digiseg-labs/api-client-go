# CountryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurements** | Pointer to [**[]Measurement**](Measurement.md) | Measurements related to this object | [optional] 
**Code** | Pointer to **string** | Country code of the country | [optional] 

## Methods

### NewCountryStats

`func NewCountryStats() *CountryStats`

NewCountryStats instantiates a new CountryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryStatsWithDefaults

`func NewCountryStatsWithDefaults() *CountryStats`

NewCountryStatsWithDefaults instantiates a new CountryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasurements

`func (o *CountryStats) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *CountryStats) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *CountryStats) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *CountryStats) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetCode

`func (o *CountryStats) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryStats) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryStats) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CountryStats) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


