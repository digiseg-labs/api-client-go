# AccountBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the logo of the account | [optional] 
**WebsiteUrl** | Pointer to **string** | URL of the account&#39;s primary website | [optional] 
**BillingCountry** | Pointer to **string** | Country code of the account. Requires &#x60;owner&#x60; role to change. | [optional] 
**CompanyType** | Pointer to **string** | The type of company that the account represents. Note that for forward-compatibility the data type here is simply a string. The values, if present, will however typically originate from the &#x60;CompanyType&#x60; enum.  | [optional] 
**CompanySize** | Pointer to [**CompanySize**](CompanySize.md) |  | [optional] 
**HasClients** | Pointer to **bool** | Determines whether the account has clients that they work for, or if their activities are for themselves. | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 

## Methods

### NewAccountBase

`func NewAccountBase() *AccountBase`

NewAccountBase instantiates a new AccountBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBaseWithDefaults

`func NewAccountBaseWithDefaults() *AccountBase`

NewAccountBaseWithDefaults instantiates a new AccountBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountBase) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountBase) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountBase) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountBase) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *AccountBase) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *AccountBase) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *AccountBase) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *AccountBase) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetBillingCountry

`func (o *AccountBase) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *AccountBase) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *AccountBase) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *AccountBase) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetCompanyType

`func (o *AccountBase) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *AccountBase) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *AccountBase) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *AccountBase) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountBase) GetCompanySize() CompanySize`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountBase) GetCompanySizeOk() (*CompanySize, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountBase) SetCompanySize(v CompanySize)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountBase) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetHasClients

`func (o *AccountBase) GetHasClients() bool`

GetHasClients returns the HasClients field if non-nil, zero value otherwise.

### GetHasClientsOk

`func (o *AccountBase) GetHasClientsOk() (*bool, bool)`

GetHasClientsOk returns a tuple with the HasClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClients

`func (o *AccountBase) SetHasClients(v bool)`

SetHasClients sets HasClients field to given value.

### HasHasClients

`func (o *AccountBase) HasHasClients() bool`

HasHasClients returns a boolean if a field has been set.

### GetSlug

`func (o *AccountBase) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountBase) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountBase) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountBase) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


