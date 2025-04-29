# SharedReportCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** | URL of the publicly available asset | [optional] [readonly] 
**ReportType** | [**SharedReportType**](SharedReportType.md) |  | 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the shared report will expire. Default is 28 days after the time of creation. | [optional] 
**StudyId** | Pointer to **string** | ID of the study to which the report belongs | [optional] [readonly] 

## Methods

### NewSharedReportCreation

`func NewSharedReportCreation(reportType SharedReportType, ) *SharedReportCreation`

NewSharedReportCreation instantiates a new SharedReportCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedReportCreationWithDefaults

`func NewSharedReportCreationWithDefaults() *SharedReportCreation`

NewSharedReportCreationWithDefaults instantiates a new SharedReportCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *SharedReportCreation) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *SharedReportCreation) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *SharedReportCreation) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *SharedReportCreation) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetReportType

`func (o *SharedReportCreation) GetReportType() SharedReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *SharedReportCreation) GetReportTypeOk() (*SharedReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *SharedReportCreation) SetReportType(v SharedReportType)`

SetReportType sets ReportType field to given value.


### GetExpiresAt

`func (o *SharedReportCreation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SharedReportCreation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SharedReportCreation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SharedReportCreation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStudyId

`func (o *SharedReportCreation) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *SharedReportCreation) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *SharedReportCreation) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *SharedReportCreation) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


