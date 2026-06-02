# AudienceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to represent the audience | 
**DisplayName** | **string** |  | 
**Description** | Pointer to **string** | A textual description of the audience | [optional] 
**Identifier** | Pointer to **string** | A single-letter identifier for the audience (e.g., A, B, C) used for simplified display | [optional] 
**PlatformCode** | Pointer to **string** | Platform-specific code, provided when &#x60;platform&#x60; (and &#x60;country&#x60; if needed for the platform) is provided | [optional] 
**Labels** | **[]string** | A list of labels associated with the audience | 
**ReachStats** | Pointer to [**AudienceReachStats**](AudienceReachStats.md) |  | [optional] 
**Category** | Pointer to [**AudienceCategoryRef**](AudienceCategoryRef.md) |  | [optional] 

## Methods

### NewAudienceItem

`func NewAudienceItem(code string, displayName string, labels []string, ) *AudienceItem`

NewAudienceItem instantiates a new AudienceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceItemWithDefaults

`func NewAudienceItemWithDefaults() *AudienceItem`

NewAudienceItemWithDefaults instantiates a new AudienceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AudienceItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudienceItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudienceItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *AudienceItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AudienceItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AudienceItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *AudienceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AudienceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AudienceItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AudienceItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *AudienceItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AudienceItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AudienceItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AudienceItem) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPlatformCode

`func (o *AudienceItem) GetPlatformCode() string`

GetPlatformCode returns the PlatformCode field if non-nil, zero value otherwise.

### GetPlatformCodeOk

`func (o *AudienceItem) GetPlatformCodeOk() (*string, bool)`

GetPlatformCodeOk returns a tuple with the PlatformCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCode

`func (o *AudienceItem) SetPlatformCode(v string)`

SetPlatformCode sets PlatformCode field to given value.

### HasPlatformCode

`func (o *AudienceItem) HasPlatformCode() bool`

HasPlatformCode returns a boolean if a field has been set.

### GetLabels

`func (o *AudienceItem) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AudienceItem) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AudienceItem) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetReachStats

`func (o *AudienceItem) GetReachStats() AudienceReachStats`

GetReachStats returns the ReachStats field if non-nil, zero value otherwise.

### GetReachStatsOk

`func (o *AudienceItem) GetReachStatsOk() (*AudienceReachStats, bool)`

GetReachStatsOk returns a tuple with the ReachStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachStats

`func (o *AudienceItem) SetReachStats(v AudienceReachStats)`

SetReachStats sets ReachStats field to given value.

### HasReachStats

`func (o *AudienceItem) HasReachStats() bool`

HasReachStats returns a boolean if a field has been set.

### GetCategory

`func (o *AudienceItem) GetCategory() AudienceCategoryRef`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AudienceItem) GetCategoryOk() (*AudienceCategoryRef, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AudienceItem) SetCategory(v AudienceCategoryRef)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AudienceItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


