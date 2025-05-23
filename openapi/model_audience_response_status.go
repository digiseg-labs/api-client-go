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

// AudienceResponseStatus Describes the result of resolving the audiences of an IP address. `resolved` means that the IP was resolved and audiences are returned. `unresolved` means that the IP could not be resolved and no audiences are returned. `forbidden` means that the authenticated user did not have grants to resolve audiences in the country of the IP. 
type AudienceResponseStatus string

// List of AudienceResponseStatus
const (
	AUDIENCERESPONSESTATUS_RESOLVED AudienceResponseStatus = "resolved"
	AUDIENCERESPONSESTATUS_UNRESOLVED AudienceResponseStatus = "unresolved"
	AUDIENCERESPONSESTATUS_FORBIDDEN AudienceResponseStatus = "forbidden"
)

// All allowed values of AudienceResponseStatus enum
var AllowedAudienceResponseStatusEnumValues = []AudienceResponseStatus{
	"resolved",
	"unresolved",
	"forbidden",
}

func (v *AudienceResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AudienceResponseStatus(value)
	for _, existing := range AllowedAudienceResponseStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AudienceResponseStatus", value)
}

// NewAudienceResponseStatusFromValue returns a pointer to a valid AudienceResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudienceResponseStatusFromValue(v string) (*AudienceResponseStatus, error) {
	ev := AudienceResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudienceResponseStatus: valid values are %v", v, AllowedAudienceResponseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudienceResponseStatus) IsValid() bool {
	for _, existing := range AllowedAudienceResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudienceResponseStatus value
func (v AudienceResponseStatus) Ptr() *AudienceResponseStatus {
	return &v
}

type NullableAudienceResponseStatus struct {
	value *AudienceResponseStatus
	isSet bool
}

func (v NullableAudienceResponseStatus) Get() *AudienceResponseStatus {
	return v.value
}

func (v *NullableAudienceResponseStatus) Set(val *AudienceResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceResponseStatus(val *AudienceResponseStatus) *NullableAudienceResponseStatus {
	return &NullableAudienceResponseStatus{value: val, isSet: true}
}

func (v NullableAudienceResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

