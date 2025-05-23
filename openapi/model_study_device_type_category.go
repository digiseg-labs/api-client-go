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

// StudyDeviceTypeCategory the model 'StudyDeviceTypeCategory'
type StudyDeviceTypeCategory string

// List of StudyDeviceTypeCategory
const (
	STUDYDEVICETYPECATEGORY_MOBILE StudyDeviceTypeCategory = "mobile"
	STUDYDEVICETYPECATEGORY_DESKTOP StudyDeviceTypeCategory = "desktop"
	STUDYDEVICETYPECATEGORY_TV StudyDeviceTypeCategory = "tv"
	STUDYDEVICETYPECATEGORY_BOT StudyDeviceTypeCategory = "bot"
	STUDYDEVICETYPECATEGORY_OTHER StudyDeviceTypeCategory = "other"
	STUDYDEVICETYPECATEGORY_UNKNOWN StudyDeviceTypeCategory = "unknown"
)

// All allowed values of StudyDeviceTypeCategory enum
var AllowedStudyDeviceTypeCategoryEnumValues = []StudyDeviceTypeCategory{
	"mobile",
	"desktop",
	"tv",
	"bot",
	"other",
	"unknown",
}

func (v *StudyDeviceTypeCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StudyDeviceTypeCategory(value)
	for _, existing := range AllowedStudyDeviceTypeCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StudyDeviceTypeCategory", value)
}

// NewStudyDeviceTypeCategoryFromValue returns a pointer to a valid StudyDeviceTypeCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStudyDeviceTypeCategoryFromValue(v string) (*StudyDeviceTypeCategory, error) {
	ev := StudyDeviceTypeCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StudyDeviceTypeCategory: valid values are %v", v, AllowedStudyDeviceTypeCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StudyDeviceTypeCategory) IsValid() bool {
	for _, existing := range AllowedStudyDeviceTypeCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StudyDeviceTypeCategory value
func (v StudyDeviceTypeCategory) Ptr() *StudyDeviceTypeCategory {
	return &v
}

type NullableStudyDeviceTypeCategory struct {
	value *StudyDeviceTypeCategory
	isSet bool
}

func (v NullableStudyDeviceTypeCategory) Get() *StudyDeviceTypeCategory {
	return v.value
}

func (v *NullableStudyDeviceTypeCategory) Set(val *StudyDeviceTypeCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyDeviceTypeCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyDeviceTypeCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyDeviceTypeCategory(val *StudyDeviceTypeCategory) *NullableStudyDeviceTypeCategory {
	return &NullableStudyDeviceTypeCategory{value: val, isSet: true}
}

func (v NullableStudyDeviceTypeCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyDeviceTypeCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

