# PermissionScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginDomains** | Pointer to **[]string** | A list of origin domains that will be allowed access through CORS and Referrer checks with the token or API key. A missing value, empty array or the value &#x60;*&#x60; grants access from any domain. | [optional] 
**ServiceGrants** | Pointer to **[]string** | A list of service grants to allow access to use. A missing value, empty array or the value &#x60;*&#x60; grants access to all of the services that are available to non-admin and non-owner users of the account. | [optional] 
**CountryData** | Pointer to **[]string** | A list of country-specific data sources to allow access to. A missing value, empty array or the value &#x60;*&#x60; grants access to any data source. | [optional] 

## Methods

### NewPermissionScopes

`func NewPermissionScopes() *PermissionScopes`

NewPermissionScopes instantiates a new PermissionScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionScopesWithDefaults

`func NewPermissionScopesWithDefaults() *PermissionScopes`

NewPermissionScopesWithDefaults instantiates a new PermissionScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginDomains

`func (o *PermissionScopes) GetOriginDomains() []string`

GetOriginDomains returns the OriginDomains field if non-nil, zero value otherwise.

### GetOriginDomainsOk

`func (o *PermissionScopes) GetOriginDomainsOk() (*[]string, bool)`

GetOriginDomainsOk returns a tuple with the OriginDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDomains

`func (o *PermissionScopes) SetOriginDomains(v []string)`

SetOriginDomains sets OriginDomains field to given value.

### HasOriginDomains

`func (o *PermissionScopes) HasOriginDomains() bool`

HasOriginDomains returns a boolean if a field has been set.

### GetServiceGrants

`func (o *PermissionScopes) GetServiceGrants() []string`

GetServiceGrants returns the ServiceGrants field if non-nil, zero value otherwise.

### GetServiceGrantsOk

`func (o *PermissionScopes) GetServiceGrantsOk() (*[]string, bool)`

GetServiceGrantsOk returns a tuple with the ServiceGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGrants

`func (o *PermissionScopes) SetServiceGrants(v []string)`

SetServiceGrants sets ServiceGrants field to given value.

### HasServiceGrants

`func (o *PermissionScopes) HasServiceGrants() bool`

HasServiceGrants returns a boolean if a field has been set.

### GetCountryData

`func (o *PermissionScopes) GetCountryData() []string`

GetCountryData returns the CountryData field if non-nil, zero value otherwise.

### GetCountryDataOk

`func (o *PermissionScopes) GetCountryDataOk() (*[]string, bool)`

GetCountryDataOk returns a tuple with the CountryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryData

`func (o *PermissionScopes) SetCountryData(v []string)`

SetCountryData sets CountryData field to given value.

### HasCountryData

`func (o *PermissionScopes) HasCountryData() bool`

HasCountryData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


