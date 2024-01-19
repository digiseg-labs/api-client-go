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

// checks if the CampaignEventLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignEventLinks{}

// CampaignEventLinks Links to trigger activity and events in this campaign. Two events are well-known and available in every campaign: impression and click. 
type CampaignEventLinks struct {
	Impression *CampaignEventLink `json:"impression,omitempty"`
	Click *CampaignEventLink `json:"click,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignEventLinks CampaignEventLinks

// NewCampaignEventLinks instantiates a new CampaignEventLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignEventLinks() *CampaignEventLinks {
	this := CampaignEventLinks{}
	return &this
}

// NewCampaignEventLinksWithDefaults instantiates a new CampaignEventLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignEventLinksWithDefaults() *CampaignEventLinks {
	this := CampaignEventLinks{}
	return &this
}

// GetImpression returns the Impression field value if set, zero value otherwise.
func (o *CampaignEventLinks) GetImpression() CampaignEventLink {
	if o == nil || IsNil(o.Impression) {
		var ret CampaignEventLink
		return ret
	}
	return *o.Impression
}

// GetImpressionOk returns a tuple with the Impression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignEventLinks) GetImpressionOk() (*CampaignEventLink, bool) {
	if o == nil || IsNil(o.Impression) {
		return nil, false
	}
	return o.Impression, true
}

// HasImpression returns a boolean if a field has been set.
func (o *CampaignEventLinks) HasImpression() bool {
	if o != nil && !IsNil(o.Impression) {
		return true
	}

	return false
}

// SetImpression gets a reference to the given CampaignEventLink and assigns it to the Impression field.
func (o *CampaignEventLinks) SetImpression(v CampaignEventLink) {
	o.Impression = &v
}

// GetClick returns the Click field value if set, zero value otherwise.
func (o *CampaignEventLinks) GetClick() CampaignEventLink {
	if o == nil || IsNil(o.Click) {
		var ret CampaignEventLink
		return ret
	}
	return *o.Click
}

// GetClickOk returns a tuple with the Click field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignEventLinks) GetClickOk() (*CampaignEventLink, bool) {
	if o == nil || IsNil(o.Click) {
		return nil, false
	}
	return o.Click, true
}

// HasClick returns a boolean if a field has been set.
func (o *CampaignEventLinks) HasClick() bool {
	if o != nil && !IsNil(o.Click) {
		return true
	}

	return false
}

// SetClick gets a reference to the given CampaignEventLink and assigns it to the Click field.
func (o *CampaignEventLinks) SetClick(v CampaignEventLink) {
	o.Click = &v
}

func (o CampaignEventLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignEventLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Impression) {
		toSerialize["impression"] = o.Impression
	}
	if !IsNil(o.Click) {
		toSerialize["click"] = o.Click
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignEventLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignEventLinks := _CampaignEventLinks{}

	err = json.Unmarshal(bytes, &varCampaignEventLinks)

	if err != nil {
		return err
	}

	*o = CampaignEventLinks(varCampaignEventLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "impression")
		delete(additionalProperties, "click")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignEventLinks struct {
	value *CampaignEventLinks
	isSet bool
}

func (v NullableCampaignEventLinks) Get() *CampaignEventLinks {
	return v.value
}

func (v *NullableCampaignEventLinks) Set(val *CampaignEventLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignEventLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignEventLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignEventLinks(val *CampaignEventLinks) *NullableCampaignEventLinks {
	return &NullableCampaignEventLinks{value: val, isSet: true}
}

func (v NullableCampaignEventLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignEventLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


