# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The event that triggered the measurement, typically &#x60;impression&#x60; or &#x60;click&#x60; | [optional] 
**Count** | Pointer to **int32** | The real value of the measurement, typically a counter value (integer) | [optional] 
**FractionOfTotal** | Pointer to **float64** | The fraction of events that fall within this object compared to the total of the category or segment (usually represented by the measurement&#39;s parent&#39;s parent). For example, if the measurement is \&quot;impression\&quot; on the &#x60;home_type&#x60; \&quot;Apartment\&quot; object, then the &#x60;fraction_of_total&#x60; represents the number of impressions on apartments compared to impressions from other &#x60;home_type&#x60; values.  | [optional] 
**ConversionRate** | Pointer to **float64** | The rate of conversion to this measurement. Typically applies to measurements like \&quot;click\&quot; where it will represent the rate of impressions that turn into a click.  | [optional] 

## Methods

### NewMeasurement

`func NewMeasurement() *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *Measurement) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Measurement) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Measurement) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Measurement) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetCount

`func (o *Measurement) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Measurement) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Measurement) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Measurement) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFractionOfTotal

`func (o *Measurement) GetFractionOfTotal() float64`

GetFractionOfTotal returns the FractionOfTotal field if non-nil, zero value otherwise.

### GetFractionOfTotalOk

`func (o *Measurement) GetFractionOfTotalOk() (*float64, bool)`

GetFractionOfTotalOk returns a tuple with the FractionOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionOfTotal

`func (o *Measurement) SetFractionOfTotal(v float64)`

SetFractionOfTotal sets FractionOfTotal field to given value.

### HasFractionOfTotal

`func (o *Measurement) HasFractionOfTotal() bool`

HasFractionOfTotal returns a boolean if a field has been set.

### GetConversionRate

`func (o *Measurement) GetConversionRate() float64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *Measurement) GetConversionRateOk() (*float64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *Measurement) SetConversionRate(v float64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *Measurement) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


