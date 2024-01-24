# BusinessAudienceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurements** | Pointer to [**[]Measurement**](Measurement.md) | Measurements related to this object | [optional] 
**AudienceCategories** | Pointer to [**BusinessAudienceStatsAudienceCategories**](BusinessAudienceStatsAudienceCategories.md) |  | [optional] 

## Methods

### NewBusinessAudienceStats

`func NewBusinessAudienceStats() *BusinessAudienceStats`

NewBusinessAudienceStats instantiates a new BusinessAudienceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessAudienceStatsWithDefaults

`func NewBusinessAudienceStatsWithDefaults() *BusinessAudienceStats`

NewBusinessAudienceStatsWithDefaults instantiates a new BusinessAudienceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasurements

`func (o *BusinessAudienceStats) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *BusinessAudienceStats) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *BusinessAudienceStats) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *BusinessAudienceStats) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetAudienceCategories

`func (o *BusinessAudienceStats) GetAudienceCategories() BusinessAudienceStatsAudienceCategories`

GetAudienceCategories returns the AudienceCategories field if non-nil, zero value otherwise.

### GetAudienceCategoriesOk

`func (o *BusinessAudienceStats) GetAudienceCategoriesOk() (*BusinessAudienceStatsAudienceCategories, bool)`

GetAudienceCategoriesOk returns a tuple with the AudienceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceCategories

`func (o *BusinessAudienceStats) SetAudienceCategories(v BusinessAudienceStatsAudienceCategories)`

SetAudienceCategories sets AudienceCategories field to given value.

### HasAudienceCategories

`func (o *BusinessAudienceStats) HasAudienceCategories() bool`

HasAudienceCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


