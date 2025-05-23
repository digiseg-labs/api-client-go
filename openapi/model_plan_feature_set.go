/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" /> 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the PlanFeatureSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanFeatureSet{}

// PlanFeatureSet struct for PlanFeatureSet
type PlanFeatureSet struct {
	MaxUsers int32 `json:"max_users"`
	MaxClients int32 `json:"max_clients"`
	MaxActiveStudies int32 `json:"max_active_studies"`
	MaxEventsPerStudy int64 `json:"max_events_per_study"`
	MaxStudyEventsPerMonth int64 `json:"max_study_events_per_month"`
	MaxAudienceLookupsPerMonth int64 `json:"max_audience_lookups_per_month"`
	StudyAudienceSet LimitedOrFullFeature `json:"study_audience_set"`
	StudyEventSet LimitedOrFullFeature `json:"study_event_set"`
	HasAudienceRecommendations bool `json:"has_audience_recommendations"`
	HasReportCustomization bool `json:"has_report_customization"`
	HasReportSharingClients bool `json:"has_report_sharing_clients"`
	HasReportSharingPublic bool `json:"has_report_sharing_public"`
	AdditionalProperties map[string]interface{}
}

type _PlanFeatureSet PlanFeatureSet

// NewPlanFeatureSet instantiates a new PlanFeatureSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanFeatureSet(maxUsers int32, maxClients int32, maxActiveStudies int32, maxEventsPerStudy int64, maxStudyEventsPerMonth int64, maxAudienceLookupsPerMonth int64, studyAudienceSet LimitedOrFullFeature, studyEventSet LimitedOrFullFeature, hasAudienceRecommendations bool, hasReportCustomization bool, hasReportSharingClients bool, hasReportSharingPublic bool) *PlanFeatureSet {
	this := PlanFeatureSet{}
	this.MaxUsers = maxUsers
	this.MaxClients = maxClients
	this.MaxActiveStudies = maxActiveStudies
	this.MaxEventsPerStudy = maxEventsPerStudy
	this.MaxStudyEventsPerMonth = maxStudyEventsPerMonth
	this.MaxAudienceLookupsPerMonth = maxAudienceLookupsPerMonth
	this.StudyAudienceSet = studyAudienceSet
	this.StudyEventSet = studyEventSet
	this.HasAudienceRecommendations = hasAudienceRecommendations
	this.HasReportCustomization = hasReportCustomization
	this.HasReportSharingClients = hasReportSharingClients
	this.HasReportSharingPublic = hasReportSharingPublic
	return &this
}

// NewPlanFeatureSetWithDefaults instantiates a new PlanFeatureSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanFeatureSetWithDefaults() *PlanFeatureSet {
	this := PlanFeatureSet{}
	return &this
}

// GetMaxUsers returns the MaxUsers field value
func (o *PlanFeatureSet) GetMaxUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxUsers
}

// GetMaxUsersOk returns a tuple with the MaxUsers field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUsers, true
}

// SetMaxUsers sets field value
func (o *PlanFeatureSet) SetMaxUsers(v int32) {
	o.MaxUsers = v
}

// GetMaxClients returns the MaxClients field value
func (o *PlanFeatureSet) GetMaxClients() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxClients
}

// GetMaxClientsOk returns a tuple with the MaxClients field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxClientsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxClients, true
}

// SetMaxClients sets field value
func (o *PlanFeatureSet) SetMaxClients(v int32) {
	o.MaxClients = v
}

// GetMaxActiveStudies returns the MaxActiveStudies field value
func (o *PlanFeatureSet) GetMaxActiveStudies() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxActiveStudies
}

// GetMaxActiveStudiesOk returns a tuple with the MaxActiveStudies field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxActiveStudiesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActiveStudies, true
}

// SetMaxActiveStudies sets field value
func (o *PlanFeatureSet) SetMaxActiveStudies(v int32) {
	o.MaxActiveStudies = v
}

// GetMaxEventsPerStudy returns the MaxEventsPerStudy field value
func (o *PlanFeatureSet) GetMaxEventsPerStudy() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxEventsPerStudy
}

// GetMaxEventsPerStudyOk returns a tuple with the MaxEventsPerStudy field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxEventsPerStudyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxEventsPerStudy, true
}

// SetMaxEventsPerStudy sets field value
func (o *PlanFeatureSet) SetMaxEventsPerStudy(v int64) {
	o.MaxEventsPerStudy = v
}

// GetMaxStudyEventsPerMonth returns the MaxStudyEventsPerMonth field value
func (o *PlanFeatureSet) GetMaxStudyEventsPerMonth() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxStudyEventsPerMonth
}

// GetMaxStudyEventsPerMonthOk returns a tuple with the MaxStudyEventsPerMonth field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxStudyEventsPerMonthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStudyEventsPerMonth, true
}

// SetMaxStudyEventsPerMonth sets field value
func (o *PlanFeatureSet) SetMaxStudyEventsPerMonth(v int64) {
	o.MaxStudyEventsPerMonth = v
}

// GetMaxAudienceLookupsPerMonth returns the MaxAudienceLookupsPerMonth field value
func (o *PlanFeatureSet) GetMaxAudienceLookupsPerMonth() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxAudienceLookupsPerMonth
}

// GetMaxAudienceLookupsPerMonthOk returns a tuple with the MaxAudienceLookupsPerMonth field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetMaxAudienceLookupsPerMonthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAudienceLookupsPerMonth, true
}

// SetMaxAudienceLookupsPerMonth sets field value
func (o *PlanFeatureSet) SetMaxAudienceLookupsPerMonth(v int64) {
	o.MaxAudienceLookupsPerMonth = v
}

// GetStudyAudienceSet returns the StudyAudienceSet field value
func (o *PlanFeatureSet) GetStudyAudienceSet() LimitedOrFullFeature {
	if o == nil {
		var ret LimitedOrFullFeature
		return ret
	}

	return o.StudyAudienceSet
}

// GetStudyAudienceSetOk returns a tuple with the StudyAudienceSet field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetStudyAudienceSetOk() (*LimitedOrFullFeature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StudyAudienceSet, true
}

// SetStudyAudienceSet sets field value
func (o *PlanFeatureSet) SetStudyAudienceSet(v LimitedOrFullFeature) {
	o.StudyAudienceSet = v
}

// GetStudyEventSet returns the StudyEventSet field value
func (o *PlanFeatureSet) GetStudyEventSet() LimitedOrFullFeature {
	if o == nil {
		var ret LimitedOrFullFeature
		return ret
	}

	return o.StudyEventSet
}

// GetStudyEventSetOk returns a tuple with the StudyEventSet field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetStudyEventSetOk() (*LimitedOrFullFeature, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StudyEventSet, true
}

// SetStudyEventSet sets field value
func (o *PlanFeatureSet) SetStudyEventSet(v LimitedOrFullFeature) {
	o.StudyEventSet = v
}

// GetHasAudienceRecommendations returns the HasAudienceRecommendations field value
func (o *PlanFeatureSet) GetHasAudienceRecommendations() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAudienceRecommendations
}

// GetHasAudienceRecommendationsOk returns a tuple with the HasAudienceRecommendations field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetHasAudienceRecommendationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAudienceRecommendations, true
}

// SetHasAudienceRecommendations sets field value
func (o *PlanFeatureSet) SetHasAudienceRecommendations(v bool) {
	o.HasAudienceRecommendations = v
}

// GetHasReportCustomization returns the HasReportCustomization field value
func (o *PlanFeatureSet) GetHasReportCustomization() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasReportCustomization
}

// GetHasReportCustomizationOk returns a tuple with the HasReportCustomization field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetHasReportCustomizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasReportCustomization, true
}

// SetHasReportCustomization sets field value
func (o *PlanFeatureSet) SetHasReportCustomization(v bool) {
	o.HasReportCustomization = v
}

// GetHasReportSharingClients returns the HasReportSharingClients field value
func (o *PlanFeatureSet) GetHasReportSharingClients() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasReportSharingClients
}

// GetHasReportSharingClientsOk returns a tuple with the HasReportSharingClients field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetHasReportSharingClientsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasReportSharingClients, true
}

// SetHasReportSharingClients sets field value
func (o *PlanFeatureSet) SetHasReportSharingClients(v bool) {
	o.HasReportSharingClients = v
}

// GetHasReportSharingPublic returns the HasReportSharingPublic field value
func (o *PlanFeatureSet) GetHasReportSharingPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasReportSharingPublic
}

// GetHasReportSharingPublicOk returns a tuple with the HasReportSharingPublic field value
// and a boolean to check if the value has been set.
func (o *PlanFeatureSet) GetHasReportSharingPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasReportSharingPublic, true
}

// SetHasReportSharingPublic sets field value
func (o *PlanFeatureSet) SetHasReportSharingPublic(v bool) {
	o.HasReportSharingPublic = v
}

func (o PlanFeatureSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanFeatureSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_users"] = o.MaxUsers
	toSerialize["max_clients"] = o.MaxClients
	toSerialize["max_active_studies"] = o.MaxActiveStudies
	toSerialize["max_events_per_study"] = o.MaxEventsPerStudy
	toSerialize["max_study_events_per_month"] = o.MaxStudyEventsPerMonth
	toSerialize["max_audience_lookups_per_month"] = o.MaxAudienceLookupsPerMonth
	toSerialize["study_audience_set"] = o.StudyAudienceSet
	toSerialize["study_event_set"] = o.StudyEventSet
	toSerialize["has_audience_recommendations"] = o.HasAudienceRecommendations
	toSerialize["has_report_customization"] = o.HasReportCustomization
	toSerialize["has_report_sharing_clients"] = o.HasReportSharingClients
	toSerialize["has_report_sharing_public"] = o.HasReportSharingPublic

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanFeatureSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_users",
		"max_clients",
		"max_active_studies",
		"max_events_per_study",
		"max_study_events_per_month",
		"max_audience_lookups_per_month",
		"study_audience_set",
		"study_event_set",
		"has_audience_recommendations",
		"has_report_customization",
		"has_report_sharing_clients",
		"has_report_sharing_public",
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

	varPlanFeatureSet := _PlanFeatureSet{}

	err = json.Unmarshal(data, &varPlanFeatureSet)

	if err != nil {
		return err
	}

	*o = PlanFeatureSet(varPlanFeatureSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_users")
		delete(additionalProperties, "max_clients")
		delete(additionalProperties, "max_active_studies")
		delete(additionalProperties, "max_events_per_study")
		delete(additionalProperties, "max_study_events_per_month")
		delete(additionalProperties, "max_audience_lookups_per_month")
		delete(additionalProperties, "study_audience_set")
		delete(additionalProperties, "study_event_set")
		delete(additionalProperties, "has_audience_recommendations")
		delete(additionalProperties, "has_report_customization")
		delete(additionalProperties, "has_report_sharing_clients")
		delete(additionalProperties, "has_report_sharing_public")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanFeatureSet struct {
	value *PlanFeatureSet
	isSet bool
}

func (v NullablePlanFeatureSet) Get() *PlanFeatureSet {
	return v.value
}

func (v *NullablePlanFeatureSet) Set(val *PlanFeatureSet) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanFeatureSet) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanFeatureSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanFeatureSet(val *PlanFeatureSet) *NullablePlanFeatureSet {
	return &NullablePlanFeatureSet{value: val, isSet: true}
}

func (v NullablePlanFeatureSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanFeatureSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


