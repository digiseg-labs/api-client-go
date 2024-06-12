/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young couples and singles | |  |  | c2 | Early family life | |  |  | c3 | Middle-aged families | |  |  | c4 | Mature families | |  |  | c5 | Pensioners / Retirees | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Small | |  |  | l2 | Medium | |  |  | l3 | Large | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AccountBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountBase{}

// AccountBase struct for AccountBase
type AccountBase struct {
	// Human readable name of the account
	Name *string `json:"name,omitempty"`
	// The URL to the logo of the account
	LogoUrl *string `json:"logo_url,omitempty"`
	// URL of the account's primary website
	WebsiteUrl *string `json:"website_url,omitempty"`
	// Country code of the account. Requires `owner` role to change.
	BillingCountry *string `json:"billing_country,omitempty"`
	// The type of company that the account represents. Note that for forward-compatibility the data type here is simply a string. The values, if present, will however typically originate from the `CompanyType` enum. 
	CompanyType *string `json:"company_type,omitempty"`
	CompanySize *CompanySize `json:"company_size,omitempty"`
	// Determines whether the account has clients that they work for, or if their activities are for themselves.
	HasClients *bool `json:"has_clients,omitempty"`
	// A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters.
	// Deprecated
	Slug *string `json:"slug,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBase AccountBase

// NewAccountBase instantiates a new AccountBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBase() *AccountBase {
	this := AccountBase{}
	return &this
}

// NewAccountBaseWithDefaults instantiates a new AccountBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBaseWithDefaults() *AccountBase {
	this := AccountBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountBase) SetName(v string) {
	o.Name = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AccountBase) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AccountBase) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AccountBase) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *AccountBase) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUrl) {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *AccountBase) HasWebsiteUrl() bool {
	if o != nil && !IsNil(o.WebsiteUrl) {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *AccountBase) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetBillingCountry returns the BillingCountry field value if set, zero value otherwise.
func (o *AccountBase) GetBillingCountry() string {
	if o == nil || IsNil(o.BillingCountry) {
		var ret string
		return ret
	}
	return *o.BillingCountry
}

// GetBillingCountryOk returns a tuple with the BillingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBillingCountryOk() (*string, bool) {
	if o == nil || IsNil(o.BillingCountry) {
		return nil, false
	}
	return o.BillingCountry, true
}

// HasBillingCountry returns a boolean if a field has been set.
func (o *AccountBase) HasBillingCountry() bool {
	if o != nil && !IsNil(o.BillingCountry) {
		return true
	}

	return false
}

// SetBillingCountry gets a reference to the given string and assigns it to the BillingCountry field.
func (o *AccountBase) SetBillingCountry(v string) {
	o.BillingCountry = &v
}

// GetCompanyType returns the CompanyType field value if set, zero value otherwise.
func (o *AccountBase) GetCompanyType() string {
	if o == nil || IsNil(o.CompanyType) {
		var ret string
		return ret
	}
	return *o.CompanyType
}

// GetCompanyTypeOk returns a tuple with the CompanyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCompanyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyType) {
		return nil, false
	}
	return o.CompanyType, true
}

// HasCompanyType returns a boolean if a field has been set.
func (o *AccountBase) HasCompanyType() bool {
	if o != nil && !IsNil(o.CompanyType) {
		return true
	}

	return false
}

// SetCompanyType gets a reference to the given string and assigns it to the CompanyType field.
func (o *AccountBase) SetCompanyType(v string) {
	o.CompanyType = &v
}

// GetCompanySize returns the CompanySize field value if set, zero value otherwise.
func (o *AccountBase) GetCompanySize() CompanySize {
	if o == nil || IsNil(o.CompanySize) {
		var ret CompanySize
		return ret
	}
	return *o.CompanySize
}

// GetCompanySizeOk returns a tuple with the CompanySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCompanySizeOk() (*CompanySize, bool) {
	if o == nil || IsNil(o.CompanySize) {
		return nil, false
	}
	return o.CompanySize, true
}

// HasCompanySize returns a boolean if a field has been set.
func (o *AccountBase) HasCompanySize() bool {
	if o != nil && !IsNil(o.CompanySize) {
		return true
	}

	return false
}

// SetCompanySize gets a reference to the given CompanySize and assigns it to the CompanySize field.
func (o *AccountBase) SetCompanySize(v CompanySize) {
	o.CompanySize = &v
}

// GetHasClients returns the HasClients field value if set, zero value otherwise.
func (o *AccountBase) GetHasClients() bool {
	if o == nil || IsNil(o.HasClients) {
		var ret bool
		return ret
	}
	return *o.HasClients
}

// GetHasClientsOk returns a tuple with the HasClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetHasClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasClients) {
		return nil, false
	}
	return o.HasClients, true
}

// HasHasClients returns a boolean if a field has been set.
func (o *AccountBase) HasHasClients() bool {
	if o != nil && !IsNil(o.HasClients) {
		return true
	}

	return false
}

// SetHasClients gets a reference to the given bool and assigns it to the HasClients field.
func (o *AccountBase) SetHasClients(v bool) {
	o.HasClients = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
// Deprecated
func (o *AccountBase) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountBase) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *AccountBase) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
// Deprecated
func (o *AccountBase) SetSlug(v string) {
	o.Slug = &v
}

func (o AccountBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.WebsiteUrl) {
		toSerialize["website_url"] = o.WebsiteUrl
	}
	if !IsNil(o.BillingCountry) {
		toSerialize["billing_country"] = o.BillingCountry
	}
	if !IsNil(o.CompanyType) {
		toSerialize["company_type"] = o.CompanyType
	}
	if !IsNil(o.CompanySize) {
		toSerialize["company_size"] = o.CompanySize
	}
	if !IsNil(o.HasClients) {
		toSerialize["has_clients"] = o.HasClients
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountBase) UnmarshalJSON(data []byte) (err error) {
	varAccountBase := _AccountBase{}

	err = json.Unmarshal(data, &varAccountBase)

	if err != nil {
		return err
	}

	*o = AccountBase(varAccountBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "website_url")
		delete(additionalProperties, "billing_country")
		delete(additionalProperties, "company_type")
		delete(additionalProperties, "company_size")
		delete(additionalProperties, "has_clients")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBase struct {
	value *AccountBase
	isSet bool
}

func (v NullableAccountBase) Get() *AccountBase {
	return v.value
}

func (v *NullableAccountBase) Set(val *AccountBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBase(val *AccountBase) *NullableAccountBase {
	return &NullableAccountBase{value: val, isSet: true}
}

func (v NullableAccountBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


