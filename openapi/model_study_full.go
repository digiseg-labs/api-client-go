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

// checks if the StudyFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyFull{}

// StudyFull struct for StudyFull
type StudyFull struct {
	// Unique ID for the object
	Id string `json:"id"`
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
	EventLinks *MeasurementEventLinks `json:"event_links,omitempty"`
	// If present, an upper limit on the number of events that will be processed in this study.
	EventCap *int32 `json:"event_cap,omitempty"`
	// The URL to a banner image for the study. Note that the banner image is used only for Digiseg study reporting and presentation, it does NOT represent any delivered banner ad creatives or similar. 
	BannerImageUrl *string `json:"banner_image_url,omitempty"`
	IntegrationPlatform *MeasurementIntegrationPlatform `json:"integration_platform,omitempty"`
	// Determines if the study is an example study, used to demonstrate product capabilities
	IsExample *bool `json:"is_example,omitempty"`
	// Date and time of the object creation
	CreatedAt time.Time `json:"created_at"`
	// ID of the user who created the object
	CreatedBy string `json:"created_by"`
	// Date and time of the latest update to the object
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the user who last updated the object
	UpdatedBy *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudyFull StudyFull

// NewStudyFull instantiates a new StudyFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyFull(id string, createdAt time.Time, createdBy string) *StudyFull {
	this := StudyFull{}
	this.Id = id
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	return &this
}

// NewStudyFullWithDefaults instantiates a new StudyFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyFullWithDefaults() *StudyFull {
	this := StudyFull{}
	return &this
}

// GetId returns the Id field value
func (o *StudyFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StudyFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StudyFull) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StudyFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StudyFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StudyFull) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StudyFull) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StudyFull) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *StudyFull) SetLabels(v []string) {
	o.Labels = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StudyFull) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StudyFull) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StudyFull) SetAccountId(v string) {
	o.AccountId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *StudyFull) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *StudyFull) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *StudyFull) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *StudyFull) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *StudyFull) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *StudyFull) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *StudyFull) GetLifeCycleStage() StudyLifecycleStage {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret StudyLifecycleStage
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetLifeCycleStageOk() (*StudyLifecycleStage, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *StudyFull) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given StudyLifecycleStage and assigns it to the LifeCycleStage field.
func (o *StudyFull) SetLifeCycleStage(v StudyLifecycleStage) {
	o.LifeCycleStage = &v
}

// GetIngestionStatus returns the IngestionStatus field value if set, zero value otherwise.
func (o *StudyFull) GetIngestionStatus() StudyIngestionStatus {
	if o == nil || IsNil(o.IngestionStatus) {
		var ret StudyIngestionStatus
		return ret
	}
	return *o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetIngestionStatusOk() (*StudyIngestionStatus, bool) {
	if o == nil || IsNil(o.IngestionStatus) {
		return nil, false
	}
	return o.IngestionStatus, true
}

// HasIngestionStatus returns a boolean if a field has been set.
func (o *StudyFull) HasIngestionStatus() bool {
	if o != nil && !IsNil(o.IngestionStatus) {
		return true
	}

	return false
}

// SetIngestionStatus gets a reference to the given StudyIngestionStatus and assigns it to the IngestionStatus field.
func (o *StudyFull) SetIngestionStatus(v StudyIngestionStatus) {
	o.IngestionStatus = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *StudyFull) GetSummaryStats() StudySummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret StudySummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetSummaryStatsOk() (*StudySummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *StudyFull) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given StudySummaryStats and assigns it to the SummaryStats field.
func (o *StudyFull) SetSummaryStats(v StudySummaryStats) {
	o.SummaryStats = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *StudyFull) GetClient() MeasurementClientItem {
	if o == nil || IsNil(o.Client) {
		var ret MeasurementClientItem
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetClientOk() (*MeasurementClientItem, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *StudyFull) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given MeasurementClientItem and assigns it to the Client field.
func (o *StudyFull) SetClient(v MeasurementClientItem) {
	o.Client = &v
}

// GetEventLinks returns the EventLinks field value if set, zero value otherwise.
func (o *StudyFull) GetEventLinks() MeasurementEventLinks {
	if o == nil || IsNil(o.EventLinks) {
		var ret MeasurementEventLinks
		return ret
	}
	return *o.EventLinks
}

// GetEventLinksOk returns a tuple with the EventLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetEventLinksOk() (*MeasurementEventLinks, bool) {
	if o == nil || IsNil(o.EventLinks) {
		return nil, false
	}
	return o.EventLinks, true
}

// HasEventLinks returns a boolean if a field has been set.
func (o *StudyFull) HasEventLinks() bool {
	if o != nil && !IsNil(o.EventLinks) {
		return true
	}

	return false
}

// SetEventLinks gets a reference to the given MeasurementEventLinks and assigns it to the EventLinks field.
func (o *StudyFull) SetEventLinks(v MeasurementEventLinks) {
	o.EventLinks = &v
}

// GetEventCap returns the EventCap field value if set, zero value otherwise.
func (o *StudyFull) GetEventCap() int32 {
	if o == nil || IsNil(o.EventCap) {
		var ret int32
		return ret
	}
	return *o.EventCap
}

// GetEventCapOk returns a tuple with the EventCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetEventCapOk() (*int32, bool) {
	if o == nil || IsNil(o.EventCap) {
		return nil, false
	}
	return o.EventCap, true
}

// HasEventCap returns a boolean if a field has been set.
func (o *StudyFull) HasEventCap() bool {
	if o != nil && !IsNil(o.EventCap) {
		return true
	}

	return false
}

// SetEventCap gets a reference to the given int32 and assigns it to the EventCap field.
func (o *StudyFull) SetEventCap(v int32) {
	o.EventCap = &v
}

// GetBannerImageUrl returns the BannerImageUrl field value if set, zero value otherwise.
func (o *StudyFull) GetBannerImageUrl() string {
	if o == nil || IsNil(o.BannerImageUrl) {
		var ret string
		return ret
	}
	return *o.BannerImageUrl
}

// GetBannerImageUrlOk returns a tuple with the BannerImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetBannerImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImageUrl) {
		return nil, false
	}
	return o.BannerImageUrl, true
}

// HasBannerImageUrl returns a boolean if a field has been set.
func (o *StudyFull) HasBannerImageUrl() bool {
	if o != nil && !IsNil(o.BannerImageUrl) {
		return true
	}

	return false
}

// SetBannerImageUrl gets a reference to the given string and assigns it to the BannerImageUrl field.
func (o *StudyFull) SetBannerImageUrl(v string) {
	o.BannerImageUrl = &v
}

// GetIntegrationPlatform returns the IntegrationPlatform field value if set, zero value otherwise.
func (o *StudyFull) GetIntegrationPlatform() MeasurementIntegrationPlatform {
	if o == nil || IsNil(o.IntegrationPlatform) {
		var ret MeasurementIntegrationPlatform
		return ret
	}
	return *o.IntegrationPlatform
}

// GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetIntegrationPlatformOk() (*MeasurementIntegrationPlatform, bool) {
	if o == nil || IsNil(o.IntegrationPlatform) {
		return nil, false
	}
	return o.IntegrationPlatform, true
}

// HasIntegrationPlatform returns a boolean if a field has been set.
func (o *StudyFull) HasIntegrationPlatform() bool {
	if o != nil && !IsNil(o.IntegrationPlatform) {
		return true
	}

	return false
}

// SetIntegrationPlatform gets a reference to the given MeasurementIntegrationPlatform and assigns it to the IntegrationPlatform field.
func (o *StudyFull) SetIntegrationPlatform(v MeasurementIntegrationPlatform) {
	o.IntegrationPlatform = &v
}

// GetIsExample returns the IsExample field value if set, zero value otherwise.
func (o *StudyFull) GetIsExample() bool {
	if o == nil || IsNil(o.IsExample) {
		var ret bool
		return ret
	}
	return *o.IsExample
}

// GetIsExampleOk returns a tuple with the IsExample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetIsExampleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExample) {
		return nil, false
	}
	return o.IsExample, true
}

// HasIsExample returns a boolean if a field has been set.
func (o *StudyFull) HasIsExample() bool {
	if o != nil && !IsNil(o.IsExample) {
		return true
	}

	return false
}

// SetIsExample gets a reference to the given bool and assigns it to the IsExample field.
func (o *StudyFull) SetIsExample(v bool) {
	o.IsExample = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *StudyFull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *StudyFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *StudyFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *StudyFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *StudyFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *StudyFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *StudyFull) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StudyFull) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *StudyFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *StudyFull) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyFull) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *StudyFull) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *StudyFull) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o StudyFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StudyFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"created_by",
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

	varStudyFull := _StudyFull{}

	err = json.Unmarshal(data, &varStudyFull)

	if err != nil {
		return err
	}

	*o = StudyFull(varStudyFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
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
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudyFull struct {
	value *StudyFull
	isSet bool
}

func (v NullableStudyFull) Get() *StudyFull {
	return v.value
}

func (v *NullableStudyFull) Set(val *StudyFull) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyFull) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyFull(val *StudyFull) *NullableStudyFull {
	return &NullableStudyFull{value: val, isSet: true}
}

func (v NullableStudyFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


