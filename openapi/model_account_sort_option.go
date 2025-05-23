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

// AccountSortOption the model 'AccountSortOption'
type AccountSortOption string

// List of AccountSortOption
const (
	ACCOUNTSORTOPTION_CREATED_AT AccountSortOption = "created_at"
	ACCOUNTSORTOPTION_CREATED_AT2 AccountSortOption = "-created_at"
	ACCOUNTSORTOPTION_NAME AccountSortOption = "name"
	ACCOUNTSORTOPTION_NAME2 AccountSortOption = "-name"
)

// All allowed values of AccountSortOption enum
var AllowedAccountSortOptionEnumValues = []AccountSortOption{
	"created_at",
	"-created_at",
	"name",
	"-name",
}

func (v *AccountSortOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountSortOption(value)
	for _, existing := range AllowedAccountSortOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountSortOption", value)
}

// NewAccountSortOptionFromValue returns a pointer to a valid AccountSortOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountSortOptionFromValue(v string) (*AccountSortOption, error) {
	ev := AccountSortOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountSortOption: valid values are %v", v, AllowedAccountSortOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountSortOption) IsValid() bool {
	for _, existing := range AllowedAccountSortOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountSortOption value
func (v AccountSortOption) Ptr() *AccountSortOption {
	return &v
}

type NullableAccountSortOption struct {
	value *AccountSortOption
	isSet bool
}

func (v NullableAccountSortOption) Get() *AccountSortOption {
	return v.value
}

func (v *NullableAccountSortOption) Set(val *AccountSortOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSortOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSortOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSortOption(val *AccountSortOption) *NullableAccountSortOption {
	return &NullableAccountSortOption{value: val, isSet: true}
}

func (v NullableAccountSortOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSortOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

