# StudyCountryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredominantCountry** | Pointer to **string** | The country code of the predominant country of the study | [optional] 
**Countries** | Pointer to [**[]CountryStats**](CountryStats.md) | A listing of each countries observed and the relevant measurements for each | [optional] 
**Resolved** | Pointer to [**MeasurementsContainer**](MeasurementsContainer.md) |  | [optional] 
**NotResolved** | Pointer to [**MeasurementsContainer**](MeasurementsContainer.md) |  | [optional] 

## Methods

### NewStudyCountryStats

`func NewStudyCountryStats() *StudyCountryStats`

NewStudyCountryStats instantiates a new StudyCountryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyCountryStatsWithDefaults

`func NewStudyCountryStatsWithDefaults() *StudyCountryStats`

NewStudyCountryStatsWithDefaults instantiates a new StudyCountryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredominantCountry

`func (o *StudyCountryStats) GetPredominantCountry() string`

GetPredominantCountry returns the PredominantCountry field if non-nil, zero value otherwise.

### GetPredominantCountryOk

`func (o *StudyCountryStats) GetPredominantCountryOk() (*string, bool)`

GetPredominantCountryOk returns a tuple with the PredominantCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredominantCountry

`func (o *StudyCountryStats) SetPredominantCountry(v string)`

SetPredominantCountry sets PredominantCountry field to given value.

### HasPredominantCountry

`func (o *StudyCountryStats) HasPredominantCountry() bool`

HasPredominantCountry returns a boolean if a field has been set.

### GetCountries

`func (o *StudyCountryStats) GetCountries() []CountryStats`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *StudyCountryStats) GetCountriesOk() (*[]CountryStats, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *StudyCountryStats) SetCountries(v []CountryStats)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *StudyCountryStats) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetResolved

`func (o *StudyCountryStats) GetResolved() MeasurementsContainer`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *StudyCountryStats) GetResolvedOk() (*MeasurementsContainer, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *StudyCountryStats) SetResolved(v MeasurementsContainer)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *StudyCountryStats) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetNotResolved

`func (o *StudyCountryStats) GetNotResolved() MeasurementsContainer`

GetNotResolved returns the NotResolved field if non-nil, zero value otherwise.

### GetNotResolvedOk

`func (o *StudyCountryStats) GetNotResolvedOk() (*MeasurementsContainer, bool)`

GetNotResolvedOk returns a tuple with the NotResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResolved

`func (o *StudyCountryStats) SetNotResolved(v MeasurementsContainer)`

SetNotResolved sets NotResolved field to given value.

### HasNotResolved

`func (o *StudyCountryStats) HasNotResolved() bool`

HasNotResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


