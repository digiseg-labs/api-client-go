/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BusinessAudienceStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessAudienceStats{}

// BusinessAudienceStats Measurements (overall and per audience) for traffic resolved as business users
type BusinessAudienceStats struct {
	// Measurements related to this object
	Measurements []Measurement `json:"measurements,omitempty"`
	AudienceCategories *BusinessAudienceStatsAllOfAudienceCategories `json:"audience_categories,omitempty"`
}

// NewBusinessAudienceStats instantiates a new BusinessAudienceStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessAudienceStats() *BusinessAudienceStats {
	this := BusinessAudienceStats{}
	return &this
}

// NewBusinessAudienceStatsWithDefaults instantiates a new BusinessAudienceStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessAudienceStatsWithDefaults() *BusinessAudienceStats {
	this := BusinessAudienceStats{}
	return &this
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *BusinessAudienceStats) GetMeasurements() []Measurement {
	if o == nil || IsNil(o.Measurements) {
		var ret []Measurement
		return ret
	}
	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessAudienceStats) GetMeasurementsOk() ([]Measurement, bool) {
	if o == nil || IsNil(o.Measurements) {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *BusinessAudienceStats) HasMeasurements() bool {
	if o != nil && !IsNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given []Measurement and assigns it to the Measurements field.
func (o *BusinessAudienceStats) SetMeasurements(v []Measurement) {
	o.Measurements = v
}

// GetAudienceCategories returns the AudienceCategories field value if set, zero value otherwise.
func (o *BusinessAudienceStats) GetAudienceCategories() BusinessAudienceStatsAllOfAudienceCategories {
	if o == nil || IsNil(o.AudienceCategories) {
		var ret BusinessAudienceStatsAllOfAudienceCategories
		return ret
	}
	return *o.AudienceCategories
}

// GetAudienceCategoriesOk returns a tuple with the AudienceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessAudienceStats) GetAudienceCategoriesOk() (*BusinessAudienceStatsAllOfAudienceCategories, bool) {
	if o == nil || IsNil(o.AudienceCategories) {
		return nil, false
	}
	return o.AudienceCategories, true
}

// HasAudienceCategories returns a boolean if a field has been set.
func (o *BusinessAudienceStats) HasAudienceCategories() bool {
	if o != nil && !IsNil(o.AudienceCategories) {
		return true
	}

	return false
}

// SetAudienceCategories gets a reference to the given BusinessAudienceStatsAllOfAudienceCategories and assigns it to the AudienceCategories field.
func (o *BusinessAudienceStats) SetAudienceCategories(v BusinessAudienceStatsAllOfAudienceCategories) {
	o.AudienceCategories = &v
}

func (o BusinessAudienceStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessAudienceStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Measurements) {
		toSerialize["measurements"] = o.Measurements
	}
	if !IsNil(o.AudienceCategories) {
		toSerialize["audience_categories"] = o.AudienceCategories
	}
	return toSerialize, nil
}

type NullableBusinessAudienceStats struct {
	value *BusinessAudienceStats
	isSet bool
}

func (v NullableBusinessAudienceStats) Get() *BusinessAudienceStats {
	return v.value
}

func (v *NullableBusinessAudienceStats) Set(val *BusinessAudienceStats) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessAudienceStats) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessAudienceStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessAudienceStats(val *BusinessAudienceStats) *NullableBusinessAudienceStats {
	return &NullableBusinessAudienceStats{value: val, isSet: true}
}

func (v NullableBusinessAudienceStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessAudienceStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

