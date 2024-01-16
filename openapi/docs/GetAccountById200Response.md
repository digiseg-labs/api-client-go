# GetAccountById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AccountFull**](AccountFull.md) |  | [optional] 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 

## Methods

### NewGetAccountById200Response

`func NewGetAccountById200Response() *GetAccountById200Response`

NewGetAccountById200Response instantiates a new GetAccountById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountById200ResponseWithDefaults

`func NewGetAccountById200ResponseWithDefaults() *GetAccountById200Response`

NewGetAccountById200ResponseWithDefaults instantiates a new GetAccountById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAccountById200Response) GetData() AccountFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountById200Response) GetDataOk() (*AccountFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountById200Response) SetData(v AccountFull)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountById200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetAccountById200Response) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAccountById200Response) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAccountById200Response) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAccountById200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


