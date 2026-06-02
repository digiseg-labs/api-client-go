# SubscriptionPlanFree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the price/plan | [optional] 
**Code** | Pointer to **string** | An optional code, typically provided if the plan/price is public and advertised | [optional] 
**IsPublic** | **bool** | Is the plan/price a public price or custom? | 
**ProductType** | [**SubscriptionProductType**](SubscriptionProductType.md) |  | 
**ListPrice** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**StripeProductId** | Pointer to **string** |  | [optional] [readonly] 
**StripePriceId** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the plan/price that the account is subscribed to | [optional] [readonly] 
**FeatureSet** | [**PlanFeatureSet**](PlanFeatureSet.md) |  | 

## Methods

### NewSubscriptionPlanFree

`func NewSubscriptionPlanFree(isPublic bool, productType SubscriptionProductType, featureSet PlanFeatureSet, ) *SubscriptionPlanFree`

NewSubscriptionPlanFree instantiates a new SubscriptionPlanFree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanFreeWithDefaults

`func NewSubscriptionPlanFreeWithDefaults() *SubscriptionPlanFree`

NewSubscriptionPlanFreeWithDefaults instantiates a new SubscriptionPlanFree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SubscriptionPlanFree) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanFree) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanFree) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanFree) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionPlanFree) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanFree) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanFree) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionPlanFree) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsPublic

`func (o *SubscriptionPlanFree) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubscriptionPlanFree) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubscriptionPlanFree) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetProductType

`func (o *SubscriptionPlanFree) GetProductType() SubscriptionProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SubscriptionPlanFree) GetProductTypeOk() (*SubscriptionProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SubscriptionPlanFree) SetProductType(v SubscriptionProductType)`

SetProductType sets ProductType field to given value.


### GetListPrice

`func (o *SubscriptionPlanFree) GetListPrice() SubscriptionPrice`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *SubscriptionPlanFree) GetListPriceOk() (*SubscriptionPrice, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *SubscriptionPlanFree) SetListPrice(v SubscriptionPrice)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *SubscriptionPlanFree) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetStripeProductId

`func (o *SubscriptionPlanFree) GetStripeProductId() string`

GetStripeProductId returns the StripeProductId field if non-nil, zero value otherwise.

### GetStripeProductIdOk

`func (o *SubscriptionPlanFree) GetStripeProductIdOk() (*string, bool)`

GetStripeProductIdOk returns a tuple with the StripeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeProductId

`func (o *SubscriptionPlanFree) SetStripeProductId(v string)`

SetStripeProductId sets StripeProductId field to given value.

### HasStripeProductId

`func (o *SubscriptionPlanFree) HasStripeProductId() bool`

HasStripeProductId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *SubscriptionPlanFree) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *SubscriptionPlanFree) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *SubscriptionPlanFree) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *SubscriptionPlanFree) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionPlanFree) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanFree) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanFree) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPlanFree) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeatureSet

`func (o *SubscriptionPlanFree) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *SubscriptionPlanFree) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *SubscriptionPlanFree) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


