# StudyFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | A set of labels that users can use to categorize their measurements. Can be used to indicate type of study, customer names or other traits.  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this study | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date for which the study and its data ingestion will start | [optional] 
**EndDate** | Pointer to **time.Time** | The date for which the study and its data ingestion will end | [optional] 
**LifeCycleStage** | Pointer to [**StudyLifecycleStage**](StudyLifecycleStage.md) |  | [optional] 
**IngestionStatus** | Pointer to [**StudyIngestionStatus**](StudyIngestionStatus.md) |  | [optional] 
**SummaryStats** | Pointer to [**StudySummaryStats**](StudySummaryStats.md) |  | [optional] 
**Client** | Pointer to [**MeasurementClientItem**](MeasurementClientItem.md) |  | [optional] 
**EventLinks** | Pointer to [**MeasurementEventLinks**](MeasurementEventLinks.md) |  | [optional] 
**EventCap** | Pointer to **int32** | If present, an upper limit on the number of events that will be processed in this study. | [optional] 
**BannerImageUrl** | Pointer to **string** | The URL to a banner image for the study. Note that the banner image is used only for Digiseg study reporting and presentation, it does NOT represent any delivered banner ad creatives or similar.  | [optional] [readonly] 
**IntegrationPlatform** | Pointer to [**MeasurementIntegrationPlatform**](MeasurementIntegrationPlatform.md) |  | [optional] 
**IsExample** | Pointer to **bool** | Determines if the study is an example study, used to demonstrate product capabilities | [optional] [readonly] 
**CreatedAt** | **time.Time** | Date and time of the object creation | 
**CreatedBy** | **string** | ID of the user who created the object | 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] 

## Methods

### NewStudyFull

`func NewStudyFull(id string, createdAt time.Time, createdBy string, ) *StudyFull`

NewStudyFull instantiates a new StudyFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyFullWithDefaults

`func NewStudyFullWithDefaults() *StudyFull`

NewStudyFullWithDefaults instantiates a new StudyFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StudyFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudyFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudyFull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StudyFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudyFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudyFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudyFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *StudyFull) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StudyFull) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StudyFull) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StudyFull) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *StudyFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StudyFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StudyFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StudyFull) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *StudyFull) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StudyFull) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StudyFull) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StudyFull) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *StudyFull) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StudyFull) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StudyFull) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StudyFull) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *StudyFull) GetLifeCycleStage() StudyLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *StudyFull) GetLifeCycleStageOk() (*StudyLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *StudyFull) SetLifeCycleStage(v StudyLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *StudyFull) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *StudyFull) GetIngestionStatus() StudyIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *StudyFull) GetIngestionStatusOk() (*StudyIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *StudyFull) SetIngestionStatus(v StudyIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *StudyFull) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetSummaryStats

`func (o *StudyFull) GetSummaryStats() StudySummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *StudyFull) GetSummaryStatsOk() (*StudySummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *StudyFull) SetSummaryStats(v StudySummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *StudyFull) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetClient

`func (o *StudyFull) GetClient() MeasurementClientItem`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *StudyFull) GetClientOk() (*MeasurementClientItem, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *StudyFull) SetClient(v MeasurementClientItem)`

SetClient sets Client field to given value.

### HasClient

`func (o *StudyFull) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetEventLinks

`func (o *StudyFull) GetEventLinks() MeasurementEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *StudyFull) GetEventLinksOk() (*MeasurementEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *StudyFull) SetEventLinks(v MeasurementEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *StudyFull) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetEventCap

`func (o *StudyFull) GetEventCap() int32`

GetEventCap returns the EventCap field if non-nil, zero value otherwise.

### GetEventCapOk

`func (o *StudyFull) GetEventCapOk() (*int32, bool)`

GetEventCapOk returns a tuple with the EventCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCap

`func (o *StudyFull) SetEventCap(v int32)`

SetEventCap sets EventCap field to given value.

### HasEventCap

`func (o *StudyFull) HasEventCap() bool`

HasEventCap returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *StudyFull) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *StudyFull) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *StudyFull) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *StudyFull) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *StudyFull) GetIntegrationPlatform() MeasurementIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *StudyFull) GetIntegrationPlatformOk() (*MeasurementIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *StudyFull) SetIntegrationPlatform(v MeasurementIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *StudyFull) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.

### GetIsExample

`func (o *StudyFull) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *StudyFull) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *StudyFull) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *StudyFull) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StudyFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StudyFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StudyFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *StudyFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StudyFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StudyFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *StudyFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StudyFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StudyFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StudyFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *StudyFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *StudyFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *StudyFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *StudyFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


