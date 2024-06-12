# StudyMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**StudyPermissions**](StudyPermissions.md) |  | [optional] 

## Methods

### NewStudyMeta

`func NewStudyMeta() *StudyMeta`

NewStudyMeta instantiates a new StudyMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyMetaWithDefaults

`func NewStudyMetaWithDefaults() *StudyMeta`

NewStudyMetaWithDefaults instantiates a new StudyMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *StudyMeta) GetPermissions() StudyPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *StudyMeta) GetPermissionsOk() (*StudyPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *StudyMeta) SetPermissions(v StudyPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *StudyMeta) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


