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

// checks if the CountryReachStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryReachStats{}

// CountryReachStats Reach statistics for a country
type CountryReachStats struct {
	NumHouseholds *int32 `json:"num_households,omitempty"`
	NumPersons *int32 `json:"num_persons,omitempty"`
	NumDevices *int32 `json:"num_devices,omitempty"`
	NumDailyImpressions *int32 `json:"num_daily_impressions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountryReachStats CountryReachStats

// NewCountryReachStats instantiates a new CountryReachStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryReachStats() *CountryReachStats {
	this := CountryReachStats{}
	return &this
}

// NewCountryReachStatsWithDefaults instantiates a new CountryReachStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryReachStatsWithDefaults() *CountryReachStats {
	this := CountryReachStats{}
	return &this
}

// GetNumHouseholds returns the NumHouseholds field value if set, zero value otherwise.
func (o *CountryReachStats) GetNumHouseholds() int32 {
	if o == nil || IsNil(o.NumHouseholds) {
		var ret int32
		return ret
	}
	return *o.NumHouseholds
}

// GetNumHouseholdsOk returns a tuple with the NumHouseholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryReachStats) GetNumHouseholdsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumHouseholds) {
		return nil, false
	}
	return o.NumHouseholds, true
}

// HasNumHouseholds returns a boolean if a field has been set.
func (o *CountryReachStats) HasNumHouseholds() bool {
	if o != nil && !IsNil(o.NumHouseholds) {
		return true
	}

	return false
}

// SetNumHouseholds gets a reference to the given int32 and assigns it to the NumHouseholds field.
func (o *CountryReachStats) SetNumHouseholds(v int32) {
	o.NumHouseholds = &v
}

// GetNumPersons returns the NumPersons field value if set, zero value otherwise.
func (o *CountryReachStats) GetNumPersons() int32 {
	if o == nil || IsNil(o.NumPersons) {
		var ret int32
		return ret
	}
	return *o.NumPersons
}

// GetNumPersonsOk returns a tuple with the NumPersons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryReachStats) GetNumPersonsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumPersons) {
		return nil, false
	}
	return o.NumPersons, true
}

// HasNumPersons returns a boolean if a field has been set.
func (o *CountryReachStats) HasNumPersons() bool {
	if o != nil && !IsNil(o.NumPersons) {
		return true
	}

	return false
}

// SetNumPersons gets a reference to the given int32 and assigns it to the NumPersons field.
func (o *CountryReachStats) SetNumPersons(v int32) {
	o.NumPersons = &v
}

// GetNumDevices returns the NumDevices field value if set, zero value otherwise.
func (o *CountryReachStats) GetNumDevices() int32 {
	if o == nil || IsNil(o.NumDevices) {
		var ret int32
		return ret
	}
	return *o.NumDevices
}

// GetNumDevicesOk returns a tuple with the NumDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryReachStats) GetNumDevicesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumDevices) {
		return nil, false
	}
	return o.NumDevices, true
}

// HasNumDevices returns a boolean if a field has been set.
func (o *CountryReachStats) HasNumDevices() bool {
	if o != nil && !IsNil(o.NumDevices) {
		return true
	}

	return false
}

// SetNumDevices gets a reference to the given int32 and assigns it to the NumDevices field.
func (o *CountryReachStats) SetNumDevices(v int32) {
	o.NumDevices = &v
}

// GetNumDailyImpressions returns the NumDailyImpressions field value if set, zero value otherwise.
func (o *CountryReachStats) GetNumDailyImpressions() int32 {
	if o == nil || IsNil(o.NumDailyImpressions) {
		var ret int32
		return ret
	}
	return *o.NumDailyImpressions
}

// GetNumDailyImpressionsOk returns a tuple with the NumDailyImpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryReachStats) GetNumDailyImpressionsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumDailyImpressions) {
		return nil, false
	}
	return o.NumDailyImpressions, true
}

// HasNumDailyImpressions returns a boolean if a field has been set.
func (o *CountryReachStats) HasNumDailyImpressions() bool {
	if o != nil && !IsNil(o.NumDailyImpressions) {
		return true
	}

	return false
}

// SetNumDailyImpressions gets a reference to the given int32 and assigns it to the NumDailyImpressions field.
func (o *CountryReachStats) SetNumDailyImpressions(v int32) {
	o.NumDailyImpressions = &v
}

func (o CountryReachStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryReachStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumHouseholds) {
		toSerialize["num_households"] = o.NumHouseholds
	}
	if !IsNil(o.NumPersons) {
		toSerialize["num_persons"] = o.NumPersons
	}
	if !IsNil(o.NumDevices) {
		toSerialize["num_devices"] = o.NumDevices
	}
	if !IsNil(o.NumDailyImpressions) {
		toSerialize["num_daily_impressions"] = o.NumDailyImpressions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountryReachStats) UnmarshalJSON(data []byte) (err error) {
	varCountryReachStats := _CountryReachStats{}

	err = json.Unmarshal(data, &varCountryReachStats)

	if err != nil {
		return err
	}

	*o = CountryReachStats(varCountryReachStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "num_households")
		delete(additionalProperties, "num_persons")
		delete(additionalProperties, "num_devices")
		delete(additionalProperties, "num_daily_impressions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountryReachStats struct {
	value *CountryReachStats
	isSet bool
}

func (v NullableCountryReachStats) Get() *CountryReachStats {
	return v.value
}

func (v *NullableCountryReachStats) Set(val *CountryReachStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryReachStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryReachStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryReachStats(val *CountryReachStats) *NullableCountryReachStats {
	return &NullableCountryReachStats{value: val, isSet: true}
}

func (v NullableCountryReachStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryReachStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


