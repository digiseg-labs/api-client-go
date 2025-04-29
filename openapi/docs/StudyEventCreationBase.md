# StudyEventCreationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** | The user agent of the event | [optional] 
**Referer** | Pointer to **string** | The referer value of the event | [optional] 
**EventTime** | Pointer to **time.Time** | Optionally, the time of the event | [optional] 
**EventType** | Pointer to **string** | The event type to ingest, typically &#x60;impression&#x60; or &#x60;click&#x60; | [optional] 

## Methods

### NewStudyEventCreationBase

`func NewStudyEventCreationBase() *StudyEventCreationBase`

NewStudyEventCreationBase instantiates a new StudyEventCreationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyEventCreationBaseWithDefaults

`func NewStudyEventCreationBaseWithDefaults() *StudyEventCreationBase`

NewStudyEventCreationBaseWithDefaults instantiates a new StudyEventCreationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *StudyEventCreationBase) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StudyEventCreationBase) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StudyEventCreationBase) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StudyEventCreationBase) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *StudyEventCreationBase) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *StudyEventCreationBase) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *StudyEventCreationBase) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *StudyEventCreationBase) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetReferer

`func (o *StudyEventCreationBase) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *StudyEventCreationBase) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *StudyEventCreationBase) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *StudyEventCreationBase) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetEventTime

`func (o *StudyEventCreationBase) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *StudyEventCreationBase) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *StudyEventCreationBase) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *StudyEventCreationBase) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *StudyEventCreationBase) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StudyEventCreationBase) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StudyEventCreationBase) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StudyEventCreationBase) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


