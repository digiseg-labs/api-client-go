# ListApiKeysByAccountId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]ApiKeyItem**](ApiKeyItem.md) |  | [optional] 

## Methods

### NewListApiKeysByAccountId200Response

`func NewListApiKeysByAccountId200Response() *ListApiKeysByAccountId200Response`

NewListApiKeysByAccountId200Response instantiates a new ListApiKeysByAccountId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysByAccountId200ResponseWithDefaults

`func NewListApiKeysByAccountId200ResponseWithDefaults() *ListApiKeysByAccountId200Response`

NewListApiKeysByAccountId200ResponseWithDefaults instantiates a new ListApiKeysByAccountId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListApiKeysByAccountId200Response) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListApiKeysByAccountId200Response) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListApiKeysByAccountId200Response) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListApiKeysByAccountId200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ListApiKeysByAccountId200Response) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListApiKeysByAccountId200Response) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListApiKeysByAccountId200Response) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListApiKeysByAccountId200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ListApiKeysByAccountId200Response) GetData() []ApiKeyItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApiKeysByAccountId200Response) GetDataOk() (*[]ApiKeyItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApiKeysByAccountId200Response) SetData(v []ApiKeyItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListApiKeysByAccountId200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


