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

// checks if the MeasurementIntegrationPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasurementIntegrationPlatform{}

// MeasurementIntegrationPlatform The integration/platform with which a study is being delivered 
type MeasurementIntegrationPlatform struct {
	// An ID for the integration platform, if the integration platform is a known platform. Note that integration platform ID uniqueness is a responsibility of the client since this is simply an optional reference point to keep. Can be null/omitted if the platform name is a one-off value with just a string. 
	Id *string `json:"id,omitempty"`
	// The name of the integration platform.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MeasurementIntegrationPlatform MeasurementIntegrationPlatform

// NewMeasurementIntegrationPlatform instantiates a new MeasurementIntegrationPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementIntegrationPlatform() *MeasurementIntegrationPlatform {
	this := MeasurementIntegrationPlatform{}
	return &this
}

// NewMeasurementIntegrationPlatformWithDefaults instantiates a new MeasurementIntegrationPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementIntegrationPlatformWithDefaults() *MeasurementIntegrationPlatform {
	this := MeasurementIntegrationPlatform{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MeasurementIntegrationPlatform) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementIntegrationPlatform) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MeasurementIntegrationPlatform) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MeasurementIntegrationPlatform) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MeasurementIntegrationPlatform) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementIntegrationPlatform) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MeasurementIntegrationPlatform) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MeasurementIntegrationPlatform) SetName(v string) {
	o.Name = &v
}

func (o MeasurementIntegrationPlatform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasurementIntegrationPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MeasurementIntegrationPlatform) UnmarshalJSON(data []byte) (err error) {
	varMeasurementIntegrationPlatform := _MeasurementIntegrationPlatform{}

	err = json.Unmarshal(data, &varMeasurementIntegrationPlatform)

	if err != nil {
		return err
	}

	*o = MeasurementIntegrationPlatform(varMeasurementIntegrationPlatform)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeasurementIntegrationPlatform struct {
	value *MeasurementIntegrationPlatform
	isSet bool
}

func (v NullableMeasurementIntegrationPlatform) Get() *MeasurementIntegrationPlatform {
	return v.value
}

func (v *NullableMeasurementIntegrationPlatform) Set(val *MeasurementIntegrationPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementIntegrationPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementIntegrationPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementIntegrationPlatform(val *MeasurementIntegrationPlatform) *NullableMeasurementIntegrationPlatform {
	return &NullableMeasurementIntegrationPlatform{value: val, isSet: true}
}

func (v NullableMeasurementIntegrationPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementIntegrationPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


