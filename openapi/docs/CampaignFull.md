# CampaignFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
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
**CreatedAt** | **time.Time** | Date and time of the object creation | 
**CreatedBy** | **string** | ID of the user who created the object | 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] 

## Methods

### NewCampaignFull

`func NewCampaignFull(id string, createdAt time.Time, createdBy string, ) *CampaignFull`

NewCampaignFull instantiates a new CampaignFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignFullWithDefaults

`func NewCampaignFullWithDefaults() *CampaignFull`

NewCampaignFullWithDefaults instantiates a new CampaignFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignFull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CampaignFull) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CampaignFull) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CampaignFull) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CampaignFull) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *CampaignFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CampaignFull) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignFull) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignFull) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignFull) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignFull) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *CampaignFull) GetLifeCycleStage() CampaignLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *CampaignFull) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *CampaignFull) SetLifeCycleStage(v CampaignLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *CampaignFull) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *CampaignFull) GetIngestionStatus() CampaignIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *CampaignFull) GetIngestionStatusOk() (*CampaignIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *CampaignFull) SetIngestionStatus(v CampaignIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *CampaignFull) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetSummaryStats

`func (o *CampaignFull) GetSummaryStats() CampaignSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *CampaignFull) GetSummaryStatsOk() (*CampaignSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *CampaignFull) SetSummaryStats(v CampaignSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *CampaignFull) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetClient

`func (o *CampaignFull) GetClient() MeasurementClientItem`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CampaignFull) GetClientOk() (*MeasurementClientItem, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CampaignFull) SetClient(v MeasurementClientItem)`

SetClient sets Client field to given value.

### HasClient

`func (o *CampaignFull) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetEventLinks

`func (o *CampaignFull) GetEventLinks() CampaignEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *CampaignFull) GetEventLinksOk() (*CampaignEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *CampaignFull) SetEventLinks(v CampaignEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *CampaignFull) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *CampaignFull) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *CampaignFull) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *CampaignFull) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *CampaignFull) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *CampaignFull) GetIntegrationPlatform() CampaignIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *CampaignFull) GetIntegrationPlatformOk() (*CampaignIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *CampaignFull) SetIntegrationPlatform(v CampaignIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *CampaignFull) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CampaignFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *CampaignFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *CampaignFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CampaignFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CampaignFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CampaignFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CampaignFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


