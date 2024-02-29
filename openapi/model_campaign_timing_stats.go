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

// checks if the CampaignTimingStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignTimingStats{}

// CampaignTimingStats struct for CampaignTimingStats
type CampaignTimingStats struct {
	HourOfDay *HourOfDayStats `json:"hour_of_day,omitempty"`
	DayOfWeek *DayOfWeekStats `json:"day_of_week,omitempty"`
	DayOfMonth *DayOfMonthStats `json:"day_of_month,omitempty"`
}

// NewCampaignTimingStats instantiates a new CampaignTimingStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignTimingStats() *CampaignTimingStats {
	this := CampaignTimingStats{}
	return &this
}

// NewCampaignTimingStatsWithDefaults instantiates a new CampaignTimingStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignTimingStatsWithDefaults() *CampaignTimingStats {
	this := CampaignTimingStats{}
	return &this
}

// GetHourOfDay returns the HourOfDay field value if set, zero value otherwise.
func (o *CampaignTimingStats) GetHourOfDay() HourOfDayStats {
	if o == nil || IsNil(o.HourOfDay) {
		var ret HourOfDayStats
		return ret
	}
	return *o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTimingStats) GetHourOfDayOk() (*HourOfDayStats, bool) {
	if o == nil || IsNil(o.HourOfDay) {
		return nil, false
	}
	return o.HourOfDay, true
}

// HasHourOfDay returns a boolean if a field has been set.
func (o *CampaignTimingStats) HasHourOfDay() bool {
	if o != nil && !IsNil(o.HourOfDay) {
		return true
	}

	return false
}

// SetHourOfDay gets a reference to the given HourOfDayStats and assigns it to the HourOfDay field.
func (o *CampaignTimingStats) SetHourOfDay(v HourOfDayStats) {
	o.HourOfDay = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *CampaignTimingStats) GetDayOfWeek() DayOfWeekStats {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret DayOfWeekStats
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTimingStats) GetDayOfWeekOk() (*DayOfWeekStats, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *CampaignTimingStats) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given DayOfWeekStats and assigns it to the DayOfWeek field.
func (o *CampaignTimingStats) SetDayOfWeek(v DayOfWeekStats) {
	o.DayOfWeek = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *CampaignTimingStats) GetDayOfMonth() DayOfMonthStats {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret DayOfMonthStats
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTimingStats) GetDayOfMonthOk() (*DayOfMonthStats, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *CampaignTimingStats) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given DayOfMonthStats and assigns it to the DayOfMonth field.
func (o *CampaignTimingStats) SetDayOfMonth(v DayOfMonthStats) {
	o.DayOfMonth = &v
}

func (o CampaignTimingStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignTimingStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HourOfDay) {
		toSerialize["hour_of_day"] = o.HourOfDay
	}
	if !IsNil(o.DayOfWeek) {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if !IsNil(o.DayOfMonth) {
		toSerialize["day_of_month"] = o.DayOfMonth
	}
	return toSerialize, nil
}

type NullableCampaignTimingStats struct {
	value *CampaignTimingStats
	isSet bool
}

func (v NullableCampaignTimingStats) Get() *CampaignTimingStats {
	return v.value
}

func (v *NullableCampaignTimingStats) Set(val *CampaignTimingStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignTimingStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignTimingStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignTimingStats(val *CampaignTimingStats) *NullableCampaignTimingStats {
	return &NullableCampaignTimingStats{value: val, isSet: true}
}

func (v NullableCampaignTimingStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignTimingStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


