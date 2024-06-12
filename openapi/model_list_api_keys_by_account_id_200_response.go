/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young couples and singles | |  |  | c2 | Early family life | |  |  | c3 | Middle-aged families | |  |  | c4 | Mature families | |  |  | c5 | Pensioners / Retirees | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Small | |  |  | l2 | Medium | |  |  | l3 | Large | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListApiKeysByAccountId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApiKeysByAccountId200Response{}

// ListApiKeysByAccountId200Response struct for ListApiKeysByAccountId200Response
type ListApiKeysByAccountId200Response struct {
	Meta *ListPaginationMeta `json:"meta,omitempty"`
	Links *ListPaginationLinks `json:"links,omitempty"`
	Data []ApiKeyItem `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListApiKeysByAccountId200Response ListApiKeysByAccountId200Response

// NewListApiKeysByAccountId200Response instantiates a new ListApiKeysByAccountId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApiKeysByAccountId200Response() *ListApiKeysByAccountId200Response {
	this := ListApiKeysByAccountId200Response{}
	return &this
}

// NewListApiKeysByAccountId200ResponseWithDefaults instantiates a new ListApiKeysByAccountId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApiKeysByAccountId200ResponseWithDefaults() *ListApiKeysByAccountId200Response {
	this := ListApiKeysByAccountId200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListApiKeysByAccountId200Response) GetMeta() ListPaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApiKeysByAccountId200Response) GetMetaOk() (*ListPaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListApiKeysByAccountId200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListPaginationMeta and assigns it to the Meta field.
func (o *ListApiKeysByAccountId200Response) SetMeta(v ListPaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListApiKeysByAccountId200Response) GetLinks() ListPaginationLinks {
	if o == nil || IsNil(o.Links) {
		var ret ListPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApiKeysByAccountId200Response) GetLinksOk() (*ListPaginationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListApiKeysByAccountId200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListPaginationLinks and assigns it to the Links field.
func (o *ListApiKeysByAccountId200Response) SetLinks(v ListPaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListApiKeysByAccountId200Response) GetData() []ApiKeyItem {
	if o == nil || IsNil(o.Data) {
		var ret []ApiKeyItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApiKeysByAccountId200Response) GetDataOk() ([]ApiKeyItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListApiKeysByAccountId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApiKeyItem and assigns it to the Data field.
func (o *ListApiKeysByAccountId200Response) SetData(v []ApiKeyItem) {
	o.Data = v
}

func (o ListApiKeysByAccountId200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApiKeysByAccountId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListApiKeysByAccountId200Response) UnmarshalJSON(data []byte) (err error) {
	varListApiKeysByAccountId200Response := _ListApiKeysByAccountId200Response{}

	err = json.Unmarshal(data, &varListApiKeysByAccountId200Response)

	if err != nil {
		return err
	}

	*o = ListApiKeysByAccountId200Response(varListApiKeysByAccountId200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListApiKeysByAccountId200Response struct {
	value *ListApiKeysByAccountId200Response
	isSet bool
}

func (v NullableListApiKeysByAccountId200Response) Get() *ListApiKeysByAccountId200Response {
	return v.value
}

func (v *NullableListApiKeysByAccountId200Response) Set(val *ListApiKeysByAccountId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListApiKeysByAccountId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListApiKeysByAccountId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApiKeysByAccountId200Response(val *ListApiKeysByAccountId200Response) *NullableListApiKeysByAccountId200Response {
	return &NullableListApiKeysByAccountId200Response{value: val, isSet: true}
}

func (v NullableListApiKeysByAccountId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApiKeysByAccountId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


