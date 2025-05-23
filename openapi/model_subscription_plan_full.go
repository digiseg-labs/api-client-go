/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" /> 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SubscriptionPlanFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlanFull{}

// SubscriptionPlanFull struct for SubscriptionPlanFull
type SubscriptionPlanFull struct {
	// The ID of the plan/price that the account is subscribed to
	Id *string `json:"id,omitempty"`
	// Date and time of the object creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ID of the user who created the object
	CreatedBy *string `json:"created_by,omitempty"`
	// Date and time of the latest update to the object
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the user who last updated the object
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The display name of the price/plan
	DisplayName *string `json:"display_name,omitempty"`
	// An optional code, typically provided if the plan/price is public and advertised
	Code *string `json:"code,omitempty"`
	// Is the plan/price a public price or custom?
	IsPublic bool `json:"is_public"`
	ProductType SubscriptionProductType `json:"product_type"`
	ListPrice *SubscriptionPrice `json:"list_price,omitempty"`
	StripeProductId *string `json:"stripe_product_id,omitempty"`
	StripePriceId *string `json:"stripe_price_id,omitempty"`
	FeatureSet PlanFeatureSet `json:"feature_set"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionPlanFull SubscriptionPlanFull

// NewSubscriptionPlanFull instantiates a new SubscriptionPlanFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlanFull(isPublic bool, productType SubscriptionProductType, featureSet PlanFeatureSet) *SubscriptionPlanFull {
	this := SubscriptionPlanFull{}
	this.IsPublic = isPublic
	this.ProductType = productType
	this.FeatureSet = featureSet
	return &this
}

// NewSubscriptionPlanFullWithDefaults instantiates a new SubscriptionPlanFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanFullWithDefaults() *SubscriptionPlanFull {
	this := SubscriptionPlanFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionPlanFull) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SubscriptionPlanFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SubscriptionPlanFull) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SubscriptionPlanFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SubscriptionPlanFull) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionPlanFull) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SubscriptionPlanFull) SetCode(v string) {
	o.Code = &v
}

// GetIsPublic returns the IsPublic field value
func (o *SubscriptionPlanFull) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *SubscriptionPlanFull) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetProductType returns the ProductType field value
func (o *SubscriptionPlanFull) GetProductType() SubscriptionProductType {
	if o == nil {
		var ret SubscriptionProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetProductTypeOk() (*SubscriptionProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *SubscriptionPlanFull) SetProductType(v SubscriptionProductType) {
	o.ProductType = v
}

// GetListPrice returns the ListPrice field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetListPrice() SubscriptionPrice {
	if o == nil || IsNil(o.ListPrice) {
		var ret SubscriptionPrice
		return ret
	}
	return *o.ListPrice
}

// GetListPriceOk returns a tuple with the ListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetListPriceOk() (*SubscriptionPrice, bool) {
	if o == nil || IsNil(o.ListPrice) {
		return nil, false
	}
	return o.ListPrice, true
}

// HasListPrice returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasListPrice() bool {
	if o != nil && !IsNil(o.ListPrice) {
		return true
	}

	return false
}

// SetListPrice gets a reference to the given SubscriptionPrice and assigns it to the ListPrice field.
func (o *SubscriptionPlanFull) SetListPrice(v SubscriptionPrice) {
	o.ListPrice = &v
}

// GetStripeProductId returns the StripeProductId field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetStripeProductId() string {
	if o == nil || IsNil(o.StripeProductId) {
		var ret string
		return ret
	}
	return *o.StripeProductId
}

// GetStripeProductIdOk returns a tuple with the StripeProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetStripeProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeProductId) {
		return nil, false
	}
	return o.StripeProductId, true
}

// HasStripeProductId returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasStripeProductId() bool {
	if o != nil && !IsNil(o.StripeProductId) {
		return true
	}

	return false
}

// SetStripeProductId gets a reference to the given string and assigns it to the StripeProductId field.
func (o *SubscriptionPlanFull) SetStripeProductId(v string) {
	o.StripeProductId = &v
}

// GetStripePriceId returns the StripePriceId field value if set, zero value otherwise.
func (o *SubscriptionPlanFull) GetStripePriceId() string {
	if o == nil || IsNil(o.StripePriceId) {
		var ret string
		return ret
	}
	return *o.StripePriceId
}

// GetStripePriceIdOk returns a tuple with the StripePriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetStripePriceIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripePriceId) {
		return nil, false
	}
	return o.StripePriceId, true
}

// HasStripePriceId returns a boolean if a field has been set.
func (o *SubscriptionPlanFull) HasStripePriceId() bool {
	if o != nil && !IsNil(o.StripePriceId) {
		return true
	}

	return false
}

// SetStripePriceId gets a reference to the given string and assigns it to the StripePriceId field.
func (o *SubscriptionPlanFull) SetStripePriceId(v string) {
	o.StripePriceId = &v
}

// GetFeatureSet returns the FeatureSet field value
func (o *SubscriptionPlanFull) GetFeatureSet() PlanFeatureSet {
	if o == nil {
		var ret PlanFeatureSet
		return ret
	}

	return o.FeatureSet
}

// GetFeatureSetOk returns a tuple with the FeatureSet field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanFull) GetFeatureSetOk() (*PlanFeatureSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureSet, true
}

// SetFeatureSet sets field value
func (o *SubscriptionPlanFull) SetFeatureSet(v PlanFeatureSet) {
	o.FeatureSet = v
}

func (o SubscriptionPlanFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlanFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["is_public"] = o.IsPublic
	toSerialize["product_type"] = o.ProductType
	if !IsNil(o.ListPrice) {
		toSerialize["list_price"] = o.ListPrice
	}
	if !IsNil(o.StripeProductId) {
		toSerialize["stripe_product_id"] = o.StripeProductId
	}
	if !IsNil(o.StripePriceId) {
		toSerialize["stripe_price_id"] = o.StripePriceId
	}
	toSerialize["feature_set"] = o.FeatureSet

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionPlanFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_public",
		"product_type",
		"feature_set",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscriptionPlanFull := _SubscriptionPlanFull{}

	err = json.Unmarshal(data, &varSubscriptionPlanFull)

	if err != nil {
		return err
	}

	*o = SubscriptionPlanFull(varSubscriptionPlanFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "code")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "list_price")
		delete(additionalProperties, "stripe_product_id")
		delete(additionalProperties, "stripe_price_id")
		delete(additionalProperties, "feature_set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionPlanFull struct {
	value *SubscriptionPlanFull
	isSet bool
}

func (v NullableSubscriptionPlanFull) Get() *SubscriptionPlanFull {
	return v.value
}

func (v *NullableSubscriptionPlanFull) Set(val *SubscriptionPlanFull) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlanFull) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlanFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlanFull(val *SubscriptionPlanFull) *NullableSubscriptionPlanFull {
	return &NullableSubscriptionPlanFull{value: val, isSet: true}
}

func (v NullableSubscriptionPlanFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlanFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


