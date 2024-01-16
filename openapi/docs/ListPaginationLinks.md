# ListPaginationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to **string** | Link to the first page of the list | [optional] 
**Next** | Pointer to **string** | Link to the next page of the list | [optional] 

## Methods

### NewListPaginationLinks

`func NewListPaginationLinks() *ListPaginationLinks`

NewListPaginationLinks instantiates a new ListPaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaginationLinksWithDefaults

`func NewListPaginationLinksWithDefaults() *ListPaginationLinks`

NewListPaginationLinksWithDefaults instantiates a new ListPaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *ListPaginationLinks) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ListPaginationLinks) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ListPaginationLinks) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *ListPaginationLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetNext

`func (o *ListPaginationLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListPaginationLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListPaginationLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListPaginationLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


