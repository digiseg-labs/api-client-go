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

// checks if the StudyMutationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyMutationData{}

// StudyMutationData struct for StudyMutationData
type StudyMutationData struct {
	EventSet *MeasurementEventSet `json:"event_set,omitempty"`
	// The ID of the measurement client that this study is for
	ClientId *string `json:"client_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyMutationData StudyMutationData

// NewStudyMutationData instantiates a new StudyMutationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyMutationData() *StudyMutationData {
	this := StudyMutationData{}
	return &this
}

// NewStudyMutationDataWithDefaults instantiates a new StudyMutationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyMutationDataWithDefaults() *StudyMutationData {
	this := StudyMutationData{}
	return &this
}

// GetEventSet returns the EventSet field value if set, zero value otherwise.
func (o *StudyMutationData) GetEventSet() MeasurementEventSet {
	if o == nil || IsNil(o.EventSet) {
		var ret MeasurementEventSet
		return ret
	}
	return *o.EventSet
}

// GetEventSetOk returns a tuple with the EventSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyMutationData) GetEventSetOk() (*MeasurementEventSet, bool) {
	if o == nil || IsNil(o.EventSet) {
		return nil, false
	}
	return o.EventSet, true
}

// HasEventSet returns a boolean if a field has been set.
func (o *StudyMutationData) HasEventSet() bool {
	if o != nil && !IsNil(o.EventSet) {
		return true
	}

	return false
}

// SetEventSet gets a reference to the given MeasurementEventSet and assigns it to the EventSet field.
func (o *StudyMutationData) SetEventSet(v MeasurementEventSet) {
	o.EventSet = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *StudyMutationData) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyMutationData) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *StudyMutationData) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *StudyMutationData) SetClientId(v string) {
	o.ClientId = &v
}

func (o StudyMutationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyMutationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventSet) {
		toSerialize["event_set"] = o.EventSet
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyMutationData) UnmarshalJSON(data []byte) (err error) {
	varStudyMutationData := _StudyMutationData{}

	err = json.Unmarshal(data, &varStudyMutationData)

	if err != nil {
		return err
	}

	*o = StudyMutationData(varStudyMutationData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event_set")
		delete(additionalProperties, "client_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyMutationData struct {
	value *StudyMutationData
	isSet bool
}

func (v NullableStudyMutationData) Get() *StudyMutationData {
	return v.value
}

func (v *NullableStudyMutationData) Set(val *StudyMutationData) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyMutationData) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyMutationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyMutationData(val *StudyMutationData) *NullableStudyMutationData {
	return &NullableStudyMutationData{value: val, isSet: true}
}

func (v NullableStudyMutationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyMutationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


