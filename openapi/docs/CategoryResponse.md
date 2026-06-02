# CategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AudienceCategoryItem**](AudienceCategoryItem.md) |  | [optional] 
**Meta** | Pointer to [**AudienceResponseMeta**](AudienceResponseMeta.md) |  | [optional] 

## Methods

### NewCategoryResponse

`func NewCategoryResponse() *CategoryResponse`

NewCategoryResponse instantiates a new CategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryResponseWithDefaults

`func NewCategoryResponseWithDefaults() *CategoryResponse`

NewCategoryResponseWithDefaults instantiates a new CategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CategoryResponse) GetData() AudienceCategoryItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CategoryResponse) GetDataOk() (*AudienceCategoryItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CategoryResponse) SetData(v AudienceCategoryItem)`

SetData sets Data field to given value.

### HasData

`func (o *CategoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *CategoryResponse) GetMeta() AudienceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CategoryResponse) GetMetaOk() (*AudienceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CategoryResponse) SetMeta(v AudienceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CategoryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


