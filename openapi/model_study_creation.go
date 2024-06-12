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
	"time"
	"fmt"
)

// checks if the StudyCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyCreation{}

// StudyCreation struct for StudyCreation
type StudyCreation struct {
	Name string `json:"name"`
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
	EventLinks *MeasurementEventLinks `json:"event_links,omitempty"`
	// If present, an upper limit on the number of events that will be processed in this study.
	EventCap *int32 `json:"event_cap,omitempty"`
	// The URL to a banner image for the study. Note that the banner image is used only for Digiseg study reporting and presentation, it does NOT represent any delivered banner ad creatives or similar. 
	BannerImageUrl *string `json:"banner_image_url,omitempty"`
	IntegrationPlatform *MeasurementIntegrationPlatform `json:"integration_platform,omitempty"`
	// Determines if the study is an example study, used to demonstrate product capabilities
	IsExample *bool `json:"is_example,omitempty"`
	EventSet MeasurementEventSet `json:"event_set"`
	// The ID of the measurement client that this study is for
	ClientId *string `json:"client_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyCreation StudyCreation

// NewStudyCreation instantiates a new StudyCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyCreation(name string, eventSet MeasurementEventSet) *StudyCreation {
	this := StudyCreation{}
	this.Name = name
	this.EventSet = eventSet
	return &this
}

// NewStudyCreationWithDefaults instantiates a new StudyCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyCreationWithDefaults() *StudyCreation {
	this := StudyCreation{}
	return &this
}

// GetName returns the Name field value
func (o *StudyCreation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StudyCreation) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StudyCreation) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StudyCreation) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *StudyCreation) SetLabels(v []string) {
	o.Labels = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StudyCreation) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StudyCreation) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StudyCreation) SetAccountId(v string) {
	o.AccountId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *StudyCreation) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *StudyCreation) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *StudyCreation) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *StudyCreation) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *StudyCreation) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *StudyCreation) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *StudyCreation) GetLifeCycleStage() StudyLifecycleStage {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret StudyLifecycleStage
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetLifeCycleStageOk() (*StudyLifecycleStage, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *StudyCreation) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given StudyLifecycleStage and assigns it to the LifeCycleStage field.
func (o *StudyCreation) SetLifeCycleStage(v StudyLifecycleStage) {
	o.LifeCycleStage = &v
}

// GetIngestionStatus returns the IngestionStatus field value if set, zero value otherwise.
func (o *StudyCreation) GetIngestionStatus() StudyIngestionStatus {
	if o == nil || IsNil(o.IngestionStatus) {
		var ret StudyIngestionStatus
		return ret
	}
	return *o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetIngestionStatusOk() (*StudyIngestionStatus, bool) {
	if o == nil || IsNil(o.IngestionStatus) {
		return nil, false
	}
	return o.IngestionStatus, true
}

// HasIngestionStatus returns a boolean if a field has been set.
func (o *StudyCreation) HasIngestionStatus() bool {
	if o != nil && !IsNil(o.IngestionStatus) {
		return true
	}

	return false
}

// SetIngestionStatus gets a reference to the given StudyIngestionStatus and assigns it to the IngestionStatus field.
func (o *StudyCreation) SetIngestionStatus(v StudyIngestionStatus) {
	o.IngestionStatus = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *StudyCreation) GetSummaryStats() StudySummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret StudySummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetSummaryStatsOk() (*StudySummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *StudyCreation) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given StudySummaryStats and assigns it to the SummaryStats field.
func (o *StudyCreation) SetSummaryStats(v StudySummaryStats) {
	o.SummaryStats = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *StudyCreation) GetClient() MeasurementClientItem {
	if o == nil || IsNil(o.Client) {
		var ret MeasurementClientItem
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetClientOk() (*MeasurementClientItem, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *StudyCreation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given MeasurementClientItem and assigns it to the Client field.
func (o *StudyCreation) SetClient(v MeasurementClientItem) {
	o.Client = &v
}

// GetEventLinks returns the EventLinks field value if set, zero value otherwise.
func (o *StudyCreation) GetEventLinks() MeasurementEventLinks {
	if o == nil || IsNil(o.EventLinks) {
		var ret MeasurementEventLinks
		return ret
	}
	return *o.EventLinks
}

// GetEventLinksOk returns a tuple with the EventLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetEventLinksOk() (*MeasurementEventLinks, bool) {
	if o == nil || IsNil(o.EventLinks) {
		return nil, false
	}
	return o.EventLinks, true
}

// HasEventLinks returns a boolean if a field has been set.
func (o *StudyCreation) HasEventLinks() bool {
	if o != nil && !IsNil(o.EventLinks) {
		return true
	}

	return false
}

// SetEventLinks gets a reference to the given MeasurementEventLinks and assigns it to the EventLinks field.
func (o *StudyCreation) SetEventLinks(v MeasurementEventLinks) {
	o.EventLinks = &v
}

// GetEventCap returns the EventCap field value if set, zero value otherwise.
func (o *StudyCreation) GetEventCap() int32 {
	if o == nil || IsNil(o.EventCap) {
		var ret int32
		return ret
	}
	return *o.EventCap
}

// GetEventCapOk returns a tuple with the EventCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetEventCapOk() (*int32, bool) {
	if o == nil || IsNil(o.EventCap) {
		return nil, false
	}
	return o.EventCap, true
}

// HasEventCap returns a boolean if a field has been set.
func (o *StudyCreation) HasEventCap() bool {
	if o != nil && !IsNil(o.EventCap) {
		return true
	}

	return false
}

// SetEventCap gets a reference to the given int32 and assigns it to the EventCap field.
func (o *StudyCreation) SetEventCap(v int32) {
	o.EventCap = &v
}

// GetBannerImageUrl returns the BannerImageUrl field value if set, zero value otherwise.
func (o *StudyCreation) GetBannerImageUrl() string {
	if o == nil || IsNil(o.BannerImageUrl) {
		var ret string
		return ret
	}
	return *o.BannerImageUrl
}

// GetBannerImageUrlOk returns a tuple with the BannerImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetBannerImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImageUrl) {
		return nil, false
	}
	return o.BannerImageUrl, true
}

// HasBannerImageUrl returns a boolean if a field has been set.
func (o *StudyCreation) HasBannerImageUrl() bool {
	if o != nil && !IsNil(o.BannerImageUrl) {
		return true
	}

	return false
}

// SetBannerImageUrl gets a reference to the given string and assigns it to the BannerImageUrl field.
func (o *StudyCreation) SetBannerImageUrl(v string) {
	o.BannerImageUrl = &v
}

// GetIntegrationPlatform returns the IntegrationPlatform field value if set, zero value otherwise.
func (o *StudyCreation) GetIntegrationPlatform() MeasurementIntegrationPlatform {
	if o == nil || IsNil(o.IntegrationPlatform) {
		var ret MeasurementIntegrationPlatform
		return ret
	}
	return *o.IntegrationPlatform
}

// GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetIntegrationPlatformOk() (*MeasurementIntegrationPlatform, bool) {
	if o == nil || IsNil(o.IntegrationPlatform) {
		return nil, false
	}
	return o.IntegrationPlatform, true
}

// HasIntegrationPlatform returns a boolean if a field has been set.
func (o *StudyCreation) HasIntegrationPlatform() bool {
	if o != nil && !IsNil(o.IntegrationPlatform) {
		return true
	}

	return false
}

// SetIntegrationPlatform gets a reference to the given MeasurementIntegrationPlatform and assigns it to the IntegrationPlatform field.
func (o *StudyCreation) SetIntegrationPlatform(v MeasurementIntegrationPlatform) {
	o.IntegrationPlatform = &v
}

// GetIsExample returns the IsExample field value if set, zero value otherwise.
func (o *StudyCreation) GetIsExample() bool {
	if o == nil || IsNil(o.IsExample) {
		var ret bool
		return ret
	}
	return *o.IsExample
}

// GetIsExampleOk returns a tuple with the IsExample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetIsExampleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExample) {
		return nil, false
	}
	return o.IsExample, true
}

// HasIsExample returns a boolean if a field has been set.
func (o *StudyCreation) HasIsExample() bool {
	if o != nil && !IsNil(o.IsExample) {
		return true
	}

	return false
}

// SetIsExample gets a reference to the given bool and assigns it to the IsExample field.
func (o *StudyCreation) SetIsExample(v bool) {
	o.IsExample = &v
}

// GetEventSet returns the EventSet field value
func (o *StudyCreation) GetEventSet() MeasurementEventSet {
	if o == nil {
		var ret MeasurementEventSet
		return ret
	}

	return o.EventSet
}

// GetEventSetOk returns a tuple with the EventSet field value
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetEventSetOk() (*MeasurementEventSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSet, true
}

// SetEventSet sets field value
func (o *StudyCreation) SetEventSet(v MeasurementEventSet) {
	o.EventSet = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *StudyCreation) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyCreation) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *StudyCreation) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *StudyCreation) SetClientId(v string) {
	o.ClientId = &v
}

func (o StudyCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	if !IsNil(o.EventLinks) {
		toSerialize["event_links"] = o.EventLinks
	}
	if !IsNil(o.EventCap) {
		toSerialize["event_cap"] = o.EventCap
	}
	if !IsNil(o.BannerImageUrl) {
		toSerialize["banner_image_url"] = o.BannerImageUrl
	}
	if !IsNil(o.IntegrationPlatform) {
		toSerialize["integration_platform"] = o.IntegrationPlatform
	}
	if !IsNil(o.IsExample) {
		toSerialize["is_example"] = o.IsExample
	}
	toSerialize["event_set"] = o.EventSet
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyCreation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"event_set",
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

	varStudyCreation := _StudyCreation{}

	err = json.Unmarshal(data, &varStudyCreation)

	if err != nil {
		return err
	}

	*o = StudyCreation(varStudyCreation)

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
		delete(additionalProperties, "event_links")
		delete(additionalProperties, "event_cap")
		delete(additionalProperties, "banner_image_url")
		delete(additionalProperties, "integration_platform")
		delete(additionalProperties, "is_example")
		delete(additionalProperties, "event_set")
		delete(additionalProperties, "client_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyCreation struct {
	value *StudyCreation
	isSet bool
}

func (v NullableStudyCreation) Get() *StudyCreation {
	return v.value
}

func (v *NullableStudyCreation) Set(val *StudyCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyCreation(val *StudyCreation) *NullableStudyCreation {
	return &NullableStudyCreation{value: val, isSet: true}
}

func (v NullableStudyCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


