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
	"time"
)

// checks if the ApiKeyBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyBase{}

// ApiKeyBase struct for ApiKeyBase
type ApiKeyBase struct {
	// Human readable name of the API key
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	// Optional date/time that the key will expire
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The ID of the API key's user
	UserId *string `json:"user_id,omitempty"`
}

// NewApiKeyBase instantiates a new ApiKeyBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyBase() *ApiKeyBase {
	this := ApiKeyBase{}
	return &this
}

// NewApiKeyBaseWithDefaults instantiates a new ApiKeyBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyBaseWithDefaults() *ApiKeyBase {
	this := ApiKeyBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiKeyBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiKeyBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiKeyBase) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiKeyBase) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyBase) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiKeyBase) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiKeyBase) SetStatus(v string) {
	o.Status = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApiKeyBase) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyBase) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiKeyBase) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApiKeyBase) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiKeyBase) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyBase) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiKeyBase) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiKeyBase) SetUserId(v string) {
	o.UserId = &v
}

func (o ApiKeyBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableApiKeyBase struct {
	value *ApiKeyBase
	isSet bool
}

func (v NullableApiKeyBase) Get() *ApiKeyBase {
	return v.value
}

func (v *NullableApiKeyBase) Set(val *ApiKeyBase) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyBase) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyBase(val *ApiKeyBase) *NullableApiKeyBase {
	return &NullableApiKeyBase{value: val, isSet: true}
}

func (v NullableApiKeyBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

