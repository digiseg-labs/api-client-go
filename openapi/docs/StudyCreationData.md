# StudyCreationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSet** | [**MeasurementEventSet**](MeasurementEventSet.md) |  | 
**ClientId** | Pointer to **string** | The ID of the measurement client that this study is for | [optional] 

## Methods

### NewStudyCreationData

`func NewStudyCreationData(eventSet MeasurementEventSet, ) *StudyCreationData`

NewStudyCreationData instantiates a new StudyCreationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyCreationDataWithDefaults

`func NewStudyCreationDataWithDefaults() *StudyCreationData`

NewStudyCreationDataWithDefaults instantiates a new StudyCreationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSet

`func (o *StudyCreationData) GetEventSet() MeasurementEventSet`

GetEventSet returns the EventSet field if non-nil, zero value otherwise.

### GetEventSetOk

`func (o *StudyCreationData) GetEventSetOk() (*MeasurementEventSet, bool)`

GetEventSetOk returns a tuple with the EventSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSet

`func (o *StudyCreationData) SetEventSet(v MeasurementEventSet)`

SetEventSet sets EventSet field to given value.


### GetClientId

`func (o *StudyCreationData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StudyCreationData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StudyCreationData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StudyCreationData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


