# DataUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFulfilled** | Pointer to **int64** | Number of records fulfilled | [optional] 
**RecordsProcessed** | Pointer to **int64** | Number of records processed | [optional] 

## Methods

### NewDataUsage

`func NewDataUsage() *DataUsage`

NewDataUsage instantiates a new DataUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUsageWithDefaults

`func NewDataUsageWithDefaults() *DataUsage`

NewDataUsageWithDefaults instantiates a new DataUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFulfilled

`func (o *DataUsage) GetRecordsFulfilled() int64`

GetRecordsFulfilled returns the RecordsFulfilled field if non-nil, zero value otherwise.

### GetRecordsFulfilledOk

`func (o *DataUsage) GetRecordsFulfilledOk() (*int64, bool)`

GetRecordsFulfilledOk returns a tuple with the RecordsFulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFulfilled

`func (o *DataUsage) SetRecordsFulfilled(v int64)`

SetRecordsFulfilled sets RecordsFulfilled field to given value.

### HasRecordsFulfilled

`func (o *DataUsage) HasRecordsFulfilled() bool`

HasRecordsFulfilled returns a boolean if a field has been set.

### GetRecordsProcessed

`func (o *DataUsage) GetRecordsProcessed() int64`

GetRecordsProcessed returns the RecordsProcessed field if non-nil, zero value otherwise.

### GetRecordsProcessedOk

`func (o *DataUsage) GetRecordsProcessedOk() (*int64, bool)`

GetRecordsProcessedOk returns a tuple with the RecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsProcessed

`func (o *DataUsage) SetRecordsProcessed(v int64)`

SetRecordsProcessed sets RecordsProcessed field to given value.

### HasRecordsProcessed

`func (o *DataUsage) HasRecordsProcessed() bool`

HasRecordsProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


