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
	"time"
)

// checks if the UserFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFull{}

// UserFull struct for UserFull
type UserFull struct {
	// Unique ID for the object
	Id *string `json:"id,omitempty"`
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
	// The approximate last time that the user logged in
	LoggedInAt *time.Time `json:"logged_in_at,omitempty"`
	AccountMemberships []UserAccountMembership `json:"account_memberships,omitempty"`
	// Determines if the user is a super admin of Digiseg API services
	// Deprecated
	IsSuperAdmin *bool `json:"is_super_admin,omitempty"`
	PlatformRoles []UserPlatformRole `json:"platform_roles,omitempty"`
	// Date and time of the object creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ID of the user who created the object
	CreatedBy *string `json:"created_by,omitempty"`
	// Date and time of the latest update to the object
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the user who last updated the object
	UpdatedBy *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFull UserFull

// NewUserFull instantiates a new UserFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFull() *UserFull {
	this := UserFull{}
	return &this
}

// NewUserFullWithDefaults instantiates a new UserFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFullWithDefaults() *UserFull {
	this := UserFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserFull) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserFull) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserFull) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserFull) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserFull) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UserFull) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UserFull) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UserFull) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserFull) GetRoles() []UserAccountRole {
	if o == nil || IsNil(o.Roles) {
		var ret []UserAccountRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetRolesOk() ([]UserAccountRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserFull) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserAccountRole and assigns it to the Roles field.
func (o *UserFull) SetRoles(v []UserAccountRole) {
	o.Roles = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserFull) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserFull) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserFull) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetLoggedInAt returns the LoggedInAt field value if set, zero value otherwise.
func (o *UserFull) GetLoggedInAt() time.Time {
	if o == nil || IsNil(o.LoggedInAt) {
		var ret time.Time
		return ret
	}
	return *o.LoggedInAt
}

// GetLoggedInAtOk returns a tuple with the LoggedInAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetLoggedInAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LoggedInAt) {
		return nil, false
	}
	return o.LoggedInAt, true
}

// HasLoggedInAt returns a boolean if a field has been set.
func (o *UserFull) HasLoggedInAt() bool {
	if o != nil && !IsNil(o.LoggedInAt) {
		return true
	}

	return false
}

// SetLoggedInAt gets a reference to the given time.Time and assigns it to the LoggedInAt field.
func (o *UserFull) SetLoggedInAt(v time.Time) {
	o.LoggedInAt = &v
}

// GetAccountMemberships returns the AccountMemberships field value if set, zero value otherwise.
func (o *UserFull) GetAccountMemberships() []UserAccountMembership {
	if o == nil || IsNil(o.AccountMemberships) {
		var ret []UserAccountMembership
		return ret
	}
	return o.AccountMemberships
}

// GetAccountMembershipsOk returns a tuple with the AccountMemberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetAccountMembershipsOk() ([]UserAccountMembership, bool) {
	if o == nil || IsNil(o.AccountMemberships) {
		return nil, false
	}
	return o.AccountMemberships, true
}

// HasAccountMemberships returns a boolean if a field has been set.
func (o *UserFull) HasAccountMemberships() bool {
	if o != nil && !IsNil(o.AccountMemberships) {
		return true
	}

	return false
}

// SetAccountMemberships gets a reference to the given []UserAccountMembership and assigns it to the AccountMemberships field.
func (o *UserFull) SetAccountMemberships(v []UserAccountMembership) {
	o.AccountMemberships = v
}

// GetIsSuperAdmin returns the IsSuperAdmin field value if set, zero value otherwise.
// Deprecated
func (o *UserFull) GetIsSuperAdmin() bool {
	if o == nil || IsNil(o.IsSuperAdmin) {
		var ret bool
		return ret
	}
	return *o.IsSuperAdmin
}

// GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserFull) GetIsSuperAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperAdmin) {
		return nil, false
	}
	return o.IsSuperAdmin, true
}

// HasIsSuperAdmin returns a boolean if a field has been set.
func (o *UserFull) HasIsSuperAdmin() bool {
	if o != nil && !IsNil(o.IsSuperAdmin) {
		return true
	}

	return false
}

// SetIsSuperAdmin gets a reference to the given bool and assigns it to the IsSuperAdmin field.
// Deprecated
func (o *UserFull) SetIsSuperAdmin(v bool) {
	o.IsSuperAdmin = &v
}

// GetPlatformRoles returns the PlatformRoles field value if set, zero value otherwise.
func (o *UserFull) GetPlatformRoles() []UserPlatformRole {
	if o == nil || IsNil(o.PlatformRoles) {
		var ret []UserPlatformRole
		return ret
	}
	return o.PlatformRoles
}

// GetPlatformRolesOk returns a tuple with the PlatformRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetPlatformRolesOk() ([]UserPlatformRole, bool) {
	if o == nil || IsNil(o.PlatformRoles) {
		return nil, false
	}
	return o.PlatformRoles, true
}

// HasPlatformRoles returns a boolean if a field has been set.
func (o *UserFull) HasPlatformRoles() bool {
	if o != nil && !IsNil(o.PlatformRoles) {
		return true
	}

	return false
}

// SetPlatformRoles gets a reference to the given []UserPlatformRole and assigns it to the PlatformRoles field.
func (o *UserFull) SetPlatformRoles(v []UserPlatformRole) {
	o.PlatformRoles = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserFull) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserFull) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *UserFull) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UserFull) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *UserFull) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserFull) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserFull) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *UserFull) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFull) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *UserFull) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *UserFull) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o UserFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
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
	if !IsNil(o.LoggedInAt) {
		toSerialize["logged_in_at"] = o.LoggedInAt
	}
	if !IsNil(o.AccountMemberships) {
		toSerialize["account_memberships"] = o.AccountMemberships
	}
	if !IsNil(o.IsSuperAdmin) {
		toSerialize["is_super_admin"] = o.IsSuperAdmin
	}
	if !IsNil(o.PlatformRoles) {
		toSerialize["platform_roles"] = o.PlatformRoles
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFull) UnmarshalJSON(data []byte) (err error) {
	varUserFull := _UserFull{}

	err = json.Unmarshal(data, &varUserFull)

	if err != nil {
		return err
	}

	*o = UserFull(varUserFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "logged_in_at")
		delete(additionalProperties, "account_memberships")
		delete(additionalProperties, "is_super_admin")
		delete(additionalProperties, "platform_roles")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFull struct {
	value *UserFull
	isSet bool
}

func (v NullableUserFull) Get() *UserFull {
	return v.value
}

func (v *NullableUserFull) Set(val *UserFull) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFull) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFull(val *UserFull) *NullableUserFull {
	return &NullableUserFull{value: val, isSet: true}
}

func (v NullableUserFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


