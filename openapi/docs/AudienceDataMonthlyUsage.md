# AudienceDataMonthlyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** | Year number | 
**Month** | **int32** | Month number (1-12) | 
**Usage** | [**AudienceDataUsage**](AudienceDataUsage.md) |  | 

## Methods

### NewAudienceDataMonthlyUsage

`func NewAudienceDataMonthlyUsage(year int32, month int32, usage AudienceDataUsage, ) *AudienceDataMonthlyUsage`

NewAudienceDataMonthlyUsage instantiates a new AudienceDataMonthlyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceDataMonthlyUsageWithDefaults

`func NewAudienceDataMonthlyUsageWithDefaults() *AudienceDataMonthlyUsage`

NewAudienceDataMonthlyUsageWithDefaults instantiates a new AudienceDataMonthlyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *AudienceDataMonthlyUsage) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AudienceDataMonthlyUsage) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AudienceDataMonthlyUsage) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *AudienceDataMonthlyUsage) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *AudienceDataMonthlyUsage) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *AudienceDataMonthlyUsage) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetUsage

`func (o *AudienceDataMonthlyUsage) GetUsage() AudienceDataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AudienceDataMonthlyUsage) GetUsageOk() (*AudienceDataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AudienceDataMonthlyUsage) SetUsage(v AudienceDataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


