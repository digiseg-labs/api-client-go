# EnrichmentJobSourceDataPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]EnrichmentJobSourceDataColumnDefinition**](EnrichmentJobSourceDataColumnDefinition.md) |  | [optional] 
**Rows** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewEnrichmentJobSourceDataPreview

`func NewEnrichmentJobSourceDataPreview() *EnrichmentJobSourceDataPreview`

NewEnrichmentJobSourceDataPreview instantiates a new EnrichmentJobSourceDataPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentJobSourceDataPreviewWithDefaults

`func NewEnrichmentJobSourceDataPreviewWithDefaults() *EnrichmentJobSourceDataPreview`

NewEnrichmentJobSourceDataPreviewWithDefaults instantiates a new EnrichmentJobSourceDataPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *EnrichmentJobSourceDataPreview) GetColumns() []EnrichmentJobSourceDataColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *EnrichmentJobSourceDataPreview) GetColumnsOk() (*[]EnrichmentJobSourceDataColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *EnrichmentJobSourceDataPreview) SetColumns(v []EnrichmentJobSourceDataColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *EnrichmentJobSourceDataPreview) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *EnrichmentJobSourceDataPreview) GetRows() [][]string`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EnrichmentJobSourceDataPreview) GetRowsOk() (*[][]string, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EnrichmentJobSourceDataPreview) SetRows(v [][]string)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EnrichmentJobSourceDataPreview) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


