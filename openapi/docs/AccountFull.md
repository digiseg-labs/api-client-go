# AccountFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] [readonly] 
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the logo of the account | [optional] 
**WebsiteUrl** | Pointer to **string** | URL of the account&#39;s primary website | [optional] 
**BillingCountry** | Pointer to **string** | Country code of the account. Requires &#x60;owner&#x60; role to change. | [optional] 
**CompanyType** | Pointer to **string** | The type of company that the account represents. Note that for forward-compatibility the data type here is simply a string. The values, if present, will however typically originate from the &#x60;CompanyType&#x60; enum.  | [optional] 
**CompanySize** | Pointer to [**CompanySize**](CompanySize.md) |  | [optional] 
**HasClients** | Pointer to **bool** | Determines whether the account has clients that they work for, or if their activities are for themselves. | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account. Deprecated in favor of the &#x60;owner&#x60; role of the user&#39;s account membership. | [optional] 
**BillingEmail** | Pointer to **string** | The email address to send billing information to. Requires &#x60;owner&#x60; role to change. | [optional] 
**BillingAddress** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 
**BillingTaxIds** | Pointer to [**[]TaxId**](TaxId.md) | A list of Tax IDs used by the account, for billing purposes. | [optional] [readonly] 
**BillingCurrency** | Pointer to [**SubscriptionPriceCurrency**](SubscriptionPriceCurrency.md) |  | [optional] [default to SUBSCRIPTIONPRICECURRENCY_EUR]
**BillingName** | Pointer to **string** | An optional official name to use for billing purposes. Requires &#x60;owner&#x60; role to change. | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] [readonly] 
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time of the object creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | ID of the user who created the object | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Date and time of the latest update to the object | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | ID of the user who last updated the object | [optional] [readonly] 

## Methods

### NewAccountFull

`func NewAccountFull() *AccountFull`

NewAccountFull instantiates a new AccountFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFullWithDefaults

`func NewAccountFullWithDefaults() *AccountFull`

NewAccountFullWithDefaults instantiates a new AccountFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountFull) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountFull) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountFull) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountFull) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *AccountFull) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *AccountFull) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *AccountFull) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *AccountFull) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetBillingCountry

`func (o *AccountFull) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *AccountFull) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *AccountFull) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *AccountFull) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetCompanyType

`func (o *AccountFull) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *AccountFull) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *AccountFull) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *AccountFull) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountFull) GetCompanySize() CompanySize`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountFull) GetCompanySizeOk() (*CompanySize, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountFull) SetCompanySize(v CompanySize)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountFull) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetHasClients

`func (o *AccountFull) GetHasClients() bool`

GetHasClients returns the HasClients field if non-nil, zero value otherwise.

### GetHasClientsOk

`func (o *AccountFull) GetHasClientsOk() (*bool, bool)`

GetHasClientsOk returns a tuple with the HasClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClients

`func (o *AccountFull) SetHasClients(v bool)`

SetHasClients sets HasClients field to given value.

### HasHasClients

`func (o *AccountFull) HasHasClients() bool`

HasHasClients returns a boolean if a field has been set.

### GetSlug

`func (o *AccountFull) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountFull) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountFull) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountFull) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountFull) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetBillingEmail

`func (o *AccountFull) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountFull) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *AccountFull) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *AccountFull) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *AccountFull) GetBillingAddress() PostalAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *AccountFull) GetBillingAddressOk() (*PostalAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *AccountFull) SetBillingAddress(v PostalAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *AccountFull) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBillingTaxIds

`func (o *AccountFull) GetBillingTaxIds() []TaxId`

GetBillingTaxIds returns the BillingTaxIds field if non-nil, zero value otherwise.

### GetBillingTaxIdsOk

`func (o *AccountFull) GetBillingTaxIdsOk() (*[]TaxId, bool)`

GetBillingTaxIdsOk returns a tuple with the BillingTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTaxIds

`func (o *AccountFull) SetBillingTaxIds(v []TaxId)`

SetBillingTaxIds sets BillingTaxIds field to given value.

### HasBillingTaxIds

`func (o *AccountFull) HasBillingTaxIds() bool`

HasBillingTaxIds returns a boolean if a field has been set.

### GetBillingCurrency

`func (o *AccountFull) GetBillingCurrency() SubscriptionPriceCurrency`

GetBillingCurrency returns the BillingCurrency field if non-nil, zero value otherwise.

### GetBillingCurrencyOk

`func (o *AccountFull) GetBillingCurrencyOk() (*SubscriptionPriceCurrency, bool)`

GetBillingCurrencyOk returns a tuple with the BillingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCurrency

`func (o *AccountFull) SetBillingCurrency(v SubscriptionPriceCurrency)`

SetBillingCurrency sets BillingCurrency field to given value.

### HasBillingCurrency

`func (o *AccountFull) HasBillingCurrency() bool`

HasBillingCurrency returns a boolean if a field has been set.

### GetBillingName

`func (o *AccountFull) GetBillingName() string`

GetBillingName returns the BillingName field if non-nil, zero value otherwise.

### GetBillingNameOk

`func (o *AccountFull) GetBillingNameOk() (*string, bool)`

GetBillingNameOk returns a tuple with the BillingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingName

`func (o *AccountFull) SetBillingName(v string)`

SetBillingName sets BillingName field to given value.

### HasBillingName

`func (o *AccountFull) HasBillingName() bool`

HasBillingName returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *AccountFull) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *AccountFull) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *AccountFull) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *AccountFull) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetFeatureSet

`func (o *AccountFull) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *AccountFull) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *AccountFull) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *AccountFull) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetSubscriptions

`func (o *AccountFull) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AccountFull) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AccountFull) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AccountFull) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountFull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AccountFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccountFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountFull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AccountFull) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AccountFull) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AccountFull) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AccountFull) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


