# AccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object | [optional] 
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the logo of the account | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 

## Methods

### NewAccountItem

`func NewAccountItem() *AccountItem`

NewAccountItem instantiates a new AccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemWithDefaults

`func NewAccountItemWithDefaults() *AccountItem`

NewAccountItemWithDefaults instantiates a new AccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountItem) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetSlug

`func (o *AccountItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountItem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountItem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


