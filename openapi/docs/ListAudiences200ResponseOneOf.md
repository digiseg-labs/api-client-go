# ListAudiences200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**AudienceSetListResponseMeta**](AudienceSetListResponseMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AudienceSetItem**](AudienceSetItem.md) |  | [optional] 

## Methods

### NewListAudiences200ResponseOneOf

`func NewListAudiences200ResponseOneOf() *ListAudiences200ResponseOneOf`

NewListAudiences200ResponseOneOf instantiates a new ListAudiences200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudiences200ResponseOneOfWithDefaults

`func NewListAudiences200ResponseOneOfWithDefaults() *ListAudiences200ResponseOneOf`

NewListAudiences200ResponseOneOfWithDefaults instantiates a new ListAudiences200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListAudiences200ResponseOneOf) GetMeta() AudienceSetListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAudiences200ResponseOneOf) GetMetaOk() (*AudienceSetListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAudiences200ResponseOneOf) SetMeta(v AudienceSetListResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAudiences200ResponseOneOf) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *ListAudiences200ResponseOneOf) GetData() []AudienceSetItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAudiences200ResponseOneOf) GetDataOk() (*[]AudienceSetItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAudiences200ResponseOneOf) SetData(v []AudienceSetItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListAudiences200ResponseOneOf) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


