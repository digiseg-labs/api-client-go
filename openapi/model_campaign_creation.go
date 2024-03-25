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
	"time"
	"bytes"
	"fmt"
)

// checks if the CampaignCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignCreation{}

// CampaignCreation struct for CampaignCreation
type CampaignCreation struct {
	Name string `json:"name"`
	// A set of labels that users can use to categorize their measurements. Can be used to indicate type of campaign, customer names or other traits. 
	Labels []string `json:"labels,omitempty"`
	// The ID of the account that owns this campaign
	AccountId *string `json:"account_id,omitempty"`
	// The date for which the campaign and its data ingestion will start.
	StartDate *time.Time `json:"start_date,omitempty"`
	LifeCycleStage *CampaignLifecycleStage `json:"life_cycle_stage,omitempty"`
	IngestionStatus *CampaignIngestionStatus `json:"ingestion_status,omitempty"`
	SummaryStats *CampaignSummaryStats `json:"summary_stats,omitempty"`
	Client *MeasurementClientItem `json:"client,omitempty"`
	EventLinks *CampaignEventLinks `json:"event_links,omitempty"`
	// The URL to a banner image for the campaign. Note that the banner image is used only for Digiseg campaign reporting and presentation, it does NOT represent any delivered banner ad creatives or similar. 
	BannerImageUrl *string `json:"banner_image_url,omitempty"`
	IntegrationPlatform *CampaignIntegrationPlatform `json:"integration_platform,omitempty"`
	EventSet CampaignEventSet `json:"event_set"`
	// The ID of the measurement client that this campaign is for
	ClientId *string `json:"client_id,omitempty"`
}

type _CampaignCreation CampaignCreation

// NewCampaignCreation instantiates a new CampaignCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignCreation(name string, eventSet CampaignEventSet) *CampaignCreation {
	this := CampaignCreation{}
	this.Name = name
	this.EventSet = eventSet
	return &this
}

// NewCampaignCreationWithDefaults instantiates a new CampaignCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignCreationWithDefaults() *CampaignCreation {
	this := CampaignCreation{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignCreation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignCreation) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CampaignCreation) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CampaignCreation) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *CampaignCreation) SetLabels(v []string) {
	o.Labels = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CampaignCreation) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CampaignCreation) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CampaignCreation) SetAccountId(v string) {
	o.AccountId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CampaignCreation) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CampaignCreation) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CampaignCreation) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *CampaignCreation) GetLifeCycleStage() CampaignLifecycleStage {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret CampaignLifecycleStage
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *CampaignCreation) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given CampaignLifecycleStage and assigns it to the LifeCycleStage field.
func (o *CampaignCreation) SetLifeCycleStage(v CampaignLifecycleStage) {
	o.LifeCycleStage = &v
}

// GetIngestionStatus returns the IngestionStatus field value if set, zero value otherwise.
func (o *CampaignCreation) GetIngestionStatus() CampaignIngestionStatus {
	if o == nil || IsNil(o.IngestionStatus) {
		var ret CampaignIngestionStatus
		return ret
	}
	return *o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetIngestionStatusOk() (*CampaignIngestionStatus, bool) {
	if o == nil || IsNil(o.IngestionStatus) {
		return nil, false
	}
	return o.IngestionStatus, true
}

// HasIngestionStatus returns a boolean if a field has been set.
func (o *CampaignCreation) HasIngestionStatus() bool {
	if o != nil && !IsNil(o.IngestionStatus) {
		return true
	}

	return false
}

// SetIngestionStatus gets a reference to the given CampaignIngestionStatus and assigns it to the IngestionStatus field.
func (o *CampaignCreation) SetIngestionStatus(v CampaignIngestionStatus) {
	o.IngestionStatus = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *CampaignCreation) GetSummaryStats() CampaignSummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret CampaignSummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetSummaryStatsOk() (*CampaignSummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *CampaignCreation) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given CampaignSummaryStats and assigns it to the SummaryStats field.
func (o *CampaignCreation) SetSummaryStats(v CampaignSummaryStats) {
	o.SummaryStats = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *CampaignCreation) GetClient() MeasurementClientItem {
	if o == nil || IsNil(o.Client) {
		var ret MeasurementClientItem
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetClientOk() (*MeasurementClientItem, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *CampaignCreation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given MeasurementClientItem and assigns it to the Client field.
func (o *CampaignCreation) SetClient(v MeasurementClientItem) {
	o.Client = &v
}

// GetEventLinks returns the EventLinks field value if set, zero value otherwise.
func (o *CampaignCreation) GetEventLinks() CampaignEventLinks {
	if o == nil || IsNil(o.EventLinks) {
		var ret CampaignEventLinks
		return ret
	}
	return *o.EventLinks
}

// GetEventLinksOk returns a tuple with the EventLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetEventLinksOk() (*CampaignEventLinks, bool) {
	if o == nil || IsNil(o.EventLinks) {
		return nil, false
	}
	return o.EventLinks, true
}

// HasEventLinks returns a boolean if a field has been set.
func (o *CampaignCreation) HasEventLinks() bool {
	if o != nil && !IsNil(o.EventLinks) {
		return true
	}

	return false
}

// SetEventLinks gets a reference to the given CampaignEventLinks and assigns it to the EventLinks field.
func (o *CampaignCreation) SetEventLinks(v CampaignEventLinks) {
	o.EventLinks = &v
}

// GetBannerImageUrl returns the BannerImageUrl field value if set, zero value otherwise.
func (o *CampaignCreation) GetBannerImageUrl() string {
	if o == nil || IsNil(o.BannerImageUrl) {
		var ret string
		return ret
	}
	return *o.BannerImageUrl
}

// GetBannerImageUrlOk returns a tuple with the BannerImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetBannerImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImageUrl) {
		return nil, false
	}
	return o.BannerImageUrl, true
}

// HasBannerImageUrl returns a boolean if a field has been set.
func (o *CampaignCreation) HasBannerImageUrl() bool {
	if o != nil && !IsNil(o.BannerImageUrl) {
		return true
	}

	return false
}

// SetBannerImageUrl gets a reference to the given string and assigns it to the BannerImageUrl field.
func (o *CampaignCreation) SetBannerImageUrl(v string) {
	o.BannerImageUrl = &v
}

// GetIntegrationPlatform returns the IntegrationPlatform field value if set, zero value otherwise.
func (o *CampaignCreation) GetIntegrationPlatform() CampaignIntegrationPlatform {
	if o == nil || IsNil(o.IntegrationPlatform) {
		var ret CampaignIntegrationPlatform
		return ret
	}
	return *o.IntegrationPlatform
}

// GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetIntegrationPlatformOk() (*CampaignIntegrationPlatform, bool) {
	if o == nil || IsNil(o.IntegrationPlatform) {
		return nil, false
	}
	return o.IntegrationPlatform, true
}

// HasIntegrationPlatform returns a boolean if a field has been set.
func (o *CampaignCreation) HasIntegrationPlatform() bool {
	if o != nil && !IsNil(o.IntegrationPlatform) {
		return true
	}

	return false
}

// SetIntegrationPlatform gets a reference to the given CampaignIntegrationPlatform and assigns it to the IntegrationPlatform field.
func (o *CampaignCreation) SetIntegrationPlatform(v CampaignIntegrationPlatform) {
	o.IntegrationPlatform = &v
}

// GetEventSet returns the EventSet field value
func (o *CampaignCreation) GetEventSet() CampaignEventSet {
	if o == nil {
		var ret CampaignEventSet
		return ret
	}

	return o.EventSet
}

// GetEventSetOk returns a tuple with the EventSet field value
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetEventSetOk() (*CampaignEventSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSet, true
}

// SetEventSet sets field value
func (o *CampaignCreation) SetEventSet(v CampaignEventSet) {
	o.EventSet = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CampaignCreation) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreation) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CampaignCreation) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CampaignCreation) SetClientId(v string) {
	o.ClientId = &v
}

func (o CampaignCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignCreation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BannerImageUrl) {
		toSerialize["banner_image_url"] = o.BannerImageUrl
	}
	if !IsNil(o.IntegrationPlatform) {
		toSerialize["integration_platform"] = o.IntegrationPlatform
	}
	toSerialize["event_set"] = o.EventSet
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	return toSerialize, nil
}

func (o *CampaignCreation) UnmarshalJSON(data []byte) (err error) {
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

	varCampaignCreation := _CampaignCreation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignCreation)

	if err != nil {
		return err
	}

	*o = CampaignCreation(varCampaignCreation)

	return err
}

type NullableCampaignCreation struct {
	value *CampaignCreation
	isSet bool
}

func (v NullableCampaignCreation) Get() *CampaignCreation {
	return v.value
}

func (v *NullableCampaignCreation) Set(val *CampaignCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignCreation(val *CampaignCreation) *NullableCampaignCreation {
	return &NullableCampaignCreation{value: val, isSet: true}
}

func (v NullableCampaignCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


