# AudienceDataRealtimeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | **time.Time** | The point in time of recording the usage data | 
**Usage** | [**AudienceDataUsage**](AudienceDataUsage.md) |  | 

## Methods

### NewAudienceDataRealtimeUsage

`func NewAudienceDataRealtimeUsage(when time.Time, usage AudienceDataUsage, ) *AudienceDataRealtimeUsage`

NewAudienceDataRealtimeUsage instantiates a new AudienceDataRealtimeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceDataRealtimeUsageWithDefaults

`func NewAudienceDataRealtimeUsageWithDefaults() *AudienceDataRealtimeUsage`

NewAudienceDataRealtimeUsageWithDefaults instantiates a new AudienceDataRealtimeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *AudienceDataRealtimeUsage) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *AudienceDataRealtimeUsage) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *AudienceDataRealtimeUsage) SetWhen(v time.Time)`

SetWhen sets When field to given value.


### GetUsage

`func (o *AudienceDataRealtimeUsage) GetUsage() AudienceDataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AudienceDataRealtimeUsage) GetUsageOk() (*AudienceDataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AudienceDataRealtimeUsage) SetUsage(v AudienceDataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


