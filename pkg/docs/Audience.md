# Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The audience category | [optional] 
**Code** | Pointer to **string** | The unique audience code | [optional] 
**Name** | Pointer to **string** | The name of the audience, typically of the form \&quot;{category name} &gt; {audience name}\&quot; | [optional] 

## Methods

### NewAudience

`func NewAudience() *Audience`

NewAudience instantiates a new Audience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceWithDefaults

`func NewAudienceWithDefaults() *Audience`

NewAudienceWithDefaults instantiates a new Audience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Audience) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Audience) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Audience) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Audience) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCode

`func (o *Audience) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Audience) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Audience) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Audience) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Audience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Audience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Audience) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Audience) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


