# EnrichmentJobBase

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

## Methods

### NewEnrichmentJobBase

`func NewEnrichmentJobBase(name string, ) *EnrichmentJobBase`

NewEnrichmentJobBase instantiates a new EnrichmentJobBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobBaseWithDefaults

`func NewEnrichmentJobBaseWithDefaults() *EnrichmentJobBase`

NewEnrichmentJobBaseWithDefaults instantiates a new EnrichmentJobBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnrichmentJobBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrichmentJobBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrichmentJobBase) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *EnrichmentJobBase) GetStatus() EnrichmentJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrichmentJobBase) GetStatusOk() (*EnrichmentJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrichmentJobBase) SetStatus(v EnrichmentJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrichmentJobBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartingAt

`func (o *EnrichmentJobBase) GetStartingAt() time.Time`

GetStartingAt returns the StartingAt field if non-nil, zero value otherwise.

### GetStartingAtOk

`func (o *EnrichmentJobBase) GetStartingAtOk() (*time.Time, bool)`

GetStartingAtOk returns a tuple with the StartingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAt

`func (o *EnrichmentJobBase) SetStartingAt(v time.Time)`

SetStartingAt sets StartingAt field to given value.

### HasStartingAt

`func (o *EnrichmentJobBase) HasStartingAt() bool`

HasStartingAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *EnrichmentJobBase) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnrichmentJobBase) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnrichmentJobBase) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnrichmentJobBase) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *EnrichmentJobBase) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *EnrichmentJobBase) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *EnrichmentJobBase) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *EnrichmentJobBase) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetFailedAt

`func (o *EnrichmentJobBase) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *EnrichmentJobBase) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *EnrichmentJobBase) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *EnrichmentJobBase) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetResult

`func (o *EnrichmentJobBase) GetResult() EnrichmentJobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EnrichmentJobBase) GetResultOk() (*EnrichmentJobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EnrichmentJobBase) SetResult(v EnrichmentJobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *EnrichmentJobBase) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


