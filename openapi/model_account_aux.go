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

// checks if the AccountAux type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAux{}

// AccountAux struct for AccountAux
type AccountAux struct {
	// ID of the user who is the ultimate owner of the account
	OwnerId *string `json:"owner_id,omitempty"`
}

// NewAccountAux instantiates a new AccountAux object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAux() *AccountAux {
	this := AccountAux{}
	return &this
}

// NewAccountAuxWithDefaults instantiates a new AccountAux object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAuxWithDefaults() *AccountAux {
	this := AccountAux{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AccountAux) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAux) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AccountAux) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AccountAux) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o AccountAux) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAux) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	return toSerialize, nil
}

type NullableAccountAux struct {
	value *AccountAux
	isSet bool
}

func (v NullableAccountAux) Get() *AccountAux {
	return v.value
}

func (v *NullableAccountAux) Set(val *AccountAux) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAux) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAux) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAux(val *AccountAux) *NullableAccountAux {
	return &NullableAccountAux{value: val, isSet: true}
}

func (v NullableAccountAux) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAux) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

