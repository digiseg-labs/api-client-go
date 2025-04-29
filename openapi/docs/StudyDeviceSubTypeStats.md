# StudyDeviceSubTypeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Count** | **int32** | The total count of events recorded for the device. | 
**Impressions** | Pointer to **int32** | The total count of impressions recorded for the device. | [optional] 
**Clicks** | Pointer to **int32** | The total count of clicks recorded for the device. | [optional] 

## Methods

### NewStudyDeviceSubTypeStats

`func NewStudyDeviceSubTypeStats(name string, count int32, ) *StudyDeviceSubTypeStats`

NewStudyDeviceSubTypeStats instantiates a new StudyDeviceSubTypeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyDeviceSubTypeStatsWithDefaults

`func NewStudyDeviceSubTypeStatsWithDefaults() *StudyDeviceSubTypeStats`

NewStudyDeviceSubTypeStatsWithDefaults instantiates a new StudyDeviceSubTypeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StudyDeviceSubTypeStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudyDeviceSubTypeStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudyDeviceSubTypeStats) SetName(v string)`

SetName sets Name field to given value.


### GetCount

`func (o *StudyDeviceSubTypeStats) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StudyDeviceSubTypeStats) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StudyDeviceSubTypeStats) SetCount(v int32)`

SetCount sets Count field to given value.


### GetImpressions

`func (o *StudyDeviceSubTypeStats) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *StudyDeviceSubTypeStats) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *StudyDeviceSubTypeStats) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *StudyDeviceSubTypeStats) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetClicks

`func (o *StudyDeviceSubTypeStats) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *StudyDeviceSubTypeStats) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *StudyDeviceSubTypeStats) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *StudyDeviceSubTypeStats) HasClicks() bool`

HasClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


