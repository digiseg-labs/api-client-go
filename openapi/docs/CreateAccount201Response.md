# CreateAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AccountFull**](AccountFull.md) |  | [optional] 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 

## Methods

### NewCreateAccount201Response

`func NewCreateAccount201Response() *CreateAccount201Response`

NewCreateAccount201Response instantiates a new CreateAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccount201ResponseWithDefaults

`func NewCreateAccount201ResponseWithDefaults() *CreateAccount201Response`

NewCreateAccount201ResponseWithDefaults instantiates a new CreateAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateAccount201Response) GetData() AccountFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateAccount201Response) GetDataOk() (*AccountFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateAccount201Response) SetData(v AccountFull)`

SetData sets Data field to given value.

### HasData

`func (o *CreateAccount201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CreateAccount201Response) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateAccount201Response) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateAccount201Response) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateAccount201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


