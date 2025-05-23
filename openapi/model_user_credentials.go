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

// checks if the UserCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCredentials{}

// UserCredentials struct for UserCredentials
type UserCredentials struct {
	// Password of the user
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCredentials UserCredentials

// NewUserCredentials instantiates a new UserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredentials() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// NewUserCredentialsWithDefaults instantiates a new UserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialsWithDefaults() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o UserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCredentials) UnmarshalJSON(data []byte) (err error) {
	varUserCredentials := _UserCredentials{}

	err = json.Unmarshal(data, &varUserCredentials)

	if err != nil {
		return err
	}

	*o = UserCredentials(varUserCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCredentials struct {
	value *UserCredentials
	isSet bool
}

func (v NullableUserCredentials) Get() *UserCredentials {
	return v.value
}

func (v *NullableUserCredentials) Set(val *UserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredentials(val *UserCredentials) *NullableUserCredentials {
	return &NullableUserCredentials{value: val, isSet: true}
}

func (v NullableUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


