# StudyFrequencyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The event that is represented in this frequency stats. Current values include &#x60;all&#x60; (all events), &#x60;impression&#x60; or &#x60;click&#x60; | [optional] 
**SubFrequencies** | Pointer to [**[]StudyFrequencyStats**](StudyFrequencyStats.md) | Optionally an array of additional frequency stats for more fine-grained event types | [optional] 
**AverageFrequency** | Pointer to **float32** | The average frequency of events per user.  | [optional] 
**Frequencies** | Pointer to [**[]FrequencyStats**](FrequencyStats.md) | A listing of frequencies observed and the relevant measurements for each. The returned list may be truncated to cut off the \&quot;long tail\&quot; of frequency values.  | [optional] 
**CountAboveCap** | Pointer to **int32** | The number of users that have generated events at a frequency value greater than those represented in &#x60;frequencies&#x60;.  | [optional] 

## Methods

### NewStudyFrequencyStats

`func NewStudyFrequencyStats() *StudyFrequencyStats`

NewStudyFrequencyStats instantiates a new StudyFrequencyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyFrequencyStatsWithDefaults

`func NewStudyFrequencyStatsWithDefaults() *StudyFrequencyStats`

NewStudyFrequencyStatsWithDefaults instantiates a new StudyFrequencyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *StudyFrequencyStats) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StudyFrequencyStats) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StudyFrequencyStats) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StudyFrequencyStats) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSubFrequencies

`func (o *StudyFrequencyStats) GetSubFrequencies() []StudyFrequencyStats`

GetSubFrequencies returns the SubFrequencies field if non-nil, zero value otherwise.

### GetSubFrequenciesOk

`func (o *StudyFrequencyStats) GetSubFrequenciesOk() (*[]StudyFrequencyStats, bool)`

GetSubFrequenciesOk returns a tuple with the SubFrequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFrequencies

`func (o *StudyFrequencyStats) SetSubFrequencies(v []StudyFrequencyStats)`

SetSubFrequencies sets SubFrequencies field to given value.

### HasSubFrequencies

`func (o *StudyFrequencyStats) HasSubFrequencies() bool`

HasSubFrequencies returns a boolean if a field has been set.

### GetAverageFrequency

`func (o *StudyFrequencyStats) GetAverageFrequency() float32`

GetAverageFrequency returns the AverageFrequency field if non-nil, zero value otherwise.

### GetAverageFrequencyOk

`func (o *StudyFrequencyStats) GetAverageFrequencyOk() (*float32, bool)`

GetAverageFrequencyOk returns a tuple with the AverageFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFrequency

`func (o *StudyFrequencyStats) SetAverageFrequency(v float32)`

SetAverageFrequency sets AverageFrequency field to given value.

### HasAverageFrequency

`func (o *StudyFrequencyStats) HasAverageFrequency() bool`

HasAverageFrequency returns a boolean if a field has been set.

### GetFrequencies

`func (o *StudyFrequencyStats) GetFrequencies() []FrequencyStats`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *StudyFrequencyStats) GetFrequenciesOk() (*[]FrequencyStats, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *StudyFrequencyStats) SetFrequencies(v []FrequencyStats)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *StudyFrequencyStats) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.

### GetCountAboveCap

`func (o *StudyFrequencyStats) GetCountAboveCap() int32`

GetCountAboveCap returns the CountAboveCap field if non-nil, zero value otherwise.

### GetCountAboveCapOk

`func (o *StudyFrequencyStats) GetCountAboveCapOk() (*int32, bool)`

GetCountAboveCapOk returns a tuple with the CountAboveCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAboveCap

`func (o *StudyFrequencyStats) SetCountAboveCap(v int32)`

SetCountAboveCap sets CountAboveCap field to given value.

### HasCountAboveCap

`func (o *StudyFrequencyStats) HasCountAboveCap() bool`

HasCountAboveCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


