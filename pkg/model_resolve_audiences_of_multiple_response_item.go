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
	"fmt"
)

// checks if the ResolveAudiencesOfMultipleResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveAudiencesOfMultipleResponseItem{}

// ResolveAudiencesOfMultipleResponseItem struct for ResolveAudiencesOfMultipleResponseItem
type ResolveAudiencesOfMultipleResponseItem struct {
	// The identifier (if provided) of the item as it was provided in the request.
	Id *string `json:"id,omitempty"`
	Audiences []Audience `json:"audiences,omitempty"`
	Status AudienceResponseStatus `json:"status"`
}

type _ResolveAudiencesOfMultipleResponseItem ResolveAudiencesOfMultipleResponseItem

// NewResolveAudiencesOfMultipleResponseItem instantiates a new ResolveAudiencesOfMultipleResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveAudiencesOfMultipleResponseItem(status AudienceResponseStatus) *ResolveAudiencesOfMultipleResponseItem {
	this := ResolveAudiencesOfMultipleResponseItem{}
	this.Status = status
	return &this
}

// NewResolveAudiencesOfMultipleResponseItemWithDefaults instantiates a new ResolveAudiencesOfMultipleResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveAudiencesOfMultipleResponseItemWithDefaults() *ResolveAudiencesOfMultipleResponseItem {
	this := ResolveAudiencesOfMultipleResponseItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResolveAudiencesOfMultipleResponseItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleResponseItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResolveAudiencesOfMultipleResponseItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResolveAudiencesOfMultipleResponseItem) SetId(v string) {
	o.Id = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *ResolveAudiencesOfMultipleResponseItem) GetAudiences() []Audience {
	if o == nil || IsNil(o.Audiences) {
		var ret []Audience
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleResponseItem) GetAudiencesOk() ([]Audience, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *ResolveAudiencesOfMultipleResponseItem) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []Audience and assigns it to the Audiences field.
func (o *ResolveAudiencesOfMultipleResponseItem) SetAudiences(v []Audience) {
	o.Audiences = v
}

// GetStatus returns the Status field value
func (o *ResolveAudiencesOfMultipleResponseItem) GetStatus() AudienceResponseStatus {
	if o == nil {
		var ret AudienceResponseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleResponseItem) GetStatusOk() (*AudienceResponseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResolveAudiencesOfMultipleResponseItem) SetStatus(v AudienceResponseStatus) {
	o.Status = v
}

func (o ResolveAudiencesOfMultipleResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveAudiencesOfMultipleResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ResolveAudiencesOfMultipleResponseItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varResolveAudiencesOfMultipleResponseItem := _ResolveAudiencesOfMultipleResponseItem{}

	err = json.Unmarshal(bytes, &varResolveAudiencesOfMultipleResponseItem)

	if err != nil {
		return err
	}

	*o = ResolveAudiencesOfMultipleResponseItem(varResolveAudiencesOfMultipleResponseItem)

	return err
}

type NullableResolveAudiencesOfMultipleResponseItem struct {
	value *ResolveAudiencesOfMultipleResponseItem
	isSet bool
}

func (v NullableResolveAudiencesOfMultipleResponseItem) Get() *ResolveAudiencesOfMultipleResponseItem {
	return v.value
}

func (v *NullableResolveAudiencesOfMultipleResponseItem) Set(val *ResolveAudiencesOfMultipleResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveAudiencesOfMultipleResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveAudiencesOfMultipleResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveAudiencesOfMultipleResponseItem(val *ResolveAudiencesOfMultipleResponseItem) *NullableResolveAudiencesOfMultipleResponseItem {
	return &NullableResolveAudiencesOfMultipleResponseItem{value: val, isSet: true}
}

func (v NullableResolveAudiencesOfMultipleResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveAudiencesOfMultipleResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


