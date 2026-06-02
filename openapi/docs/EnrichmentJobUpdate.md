# EnrichmentJobUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for the job | 
**Status** | Pointer to [**EnrichmentJobStatus**](EnrichmentJobStatus.md) |  | [optional] 
**StartingAt** | Pointer to **time.Time** | Date and time of the trigger to start the job | [optional] [readonly] 
**StartedAt** | Pointer to **time.Time** | Date and time of the job started running | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Date and time of the job finishing | [optional] [readonly] 
**FailedAt** | Pointer to **time.Time** | Date and time of the job failing | [optional] [readonly] 
**Result** | Pointer to [**EnrichmentJobResult**](EnrichmentJobResult.md) |  | [optional] 
**EstimatedRecords** | Pointer to **int32** |  | [optional] [readonly] 
**EstimatedIps** | Pointer to **int32** |  | [optional] [readonly] 
**FileSource** | Pointer to [**EnrichmentJobFileSource**](EnrichmentJobFileSource.md) |  | [optional] 

## Methods

### NewEnrichmentJobUpdate

`func NewEnrichmentJobUpdate(name string, ) *EnrichmentJobUpdate`

NewEnrichmentJobUpdate instantiates a new EnrichmentJobUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobUpdateWithDefaults

`func NewEnrichmentJobUpdateWithDefaults() *EnrichmentJobUpdate`

NewEnrichmentJobUpdateWithDefaults instantiates a new EnrichmentJobUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnrichmentJobUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrichmentJobUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrichmentJobUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *EnrichmentJobUpdate) GetStatus() EnrichmentJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrichmentJobUpdate) GetStatusOk() (*EnrichmentJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrichmentJobUpdate) SetStatus(v EnrichmentJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrichmentJobUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartingAt

`func (o *EnrichmentJobUpdate) GetStartingAt() time.Time`

GetStartingAt returns the StartingAt field if non-nil, zero value otherwise.

### GetStartingAtOk

`func (o *EnrichmentJobUpdate) GetStartingAtOk() (*time.Time, bool)`

GetStartingAtOk returns a tuple with the StartingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAt

`func (o *EnrichmentJobUpdate) SetStartingAt(v time.Time)`

SetStartingAt sets StartingAt field to given value.

### HasStartingAt

`func (o *EnrichmentJobUpdate) HasStartingAt() bool`

HasStartingAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *EnrichmentJobUpdate) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnrichmentJobUpdate) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnrichmentJobUpdate) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnrichmentJobUpdate) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *EnrichmentJobUpdate) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *EnrichmentJobUpdate) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *EnrichmentJobUpdate) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *EnrichmentJobUpdate) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetFailedAt

`func (o *EnrichmentJobUpdate) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *EnrichmentJobUpdate) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *EnrichmentJobUpdate) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *EnrichmentJobUpdate) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetResult

`func (o *EnrichmentJobUpdate) GetResult() EnrichmentJobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EnrichmentJobUpdate) GetResultOk() (*EnrichmentJobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EnrichmentJobUpdate) SetResult(v EnrichmentJobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *EnrichmentJobUpdate) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetEstimatedRecords

`func (o *EnrichmentJobUpdate) GetEstimatedRecords() int32`

GetEstimatedRecords returns the EstimatedRecords field if non-nil, zero value otherwise.

### GetEstimatedRecordsOk

`func (o *EnrichmentJobUpdate) GetEstimatedRecordsOk() (*int32, bool)`

GetEstimatedRecordsOk returns a tuple with the EstimatedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRecords

`func (o *EnrichmentJobUpdate) SetEstimatedRecords(v int32)`

SetEstimatedRecords sets EstimatedRecords field to given value.

### HasEstimatedRecords

`func (o *EnrichmentJobUpdate) HasEstimatedRecords() bool`

HasEstimatedRecords returns a boolean if a field has been set.

### GetEstimatedIps

`func (o *EnrichmentJobUpdate) GetEstimatedIps() int32`

GetEstimatedIps returns the EstimatedIps field if non-nil, zero value otherwise.

### GetEstimatedIpsOk

`func (o *EnrichmentJobUpdate) GetEstimatedIpsOk() (*int32, bool)`

GetEstimatedIpsOk returns a tuple with the EstimatedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIps

`func (o *EnrichmentJobUpdate) SetEstimatedIps(v int32)`

SetEstimatedIps sets EstimatedIps field to given value.

### HasEstimatedIps

`func (o *EnrichmentJobUpdate) HasEstimatedIps() bool`

HasEstimatedIps returns a boolean if a field has been set.

### GetFileSource

`func (o *EnrichmentJobUpdate) GetFileSource() EnrichmentJobFileSource`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *EnrichmentJobUpdate) GetFileSourceOk() (*EnrichmentJobFileSource, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *EnrichmentJobUpdate) SetFileSource(v EnrichmentJobFileSource)`

SetFileSource sets FileSource field to given value.

### HasFileSource

`func (o *EnrichmentJobUpdate) HasFileSource() bool`

HasFileSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


