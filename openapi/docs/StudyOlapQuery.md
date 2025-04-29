# StudyOlapQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterTimeFrom** | Pointer to **time.Time** | The from-date of the events to query | [optional] 
**FilterTimeTo** | Pointer to **time.Time** | The to-date of the events to query | [optional] 
**GroupBy** | Pointer to [**[]StudyOlapDimensionKey**](StudyOlapDimensionKey.md) |  | [optional] 
**FilterDimensionsInclude** | Pointer to [**[]StudyOlapFilterItem**](StudyOlapFilterItem.md) |  | [optional] 
**FilterDimensionsExclude** | Pointer to [**[]StudyOlapFilterItem**](StudyOlapFilterItem.md) |  | [optional] 

## Methods

### NewStudyOlapQuery

`func NewStudyOlapQuery() *StudyOlapQuery`

NewStudyOlapQuery instantiates a new StudyOlapQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyOlapQueryWithDefaults

`func NewStudyOlapQueryWithDefaults() *StudyOlapQuery`

NewStudyOlapQueryWithDefaults instantiates a new StudyOlapQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterTimeFrom

`func (o *StudyOlapQuery) GetFilterTimeFrom() time.Time`

GetFilterTimeFrom returns the FilterTimeFrom field if non-nil, zero value otherwise.

### GetFilterTimeFromOk

`func (o *StudyOlapQuery) GetFilterTimeFromOk() (*time.Time, bool)`

GetFilterTimeFromOk returns a tuple with the FilterTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTimeFrom

`func (o *StudyOlapQuery) SetFilterTimeFrom(v time.Time)`

SetFilterTimeFrom sets FilterTimeFrom field to given value.

### HasFilterTimeFrom

`func (o *StudyOlapQuery) HasFilterTimeFrom() bool`

HasFilterTimeFrom returns a boolean if a field has been set.

### GetFilterTimeTo

`func (o *StudyOlapQuery) GetFilterTimeTo() time.Time`

GetFilterTimeTo returns the FilterTimeTo field if non-nil, zero value otherwise.

### GetFilterTimeToOk

`func (o *StudyOlapQuery) GetFilterTimeToOk() (*time.Time, bool)`

GetFilterTimeToOk returns a tuple with the FilterTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTimeTo

`func (o *StudyOlapQuery) SetFilterTimeTo(v time.Time)`

SetFilterTimeTo sets FilterTimeTo field to given value.

### HasFilterTimeTo

`func (o *StudyOlapQuery) HasFilterTimeTo() bool`

HasFilterTimeTo returns a boolean if a field has been set.

### GetGroupBy

`func (o *StudyOlapQuery) GetGroupBy() []StudyOlapDimensionKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *StudyOlapQuery) GetGroupByOk() (*[]StudyOlapDimensionKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *StudyOlapQuery) SetGroupBy(v []StudyOlapDimensionKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *StudyOlapQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilterDimensionsInclude

`func (o *StudyOlapQuery) GetFilterDimensionsInclude() []StudyOlapFilterItem`

GetFilterDimensionsInclude returns the FilterDimensionsInclude field if non-nil, zero value otherwise.

### GetFilterDimensionsIncludeOk

`func (o *StudyOlapQuery) GetFilterDimensionsIncludeOk() (*[]StudyOlapFilterItem, bool)`

GetFilterDimensionsIncludeOk returns a tuple with the FilterDimensionsInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDimensionsInclude

`func (o *StudyOlapQuery) SetFilterDimensionsInclude(v []StudyOlapFilterItem)`

SetFilterDimensionsInclude sets FilterDimensionsInclude field to given value.

### HasFilterDimensionsInclude

`func (o *StudyOlapQuery) HasFilterDimensionsInclude() bool`

HasFilterDimensionsInclude returns a boolean if a field has been set.

### GetFilterDimensionsExclude

`func (o *StudyOlapQuery) GetFilterDimensionsExclude() []StudyOlapFilterItem`

GetFilterDimensionsExclude returns the FilterDimensionsExclude field if non-nil, zero value otherwise.

### GetFilterDimensionsExcludeOk

`func (o *StudyOlapQuery) GetFilterDimensionsExcludeOk() (*[]StudyOlapFilterItem, bool)`

GetFilterDimensionsExcludeOk returns a tuple with the FilterDimensionsExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDimensionsExclude

`func (o *StudyOlapQuery) SetFilterDimensionsExclude(v []StudyOlapFilterItem)`

SetFilterDimensionsExclude sets FilterDimensionsExclude field to given value.

### HasFilterDimensionsExclude

`func (o *StudyOlapQuery) HasFilterDimensionsExclude() bool`

HasFilterDimensionsExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


