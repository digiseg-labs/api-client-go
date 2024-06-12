# StudyAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLinks** | Pointer to [**MeasurementEventLinks**](MeasurementEventLinks.md) |  | [optional] 
**EventCap** | Pointer to **int32** | If present, an upper limit on the number of events that will be processed in this study. | [optional] 
**BannerImageUrl** | Pointer to **string** | The URL to a banner image for the study. Note that the banner image is used only for Digiseg study reporting and presentation, it does NOT represent any delivered banner ad creatives or similar.  | [optional] [readonly] 
**IntegrationPlatform** | Pointer to [**MeasurementIntegrationPlatform**](MeasurementIntegrationPlatform.md) |  | [optional] 
**IsExample** | Pointer to **bool** | Determines if the study is an example study, used to demonstrate product capabilities | [optional] [readonly] 

## Methods

### NewStudyAux

`func NewStudyAux() *StudyAux`

NewStudyAux instantiates a new StudyAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyAuxWithDefaults

`func NewStudyAuxWithDefaults() *StudyAux`

NewStudyAuxWithDefaults instantiates a new StudyAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLinks

`func (o *StudyAux) GetEventLinks() MeasurementEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *StudyAux) GetEventLinksOk() (*MeasurementEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *StudyAux) SetEventLinks(v MeasurementEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *StudyAux) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetEventCap

`func (o *StudyAux) GetEventCap() int32`

GetEventCap returns the EventCap field if non-nil, zero value otherwise.

### GetEventCapOk

`func (o *StudyAux) GetEventCapOk() (*int32, bool)`

GetEventCapOk returns a tuple with the EventCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCap

`func (o *StudyAux) SetEventCap(v int32)`

SetEventCap sets EventCap field to given value.

### HasEventCap

`func (o *StudyAux) HasEventCap() bool`

HasEventCap returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *StudyAux) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *StudyAux) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *StudyAux) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *StudyAux) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *StudyAux) GetIntegrationPlatform() MeasurementIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *StudyAux) GetIntegrationPlatformOk() (*MeasurementIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *StudyAux) SetIntegrationPlatform(v MeasurementIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *StudyAux) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.

### GetIsExample

`func (o *StudyAux) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *StudyAux) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *StudyAux) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *StudyAux) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


