# AccountCreation

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
**StripeCustomerId** | Pointer to **string** |  | [optional] [readonly] 
**Owner** | Pointer to [**AccountOwnerCreation**](AccountOwnerCreation.md) |  | [optional] 
**NotifyUser** | Pointer to **bool** | Whether or not to notify the user that they have been registered | [optional] [default to true]

## Methods

### NewAccountCreation

`func NewAccountCreation() *AccountCreation`

NewAccountCreation instantiates a new AccountCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationWithDefaults

`func NewAccountCreationWithDefaults() *AccountCreation`

NewAccountCreationWithDefaults instantiates a new AccountCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCreation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountCreation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountCreation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountCreation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountCreation) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *AccountCreation) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *AccountCreation) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *AccountCreation) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *AccountCreation) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetBillingCountry

`func (o *AccountCreation) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *AccountCreation) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *AccountCreation) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *AccountCreation) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetCompanyType

`func (o *AccountCreation) GetCompanyType() CompanyType`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *AccountCreation) GetCompanyTypeOk() (*CompanyType, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *AccountCreation) SetCompanyType(v CompanyType)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *AccountCreation) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountCreation) GetCompanySize() CompanySize`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountCreation) GetCompanySizeOk() (*CompanySize, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountCreation) SetCompanySize(v CompanySize)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountCreation) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetHasClients

`func (o *AccountCreation) GetHasClients() bool`

GetHasClients returns the HasClients field if non-nil, zero value otherwise.

### GetHasClientsOk

`func (o *AccountCreation) GetHasClientsOk() (*bool, bool)`

GetHasClientsOk returns a tuple with the HasClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClients

`func (o *AccountCreation) SetHasClients(v bool)`

SetHasClients sets HasClients field to given value.

### HasHasClients

`func (o *AccountCreation) HasHasClients() bool`

HasHasClients returns a boolean if a field has been set.

### GetSlug

`func (o *AccountCreation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountCreation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountCreation) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountCreation) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountCreation) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountCreation) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountCreation) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountCreation) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetBillingEmail

`func (o *AccountCreation) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountCreation) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *AccountCreation) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *AccountCreation) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *AccountCreation) GetBillingAddress() PostalAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountCreation) GetBillingAddressOk() (*PostalAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountCreation) SetBillingAddress(v PostalAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountCreation) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *AccountCreation) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *AccountCreation) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *AccountCreation) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *AccountCreation) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetOwner

`func (o *AccountCreation) GetOwner() AccountOwnerCreation`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountCreation) GetOwnerOk() (*AccountOwnerCreation, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountCreation) SetOwner(v AccountOwnerCreation)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountCreation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotifyUser

`func (o *AccountCreation) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *AccountCreation) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *AccountCreation) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *AccountCreation) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


