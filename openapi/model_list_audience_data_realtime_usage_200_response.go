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

// checks if the ListAudienceDataRealtimeUsage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudienceDataRealtimeUsage200Response{}

// ListAudienceDataRealtimeUsage200Response struct for ListAudienceDataRealtimeUsage200Response
type ListAudienceDataRealtimeUsage200Response struct {
	Data []AudienceDataRealtimeUsage `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAudienceDataRealtimeUsage200Response ListAudienceDataRealtimeUsage200Response

// NewListAudienceDataRealtimeUsage200Response instantiates a new ListAudienceDataRealtimeUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudienceDataRealtimeUsage200Response() *ListAudienceDataRealtimeUsage200Response {
	this := ListAudienceDataRealtimeUsage200Response{}
	return &this
}

// NewListAudienceDataRealtimeUsage200ResponseWithDefaults instantiates a new ListAudienceDataRealtimeUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudienceDataRealtimeUsage200ResponseWithDefaults() *ListAudienceDataRealtimeUsage200Response {
	this := ListAudienceDataRealtimeUsage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAudienceDataRealtimeUsage200Response) GetData() []AudienceDataRealtimeUsage {
	if o == nil || IsNil(o.Data) {
		var ret []AudienceDataRealtimeUsage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAudienceDataRealtimeUsage200Response) GetDataOk() ([]AudienceDataRealtimeUsage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAudienceDataRealtimeUsage200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AudienceDataRealtimeUsage and assigns it to the Data field.
func (o *ListAudienceDataRealtimeUsage200Response) SetData(v []AudienceDataRealtimeUsage) {
	o.Data = v
}

func (o ListAudienceDataRealtimeUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudienceDataRealtimeUsage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAudienceDataRealtimeUsage200Response) UnmarshalJSON(data []byte) (err error) {
	varListAudienceDataRealtimeUsage200Response := _ListAudienceDataRealtimeUsage200Response{}

	err = json.Unmarshal(data, &varListAudienceDataRealtimeUsage200Response)

	if err != nil {
		return err
	}

	*o = ListAudienceDataRealtimeUsage200Response(varListAudienceDataRealtimeUsage200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAudienceDataRealtimeUsage200Response struct {
	value *ListAudienceDataRealtimeUsage200Response
	isSet bool
}

func (v NullableListAudienceDataRealtimeUsage200Response) Get() *ListAudienceDataRealtimeUsage200Response {
	return v.value
}

func (v *NullableListAudienceDataRealtimeUsage200Response) Set(val *ListAudienceDataRealtimeUsage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudienceDataRealtimeUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudienceDataRealtimeUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudienceDataRealtimeUsage200Response(val *ListAudienceDataRealtimeUsage200Response) *NullableListAudienceDataRealtimeUsage200Response {
	return &NullableListAudienceDataRealtimeUsage200Response{value: val, isSet: true}
}

func (v NullableListAudienceDataRealtimeUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudienceDataRealtimeUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


