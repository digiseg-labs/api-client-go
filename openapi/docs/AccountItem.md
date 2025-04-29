# AccountItem

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
**FeatureSet** | Pointer to [**PlanFeatureSet**](PlanFeatureSet.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]AccountSubscriptionItem**](AccountSubscriptionItem.md) |  | [optional] [readonly] 

## Methods

### NewAccountItem

`func NewAccountItem() *AccountItem`

NewAccountItem instantiates a new AccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemWithDefaults

`func NewAccountItemWithDefaults() *AccountItem`

NewAccountItemWithDefaults instantiates a new AccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountItem) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *AccountItem) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *AccountItem) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *AccountItem) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *AccountItem) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetBillingCountry

`func (o *AccountItem) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *AccountItem) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *AccountItem) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *AccountItem) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetCompanyType

`func (o *AccountItem) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *AccountItem) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *AccountItem) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *AccountItem) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySize

`func (o *AccountItem) GetCompanySize() CompanySize`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *AccountItem) GetCompanySizeOk() (*CompanySize, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *AccountItem) SetCompanySize(v CompanySize)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *AccountItem) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetHasClients

`func (o *AccountItem) GetHasClients() bool`

GetHasClients returns the HasClients field if non-nil, zero value otherwise.

### GetHasClientsOk

`func (o *AccountItem) GetHasClientsOk() (*bool, bool)`

GetHasClientsOk returns a tuple with the HasClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClients

`func (o *AccountItem) SetHasClients(v bool)`

SetHasClients sets HasClients field to given value.

### HasHasClients

`func (o *AccountItem) HasHasClients() bool`

HasHasClients returns a boolean if a field has been set.

### GetSlug

`func (o *AccountItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountItem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountItem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetFeatureSet

`func (o *AccountItem) GetFeatureSet() PlanFeatureSet`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *AccountItem) GetFeatureSetOk() (*PlanFeatureSet, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *AccountItem) SetFeatureSet(v PlanFeatureSet)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *AccountItem) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetSubscriptions

`func (o *AccountItem) GetSubscriptions() []AccountSubscriptionItem`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AccountItem) GetSubscriptionsOk() (*[]AccountSubscriptionItem, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AccountItem) SetSubscriptions(v []AccountSubscriptionItem)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AccountItem) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


