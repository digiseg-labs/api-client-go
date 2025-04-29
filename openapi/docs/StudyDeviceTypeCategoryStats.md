# StudyDeviceTypeCategoryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**StudyDeviceTypeCategory**](StudyDeviceTypeCategory.md) |  | 
**Count** | **int32** | The total count of events recorded for the device category. | 
**Impressions** | Pointer to **int32** | The total count of impressions recorded for the device category. | [optional] 
**Clicks** | Pointer to **int32** | The total count of clicks recorded for the device category. | [optional] 
**SubTypes** | Pointer to [**[]StudyDeviceSubTypeStats**](StudyDeviceSubTypeStats.md) |  | [optional] 

## Methods

### NewStudyDeviceTypeCategoryStats

`func NewStudyDeviceTypeCategoryStats(category StudyDeviceTypeCategory, count int32, ) *StudyDeviceTypeCategoryStats`

NewStudyDeviceTypeCategoryStats instantiates a new StudyDeviceTypeCategoryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyDeviceTypeCategoryStatsWithDefaults

`func NewStudyDeviceTypeCategoryStatsWithDefaults() *StudyDeviceTypeCategoryStats`

NewStudyDeviceTypeCategoryStatsWithDefaults instantiates a new StudyDeviceTypeCategoryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *StudyDeviceTypeCategoryStats) GetCategory() StudyDeviceTypeCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StudyDeviceTypeCategoryStats) GetCategoryOk() (*StudyDeviceTypeCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StudyDeviceTypeCategoryStats) SetCategory(v StudyDeviceTypeCategory)`

SetCategory sets Category field to given value.


### GetCount

`func (o *StudyDeviceTypeCategoryStats) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StudyDeviceTypeCategoryStats) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StudyDeviceTypeCategoryStats) SetCount(v int32)`

SetCount sets Count field to given value.


### GetImpressions

`func (o *StudyDeviceTypeCategoryStats) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *StudyDeviceTypeCategoryStats) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *StudyDeviceTypeCategoryStats) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *StudyDeviceTypeCategoryStats) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetClicks

`func (o *StudyDeviceTypeCategoryStats) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *StudyDeviceTypeCategoryStats) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *StudyDeviceTypeCategoryStats) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *StudyDeviceTypeCategoryStats) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetSubTypes

`func (o *StudyDeviceTypeCategoryStats) GetSubTypes() []StudyDeviceSubTypeStats`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *StudyDeviceTypeCategoryStats) GetSubTypesOk() (*[]StudyDeviceSubTypeStats, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *StudyDeviceTypeCategoryStats) SetSubTypes(v []StudyDeviceSubTypeStats)`

SetSubTypes sets SubTypes field to given value.

### HasSubTypes

`func (o *StudyDeviceTypeCategoryStats) HasSubTypes() bool`

HasSubTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


