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
)

// checks if the CampaignMutation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutation{}

// CampaignMutation struct for CampaignMutation
type CampaignMutation struct {
	Name *string `json:"name,omitempty"`
	// A set of labels that users can use to categorize their measurements. Can be used to indicate type of campaign, customer names or other traits. 
	Labels []string `json:"labels,omitempty"`
	// The ID of the account that owns this campaign
	AccountId *string `json:"account_id,omitempty"`
	// The date for which the campaign and its data ingestion will start.
	StartDate *time.Time `json:"start_date,omitempty"`
	LifeCycleStage *CampaignLifecycleStage `json:"life_cycle_stage,omitempty"`
	IngestionStatus *CampaignIngestionStatus `json:"ingestion_status,omitempty"`
	EventLinks *CampaignEventLinks `json:"event_links,omitempty"`
	IntegrationPlatform *CampaignIntegrationPlatform `json:"integration_platform,omitempty"`
	// The URL to a banner image for the campaign. Note that the banner image is used only for Digiseg campaign reporting and presentation, it does NOT represent any delivered banner ad creatives or similar. 
	BannerImageUrl *string `json:"banner_image_url,omitempty"`
	Client *MeasurementClientItem `json:"client,omitempty"`
}

// NewCampaignMutation instantiates a new CampaignMutation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutation() *CampaignMutation {
	this := CampaignMutation{}
	return &this
}

// NewCampaignMutationWithDefaults instantiates a new CampaignMutation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutationWithDefaults() *CampaignMutation {
	this := CampaignMutation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignMutation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignMutation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignMutation) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CampaignMutation) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CampaignMutation) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *CampaignMutation) SetLabels(v []string) {
	o.Labels = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CampaignMutation) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CampaignMutation) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CampaignMutation) SetAccountId(v string) {
	o.AccountId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CampaignMutation) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CampaignMutation) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CampaignMutation) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *CampaignMutation) GetLifeCycleStage() CampaignLifecycleStage {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret CampaignLifecycleStage
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *CampaignMutation) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given CampaignLifecycleStage and assigns it to the LifeCycleStage field.
func (o *CampaignMutation) SetLifeCycleStage(v CampaignLifecycleStage) {
	o.LifeCycleStage = &v
}

// GetIngestionStatus returns the IngestionStatus field value if set, zero value otherwise.
func (o *CampaignMutation) GetIngestionStatus() CampaignIngestionStatus {
	if o == nil || IsNil(o.IngestionStatus) {
		var ret CampaignIngestionStatus
		return ret
	}
	return *o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetIngestionStatusOk() (*CampaignIngestionStatus, bool) {
	if o == nil || IsNil(o.IngestionStatus) {
		return nil, false
	}
	return o.IngestionStatus, true
}

// HasIngestionStatus returns a boolean if a field has been set.
func (o *CampaignMutation) HasIngestionStatus() bool {
	if o != nil && !IsNil(o.IngestionStatus) {
		return true
	}

	return false
}

// SetIngestionStatus gets a reference to the given CampaignIngestionStatus and assigns it to the IngestionStatus field.
func (o *CampaignMutation) SetIngestionStatus(v CampaignIngestionStatus) {
	o.IngestionStatus = &v
}

// GetEventLinks returns the EventLinks field value if set, zero value otherwise.
func (o *CampaignMutation) GetEventLinks() CampaignEventLinks {
	if o == nil || IsNil(o.EventLinks) {
		var ret CampaignEventLinks
		return ret
	}
	return *o.EventLinks
}

// GetEventLinksOk returns a tuple with the EventLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetEventLinksOk() (*CampaignEventLinks, bool) {
	if o == nil || IsNil(o.EventLinks) {
		return nil, false
	}
	return o.EventLinks, true
}

// HasEventLinks returns a boolean if a field has been set.
func (o *CampaignMutation) HasEventLinks() bool {
	if o != nil && !IsNil(o.EventLinks) {
		return true
	}

	return false
}

// SetEventLinks gets a reference to the given CampaignEventLinks and assigns it to the EventLinks field.
func (o *CampaignMutation) SetEventLinks(v CampaignEventLinks) {
	o.EventLinks = &v
}

// GetIntegrationPlatform returns the IntegrationPlatform field value if set, zero value otherwise.
func (o *CampaignMutation) GetIntegrationPlatform() CampaignIntegrationPlatform {
	if o == nil || IsNil(o.IntegrationPlatform) {
		var ret CampaignIntegrationPlatform
		return ret
	}
	return *o.IntegrationPlatform
}

// GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetIntegrationPlatformOk() (*CampaignIntegrationPlatform, bool) {
	if o == nil || IsNil(o.IntegrationPlatform) {
		return nil, false
	}
	return o.IntegrationPlatform, true
}

// HasIntegrationPlatform returns a boolean if a field has been set.
func (o *CampaignMutation) HasIntegrationPlatform() bool {
	if o != nil && !IsNil(o.IntegrationPlatform) {
		return true
	}

	return false
}

// SetIntegrationPlatform gets a reference to the given CampaignIntegrationPlatform and assigns it to the IntegrationPlatform field.
func (o *CampaignMutation) SetIntegrationPlatform(v CampaignIntegrationPlatform) {
	o.IntegrationPlatform = &v
}

// GetBannerImageUrl returns the BannerImageUrl field value if set, zero value otherwise.
func (o *CampaignMutation) GetBannerImageUrl() string {
	if o == nil || IsNil(o.BannerImageUrl) {
		var ret string
		return ret
	}
	return *o.BannerImageUrl
}

// GetBannerImageUrlOk returns a tuple with the BannerImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetBannerImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImageUrl) {
		return nil, false
	}
	return o.BannerImageUrl, true
}

// HasBannerImageUrl returns a boolean if a field has been set.
func (o *CampaignMutation) HasBannerImageUrl() bool {
	if o != nil && !IsNil(o.BannerImageUrl) {
		return true
	}

	return false
}

// SetBannerImageUrl gets a reference to the given string and assigns it to the BannerImageUrl field.
func (o *CampaignMutation) SetBannerImageUrl(v string) {
	o.BannerImageUrl = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *CampaignMutation) GetClient() MeasurementClientItem {
	if o == nil || IsNil(o.Client) {
		var ret MeasurementClientItem
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutation) GetClientOk() (*MeasurementClientItem, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *CampaignMutation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given MeasurementClientItem and assigns it to the Client field.
func (o *CampaignMutation) SetClient(v MeasurementClientItem) {
	o.Client = &v
}

func (o CampaignMutation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignMutation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LifeCycleStage) {
		toSerialize["life_cycle_stage"] = o.LifeCycleStage
	}
	if !IsNil(o.IngestionStatus) {
		toSerialize["ingestion_status"] = o.IngestionStatus
	}
	if !IsNil(o.EventLinks) {
		toSerialize["event_links"] = o.EventLinks
	}
	if !IsNil(o.IntegrationPlatform) {
		toSerialize["integration_platform"] = o.IntegrationPlatform
	}
	if !IsNil(o.BannerImageUrl) {
		toSerialize["banner_image_url"] = o.BannerImageUrl
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	return toSerialize, nil
}

type NullableCampaignMutation struct {
	value *CampaignMutation
	isSet bool
}

func (v NullableCampaignMutation) Get() *CampaignMutation {
	return v.value
}

func (v *NullableCampaignMutation) Set(val *CampaignMutation) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutation) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutation(val *CampaignMutation) *NullableCampaignMutation {
	return &NullableCampaignMutation{value: val, isSet: true}
}

func (v NullableCampaignMutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMutation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


