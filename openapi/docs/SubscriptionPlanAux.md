# SubscriptionPlanAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the plan/price that the account is subscribed to | [optional] [readonly] 
**FeatureSet** | [**PlanFeatureSet**](PlanFeatureSet.md) |  | 

## Methods

### NewSubscriptionPlanAux

`func NewSubscriptionPlanAux(featureSet PlanFeatureSet, ) *SubscriptionPlanAux`

NewSubscriptionPlanAux instantiates a new SubscriptionPlanAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanAuxWithDefaults

`func NewSubscriptionPlanAuxWithDefaults() *SubscriptionPlanAux`

NewSubscriptionPlanAuxWithDefaults instantiates a new SubscriptionPlanAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPlanAux) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanAux) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanAux) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPlanAux) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeatureSet

`func (o *SubscriptionPlanAux) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *SubscriptionPlanAux) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *SubscriptionPlanAux) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


