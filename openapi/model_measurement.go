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

// checks if the Measurement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Measurement{}

// Measurement Represents a single measurement
type Measurement struct {
	// The event that triggered the measurement, typically `impression` or `click`
	Event string `json:"event"`
	// The real value of the measurement, typically a counter value (integer)
	Count *int32 `json:"count,omitempty"`
	// The fraction of events that fall within this object compared to the total of the category or segment (usually represented by the measurement's parent's parent). For example, if the measurement is \"impression\" on the `home_type` \"Apartment\" object, then the `fraction_of_total` represents the number of impressions on apartments compared to impressions from other `home_type` values. 
	FractionOfTotal *float64 `json:"fraction_of_total,omitempty"`
	// The rate of conversion to this measurement. Typically applies to measurements like \"click\" where it will represent the rate of impressions that turn into a click. 
	ConversionRate *float64 `json:"conversion_rate,omitempty"`
}

type _Measurement Measurement

// NewMeasurement instantiates a new Measurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurement(event string) *Measurement {
	this := Measurement{}
	this.Event = event
	return &this
}

// NewMeasurementWithDefaults instantiates a new Measurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementWithDefaults() *Measurement {
	this := Measurement{}
	return &this
}

// GetEvent returns the Event field value
func (o *Measurement) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *Measurement) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *Measurement) SetEvent(v string) {
	o.Event = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Measurement) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Measurement) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Measurement) SetCount(v int32) {
	o.Count = &v
}

// GetFractionOfTotal returns the FractionOfTotal field value if set, zero value otherwise.
func (o *Measurement) GetFractionOfTotal() float64 {
	if o == nil || IsNil(o.FractionOfTotal) {
		var ret float64
		return ret
	}
	return *o.FractionOfTotal
}

// GetFractionOfTotalOk returns a tuple with the FractionOfTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetFractionOfTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.FractionOfTotal) {
		return nil, false
	}
	return o.FractionOfTotal, true
}

// HasFractionOfTotal returns a boolean if a field has been set.
func (o *Measurement) HasFractionOfTotal() bool {
	if o != nil && !IsNil(o.FractionOfTotal) {
		return true
	}

	return false
}

// SetFractionOfTotal gets a reference to the given float64 and assigns it to the FractionOfTotal field.
func (o *Measurement) SetFractionOfTotal(v float64) {
	o.FractionOfTotal = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *Measurement) GetConversionRate() float64 {
	if o == nil || IsNil(o.ConversionRate) {
		var ret float64
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetConversionRateOk() (*float64, bool) {
	if o == nil || IsNil(o.ConversionRate) {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *Measurement) HasConversionRate() bool {
	if o != nil && !IsNil(o.ConversionRate) {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given float64 and assigns it to the ConversionRate field.
func (o *Measurement) SetConversionRate(v float64) {
	o.ConversionRate = &v
}

func (o Measurement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Measurement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.FractionOfTotal) {
		toSerialize["fraction_of_total"] = o.FractionOfTotal
	}
	if !IsNil(o.ConversionRate) {
		toSerialize["conversion_rate"] = o.ConversionRate
	}
	return toSerialize, nil
}

func (o *Measurement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
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

	varMeasurement := _Measurement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeasurement)

	if err != nil {
		return err
	}

	*o = Measurement(varMeasurement)

	return err
}

type NullableMeasurement struct {
	value *Measurement
	isSet bool
}

func (v NullableMeasurement) Get() *Measurement {
	return v.value
}

func (v *NullableMeasurement) Set(val *Measurement) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurement(val *Measurement) *NullableMeasurement {
	return &NullableMeasurement{value: val, isSet: true}
}

func (v NullableMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


