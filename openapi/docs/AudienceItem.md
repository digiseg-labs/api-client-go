# AudienceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to represent the audience | 
**DisplayName** | **string** |  | 
**Description** | Pointer to **string** | A textual description of the audience | [optional] 
**Composition** | Pointer to **[]string** | A list of audience codes that this audience is composed from | [optional] 
**PlatformCode** | Pointer to **string** | Platform-specific code, provided when &#x60;platform&#x60; (and &#x60;country&#x60; if needed for the platform) is provided | [optional] 
**ReachStats** | Pointer to [**AudienceReachStats**](AudienceReachStats.md) |  | [optional] 

## Methods

### NewAudienceItem

`func NewAudienceItem(code string, displayName string, ) *AudienceItem`

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

### GetComposition

`func (o *AudienceItem) GetComposition() []string`

GetComposition returns the Composition field if non-nil, zero value otherwise.

### GetCompositionOk

`func (o *AudienceItem) GetCompositionOk() (*[]string, bool)`

GetCompositionOk returns a tuple with the Composition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposition

`func (o *AudienceItem) SetComposition(v []string)`

SetComposition sets Composition field to given value.

### HasComposition

`func (o *AudienceItem) HasComposition() bool`

HasComposition returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


