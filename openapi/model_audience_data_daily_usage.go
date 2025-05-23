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

// checks if the AudienceDataDailyUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceDataDailyUsage{}

// AudienceDataDailyUsage struct for AudienceDataDailyUsage
type AudienceDataDailyUsage struct {
	// Date in ISO 8601 format
	Date string `json:"date"`
	Usage AudienceDataUsage `json:"usage"`
	AdditionalProperties map[string]interface{}
}

type _AudienceDataDailyUsage AudienceDataDailyUsage

// NewAudienceDataDailyUsage instantiates a new AudienceDataDailyUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceDataDailyUsage(date string, usage AudienceDataUsage) *AudienceDataDailyUsage {
	this := AudienceDataDailyUsage{}
	this.Date = date
	this.Usage = usage
	return &this
}

// NewAudienceDataDailyUsageWithDefaults instantiates a new AudienceDataDailyUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceDataDailyUsageWithDefaults() *AudienceDataDailyUsage {
	this := AudienceDataDailyUsage{}
	return &this
}

// GetDate returns the Date field value
func (o *AudienceDataDailyUsage) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *AudienceDataDailyUsage) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *AudienceDataDailyUsage) SetDate(v string) {
	o.Date = v
}

// GetUsage returns the Usage field value
func (o *AudienceDataDailyUsage) GetUsage() AudienceDataUsage {
	if o == nil {
		var ret AudienceDataUsage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *AudienceDataDailyUsage) GetUsageOk() (*AudienceDataUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *AudienceDataDailyUsage) SetUsage(v AudienceDataUsage) {
	o.Usage = v
}

func (o AudienceDataDailyUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceDataDailyUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AudienceDataDailyUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"usage",
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

	varAudienceDataDailyUsage := _AudienceDataDailyUsage{}

	err = json.Unmarshal(data, &varAudienceDataDailyUsage)

	if err != nil {
		return err
	}

	*o = AudienceDataDailyUsage(varAudienceDataDailyUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAudienceDataDailyUsage struct {
	value *AudienceDataDailyUsage
	isSet bool
}

func (v NullableAudienceDataDailyUsage) Get() *AudienceDataDailyUsage {
	return v.value
}

func (v *NullableAudienceDataDailyUsage) Set(val *AudienceDataDailyUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceDataDailyUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceDataDailyUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceDataDailyUsage(val *AudienceDataDailyUsage) *NullableAudienceDataDailyUsage {
	return &NullableAudienceDataDailyUsage{value: val, isSet: true}
}

func (v NullableAudienceDataDailyUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceDataDailyUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


