# ListAudiences200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**AudienceSetListResponseMeta**](AudienceSetListResponseMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AudienceSetItem**](AudienceSetItem.md) |  | [optional] 

## Methods

### NewListAudiences200Response

`func NewListAudiences200Response() *ListAudiences200Response`

NewListAudiences200Response instantiates a new ListAudiences200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudiences200ResponseWithDefaults

`func NewListAudiences200ResponseWithDefaults() *ListAudiences200Response`

NewListAudiences200ResponseWithDefaults instantiates a new ListAudiences200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListAudiences200Response) GetMeta() AudienceSetListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAudiences200Response) GetMetaOk() (*AudienceSetListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAudiences200Response) SetMeta(v AudienceSetListResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAudiences200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *ListAudiences200Response) GetData() []AudienceSetItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAudiences200Response) GetDataOk() (*[]AudienceSetItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAudiences200Response) SetData(v []AudienceSetItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListAudiences200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


