# StudyEventCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]StudyEventCreationBase**](StudyEventCreationBase.md) |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** | The user agent of the event | [optional] 
**Referer** | Pointer to **string** | The referer value of the event | [optional] 
**EventTime** | Pointer to **time.Time** | Optionally, the time of the event | [optional] 
**EventType** | Pointer to **string** | The event type to ingest, typically &#x60;impression&#x60; or &#x60;click&#x60; | [optional] 

## Methods

### NewStudyEventCreation

`func NewStudyEventCreation() *StudyEventCreation`

NewStudyEventCreation instantiates a new StudyEventCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyEventCreationWithDefaults

`func NewStudyEventCreationWithDefaults() *StudyEventCreation`

NewStudyEventCreationWithDefaults instantiates a new StudyEventCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *StudyEventCreation) GetEvents() []StudyEventCreationBase`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StudyEventCreation) GetEventsOk() (*[]StudyEventCreationBase, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StudyEventCreation) SetEvents(v []StudyEventCreationBase)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StudyEventCreation) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIpAddress

`func (o *StudyEventCreation) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StudyEventCreation) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StudyEventCreation) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StudyEventCreation) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *StudyEventCreation) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *StudyEventCreation) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *StudyEventCreation) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *StudyEventCreation) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetReferer

`func (o *StudyEventCreation) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *StudyEventCreation) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *StudyEventCreation) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *StudyEventCreation) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetEventTime

`func (o *StudyEventCreation) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *StudyEventCreation) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *StudyEventCreation) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *StudyEventCreation) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *StudyEventCreation) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StudyEventCreation) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StudyEventCreation) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StudyEventCreation) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


