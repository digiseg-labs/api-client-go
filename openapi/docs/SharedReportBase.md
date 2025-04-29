# SharedReportBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** | URL of the publicly available asset | [optional] [readonly] 
**ReportType** | [**SharedReportType**](SharedReportType.md) |  | 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the shared report will expire. Default is 28 days after the time of creation. | [optional] 
**StudyId** | Pointer to **string** | ID of the study to which the report belongs | [optional] [readonly] 

## Methods

### NewSharedReportBase

`func NewSharedReportBase(reportType SharedReportType, ) *SharedReportBase`

NewSharedReportBase instantiates a new SharedReportBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedReportBaseWithDefaults

`func NewSharedReportBaseWithDefaults() *SharedReportBase`

NewSharedReportBaseWithDefaults instantiates a new SharedReportBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *SharedReportBase) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *SharedReportBase) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *SharedReportBase) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *SharedReportBase) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetReportType

`func (o *SharedReportBase) GetReportType() SharedReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *SharedReportBase) GetReportTypeOk() (*SharedReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *SharedReportBase) SetReportType(v SharedReportType)`

SetReportType sets ReportType field to given value.


### GetExpiresAt

`func (o *SharedReportBase) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SharedReportBase) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SharedReportBase) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SharedReportBase) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStudyId

`func (o *SharedReportBase) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *SharedReportBase) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *SharedReportBase) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *SharedReportBase) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


