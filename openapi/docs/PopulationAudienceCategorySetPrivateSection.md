# PopulationAudienceCategorySetPrivateSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FractionOfTotal** | Pointer to **float64** | The fraction of events that fall within this object compared to the total of the category or segment (usually represented by the measurement&#39;s parent&#39;s parent). For example, if the measurement is \&quot;impression\&quot; on the &#x60;home_type&#x60; \&quot;Apartment\&quot; object, then the &#x60;fraction_of_total&#x60; represents the number of impressions on apartments compared to impressions from other &#x60;home_type&#x60; values.  | [optional] 
**AudienceCategories** | Pointer to **map[string]map[string]float64** | An object with category codes as keys, objects with audience codes and fractions of totals as keys. | [optional] 

## Methods

### NewPopulationAudienceCategorySetPrivateSection

`func NewPopulationAudienceCategorySetPrivateSection() *PopulationAudienceCategorySetPrivateSection`

NewPopulationAudienceCategorySetPrivateSection instantiates a new PopulationAudienceCategorySetPrivateSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationAudienceCategorySetPrivateSectionWithDefaults

`func NewPopulationAudienceCategorySetPrivateSectionWithDefaults() *PopulationAudienceCategorySetPrivateSection`

NewPopulationAudienceCategorySetPrivateSectionWithDefaults instantiates a new PopulationAudienceCategorySetPrivateSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFractionOfTotal

`func (o *PopulationAudienceCategorySetPrivateSection) GetFractionOfTotal() float64`

GetFractionOfTotal returns the FractionOfTotal field if non-nil, zero value otherwise.

### GetFractionOfTotalOk

`func (o *PopulationAudienceCategorySetPrivateSection) GetFractionOfTotalOk() (*float64, bool)`

GetFractionOfTotalOk returns a tuple with the FractionOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionOfTotal

`func (o *PopulationAudienceCategorySetPrivateSection) SetFractionOfTotal(v float64)`

SetFractionOfTotal sets FractionOfTotal field to given value.

### HasFractionOfTotal

`func (o *PopulationAudienceCategorySetPrivateSection) HasFractionOfTotal() bool`

HasFractionOfTotal returns a boolean if a field has been set.

### GetAudienceCategories

`func (o *PopulationAudienceCategorySetPrivateSection) GetAudienceCategories() map[string]map[string]float64`

GetAudienceCategories returns the AudienceCategories field if non-nil, zero value otherwise.

### GetAudienceCategoriesOk

`func (o *PopulationAudienceCategorySetPrivateSection) GetAudienceCategoriesOk() (*map[string]map[string]float64, bool)`

GetAudienceCategoriesOk returns a tuple with the AudienceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceCategories

`func (o *PopulationAudienceCategorySetPrivateSection) SetAudienceCategories(v map[string]map[string]float64)`

SetAudienceCategories sets AudienceCategories field to given value.

### HasAudienceCategories

`func (o *PopulationAudienceCategorySetPrivateSection) HasAudienceCategories() bool`

HasAudienceCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


