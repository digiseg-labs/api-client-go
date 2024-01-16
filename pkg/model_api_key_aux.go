/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiKeyAux type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyAux{}

// ApiKeyAux struct for ApiKeyAux
type ApiKeyAux struct {
	Scopes *PermissionScopes `json:"scopes,omitempty"`
}

// NewApiKeyAux instantiates a new ApiKeyAux object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyAux() *ApiKeyAux {
	this := ApiKeyAux{}
	return &this
}

// NewApiKeyAuxWithDefaults instantiates a new ApiKeyAux object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyAuxWithDefaults() *ApiKeyAux {
	this := ApiKeyAux{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ApiKeyAux) GetScopes() PermissionScopes {
	if o == nil || IsNil(o.Scopes) {
		var ret PermissionScopes
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyAux) GetScopesOk() (*PermissionScopes, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApiKeyAux) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given PermissionScopes and assigns it to the Scopes field.
func (o *ApiKeyAux) SetScopes(v PermissionScopes) {
	o.Scopes = &v
}

func (o ApiKeyAux) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyAux) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableApiKeyAux struct {
	value *ApiKeyAux
	isSet bool
}

func (v NullableApiKeyAux) Get() *ApiKeyAux {
	return v.value
}

func (v *NullableApiKeyAux) Set(val *ApiKeyAux) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyAux) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyAux) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyAux(val *ApiKeyAux) *NullableApiKeyAux {
	return &NullableApiKeyAux{value: val, isSet: true}
}

func (v NullableApiKeyAux) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyAux) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


