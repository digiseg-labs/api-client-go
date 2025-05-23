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
)

// checks if the PopulationAudienceCategorySetSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationAudienceCategorySetSection{}

// PopulationAudienceCategorySetSection struct for PopulationAudienceCategorySetSection
type PopulationAudienceCategorySetSection struct {
	// The fraction of events that fall within this object compared to the total of the category or segment (usually represented by the measurement's parent's parent). For example, if the measurement is \"impression\" on the `home_type` \"Apartment\" object, then the `fraction_of_total` represents the number of impressions on apartments compared to impressions from other `home_type` values. 
	FractionOfTotal *float64 `json:"fraction_of_total,omitempty"`
	// An object with category codes as keys, objects with audience codes and fractions of totals as keys.
	AudienceCategories *map[string]map[string]float64 `json:"audience_categories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopulationAudienceCategorySetSection PopulationAudienceCategorySetSection

// NewPopulationAudienceCategorySetSection instantiates a new PopulationAudienceCategorySetSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationAudienceCategorySetSection() *PopulationAudienceCategorySetSection {
	this := PopulationAudienceCategorySetSection{}
	return &this
}

// NewPopulationAudienceCategorySetSectionWithDefaults instantiates a new PopulationAudienceCategorySetSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationAudienceCategorySetSectionWithDefaults() *PopulationAudienceCategorySetSection {
	this := PopulationAudienceCategorySetSection{}
	return &this
}

// GetFractionOfTotal returns the FractionOfTotal field value if set, zero value otherwise.
func (o *PopulationAudienceCategorySetSection) GetFractionOfTotal() float64 {
	if o == nil || IsNil(o.FractionOfTotal) {
		var ret float64
		return ret
	}
	return *o.FractionOfTotal
}

// GetFractionOfTotalOk returns a tuple with the FractionOfTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationAudienceCategorySetSection) GetFractionOfTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.FractionOfTotal) {
		return nil, false
	}
	return o.FractionOfTotal, true
}

// HasFractionOfTotal returns a boolean if a field has been set.
func (o *PopulationAudienceCategorySetSection) HasFractionOfTotal() bool {
	if o != nil && !IsNil(o.FractionOfTotal) {
		return true
	}

	return false
}

// SetFractionOfTotal gets a reference to the given float64 and assigns it to the FractionOfTotal field.
func (o *PopulationAudienceCategorySetSection) SetFractionOfTotal(v float64) {
	o.FractionOfTotal = &v
}

// GetAudienceCategories returns the AudienceCategories field value if set, zero value otherwise.
func (o *PopulationAudienceCategorySetSection) GetAudienceCategories() map[string]map[string]float64 {
	if o == nil || IsNil(o.AudienceCategories) {
		var ret map[string]map[string]float64
		return ret
	}
	return *o.AudienceCategories
}

// GetAudienceCategoriesOk returns a tuple with the AudienceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationAudienceCategorySetSection) GetAudienceCategoriesOk() (*map[string]map[string]float64, bool) {
	if o == nil || IsNil(o.AudienceCategories) {
		return nil, false
	}
	return o.AudienceCategories, true
}

// HasAudienceCategories returns a boolean if a field has been set.
func (o *PopulationAudienceCategorySetSection) HasAudienceCategories() bool {
	if o != nil && !IsNil(o.AudienceCategories) {
		return true
	}

	return false
}

// SetAudienceCategories gets a reference to the given map[string]map[string]float64 and assigns it to the AudienceCategories field.
func (o *PopulationAudienceCategorySetSection) SetAudienceCategories(v map[string]map[string]float64) {
	o.AudienceCategories = &v
}

func (o PopulationAudienceCategorySetSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationAudienceCategorySetSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FractionOfTotal) {
		toSerialize["fraction_of_total"] = o.FractionOfTotal
	}
	if !IsNil(o.AudienceCategories) {
		toSerialize["audience_categories"] = o.AudienceCategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PopulationAudienceCategorySetSection) UnmarshalJSON(data []byte) (err error) {
	varPopulationAudienceCategorySetSection := _PopulationAudienceCategorySetSection{}

	err = json.Unmarshal(data, &varPopulationAudienceCategorySetSection)

	if err != nil {
		return err
	}

	*o = PopulationAudienceCategorySetSection(varPopulationAudienceCategorySetSection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fraction_of_total")
		delete(additionalProperties, "audience_categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopulationAudienceCategorySetSection struct {
	value *PopulationAudienceCategorySetSection
	isSet bool
}

func (v NullablePopulationAudienceCategorySetSection) Get() *PopulationAudienceCategorySetSection {
	return v.value
}

func (v *NullablePopulationAudienceCategorySetSection) Set(val *PopulationAudienceCategorySetSection) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationAudienceCategorySetSection) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationAudienceCategorySetSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationAudienceCategorySetSection(val *PopulationAudienceCategorySetSection) *NullablePopulationAudienceCategorySetSection {
	return &NullablePopulationAudienceCategorySetSection{value: val, isSet: true}
}

func (v NullablePopulationAudienceCategorySetSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationAudienceCategorySetSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


