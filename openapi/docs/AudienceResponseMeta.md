# AudienceResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to [**CountryItem**](CountryItem.md) |  | [optional] 
**Platform** | Pointer to [**AudiencePlatformItem**](AudiencePlatformItem.md) |  | [optional] 

## Methods

### NewAudienceResponseMeta

`func NewAudienceResponseMeta() *AudienceResponseMeta`

NewAudienceResponseMeta instantiates a new AudienceResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceResponseMetaWithDefaults

`func NewAudienceResponseMetaWithDefaults() *AudienceResponseMeta`

NewAudienceResponseMetaWithDefaults instantiates a new AudienceResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *AudienceResponseMeta) GetCountry() CountryItem`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AudienceResponseMeta) GetCountryOk() (*CountryItem, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AudienceResponseMeta) SetCountry(v CountryItem)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AudienceResponseMeta) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPlatform

`func (o *AudienceResponseMeta) GetPlatform() AudiencePlatformItem`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AudienceResponseMeta) GetPlatformOk() (*AudiencePlatformItem, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AudienceResponseMeta) SetPlatform(v AudiencePlatformItem)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AudienceResponseMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


