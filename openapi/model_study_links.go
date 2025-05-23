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

// checks if the StudyLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyLinks{}

// StudyLinks struct for StudyLinks
type StudyLinks struct {
	// Link to the country statistics for the study
	StatsCountries *string `json:"stats_countries,omitempty"`
	// Link to the audience statistics for the study
	StatsAudiences *string `json:"stats_audiences,omitempty"`
	// Link to the timing statistics for the study
	StatsTiming *string `json:"stats_timing,omitempty"`
	// Link to the frequency statistics for the study
	StatsFrequencies *string `json:"stats_frequencies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyLinks StudyLinks

// NewStudyLinks instantiates a new StudyLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyLinks() *StudyLinks {
	this := StudyLinks{}
	return &this
}

// NewStudyLinksWithDefaults instantiates a new StudyLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyLinksWithDefaults() *StudyLinks {
	this := StudyLinks{}
	return &this
}

// GetStatsCountries returns the StatsCountries field value if set, zero value otherwise.
func (o *StudyLinks) GetStatsCountries() string {
	if o == nil || IsNil(o.StatsCountries) {
		var ret string
		return ret
	}
	return *o.StatsCountries
}

// GetStatsCountriesOk returns a tuple with the StatsCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyLinks) GetStatsCountriesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsCountries) {
		return nil, false
	}
	return o.StatsCountries, true
}

// HasStatsCountries returns a boolean if a field has been set.
func (o *StudyLinks) HasStatsCountries() bool {
	if o != nil && !IsNil(o.StatsCountries) {
		return true
	}

	return false
}

// SetStatsCountries gets a reference to the given string and assigns it to the StatsCountries field.
func (o *StudyLinks) SetStatsCountries(v string) {
	o.StatsCountries = &v
}

// GetStatsAudiences returns the StatsAudiences field value if set, zero value otherwise.
func (o *StudyLinks) GetStatsAudiences() string {
	if o == nil || IsNil(o.StatsAudiences) {
		var ret string
		return ret
	}
	return *o.StatsAudiences
}

// GetStatsAudiencesOk returns a tuple with the StatsAudiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyLinks) GetStatsAudiencesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsAudiences) {
		return nil, false
	}
	return o.StatsAudiences, true
}

// HasStatsAudiences returns a boolean if a field has been set.
func (o *StudyLinks) HasStatsAudiences() bool {
	if o != nil && !IsNil(o.StatsAudiences) {
		return true
	}

	return false
}

// SetStatsAudiences gets a reference to the given string and assigns it to the StatsAudiences field.
func (o *StudyLinks) SetStatsAudiences(v string) {
	o.StatsAudiences = &v
}

// GetStatsTiming returns the StatsTiming field value if set, zero value otherwise.
func (o *StudyLinks) GetStatsTiming() string {
	if o == nil || IsNil(o.StatsTiming) {
		var ret string
		return ret
	}
	return *o.StatsTiming
}

// GetStatsTimingOk returns a tuple with the StatsTiming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyLinks) GetStatsTimingOk() (*string, bool) {
	if o == nil || IsNil(o.StatsTiming) {
		return nil, false
	}
	return o.StatsTiming, true
}

// HasStatsTiming returns a boolean if a field has been set.
func (o *StudyLinks) HasStatsTiming() bool {
	if o != nil && !IsNil(o.StatsTiming) {
		return true
	}

	return false
}

// SetStatsTiming gets a reference to the given string and assigns it to the StatsTiming field.
func (o *StudyLinks) SetStatsTiming(v string) {
	o.StatsTiming = &v
}

// GetStatsFrequencies returns the StatsFrequencies field value if set, zero value otherwise.
func (o *StudyLinks) GetStatsFrequencies() string {
	if o == nil || IsNil(o.StatsFrequencies) {
		var ret string
		return ret
	}
	return *o.StatsFrequencies
}

// GetStatsFrequenciesOk returns a tuple with the StatsFrequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyLinks) GetStatsFrequenciesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsFrequencies) {
		return nil, false
	}
	return o.StatsFrequencies, true
}

// HasStatsFrequencies returns a boolean if a field has been set.
func (o *StudyLinks) HasStatsFrequencies() bool {
	if o != nil && !IsNil(o.StatsFrequencies) {
		return true
	}

	return false
}

// SetStatsFrequencies gets a reference to the given string and assigns it to the StatsFrequencies field.
func (o *StudyLinks) SetStatsFrequencies(v string) {
	o.StatsFrequencies = &v
}

func (o StudyLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatsCountries) {
		toSerialize["stats_countries"] = o.StatsCountries
	}
	if !IsNil(o.StatsAudiences) {
		toSerialize["stats_audiences"] = o.StatsAudiences
	}
	if !IsNil(o.StatsTiming) {
		toSerialize["stats_timing"] = o.StatsTiming
	}
	if !IsNil(o.StatsFrequencies) {
		toSerialize["stats_frequencies"] = o.StatsFrequencies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyLinks) UnmarshalJSON(data []byte) (err error) {
	varStudyLinks := _StudyLinks{}

	err = json.Unmarshal(data, &varStudyLinks)

	if err != nil {
		return err
	}

	*o = StudyLinks(varStudyLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stats_countries")
		delete(additionalProperties, "stats_audiences")
		delete(additionalProperties, "stats_timing")
		delete(additionalProperties, "stats_frequencies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyLinks struct {
	value *StudyLinks
	isSet bool
}

func (v NullableStudyLinks) Get() *StudyLinks {
	return v.value
}

func (v *NullableStudyLinks) Set(val *StudyLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyLinks(val *StudyLinks) *NullableStudyLinks {
	return &NullableStudyLinks{value: val, isSet: true}
}

func (v NullableStudyLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


