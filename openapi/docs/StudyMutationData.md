# StudyMutationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSet** | Pointer to [**MeasurementEventSet**](MeasurementEventSet.md) |  | [optional] 
**ClientId** | Pointer to **string** | The ID of the measurement client that this study is for | [optional] 

## Methods

### NewStudyMutationData

`func NewStudyMutationData() *StudyMutationData`

NewStudyMutationData instantiates a new StudyMutationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyMutationDataWithDefaults

`func NewStudyMutationDataWithDefaults() *StudyMutationData`

NewStudyMutationDataWithDefaults instantiates a new StudyMutationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSet

`func (o *StudyMutationData) GetEventSet() MeasurementEventSet`

GetEventSet returns the EventSet field if non-nil, zero value otherwise.

### GetEventSetOk

`func (o *StudyMutationData) GetEventSetOk() (*MeasurementEventSet, bool)`

GetEventSetOk returns a tuple with the EventSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSet

`func (o *StudyMutationData) SetEventSet(v MeasurementEventSet)`

SetEventSet sets EventSet field to given value.

### HasEventSet

`func (o *StudyMutationData) HasEventSet() bool`

HasEventSet returns a boolean if a field has been set.

### GetClientId

`func (o *StudyMutationData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StudyMutationData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StudyMutationData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StudyMutationData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


