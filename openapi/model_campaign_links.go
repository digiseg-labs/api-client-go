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

// checks if the CampaignLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignLinks{}

// CampaignLinks struct for CampaignLinks
type CampaignLinks struct {
	// Link to the country statistics for the campaign
	StatsCountries *string `json:"stats_countries,omitempty"`
	// Link to the audience statistics for the campaign
	StatsAudiences *string `json:"stats_audiences,omitempty"`
	// Link to the timing statistics for the campaign
	StatsTiming *string `json:"stats_timing,omitempty"`
	// Link to the frequency statistics for the campaign
	StatsFrequencies *string `json:"stats_frequencies,omitempty"`
}

// NewCampaignLinks instantiates a new CampaignLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignLinks() *CampaignLinks {
	this := CampaignLinks{}
	return &this
}

// NewCampaignLinksWithDefaults instantiates a new CampaignLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignLinksWithDefaults() *CampaignLinks {
	this := CampaignLinks{}
	return &this
}

// GetStatsCountries returns the StatsCountries field value if set, zero value otherwise.
func (o *CampaignLinks) GetStatsCountries() string {
	if o == nil || IsNil(o.StatsCountries) {
		var ret string
		return ret
	}
	return *o.StatsCountries
}

// GetStatsCountriesOk returns a tuple with the StatsCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetStatsCountriesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsCountries) {
		return nil, false
	}
	return o.StatsCountries, true
}

// HasStatsCountries returns a boolean if a field has been set.
func (o *CampaignLinks) HasStatsCountries() bool {
	if o != nil && !IsNil(o.StatsCountries) {
		return true
	}

	return false
}

// SetStatsCountries gets a reference to the given string and assigns it to the StatsCountries field.
func (o *CampaignLinks) SetStatsCountries(v string) {
	o.StatsCountries = &v
}

// GetStatsAudiences returns the StatsAudiences field value if set, zero value otherwise.
func (o *CampaignLinks) GetStatsAudiences() string {
	if o == nil || IsNil(o.StatsAudiences) {
		var ret string
		return ret
	}
	return *o.StatsAudiences
}

// GetStatsAudiencesOk returns a tuple with the StatsAudiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetStatsAudiencesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsAudiences) {
		return nil, false
	}
	return o.StatsAudiences, true
}

// HasStatsAudiences returns a boolean if a field has been set.
func (o *CampaignLinks) HasStatsAudiences() bool {
	if o != nil && !IsNil(o.StatsAudiences) {
		return true
	}

	return false
}

// SetStatsAudiences gets a reference to the given string and assigns it to the StatsAudiences field.
func (o *CampaignLinks) SetStatsAudiences(v string) {
	o.StatsAudiences = &v
}

// GetStatsTiming returns the StatsTiming field value if set, zero value otherwise.
func (o *CampaignLinks) GetStatsTiming() string {
	if o == nil || IsNil(o.StatsTiming) {
		var ret string
		return ret
	}
	return *o.StatsTiming
}

// GetStatsTimingOk returns a tuple with the StatsTiming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetStatsTimingOk() (*string, bool) {
	if o == nil || IsNil(o.StatsTiming) {
		return nil, false
	}
	return o.StatsTiming, true
}

// HasStatsTiming returns a boolean if a field has been set.
func (o *CampaignLinks) HasStatsTiming() bool {
	if o != nil && !IsNil(o.StatsTiming) {
		return true
	}

	return false
}

// SetStatsTiming gets a reference to the given string and assigns it to the StatsTiming field.
func (o *CampaignLinks) SetStatsTiming(v string) {
	o.StatsTiming = &v
}

// GetStatsFrequencies returns the StatsFrequencies field value if set, zero value otherwise.
func (o *CampaignLinks) GetStatsFrequencies() string {
	if o == nil || IsNil(o.StatsFrequencies) {
		var ret string
		return ret
	}
	return *o.StatsFrequencies
}

// GetStatsFrequenciesOk returns a tuple with the StatsFrequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetStatsFrequenciesOk() (*string, bool) {
	if o == nil || IsNil(o.StatsFrequencies) {
		return nil, false
	}
	return o.StatsFrequencies, true
}

// HasStatsFrequencies returns a boolean if a field has been set.
func (o *CampaignLinks) HasStatsFrequencies() bool {
	if o != nil && !IsNil(o.StatsFrequencies) {
		return true
	}

	return false
}

// SetStatsFrequencies gets a reference to the given string and assigns it to the StatsFrequencies field.
func (o *CampaignLinks) SetStatsFrequencies(v string) {
	o.StatsFrequencies = &v
}

func (o CampaignLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatsCountries) {
		toSerialize["stats_countries"] = o.StatsCountries
	}
	if !IsNil(o.StatsAudiences) {
		toSerialize["stats_audiences"] = o.StatsAudiences
	}
	if !IsNil(o.StatsTiming) {
		toSerialize["stats_timing"] = o.StatsTiming
	}
	if !IsNil(o.StatsFrequencies) {
		toSerialize["stats_frequencies"] = o.StatsFrequencies
	}
	return toSerialize, nil
}

type NullableCampaignLinks struct {
	value *CampaignLinks
	isSet bool
}

func (v NullableCampaignLinks) Get() *CampaignLinks {
	return v.value
}

func (v *NullableCampaignLinks) Set(val *CampaignLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignLinks(val *CampaignLinks) *NullableCampaignLinks {
	return &NullableCampaignLinks{value: val, isSet: true}
}

func (v NullableCampaignLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


