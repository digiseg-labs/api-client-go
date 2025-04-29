# AudienceRecommendationResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | The ID of the conversation that this message pertains to. | 
**MessageText** | **string** | The text message sent back to the audience recommendation conversation. | 
**Recommendations** | Pointer to [**[]AudienceRecommendationRecommendationItem**](AudienceRecommendationRecommendationItem.md) |  | [optional] 

## Methods

### NewAudienceRecommendationResponseMessage

`func NewAudienceRecommendationResponseMessage(conversationId string, messageText string, ) *AudienceRecommendationResponseMessage`

NewAudienceRecommendationResponseMessage instantiates a new AudienceRecommendationResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecommendationResponseMessageWithDefaults

`func NewAudienceRecommendationResponseMessageWithDefaults() *AudienceRecommendationResponseMessage`

NewAudienceRecommendationResponseMessageWithDefaults instantiates a new AudienceRecommendationResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *AudienceRecommendationResponseMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AudienceRecommendationResponseMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AudienceRecommendationResponseMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetMessageText

`func (o *AudienceRecommendationResponseMessage) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *AudienceRecommendationResponseMessage) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *AudienceRecommendationResponseMessage) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.


### GetRecommendations

`func (o *AudienceRecommendationResponseMessage) GetRecommendations() []AudienceRecommendationRecommendationItem`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *AudienceRecommendationResponseMessage) GetRecommendationsOk() (*[]AudienceRecommendationRecommendationItem, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *AudienceRecommendationResponseMessage) SetRecommendations(v []AudienceRecommendationRecommendationItem)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *AudienceRecommendationResponseMessage) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


