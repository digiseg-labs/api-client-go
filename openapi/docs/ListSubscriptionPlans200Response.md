# ListSubscriptionPlans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SubscriptionPlanItem**](SubscriptionPlanItem.md) |  | [optional] 
**Meta** | Pointer to [**ListPaginationMeta**](ListPaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**ListPaginationLinks**](ListPaginationLinks.md) |  | [optional] 

## Methods

### NewListSubscriptionPlans200Response

`func NewListSubscriptionPlans200Response() *ListSubscriptionPlans200Response`

NewListSubscriptionPlans200Response instantiates a new ListSubscriptionPlans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionPlans200ResponseWithDefaults

`func NewListSubscriptionPlans200ResponseWithDefaults() *ListSubscriptionPlans200Response`

NewListSubscriptionPlans200ResponseWithDefaults instantiates a new ListSubscriptionPlans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSubscriptionPlans200Response) GetData() []SubscriptionPlanItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSubscriptionPlans200Response) GetDataOk() (*[]SubscriptionPlanItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSubscriptionPlans200Response) SetData(v []SubscriptionPlanItem)`

SetData sets Data field to given value.

### HasData

`func (o *ListSubscriptionPlans200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListSubscriptionPlans200Response) GetMeta() ListPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubscriptionPlans200Response) GetMetaOk() (*ListPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubscriptionPlans200Response) SetMeta(v ListPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubscriptionPlans200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ListSubscriptionPlans200Response) GetLinks() ListPaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSubscriptionPlans200Response) GetLinksOk() (*ListPaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSubscriptionPlans200Response) SetLinks(v ListPaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSubscriptionPlans200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


