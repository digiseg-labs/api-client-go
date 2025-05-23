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

// CompanySize Describes the size of a company
type CompanySize string

// List of CompanySize
const (
	COMPANYSIZE_SELF_EMPLOYED CompanySize = "Self-employed"
	COMPANYSIZE__1_10_EMPLOYEES CompanySize = "1-10 employees"
	COMPANYSIZE__11_50_EMPLOYEES CompanySize = "11-50 employees"
	COMPANYSIZE__51_200_EMPLOYEES CompanySize = "51-200 employees"
	COMPANYSIZE__201_500_EMPLOYEES CompanySize = "201-500 employees"
	COMPANYSIZE__501_1000_EMPLOYEES CompanySize = "501-1000 employees"
	COMPANYSIZE__1001_5000_EMPLOYEES CompanySize = "1001-5000 employees"
	COMPANYSIZE__5001_10000_EMPLOYEES CompanySize = "5001-10,000 employees"
	COMPANYSIZE__10001_EMPLOYEES CompanySize = "10,001+ employees"
)

// All allowed values of CompanySize enum
var AllowedCompanySizeEnumValues = []CompanySize{
	"Self-employed",
	"1-10 employees",
	"11-50 employees",
	"51-200 employees",
	"201-500 employees",
	"501-1000 employees",
	"1001-5000 employees",
	"5001-10,000 employees",
	"10,001+ employees",
}

func (v *CompanySize) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompanySize(value)
	for _, existing := range AllowedCompanySizeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompanySize", value)
}

// NewCompanySizeFromValue returns a pointer to a valid CompanySize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompanySizeFromValue(v string) (*CompanySize, error) {
	ev := CompanySize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompanySize: valid values are %v", v, AllowedCompanySizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompanySize) IsValid() bool {
	for _, existing := range AllowedCompanySizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompanySize value
func (v CompanySize) Ptr() *CompanySize {
	return &v
}

type NullableCompanySize struct {
	value *CompanySize
	isSet bool
}

func (v NullableCompanySize) Get() *CompanySize {
	return v.value
}

func (v *NullableCompanySize) Set(val *CompanySize) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanySize) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanySize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanySize(val *CompanySize) *NullableCompanySize {
	return &NullableCompanySize{value: val, isSet: true}
}

func (v NullableCompanySize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanySize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

