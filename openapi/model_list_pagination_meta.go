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

// checks if the ListPaginationMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPaginationMeta{}

// ListPaginationMeta struct for ListPaginationMeta
type ListPaginationMeta struct {
	Page *ListPaginationMetaPage `json:"page,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPaginationMeta ListPaginationMeta

// NewListPaginationMeta instantiates a new ListPaginationMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaginationMeta() *ListPaginationMeta {
	this := ListPaginationMeta{}
	return &this
}

// NewListPaginationMetaWithDefaults instantiates a new ListPaginationMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaginationMetaWithDefaults() *ListPaginationMeta {
	this := ListPaginationMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListPaginationMeta) GetPage() ListPaginationMetaPage {
	if o == nil || IsNil(o.Page) {
		var ret ListPaginationMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaginationMeta) GetPageOk() (*ListPaginationMetaPage, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListPaginationMeta) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ListPaginationMetaPage and assigns it to the Page field.
func (o *ListPaginationMeta) SetPage(v ListPaginationMetaPage) {
	o.Page = &v
}

func (o ListPaginationMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPaginationMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPaginationMeta) UnmarshalJSON(data []byte) (err error) {
	varListPaginationMeta := _ListPaginationMeta{}

	err = json.Unmarshal(data, &varListPaginationMeta)

	if err != nil {
		return err
	}

	*o = ListPaginationMeta(varListPaginationMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPaginationMeta struct {
	value *ListPaginationMeta
	isSet bool
}

func (v NullableListPaginationMeta) Get() *ListPaginationMeta {
	return v.value
}

func (v *NullableListPaginationMeta) Set(val *ListPaginationMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaginationMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaginationMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaginationMeta(val *ListPaginationMeta) *NullableListPaginationMeta {
	return &NullableListPaginationMeta{value: val, isSet: true}
}

func (v NullableListPaginationMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaginationMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


