# AudienceDataDailyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date in ISO 8601 format | 
**Usage** | [**AudienceDataUsage**](AudienceDataUsage.md) |  | 

## Methods

### NewAudienceDataDailyUsage

`func NewAudienceDataDailyUsage(date string, usage AudienceDataUsage, ) *AudienceDataDailyUsage`

NewAudienceDataDailyUsage instantiates a new AudienceDataDailyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceDataDailyUsageWithDefaults

`func NewAudienceDataDailyUsageWithDefaults() *AudienceDataDailyUsage`

NewAudienceDataDailyUsageWithDefaults instantiates a new AudienceDataDailyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AudienceDataDailyUsage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AudienceDataDailyUsage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AudienceDataDailyUsage) SetDate(v string)`

SetDate sets Date field to given value.


### GetUsage

`func (o *AudienceDataDailyUsage) GetUsage() AudienceDataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AudienceDataDailyUsage) GetUsageOk() (*AudienceDataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AudienceDataDailyUsage) SetUsage(v AudienceDataUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


