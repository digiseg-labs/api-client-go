# CampaignCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Labels** | Pointer to **[]string** | A set of labels that users can use to categorize their campaigns. Can be used to indicate type of campaign, customer names or other traits.  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this campaign | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date for which the campaign and its data ingestion will start. | [optional] [readonly] 
**LifeCycleStage** | Pointer to [**CampaignLifecycleStage**](CampaignLifecycleStage.md) |  | [optional] 
**IngestionStatus** | Pointer to [**CampaignIngestionStatus**](CampaignIngestionStatus.md) |  | [optional] 
**EventSet** | Pointer to [**CampaignEventSet**](CampaignEventSet.md) |  | [optional] 

## Methods

### NewCampaignCreation

`func NewCampaignCreation(name string, ) *CampaignCreation`

NewCampaignCreation instantiates a new CampaignCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreationWithDefaults

`func NewCampaignCreationWithDefaults() *CampaignCreation`

NewCampaignCreationWithDefaults instantiates a new CampaignCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCreation) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *CampaignCreation) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CampaignCreation) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CampaignCreation) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CampaignCreation) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *CampaignCreation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignCreation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignCreation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CampaignCreation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignCreation) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignCreation) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignCreation) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignCreation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *CampaignCreation) GetLifeCycleStage() CampaignLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *CampaignCreation) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *CampaignCreation) SetLifeCycleStage(v CampaignLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *CampaignCreation) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *CampaignCreation) GetIngestionStatus() CampaignIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *CampaignCreation) GetIngestionStatusOk() (*CampaignIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *CampaignCreation) SetIngestionStatus(v CampaignIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *CampaignCreation) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.

### GetEventSet

`func (o *CampaignCreation) GetEventSet() CampaignEventSet`

GetEventSet returns the EventSet field if non-nil, zero value otherwise.

### GetEventSetOk

`func (o *CampaignCreation) GetEventSetOk() (*CampaignEventSet, bool)`

GetEventSetOk returns a tuple with the EventSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSet

`func (o *CampaignCreation) SetEventSet(v CampaignEventSet)`

SetEventSet sets EventSet field to given value.

### HasEventSet

`func (o *CampaignCreation) HasEventSet() bool`

HasEventSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

