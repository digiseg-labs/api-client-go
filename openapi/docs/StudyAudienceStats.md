# StudyAudienceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotResolved** | Pointer to [**MeasurementsContainer**](MeasurementsContainer.md) |  | [optional] 
**Business** | Pointer to [**BusinessAudienceStats**](BusinessAudienceStats.md) |  | [optional] 
**Private** | Pointer to [**PrivateAudienceStats**](PrivateAudienceStats.md) |  | [optional] 

## Methods

### NewStudyAudienceStats

`func NewStudyAudienceStats() *StudyAudienceStats`

NewStudyAudienceStats instantiates a new StudyAudienceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyAudienceStatsWithDefaults

`func NewStudyAudienceStatsWithDefaults() *StudyAudienceStats`

NewStudyAudienceStatsWithDefaults instantiates a new StudyAudienceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotResolved

`func (o *StudyAudienceStats) GetNotResolved() MeasurementsContainer`

GetNotResolved returns the NotResolved field if non-nil, zero value otherwise.

### GetNotResolvedOk

`func (o *StudyAudienceStats) GetNotResolvedOk() (*MeasurementsContainer, bool)`

GetNotResolvedOk returns a tuple with the NotResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResolved

`func (o *StudyAudienceStats) SetNotResolved(v MeasurementsContainer)`

SetNotResolved sets NotResolved field to given value.

### HasNotResolved

`func (o *StudyAudienceStats) HasNotResolved() bool`

HasNotResolved returns a boolean if a field has been set.

### GetBusiness

`func (o *StudyAudienceStats) GetBusiness() BusinessAudienceStats`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *StudyAudienceStats) GetBusinessOk() (*BusinessAudienceStats, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *StudyAudienceStats) SetBusiness(v BusinessAudienceStats)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *StudyAudienceStats) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetPrivate

`func (o *StudyAudienceStats) GetPrivate() PrivateAudienceStats`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *StudyAudienceStats) GetPrivateOk() (*PrivateAudienceStats, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *StudyAudienceStats) SetPrivate(v PrivateAudienceStats)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *StudyAudienceStats) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


