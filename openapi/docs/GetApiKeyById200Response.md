# GetApiKeyById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiKeyFull**](ApiKeyFull.md) |  | [optional] 
**Links** | Pointer to [**ApiKeyLinks**](ApiKeyLinks.md) |  | [optional] 

## Methods

### NewGetApiKeyById200Response

`func NewGetApiKeyById200Response() *GetApiKeyById200Response`

NewGetApiKeyById200Response instantiates a new GetApiKeyById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeyById200ResponseWithDefaults

`func NewGetApiKeyById200ResponseWithDefaults() *GetApiKeyById200Response`

NewGetApiKeyById200ResponseWithDefaults instantiates a new GetApiKeyById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetApiKeyById200Response) GetData() ApiKeyFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApiKeyById200Response) GetDataOk() (*ApiKeyFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApiKeyById200Response) SetData(v ApiKeyFull)`

SetData sets Data field to given value.

### HasData

`func (o *GetApiKeyById200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetApiKeyById200Response) GetLinks() ApiKeyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetApiKeyById200Response) GetLinksOk() (*ApiKeyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetApiKeyById200Response) SetLinks(v ApiKeyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetApiKeyById200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


