# AudiencePlatformItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to represent the platform when querying platform-specific information | 
**DisplayName** | **string** | Name of the audience platform | 
**HasGlobalTaxonomy** | **bool** | Whether the platform-specific codes for the particular platform are global (or country-specific when NOT global) | 
**SupportedCountries** | **[]string** | A list of supported countries, each represented by their country code | 

## Methods

### NewAudiencePlatformItem

`func NewAudiencePlatformItem(code string, displayName string, hasGlobalTaxonomy bool, supportedCountries []string, ) *AudiencePlatformItem`

NewAudiencePlatformItem instantiates a new AudiencePlatformItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencePlatformItemWithDefaults

`func NewAudiencePlatformItemWithDefaults() *AudiencePlatformItem`

NewAudiencePlatformItemWithDefaults instantiates a new AudiencePlatformItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AudiencePlatformItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudiencePlatformItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudiencePlatformItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *AudiencePlatformItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AudiencePlatformItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AudiencePlatformItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHasGlobalTaxonomy

`func (o *AudiencePlatformItem) GetHasGlobalTaxonomy() bool`

GetHasGlobalTaxonomy returns the HasGlobalTaxonomy field if non-nil, zero value otherwise.

### GetHasGlobalTaxonomyOk

`func (o *AudiencePlatformItem) GetHasGlobalTaxonomyOk() (*bool, bool)`

GetHasGlobalTaxonomyOk returns a tuple with the HasGlobalTaxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGlobalTaxonomy

`func (o *AudiencePlatformItem) SetHasGlobalTaxonomy(v bool)`

SetHasGlobalTaxonomy sets HasGlobalTaxonomy field to given value.


### GetSupportedCountries

`func (o *AudiencePlatformItem) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *AudiencePlatformItem) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *AudiencePlatformItem) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


