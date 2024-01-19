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

// checks if the AudienceCategoryStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceCategoryStats{}

// AudienceCategoryStats struct for AudienceCategoryStats
type AudienceCategoryStats struct {
	Name *string `json:"name,omitempty"`
	Audiences []AudienceStats `json:"audiences,omitempty"`
}

// NewAudienceCategoryStats instantiates a new AudienceCategoryStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceCategoryStats() *AudienceCategoryStats {
	this := AudienceCategoryStats{}
	return &this
}

// NewAudienceCategoryStatsWithDefaults instantiates a new AudienceCategoryStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceCategoryStatsWithDefaults() *AudienceCategoryStats {
	this := AudienceCategoryStats{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AudienceCategoryStats) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceCategoryStats) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AudienceCategoryStats) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AudienceCategoryStats) SetName(v string) {
	o.Name = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *AudienceCategoryStats) GetAudiences() []AudienceStats {
	if o == nil || IsNil(o.Audiences) {
		var ret []AudienceStats
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceCategoryStats) GetAudiencesOk() ([]AudienceStats, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *AudienceCategoryStats) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []AudienceStats and assigns it to the Audiences field.
func (o *AudienceCategoryStats) SetAudiences(v []AudienceStats) {
	o.Audiences = v
}

func (o AudienceCategoryStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceCategoryStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	return toSerialize, nil
}

type NullableAudienceCategoryStats struct {
	value *AudienceCategoryStats
	isSet bool
}

func (v NullableAudienceCategoryStats) Get() *AudienceCategoryStats {
	return v.value
}

func (v *NullableAudienceCategoryStats) Set(val *AudienceCategoryStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceCategoryStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceCategoryStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceCategoryStats(val *AudienceCategoryStats) *NullableAudienceCategoryStats {
	return &NullableAudienceCategoryStats{value: val, isSet: true}
}

func (v NullableAudienceCategoryStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceCategoryStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

