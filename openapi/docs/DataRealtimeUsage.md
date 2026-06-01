# DataRealtimeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | **time.Time** | The point in time of recording the usage data | 
**Usage** | [**DataUsage**](DataUsage.md) |  | 

## Methods

### NewDataRealtimeUsage

`func NewDataRealtimeUsage(when time.Time, usage DataUsage, ) *DataRealtimeUsage`

NewDataRealtimeUsage instantiates a new DataRealtimeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRealtimeUsageWithDefaults

`func NewDataRealtimeUsageWithDefaults() *DataRealtimeUsage`

NewDataRealtimeUsageWithDefaults instantiates a new DataRealtimeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *DataRealtimeUsage) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *DataRealtimeUsage) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *DataRealtimeUsage) SetWhen(v time.Time)`

SetWhen sets When field to given value.


### GetUsage

`func (o *DataRealtimeUsage) GetUsage() DataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DataRealtimeUsage) GetUsageOk() (*DataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DataRealtimeUsage) SetUsage(v DataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


