# EnrichmentJobAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedRecords** | Pointer to **int32** |  | [optional] [readonly] 
**EstimatedIps** | Pointer to **int32** |  | [optional] [readonly] 
**FileSource** | Pointer to [**EnrichmentJobFileSource**](EnrichmentJobFileSource.md) |  | [optional] 

## Methods

### NewEnrichmentJobAux

`func NewEnrichmentJobAux() *EnrichmentJobAux`

NewEnrichmentJobAux instantiates a new EnrichmentJobAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobAuxWithDefaults

`func NewEnrichmentJobAuxWithDefaults() *EnrichmentJobAux`

NewEnrichmentJobAuxWithDefaults instantiates a new EnrichmentJobAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedRecords

`func (o *EnrichmentJobAux) GetEstimatedRecords() int32`

GetEstimatedRecords returns the EstimatedRecords field if non-nil, zero value otherwise.

### GetEstimatedRecordsOk

`func (o *EnrichmentJobAux) GetEstimatedRecordsOk() (*int32, bool)`

GetEstimatedRecordsOk returns a tuple with the EstimatedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRecords

`func (o *EnrichmentJobAux) SetEstimatedRecords(v int32)`

SetEstimatedRecords sets EstimatedRecords field to given value.

### HasEstimatedRecords

`func (o *EnrichmentJobAux) HasEstimatedRecords() bool`

HasEstimatedRecords returns a boolean if a field has been set.

### GetEstimatedIps

`func (o *EnrichmentJobAux) GetEstimatedIps() int32`

GetEstimatedIps returns the EstimatedIps field if non-nil, zero value otherwise.

### GetEstimatedIpsOk

`func (o *EnrichmentJobAux) GetEstimatedIpsOk() (*int32, bool)`

GetEstimatedIpsOk returns a tuple with the EstimatedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIps

`func (o *EnrichmentJobAux) SetEstimatedIps(v int32)`

SetEstimatedIps sets EstimatedIps field to given value.

### HasEstimatedIps

`func (o *EnrichmentJobAux) HasEstimatedIps() bool`

HasEstimatedIps returns a boolean if a field has been set.

### GetFileSource

`func (o *EnrichmentJobAux) GetFileSource() EnrichmentJobFileSource`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *EnrichmentJobAux) GetFileSourceOk() (*EnrichmentJobFileSource, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *EnrichmentJobAux) SetFileSource(v EnrichmentJobFileSource)`

SetFileSource sets FileSource field to given value.

### HasFileSource

`func (o *EnrichmentJobAux) HasFileSource() bool`

HasFileSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


