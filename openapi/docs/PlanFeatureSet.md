# PlanFeatureSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUsers** | **int32** |  | 
**MaxClients** | **int32** |  | 
**MaxActiveStudies** | **int32** |  | 
**MaxEventsPerStudy** | **int64** |  | 
**MaxStudyEventsPerMonth** | **int64** |  | 
**MaxAudienceLookupsPerMonth** | **int64** |  | 
**StudyAudienceSet** | [**LimitedOrFullFeature**](LimitedOrFullFeature.md) |  | 
**StudyEventSet** | [**LimitedOrFullFeature**](LimitedOrFullFeature.md) |  | 
**HasAudienceRecommendations** | **bool** |  | 
**HasReportCustomization** | **bool** |  | 
**HasReportSharingClients** | **bool** |  | 
**HasReportSharingPublic** | **bool** |  | 

## Methods

### NewPlanFeatureSet

`func NewPlanFeatureSet(maxUsers int32, maxClients int32, maxActiveStudies int32, maxEventsPerStudy int64, maxStudyEventsPerMonth int64, maxAudienceLookupsPerMonth int64, studyAudienceSet LimitedOrFullFeature, studyEventSet LimitedOrFullFeature, hasAudienceRecommendations bool, hasReportCustomization bool, hasReportSharingClients bool, hasReportSharingPublic bool, ) *PlanFeatureSet`

NewPlanFeatureSet instantiates a new PlanFeatureSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanFeatureSetWithDefaults

`func NewPlanFeatureSetWithDefaults() *PlanFeatureSet`

NewPlanFeatureSetWithDefaults instantiates a new PlanFeatureSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUsers

`func (o *PlanFeatureSet) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *PlanFeatureSet) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *PlanFeatureSet) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.


### GetMaxClients

`func (o *PlanFeatureSet) GetMaxClients() int32`

GetMaxClients returns the MaxClients field if non-nil, zero value otherwise.

### GetMaxClientsOk

`func (o *PlanFeatureSet) GetMaxClientsOk() (*int32, bool)`

GetMaxClientsOk returns a tuple with the MaxClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients

`func (o *PlanFeatureSet) SetMaxClients(v int32)`

SetMaxClients sets MaxClients field to given value.


### GetMaxActiveStudies

`func (o *PlanFeatureSet) GetMaxActiveStudies() int32`

GetMaxActiveStudies returns the MaxActiveStudies field if non-nil, zero value otherwise.

### GetMaxActiveStudiesOk

`func (o *PlanFeatureSet) GetMaxActiveStudiesOk() (*int32, bool)`

GetMaxActiveStudiesOk returns a tuple with the MaxActiveStudies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveStudies

`func (o *PlanFeatureSet) SetMaxActiveStudies(v int32)`

SetMaxActiveStudies sets MaxActiveStudies field to given value.


### GetMaxEventsPerStudy

`func (o *PlanFeatureSet) GetMaxEventsPerStudy() int64`

GetMaxEventsPerStudy returns the MaxEventsPerStudy field if non-nil, zero value otherwise.

### GetMaxEventsPerStudyOk

`func (o *PlanFeatureSet) GetMaxEventsPerStudyOk() (*int64, bool)`

GetMaxEventsPerStudyOk returns a tuple with the MaxEventsPerStudy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerStudy

`func (o *PlanFeatureSet) SetMaxEventsPerStudy(v int64)`

SetMaxEventsPerStudy sets MaxEventsPerStudy field to given value.


### GetMaxStudyEventsPerMonth

`func (o *PlanFeatureSet) GetMaxStudyEventsPerMonth() int64`

GetMaxStudyEventsPerMonth returns the MaxStudyEventsPerMonth field if non-nil, zero value otherwise.

### GetMaxStudyEventsPerMonthOk

`func (o *PlanFeatureSet) GetMaxStudyEventsPerMonthOk() (*int64, bool)`

GetMaxStudyEventsPerMonthOk returns a tuple with the MaxStudyEventsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStudyEventsPerMonth

`func (o *PlanFeatureSet) SetMaxStudyEventsPerMonth(v int64)`

SetMaxStudyEventsPerMonth sets MaxStudyEventsPerMonth field to given value.


### GetMaxAudienceLookupsPerMonth

`func (o *PlanFeatureSet) GetMaxAudienceLookupsPerMonth() int64`

GetMaxAudienceLookupsPerMonth returns the MaxAudienceLookupsPerMonth field if non-nil, zero value otherwise.

### GetMaxAudienceLookupsPerMonthOk

`func (o *PlanFeatureSet) GetMaxAudienceLookupsPerMonthOk() (*int64, bool)`

GetMaxAudienceLookupsPerMonthOk returns a tuple with the MaxAudienceLookupsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudienceLookupsPerMonth

`func (o *PlanFeatureSet) SetMaxAudienceLookupsPerMonth(v int64)`

SetMaxAudienceLookupsPerMonth sets MaxAudienceLookupsPerMonth field to given value.


### GetStudyAudienceSet

`func (o *PlanFeatureSet) GetStudyAudienceSet() LimitedOrFullFeature`

GetStudyAudienceSet returns the StudyAudienceSet field if non-nil, zero value otherwise.

### GetStudyAudienceSetOk

`func (o *PlanFeatureSet) GetStudyAudienceSetOk() (*LimitedOrFullFeature, bool)`

GetStudyAudienceSetOk returns a tuple with the StudyAudienceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyAudienceSet

`func (o *PlanFeatureSet) SetStudyAudienceSet(v LimitedOrFullFeature)`

SetStudyAudienceSet sets StudyAudienceSet field to given value.


### GetStudyEventSet

`func (o *PlanFeatureSet) GetStudyEventSet() LimitedOrFullFeature`

GetStudyEventSet returns the StudyEventSet field if non-nil, zero value otherwise.

### GetStudyEventSetOk

`func (o *PlanFeatureSet) GetStudyEventSetOk() (*LimitedOrFullFeature, bool)`

GetStudyEventSetOk returns a tuple with the StudyEventSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyEventSet

`func (o *PlanFeatureSet) SetStudyEventSet(v LimitedOrFullFeature)`

SetStudyEventSet sets StudyEventSet field to given value.


### GetHasAudienceRecommendations

`func (o *PlanFeatureSet) GetHasAudienceRecommendations() bool`

GetHasAudienceRecommendations returns the HasAudienceRecommendations field if non-nil, zero value otherwise.

### GetHasAudienceRecommendationsOk

`func (o *PlanFeatureSet) GetHasAudienceRecommendationsOk() (*bool, bool)`

GetHasAudienceRecommendationsOk returns a tuple with the HasAudienceRecommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAudienceRecommendations

`func (o *PlanFeatureSet) SetHasAudienceRecommendations(v bool)`

SetHasAudienceRecommendations sets HasAudienceRecommendations field to given value.


### GetHasReportCustomization

`func (o *PlanFeatureSet) GetHasReportCustomization() bool`

GetHasReportCustomization returns the HasReportCustomization field if non-nil, zero value otherwise.

### GetHasReportCustomizationOk

`func (o *PlanFeatureSet) GetHasReportCustomizationOk() (*bool, bool)`

GetHasReportCustomizationOk returns a tuple with the HasReportCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReportCustomization

`func (o *PlanFeatureSet) SetHasReportCustomization(v bool)`

SetHasReportCustomization sets HasReportCustomization field to given value.


### GetHasReportSharingClients

`func (o *PlanFeatureSet) GetHasReportSharingClients() bool`

GetHasReportSharingClients returns the HasReportSharingClients field if non-nil, zero value otherwise.

### GetHasReportSharingClientsOk

`func (o *PlanFeatureSet) GetHasReportSharingClientsOk() (*bool, bool)`

GetHasReportSharingClientsOk returns a tuple with the HasReportSharingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReportSharingClients

`func (o *PlanFeatureSet) SetHasReportSharingClients(v bool)`

SetHasReportSharingClients sets HasReportSharingClients field to given value.


### GetHasReportSharingPublic

`func (o *PlanFeatureSet) GetHasReportSharingPublic() bool`

GetHasReportSharingPublic returns the HasReportSharingPublic field if non-nil, zero value otherwise.

### GetHasReportSharingPublicOk

`func (o *PlanFeatureSet) GetHasReportSharingPublicOk() (*bool, bool)`

GetHasReportSharingPublicOk returns a tuple with the HasReportSharingPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReportSharingPublic

`func (o *PlanFeatureSet) SetHasReportSharingPublic(v bool)`

SetHasReportSharingPublic sets HasReportSharingPublic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


