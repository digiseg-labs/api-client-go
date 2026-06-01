# AudienceBuilderSegmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Unique code for the audience builder segment. | 
**ShortName** | **string** | Short display name for the segment. | 
**SegmentReach** | **float32** | Estimated reach for the segment as a fraction of the audience. | 
**InTargetReach** | **float32** | Estimated in-target reach for the segment as a fraction of the audience. | 

## Methods

### NewAudienceBuilderSegmentItem

`func NewAudienceBuilderSegmentItem(code string, shortName string, segmentReach float32, inTargetReach float32, ) *AudienceBuilderSegmentItem`

NewAudienceBuilderSegmentItem instantiates a new AudienceBuilderSegmentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceBuilderSegmentItemWithDefaults

`func NewAudienceBuilderSegmentItemWithDefaults() *AudienceBuilderSegmentItem`

NewAudienceBuilderSegmentItemWithDefaults instantiates a new AudienceBuilderSegmentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AudienceBuilderSegmentItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AudienceBuilderSegmentItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AudienceBuilderSegmentItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetShortName

`func (o *AudienceBuilderSegmentItem) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *AudienceBuilderSegmentItem) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *AudienceBuilderSegmentItem) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetSegmentReach

`func (o *AudienceBuilderSegmentItem) GetSegmentReach() float32`

GetSegmentReach returns the SegmentReach field if non-nil, zero value otherwise.

### GetSegmentReachOk

`func (o *AudienceBuilderSegmentItem) GetSegmentReachOk() (*float32, bool)`

GetSegmentReachOk returns a tuple with the SegmentReach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentReach

`func (o *AudienceBuilderSegmentItem) SetSegmentReach(v float32)`

SetSegmentReach sets SegmentReach field to given value.


### GetInTargetReach

`func (o *AudienceBuilderSegmentItem) GetInTargetReach() float32`

GetInTargetReach returns the InTargetReach field if non-nil, zero value otherwise.

### GetInTargetReachOk

`func (o *AudienceBuilderSegmentItem) GetInTargetReachOk() (*float32, bool)`

GetInTargetReachOk returns a tuple with the InTargetReach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTargetReach

`func (o *AudienceBuilderSegmentItem) SetInTargetReach(v float32)`

SetInTargetReach sets InTargetReach field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


