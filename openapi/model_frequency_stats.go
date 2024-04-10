/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young singles and couples | |  |  | c2 | Young couples with children | |  |  | c3 | Families with school children | |  |  | c4 | Older families | |  |  | c5 | Pensioners | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Up to 80 m² | |  |  | l2 | 80-119 m² | |  |  | l3 | Above 120 m² | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FrequencyStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrequencyStats{}

// FrequencyStats struct for FrequencyStats
type FrequencyStats struct {
	// A frequency of study impressions to users. * The value 1 means that the study has been exposed just once. * The value 2 means that the study has been exposed twice. * And so on... 
	Frequency *int32 `json:"frequency,omitempty"`
	// The number of users that have generated impressions at the corresponding frequency
	Count *int32 `json:"count,omitempty"`
}

// NewFrequencyStats instantiates a new FrequencyStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrequencyStats() *FrequencyStats {
	this := FrequencyStats{}
	return &this
}

// NewFrequencyStatsWithDefaults instantiates a new FrequencyStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrequencyStatsWithDefaults() *FrequencyStats {
	this := FrequencyStats{}
	return &this
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *FrequencyStats) GetFrequency() int32 {
	if o == nil || IsNil(o.Frequency) {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyStats) GetFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *FrequencyStats) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *FrequencyStats) SetFrequency(v int32) {
	o.Frequency = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FrequencyStats) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyStats) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FrequencyStats) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FrequencyStats) SetCount(v int32) {
	o.Count = &v
}

func (o FrequencyStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrequencyStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableFrequencyStats struct {
	value *FrequencyStats
	isSet bool
}

func (v NullableFrequencyStats) Get() *FrequencyStats {
	return v.value
}

func (v *NullableFrequencyStats) Set(val *FrequencyStats) {
	v.value = val
	v.isSet = true
}

func (v NullableFrequencyStats) IsSet() bool {
	return v.isSet
}

func (v *NullableFrequencyStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrequencyStats(val *FrequencyStats) *NullableFrequencyStats {
	return &NullableFrequencyStats{value: val, isSet: true}
}

func (v NullableFrequencyStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrequencyStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


