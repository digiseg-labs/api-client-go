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

// checks if the IdentifyableObject1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentifyableObject1{}

// IdentifyableObject1 struct for IdentifyableObject1
type IdentifyableObject1 struct {
	// Unique ID for the object
	Id string `json:"id"`
}

type _IdentifyableObject1 IdentifyableObject1

// NewIdentifyableObject1 instantiates a new IdentifyableObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifyableObject1(id string) *IdentifyableObject1 {
	this := IdentifyableObject1{}
	this.Id = id
	return &this
}

// NewIdentifyableObject1WithDefaults instantiates a new IdentifyableObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifyableObject1WithDefaults() *IdentifyableObject1 {
	this := IdentifyableObject1{}
	return &this
}

// GetId returns the Id field value
func (o *IdentifyableObject1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentifyableObject1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentifyableObject1) SetId(v string) {
	o.Id = v
}

func (o IdentifyableObject1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentifyableObject1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *IdentifyableObject1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varIdentifyableObject1 := _IdentifyableObject1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentifyableObject1)

	if err != nil {
		return err
	}

	*o = IdentifyableObject1(varIdentifyableObject1)

	return err
}

type NullableIdentifyableObject1 struct {
	value *IdentifyableObject1
	isSet bool
}

func (v NullableIdentifyableObject1) Get() *IdentifyableObject1 {
	return v.value
}

func (v *NullableIdentifyableObject1) Set(val *IdentifyableObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifyableObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifyableObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifyableObject1(val *IdentifyableObject1) *NullableIdentifyableObject1 {
	return &NullableIdentifyableObject1{value: val, isSet: true}
}

func (v NullableIdentifyableObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifyableObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


