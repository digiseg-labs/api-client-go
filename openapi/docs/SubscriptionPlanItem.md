# SubscriptionPlanItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the price/plan | [optional] 
**Code** | Pointer to **string** | An optional code, typically provided if the plan/price is public and advertised | [optional] 
**IsPublic** | **bool** | Is the plan/price a public price or custom? | 
**ProductType** | [**SubscriptionProductType**](SubscriptionProductType.md) |  | 
**ListPrice** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**StripeProductId** | Pointer to **string** |  | [optional] [readonly] 
**StripePriceId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSubscriptionPlanItem

`func NewSubscriptionPlanItem(isPublic bool, productType SubscriptionProductType, ) *SubscriptionPlanItem`

NewSubscriptionPlanItem instantiates a new SubscriptionPlanItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanItemWithDefaults

`func NewSubscriptionPlanItemWithDefaults() *SubscriptionPlanItem`

NewSubscriptionPlanItemWithDefaults instantiates a new SubscriptionPlanItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPlanItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPlanItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionPlanItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionPlanItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionPlanItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsPublic

`func (o *SubscriptionPlanItem) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubscriptionPlanItem) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubscriptionPlanItem) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetProductType

`func (o *SubscriptionPlanItem) GetProductType() SubscriptionProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SubscriptionPlanItem) GetProductTypeOk() (*SubscriptionProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SubscriptionPlanItem) SetProductType(v SubscriptionProductType)`

SetProductType sets ProductType field to given value.


### GetListPrice

`func (o *SubscriptionPlanItem) GetListPrice() SubscriptionPrice`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *SubscriptionPlanItem) GetListPriceOk() (*SubscriptionPrice, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *SubscriptionPlanItem) SetListPrice(v SubscriptionPrice)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *SubscriptionPlanItem) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetStripeProductId

`func (o *SubscriptionPlanItem) GetStripeProductId() string`

GetStripeProductId returns the StripeProductId field if non-nil, zero value otherwise.

### GetStripeProductIdOk

`func (o *SubscriptionPlanItem) GetStripeProductIdOk() (*string, bool)`

GetStripeProductIdOk returns a tuple with the StripeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeProductId

`func (o *SubscriptionPlanItem) SetStripeProductId(v string)`

SetStripeProductId sets StripeProductId field to given value.

### HasStripeProductId

`func (o *SubscriptionPlanItem) HasStripeProductId() bool`

HasStripeProductId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *SubscriptionPlanItem) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *SubscriptionPlanItem) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *SubscriptionPlanItem) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *SubscriptionPlanItem) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


