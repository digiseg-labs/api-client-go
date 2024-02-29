# CampaignIntegrationPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An ID for the integration platform, if the integration platform is a known platform to Digiseg. Will be null/omitted if the platform name was entered manually as a string.  | [optional] 
**Name** | Pointer to **string** | The name of the integration platform. | [optional] 

## Methods

### NewCampaignIntegrationPlatform

`func NewCampaignIntegrationPlatform() *CampaignIntegrationPlatform`

NewCampaignIntegrationPlatform instantiates a new CampaignIntegrationPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignIntegrationPlatformWithDefaults

`func NewCampaignIntegrationPlatformWithDefaults() *CampaignIntegrationPlatform`

NewCampaignIntegrationPlatformWithDefaults instantiates a new CampaignIntegrationPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignIntegrationPlatform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignIntegrationPlatform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignIntegrationPlatform) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignIntegrationPlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignIntegrationPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignIntegrationPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignIntegrationPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignIntegrationPlatform) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


