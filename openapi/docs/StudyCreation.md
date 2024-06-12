# StudyCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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
**EventSet** | [**MeasurementEventSet**](MeasurementEventSet.md) |  | 
**ClientId** | Pointer to **string** | The ID of the measurement client that this study is for | [optional] 

## Methods

### NewStudyCreation

`func NewStudyCreation(name string, eventSet MeasurementEventSet, ) *StudyCreation`

NewStudyCreation instantiates a new StudyCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyCreationWithDefaults

`func NewStudyCreationWithDefaults() *StudyCreation`

NewStudyCreationWithDefaults instantiates a new StudyCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StudyCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudyCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudyCreation) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *StudyCreation) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StudyCreation) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StudyCreation) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StudyCreation) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *StudyCreation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StudyCreation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StudyCreation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StudyCreation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *StudyCreation) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StudyCreation) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StudyCreation) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StudyCreation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *StudyCreation) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StudyCreation) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StudyCreation) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StudyCreation) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *StudyCreation) GetLifeCycleStage() StudyLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *StudyCreation) GetLifeCycleStageOk() (*StudyLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *StudyCreation) SetLifeCycleStage(v StudyLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *StudyCreation) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *StudyCreation) GetIngestionStatus() StudyIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *StudyCreation) GetIngestionStatusOk() (*StudyIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *StudyCreation) SetIngestionStatus(v StudyIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *StudyCreation) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetSummaryStats

`func (o *StudyCreation) GetSummaryStats() StudySummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *StudyCreation) GetSummaryStatsOk() (*StudySummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *StudyCreation) SetSummaryStats(v StudySummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *StudyCreation) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetClient

`func (o *StudyCreation) GetClient() MeasurementClientItem`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *StudyCreation) GetClientOk() (*MeasurementClientItem, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *StudyCreation) SetClient(v MeasurementClientItem)`

SetClient sets Client field to given value.

### HasClient

`func (o *StudyCreation) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetEventLinks

`func (o *StudyCreation) GetEventLinks() MeasurementEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *StudyCreation) GetEventLinksOk() (*MeasurementEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *StudyCreation) SetEventLinks(v MeasurementEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *StudyCreation) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetEventCap

`func (o *StudyCreation) GetEventCap() int32`

GetEventCap returns the EventCap field if non-nil, zero value otherwise.

### GetEventCapOk

`func (o *StudyCreation) GetEventCapOk() (*int32, bool)`

GetEventCapOk returns a tuple with the EventCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCap

`func (o *StudyCreation) SetEventCap(v int32)`

SetEventCap sets EventCap field to given value.

### HasEventCap

`func (o *StudyCreation) HasEventCap() bool`

HasEventCap returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *StudyCreation) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *StudyCreation) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *StudyCreation) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *StudyCreation) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *StudyCreation) GetIntegrationPlatform() MeasurementIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *StudyCreation) GetIntegrationPlatformOk() (*MeasurementIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *StudyCreation) SetIntegrationPlatform(v MeasurementIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *StudyCreation) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.

### GetIsExample

`func (o *StudyCreation) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *StudyCreation) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *StudyCreation) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *StudyCreation) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.

### GetEventSet

`func (o *StudyCreation) GetEventSet() MeasurementEventSet`

GetEventSet returns the EventSet field if non-nil, zero value otherwise.

### GetEventSetOk

`func (o *StudyCreation) GetEventSetOk() (*MeasurementEventSet, bool)`

GetEventSetOk returns a tuple with the EventSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSet

`func (o *StudyCreation) SetEventSet(v MeasurementEventSet)`

SetEventSet sets EventSet field to given value.


### GetClientId

`func (o *StudyCreation) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StudyCreation) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StudyCreation) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StudyCreation) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


