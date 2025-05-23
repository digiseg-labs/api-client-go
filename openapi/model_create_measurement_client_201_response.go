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

// checks if the CreateMeasurementClient201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMeasurementClient201Response{}

// CreateMeasurementClient201Response struct for CreateMeasurementClient201Response
type CreateMeasurementClient201Response struct {
	Data *MeasurementClientFull `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMeasurementClient201Response CreateMeasurementClient201Response

// NewCreateMeasurementClient201Response instantiates a new CreateMeasurementClient201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMeasurementClient201Response() *CreateMeasurementClient201Response {
	this := CreateMeasurementClient201Response{}
	return &this
}

// NewCreateMeasurementClient201ResponseWithDefaults instantiates a new CreateMeasurementClient201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMeasurementClient201ResponseWithDefaults() *CreateMeasurementClient201Response {
	this := CreateMeasurementClient201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateMeasurementClient201Response) GetData() MeasurementClientFull {
	if o == nil || IsNil(o.Data) {
		var ret MeasurementClientFull
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeasurementClient201Response) GetDataOk() (*MeasurementClientFull, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateMeasurementClient201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MeasurementClientFull and assigns it to the Data field.
func (o *CreateMeasurementClient201Response) SetData(v MeasurementClientFull) {
	o.Data = &v
}

func (o CreateMeasurementClient201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMeasurementClient201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMeasurementClient201Response) UnmarshalJSON(data []byte) (err error) {
	varCreateMeasurementClient201Response := _CreateMeasurementClient201Response{}

	err = json.Unmarshal(data, &varCreateMeasurementClient201Response)

	if err != nil {
		return err
	}

	*o = CreateMeasurementClient201Response(varCreateMeasurementClient201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMeasurementClient201Response struct {
	value *CreateMeasurementClient201Response
	isSet bool
}

func (v NullableCreateMeasurementClient201Response) Get() *CreateMeasurementClient201Response {
	return v.value
}

func (v *NullableCreateMeasurementClient201Response) Set(val *CreateMeasurementClient201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMeasurementClient201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMeasurementClient201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMeasurementClient201Response(val *CreateMeasurementClient201Response) *NullableCreateMeasurementClient201Response {
	return &NullableCreateMeasurementClient201Response{value: val, isSet: true}
}

func (v NullableCreateMeasurementClient201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMeasurementClient201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


