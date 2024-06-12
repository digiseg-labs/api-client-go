# StudyPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Get** | **bool** | Can the current user GET (read) study information? | 
**Put** | **bool** | Can the current user PUT (update) study information? | 
**Delete** | **bool** | Can the current user DELETE the study? | 

## Methods

### NewStudyPermissions

`func NewStudyPermissions(get bool, put bool, delete bool, ) *StudyPermissions`

NewStudyPermissions instantiates a new StudyPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyPermissionsWithDefaults

`func NewStudyPermissionsWithDefaults() *StudyPermissions`

NewStudyPermissionsWithDefaults instantiates a new StudyPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGet

`func (o *StudyPermissions) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *StudyPermissions) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *StudyPermissions) SetGet(v bool)`

SetGet sets Get field to given value.


### GetPut

`func (o *StudyPermissions) GetPut() bool`

GetPut returns the Put field if non-nil, zero value otherwise.

### GetPutOk

`func (o *StudyPermissions) GetPutOk() (*bool, bool)`

GetPutOk returns a tuple with the Put field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPut

`func (o *StudyPermissions) SetPut(v bool)`

SetPut sets Put field to given value.


### GetDelete

`func (o *StudyPermissions) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *StudyPermissions) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *StudyPermissions) SetDelete(v bool)`

SetDelete sets Delete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


