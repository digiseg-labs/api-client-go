# SubscriptionPlanMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The display name of the price/plan | [optional] 
**Code** | Pointer to **string** | An optional code, typically provided if the plan/price is public and advertised | [optional] 
**IsPublic** | Pointer to **bool** | Is the plan/price a public price or custom? | [optional] 

## Methods

### NewSubscriptionPlanMutation

`func NewSubscriptionPlanMutation() *SubscriptionPlanMutation`

NewSubscriptionPlanMutation instantiates a new SubscriptionPlanMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanMutationWithDefaults

`func NewSubscriptionPlanMutationWithDefaults() *SubscriptionPlanMutation`

NewSubscriptionPlanMutationWithDefaults instantiates a new SubscriptionPlanMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureSet

`func (o *SubscriptionPlanMutation) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *SubscriptionPlanMutation) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *SubscriptionPlanMutation) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *SubscriptionPlanMutation) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionPlanMutation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanMutation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanMutation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanMutation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionPlanMutation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanMutation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanMutation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionPlanMutation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsPublic

`func (o *SubscriptionPlanMutation) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubscriptionPlanMutation) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubscriptionPlanMutation) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *SubscriptionPlanMutation) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


