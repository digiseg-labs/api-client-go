# StudyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | A set of labels that users can use to categorize their measurements. Can be used to indicate type of study, customer names or other traits.  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this study | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date for which the study and its data ingestion will start. | [optional] [readonly] 
**LifeCycleStage** | Pointer to [**StudyLifecycleStage**](StudyLifecycleStage.md) |  | [optional] 
**IngestionStatus** | Pointer to [**StudyIngestionStatus**](StudyIngestionStatus.md) |  | [optional] 
**SummaryStats** | Pointer to [**StudySummaryStats**](StudySummaryStats.md) |  | [optional] 
**Client** | Pointer to [**MeasurementClientItem**](MeasurementClientItem.md) |  | [optional] 

## Methods

### NewStudyItem

`func NewStudyItem(id string, ) *StudyItem`

NewStudyItem instantiates a new StudyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyItemWithDefaults

`func NewStudyItemWithDefaults() *StudyItem`

NewStudyItemWithDefaults instantiates a new StudyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StudyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudyItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StudyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudyItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudyItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *StudyItem) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StudyItem) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StudyItem) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StudyItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *StudyItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StudyItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StudyItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StudyItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *StudyItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StudyItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StudyItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StudyItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *StudyItem) GetLifeCycleStage() StudyLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *StudyItem) GetLifeCycleStageOk() (*StudyLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *StudyItem) SetLifeCycleStage(v StudyLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *StudyItem) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *StudyItem) GetIngestionStatus() StudyIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *StudyItem) GetIngestionStatusOk() (*StudyIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *StudyItem) SetIngestionStatus(v StudyIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *StudyItem) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetSummaryStats

`func (o *StudyItem) GetSummaryStats() StudySummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *StudyItem) GetSummaryStatsOk() (*StudySummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *StudyItem) SetSummaryStats(v StudySummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *StudyItem) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetClient

`func (o *StudyItem) GetClient() MeasurementClientItem`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *StudyItem) GetClientOk() (*MeasurementClientItem, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *StudyItem) SetClient(v MeasurementClientItem)`

SetClient sets Client field to given value.

### HasClient

`func (o *StudyItem) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


