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
	"fmt"
)

// checks if the ResolveAudiencesOfMultipleRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveAudiencesOfMultipleRequestItem{}

// ResolveAudiencesOfMultipleRequestItem struct for ResolveAudiencesOfMultipleRequestItem
type ResolveAudiencesOfMultipleRequestItem struct {
	// An optional identifier for the item. The identifier will also be available in the response.
	Id *string `json:"id,omitempty"`
	// The IP address to resolve audiences for
	IpAddress string `json:"ip_address"`
	AdditionalProperties map[string]interface{}
}

type _ResolveAudiencesOfMultipleRequestItem ResolveAudiencesOfMultipleRequestItem

// NewResolveAudiencesOfMultipleRequestItem instantiates a new ResolveAudiencesOfMultipleRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveAudiencesOfMultipleRequestItem(ipAddress string) *ResolveAudiencesOfMultipleRequestItem {
	this := ResolveAudiencesOfMultipleRequestItem{}
	this.IpAddress = ipAddress
	return &this
}

// NewResolveAudiencesOfMultipleRequestItemWithDefaults instantiates a new ResolveAudiencesOfMultipleRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveAudiencesOfMultipleRequestItemWithDefaults() *ResolveAudiencesOfMultipleRequestItem {
	this := ResolveAudiencesOfMultipleRequestItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResolveAudiencesOfMultipleRequestItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleRequestItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResolveAudiencesOfMultipleRequestItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResolveAudiencesOfMultipleRequestItem) SetId(v string) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value
func (o *ResolveAudiencesOfMultipleRequestItem) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleRequestItem) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *ResolveAudiencesOfMultipleRequestItem) SetIpAddress(v string) {
	o.IpAddress = v
}

func (o ResolveAudiencesOfMultipleRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveAudiencesOfMultipleRequestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["ip_address"] = o.IpAddress

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResolveAudiencesOfMultipleRequestItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip_address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResolveAudiencesOfMultipleRequestItem := _ResolveAudiencesOfMultipleRequestItem{}

	err = json.Unmarshal(data, &varResolveAudiencesOfMultipleRequestItem)

	if err != nil {
		return err
	}

	*o = ResolveAudiencesOfMultipleRequestItem(varResolveAudiencesOfMultipleRequestItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResolveAudiencesOfMultipleRequestItem struct {
	value *ResolveAudiencesOfMultipleRequestItem
	isSet bool
}

func (v NullableResolveAudiencesOfMultipleRequestItem) Get() *ResolveAudiencesOfMultipleRequestItem {
	return v.value
}

func (v *NullableResolveAudiencesOfMultipleRequestItem) Set(val *ResolveAudiencesOfMultipleRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveAudiencesOfMultipleRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveAudiencesOfMultipleRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveAudiencesOfMultipleRequestItem(val *ResolveAudiencesOfMultipleRequestItem) *NullableResolveAudiencesOfMultipleRequestItem {
	return &NullableResolveAudiencesOfMultipleRequestItem{value: val, isSet: true}
}

func (v NullableResolveAudiencesOfMultipleRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveAudiencesOfMultipleRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


