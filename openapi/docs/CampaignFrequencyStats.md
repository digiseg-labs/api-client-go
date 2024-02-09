# CampaignFrequencyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageFrequency** | Pointer to **float32** | The average frequency of impressions per user.  | [optional] 
**Frequencies** | Pointer to [**[]FrequencyStats**](FrequencyStats.md) | A listing of frequencies observed and the relevant measurements for each. The returned list may be truncated to cut off the \&quot;long tail\&quot; of frequency values.  | [optional] 
**CountAboveCap** | Pointer to **int32** | The number of users that have generated impressions at a frequency value greater than those represented in &#x60;frequencies&#x60;.  | [optional] 

## Methods

### NewCampaignFrequencyStats

`func NewCampaignFrequencyStats() *CampaignFrequencyStats`

NewCampaignFrequencyStats instantiates a new CampaignFrequencyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignFrequencyStatsWithDefaults

`func NewCampaignFrequencyStatsWithDefaults() *CampaignFrequencyStats`

NewCampaignFrequencyStatsWithDefaults instantiates a new CampaignFrequencyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageFrequency

`func (o *CampaignFrequencyStats) GetAverageFrequency() float32`

GetAverageFrequency returns the AverageFrequency field if non-nil, zero value otherwise.

### GetAverageFrequencyOk

`func (o *CampaignFrequencyStats) GetAverageFrequencyOk() (*float32, bool)`

GetAverageFrequencyOk returns a tuple with the AverageFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFrequency

`func (o *CampaignFrequencyStats) SetAverageFrequency(v float32)`

SetAverageFrequency sets AverageFrequency field to given value.

### HasAverageFrequency

`func (o *CampaignFrequencyStats) HasAverageFrequency() bool`

HasAverageFrequency returns a boolean if a field has been set.

### GetFrequencies

`func (o *CampaignFrequencyStats) GetFrequencies() []FrequencyStats`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *CampaignFrequencyStats) GetFrequenciesOk() (*[]FrequencyStats, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *CampaignFrequencyStats) SetFrequencies(v []FrequencyStats)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *CampaignFrequencyStats) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.

### GetCountAboveCap

`func (o *CampaignFrequencyStats) GetCountAboveCap() int32`

GetCountAboveCap returns the CountAboveCap field if non-nil, zero value otherwise.

### GetCountAboveCapOk

`func (o *CampaignFrequencyStats) GetCountAboveCapOk() (*int32, bool)`

GetCountAboveCapOk returns a tuple with the CountAboveCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAboveCap

`func (o *CampaignFrequencyStats) SetCountAboveCap(v int32)`

SetCountAboveCap sets CountAboveCap field to given value.

### HasCountAboveCap

`func (o *CampaignFrequencyStats) HasCountAboveCap() bool`

HasCountAboveCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


