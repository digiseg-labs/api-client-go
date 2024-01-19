# CampaignEventLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The base URI of the link to an event | [optional] 
**Parameters** | Pointer to [**map[string]CampaignEventLinkParameterInfo**](CampaignEventLinkParameterInfo.md) | Describes any parameters that can be added to the event link | [optional] 

## Methods

### NewCampaignEventLink

`func NewCampaignEventLink() *CampaignEventLink`

NewCampaignEventLink instantiates a new CampaignEventLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEventLinkWithDefaults

`func NewCampaignEventLinkWithDefaults() *CampaignEventLink`

NewCampaignEventLinkWithDefaults instantiates a new CampaignEventLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *CampaignEventLink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CampaignEventLink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CampaignEventLink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CampaignEventLink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetParameters

`func (o *CampaignEventLink) GetParameters() map[string]CampaignEventLinkParameterInfo`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CampaignEventLink) GetParametersOk() (*map[string]CampaignEventLinkParameterInfo, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CampaignEventLink) SetParameters(v map[string]CampaignEventLinkParameterInfo)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CampaignEventLink) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


