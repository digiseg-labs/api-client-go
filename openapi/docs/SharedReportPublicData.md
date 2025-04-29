# SharedReportPublicData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | **string** |  | 
**AccessToken** | Pointer to **string** | Access token that can be used for API requests, if the user is not already authenticated | [optional] [readonly] 

## Methods

### NewSharedReportPublicData

`func NewSharedReportPublicData(studyId string, ) *SharedReportPublicData`

NewSharedReportPublicData instantiates a new SharedReportPublicData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedReportPublicDataWithDefaults

`func NewSharedReportPublicDataWithDefaults() *SharedReportPublicData`

NewSharedReportPublicDataWithDefaults instantiates a new SharedReportPublicData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *SharedReportPublicData) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *SharedReportPublicData) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *SharedReportPublicData) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.


### GetAccessToken

`func (o *SharedReportPublicData) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SharedReportPublicData) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SharedReportPublicData) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *SharedReportPublicData) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


