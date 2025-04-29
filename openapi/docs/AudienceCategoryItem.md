# AudienceCategoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to represent the audience category | 
**DisplayName** | **string** |  | 
**Audiences** | [**[]AudienceItem**](AudienceItem.md) |  | 

## Methods

### NewAudienceCategoryItem

`func NewAudienceCategoryItem(code string, displayName string, audiences []AudienceItem, ) *AudienceCategoryItem`

NewAudienceCategoryItem instantiates a new AudienceCategoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceCategoryItemWithDefaults

`func NewAudienceCategoryItemWithDefaults() *AudienceCategoryItem`

NewAudienceCategoryItemWithDefaults instantiates a new AudienceCategoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AudienceCategoryItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudienceCategoryItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudienceCategoryItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *AudienceCategoryItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AudienceCategoryItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AudienceCategoryItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAudiences

`func (o *AudienceCategoryItem) GetAudiences() []AudienceItem`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *AudienceCategoryItem) GetAudiencesOk() (*[]AudienceItem, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *AudienceCategoryItem) SetAudiences(v []AudienceItem)`

SetAudiences sets Audiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


