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
	"bytes"
	"fmt"
)

// checks if the AuthTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenRequest{}

// AuthTokenRequest struct for AuthTokenRequest
type AuthTokenRequest struct {
	// The username (typically an email address) of the user to authenticate
	Username string `json:"username"`
	// A one-time password provided to perform passwordless auth
	Otp *string `json:"otp,omitempty"`
	// The password for the given username
	Password *string `json:"password,omitempty"`
	// A previously issued refresh token for the given username
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scopes *PermissionScopes `json:"scopes,omitempty"`
}

type _AuthTokenRequest AuthTokenRequest

// NewAuthTokenRequest instantiates a new AuthTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenRequest(username string) *AuthTokenRequest {
	this := AuthTokenRequest{}
	this.Username = username
	return &this
}

// NewAuthTokenRequestWithDefaults instantiates a new AuthTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenRequestWithDefaults() *AuthTokenRequest {
	this := AuthTokenRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthTokenRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthTokenRequest) SetUsername(v string) {
	o.Username = v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *AuthTokenRequest) GetOtp() string {
	if o == nil || IsNil(o.Otp) {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetOtpOk() (*string, bool) {
	if o == nil || IsNil(o.Otp) {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *AuthTokenRequest) HasOtp() bool {
	if o != nil && !IsNil(o.Otp) {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *AuthTokenRequest) SetOtp(v string) {
	o.Otp = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthTokenRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthTokenRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthTokenRequest) SetPassword(v string) {
	o.Password = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AuthTokenRequest) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AuthTokenRequest) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AuthTokenRequest) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthTokenRequest) GetScopes() PermissionScopes {
	if o == nil || IsNil(o.Scopes) {
		var ret PermissionScopes
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetScopesOk() (*PermissionScopes, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthTokenRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given PermissionScopes and assigns it to the Scopes field.
func (o *AuthTokenRequest) SetScopes(v PermissionScopes) {
	o.Scopes = &v
}

func (o AuthTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.Otp) {
		toSerialize["otp"] = o.Otp
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

func (o *AuthTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthTokenRequest := _AuthTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthTokenRequest)

	if err != nil {
		return err
	}

	*o = AuthTokenRequest(varAuthTokenRequest)

	return err
}

type NullableAuthTokenRequest struct {
	value *AuthTokenRequest
	isSet bool
}

func (v NullableAuthTokenRequest) Get() *AuthTokenRequest {
	return v.value
}

func (v *NullableAuthTokenRequest) Set(val *AuthTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenRequest(val *AuthTokenRequest) *NullableAuthTokenRequest {
	return &NullableAuthTokenRequest{value: val, isSet: true}
}

func (v NullableAuthTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


