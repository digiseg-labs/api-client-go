# SubscriptionPlanFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the plan/price that the account is subscribed to | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the price/plan | [optional] 
**Code** | Pointer to **string** | An optional code, typically provided if the plan/price is public and advertised | [optional] 
**IsPublic** | **bool** | Is the plan/price a public price or custom? | 
**ProductType** | [**SubscriptionProductType**](SubscriptionProductType.md) |  | 
**ListPrice** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**StripeProductId** | Pointer to **string** |  | [optional] [readonly] 
**StripePriceId** | Pointer to **string** |  | [optional] [readonly] 
**FeatureSet** | [**PlanFeatureSet**](PlanFeatureSet.md) |  | 

## Methods

### NewSubscriptionPlanFull

`func NewSubscriptionPlanFull(isPublic bool, productType SubscriptionProductType, featureSet PlanFeatureSet, ) *SubscriptionPlanFull`

NewSubscriptionPlanFull instantiates a new SubscriptionPlanFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanFullWithDefaults

`func NewSubscriptionPlanFullWithDefaults() *SubscriptionPlanFull`

NewSubscriptionPlanFullWithDefaults instantiates a new SubscriptionPlanFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPlanFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPlanFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionPlanFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionPlanFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionPlanFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionPlanFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SubscriptionPlanFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubscriptionPlanFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubscriptionPlanFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SubscriptionPlanFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionPlanFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionPlanFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionPlanFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionPlanFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SubscriptionPlanFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SubscriptionPlanFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SubscriptionPlanFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SubscriptionPlanFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionPlanFull) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanFull) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanFull) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanFull) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionPlanFull) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanFull) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanFull) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionPlanFull) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsPublic

`func (o *SubscriptionPlanFull) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubscriptionPlanFull) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubscriptionPlanFull) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetProductType

`func (o *SubscriptionPlanFull) GetProductType() SubscriptionProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SubscriptionPlanFull) GetProductTypeOk() (*SubscriptionProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SubscriptionPlanFull) SetProductType(v SubscriptionProductType)`

SetProductType sets ProductType field to given value.


### GetListPrice

`func (o *SubscriptionPlanFull) GetListPrice() SubscriptionPrice`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *SubscriptionPlanFull) GetListPriceOk() (*SubscriptionPrice, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *SubscriptionPlanFull) SetListPrice(v SubscriptionPrice)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *SubscriptionPlanFull) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetStripeProductId

`func (o *SubscriptionPlanFull) GetStripeProductId() string`

GetStripeProductId returns the StripeProductId field if non-nil, zero value otherwise.

### GetStripeProductIdOk

`func (o *SubscriptionPlanFull) GetStripeProductIdOk() (*string, bool)`

GetStripeProductIdOk returns a tuple with the StripeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeProductId

`func (o *SubscriptionPlanFull) SetStripeProductId(v string)`

SetStripeProductId sets StripeProductId field to given value.

### HasStripeProductId

`func (o *SubscriptionPlanFull) HasStripeProductId() bool`

HasStripeProductId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *SubscriptionPlanFull) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *SubscriptionPlanFull) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *SubscriptionPlanFull) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *SubscriptionPlanFull) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetFeatureSet

`func (o *SubscriptionPlanFull) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *SubscriptionPlanFull) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *SubscriptionPlanFull) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


