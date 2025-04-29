# AudienceSetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to represent the audience set | 
**DisplayName** | **string** |  | 
**Categories** | [**[]AudienceCategoryItem**](AudienceCategoryItem.md) |  | 

## Methods

### NewAudienceSetItem

`func NewAudienceSetItem(code string, displayName string, categories []AudienceCategoryItem, ) *AudienceSetItem`

NewAudienceSetItem instantiates a new AudienceSetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSetItemWithDefaults

`func NewAudienceSetItemWithDefaults() *AudienceSetItem`

NewAudienceSetItemWithDefaults instantiates a new AudienceSetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AudienceSetItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudienceSetItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudienceSetItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *AudienceSetItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AudienceSetItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AudienceSetItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCategories

`func (o *AudienceSetItem) GetCategories() []AudienceCategoryItem`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AudienceSetItem) GetCategoriesOk() (*[]AudienceCategoryItem, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AudienceSetItem) SetCategories(v []AudienceCategoryItem)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


