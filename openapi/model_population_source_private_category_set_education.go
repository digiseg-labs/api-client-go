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
	"bytes"
	"fmt"
)

// checks if the PopulationSourcePrivateCategorySetEducation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationSourcePrivateCategorySetEducation{}

// PopulationSourcePrivateCategorySetEducation struct for PopulationSourcePrivateCategorySetEducation
type PopulationSourcePrivateCategorySetEducation struct {
	F1 int32 `json:"f1"`
	F2 int32 `json:"f2"`
	F3 int32 `json:"f3"`
}

type _PopulationSourcePrivateCategorySetEducation PopulationSourcePrivateCategorySetEducation

// NewPopulationSourcePrivateCategorySetEducation instantiates a new PopulationSourcePrivateCategorySetEducation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationSourcePrivateCategorySetEducation(f1 int32, f2 int32, f3 int32) *PopulationSourcePrivateCategorySetEducation {
	this := PopulationSourcePrivateCategorySetEducation{}
	this.F1 = f1
	this.F2 = f2
	this.F3 = f3
	return &this
}

// NewPopulationSourcePrivateCategorySetEducationWithDefaults instantiates a new PopulationSourcePrivateCategorySetEducation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationSourcePrivateCategorySetEducationWithDefaults() *PopulationSourcePrivateCategorySetEducation {
	this := PopulationSourcePrivateCategorySetEducation{}
	return &this
}

// GetF1 returns the F1 field value
func (o *PopulationSourcePrivateCategorySetEducation) GetF1() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.F1
}

// GetF1Ok returns a tuple with the F1 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetEducation) GetF1Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.F1, true
}

// SetF1 sets field value
func (o *PopulationSourcePrivateCategorySetEducation) SetF1(v int32) {
	o.F1 = v
}

// GetF2 returns the F2 field value
func (o *PopulationSourcePrivateCategorySetEducation) GetF2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.F2
}

// GetF2Ok returns a tuple with the F2 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetEducation) GetF2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.F2, true
}

// SetF2 sets field value
func (o *PopulationSourcePrivateCategorySetEducation) SetF2(v int32) {
	o.F2 = v
}

// GetF3 returns the F3 field value
func (o *PopulationSourcePrivateCategorySetEducation) GetF3() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.F3
}

// GetF3Ok returns a tuple with the F3 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetEducation) GetF3Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.F3, true
}

// SetF3 sets field value
func (o *PopulationSourcePrivateCategorySetEducation) SetF3(v int32) {
	o.F3 = v
}

func (o PopulationSourcePrivateCategorySetEducation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationSourcePrivateCategorySetEducation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["f1"] = o.F1
	toSerialize["f2"] = o.F2
	toSerialize["f3"] = o.F3
	return toSerialize, nil
}

func (o *PopulationSourcePrivateCategorySetEducation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"f1",
		"f2",
		"f3",
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

	varPopulationSourcePrivateCategorySetEducation := _PopulationSourcePrivateCategorySetEducation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPopulationSourcePrivateCategorySetEducation)

	if err != nil {
		return err
	}

	*o = PopulationSourcePrivateCategorySetEducation(varPopulationSourcePrivateCategorySetEducation)

	return err
}

type NullablePopulationSourcePrivateCategorySetEducation struct {
	value *PopulationSourcePrivateCategorySetEducation
	isSet bool
}

func (v NullablePopulationSourcePrivateCategorySetEducation) Get() *PopulationSourcePrivateCategorySetEducation {
	return v.value
}

func (v *NullablePopulationSourcePrivateCategorySetEducation) Set(val *PopulationSourcePrivateCategorySetEducation) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationSourcePrivateCategorySetEducation) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationSourcePrivateCategorySetEducation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationSourcePrivateCategorySetEducation(val *PopulationSourcePrivateCategorySetEducation) *NullablePopulationSourcePrivateCategorySetEducation {
	return &NullablePopulationSourcePrivateCategorySetEducation{value: val, isSet: true}
}

func (v NullablePopulationSourcePrivateCategorySetEducation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationSourcePrivateCategorySetEducation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


