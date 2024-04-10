# AccountMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the logo of the account | [optional] 
**WebsiteUrl** | Pointer to **string** | URL of the account&#39;s primary website | [optional] 
**BillingCountry** | Pointer to **string** | Country code of the account. Requires &#x60;owner&#x60; role to change. | [optional] 
**CompanyType** | Pointer to [**CompanyType**](CompanyType.md) |  | [optional] 
**CompanySize** | Pointer to [**CompanySize**](CompanySize.md) |  | [optional] 
**HasClients** | Pointer to **bool** | Determines whether the account has clients that they work for, or if their activities are for themselves. | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account. Deprecated in favor of the &#x60;owner&#x60; role of the user&#39;s account membership. | [optional] 
**BillingEmail** | Pointer to **string** | The email address to send billing information to. Requires &#x60;owner&#x60; role to change. | [optional] 
**BillingAddress** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 

## Methods

### NewAccountMutation

`func NewAccountMutation() *AccountMutation`

NewAccountMutation instantiates a new AccountMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMutationWithDefaults

`func NewAccountMutationWithDefaults() *AccountMutation`

NewAccountMutationWithDefaults instantiates a new AccountMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountMutation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountMutation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountMutation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountMutation) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *AccountMutation) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *AccountMutation) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *AccountMutation) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *AccountMutation) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetBillingCountry

`func (o *AccountMutation) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *AccountMutation) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *AccountMutation) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *AccountMutation) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetCompanyType

`func (o *AccountMutation) GetCompanyType() CompanyType`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *AccountMutation) GetCompanyTypeOk() (*CompanyType, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *AccountMutation) SetCompanyType(v CompanyType)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *AccountMutation) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountMutation) GetCompanySize() CompanySize`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountMutation) GetCompanySizeOk() (*CompanySize, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountMutation) SetCompanySize(v CompanySize)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountMutation) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetHasClients

`func (o *AccountMutation) GetHasClients() bool`

GetHasClients returns the HasClients field if non-nil, zero value otherwise.

### GetHasClientsOk

`func (o *AccountMutation) GetHasClientsOk() (*bool, bool)`

GetHasClientsOk returns a tuple with the HasClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClients

`func (o *AccountMutation) SetHasClients(v bool)`

SetHasClients sets HasClients field to given value.

### HasHasClients

`func (o *AccountMutation) HasHasClients() bool`

HasHasClients returns a boolean if a field has been set.

### GetSlug

`func (o *AccountMutation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountMutation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountMutation) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountMutation) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountMutation) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountMutation) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountMutation) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountMutation) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetBillingEmail

`func (o *AccountMutation) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountMutation) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *AccountMutation) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *AccountMutation) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *AccountMutation) GetBillingAddress() PostalAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountMutation) GetBillingAddressOk() (*PostalAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountMutation) SetBillingAddress(v PostalAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountMutation) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


