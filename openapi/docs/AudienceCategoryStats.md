# AudienceCategoryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the audience category | 
**Audiences** | [**[]AudienceStats**](AudienceStats.md) |  | 

## Methods

### NewAudienceCategoryStats

`func NewAudienceCategoryStats(name string, audiences []AudienceStats, ) *AudienceCategoryStats`

NewAudienceCategoryStats instantiates a new AudienceCategoryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceCategoryStatsWithDefaults

`func NewAudienceCategoryStatsWithDefaults() *AudienceCategoryStats`

NewAudienceCategoryStatsWithDefaults instantiates a new AudienceCategoryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AudienceCategoryStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceCategoryStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceCategoryStats) SetName(v string)`

SetName sets Name field to given value.


### GetAudiences

`func (o *AudienceCategoryStats) GetAudiences() []AudienceStats`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *AudienceCategoryStats) GetAudiencesOk() (*[]AudienceStats, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *AudienceCategoryStats) SetAudiences(v []AudienceStats)`

SetAudiences sets Audiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


