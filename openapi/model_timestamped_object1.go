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
	"time"
	"fmt"
)

// checks if the TimestampedObject1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimestampedObject1{}

// TimestampedObject1 struct for TimestampedObject1
type TimestampedObject1 struct {
	// Date and time of the object creation
	CreatedAt time.Time `json:"created_at"`
	// ID of the user who created the object
	CreatedBy string `json:"created_by"`
	// Date and time of the latest update to the object
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the user who last updated the object
	UpdatedBy *string `json:"updated_by,omitempty"`
}

type _TimestampedObject1 TimestampedObject1

// NewTimestampedObject1 instantiates a new TimestampedObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampedObject1(createdAt time.Time, createdBy string) *TimestampedObject1 {
	this := TimestampedObject1{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	return &this
}

// NewTimestampedObject1WithDefaults instantiates a new TimestampedObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampedObject1WithDefaults() *TimestampedObject1 {
	this := TimestampedObject1{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TimestampedObject1) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TimestampedObject1) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TimestampedObject1) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *TimestampedObject1) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TimestampedObject1) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *TimestampedObject1) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TimestampedObject1) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampedObject1) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TimestampedObject1) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TimestampedObject1) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *TimestampedObject1) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampedObject1) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *TimestampedObject1) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *TimestampedObject1) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o TimestampedObject1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimestampedObject1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

func (o *TimestampedObject1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"created_by",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTimestampedObject1 := _TimestampedObject1{}

	err = json.Unmarshal(bytes, &varTimestampedObject1)

	if err != nil {
		return err
	}

	*o = TimestampedObject1(varTimestampedObject1)

	return err
}

type NullableTimestampedObject1 struct {
	value *TimestampedObject1
	isSet bool
}

func (v NullableTimestampedObject1) Get() *TimestampedObject1 {
	return v.value
}

func (v *NullableTimestampedObject1) Set(val *TimestampedObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampedObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampedObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampedObject1(val *TimestampedObject1) *NullableTimestampedObject1 {
	return &NullableTimestampedObject1{value: val, isSet: true}
}

func (v NullableTimestampedObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampedObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

