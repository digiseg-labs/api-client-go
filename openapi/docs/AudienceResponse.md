# AudienceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to [**[]Audience**](Audience.md) |  | [optional] 
**Status** | [**AudienceResponseStatus**](AudienceResponseStatus.md) |  | 

## Methods

### NewAudienceResponse

`func NewAudienceResponse(status AudienceResponseStatus, ) *AudienceResponse`

NewAudienceResponse instantiates a new AudienceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceResponseWithDefaults

`func NewAudienceResponseWithDefaults() *AudienceResponse`

NewAudienceResponseWithDefaults instantiates a new AudienceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *AudienceResponse) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *AudienceResponse) GetAudiencesOk() (*[]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *AudienceResponse) SetAudiences(v []Audience)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *AudienceResponse) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetStatus

`func (o *AudienceResponse) GetStatus() AudienceResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudienceResponse) GetStatusOk() (*AudienceResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudienceResponse) SetStatus(v AudienceResponseStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


