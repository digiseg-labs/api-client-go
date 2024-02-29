# PopulationFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A humanly readable name of the population | [optional] 
**Private** | Pointer to [**PopulationAudienceCategorySetPrivateSection**](PopulationAudienceCategorySetPrivateSection.md) |  | [optional] 
**Business** | Pointer to [**PopulationAudienceCategorySetBusinessSection**](PopulationAudienceCategorySetBusinessSection.md) |  | [optional] 

## Methods

### NewPopulationFull

`func NewPopulationFull() *PopulationFull`

NewPopulationFull instantiates a new PopulationFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationFullWithDefaults

`func NewPopulationFullWithDefaults() *PopulationFull`

NewPopulationFullWithDefaults instantiates a new PopulationFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PopulationFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PopulationFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PopulationFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PopulationFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *PopulationFull) GetPrivate() PopulationAudienceCategorySetPrivateSection`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PopulationFull) GetPrivateOk() (*PopulationAudienceCategorySetPrivateSection, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PopulationFull) SetPrivate(v PopulationAudienceCategorySetPrivateSection)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *PopulationFull) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetBusiness

`func (o *PopulationFull) GetBusiness() PopulationAudienceCategorySetBusinessSection`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *PopulationFull) GetBusinessOk() (*PopulationAudienceCategorySetBusinessSection, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *PopulationFull) SetBusiness(v PopulationAudienceCategorySetBusinessSection)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *PopulationFull) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


