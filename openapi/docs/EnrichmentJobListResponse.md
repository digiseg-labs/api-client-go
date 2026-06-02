# EnrichmentJobListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EnrichmentJobItem**](EnrichmentJobItem.md) |  | 
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 

## Methods

### NewEnrichmentJobListResponse

`func NewEnrichmentJobListResponse(data []EnrichmentJobItem, ) *EnrichmentJobListResponse`

NewEnrichmentJobListResponse instantiates a new EnrichmentJobListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobListResponseWithDefaults

`func NewEnrichmentJobListResponseWithDefaults() *EnrichmentJobListResponse`

NewEnrichmentJobListResponseWithDefaults instantiates a new EnrichmentJobListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnrichmentJobListResponse) GetData() []EnrichmentJobItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnrichmentJobListResponse) GetDataOk() (*[]EnrichmentJobItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnrichmentJobListResponse) SetData(v []EnrichmentJobItem)`

SetData sets Data field to given value.


### GetMeta

`func (o *EnrichmentJobListResponse) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnrichmentJobListResponse) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnrichmentJobListResponse) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnrichmentJobListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *EnrichmentJobListResponse) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnrichmentJobListResponse) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnrichmentJobListResponse) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnrichmentJobListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


