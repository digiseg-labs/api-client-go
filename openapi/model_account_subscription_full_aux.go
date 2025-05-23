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
	"time"
	"fmt"
)

// checks if the AccountSubscriptionFullAux type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSubscriptionFullAux{}

// AccountSubscriptionFullAux struct for AccountSubscriptionFullAux
type AccountSubscriptionFullAux struct {
	// The ID of the account
	AccountId string `json:"account_id"`
	Plan SubscriptionPlanFull `json:"plan"`
	// Timestamp of the last processed payment
	LastPaidAt *time.Time `json:"last_paid_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountSubscriptionFullAux AccountSubscriptionFullAux

// NewAccountSubscriptionFullAux instantiates a new AccountSubscriptionFullAux object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSubscriptionFullAux(accountId string, plan SubscriptionPlanFull) *AccountSubscriptionFullAux {
	this := AccountSubscriptionFullAux{}
	this.AccountId = accountId
	this.Plan = plan
	return &this
}

// NewAccountSubscriptionFullAuxWithDefaults instantiates a new AccountSubscriptionFullAux object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSubscriptionFullAuxWithDefaults() *AccountSubscriptionFullAux {
	this := AccountSubscriptionFullAux{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountSubscriptionFullAux) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountSubscriptionFullAux) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountSubscriptionFullAux) SetAccountId(v string) {
	o.AccountId = v
}

// GetPlan returns the Plan field value
func (o *AccountSubscriptionFullAux) GetPlan() SubscriptionPlanFull {
	if o == nil {
		var ret SubscriptionPlanFull
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *AccountSubscriptionFullAux) GetPlanOk() (*SubscriptionPlanFull, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *AccountSubscriptionFullAux) SetPlan(v SubscriptionPlanFull) {
	o.Plan = v
}

// GetLastPaidAt returns the LastPaidAt field value if set, zero value otherwise.
func (o *AccountSubscriptionFullAux) GetLastPaidAt() time.Time {
	if o == nil || IsNil(o.LastPaidAt) {
		var ret time.Time
		return ret
	}
	return *o.LastPaidAt
}

// GetLastPaidAtOk returns a tuple with the LastPaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSubscriptionFullAux) GetLastPaidAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPaidAt) {
		return nil, false
	}
	return o.LastPaidAt, true
}

// HasLastPaidAt returns a boolean if a field has been set.
func (o *AccountSubscriptionFullAux) HasLastPaidAt() bool {
	if o != nil && !IsNil(o.LastPaidAt) {
		return true
	}

	return false
}

// SetLastPaidAt gets a reference to the given time.Time and assigns it to the LastPaidAt field.
func (o *AccountSubscriptionFullAux) SetLastPaidAt(v time.Time) {
	o.LastPaidAt = &v
}

func (o AccountSubscriptionFullAux) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSubscriptionFullAux) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["plan"] = o.Plan
	if !IsNil(o.LastPaidAt) {
		toSerialize["last_paid_at"] = o.LastPaidAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountSubscriptionFullAux) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"plan",
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

	varAccountSubscriptionFullAux := _AccountSubscriptionFullAux{}

	err = json.Unmarshal(data, &varAccountSubscriptionFullAux)

	if err != nil {
		return err
	}

	*o = AccountSubscriptionFullAux(varAccountSubscriptionFullAux)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "last_paid_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountSubscriptionFullAux struct {
	value *AccountSubscriptionFullAux
	isSet bool
}

func (v NullableAccountSubscriptionFullAux) Get() *AccountSubscriptionFullAux {
	return v.value
}

func (v *NullableAccountSubscriptionFullAux) Set(val *AccountSubscriptionFullAux) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSubscriptionFullAux) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSubscriptionFullAux) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSubscriptionFullAux(val *AccountSubscriptionFullAux) *NullableAccountSubscriptionFullAux {
	return &NullableAccountSubscriptionFullAux{value: val, isSet: true}
}

func (v NullableAccountSubscriptionFullAux) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSubscriptionFullAux) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


