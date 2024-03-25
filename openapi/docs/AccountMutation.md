# AccountMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the account | [optional] 
**LogoUrl** | Pointer to **string** | The URL to the logo of the account | [optional] 
**Slug** | Pointer to **string** | A short human-readable name to identify the account. Must be lower-case and between 4 and 16 characters. | [optional] 
**OwnerId** | Pointer to **string** | ID of the user who is the ultimate owner of the account | [optional] 

## Methods

### NewAccountMutation

`func NewAccountMutation() *AccountMutation`

NewAccountMutation instantiates a new AccountMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMutationWithDefaults

`func NewAccountMutationWithDefaults() *AccountMutation`

NewAccountMutationWithDefaults instantiates a new AccountMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountMutation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountMutation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountMutation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountMutation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountMutation) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountMutation) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountMutation) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountMutation) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetSlug

`func (o *AccountMutation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountMutation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountMutation) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountMutation) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccountMutation) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccountMutation) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccountMutation) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccountMutation) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


