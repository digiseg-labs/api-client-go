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
	"fmt"
)

// checks if the Comparison type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Comparison{}

// Comparison Represents a comparison of measurements with another source
type Comparison struct {
	// The name of the comparison source
	Name string `json:"name"`
	// The code of the comparison source
	Code string `json:"code"`
	// The \"fraction of total\" value that is being compared with. 
	FractionOfTotal float64 `json:"fraction_of_total"`
	// The comparison index where 100 means that the measurement is completely aligned with the compared metric. Values below 100 means that the measurement is below the compared metric, and values above 100 means that the measurement is above the compared metric. 
	Index int32 `json:"index"`
}

type _Comparison Comparison

// NewComparison instantiates a new Comparison object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComparison(name string, code string, fractionOfTotal float64, index int32) *Comparison {
	this := Comparison{}
	this.Name = name
	this.Code = code
	this.FractionOfTotal = fractionOfTotal
	this.Index = index
	return &this
}

// NewComparisonWithDefaults instantiates a new Comparison object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComparisonWithDefaults() *Comparison {
	this := Comparison{}
	return &this
}

// GetName returns the Name field value
func (o *Comparison) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Comparison) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Comparison) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *Comparison) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Comparison) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Comparison) SetCode(v string) {
	o.Code = v
}

// GetFractionOfTotal returns the FractionOfTotal field value
func (o *Comparison) GetFractionOfTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FractionOfTotal
}

// GetFractionOfTotalOk returns a tuple with the FractionOfTotal field value
// and a boolean to check if the value has been set.
func (o *Comparison) GetFractionOfTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FractionOfTotal, true
}

// SetFractionOfTotal sets field value
func (o *Comparison) SetFractionOfTotal(v float64) {
	o.FractionOfTotal = v
}

// GetIndex returns the Index field value
func (o *Comparison) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *Comparison) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *Comparison) SetIndex(v int32) {
	o.Index = v
}

func (o Comparison) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Comparison) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	toSerialize["fraction_of_total"] = o.FractionOfTotal
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

func (o *Comparison) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"code",
		"fraction_of_total",
		"index",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varComparison := _Comparison{}

	err = json.Unmarshal(bytes, &varComparison)

	if err != nil {
		return err
	}

	*o = Comparison(varComparison)

	return err
}

type NullableComparison struct {
	value *Comparison
	isSet bool
}

func (v NullableComparison) Get() *Comparison {
	return v.value
}

func (v *NullableComparison) Set(val *Comparison) {
	v.value = val
	v.isSet = true
}

func (v NullableComparison) IsSet() bool {
	return v.isSet
}

func (v *NullableComparison) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparison(val *Comparison) *NullableComparison {
	return &NullableComparison{value: val, isSet: true}
}

func (v NullableComparison) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComparison) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

