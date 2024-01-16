# GetApiKeysByUserId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]ApiKeyItem**](ApiKeyItem.md) |  | [optional] 

## Methods

### NewGetApiKeysByUserId200Response

`func NewGetApiKeysByUserId200Response() *GetApiKeysByUserId200Response`

NewGetApiKeysByUserId200Response instantiates a new GetApiKeysByUserId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeysByUserId200ResponseWithDefaults

`func NewGetApiKeysByUserId200ResponseWithDefaults() *GetApiKeysByUserId200Response`

NewGetApiKeysByUserId200ResponseWithDefaults instantiates a new GetApiKeysByUserId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetApiKeysByUserId200Response) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetApiKeysByUserId200Response) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetApiKeysByUserId200Response) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetApiKeysByUserId200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *GetApiKeysByUserId200Response) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetApiKeysByUserId200Response) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetApiKeysByUserId200Response) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetApiKeysByUserId200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *GetApiKeysByUserId200Response) GetData() []ApiKeyItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApiKeysByUserId200Response) GetDataOk() (*[]ApiKeyItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApiKeysByUserId200Response) SetData(v []ApiKeyItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetApiKeysByUserId200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


