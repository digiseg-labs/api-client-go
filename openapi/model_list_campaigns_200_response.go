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

// checks if the ListCampaigns200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCampaigns200Response{}

// ListCampaigns200Response struct for ListCampaigns200Response
type ListCampaigns200Response struct {
	Meta *ListPaginationMeta `json:"meta,omitempty"`
	Data []CampaignItem `json:"data,omitempty"`
	Links *ListPaginationLinks `json:"links,omitempty"`
}

// NewListCampaigns200Response instantiates a new ListCampaigns200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCampaigns200Response() *ListCampaigns200Response {
	this := ListCampaigns200Response{}
	return &this
}

// NewListCampaigns200ResponseWithDefaults instantiates a new ListCampaigns200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCampaigns200ResponseWithDefaults() *ListCampaigns200Response {
	this := ListCampaigns200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListCampaigns200Response) GetMeta() ListPaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCampaigns200Response) GetMetaOk() (*ListPaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListCampaigns200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListPaginationMeta and assigns it to the Meta field.
func (o *ListCampaigns200Response) SetMeta(v ListPaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListCampaigns200Response) GetData() []CampaignItem {
	if o == nil || IsNil(o.Data) {
		var ret []CampaignItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCampaigns200Response) GetDataOk() ([]CampaignItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCampaigns200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CampaignItem and assigns it to the Data field.
func (o *ListCampaigns200Response) SetData(v []CampaignItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListCampaigns200Response) GetLinks() ListPaginationLinks {
	if o == nil || IsNil(o.Links) {
		var ret ListPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCampaigns200Response) GetLinksOk() (*ListPaginationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListCampaigns200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListPaginationLinks and assigns it to the Links field.
func (o *ListCampaigns200Response) SetLinks(v ListPaginationLinks) {
	o.Links = &v
}

func (o ListCampaigns200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCampaigns200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableListCampaigns200Response struct {
	value *ListCampaigns200Response
	isSet bool
}

func (v NullableListCampaigns200Response) Get() *ListCampaigns200Response {
	return v.value
}

func (v *NullableListCampaigns200Response) Set(val *ListCampaigns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListCampaigns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListCampaigns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCampaigns200Response(val *ListCampaigns200Response) *NullableListCampaigns200Response {
	return &NullableListCampaigns200Response{value: val, isSet: true}
}

func (v NullableListCampaigns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCampaigns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

