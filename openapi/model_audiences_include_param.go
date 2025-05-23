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

// AudiencesIncludeParam the model 'AudiencesIncludeParam'
type AudiencesIncludeParam string

// List of AudiencesIncludeParam
const (
	AUDIENCESINCLUDEPARAM_CORE AudiencesIncludeParam = "core"
	AUDIENCESINCLUDEPARAM_COMPOSITE AudiencesIncludeParam = "composite"
	AUDIENCESINCLUDEPARAM_NAME AudiencesIncludeParam = "name"
	AUDIENCESINCLUDEPARAM_CATEGORY AudiencesIncludeParam = "category"
)

// All allowed values of AudiencesIncludeParam enum
var AllowedAudiencesIncludeParamEnumValues = []AudiencesIncludeParam{
	"core",
	"composite",
	"name",
	"category",
}

func (v *AudiencesIncludeParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AudiencesIncludeParam(value)
	for _, existing := range AllowedAudiencesIncludeParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AudiencesIncludeParam", value)
}

// NewAudiencesIncludeParamFromValue returns a pointer to a valid AudiencesIncludeParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencesIncludeParamFromValue(v string) (*AudiencesIncludeParam, error) {
	ev := AudiencesIncludeParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencesIncludeParam: valid values are %v", v, AllowedAudiencesIncludeParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencesIncludeParam) IsValid() bool {
	for _, existing := range AllowedAudiencesIncludeParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudiencesIncludeParam value
func (v AudiencesIncludeParam) Ptr() *AudiencesIncludeParam {
	return &v
}

type NullableAudiencesIncludeParam struct {
	value *AudiencesIncludeParam
	isSet bool
}

func (v NullableAudiencesIncludeParam) Get() *AudiencesIncludeParam {
	return v.value
}

func (v *NullableAudiencesIncludeParam) Set(val *AudiencesIncludeParam) {
	v.value = val
	v.isSet = true
}

func (v NullableAudiencesIncludeParam) IsSet() bool {
	return v.isSet
}

func (v *NullableAudiencesIncludeParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudiencesIncludeParam(val *AudiencesIncludeParam) *NullableAudiencesIncludeParam {
	return &NullableAudiencesIncludeParam{value: val, isSet: true}
}

func (v NullableAudiencesIncludeParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudiencesIncludeParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

