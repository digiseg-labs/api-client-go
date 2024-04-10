# FrequencyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **int32** | A frequency of study impressions to users. * The value 1 means that the study has been exposed just once. * The value 2 means that the study has been exposed twice. * And so on...  | [optional] 
**Count** | Pointer to **int32** | The number of users that have generated impressions at the corresponding frequency | [optional] 

## Methods

### NewFrequencyStats

`func NewFrequencyStats() *FrequencyStats`

NewFrequencyStats instantiates a new FrequencyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrequencyStatsWithDefaults

`func NewFrequencyStatsWithDefaults() *FrequencyStats`

NewFrequencyStatsWithDefaults instantiates a new FrequencyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *FrequencyStats) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *FrequencyStats) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *FrequencyStats) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *FrequencyStats) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetCount

`func (o *FrequencyStats) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FrequencyStats) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FrequencyStats) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FrequencyStats) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


