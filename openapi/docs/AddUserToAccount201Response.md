# AddUserToAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserFull**](UserFull.md) |  | [optional] 
**Links** | Pointer to [**UserLinks**](UserLinks.md) |  | [optional] 

## Methods

### NewAddUserToAccount201Response

`func NewAddUserToAccount201Response() *AddUserToAccount201Response`

NewAddUserToAccount201Response instantiates a new AddUserToAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserToAccount201ResponseWithDefaults

`func NewAddUserToAccount201ResponseWithDefaults() *AddUserToAccount201Response`

NewAddUserToAccount201ResponseWithDefaults instantiates a new AddUserToAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddUserToAccount201Response) GetData() UserFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddUserToAccount201Response) GetDataOk() (*UserFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddUserToAccount201Response) SetData(v UserFull)`

SetData sets Data field to given value.

### HasData

`func (o *AddUserToAccount201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AddUserToAccount201Response) GetLinks() UserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AddUserToAccount201Response) GetLinksOk() (*UserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AddUserToAccount201Response) SetLinks(v UserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AddUserToAccount201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


