# SendEmailQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The ID of the email template to use. | 
**ModelData** | Pointer to **map[string]string** | Model data to pass to the email template. Note that the following model data fields will be automatically populated by the service:  * &#x60;user_id&#x60;: The ID of the recipient user  * &#x60;user_name&#x60;: The name of the recipient user  * &#x60;user_email&#x60;: The email of the recipient user  * &#x60;account_id&#x60;: The ID of the recipient user&#39;s account  * &#x60;account_name&#x60;: The name of the recipient user&#39;s account  | [optional] 
**IncludeUsers** | [**SendEmailUserFilters**](SendEmailUserFilters.md) |  | 

## Methods

### NewSendEmailQuery

`func NewSendEmailQuery(templateId string, includeUsers SendEmailUserFilters, ) *SendEmailQuery`

NewSendEmailQuery instantiates a new SendEmailQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailQueryWithDefaults

`func NewSendEmailQueryWithDefaults() *SendEmailQuery`

NewSendEmailQueryWithDefaults instantiates a new SendEmailQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *SendEmailQuery) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendEmailQuery) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendEmailQuery) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetModelData

`func (o *SendEmailQuery) GetModelData() map[string]string`

GetModelData returns the ModelData field if non-nil, zero value otherwise.

### GetModelDataOk

`func (o *SendEmailQuery) GetModelDataOk() (*map[string]string, bool)`

GetModelDataOk returns a tuple with the ModelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelData

`func (o *SendEmailQuery) SetModelData(v map[string]string)`

SetModelData sets ModelData field to given value.

### HasModelData

`func (o *SendEmailQuery) HasModelData() bool`

HasModelData returns a boolean if a field has been set.

### GetIncludeUsers

`func (o *SendEmailQuery) GetIncludeUsers() SendEmailUserFilters`

GetIncludeUsers returns the IncludeUsers field if non-nil, zero value otherwise.

### GetIncludeUsersOk

`func (o *SendEmailQuery) GetIncludeUsersOk() (*SendEmailUserFilters, bool)`

GetIncludeUsersOk returns a tuple with the IncludeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUsers

`func (o *SendEmailQuery) SetIncludeUsers(v SendEmailUserFilters)`

SetIncludeUsers sets IncludeUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


