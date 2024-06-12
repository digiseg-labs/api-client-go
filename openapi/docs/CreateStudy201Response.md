# CreateStudy201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**StudyFull**](StudyFull.md) |  | [optional] 
**Links** | Pointer to [**StudyLinks**](StudyLinks.md) |  | [optional] 
**Meta** | Pointer to [**StudyMeta**](StudyMeta.md) |  | [optional] 

## Methods

### NewCreateStudy201Response

`func NewCreateStudy201Response() *CreateStudy201Response`

NewCreateStudy201Response instantiates a new CreateStudy201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStudy201ResponseWithDefaults

`func NewCreateStudy201ResponseWithDefaults() *CreateStudy201Response`

NewCreateStudy201ResponseWithDefaults instantiates a new CreateStudy201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateStudy201Response) GetData() StudyFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateStudy201Response) GetDataOk() (*StudyFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateStudy201Response) SetData(v StudyFull)`

SetData sets Data field to given value.

### HasData

`func (o *CreateStudy201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CreateStudy201Response) GetLinks() StudyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateStudy201Response) GetLinksOk() (*StudyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateStudy201Response) SetLinks(v StudyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateStudy201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *CreateStudy201Response) GetMeta() StudyMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateStudy201Response) GetMetaOk() (*StudyMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateStudy201Response) SetMeta(v StudyMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateStudy201Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


