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
)

// checks if the AccountMutationAux type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountMutationAux{}

// AccountMutationAux struct for AccountMutationAux
type AccountMutationAux struct {
	CompanyType *CompanyType `json:"company_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountMutationAux AccountMutationAux

// NewAccountMutationAux instantiates a new AccountMutationAux object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMutationAux() *AccountMutationAux {
	this := AccountMutationAux{}
	return &this
}

// NewAccountMutationAuxWithDefaults instantiates a new AccountMutationAux object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMutationAuxWithDefaults() *AccountMutationAux {
	this := AccountMutationAux{}
	return &this
}

// GetCompanyType returns the CompanyType field value if set, zero value otherwise.
func (o *AccountMutationAux) GetCompanyType() CompanyType {
	if o == nil || IsNil(o.CompanyType) {
		var ret CompanyType
		return ret
	}
	return *o.CompanyType
}

// GetCompanyTypeOk returns a tuple with the CompanyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMutationAux) GetCompanyTypeOk() (*CompanyType, bool) {
	if o == nil || IsNil(o.CompanyType) {
		return nil, false
	}
	return o.CompanyType, true
}

// HasCompanyType returns a boolean if a field has been set.
func (o *AccountMutationAux) HasCompanyType() bool {
	if o != nil && !IsNil(o.CompanyType) {
		return true
	}

	return false
}

// SetCompanyType gets a reference to the given CompanyType and assigns it to the CompanyType field.
func (o *AccountMutationAux) SetCompanyType(v CompanyType) {
	o.CompanyType = &v
}

func (o AccountMutationAux) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountMutationAux) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyType) {
		toSerialize["company_type"] = o.CompanyType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountMutationAux) UnmarshalJSON(data []byte) (err error) {
	varAccountMutationAux := _AccountMutationAux{}

	err = json.Unmarshal(data, &varAccountMutationAux)

	if err != nil {
		return err
	}

	*o = AccountMutationAux(varAccountMutationAux)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountMutationAux struct {
	value *AccountMutationAux
	isSet bool
}

func (v NullableAccountMutationAux) Get() *AccountMutationAux {
	return v.value
}

func (v *NullableAccountMutationAux) Set(val *AccountMutationAux) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMutationAux) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMutationAux) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMutationAux(val *AccountMutationAux) *NullableAccountMutationAux {
	return &NullableAccountMutationAux{value: val, isSet: true}
}

func (v NullableAccountMutationAux) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMutationAux) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


