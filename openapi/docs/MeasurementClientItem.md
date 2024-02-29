# MeasurementClientItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object | 
**Name** | Pointer to **string** | The name of the measurement client | [optional] 
**AccountId** | Pointer to **string** | The ID of the account that owns this client | [optional] [readonly] 
**LogoUrl** | Pointer to **string** | The URL of an image representing the measurement client | [optional] [readonly] 

## Methods

### NewMeasurementClientItem

`func NewMeasurementClientItem(id string, ) *MeasurementClientItem`

NewMeasurementClientItem instantiates a new MeasurementClientItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementClientItemWithDefaults

`func NewMeasurementClientItemWithDefaults() *MeasurementClientItem`

NewMeasurementClientItemWithDefaults instantiates a new MeasurementClientItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeasurementClientItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeasurementClientItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeasurementClientItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MeasurementClientItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeasurementClientItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeasurementClientItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeasurementClientItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *MeasurementClientItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MeasurementClientItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MeasurementClientItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MeasurementClientItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLogoUrl

`func (o *MeasurementClientItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *MeasurementClientItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *MeasurementClientItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *MeasurementClientItem) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


