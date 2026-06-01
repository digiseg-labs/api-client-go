# EnrichmentJobFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
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
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 

## Methods

### NewEnrichmentJobFull

`func NewEnrichmentJobFull(name string, ) *EnrichmentJobFull`

NewEnrichmentJobFull instantiates a new EnrichmentJobFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobFullWithDefaults

`func NewEnrichmentJobFullWithDefaults() *EnrichmentJobFull`

NewEnrichmentJobFullWithDefaults instantiates a new EnrichmentJobFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrichmentJobFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrichmentJobFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrichmentJobFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrichmentJobFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnrichmentJobFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrichmentJobFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrichmentJobFull) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *EnrichmentJobFull) GetStatus() EnrichmentJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrichmentJobFull) GetStatusOk() (*EnrichmentJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrichmentJobFull) SetStatus(v EnrichmentJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrichmentJobFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartingAt

`func (o *EnrichmentJobFull) GetStartingAt() time.Time`

GetStartingAt returns the StartingAt field if non-nil, zero value otherwise.

### GetStartingAtOk

`func (o *EnrichmentJobFull) GetStartingAtOk() (*time.Time, bool)`

GetStartingAtOk returns a tuple with the StartingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAt

`func (o *EnrichmentJobFull) SetStartingAt(v time.Time)`

SetStartingAt sets StartingAt field to given value.

### HasStartingAt

`func (o *EnrichmentJobFull) HasStartingAt() bool`

HasStartingAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *EnrichmentJobFull) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnrichmentJobFull) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnrichmentJobFull) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnrichmentJobFull) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *EnrichmentJobFull) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *EnrichmentJobFull) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *EnrichmentJobFull) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *EnrichmentJobFull) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetFailedAt

`func (o *EnrichmentJobFull) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *EnrichmentJobFull) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *EnrichmentJobFull) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *EnrichmentJobFull) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetResult

`func (o *EnrichmentJobFull) GetResult() EnrichmentJobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EnrichmentJobFull) GetResultOk() (*EnrichmentJobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EnrichmentJobFull) SetResult(v EnrichmentJobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *EnrichmentJobFull) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetEstimatedRecords

`func (o *EnrichmentJobFull) GetEstimatedRecords() int32`

GetEstimatedRecords returns the EstimatedRecords field if non-nil, zero value otherwise.

### GetEstimatedRecordsOk

`func (o *EnrichmentJobFull) GetEstimatedRecordsOk() (*int32, bool)`

GetEstimatedRecordsOk returns a tuple with the EstimatedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRecords

`func (o *EnrichmentJobFull) SetEstimatedRecords(v int32)`

SetEstimatedRecords sets EstimatedRecords field to given value.

### HasEstimatedRecords

`func (o *EnrichmentJobFull) HasEstimatedRecords() bool`

HasEstimatedRecords returns a boolean if a field has been set.

### GetEstimatedIps

`func (o *EnrichmentJobFull) GetEstimatedIps() int32`

GetEstimatedIps returns the EstimatedIps field if non-nil, zero value otherwise.

### GetEstimatedIpsOk

`func (o *EnrichmentJobFull) GetEstimatedIpsOk() (*int32, bool)`

GetEstimatedIpsOk returns a tuple with the EstimatedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIps

`func (o *EnrichmentJobFull) SetEstimatedIps(v int32)`

SetEstimatedIps sets EstimatedIps field to given value.

### HasEstimatedIps

`func (o *EnrichmentJobFull) HasEstimatedIps() bool`

HasEstimatedIps returns a boolean if a field has been set.

### GetFileSource

`func (o *EnrichmentJobFull) GetFileSource() EnrichmentJobFileSource`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *EnrichmentJobFull) GetFileSourceOk() (*EnrichmentJobFileSource, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *EnrichmentJobFull) SetFileSource(v EnrichmentJobFileSource)`

SetFileSource sets FileSource field to given value.

### HasFileSource

`func (o *EnrichmentJobFull) HasFileSource() bool`

HasFileSource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnrichmentJobFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnrichmentJobFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnrichmentJobFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnrichmentJobFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EnrichmentJobFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnrichmentJobFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnrichmentJobFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnrichmentJobFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnrichmentJobFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnrichmentJobFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnrichmentJobFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnrichmentJobFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *EnrichmentJobFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnrichmentJobFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnrichmentJobFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnrichmentJobFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


