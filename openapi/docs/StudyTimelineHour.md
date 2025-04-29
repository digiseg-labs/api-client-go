# StudyTimelineHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** | The hour of the day (from 0-23) in UTC time | [optional] 
**Count** | Pointer to **int32** | The total count of events recorded in the hour. | [optional] 
**Impressions** | Pointer to **int32** | The total count of impressions recorded in the hour. | [optional] 
**Clicks** | Pointer to **int32** | The total count of clicks recorded in the hour. | [optional] 

## Methods

### NewStudyTimelineHour

`func NewStudyTimelineHour() *StudyTimelineHour`

NewStudyTimelineHour instantiates a new StudyTimelineHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyTimelineHourWithDefaults

`func NewStudyTimelineHourWithDefaults() *StudyTimelineHour`

NewStudyTimelineHourWithDefaults instantiates a new StudyTimelineHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *StudyTimelineHour) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *StudyTimelineHour) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *StudyTimelineHour) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *StudyTimelineHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetCount

`func (o *StudyTimelineHour) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StudyTimelineHour) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StudyTimelineHour) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *StudyTimelineHour) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetImpressions

`func (o *StudyTimelineHour) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *StudyTimelineHour) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *StudyTimelineHour) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *StudyTimelineHour) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetClicks

`func (o *StudyTimelineHour) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *StudyTimelineHour) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *StudyTimelineHour) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *StudyTimelineHour) HasClicks() bool`

HasClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


