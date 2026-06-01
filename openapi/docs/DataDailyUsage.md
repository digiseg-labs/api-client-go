# DataDailyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date in ISO 8601 format | 
**Usage** | [**DataUsage**](DataUsage.md) |  | 

## Methods

### NewDataDailyUsage

`func NewDataDailyUsage(date string, usage DataUsage, ) *DataDailyUsage`

NewDataDailyUsage instantiates a new DataDailyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDailyUsageWithDefaults

`func NewDataDailyUsageWithDefaults() *DataDailyUsage`

NewDataDailyUsageWithDefaults instantiates a new DataDailyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DataDailyUsage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DataDailyUsage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DataDailyUsage) SetDate(v string)`

SetDate sets Date field to given value.


### GetUsage

`func (o *DataDailyUsage) GetUsage() DataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DataDailyUsage) GetUsageOk() (*DataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DataDailyUsage) SetUsage(v DataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


