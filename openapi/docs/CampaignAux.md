# CampaignAux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLinks** | Pointer to [**CampaignEventLinks**](CampaignEventLinks.md) |  | [optional] 
**BannerImageUrl** | Pointer to **string** | The URL to a banner image for the campaign. Note that the banner image is used only for Digiseg campaign reporting and presentation, it does NOT represent any delivered banner ad creatives or similar.  | [optional] [readonly] 
**IntegrationPlatform** | Pointer to [**CampaignIntegrationPlatform**](CampaignIntegrationPlatform.md) |  | [optional] 

## Methods

### NewCampaignAux

`func NewCampaignAux() *CampaignAux`

NewCampaignAux instantiates a new CampaignAux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignAuxWithDefaults

`func NewCampaignAuxWithDefaults() *CampaignAux`

NewCampaignAuxWithDefaults instantiates a new CampaignAux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLinks

`func (o *CampaignAux) GetEventLinks() CampaignEventLinks`

GetEventLinks returns the EventLinks field if non-nil, zero value otherwise.

### GetEventLinksOk

`func (o *CampaignAux) GetEventLinksOk() (*CampaignEventLinks, bool)`

GetEventLinksOk returns a tuple with the EventLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLinks

`func (o *CampaignAux) SetEventLinks(v CampaignEventLinks)`

SetEventLinks sets EventLinks field to given value.

### HasEventLinks

`func (o *CampaignAux) HasEventLinks() bool`

HasEventLinks returns a boolean if a field has been set.

### GetBannerImageUrl

`func (o *CampaignAux) GetBannerImageUrl() string`

GetBannerImageUrl returns the BannerImageUrl field if non-nil, zero value otherwise.

### GetBannerImageUrlOk

`func (o *CampaignAux) GetBannerImageUrlOk() (*string, bool)`

GetBannerImageUrlOk returns a tuple with the BannerImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImageUrl

`func (o *CampaignAux) SetBannerImageUrl(v string)`

SetBannerImageUrl sets BannerImageUrl field to given value.

### HasBannerImageUrl

`func (o *CampaignAux) HasBannerImageUrl() bool`

HasBannerImageUrl returns a boolean if a field has been set.

### GetIntegrationPlatform

`func (o *CampaignAux) GetIntegrationPlatform() CampaignIntegrationPlatform`

GetIntegrationPlatform returns the IntegrationPlatform field if non-nil, zero value otherwise.

### GetIntegrationPlatformOk

`func (o *CampaignAux) GetIntegrationPlatformOk() (*CampaignIntegrationPlatform, bool)`

GetIntegrationPlatformOk returns a tuple with the IntegrationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPlatform

`func (o *CampaignAux) SetIntegrationPlatform(v CampaignIntegrationPlatform)`

SetIntegrationPlatform sets IntegrationPlatform field to given value.

### HasIntegrationPlatform

`func (o *CampaignAux) HasIntegrationPlatform() bool`

HasIntegrationPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


