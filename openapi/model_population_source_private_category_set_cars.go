/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young singles and couples | |  |  | c2 | Young couples with children | |  |  | c3 | Families with school children | |  |  | c4 | Older families | |  |  | c5 | Pensioners | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Up to 80 m² | |  |  | l2 | 80-119 m² | |  |  | l3 | Above 120 m² | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PopulationSourcePrivateCategorySetCars type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationSourcePrivateCategorySetCars{}

// PopulationSourcePrivateCategorySetCars struct for PopulationSourcePrivateCategorySetCars
type PopulationSourcePrivateCategorySetCars struct {
	D1 int32 `json:"d1"`
	D2 int32 `json:"d2"`
	D3 int32 `json:"d3"`
}

type _PopulationSourcePrivateCategorySetCars PopulationSourcePrivateCategorySetCars

// NewPopulationSourcePrivateCategorySetCars instantiates a new PopulationSourcePrivateCategorySetCars object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationSourcePrivateCategorySetCars(d1 int32, d2 int32, d3 int32) *PopulationSourcePrivateCategorySetCars {
	this := PopulationSourcePrivateCategorySetCars{}
	this.D1 = d1
	this.D2 = d2
	this.D3 = d3
	return &this
}

// NewPopulationSourcePrivateCategorySetCarsWithDefaults instantiates a new PopulationSourcePrivateCategorySetCars object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationSourcePrivateCategorySetCarsWithDefaults() *PopulationSourcePrivateCategorySetCars {
	this := PopulationSourcePrivateCategorySetCars{}
	return &this
}

// GetD1 returns the D1 field value
func (o *PopulationSourcePrivateCategorySetCars) GetD1() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.D1
}

// GetD1Ok returns a tuple with the D1 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetCars) GetD1Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.D1, true
}

// SetD1 sets field value
func (o *PopulationSourcePrivateCategorySetCars) SetD1(v int32) {
	o.D1 = v
}

// GetD2 returns the D2 field value
func (o *PopulationSourcePrivateCategorySetCars) GetD2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.D2
}

// GetD2Ok returns a tuple with the D2 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetCars) GetD2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.D2, true
}

// SetD2 sets field value
func (o *PopulationSourcePrivateCategorySetCars) SetD2(v int32) {
	o.D2 = v
}

// GetD3 returns the D3 field value
func (o *PopulationSourcePrivateCategorySetCars) GetD3() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.D3
}

// GetD3Ok returns a tuple with the D3 field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySetCars) GetD3Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.D3, true
}

// SetD3 sets field value
func (o *PopulationSourcePrivateCategorySetCars) SetD3(v int32) {
	o.D3 = v
}

func (o PopulationSourcePrivateCategorySetCars) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationSourcePrivateCategorySetCars) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["d1"] = o.D1
	toSerialize["d2"] = o.D2
	toSerialize["d3"] = o.D3
	return toSerialize, nil
}

func (o *PopulationSourcePrivateCategorySetCars) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"d1",
		"d2",
		"d3",
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

	varPopulationSourcePrivateCategorySetCars := _PopulationSourcePrivateCategorySetCars{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPopulationSourcePrivateCategorySetCars)

	if err != nil {
		return err
	}

	*o = PopulationSourcePrivateCategorySetCars(varPopulationSourcePrivateCategorySetCars)

	return err
}

type NullablePopulationSourcePrivateCategorySetCars struct {
	value *PopulationSourcePrivateCategorySetCars
	isSet bool
}

func (v NullablePopulationSourcePrivateCategorySetCars) Get() *PopulationSourcePrivateCategorySetCars {
	return v.value
}

func (v *NullablePopulationSourcePrivateCategorySetCars) Set(val *PopulationSourcePrivateCategorySetCars) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationSourcePrivateCategorySetCars) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationSourcePrivateCategorySetCars) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationSourcePrivateCategorySetCars(val *PopulationSourcePrivateCategorySetCars) *NullablePopulationSourcePrivateCategorySetCars {
	return &NullablePopulationSourcePrivateCategorySetCars{value: val, isSet: true}
}

func (v NullablePopulationSourcePrivateCategorySetCars) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationSourcePrivateCategorySetCars) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


