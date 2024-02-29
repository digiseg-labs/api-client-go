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

// checks if the PopulationSourcePrivateCategorySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationSourcePrivateCategorySet{}

// PopulationSourcePrivateCategorySet struct for PopulationSourcePrivateCategorySet
type PopulationSourcePrivateCategorySet struct {
	HomeType PopulationSourcePrivateCategorySetHomeType `json:"home_type"`
	Savings PopulationSourcePrivateCategorySetSavings `json:"savings"`
	Lifecycle PopulationSourcePrivateCategorySetLifecycle `json:"lifecycle"`
	Cars PopulationSourcePrivateCategorySetCars `json:"cars"`
	Children PopulationSourcePrivateCategorySetChildren `json:"children"`
	Education PopulationSourcePrivateCategorySetEducation `json:"education"`
	NeighbourhoodType PopulationSourcePrivateCategorySetNeighbourhoodType `json:"neighbourhood_type"`
	Income PopulationSourcePrivateCategorySetIncome `json:"income"`
	HomeOwnership PopulationSourcePrivateCategorySetHomeOwnership `json:"home_ownership"`
	BuildingAge PopulationSourcePrivateCategorySetBuildingAge `json:"building_age"`
	LivingSpace PopulationSourcePrivateCategorySetLivingSpace `json:"living_space"`
	TechLevel PopulationSourcePrivateCategorySetTechLevel `json:"tech_level"`
}

type _PopulationSourcePrivateCategorySet PopulationSourcePrivateCategorySet

// NewPopulationSourcePrivateCategorySet instantiates a new PopulationSourcePrivateCategorySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationSourcePrivateCategorySet(homeType PopulationSourcePrivateCategorySetHomeType, savings PopulationSourcePrivateCategorySetSavings, lifecycle PopulationSourcePrivateCategorySetLifecycle, cars PopulationSourcePrivateCategorySetCars, children PopulationSourcePrivateCategorySetChildren, education PopulationSourcePrivateCategorySetEducation, neighbourhoodType PopulationSourcePrivateCategorySetNeighbourhoodType, income PopulationSourcePrivateCategorySetIncome, homeOwnership PopulationSourcePrivateCategorySetHomeOwnership, buildingAge PopulationSourcePrivateCategorySetBuildingAge, livingSpace PopulationSourcePrivateCategorySetLivingSpace, techLevel PopulationSourcePrivateCategorySetTechLevel) *PopulationSourcePrivateCategorySet {
	this := PopulationSourcePrivateCategorySet{}
	this.HomeType = homeType
	this.Savings = savings
	this.Lifecycle = lifecycle
	this.Cars = cars
	this.Children = children
	this.Education = education
	this.NeighbourhoodType = neighbourhoodType
	this.Income = income
	this.HomeOwnership = homeOwnership
	this.BuildingAge = buildingAge
	this.LivingSpace = livingSpace
	this.TechLevel = techLevel
	return &this
}

// NewPopulationSourcePrivateCategorySetWithDefaults instantiates a new PopulationSourcePrivateCategorySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationSourcePrivateCategorySetWithDefaults() *PopulationSourcePrivateCategorySet {
	this := PopulationSourcePrivateCategorySet{}
	return &this
}

// GetHomeType returns the HomeType field value
func (o *PopulationSourcePrivateCategorySet) GetHomeType() PopulationSourcePrivateCategorySetHomeType {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetHomeType
		return ret
	}

	return o.HomeType
}

// GetHomeTypeOk returns a tuple with the HomeType field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetHomeTypeOk() (*PopulationSourcePrivateCategorySetHomeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeType, true
}

// SetHomeType sets field value
func (o *PopulationSourcePrivateCategorySet) SetHomeType(v PopulationSourcePrivateCategorySetHomeType) {
	o.HomeType = v
}

// GetSavings returns the Savings field value
func (o *PopulationSourcePrivateCategorySet) GetSavings() PopulationSourcePrivateCategorySetSavings {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetSavings
		return ret
	}

	return o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetSavingsOk() (*PopulationSourcePrivateCategorySetSavings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Savings, true
}

// SetSavings sets field value
func (o *PopulationSourcePrivateCategorySet) SetSavings(v PopulationSourcePrivateCategorySetSavings) {
	o.Savings = v
}

// GetLifecycle returns the Lifecycle field value
func (o *PopulationSourcePrivateCategorySet) GetLifecycle() PopulationSourcePrivateCategorySetLifecycle {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetLifecycle
		return ret
	}

	return o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetLifecycleOk() (*PopulationSourcePrivateCategorySetLifecycle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lifecycle, true
}

// SetLifecycle sets field value
func (o *PopulationSourcePrivateCategorySet) SetLifecycle(v PopulationSourcePrivateCategorySetLifecycle) {
	o.Lifecycle = v
}

// GetCars returns the Cars field value
func (o *PopulationSourcePrivateCategorySet) GetCars() PopulationSourcePrivateCategorySetCars {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetCars
		return ret
	}

	return o.Cars
}

// GetCarsOk returns a tuple with the Cars field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetCarsOk() (*PopulationSourcePrivateCategorySetCars, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cars, true
}

// SetCars sets field value
func (o *PopulationSourcePrivateCategorySet) SetCars(v PopulationSourcePrivateCategorySetCars) {
	o.Cars = v
}

// GetChildren returns the Children field value
func (o *PopulationSourcePrivateCategorySet) GetChildren() PopulationSourcePrivateCategorySetChildren {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetChildren
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetChildrenOk() (*PopulationSourcePrivateCategorySetChildren, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *PopulationSourcePrivateCategorySet) SetChildren(v PopulationSourcePrivateCategorySetChildren) {
	o.Children = v
}

// GetEducation returns the Education field value
func (o *PopulationSourcePrivateCategorySet) GetEducation() PopulationSourcePrivateCategorySetEducation {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetEducation
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetEducationOk() (*PopulationSourcePrivateCategorySetEducation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *PopulationSourcePrivateCategorySet) SetEducation(v PopulationSourcePrivateCategorySetEducation) {
	o.Education = v
}

// GetNeighbourhoodType returns the NeighbourhoodType field value
func (o *PopulationSourcePrivateCategorySet) GetNeighbourhoodType() PopulationSourcePrivateCategorySetNeighbourhoodType {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetNeighbourhoodType
		return ret
	}

	return o.NeighbourhoodType
}

// GetNeighbourhoodTypeOk returns a tuple with the NeighbourhoodType field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetNeighbourhoodTypeOk() (*PopulationSourcePrivateCategorySetNeighbourhoodType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeighbourhoodType, true
}

// SetNeighbourhoodType sets field value
func (o *PopulationSourcePrivateCategorySet) SetNeighbourhoodType(v PopulationSourcePrivateCategorySetNeighbourhoodType) {
	o.NeighbourhoodType = v
}

// GetIncome returns the Income field value
func (o *PopulationSourcePrivateCategorySet) GetIncome() PopulationSourcePrivateCategorySetIncome {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetIncome
		return ret
	}

	return o.Income
}

// GetIncomeOk returns a tuple with the Income field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetIncomeOk() (*PopulationSourcePrivateCategorySetIncome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Income, true
}

// SetIncome sets field value
func (o *PopulationSourcePrivateCategorySet) SetIncome(v PopulationSourcePrivateCategorySetIncome) {
	o.Income = v
}

// GetHomeOwnership returns the HomeOwnership field value
func (o *PopulationSourcePrivateCategorySet) GetHomeOwnership() PopulationSourcePrivateCategorySetHomeOwnership {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetHomeOwnership
		return ret
	}

	return o.HomeOwnership
}

// GetHomeOwnershipOk returns a tuple with the HomeOwnership field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetHomeOwnershipOk() (*PopulationSourcePrivateCategorySetHomeOwnership, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeOwnership, true
}

// SetHomeOwnership sets field value
func (o *PopulationSourcePrivateCategorySet) SetHomeOwnership(v PopulationSourcePrivateCategorySetHomeOwnership) {
	o.HomeOwnership = v
}

// GetBuildingAge returns the BuildingAge field value
func (o *PopulationSourcePrivateCategorySet) GetBuildingAge() PopulationSourcePrivateCategorySetBuildingAge {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetBuildingAge
		return ret
	}

	return o.BuildingAge
}

// GetBuildingAgeOk returns a tuple with the BuildingAge field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetBuildingAgeOk() (*PopulationSourcePrivateCategorySetBuildingAge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildingAge, true
}

// SetBuildingAge sets field value
func (o *PopulationSourcePrivateCategorySet) SetBuildingAge(v PopulationSourcePrivateCategorySetBuildingAge) {
	o.BuildingAge = v
}

// GetLivingSpace returns the LivingSpace field value
func (o *PopulationSourcePrivateCategorySet) GetLivingSpace() PopulationSourcePrivateCategorySetLivingSpace {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetLivingSpace
		return ret
	}

	return o.LivingSpace
}

// GetLivingSpaceOk returns a tuple with the LivingSpace field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetLivingSpaceOk() (*PopulationSourcePrivateCategorySetLivingSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LivingSpace, true
}

// SetLivingSpace sets field value
func (o *PopulationSourcePrivateCategorySet) SetLivingSpace(v PopulationSourcePrivateCategorySetLivingSpace) {
	o.LivingSpace = v
}

// GetTechLevel returns the TechLevel field value
func (o *PopulationSourcePrivateCategorySet) GetTechLevel() PopulationSourcePrivateCategorySetTechLevel {
	if o == nil {
		var ret PopulationSourcePrivateCategorySetTechLevel
		return ret
	}

	return o.TechLevel
}

// GetTechLevelOk returns a tuple with the TechLevel field value
// and a boolean to check if the value has been set.
func (o *PopulationSourcePrivateCategorySet) GetTechLevelOk() (*PopulationSourcePrivateCategorySetTechLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechLevel, true
}

// SetTechLevel sets field value
func (o *PopulationSourcePrivateCategorySet) SetTechLevel(v PopulationSourcePrivateCategorySetTechLevel) {
	o.TechLevel = v
}

func (o PopulationSourcePrivateCategorySet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationSourcePrivateCategorySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["home_type"] = o.HomeType
	toSerialize["savings"] = o.Savings
	toSerialize["lifecycle"] = o.Lifecycle
	toSerialize["cars"] = o.Cars
	toSerialize["children"] = o.Children
	toSerialize["education"] = o.Education
	toSerialize["neighbourhood_type"] = o.NeighbourhoodType
	toSerialize["income"] = o.Income
	toSerialize["home_ownership"] = o.HomeOwnership
	toSerialize["building_age"] = o.BuildingAge
	toSerialize["living_space"] = o.LivingSpace
	toSerialize["tech_level"] = o.TechLevel
	return toSerialize, nil
}

func (o *PopulationSourcePrivateCategorySet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"home_type",
		"savings",
		"lifecycle",
		"cars",
		"children",
		"education",
		"neighbourhood_type",
		"income",
		"home_ownership",
		"building_age",
		"living_space",
		"tech_level",
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

	varPopulationSourcePrivateCategorySet := _PopulationSourcePrivateCategorySet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPopulationSourcePrivateCategorySet)

	if err != nil {
		return err
	}

	*o = PopulationSourcePrivateCategorySet(varPopulationSourcePrivateCategorySet)

	return err
}

type NullablePopulationSourcePrivateCategorySet struct {
	value *PopulationSourcePrivateCategorySet
	isSet bool
}

func (v NullablePopulationSourcePrivateCategorySet) Get() *PopulationSourcePrivateCategorySet {
	return v.value
}

func (v *NullablePopulationSourcePrivateCategorySet) Set(val *PopulationSourcePrivateCategorySet) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationSourcePrivateCategorySet) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationSourcePrivateCategorySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationSourcePrivateCategorySet(val *PopulationSourcePrivateCategorySet) *NullablePopulationSourcePrivateCategorySet {
	return &NullablePopulationSourcePrivateCategorySet{value: val, isSet: true}
}

func (v NullablePopulationSourcePrivateCategorySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationSourcePrivateCategorySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


