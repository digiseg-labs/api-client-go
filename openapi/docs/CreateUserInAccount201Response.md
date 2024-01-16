# CreateUserInAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserFull**](UserFull.md) |  | [optional] 
**Links** | Pointer to [**UserLinks**](UserLinks.md) |  | [optional] 

## Methods

### NewCreateUserInAccount201Response

`func NewCreateUserInAccount201Response() *CreateUserInAccount201Response`

NewCreateUserInAccount201Response instantiates a new CreateUserInAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserInAccount201ResponseWithDefaults

`func NewCreateUserInAccount201ResponseWithDefaults() *CreateUserInAccount201Response`

NewCreateUserInAccount201ResponseWithDefaults instantiates a new CreateUserInAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateUserInAccount201Response) GetData() UserFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateUserInAccount201Response) GetDataOk() (*UserFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateUserInAccount201Response) SetData(v UserFull)`

SetData sets Data field to given value.

### HasData

`func (o *CreateUserInAccount201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CreateUserInAccount201Response) GetLinks() UserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateUserInAccount201Response) GetLinksOk() (*UserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateUserInAccount201Response) SetLinks(v UserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateUserInAccount201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


