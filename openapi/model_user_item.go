/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserItem{}

// UserItem struct for UserItem
type UserItem struct {
	// Unique ID for the object
	Id *string `json:"id,omitempty"`
	// The email of the user (used as username when authenticating with password)
	Email *string `json:"email,omitempty"`
	// Human readable name of the user
	Name *string `json:"name,omitempty"`
	// The roles that the user has within the account
	Roles []UserAccountRole `json:"roles,omitempty"`
	// The URL to an avatar of the user
	AvatarUrl *string `json:"avatar_url,omitempty"`
}

// NewUserItem instantiates a new UserItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserItem() *UserItem {
	this := UserItem{}
	return &this
}

// NewUserItemWithDefaults instantiates a new UserItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserItemWithDefaults() *UserItem {
	this := UserItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserItem) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserItem) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserItem) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserItem) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserItem) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserItem) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserItem) GetRoles() []UserAccountRole {
	if o == nil || IsNil(o.Roles) {
		var ret []UserAccountRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserItem) GetRolesOk() ([]UserAccountRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserItem) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserAccountRole and assigns it to the Roles field.
func (o *UserItem) SetRoles(v []UserAccountRole) {
	o.Roles = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserItem) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserItem) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserItem) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserItem) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

func (o UserItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return toSerialize, nil
}

type NullableUserItem struct {
	value *UserItem
	isSet bool
}

func (v NullableUserItem) Get() *UserItem {
	return v.value
}

func (v *NullableUserItem) Set(val *UserItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserItem(val *UserItem) *NullableUserItem {
	return &NullableUserItem{value: val, isSet: true}
}

func (v NullableUserItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


