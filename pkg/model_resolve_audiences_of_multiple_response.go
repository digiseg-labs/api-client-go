/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResolveAudiencesOfMultipleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveAudiencesOfMultipleResponse{}

// ResolveAudiencesOfMultipleResponse struct for ResolveAudiencesOfMultipleResponse
type ResolveAudiencesOfMultipleResponse struct {
	Results []ResolveAudiencesOfMultipleResponseItem `json:"results,omitempty"`
}

// NewResolveAudiencesOfMultipleResponse instantiates a new ResolveAudiencesOfMultipleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveAudiencesOfMultipleResponse() *ResolveAudiencesOfMultipleResponse {
	this := ResolveAudiencesOfMultipleResponse{}
	return &this
}

// NewResolveAudiencesOfMultipleResponseWithDefaults instantiates a new ResolveAudiencesOfMultipleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveAudiencesOfMultipleResponseWithDefaults() *ResolveAudiencesOfMultipleResponse {
	this := ResolveAudiencesOfMultipleResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResolveAudiencesOfMultipleResponse) GetResults() []ResolveAudiencesOfMultipleResponseItem {
	if o == nil || IsNil(o.Results) {
		var ret []ResolveAudiencesOfMultipleResponseItem
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveAudiencesOfMultipleResponse) GetResultsOk() ([]ResolveAudiencesOfMultipleResponseItem, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResolveAudiencesOfMultipleResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResolveAudiencesOfMultipleResponseItem and assigns it to the Results field.
func (o *ResolveAudiencesOfMultipleResponse) SetResults(v []ResolveAudiencesOfMultipleResponseItem) {
	o.Results = v
}

func (o ResolveAudiencesOfMultipleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveAudiencesOfMultipleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableResolveAudiencesOfMultipleResponse struct {
	value *ResolveAudiencesOfMultipleResponse
	isSet bool
}

func (v NullableResolveAudiencesOfMultipleResponse) Get() *ResolveAudiencesOfMultipleResponse {
	return v.value
}

func (v *NullableResolveAudiencesOfMultipleResponse) Set(val *ResolveAudiencesOfMultipleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveAudiencesOfMultipleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveAudiencesOfMultipleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveAudiencesOfMultipleResponse(val *ResolveAudiencesOfMultipleResponse) *NullableResolveAudiencesOfMultipleResponse {
	return &NullableResolveAudiencesOfMultipleResponse{value: val, isSet: true}
}

func (v NullableResolveAudiencesOfMultipleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveAudiencesOfMultipleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

