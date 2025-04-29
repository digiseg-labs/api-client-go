# StudyOlapFilterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K** | [**StudyOlapDimensionKey**](StudyOlapDimensionKey.md) |  | 
**V** | **string** | Represents a value for a &#x60;StudyOlapDimensionKey&#x60;. | 

## Methods

### NewStudyOlapFilterItem

`func NewStudyOlapFilterItem(k StudyOlapDimensionKey, v string, ) *StudyOlapFilterItem`

NewStudyOlapFilterItem instantiates a new StudyOlapFilterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyOlapFilterItemWithDefaults

`func NewStudyOlapFilterItemWithDefaults() *StudyOlapFilterItem`

NewStudyOlapFilterItemWithDefaults instantiates a new StudyOlapFilterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK

`func (o *StudyOlapFilterItem) GetK() StudyOlapDimensionKey`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *StudyOlapFilterItem) GetKOk() (*StudyOlapDimensionKey, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *StudyOlapFilterItem) SetK(v StudyOlapDimensionKey)`

SetK sets K field to given value.


### GetV

`func (o *StudyOlapFilterItem) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *StudyOlapFilterItem) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *StudyOlapFilterItem) SetV(v string)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


