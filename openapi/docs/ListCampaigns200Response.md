# ListCampaigns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]CampaignItem**](CampaignItem.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 

## Methods

### NewListCampaigns200Response

`func NewListCampaigns200Response() *ListCampaigns200Response`

NewListCampaigns200Response instantiates a new ListCampaigns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCampaigns200ResponseWithDefaults

`func NewListCampaigns200ResponseWithDefaults() *ListCampaigns200Response`

NewListCampaigns200ResponseWithDefaults instantiates a new ListCampaigns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCampaigns200Response) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCampaigns200Response) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCampaigns200Response) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCampaigns200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *ListCampaigns200Response) GetData() []CampaignItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCampaigns200Response) GetDataOk() (*[]CampaignItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCampaigns200Response) SetData(v []CampaignItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListCampaigns200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ListCampaigns200Response) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCampaigns200Response) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCampaigns200Response) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListCampaigns200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


