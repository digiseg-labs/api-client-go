# CountryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 2-letter country code | 
**DisplayName** | **string** | Name of the country | 
**ReachStats** | Pointer to [**CountryReachStats**](CountryReachStats.md) |  | [optional] 

## Methods

### NewCountryItem

`func NewCountryItem(code string, displayName string, ) *CountryItem`

NewCountryItem instantiates a new CountryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryItemWithDefaults

`func NewCountryItemWithDefaults() *CountryItem`

NewCountryItemWithDefaults instantiates a new CountryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CountryItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *CountryItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CountryItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CountryItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetReachStats

`func (o *CountryItem) GetReachStats() CountryReachStats`

GetReachStats returns the ReachStats field if non-nil, zero value otherwise.

### GetReachStatsOk

`func (o *CountryItem) GetReachStatsOk() (*CountryReachStats, bool)`

GetReachStatsOk returns a tuple with the ReachStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachStats

`func (o *CountryItem) SetReachStats(v CountryReachStats)`

SetReachStats sets ReachStats field to given value.

### HasReachStats

`func (o *CountryItem) HasReachStats() bool`

HasReachStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


