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
)

// checks if the PopulationSourceBusinessSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationSourceBusinessSection{}

// PopulationSourceBusinessSection struct for PopulationSourceBusinessSection
type PopulationSourceBusinessSection struct {
	AudienceCategories *PopulationSourceBusinessCategorySet `json:"audience_categories,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewPopulationSourceBusinessSection instantiates a new PopulationSourceBusinessSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationSourceBusinessSection() *PopulationSourceBusinessSection {
	this := PopulationSourceBusinessSection{}
	return &this
}

// NewPopulationSourceBusinessSectionWithDefaults instantiates a new PopulationSourceBusinessSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationSourceBusinessSectionWithDefaults() *PopulationSourceBusinessSection {
	this := PopulationSourceBusinessSection{}
	return &this
}

// GetAudienceCategories returns the AudienceCategories field value if set, zero value otherwise.
func (o *PopulationSourceBusinessSection) GetAudienceCategories() PopulationSourceBusinessCategorySet {
	if o == nil || IsNil(o.AudienceCategories) {
		var ret PopulationSourceBusinessCategorySet
		return ret
	}
	return *o.AudienceCategories
}

// GetAudienceCategoriesOk returns a tuple with the AudienceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationSourceBusinessSection) GetAudienceCategoriesOk() (*PopulationSourceBusinessCategorySet, bool) {
	if o == nil || IsNil(o.AudienceCategories) {
		return nil, false
	}
	return o.AudienceCategories, true
}

// HasAudienceCategories returns a boolean if a field has been set.
func (o *PopulationSourceBusinessSection) HasAudienceCategories() bool {
	if o != nil && !IsNil(o.AudienceCategories) {
		return true
	}

	return false
}

// SetAudienceCategories gets a reference to the given PopulationSourceBusinessCategorySet and assigns it to the AudienceCategories field.
func (o *PopulationSourceBusinessSection) SetAudienceCategories(v PopulationSourceBusinessCategorySet) {
	o.AudienceCategories = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PopulationSourceBusinessSection) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationSourceBusinessSection) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PopulationSourceBusinessSection) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PopulationSourceBusinessSection) SetCount(v int32) {
	o.Count = &v
}

func (o PopulationSourceBusinessSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationSourceBusinessSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudienceCategories) {
		toSerialize["audience_categories"] = o.AudienceCategories
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullablePopulationSourceBusinessSection struct {
	value *PopulationSourceBusinessSection
	isSet bool
}

func (v NullablePopulationSourceBusinessSection) Get() *PopulationSourceBusinessSection {
	return v.value
}

func (v *NullablePopulationSourceBusinessSection) Set(val *PopulationSourceBusinessSection) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationSourceBusinessSection) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationSourceBusinessSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationSourceBusinessSection(val *PopulationSourceBusinessSection) *NullablePopulationSourceBusinessSection {
	return &NullablePopulationSourceBusinessSection{value: val, isSet: true}
}

func (v NullablePopulationSourceBusinessSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationSourceBusinessSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

