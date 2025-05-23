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

// checks if the AudienceExampleInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceExampleInput{}

// AudienceExampleInput struct for AudienceExampleInput
type AudienceExampleInput struct {
	// The IP address to resolve audiences for
	IpAddress *string `json:"ip_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AudienceExampleInput AudienceExampleInput

// NewAudienceExampleInput instantiates a new AudienceExampleInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceExampleInput() *AudienceExampleInput {
	this := AudienceExampleInput{}
	return &this
}

// NewAudienceExampleInputWithDefaults instantiates a new AudienceExampleInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceExampleInputWithDefaults() *AudienceExampleInput {
	this := AudienceExampleInput{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AudienceExampleInput) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceExampleInput) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AudienceExampleInput) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AudienceExampleInput) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o AudienceExampleInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceExampleInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AudienceExampleInput) UnmarshalJSON(data []byte) (err error) {
	varAudienceExampleInput := _AudienceExampleInput{}

	err = json.Unmarshal(data, &varAudienceExampleInput)

	if err != nil {
		return err
	}

	*o = AudienceExampleInput(varAudienceExampleInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAudienceExampleInput struct {
	value *AudienceExampleInput
	isSet bool
}

func (v NullableAudienceExampleInput) Get() *AudienceExampleInput {
	return v.value
}

func (v *NullableAudienceExampleInput) Set(val *AudienceExampleInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceExampleInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceExampleInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceExampleInput(val *AudienceExampleInput) *NullableAudienceExampleInput {
	return &NullableAudienceExampleInput{value: val, isSet: true}
}

func (v NullableAudienceExampleInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceExampleInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


