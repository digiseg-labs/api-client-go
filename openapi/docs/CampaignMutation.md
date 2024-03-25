# CampaignMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | A set of labels that users can use to categorize their measurements. Can be used to indicate type of campaign, customer names or other traits.  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this campaign | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date for which the campaign and its data ingestion will start. | [optional] [readonly] 
**LifeCycleStage** | Pointer to [**CampaignLifecycleStage**](CampaignLifecycleStage.md) |  | [optional] 
**IngestionStatus** | Pointer to [**CampaignIngestionStatus**](CampaignIngestionStatus.md) |  | [optional] 
**SummaryStats** | Pointer to [**CampaignSummaryStats**](CampaignSummaryStats.md) |  | [optional] 
**Client** | Pointer to [**MeasurementClientItem**](MeasurementClientItem.md) |  | [optional] 
**EventLinks** | Pointer to [**CampaignEventLinks**](CampaignEventLinks.md) |  | [optional] 
**BannerImageUrl** | Pointer to **string** | The URL to a banner image for the campaign. Note that the banner image is used only for Digiseg campaign reporting and presentation, it does NOT represent any delivered banner ad creatives or similar.  | [optional] [readonly] 
**IntegrationPlatform** | Pointer to [**CampaignIntegrationPlatform**](CampaignIntegrationPlatform.md) |  | [optional] 

## Methods

### NewCampaignMutation

`func NewCampaignMutation() *CampaignMutation`

NewCampaignMutation instantiates a new CampaignMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMutationWithDefaults

`func NewCampaignMutationWithDefaults() *CampaignMutation`

NewCampaignMutationWithDefaults instantiates a new CampaignMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CampaignMutation) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CampaignMutation) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CampaignMutation) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CampaignMutation) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *CampaignMutation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignMutation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignMutation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CampaignMutation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignMutation) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignMutation) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignMutation) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignMutation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *CampaignMutation) GetLifeCycleStage() CampaignLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *CampaignMutation) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *CampaignMutation) SetLifeCycleStage(v CampaignLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *CampaignMutation) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *CampaignMutation) GetIngestionStatus() CampaignIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *CampaignMutation) GetIngestionStatusOk() (*CampaignIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *CampaignMutation) SetIngestionStatus(v CampaignIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *CampaignMutation) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetSummaryStats

`func (o *CampaignMutation) GetSummaryStats() CampaignSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *CampaignMutation) GetSummaryStatsOk() (*CampaignSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *CampaignMutation) SetSummaryStats(v CampaignSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *CampaignMutation) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetClient

`func (o *CampaignMutation) GetClient() MeasurementClientItem`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CampaignMutation) GetClientOk() (*MeasurementClientItem, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CampaignMutation) SetClient(v MeasurementClientItem)`

SetClient sets Client field to given value.

### HasClient

`func (o *CampaignMutation) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetEventLinks

`func (o *CampaignMutation) GetEventLinks() CampaignEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *CampaignMutation) GetEventLinksOk() (*CampaignEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *CampaignMutation) SetEventLinks(v CampaignEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *CampaignMutation) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *CampaignMutation) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *CampaignMutation) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *CampaignMutation) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *CampaignMutation) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *CampaignMutation) GetIntegrationPlatform() CampaignIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *CampaignMutation) GetIntegrationPlatformOk() (*CampaignIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *CampaignMutation) SetIntegrationPlatform(v CampaignIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *CampaignMutation) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


