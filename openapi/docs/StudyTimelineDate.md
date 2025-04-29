# StudyTimelineDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Hours** | [**[]StudyTimelineHour**](StudyTimelineHour.md) |  | 

## Methods

### NewStudyTimelineDate

`func NewStudyTimelineDate(date string, hours []StudyTimelineHour, ) *StudyTimelineDate`

NewStudyTimelineDate instantiates a new StudyTimelineDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyTimelineDateWithDefaults

`func NewStudyTimelineDateWithDefaults() *StudyTimelineDate`

NewStudyTimelineDateWithDefaults instantiates a new StudyTimelineDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *StudyTimelineDate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StudyTimelineDate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StudyTimelineDate) SetDate(v string)`

SetDate sets Date field to given value.


### GetHours

`func (o *StudyTimelineDate) GetHours() []StudyTimelineHour`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *StudyTimelineDate) GetHoursOk() (*[]StudyTimelineHour, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *StudyTimelineDate) SetHours(v []StudyTimelineHour)`

SetHours sets Hours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


