# Comparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the comparison source | 
**Code** | **string** | The code of the comparison source | 
**FractionOfTotal** | **float64** | The \&quot;fraction of total\&quot; value that is being compared with.  | 
**Index** | **int32** | The comparison index where 100 means that the measurement is completely aligned with the compared metric. Values below 100 means that the measurement is below the compared metric, and values above 100 means that the measurement is above the compared metric.  | 

## Methods

### NewComparison

`func NewComparison(name string, code string, fractionOfTotal float64, index int32, ) *Comparison`

NewComparison instantiates a new Comparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonWithDefaults

`func NewComparisonWithDefaults() *Comparison`

NewComparisonWithDefaults instantiates a new Comparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Comparison) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Comparison) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Comparison) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *Comparison) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Comparison) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Comparison) SetCode(v string)`

SetCode sets Code field to given value.


### GetFractionOfTotal

`func (o *Comparison) GetFractionOfTotal() float64`

GetFractionOfTotal returns the FractionOfTotal field if non-nil, zero value otherwise.

### GetFractionOfTotalOk

`func (o *Comparison) GetFractionOfTotalOk() (*float64, bool)`

GetFractionOfTotalOk returns a tuple with the FractionOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionOfTotal

`func (o *Comparison) SetFractionOfTotal(v float64)`

SetFractionOfTotal sets FractionOfTotal field to given value.


### GetIndex

`func (o *Comparison) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Comparison) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Comparison) SetIndex(v int32)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


