# EnrichmentJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadLink** | Pointer to **string** | Link to the downloadable result (if available) | [optional] 
**DownloadAvailable** | **bool** | Indicates whether the downloadable result is available | 
**AvailableUntil** | Pointer to **time.Time** | Date and time for the downloadable result (data is automatically deleted after this time) | [optional] 
**CountPrivate** | Pointer to **int32** |  | [optional] 
**CountBusiness** | Pointer to **int32** |  | [optional] 
**CountTotal** | Pointer to **int32** |  | [optional] 
**CountNotResolved** | Pointer to **int32** |  | [optional] 
**CountSkipped** | Pointer to **int32** |  | [optional] 
**CountOutOfQuota** | Pointer to **int32** |  | [optional] 

## Methods

### NewEnrichmentJobResult

`func NewEnrichmentJobResult(downloadAvailable bool, ) *EnrichmentJobResult`

NewEnrichmentJobResult instantiates a new EnrichmentJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobResultWithDefaults

`func NewEnrichmentJobResultWithDefaults() *EnrichmentJobResult`

NewEnrichmentJobResultWithDefaults instantiates a new EnrichmentJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadLink

`func (o *EnrichmentJobResult) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *EnrichmentJobResult) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *EnrichmentJobResult) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *EnrichmentJobResult) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetDownloadAvailable

`func (o *EnrichmentJobResult) GetDownloadAvailable() bool`

GetDownloadAvailable returns the DownloadAvailable field if non-nil, zero value otherwise.

### GetDownloadAvailableOk

`func (o *EnrichmentJobResult) GetDownloadAvailableOk() (*bool, bool)`

GetDownloadAvailableOk returns a tuple with the DownloadAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadAvailable

`func (o *EnrichmentJobResult) SetDownloadAvailable(v bool)`

SetDownloadAvailable sets DownloadAvailable field to given value.


### GetAvailableUntil

`func (o *EnrichmentJobResult) GetAvailableUntil() time.Time`

GetAvailableUntil returns the AvailableUntil field if non-nil, zero value otherwise.

### GetAvailableUntilOk

`func (o *EnrichmentJobResult) GetAvailableUntilOk() (*time.Time, bool)`

GetAvailableUntilOk returns a tuple with the AvailableUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUntil

`func (o *EnrichmentJobResult) SetAvailableUntil(v time.Time)`

SetAvailableUntil sets AvailableUntil field to given value.

### HasAvailableUntil

`func (o *EnrichmentJobResult) HasAvailableUntil() bool`

HasAvailableUntil returns a boolean if a field has been set.

### GetCountPrivate

`func (o *EnrichmentJobResult) GetCountPrivate() int32`

GetCountPrivate returns the CountPrivate field if non-nil, zero value otherwise.

### GetCountPrivateOk

`func (o *EnrichmentJobResult) GetCountPrivateOk() (*int32, bool)`

GetCountPrivateOk returns a tuple with the CountPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountPrivate

`func (o *EnrichmentJobResult) SetCountPrivate(v int32)`

SetCountPrivate sets CountPrivate field to given value.

### HasCountPrivate

`func (o *EnrichmentJobResult) HasCountPrivate() bool`

HasCountPrivate returns a boolean if a field has been set.

### GetCountBusiness

`func (o *EnrichmentJobResult) GetCountBusiness() int32`

GetCountBusiness returns the CountBusiness field if non-nil, zero value otherwise.

### GetCountBusinessOk

`func (o *EnrichmentJobResult) GetCountBusinessOk() (*int32, bool)`

GetCountBusinessOk returns a tuple with the CountBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountBusiness

`func (o *EnrichmentJobResult) SetCountBusiness(v int32)`

SetCountBusiness sets CountBusiness field to given value.

### HasCountBusiness

`func (o *EnrichmentJobResult) HasCountBusiness() bool`

HasCountBusiness returns a boolean if a field has been set.

### GetCountTotal

`func (o *EnrichmentJobResult) GetCountTotal() int32`

GetCountTotal returns the CountTotal field if non-nil, zero value otherwise.

### GetCountTotalOk

`func (o *EnrichmentJobResult) GetCountTotalOk() (*int32, bool)`

GetCountTotalOk returns a tuple with the CountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountTotal

`func (o *EnrichmentJobResult) SetCountTotal(v int32)`

SetCountTotal sets CountTotal field to given value.

### HasCountTotal

`func (o *EnrichmentJobResult) HasCountTotal() bool`

HasCountTotal returns a boolean if a field has been set.

### GetCountNotResolved

`func (o *EnrichmentJobResult) GetCountNotResolved() int32`

GetCountNotResolved returns the CountNotResolved field if non-nil, zero value otherwise.

### GetCountNotResolvedOk

`func (o *EnrichmentJobResult) GetCountNotResolvedOk() (*int32, bool)`

GetCountNotResolvedOk returns a tuple with the CountNotResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountNotResolved

`func (o *EnrichmentJobResult) SetCountNotResolved(v int32)`

SetCountNotResolved sets CountNotResolved field to given value.

### HasCountNotResolved

`func (o *EnrichmentJobResult) HasCountNotResolved() bool`

HasCountNotResolved returns a boolean if a field has been set.

### GetCountSkipped

`func (o *EnrichmentJobResult) GetCountSkipped() int32`

GetCountSkipped returns the CountSkipped field if non-nil, zero value otherwise.

### GetCountSkippedOk

`func (o *EnrichmentJobResult) GetCountSkippedOk() (*int32, bool)`

GetCountSkippedOk returns a tuple with the CountSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountSkipped

`func (o *EnrichmentJobResult) SetCountSkipped(v int32)`

SetCountSkipped sets CountSkipped field to given value.

### HasCountSkipped

`func (o *EnrichmentJobResult) HasCountSkipped() bool`

HasCountSkipped returns a boolean if a field has been set.

### GetCountOutOfQuota

`func (o *EnrichmentJobResult) GetCountOutOfQuota() int32`

GetCountOutOfQuota returns the CountOutOfQuota field if non-nil, zero value otherwise.

### GetCountOutOfQuotaOk

`func (o *EnrichmentJobResult) GetCountOutOfQuotaOk() (*int32, bool)`

GetCountOutOfQuotaOk returns a tuple with the CountOutOfQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOutOfQuota

`func (o *EnrichmentJobResult) SetCountOutOfQuota(v int32)`

SetCountOutOfQuota sets CountOutOfQuota field to given value.

### HasCountOutOfQuota

`func (o *EnrichmentJobResult) HasCountOutOfQuota() bool`

HasCountOutOfQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


