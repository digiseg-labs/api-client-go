# CampaignCountryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredominantCountry** | Pointer to **string** | The country code of the predominant country of the campaign | [optional] 
**Countries** | Pointer to [**[]CountryStats**](CountryStats.md) | A listing of each countries observed and the relevant measurements for each | [optional] 
**Resolved** | Pointer to [**MeasurementsContainer**](MeasurementsContainer.md) |  | [optional] 
**NotResolved** | Pointer to [**MeasurementsContainer**](MeasurementsContainer.md) |  | [optional] 

## Methods

### NewCampaignCountryStats

`func NewCampaignCountryStats() *CampaignCountryStats`

NewCampaignCountryStats instantiates a new CampaignCountryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCountryStatsWithDefaults

`func NewCampaignCountryStatsWithDefaults() *CampaignCountryStats`

NewCampaignCountryStatsWithDefaults instantiates a new CampaignCountryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredominantCountry

`func (o *CampaignCountryStats) GetPredominantCountry() string`

GetPredominantCountry returns the PredominantCountry field if non-nil, zero value otherwise.

### GetPredominantCountryOk

`func (o *CampaignCountryStats) GetPredominantCountryOk() (*string, bool)`

GetPredominantCountryOk returns a tuple with the PredominantCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredominantCountry

`func (o *CampaignCountryStats) SetPredominantCountry(v string)`

SetPredominantCountry sets PredominantCountry field to given value.

### HasPredominantCountry

`func (o *CampaignCountryStats) HasPredominantCountry() bool`

HasPredominantCountry returns a boolean if a field has been set.

### GetCountries

`func (o *CampaignCountryStats) GetCountries() []CountryStats`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CampaignCountryStats) GetCountriesOk() (*[]CountryStats, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CampaignCountryStats) SetCountries(v []CountryStats)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CampaignCountryStats) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetResolved

`func (o *CampaignCountryStats) GetResolved() MeasurementsContainer`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *CampaignCountryStats) GetResolvedOk() (*MeasurementsContainer, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *CampaignCountryStats) SetResolved(v MeasurementsContainer)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *CampaignCountryStats) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetNotResolved

`func (o *CampaignCountryStats) GetNotResolved() MeasurementsContainer`

GetNotResolved returns the NotResolved field if non-nil, zero value otherwise.

### GetNotResolvedOk

`func (o *CampaignCountryStats) GetNotResolvedOk() (*MeasurementsContainer, bool)`

GetNotResolvedOk returns a tuple with the NotResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResolved

`func (o *CampaignCountryStats) SetNotResolved(v MeasurementsContainer)`

SetNotResolved sets NotResolved field to given value.

### HasNotResolved

`func (o *CampaignCountryStats) HasNotResolved() bool`

HasNotResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


