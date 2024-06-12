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

// checks if the MeasurementEventLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasurementEventLink{}

// MeasurementEventLink struct for MeasurementEventLink
type MeasurementEventLink struct {
	// The base URI of the link to an event
	Link *string `json:"link,omitempty"`
	// Describes any parameters that can be added to the event link
	Parameters *map[string]MeasurementEventLinkParameterInfo `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MeasurementEventLink MeasurementEventLink

// NewMeasurementEventLink instantiates a new MeasurementEventLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementEventLink() *MeasurementEventLink {
	this := MeasurementEventLink{}
	return &this
}

// NewMeasurementEventLinkWithDefaults instantiates a new MeasurementEventLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementEventLinkWithDefaults() *MeasurementEventLink {
	this := MeasurementEventLink{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MeasurementEventLink) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementEventLink) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MeasurementEventLink) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MeasurementEventLink) SetLink(v string) {
	o.Link = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *MeasurementEventLink) GetParameters() map[string]MeasurementEventLinkParameterInfo {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]MeasurementEventLinkParameterInfo
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementEventLink) GetParametersOk() (*map[string]MeasurementEventLinkParameterInfo, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *MeasurementEventLink) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]MeasurementEventLinkParameterInfo and assigns it to the Parameters field.
func (o *MeasurementEventLink) SetParameters(v map[string]MeasurementEventLinkParameterInfo) {
	o.Parameters = &v
}

func (o MeasurementEventLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasurementEventLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MeasurementEventLink) UnmarshalJSON(data []byte) (err error) {
	varMeasurementEventLink := _MeasurementEventLink{}

	err = json.Unmarshal(data, &varMeasurementEventLink)

	if err != nil {
		return err
	}

	*o = MeasurementEventLink(varMeasurementEventLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "link")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeasurementEventLink struct {
	value *MeasurementEventLink
	isSet bool
}

func (v NullableMeasurementEventLink) Get() *MeasurementEventLink {
	return v.value
}

func (v *NullableMeasurementEventLink) Set(val *MeasurementEventLink) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementEventLink) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementEventLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementEventLink(val *MeasurementEventLink) *NullableMeasurementEventLink {
	return &NullableMeasurementEventLink{value: val, isSet: true}
}

func (v NullableMeasurementEventLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementEventLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


