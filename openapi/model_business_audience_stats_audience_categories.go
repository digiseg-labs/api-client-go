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
	"bytes"
	"fmt"
)

// checks if the BusinessAudienceStatsAudienceCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessAudienceStatsAudienceCategories{}

// BusinessAudienceStatsAudienceCategories struct for BusinessAudienceStatsAudienceCategories
type BusinessAudienceStatsAudienceCategories struct {
	Size AudienceCategoryStats `json:"size"`
}

type _BusinessAudienceStatsAudienceCategories BusinessAudienceStatsAudienceCategories

// NewBusinessAudienceStatsAudienceCategories instantiates a new BusinessAudienceStatsAudienceCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessAudienceStatsAudienceCategories(size AudienceCategoryStats) *BusinessAudienceStatsAudienceCategories {
	this := BusinessAudienceStatsAudienceCategories{}
	this.Size = size
	return &this
}

// NewBusinessAudienceStatsAudienceCategoriesWithDefaults instantiates a new BusinessAudienceStatsAudienceCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessAudienceStatsAudienceCategoriesWithDefaults() *BusinessAudienceStatsAudienceCategories {
	this := BusinessAudienceStatsAudienceCategories{}
	return &this
}

// GetSize returns the Size field value
func (o *BusinessAudienceStatsAudienceCategories) GetSize() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *BusinessAudienceStatsAudienceCategories) GetSizeOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *BusinessAudienceStatsAudienceCategories) SetSize(v AudienceCategoryStats) {
	o.Size = v
}

func (o BusinessAudienceStatsAudienceCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessAudienceStatsAudienceCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *BusinessAudienceStatsAudienceCategories) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
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

	varBusinessAudienceStatsAudienceCategories := _BusinessAudienceStatsAudienceCategories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBusinessAudienceStatsAudienceCategories)

	if err != nil {
		return err
	}

	*o = BusinessAudienceStatsAudienceCategories(varBusinessAudienceStatsAudienceCategories)

	return err
}

type NullableBusinessAudienceStatsAudienceCategories struct {
	value *BusinessAudienceStatsAudienceCategories
	isSet bool
}

func (v NullableBusinessAudienceStatsAudienceCategories) Get() *BusinessAudienceStatsAudienceCategories {
	return v.value
}

func (v *NullableBusinessAudienceStatsAudienceCategories) Set(val *BusinessAudienceStatsAudienceCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessAudienceStatsAudienceCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessAudienceStatsAudienceCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessAudienceStatsAudienceCategories(val *BusinessAudienceStatsAudienceCategories) *NullableBusinessAudienceStatsAudienceCategories {
	return &NullableBusinessAudienceStatsAudienceCategories{value: val, isSet: true}
}

func (v NullableBusinessAudienceStatsAudienceCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessAudienceStatsAudienceCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


