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
	"time"
)

// checks if the StudyBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyBase{}

// StudyBase struct for StudyBase
type StudyBase struct {
	Name *string `json:"name,omitempty"`
	// A set of labels that users can use to categorize their measurements. Can be used to indicate type of study, customer names or other traits. 
	Labels []string `json:"labels,omitempty"`
	// The ID of the account that owns this study
	AccountId *string `json:"account_id,omitempty"`
	// The date for which the study and its data ingestion will start
	StartDate *time.Time `json:"start_date,omitempty"`
	// The date for which the study and its data ingestion will end
	EndDate *time.Time `json:"end_date,omitempty"`
	LifeCycleStage *StudyLifecycleStage `json:"life_cycle_stage,omitempty"`
	IngestionStatus *StudyIngestionStatus `json:"ingestion_status,omitempty"`
	SummaryStats *StudySummaryStats `json:"summary_stats,omitempty"`
	Client *MeasurementClientItem `json:"client,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyBase StudyBase

// NewStudyBase instantiates a new StudyBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyBase() *StudyBase {
	this := StudyBase{}
	return &this
}

// NewStudyBaseWithDefaults instantiates a new StudyBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyBaseWithDefaults() *StudyBase {
	this := StudyBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StudyBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StudyBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StudyBase) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StudyBase) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StudyBase) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *StudyBase) SetLabels(v []string) {
	o.Labels = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StudyBase) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StudyBase) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StudyBase) SetAccountId(v string) {
	o.AccountId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *StudyBase) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *StudyBase) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *StudyBase) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *StudyBase) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *StudyBase) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *StudyBase) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *StudyBase) GetLifeCycleStage() StudyLifecycleStage {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret StudyLifecycleStage
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetLifeCycleStageOk() (*StudyLifecycleStage, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *StudyBase) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given StudyLifecycleStage and assigns it to the LifeCycleStage field.
func (o *StudyBase) SetLifeCycleStage(v StudyLifecycleStage) {
	o.LifeCycleStage = &v
}

// GetIngestionStatus returns the IngestionStatus field value if set, zero value otherwise.
func (o *StudyBase) GetIngestionStatus() StudyIngestionStatus {
	if o == nil || IsNil(o.IngestionStatus) {
		var ret StudyIngestionStatus
		return ret
	}
	return *o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetIngestionStatusOk() (*StudyIngestionStatus, bool) {
	if o == nil || IsNil(o.IngestionStatus) {
		return nil, false
	}
	return o.IngestionStatus, true
}

// HasIngestionStatus returns a boolean if a field has been set.
func (o *StudyBase) HasIngestionStatus() bool {
	if o != nil && !IsNil(o.IngestionStatus) {
		return true
	}

	return false
}

// SetIngestionStatus gets a reference to the given StudyIngestionStatus and assigns it to the IngestionStatus field.
func (o *StudyBase) SetIngestionStatus(v StudyIngestionStatus) {
	o.IngestionStatus = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *StudyBase) GetSummaryStats() StudySummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret StudySummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetSummaryStatsOk() (*StudySummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *StudyBase) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given StudySummaryStats and assigns it to the SummaryStats field.
func (o *StudyBase) SetSummaryStats(v StudySummaryStats) {
	o.SummaryStats = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *StudyBase) GetClient() MeasurementClientItem {
	if o == nil || IsNil(o.Client) {
		var ret MeasurementClientItem
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyBase) GetClientOk() (*MeasurementClientItem, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *StudyBase) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given MeasurementClientItem and assigns it to the Client field.
func (o *StudyBase) SetClient(v MeasurementClientItem) {
	o.Client = &v
}

func (o StudyBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.LifeCycleStage) {
		toSerialize["life_cycle_stage"] = o.LifeCycleStage
	}
	if !IsNil(o.IngestionStatus) {
		toSerialize["ingestion_status"] = o.IngestionStatus
	}
	if !IsNil(o.SummaryStats) {
		toSerialize["summary_stats"] = o.SummaryStats
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyBase) UnmarshalJSON(data []byte) (err error) {
	varStudyBase := _StudyBase{}

	err = json.Unmarshal(data, &varStudyBase)

	if err != nil {
		return err
	}

	*o = StudyBase(varStudyBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "life_cycle_stage")
		delete(additionalProperties, "ingestion_status")
		delete(additionalProperties, "summary_stats")
		delete(additionalProperties, "client")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyBase struct {
	value *StudyBase
	isSet bool
}

func (v NullableStudyBase) Get() *StudyBase {
	return v.value
}

func (v *NullableStudyBase) Set(val *StudyBase) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyBase) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyBase(val *StudyBase) *NullableStudyBase {
	return &NullableStudyBase{value: val, isSet: true}
}

func (v NullableStudyBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


