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
	"fmt"
)

// checks if the StudyDeviceTypeCategoryStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyDeviceTypeCategoryStats{}

// StudyDeviceTypeCategoryStats struct for StudyDeviceTypeCategoryStats
type StudyDeviceTypeCategoryStats struct {
	Category StudyDeviceTypeCategory `json:"category"`
	// The total count of events recorded for the device category.
	Count int32 `json:"count"`
	// The total count of impressions recorded for the device category.
	Impressions *int32 `json:"impressions,omitempty"`
	// The total count of clicks recorded for the device category.
	Clicks *int32 `json:"clicks,omitempty"`
	SubTypes []StudyDeviceSubTypeStats `json:"sub_types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyDeviceTypeCategoryStats StudyDeviceTypeCategoryStats

// NewStudyDeviceTypeCategoryStats instantiates a new StudyDeviceTypeCategoryStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyDeviceTypeCategoryStats(category StudyDeviceTypeCategory, count int32) *StudyDeviceTypeCategoryStats {
	this := StudyDeviceTypeCategoryStats{}
	this.Category = category
	this.Count = count
	return &this
}

// NewStudyDeviceTypeCategoryStatsWithDefaults instantiates a new StudyDeviceTypeCategoryStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyDeviceTypeCategoryStatsWithDefaults() *StudyDeviceTypeCategoryStats {
	this := StudyDeviceTypeCategoryStats{}
	return &this
}

// GetCategory returns the Category field value
func (o *StudyDeviceTypeCategoryStats) GetCategory() StudyDeviceTypeCategory {
	if o == nil {
		var ret StudyDeviceTypeCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *StudyDeviceTypeCategoryStats) GetCategoryOk() (*StudyDeviceTypeCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *StudyDeviceTypeCategoryStats) SetCategory(v StudyDeviceTypeCategory) {
	o.Category = v
}

// GetCount returns the Count field value
func (o *StudyDeviceTypeCategoryStats) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *StudyDeviceTypeCategoryStats) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *StudyDeviceTypeCategoryStats) SetCount(v int32) {
	o.Count = v
}

// GetImpressions returns the Impressions field value if set, zero value otherwise.
func (o *StudyDeviceTypeCategoryStats) GetImpressions() int32 {
	if o == nil || IsNil(o.Impressions) {
		var ret int32
		return ret
	}
	return *o.Impressions
}

// GetImpressionsOk returns a tuple with the Impressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyDeviceTypeCategoryStats) GetImpressionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Impressions) {
		return nil, false
	}
	return o.Impressions, true
}

// HasImpressions returns a boolean if a field has been set.
func (o *StudyDeviceTypeCategoryStats) HasImpressions() bool {
	if o != nil && !IsNil(o.Impressions) {
		return true
	}

	return false
}

// SetImpressions gets a reference to the given int32 and assigns it to the Impressions field.
func (o *StudyDeviceTypeCategoryStats) SetImpressions(v int32) {
	o.Impressions = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *StudyDeviceTypeCategoryStats) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyDeviceTypeCategoryStats) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *StudyDeviceTypeCategoryStats) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *StudyDeviceTypeCategoryStats) SetClicks(v int32) {
	o.Clicks = &v
}

// GetSubTypes returns the SubTypes field value if set, zero value otherwise.
func (o *StudyDeviceTypeCategoryStats) GetSubTypes() []StudyDeviceSubTypeStats {
	if o == nil || IsNil(o.SubTypes) {
		var ret []StudyDeviceSubTypeStats
		return ret
	}
	return o.SubTypes
}

// GetSubTypesOk returns a tuple with the SubTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyDeviceTypeCategoryStats) GetSubTypesOk() ([]StudyDeviceSubTypeStats, bool) {
	if o == nil || IsNil(o.SubTypes) {
		return nil, false
	}
	return o.SubTypes, true
}

// HasSubTypes returns a boolean if a field has been set.
func (o *StudyDeviceTypeCategoryStats) HasSubTypes() bool {
	if o != nil && !IsNil(o.SubTypes) {
		return true
	}

	return false
}

// SetSubTypes gets a reference to the given []StudyDeviceSubTypeStats and assigns it to the SubTypes field.
func (o *StudyDeviceTypeCategoryStats) SetSubTypes(v []StudyDeviceSubTypeStats) {
	o.SubTypes = v
}

func (o StudyDeviceTypeCategoryStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyDeviceTypeCategoryStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["count"] = o.Count
	if !IsNil(o.Impressions) {
		toSerialize["impressions"] = o.Impressions
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.SubTypes) {
		toSerialize["sub_types"] = o.SubTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyDeviceTypeCategoryStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
		"count",
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

	varStudyDeviceTypeCategoryStats := _StudyDeviceTypeCategoryStats{}

	err = json.Unmarshal(data, &varStudyDeviceTypeCategoryStats)

	if err != nil {
		return err
	}

	*o = StudyDeviceTypeCategoryStats(varStudyDeviceTypeCategoryStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "category")
		delete(additionalProperties, "count")
		delete(additionalProperties, "impressions")
		delete(additionalProperties, "clicks")
		delete(additionalProperties, "sub_types")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyDeviceTypeCategoryStats struct {
	value *StudyDeviceTypeCategoryStats
	isSet bool
}

func (v NullableStudyDeviceTypeCategoryStats) Get() *StudyDeviceTypeCategoryStats {
	return v.value
}

func (v *NullableStudyDeviceTypeCategoryStats) Set(val *StudyDeviceTypeCategoryStats) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyDeviceTypeCategoryStats) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyDeviceTypeCategoryStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyDeviceTypeCategoryStats(val *StudyDeviceTypeCategoryStats) *NullableStudyDeviceTypeCategoryStats {
	return &NullableStudyDeviceTypeCategoryStats{value: val, isSet: true}
}

func (v NullableStudyDeviceTypeCategoryStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyDeviceTypeCategoryStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


