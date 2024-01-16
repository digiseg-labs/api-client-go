# ResolveAudiencesOfMultipleResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier (if provided) of the item as it was provided in the request. | [optional] 
**Audiences** | Pointer to [**[]Audience**](Audience.md) |  | [optional] 
**Status** | [**AudienceResponseStatus**](AudienceResponseStatus.md) |  | 

## Methods

### NewResolveAudiencesOfMultipleResponseItem

`func NewResolveAudiencesOfMultipleResponseItem(status AudienceResponseStatus, ) *ResolveAudiencesOfMultipleResponseItem`

NewResolveAudiencesOfMultipleResponseItem instantiates a new ResolveAudiencesOfMultipleResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveAudiencesOfMultipleResponseItemWithDefaults

`func NewResolveAudiencesOfMultipleResponseItemWithDefaults() *ResolveAudiencesOfMultipleResponseItem`

NewResolveAudiencesOfMultipleResponseItemWithDefaults instantiates a new ResolveAudiencesOfMultipleResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResolveAudiencesOfMultipleResponseItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResolveAudiencesOfMultipleResponseItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResolveAudiencesOfMultipleResponseItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResolveAudiencesOfMultipleResponseItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAudiences

`func (o *ResolveAudiencesOfMultipleResponseItem) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ResolveAudiencesOfMultipleResponseItem) GetAudiencesOk() (*[]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ResolveAudiencesOfMultipleResponseItem) SetAudiences(v []Audience)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *ResolveAudiencesOfMultipleResponseItem) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetStatus

`func (o *ResolveAudiencesOfMultipleResponseItem) GetStatus() AudienceResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResolveAudiencesOfMultipleResponseItem) GetStatusOk() (*AudienceResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResolveAudiencesOfMultipleResponseItem) SetStatus(v AudienceResponseStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


