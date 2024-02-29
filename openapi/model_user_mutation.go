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

// checks if the UserMutation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMutation{}

// UserMutation struct for UserMutation
type UserMutation struct {
	// The email of the user (used as username when authenticating with password)
	Email *string `json:"email,omitempty"`
	// Human readable name of the user
	Name *string `json:"name,omitempty"`
	// ID of the account that this user pertains to. If the user has multiple account memberships, this account ID will represent the primary account of the user. 
	AccountId *string `json:"account_id,omitempty"`
	// The roles that the user has within the account
	Roles []UserAccountRole `json:"roles,omitempty"`
	// The URL to an avatar of the user
	AvatarUrl *string `json:"avatar_url,omitempty"`
	AccountMemberships []UserAccountMembership `json:"account_memberships,omitempty"`
	// Determines if the user is a super admin of Digiseg API services
	IsSuperAdmin *bool `json:"is_super_admin,omitempty"`
	// Password of the user
	Password *string `json:"password,omitempty"`
}

// NewUserMutation instantiates a new UserMutation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMutation() *UserMutation {
	this := UserMutation{}
	return &this
}

// NewUserMutationWithDefaults instantiates a new UserMutation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMutationWithDefaults() *UserMutation {
	this := UserMutation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserMutation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserMutation) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserMutation) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserMutation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserMutation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserMutation) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UserMutation) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UserMutation) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UserMutation) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserMutation) GetRoles() []UserAccountRole {
	if o == nil || IsNil(o.Roles) {
		var ret []UserAccountRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetRolesOk() ([]UserAccountRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserMutation) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserAccountRole and assigns it to the Roles field.
func (o *UserMutation) SetRoles(v []UserAccountRole) {
	o.Roles = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserMutation) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserMutation) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserMutation) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetAccountMemberships returns the AccountMemberships field value if set, zero value otherwise.
func (o *UserMutation) GetAccountMemberships() []UserAccountMembership {
	if o == nil || IsNil(o.AccountMemberships) {
		var ret []UserAccountMembership
		return ret
	}
	return o.AccountMemberships
}

// GetAccountMembershipsOk returns a tuple with the AccountMemberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetAccountMembershipsOk() ([]UserAccountMembership, bool) {
	if o == nil || IsNil(o.AccountMemberships) {
		return nil, false
	}
	return o.AccountMemberships, true
}

// HasAccountMemberships returns a boolean if a field has been set.
func (o *UserMutation) HasAccountMemberships() bool {
	if o != nil && !IsNil(o.AccountMemberships) {
		return true
	}

	return false
}

// SetAccountMemberships gets a reference to the given []UserAccountMembership and assigns it to the AccountMemberships field.
func (o *UserMutation) SetAccountMemberships(v []UserAccountMembership) {
	o.AccountMemberships = v
}

// GetIsSuperAdmin returns the IsSuperAdmin field value if set, zero value otherwise.
func (o *UserMutation) GetIsSuperAdmin() bool {
	if o == nil || IsNil(o.IsSuperAdmin) {
		var ret bool
		return ret
	}
	return *o.IsSuperAdmin
}

// GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetIsSuperAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperAdmin) {
		return nil, false
	}
	return o.IsSuperAdmin, true
}

// HasIsSuperAdmin returns a boolean if a field has been set.
func (o *UserMutation) HasIsSuperAdmin() bool {
	if o != nil && !IsNil(o.IsSuperAdmin) {
		return true
	}

	return false
}

// SetIsSuperAdmin gets a reference to the given bool and assigns it to the IsSuperAdmin field.
func (o *UserMutation) SetIsSuperAdmin(v bool) {
	o.IsSuperAdmin = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserMutation) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMutation) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserMutation) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserMutation) SetPassword(v string) {
	o.Password = &v
}

func (o UserMutation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMutation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.AccountMemberships) {
		toSerialize["account_memberships"] = o.AccountMemberships
	}
	if !IsNil(o.IsSuperAdmin) {
		toSerialize["is_super_admin"] = o.IsSuperAdmin
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUserMutation struct {
	value *UserMutation
	isSet bool
}

func (v NullableUserMutation) Get() *UserMutation {
	return v.value
}

func (v *NullableUserMutation) Set(val *UserMutation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMutation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMutation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMutation(val *UserMutation) *NullableUserMutation {
	return &NullableUserMutation{value: val, isSet: true}
}

func (v NullableUserMutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMutation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


