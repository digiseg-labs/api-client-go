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

// checks if the CreateSharedReport201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSharedReport201Response{}

// CreateSharedReport201Response struct for CreateSharedReport201Response
type CreateSharedReport201Response struct {
	Data *SharedReportFull `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSharedReport201Response CreateSharedReport201Response

// NewCreateSharedReport201Response instantiates a new CreateSharedReport201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSharedReport201Response() *CreateSharedReport201Response {
	this := CreateSharedReport201Response{}
	return &this
}

// NewCreateSharedReport201ResponseWithDefaults instantiates a new CreateSharedReport201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSharedReport201ResponseWithDefaults() *CreateSharedReport201Response {
	this := CreateSharedReport201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateSharedReport201Response) GetData() SharedReportFull {
	if o == nil || IsNil(o.Data) {
		var ret SharedReportFull
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSharedReport201Response) GetDataOk() (*SharedReportFull, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateSharedReport201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SharedReportFull and assigns it to the Data field.
func (o *CreateSharedReport201Response) SetData(v SharedReportFull) {
	o.Data = &v
}

func (o CreateSharedReport201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSharedReport201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSharedReport201Response) UnmarshalJSON(data []byte) (err error) {
	varCreateSharedReport201Response := _CreateSharedReport201Response{}

	err = json.Unmarshal(data, &varCreateSharedReport201Response)

	if err != nil {
		return err
	}

	*o = CreateSharedReport201Response(varCreateSharedReport201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSharedReport201Response struct {
	value *CreateSharedReport201Response
	isSet bool
}

func (v NullableCreateSharedReport201Response) Get() *CreateSharedReport201Response {
	return v.value
}

func (v *NullableCreateSharedReport201Response) Set(val *CreateSharedReport201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSharedReport201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSharedReport201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSharedReport201Response(val *CreateSharedReport201Response) *NullableCreateSharedReport201Response {
	return &NullableCreateSharedReport201Response{value: val, isSet: true}
}

func (v NullableCreateSharedReport201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSharedReport201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


