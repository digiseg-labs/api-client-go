# ListSubscriptionOffers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SubscriptionOfferItem**](SubscriptionOfferItem.md) |  | [optional] 
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 

## Methods

### NewListSubscriptionOffers200Response

`func NewListSubscriptionOffers200Response() *ListSubscriptionOffers200Response`

NewListSubscriptionOffers200Response instantiates a new ListSubscriptionOffers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionOffers200ResponseWithDefaults

`func NewListSubscriptionOffers200ResponseWithDefaults() *ListSubscriptionOffers200Response`

NewListSubscriptionOffers200ResponseWithDefaults instantiates a new ListSubscriptionOffers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSubscriptionOffers200Response) GetData() []SubscriptionOfferItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSubscriptionOffers200Response) GetDataOk() (*[]SubscriptionOfferItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSubscriptionOffers200Response) SetData(v []SubscriptionOfferItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListSubscriptionOffers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListSubscriptionOffers200Response) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubscriptionOffers200Response) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubscriptionOffers200Response) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubscriptionOffers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ListSubscriptionOffers200Response) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSubscriptionOffers200Response) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSubscriptionOffers200Response) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSubscriptionOffers200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


