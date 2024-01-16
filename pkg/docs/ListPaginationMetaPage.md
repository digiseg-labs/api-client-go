# ListPaginationMetaPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total amount of elements in the list (the returned &#x60;data&#x60; can be paginated) | [optional] 
**FirstCursor** | Pointer to **string** |  | [optional] 
**LastCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListPaginationMetaPage

`func NewListPaginationMetaPage() *ListPaginationMetaPage`

NewListPaginationMetaPage instantiates a new ListPaginationMetaPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaginationMetaPageWithDefaults

`func NewListPaginationMetaPageWithDefaults() *ListPaginationMetaPage`

NewListPaginationMetaPageWithDefaults instantiates a new ListPaginationMetaPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListPaginationMetaPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListPaginationMetaPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListPaginationMetaPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListPaginationMetaPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFirstCursor

`func (o *ListPaginationMetaPage) GetFirstCursor() string`

GetFirstCursor returns the FirstCursor field if non-nil, zero value otherwise.

### GetFirstCursorOk

`func (o *ListPaginationMetaPage) GetFirstCursorOk() (*string, bool)`

GetFirstCursorOk returns a tuple with the FirstCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCursor

`func (o *ListPaginationMetaPage) SetFirstCursor(v string)`

SetFirstCursor sets FirstCursor field to given value.

### HasFirstCursor

`func (o *ListPaginationMetaPage) HasFirstCursor() bool`

HasFirstCursor returns a boolean if a field has been set.

### GetLastCursor

`func (o *ListPaginationMetaPage) GetLastCursor() string`

GetLastCursor returns the LastCursor field if non-nil, zero value otherwise.

### GetLastCursorOk

`func (o *ListPaginationMetaPage) GetLastCursorOk() (*string, bool)`

GetLastCursorOk returns a tuple with the LastCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCursor

`func (o *ListPaginationMetaPage) SetLastCursor(v string)`

SetLastCursor sets LastCursor field to given value.

### HasLastCursor

`func (o *ListPaginationMetaPage) HasLastCursor() bool`

HasLastCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


