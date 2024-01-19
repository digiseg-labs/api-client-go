# CampaignBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | A set of labels that users can use to categorize their campaigns. Can be used to indicate type of campaign, customer names or other traits.  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this campaign | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date for which the campaign and its data ingestion will start. | [optional] [readonly] 
**LifeCycleStage** | Pointer to [**CampaignLifecycleStage**](CampaignLifecycleStage.md) |  | [optional] 
**IngestionStatus** | Pointer to [**CampaignIngestionStatus**](CampaignIngestionStatus.md) |  | [optional] 

## Methods

### NewCampaignBase

`func NewCampaignBase() *CampaignBase`

NewCampaignBase instantiates a new CampaignBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignBaseWithDefaults

`func NewCampaignBaseWithDefaults() *CampaignBase`

NewCampaignBaseWithDefaults instantiates a new CampaignBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CampaignBase) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CampaignBase) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CampaignBase) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CampaignBase) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *CampaignBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CampaignBase) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignBase) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignBase) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignBase) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignBase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *CampaignBase) GetLifeCycleStage() CampaignLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *CampaignBase) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *CampaignBase) SetLifeCycleStage(v CampaignLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *CampaignBase) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *CampaignBase) GetIngestionStatus() CampaignIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *CampaignBase) GetIngestionStatusOk() (*CampaignIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *CampaignBase) SetIngestionStatus(v CampaignIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *CampaignBase) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


