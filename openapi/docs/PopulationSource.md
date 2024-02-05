# PopulationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A humanly readable name of the population | 
**Source** | **string** | Name or description of the data source | 
**Meta** | **map[string]interface{}** |  | 
**NotResolved** | Pointer to [**PopulationSourceNotResolvedSection**](PopulationSourceNotResolvedSection.md) |  | [optional] 
**Private** | [**PopulationSourcePrivateSection**](PopulationSourcePrivateSection.md) |  | 
**Business** | [**PopulationSourceBusinessSection**](PopulationSourceBusinessSection.md) |  | 

## Methods

### NewPopulationSource

`func NewPopulationSource(name string, source string, meta map[string]interface{}, private PopulationSourcePrivateSection, business PopulationSourceBusinessSection, ) *PopulationSource`

NewPopulationSource instantiates a new PopulationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationSourceWithDefaults

`func NewPopulationSourceWithDefaults() *PopulationSource`

NewPopulationSourceWithDefaults instantiates a new PopulationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PopulationSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PopulationSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PopulationSource) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *PopulationSource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PopulationSource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PopulationSource) SetSource(v string)`

SetSource sets Source field to given value.


### GetMeta

`func (o *PopulationSource) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PopulationSource) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PopulationSource) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetNotResolved

`func (o *PopulationSource) GetNotResolved() PopulationSourceNotResolvedSection`

GetNotResolved returns the NotResolved field if non-nil, zero value otherwise.

### GetNotResolvedOk

`func (o *PopulationSource) GetNotResolvedOk() (*PopulationSourceNotResolvedSection, bool)`

GetNotResolvedOk returns a tuple with the NotResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResolved

`func (o *PopulationSource) SetNotResolved(v PopulationSourceNotResolvedSection)`

SetNotResolved sets NotResolved field to given value.

### HasNotResolved

`func (o *PopulationSource) HasNotResolved() bool`

HasNotResolved returns a boolean if a field has been set.

### GetPrivate

`func (o *PopulationSource) GetPrivate() PopulationSourcePrivateSection`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PopulationSource) GetPrivateOk() (*PopulationSourcePrivateSection, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PopulationSource) SetPrivate(v PopulationSourcePrivateSection)`

SetPrivate sets Private field to given value.


### GetBusiness

`func (o *PopulationSource) GetBusiness() PopulationSourceBusinessSection`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *PopulationSource) GetBusinessOk() (*PopulationSourceBusinessSection, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *PopulationSource) SetBusiness(v PopulationSourceBusinessSection)`

SetBusiness sets Business field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


