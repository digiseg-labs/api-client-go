# CategoryListResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to [**CountryItem**](CountryItem.md) |  | [optional] 
**Platform** | Pointer to [**AudiencePlatformItem**](AudiencePlatformItem.md) |  | [optional] 
**Page** | Pointer to [**ListPaginationMetaPage**](ListPaginationMetaPage.md) |  | [optional] 

## Methods

### NewCategoryListResponseMeta

`func NewCategoryListResponseMeta() *CategoryListResponseMeta`

NewCategoryListResponseMeta instantiates a new CategoryListResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryListResponseMetaWithDefaults

`func NewCategoryListResponseMetaWithDefaults() *CategoryListResponseMeta`

NewCategoryListResponseMetaWithDefaults instantiates a new CategoryListResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CategoryListResponseMeta) GetCountry() CountryItem`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CategoryListResponseMeta) GetCountryOk() (*CountryItem, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CategoryListResponseMeta) SetCountry(v CountryItem)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CategoryListResponseMeta) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPlatform

`func (o *CategoryListResponseMeta) GetPlatform() AudiencePlatformItem`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CategoryListResponseMeta) GetPlatformOk() (*AudiencePlatformItem, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CategoryListResponseMeta) SetPlatform(v AudiencePlatformItem)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CategoryListResponseMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPage

`func (o *CategoryListResponseMeta) GetPage() ListPaginationMetaPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CategoryListResponseMeta) GetPageOk() (*ListPaginationMetaPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CategoryListResponseMeta) SetPage(v ListPaginationMetaPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *CategoryListResponseMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


