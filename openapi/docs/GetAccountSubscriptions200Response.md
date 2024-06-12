# GetAccountSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AccountSubscriptionsSummary**](AccountSubscriptionsSummary.md) |  | [optional] 

## Methods

### NewGetAccountSubscriptions200Response

`func NewGetAccountSubscriptions200Response() *GetAccountSubscriptions200Response`

NewGetAccountSubscriptions200Response instantiates a new GetAccountSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSubscriptions200ResponseWithDefaults

`func NewGetAccountSubscriptions200ResponseWithDefaults() *GetAccountSubscriptions200Response`

NewGetAccountSubscriptions200ResponseWithDefaults instantiates a new GetAccountSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAccountSubscriptions200Response) GetData() AccountSubscriptionsSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountSubscriptions200Response) GetDataOk() (*AccountSubscriptionsSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountSubscriptions200Response) SetData(v AccountSubscriptionsSummary)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountSubscriptions200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


