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

// checks if the PopulationAudienceCategorySetBusinessSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationAudienceCategorySetBusinessSection{}

// PopulationAudienceCategorySetBusinessSection struct for PopulationAudienceCategorySetBusinessSection
type PopulationAudienceCategorySetBusinessSection struct {
	// The fraction of events that fall within this object compared to the total of the category or segment (usually represented by the measurement's parent's parent). For example, if the measurement is \"impression\" on the `home_type` \"Apartment\" object, then the `fraction_of_total` represents the number of impressions on apartments compared to impressions from other `home_type` values. 
	FractionOfTotal *float64 `json:"fraction_of_total,omitempty"`
	// An object with category codes as keys, objects with audience codes and fractions of totals as keys.
	AudienceCategories *map[string]map[string]float64 `json:"audience_categories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopulationAudienceCategorySetBusinessSection PopulationAudienceCategorySetBusinessSection

// NewPopulationAudienceCategorySetBusinessSection instantiates a new PopulationAudienceCategorySetBusinessSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationAudienceCategorySetBusinessSection() *PopulationAudienceCategorySetBusinessSection {
	this := PopulationAudienceCategorySetBusinessSection{}
	return &this
}

// NewPopulationAudienceCategorySetBusinessSectionWithDefaults instantiates a new PopulationAudienceCategorySetBusinessSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationAudienceCategorySetBusinessSectionWithDefaults() *PopulationAudienceCategorySetBusinessSection {
	this := PopulationAudienceCategorySetBusinessSection{}
	return &this
}

// GetFractionOfTotal returns the FractionOfTotal field value if set, zero value otherwise.
func (o *PopulationAudienceCategorySetBusinessSection) GetFractionOfTotal() float64 {
	if o == nil || IsNil(o.FractionOfTotal) {
		var ret float64
		return ret
	}
	return *o.FractionOfTotal
}

// GetFractionOfTotalOk returns a tuple with the FractionOfTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationAudienceCategorySetBusinessSection) GetFractionOfTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.FractionOfTotal) {
		return nil, false
	}
	return o.FractionOfTotal, true
}

// HasFractionOfTotal returns a boolean if a field has been set.
func (o *PopulationAudienceCategorySetBusinessSection) HasFractionOfTotal() bool {
	if o != nil && !IsNil(o.FractionOfTotal) {
		return true
	}

	return false
}

// SetFractionOfTotal gets a reference to the given float64 and assigns it to the FractionOfTotal field.
func (o *PopulationAudienceCategorySetBusinessSection) SetFractionOfTotal(v float64) {
	o.FractionOfTotal = &v
}

// GetAudienceCategories returns the AudienceCategories field value if set, zero value otherwise.
func (o *PopulationAudienceCategorySetBusinessSection) GetAudienceCategories() map[string]map[string]float64 {
	if o == nil || IsNil(o.AudienceCategories) {
		var ret map[string]map[string]float64
		return ret
	}
	return *o.AudienceCategories
}

// GetAudienceCategoriesOk returns a tuple with the AudienceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationAudienceCategorySetBusinessSection) GetAudienceCategoriesOk() (*map[string]map[string]float64, bool) {
	if o == nil || IsNil(o.AudienceCategories) {
		return nil, false
	}
	return o.AudienceCategories, true
}

// HasAudienceCategories returns a boolean if a field has been set.
func (o *PopulationAudienceCategorySetBusinessSection) HasAudienceCategories() bool {
	if o != nil && !IsNil(o.AudienceCategories) {
		return true
	}

	return false
}

// SetAudienceCategories gets a reference to the given map[string]map[string]float64 and assigns it to the AudienceCategories field.
func (o *PopulationAudienceCategorySetBusinessSection) SetAudienceCategories(v map[string]map[string]float64) {
	o.AudienceCategories = &v
}

func (o PopulationAudienceCategorySetBusinessSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationAudienceCategorySetBusinessSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FractionOfTotal) {
		toSerialize["fraction_of_total"] = o.FractionOfTotal
	}
	if !IsNil(o.AudienceCategories) {
		toSerialize["audience_categories"] = o.AudienceCategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PopulationAudienceCategorySetBusinessSection) UnmarshalJSON(data []byte) (err error) {
	varPopulationAudienceCategorySetBusinessSection := _PopulationAudienceCategorySetBusinessSection{}

	err = json.Unmarshal(data, &varPopulationAudienceCategorySetBusinessSection)

	if err != nil {
		return err
	}

	*o = PopulationAudienceCategorySetBusinessSection(varPopulationAudienceCategorySetBusinessSection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fraction_of_total")
		delete(additionalProperties, "audience_categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopulationAudienceCategorySetBusinessSection struct {
	value *PopulationAudienceCategorySetBusinessSection
	isSet bool
}

func (v NullablePopulationAudienceCategorySetBusinessSection) Get() *PopulationAudienceCategorySetBusinessSection {
	return v.value
}

func (v *NullablePopulationAudienceCategorySetBusinessSection) Set(val *PopulationAudienceCategorySetBusinessSection) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationAudienceCategorySetBusinessSection) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationAudienceCategorySetBusinessSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationAudienceCategorySetBusinessSection(val *PopulationAudienceCategorySetBusinessSection) *NullablePopulationAudienceCategorySetBusinessSection {
	return &NullablePopulationAudienceCategorySetBusinessSection{value: val, isSet: true}
}

func (v NullablePopulationAudienceCategorySetBusinessSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationAudienceCategorySetBusinessSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


