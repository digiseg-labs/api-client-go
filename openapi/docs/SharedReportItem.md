# SharedReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
**AssetUrl** | Pointer to **string** | URL of the publicly available asset | [optional] [readonly] 
**ReportType** | [**SharedReportType**](SharedReportType.md) |  | 
**ExpiresAt** | Pointer to **time.Time** | Optional date/time that the shared report will expire. Default is 28 days after the time of creation. | [optional] 
**StudyId** | Pointer to **string** | ID of the study to which the report belongs | [optional] [readonly] 

## Methods

### NewSharedReportItem

`func NewSharedReportItem(id string, reportType SharedReportType, ) *SharedReportItem`

NewSharedReportItem instantiates a new SharedReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedReportItemWithDefaults

`func NewSharedReportItemWithDefaults() *SharedReportItem`

NewSharedReportItemWithDefaults instantiates a new SharedReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedReportItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedReportItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedReportItem) SetId(v string)`

SetId sets Id field to given value.


### GetAssetUrl

`func (o *SharedReportItem) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *SharedReportItem) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *SharedReportItem) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *SharedReportItem) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetReportType

`func (o *SharedReportItem) GetReportType() SharedReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *SharedReportItem) GetReportTypeOk() (*SharedReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *SharedReportItem) SetReportType(v SharedReportType)`

SetReportType sets ReportType field to given value.


### GetExpiresAt

`func (o *SharedReportItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SharedReportItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SharedReportItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SharedReportItem) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStudyId

`func (o *SharedReportItem) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *SharedReportItem) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *SharedReportItem) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *SharedReportItem) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


