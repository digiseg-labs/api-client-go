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

// checks if the ListAudienceDataDailyUsage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudienceDataDailyUsage200Response{}

// ListAudienceDataDailyUsage200Response struct for ListAudienceDataDailyUsage200Response
type ListAudienceDataDailyUsage200Response struct {
	Data []AudienceDataDailyUsage `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAudienceDataDailyUsage200Response ListAudienceDataDailyUsage200Response

// NewListAudienceDataDailyUsage200Response instantiates a new ListAudienceDataDailyUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudienceDataDailyUsage200Response() *ListAudienceDataDailyUsage200Response {
	this := ListAudienceDataDailyUsage200Response{}
	return &this
}

// NewListAudienceDataDailyUsage200ResponseWithDefaults instantiates a new ListAudienceDataDailyUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudienceDataDailyUsage200ResponseWithDefaults() *ListAudienceDataDailyUsage200Response {
	this := ListAudienceDataDailyUsage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAudienceDataDailyUsage200Response) GetData() []AudienceDataDailyUsage {
	if o == nil || IsNil(o.Data) {
		var ret []AudienceDataDailyUsage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAudienceDataDailyUsage200Response) GetDataOk() ([]AudienceDataDailyUsage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAudienceDataDailyUsage200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AudienceDataDailyUsage and assigns it to the Data field.
func (o *ListAudienceDataDailyUsage200Response) SetData(v []AudienceDataDailyUsage) {
	o.Data = v
}

func (o ListAudienceDataDailyUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudienceDataDailyUsage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAudienceDataDailyUsage200Response) UnmarshalJSON(data []byte) (err error) {
	varListAudienceDataDailyUsage200Response := _ListAudienceDataDailyUsage200Response{}

	err = json.Unmarshal(data, &varListAudienceDataDailyUsage200Response)

	if err != nil {
		return err
	}

	*o = ListAudienceDataDailyUsage200Response(varListAudienceDataDailyUsage200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAudienceDataDailyUsage200Response struct {
	value *ListAudienceDataDailyUsage200Response
	isSet bool
}

func (v NullableListAudienceDataDailyUsage200Response) Get() *ListAudienceDataDailyUsage200Response {
	return v.value
}

func (v *NullableListAudienceDataDailyUsage200Response) Set(val *ListAudienceDataDailyUsage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudienceDataDailyUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudienceDataDailyUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudienceDataDailyUsage200Response(val *ListAudienceDataDailyUsage200Response) *NullableListAudienceDataDailyUsage200Response {
	return &NullableListAudienceDataDailyUsage200Response{value: val, isSet: true}
}

func (v NullableListAudienceDataDailyUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudienceDataDailyUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


