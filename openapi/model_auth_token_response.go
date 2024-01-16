/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  For a catalog of Digisegs audiences, refer to the [Audience list](https://digiseg.io/audiences-list).  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the AuthTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenResponse{}

// AuthTokenResponse struct for AuthTokenResponse
type AuthTokenResponse struct {
	// A short-lived token (usable for 1 hour) to be used in subsequent requests
	AccessToken string `json:"access_token"`
	// The type of access token returned
	TokenType string `json:"token_type"`
	// The duration of time (in seconds) the access token is granted for
	ExpiresIn int32 `json:"expires_in"`
	// A long-lived token that can be used to generate new access tokens even after the returned access token expires.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type _AuthTokenResponse AuthTokenResponse

// NewAuthTokenResponse instantiates a new AuthTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenResponse(accessToken string, tokenType string, expiresIn int32) *AuthTokenResponse {
	this := AuthTokenResponse{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	this.ExpiresIn = expiresIn
	return &this
}

// NewAuthTokenResponseWithDefaults instantiates a new AuthTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenResponseWithDefaults() *AuthTokenResponse {
	this := AuthTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthTokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthTokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *AuthTokenResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AuthTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AuthTokenResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *AuthTokenResponse) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *AuthTokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *AuthTokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AuthTokenResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AuthTokenResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AuthTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o AuthTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["token_type"] = o.TokenType
	toSerialize["expires_in"] = o.ExpiresIn
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

func (o *AuthTokenResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"token_type",
		"expires_in",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthTokenResponse := _AuthTokenResponse{}

	err = json.Unmarshal(bytes, &varAuthTokenResponse)

	if err != nil {
		return err
	}

	*o = AuthTokenResponse(varAuthTokenResponse)

	return err
}

type NullableAuthTokenResponse struct {
	value *AuthTokenResponse
	isSet bool
}

func (v NullableAuthTokenResponse) Get() *AuthTokenResponse {
	return v.value
}

func (v *NullableAuthTokenResponse) Set(val *AuthTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenResponse(val *AuthTokenResponse) *NullableAuthTokenResponse {
	return &NullableAuthTokenResponse{value: val, isSet: true}
}

func (v NullableAuthTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

