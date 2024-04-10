# StudyTimingStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HourOfDay** | Pointer to [**HourOfDayStats**](HourOfDayStats.md) |  | [optional] 
**DayOfWeek** | Pointer to [**DayOfWeekStats**](DayOfWeekStats.md) |  | [optional] 
**DayOfMonth** | Pointer to [**DayOfMonthStats**](DayOfMonthStats.md) |  | [optional] 

## Methods

### NewStudyTimingStats

`func NewStudyTimingStats() *StudyTimingStats`

NewStudyTimingStats instantiates a new StudyTimingStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyTimingStatsWithDefaults

`func NewStudyTimingStatsWithDefaults() *StudyTimingStats`

NewStudyTimingStatsWithDefaults instantiates a new StudyTimingStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourOfDay

`func (o *StudyTimingStats) GetHourOfDay() HourOfDayStats`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *StudyTimingStats) GetHourOfDayOk() (*HourOfDayStats, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *StudyTimingStats) SetHourOfDay(v HourOfDayStats)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *StudyTimingStats) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *StudyTimingStats) GetDayOfWeek() DayOfWeekStats`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *StudyTimingStats) GetDayOfWeekOk() (*DayOfWeekStats, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *StudyTimingStats) SetDayOfWeek(v DayOfWeekStats)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *StudyTimingStats) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *StudyTimingStats) GetDayOfMonth() DayOfMonthStats`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *StudyTimingStats) GetDayOfMonthOk() (*DayOfMonthStats, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *StudyTimingStats) SetDayOfMonth(v DayOfMonthStats)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *StudyTimingStats) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


