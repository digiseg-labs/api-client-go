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

// checks if the QueryCampaignTimingStats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryCampaignTimingStats200Response{}

// QueryCampaignTimingStats200Response struct for QueryCampaignTimingStats200Response
type QueryCampaignTimingStats200Response struct {
	Data *CampaignTimingStats `json:"data,omitempty"`
}

// NewQueryCampaignTimingStats200Response instantiates a new QueryCampaignTimingStats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCampaignTimingStats200Response() *QueryCampaignTimingStats200Response {
	this := QueryCampaignTimingStats200Response{}
	return &this
}

// NewQueryCampaignTimingStats200ResponseWithDefaults instantiates a new QueryCampaignTimingStats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCampaignTimingStats200ResponseWithDefaults() *QueryCampaignTimingStats200Response {
	this := QueryCampaignTimingStats200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryCampaignTimingStats200Response) GetData() CampaignTimingStats {
	if o == nil || IsNil(o.Data) {
		var ret CampaignTimingStats
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCampaignTimingStats200Response) GetDataOk() (*CampaignTimingStats, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryCampaignTimingStats200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CampaignTimingStats and assigns it to the Data field.
func (o *QueryCampaignTimingStats200Response) SetData(v CampaignTimingStats) {
	o.Data = &v
}

func (o QueryCampaignTimingStats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCampaignTimingStats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableQueryCampaignTimingStats200Response struct {
	value *QueryCampaignTimingStats200Response
	isSet bool
}

func (v NullableQueryCampaignTimingStats200Response) Get() *QueryCampaignTimingStats200Response {
	return v.value
}

func (v *NullableQueryCampaignTimingStats200Response) Set(val *QueryCampaignTimingStats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCampaignTimingStats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCampaignTimingStats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCampaignTimingStats200Response(val *QueryCampaignTimingStats200Response) *NullableQueryCampaignTimingStats200Response {
	return &NullableQueryCampaignTimingStats200Response{value: val, isSet: true}
}

func (v NullableQueryCampaignTimingStats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCampaignTimingStats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


