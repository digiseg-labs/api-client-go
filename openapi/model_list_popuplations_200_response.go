/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListPopuplations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPopuplations200Response{}

// ListPopuplations200Response struct for ListPopuplations200Response
type ListPopuplations200Response struct {
	Data *CategoryPopulationsFull `json:"data,omitempty"`
}

// NewListPopuplations200Response instantiates a new ListPopuplations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPopuplations200Response() *ListPopuplations200Response {
	this := ListPopuplations200Response{}
	return &this
}

// NewListPopuplations200ResponseWithDefaults instantiates a new ListPopuplations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPopuplations200ResponseWithDefaults() *ListPopuplations200Response {
	this := ListPopuplations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListPopuplations200Response) GetData() CategoryPopulationsFull {
	if o == nil || IsNil(o.Data) {
		var ret CategoryPopulationsFull
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPopuplations200Response) GetDataOk() (*CategoryPopulationsFull, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListPopuplations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CategoryPopulationsFull and assigns it to the Data field.
func (o *ListPopuplations200Response) SetData(v CategoryPopulationsFull) {
	o.Data = &v
}

func (o ListPopuplations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPopuplations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListPopuplations200Response struct {
	value *ListPopuplations200Response
	isSet bool
}

func (v NullableListPopuplations200Response) Get() *ListPopuplations200Response {
	return v.value
}

func (v *NullableListPopuplations200Response) Set(val *ListPopuplations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPopuplations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPopuplations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPopuplations200Response(val *ListPopuplations200Response) *NullableListPopuplations200Response {
	return &NullableListPopuplations200Response{value: val, isSet: true}
}

func (v NullableListPopuplations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPopuplations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


