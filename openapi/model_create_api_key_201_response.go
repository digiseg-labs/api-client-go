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

// checks if the CreateApiKey201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiKey201Response{}

// CreateApiKey201Response struct for CreateApiKey201Response
type CreateApiKey201Response struct {
	Data *ApiKeyFullWithToken `json:"data,omitempty"`
}

// NewCreateApiKey201Response instantiates a new CreateApiKey201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKey201Response() *CreateApiKey201Response {
	this := CreateApiKey201Response{}
	return &this
}

// NewCreateApiKey201ResponseWithDefaults instantiates a new CreateApiKey201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKey201ResponseWithDefaults() *CreateApiKey201Response {
	this := CreateApiKey201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateApiKey201Response) GetData() ApiKeyFullWithToken {
	if o == nil || IsNil(o.Data) {
		var ret ApiKeyFullWithToken
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKey201Response) GetDataOk() (*ApiKeyFullWithToken, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateApiKey201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ApiKeyFullWithToken and assigns it to the Data field.
func (o *CreateApiKey201Response) SetData(v ApiKeyFullWithToken) {
	o.Data = &v
}

func (o CreateApiKey201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiKey201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateApiKey201Response struct {
	value *CreateApiKey201Response
	isSet bool
}

func (v NullableCreateApiKey201Response) Get() *CreateApiKey201Response {
	return v.value
}

func (v *NullableCreateApiKey201Response) Set(val *CreateApiKey201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKey201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKey201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKey201Response(val *CreateApiKey201Response) *NullableCreateApiKey201Response {
	return &NullableCreateApiKey201Response{value: val, isSet: true}
}

func (v NullableCreateApiKey201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKey201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


