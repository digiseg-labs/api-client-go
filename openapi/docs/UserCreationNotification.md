# UserCreationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyUser** | Pointer to **bool** | Whether or not to notify the user that they have been registered | [optional] [default to true]

## Methods

### NewUserCreationNotification

`func NewUserCreationNotification() *UserCreationNotification`

NewUserCreationNotification instantiates a new UserCreationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreationNotificationWithDefaults

`func NewUserCreationNotificationWithDefaults() *UserCreationNotification`

NewUserCreationNotificationWithDefaults instantiates a new UserCreationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyUser

`func (o *UserCreationNotification) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *UserCreationNotification) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *UserCreationNotification) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *UserCreationNotification) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


