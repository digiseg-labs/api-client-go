# UserLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Link for getting to the user&#39;s account | [optional] 
**Apikeys** | Pointer to **string** | Link for getting to the user&#39;s api keys | [optional] 

## Methods

### NewUserLinks

`func NewUserLinks() *UserLinks`

NewUserLinks instantiates a new UserLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLinksWithDefaults

`func NewUserLinksWithDefaults() *UserLinks`

NewUserLinksWithDefaults instantiates a new UserLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserLinks) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserLinks) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserLinks) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserLinks) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApikeys

`func (o *UserLinks) GetApikeys() string`

GetApikeys returns the Apikeys field if non-nil, zero value otherwise.

### GetApikeysOk

`func (o *UserLinks) GetApikeysOk() (*string, bool)`

GetApikeysOk returns a tuple with the Apikeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikeys

`func (o *UserLinks) SetApikeys(v string)`

SetApikeys sets Apikeys field to given value.

### HasApikeys

`func (o *UserLinks) HasApikeys() bool`

HasApikeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


