# DataMonthlyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** | Year number | 
**Month** | **int32** | Month number (1-12) | 
**Usage** | [**DataUsage**](DataUsage.md) |  | 

## Methods

### NewDataMonthlyUsage

`func NewDataMonthlyUsage(year int32, month int32, usage DataUsage, ) *DataMonthlyUsage`

NewDataMonthlyUsage instantiates a new DataMonthlyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMonthlyUsageWithDefaults

`func NewDataMonthlyUsageWithDefaults() *DataMonthlyUsage`

NewDataMonthlyUsageWithDefaults instantiates a new DataMonthlyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *DataMonthlyUsage) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *DataMonthlyUsage) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *DataMonthlyUsage) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *DataMonthlyUsage) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *DataMonthlyUsage) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *DataMonthlyUsage) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetUsage

`func (o *DataMonthlyUsage) GetUsage() DataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DataMonthlyUsage) GetUsageOk() (*DataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DataMonthlyUsage) SetUsage(v DataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


