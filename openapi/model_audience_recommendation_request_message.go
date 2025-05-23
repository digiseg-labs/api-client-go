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
	"fmt"
)

// checks if the AudienceRecommendationRequestMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceRecommendationRequestMessage{}

// AudienceRecommendationRequestMessage A user's message in an audience recommendation conversation
type AudienceRecommendationRequestMessage struct {
	// The ID of the conversation to add this message to. Do not provide a conversation ID in the initial/first message.
	ConversationId *string `json:"conversation_id,omitempty"`
	// The text message to send to the audience recommendation conversation.
	MessageText string `json:"message_text"`
	AdditionalProperties map[string]interface{}
}

type _AudienceRecommendationRequestMessage AudienceRecommendationRequestMessage

// NewAudienceRecommendationRequestMessage instantiates a new AudienceRecommendationRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceRecommendationRequestMessage(messageText string) *AudienceRecommendationRequestMessage {
	this := AudienceRecommendationRequestMessage{}
	this.MessageText = messageText
	return &this
}

// NewAudienceRecommendationRequestMessageWithDefaults instantiates a new AudienceRecommendationRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceRecommendationRequestMessageWithDefaults() *AudienceRecommendationRequestMessage {
	this := AudienceRecommendationRequestMessage{}
	return &this
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *AudienceRecommendationRequestMessage) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceRecommendationRequestMessage) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *AudienceRecommendationRequestMessage) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *AudienceRecommendationRequestMessage) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetMessageText returns the MessageText field value
func (o *AudienceRecommendationRequestMessage) GetMessageText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageText
}

// GetMessageTextOk returns a tuple with the MessageText field value
// and a boolean to check if the value has been set.
func (o *AudienceRecommendationRequestMessage) GetMessageTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageText, true
}

// SetMessageText sets field value
func (o *AudienceRecommendationRequestMessage) SetMessageText(v string) {
	o.MessageText = v
}

func (o AudienceRecommendationRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceRecommendationRequestMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	toSerialize["message_text"] = o.MessageText

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AudienceRecommendationRequestMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_text",
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

	varAudienceRecommendationRequestMessage := _AudienceRecommendationRequestMessage{}

	err = json.Unmarshal(data, &varAudienceRecommendationRequestMessage)

	if err != nil {
		return err
	}

	*o = AudienceRecommendationRequestMessage(varAudienceRecommendationRequestMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conversation_id")
		delete(additionalProperties, "message_text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAudienceRecommendationRequestMessage struct {
	value *AudienceRecommendationRequestMessage
	isSet bool
}

func (v NullableAudienceRecommendationRequestMessage) Get() *AudienceRecommendationRequestMessage {
	return v.value
}

func (v *NullableAudienceRecommendationRequestMessage) Set(val *AudienceRecommendationRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceRecommendationRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceRecommendationRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceRecommendationRequestMessage(val *AudienceRecommendationRequestMessage) *NullableAudienceRecommendationRequestMessage {
	return &NullableAudienceRecommendationRequestMessage{value: val, isSet: true}
}

func (v NullableAudienceRecommendationRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceRecommendationRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


