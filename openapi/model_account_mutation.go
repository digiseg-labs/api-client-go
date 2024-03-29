/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young singles and couples | |  |  | c2 | Young couples with children | |  |  | c3 | Families with school children | |  |  | c4 | Older families | |  |  | c5 | Pensioners | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Up to 80 m² | |  |  | l2 | 80-119 m² | |  |  | l3 | Above 120 m² | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AccountMutation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountMutation{}

// AccountMutation struct for AccountMutation
type AccountMutation struct {
	// Human readable name of the account
	Name *string `json:"name,omitempty"`
	// The URL to the logo of the account
	LogoUrl *string `json:"logo_url,omitempty"`
	// A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters.
	// Deprecated
	Slug *string `json:"slug,omitempty"`
	// ID of the user who is the ultimate owner of the account
	OwnerId *string `json:"owner_id,omitempty"`
}

// NewAccountMutation instantiates a new AccountMutation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMutation() *AccountMutation {
	this := AccountMutation{}
	return &this
}

// NewAccountMutationWithDefaults instantiates a new AccountMutation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMutationWithDefaults() *AccountMutation {
	this := AccountMutation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountMutation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMutation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountMutation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountMutation) SetName(v string) {
	o.Name = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AccountMutation) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMutation) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AccountMutation) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AccountMutation) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
// Deprecated
func (o *AccountMutation) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountMutation) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *AccountMutation) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
// Deprecated
func (o *AccountMutation) SetSlug(v string) {
	o.Slug = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AccountMutation) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMutation) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AccountMutation) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AccountMutation) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o AccountMutation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountMutation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	return toSerialize, nil
}

type NullableAccountMutation struct {
	value *AccountMutation
	isSet bool
}

func (v NullableAccountMutation) Get() *AccountMutation {
	return v.value
}

func (v *NullableAccountMutation) Set(val *AccountMutation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMutation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMutation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMutation(val *AccountMutation) *NullableAccountMutation {
	return &NullableAccountMutation{value: val, isSet: true}
}

func (v NullableAccountMutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMutation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


