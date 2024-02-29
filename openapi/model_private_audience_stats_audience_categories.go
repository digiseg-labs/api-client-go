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

// checks if the PrivateAudienceStatsAudienceCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateAudienceStatsAudienceCategories{}

// PrivateAudienceStatsAudienceCategories struct for PrivateAudienceStatsAudienceCategories
type PrivateAudienceStatsAudienceCategories struct {
	HomeType AudienceCategoryStats `json:"home_type"`
	Savings AudienceCategoryStats `json:"savings"`
	Lifecycle AudienceCategoryStats `json:"lifecycle"`
	Cars AudienceCategoryStats `json:"cars"`
	Children AudienceCategoryStats `json:"children"`
	Education AudienceCategoryStats `json:"education"`
	NeighbourhoodType AudienceCategoryStats `json:"neighbourhood_type"`
	Income AudienceCategoryStats `json:"income"`
	HomeOwnership AudienceCategoryStats `json:"home_ownership"`
	BuildingAge AudienceCategoryStats `json:"building_age"`
	LivingSpace AudienceCategoryStats `json:"living_space"`
	TechLevel AudienceCategoryStats `json:"tech_level"`
}

type _PrivateAudienceStatsAudienceCategories PrivateAudienceStatsAudienceCategories

// NewPrivateAudienceStatsAudienceCategories instantiates a new PrivateAudienceStatsAudienceCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateAudienceStatsAudienceCategories(homeType AudienceCategoryStats, savings AudienceCategoryStats, lifecycle AudienceCategoryStats, cars AudienceCategoryStats, children AudienceCategoryStats, education AudienceCategoryStats, neighbourhoodType AudienceCategoryStats, income AudienceCategoryStats, homeOwnership AudienceCategoryStats, buildingAge AudienceCategoryStats, livingSpace AudienceCategoryStats, techLevel AudienceCategoryStats) *PrivateAudienceStatsAudienceCategories {
	this := PrivateAudienceStatsAudienceCategories{}
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

// NewPrivateAudienceStatsAudienceCategoriesWithDefaults instantiates a new PrivateAudienceStatsAudienceCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateAudienceStatsAudienceCategoriesWithDefaults() *PrivateAudienceStatsAudienceCategories {
	this := PrivateAudienceStatsAudienceCategories{}
	return &this
}

// GetHomeType returns the HomeType field value
func (o *PrivateAudienceStatsAudienceCategories) GetHomeType() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.HomeType
}

// GetHomeTypeOk returns a tuple with the HomeType field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetHomeTypeOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeType, true
}

// SetHomeType sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetHomeType(v AudienceCategoryStats) {
	o.HomeType = v
}

// GetSavings returns the Savings field value
func (o *PrivateAudienceStatsAudienceCategories) GetSavings() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetSavingsOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Savings, true
}

// SetSavings sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetSavings(v AudienceCategoryStats) {
	o.Savings = v
}

// GetLifecycle returns the Lifecycle field value
func (o *PrivateAudienceStatsAudienceCategories) GetLifecycle() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetLifecycleOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lifecycle, true
}

// SetLifecycle sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetLifecycle(v AudienceCategoryStats) {
	o.Lifecycle = v
}

// GetCars returns the Cars field value
func (o *PrivateAudienceStatsAudienceCategories) GetCars() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Cars
}

// GetCarsOk returns a tuple with the Cars field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetCarsOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cars, true
}

// SetCars sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetCars(v AudienceCategoryStats) {
	o.Cars = v
}

// GetChildren returns the Children field value
func (o *PrivateAudienceStatsAudienceCategories) GetChildren() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetChildrenOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetChildren(v AudienceCategoryStats) {
	o.Children = v
}

// GetEducation returns the Education field value
func (o *PrivateAudienceStatsAudienceCategories) GetEducation() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetEducationOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetEducation(v AudienceCategoryStats) {
	o.Education = v
}

// GetNeighbourhoodType returns the NeighbourhoodType field value
func (o *PrivateAudienceStatsAudienceCategories) GetNeighbourhoodType() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.NeighbourhoodType
}

// GetNeighbourhoodTypeOk returns a tuple with the NeighbourhoodType field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetNeighbourhoodTypeOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeighbourhoodType, true
}

// SetNeighbourhoodType sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetNeighbourhoodType(v AudienceCategoryStats) {
	o.NeighbourhoodType = v
}

// GetIncome returns the Income field value
func (o *PrivateAudienceStatsAudienceCategories) GetIncome() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.Income
}

// GetIncomeOk returns a tuple with the Income field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetIncomeOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Income, true
}

// SetIncome sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetIncome(v AudienceCategoryStats) {
	o.Income = v
}

// GetHomeOwnership returns the HomeOwnership field value
func (o *PrivateAudienceStatsAudienceCategories) GetHomeOwnership() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.HomeOwnership
}

// GetHomeOwnershipOk returns a tuple with the HomeOwnership field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetHomeOwnershipOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeOwnership, true
}

// SetHomeOwnership sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetHomeOwnership(v AudienceCategoryStats) {
	o.HomeOwnership = v
}

// GetBuildingAge returns the BuildingAge field value
func (o *PrivateAudienceStatsAudienceCategories) GetBuildingAge() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.BuildingAge
}

// GetBuildingAgeOk returns a tuple with the BuildingAge field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetBuildingAgeOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildingAge, true
}

// SetBuildingAge sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetBuildingAge(v AudienceCategoryStats) {
	o.BuildingAge = v
}

// GetLivingSpace returns the LivingSpace field value
func (o *PrivateAudienceStatsAudienceCategories) GetLivingSpace() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.LivingSpace
}

// GetLivingSpaceOk returns a tuple with the LivingSpace field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetLivingSpaceOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LivingSpace, true
}

// SetLivingSpace sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetLivingSpace(v AudienceCategoryStats) {
	o.LivingSpace = v
}

// GetTechLevel returns the TechLevel field value
func (o *PrivateAudienceStatsAudienceCategories) GetTechLevel() AudienceCategoryStats {
	if o == nil {
		var ret AudienceCategoryStats
		return ret
	}

	return o.TechLevel
}

// GetTechLevelOk returns a tuple with the TechLevel field value
// and a boolean to check if the value has been set.
func (o *PrivateAudienceStatsAudienceCategories) GetTechLevelOk() (*AudienceCategoryStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechLevel, true
}

// SetTechLevel sets field value
func (o *PrivateAudienceStatsAudienceCategories) SetTechLevel(v AudienceCategoryStats) {
	o.TechLevel = v
}

func (o PrivateAudienceStatsAudienceCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateAudienceStatsAudienceCategories) ToMap() (map[string]interface{}, error) {
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

func (o *PrivateAudienceStatsAudienceCategories) UnmarshalJSON(data []byte) (err error) {
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

	varPrivateAudienceStatsAudienceCategories := _PrivateAudienceStatsAudienceCategories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateAudienceStatsAudienceCategories)

	if err != nil {
		return err
	}

	*o = PrivateAudienceStatsAudienceCategories(varPrivateAudienceStatsAudienceCategories)

	return err
}

type NullablePrivateAudienceStatsAudienceCategories struct {
	value *PrivateAudienceStatsAudienceCategories
	isSet bool
}

func (v NullablePrivateAudienceStatsAudienceCategories) Get() *PrivateAudienceStatsAudienceCategories {
	return v.value
}

func (v *NullablePrivateAudienceStatsAudienceCategories) Set(val *PrivateAudienceStatsAudienceCategories) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateAudienceStatsAudienceCategories) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateAudienceStatsAudienceCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateAudienceStatsAudienceCategories(val *PrivateAudienceStatsAudienceCategories) *NullablePrivateAudienceStatsAudienceCategories {
	return &NullablePrivateAudienceStatsAudienceCategories{value: val, isSet: true}
}

func (v NullablePrivateAudienceStatsAudienceCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateAudienceStatsAudienceCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


