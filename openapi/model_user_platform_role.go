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

// UserPlatformRole Defines the role of the user in the Digiseg platform. Most users will not have any platform roles, since this is preserved for account-agnostic operational privileges. 
type UserPlatformRole string

// List of UserPlatformRole
const (
	USERPLATFORMROLE_SUPER_ADMIN UserPlatformRole = "super_admin"
	USERPLATFORMROLE_DEV_OPS UserPlatformRole = "dev_ops"
	USERPLATFORMROLE_CUSTOMER_OPS UserPlatformRole = "customer_ops"
)

// All allowed values of UserPlatformRole enum
var AllowedUserPlatformRoleEnumValues = []UserPlatformRole{
	"super_admin",
	"dev_ops",
	"customer_ops",
}

func (v *UserPlatformRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserPlatformRole(value)
	for _, existing := range AllowedUserPlatformRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserPlatformRole", value)
}

// NewUserPlatformRoleFromValue returns a pointer to a valid UserPlatformRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserPlatformRoleFromValue(v string) (*UserPlatformRole, error) {
	ev := UserPlatformRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserPlatformRole: valid values are %v", v, AllowedUserPlatformRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserPlatformRole) IsValid() bool {
	for _, existing := range AllowedUserPlatformRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserPlatformRole value
func (v UserPlatformRole) Ptr() *UserPlatformRole {
	return &v
}

type NullableUserPlatformRole struct {
	value *UserPlatformRole
	isSet bool
}

func (v NullableUserPlatformRole) Get() *UserPlatformRole {
	return v.value
}

func (v *NullableUserPlatformRole) Set(val *UserPlatformRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlatformRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlatformRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlatformRole(val *UserPlatformRole) *NullableUserPlatformRole {
	return &NullableUserPlatformRole{value: val, isSet: true}
}

func (v NullableUserPlatformRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlatformRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

