# SubscriptionPlanBase

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

## Methods

### NewSubscriptionPlanBase

`func NewSubscriptionPlanBase(isPublic bool, productType SubscriptionProductType, ) *SubscriptionPlanBase`

NewSubscriptionPlanBase instantiates a new SubscriptionPlanBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanBaseWithDefaults

`func NewSubscriptionPlanBaseWithDefaults() *SubscriptionPlanBase`

NewSubscriptionPlanBaseWithDefaults instantiates a new SubscriptionPlanBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SubscriptionPlanBase) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanBase) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanBase) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanBase) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionPlanBase) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanBase) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanBase) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionPlanBase) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsPublic

`func (o *SubscriptionPlanBase) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubscriptionPlanBase) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubscriptionPlanBase) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetProductType

`func (o *SubscriptionPlanBase) GetProductType() SubscriptionProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SubscriptionPlanBase) GetProductTypeOk() (*SubscriptionProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SubscriptionPlanBase) SetProductType(v SubscriptionProductType)`

SetProductType sets ProductType field to given value.


### GetListPrice

`func (o *SubscriptionPlanBase) GetListPrice() SubscriptionPrice`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *SubscriptionPlanBase) GetListPriceOk() (*SubscriptionPrice, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *SubscriptionPlanBase) SetListPrice(v SubscriptionPrice)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *SubscriptionPlanBase) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetStripeProductId

`func (o *SubscriptionPlanBase) GetStripeProductId() string`

GetStripeProductId returns the StripeProductId field if non-nil, zero value otherwise.

### GetStripeProductIdOk

`func (o *SubscriptionPlanBase) GetStripeProductIdOk() (*string, bool)`

GetStripeProductIdOk returns a tuple with the StripeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeProductId

`func (o *SubscriptionPlanBase) SetStripeProductId(v string)`

SetStripeProductId sets StripeProductId field to given value.

### HasStripeProductId

`func (o *SubscriptionPlanBase) HasStripeProductId() bool`

HasStripeProductId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *SubscriptionPlanBase) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *SubscriptionPlanBase) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *SubscriptionPlanBase) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *SubscriptionPlanBase) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


