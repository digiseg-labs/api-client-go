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

// checks if the ComparisonsContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComparisonsContainer{}

// ComparisonsContainer An object containing comparisons
type ComparisonsContainer struct {
	Comparisons []Comparison `json:"comparisons,omitempty"`
}

// NewComparisonsContainer instantiates a new ComparisonsContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComparisonsContainer() *ComparisonsContainer {
	this := ComparisonsContainer{}
	return &this
}

// NewComparisonsContainerWithDefaults instantiates a new ComparisonsContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComparisonsContainerWithDefaults() *ComparisonsContainer {
	this := ComparisonsContainer{}
	return &this
}

// GetComparisons returns the Comparisons field value if set, zero value otherwise.
func (o *ComparisonsContainer) GetComparisons() []Comparison {
	if o == nil || IsNil(o.Comparisons) {
		var ret []Comparison
		return ret
	}
	return o.Comparisons
}

// GetComparisonsOk returns a tuple with the Comparisons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonsContainer) GetComparisonsOk() ([]Comparison, bool) {
	if o == nil || IsNil(o.Comparisons) {
		return nil, false
	}
	return o.Comparisons, true
}

// HasComparisons returns a boolean if a field has been set.
func (o *ComparisonsContainer) HasComparisons() bool {
	if o != nil && !IsNil(o.Comparisons) {
		return true
	}

	return false
}

// SetComparisons gets a reference to the given []Comparison and assigns it to the Comparisons field.
func (o *ComparisonsContainer) SetComparisons(v []Comparison) {
	o.Comparisons = v
}

func (o ComparisonsContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComparisonsContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comparisons) {
		toSerialize["comparisons"] = o.Comparisons
	}
	return toSerialize, nil
}

type NullableComparisonsContainer struct {
	value *ComparisonsContainer
	isSet bool
}

func (v NullableComparisonsContainer) Get() *ComparisonsContainer {
	return v.value
}

func (v *NullableComparisonsContainer) Set(val *ComparisonsContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableComparisonsContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableComparisonsContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparisonsContainer(val *ComparisonsContainer) *NullableComparisonsContainer {
	return &NullableComparisonsContainer{value: val, isSet: true}
}

func (v NullableComparisonsContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComparisonsContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


