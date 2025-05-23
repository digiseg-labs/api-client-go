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

// checks if the AccountOwnerCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountOwnerCreation{}

// AccountOwnerCreation Information about the account owner of a newly created account. Can only be specified by super-admins.
type AccountOwnerCreation struct {
	// The account owner's name
	Name string `json:"name"`
	// The account owner's email
	Email string `json:"email"`
	// The account owner's password
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountOwnerCreation AccountOwnerCreation

// NewAccountOwnerCreation instantiates a new AccountOwnerCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOwnerCreation(name string, email string) *AccountOwnerCreation {
	this := AccountOwnerCreation{}
	this.Name = name
	this.Email = email
	return &this
}

// NewAccountOwnerCreationWithDefaults instantiates a new AccountOwnerCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOwnerCreationWithDefaults() *AccountOwnerCreation {
	this := AccountOwnerCreation{}
	return &this
}

// GetName returns the Name field value
func (o *AccountOwnerCreation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountOwnerCreation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountOwnerCreation) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *AccountOwnerCreation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccountOwnerCreation) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccountOwnerCreation) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AccountOwnerCreation) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOwnerCreation) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AccountOwnerCreation) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AccountOwnerCreation) SetPassword(v string) {
	o.Password = &v
}

func (o AccountOwnerCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOwnerCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountOwnerCreation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
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

	varAccountOwnerCreation := _AccountOwnerCreation{}

	err = json.Unmarshal(data, &varAccountOwnerCreation)

	if err != nil {
		return err
	}

	*o = AccountOwnerCreation(varAccountOwnerCreation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountOwnerCreation struct {
	value *AccountOwnerCreation
	isSet bool
}

func (v NullableAccountOwnerCreation) Get() *AccountOwnerCreation {
	return v.value
}

func (v *NullableAccountOwnerCreation) Set(val *AccountOwnerCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOwnerCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOwnerCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOwnerCreation(val *AccountOwnerCreation) *NullableAccountOwnerCreation {
	return &NullableAccountOwnerCreation{value: val, isSet: true}
}

func (v NullableAccountOwnerCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOwnerCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


