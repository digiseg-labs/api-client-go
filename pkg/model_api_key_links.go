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

// checks if the ApiKeyLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyLinks{}

// ApiKeyLinks struct for ApiKeyLinks
type ApiKeyLinks struct {
	// Link for getting to the api key's user
	User *string `json:"user,omitempty"`
}

// NewApiKeyLinks instantiates a new ApiKeyLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyLinks() *ApiKeyLinks {
	this := ApiKeyLinks{}
	return &this
}

// NewApiKeyLinksWithDefaults instantiates a new ApiKeyLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyLinksWithDefaults() *ApiKeyLinks {
	this := ApiKeyLinks{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ApiKeyLinks) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyLinks) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ApiKeyLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ApiKeyLinks) SetUser(v string) {
	o.User = &v
}

func (o ApiKeyLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableApiKeyLinks struct {
	value *ApiKeyLinks
	isSet bool
}

func (v NullableApiKeyLinks) Get() *ApiKeyLinks {
	return v.value
}

func (v *NullableApiKeyLinks) Set(val *ApiKeyLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyLinks(val *ApiKeyLinks) *NullableApiKeyLinks {
	return &NullableApiKeyLinks{value: val, isSet: true}
}

func (v NullableApiKeyLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

