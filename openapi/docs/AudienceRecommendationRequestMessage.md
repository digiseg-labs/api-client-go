# AudienceRecommendationRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** | The ID of the conversation to add this message to. Do not provide a conversation ID in the initial/first message. | [optional] 
**MessageText** | **string** | The text message to send to the audience recommendation conversation. | 

## Methods

### NewAudienceRecommendationRequestMessage

`func NewAudienceRecommendationRequestMessage(messageText string, ) *AudienceRecommendationRequestMessage`

NewAudienceRecommendationRequestMessage instantiates a new AudienceRecommendationRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecommendationRequestMessageWithDefaults

`func NewAudienceRecommendationRequestMessageWithDefaults() *AudienceRecommendationRequestMessage`

NewAudienceRecommendationRequestMessageWithDefaults instantiates a new AudienceRecommendationRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *AudienceRecommendationRequestMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AudienceRecommendationRequestMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AudienceRecommendationRequestMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *AudienceRecommendationRequestMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetMessageText

`func (o *AudienceRecommendationRequestMessage) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *AudienceRecommendationRequestMessage) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *AudienceRecommendationRequestMessage) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


