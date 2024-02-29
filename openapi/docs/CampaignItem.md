# CampaignItem

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

## Methods

### NewCampaignItem

`func NewCampaignItem(id string, ) *CampaignItem`

NewCampaignItem instantiates a new CampaignItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemWithDefaults

`func NewCampaignItemWithDefaults() *CampaignItem`

NewCampaignItemWithDefaults instantiates a new CampaignItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CampaignItem) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CampaignItem) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CampaignItem) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CampaignItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAccountId

`func (o *CampaignItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CampaignItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *CampaignItem) GetLifeCycleStage() CampaignLifecycleStage`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *CampaignItem) GetLifeCycleStageOk() (*CampaignLifecycleStage, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *CampaignItem) SetLifeCycleStage(v CampaignLifecycleStage)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *CampaignItem) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetIngestionStatus

`func (o *CampaignItem) GetIngestionStatus() CampaignIngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *CampaignItem) GetIngestionStatusOk() (*CampaignIngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *CampaignItem) SetIngestionStatus(v CampaignIngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.

### HasIngestionStatus

`func (o *CampaignItem) HasIngestionStatus() bool`

HasIngestionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


