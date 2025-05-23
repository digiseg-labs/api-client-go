/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" /> 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SharedReportFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedReportFull{}

// SharedReportFull struct for SharedReportFull
type SharedReportFull struct {
	// Unique ID for the object
	Id string `json:"id"`
	// URL of the publicly available asset
	AssetUrl *string `json:"asset_url,omitempty"`
	ReportType SharedReportType `json:"report_type"`
	// Optional date/time that the shared report will expire. Default is 28 days after the time of creation.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ID of the study to which the report belongs
	StudyId *string `json:"study_id,omitempty"`
	// ID of the account that owns the study
	AccountId *string `json:"account_id,omitempty"`
	// Date and time of the object creation
	CreatedAt time.Time `json:"created_at"`
	// ID of the user who created the object
	CreatedBy string `json:"created_by"`
	// Date and time of the latest update to the object
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the user who last updated the object
	UpdatedBy *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SharedReportFull SharedReportFull

// NewSharedReportFull instantiates a new SharedReportFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedReportFull(id string, reportType SharedReportType, createdAt time.Time, createdBy string) *SharedReportFull {
	this := SharedReportFull{}
	this.Id = id
	this.ReportType = reportType
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	return &this
}

// NewSharedReportFullWithDefaults instantiates a new SharedReportFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedReportFullWithDefaults() *SharedReportFull {
	this := SharedReportFull{}
	return &this
}

// GetId returns the Id field value
func (o *SharedReportFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SharedReportFull) SetId(v string) {
	o.Id = v
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *SharedReportFull) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *SharedReportFull) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *SharedReportFull) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetReportType returns the ReportType field value
func (o *SharedReportFull) GetReportType() SharedReportType {
	if o == nil {
		var ret SharedReportType
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetReportTypeOk() (*SharedReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *SharedReportFull) SetReportType(v SharedReportType) {
	o.ReportType = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SharedReportFull) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SharedReportFull) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *SharedReportFull) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetStudyId returns the StudyId field value if set, zero value otherwise.
func (o *SharedReportFull) GetStudyId() string {
	if o == nil || IsNil(o.StudyId) {
		var ret string
		return ret
	}
	return *o.StudyId
}

// GetStudyIdOk returns a tuple with the StudyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetStudyIdOk() (*string, bool) {
	if o == nil || IsNil(o.StudyId) {
		return nil, false
	}
	return o.StudyId, true
}

// HasStudyId returns a boolean if a field has been set.
func (o *SharedReportFull) HasStudyId() bool {
	if o != nil && !IsNil(o.StudyId) {
		return true
	}

	return false
}

// SetStudyId gets a reference to the given string and assigns it to the StudyId field.
func (o *SharedReportFull) SetStudyId(v string) {
	o.StudyId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SharedReportFull) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SharedReportFull) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SharedReportFull) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SharedReportFull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SharedReportFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SharedReportFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SharedReportFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SharedReportFull) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SharedReportFull) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SharedReportFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SharedReportFull) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedReportFull) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SharedReportFull) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SharedReportFull) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o SharedReportFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedReportFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.AssetUrl) {
		toSerialize["asset_url"] = o.AssetUrl
	}
	toSerialize["report_type"] = o.ReportType
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.StudyId) {
		toSerialize["study_id"] = o.StudyId
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
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

func (o *SharedReportFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"report_type",
		"created_at",
		"created_by",
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

	varSharedReportFull := _SharedReportFull{}

	err = json.Unmarshal(data, &varSharedReportFull)

	if err != nil {
		return err
	}

	*o = SharedReportFull(varSharedReportFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "asset_url")
		delete(additionalProperties, "report_type")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "study_id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharedReportFull struct {
	value *SharedReportFull
	isSet bool
}

func (v NullableSharedReportFull) Get() *SharedReportFull {
	return v.value
}

func (v *NullableSharedReportFull) Set(val *SharedReportFull) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedReportFull) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedReportFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedReportFull(val *SharedReportFull) *NullableSharedReportFull {
	return &NullableSharedReportFull{value: val, isSet: true}
}

func (v NullableSharedReportFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedReportFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


