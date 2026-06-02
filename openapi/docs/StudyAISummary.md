# StudyAISummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**ExecutiveSummary** | **string** |  | 
**KeyFindings** | **[]string** |  | 
**ImportantNumbers** | [**[]StudyAIImportantNumber**](StudyAIImportantNumber.md) |  | 
**Caveats** | **[]string** |  | 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewStudyAISummary

`func NewStudyAISummary(title string, executiveSummary string, keyFindings []string, importantNumbers []StudyAIImportantNumber, caveats []string, ) *StudyAISummary`

NewStudyAISummary instantiates a new StudyAISummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyAISummaryWithDefaults

`func NewStudyAISummaryWithDefaults() *StudyAISummary`

NewStudyAISummaryWithDefaults instantiates a new StudyAISummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StudyAISummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudyAISummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudyAISummary) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExecutiveSummary

`func (o *StudyAISummary) GetExecutiveSummary() string`

GetExecutiveSummary returns the ExecutiveSummary field if non-nil, zero value otherwise.

### GetExecutiveSummaryOk

`func (o *StudyAISummary) GetExecutiveSummaryOk() (*string, bool)`

GetExecutiveSummaryOk returns a tuple with the ExecutiveSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummary

`func (o *StudyAISummary) SetExecutiveSummary(v string)`

SetExecutiveSummary sets ExecutiveSummary field to given value.


### GetKeyFindings

`func (o *StudyAISummary) GetKeyFindings() []string`

GetKeyFindings returns the KeyFindings field if non-nil, zero value otherwise.

### GetKeyFindingsOk

`func (o *StudyAISummary) GetKeyFindingsOk() (*[]string, bool)`

GetKeyFindingsOk returns a tuple with the KeyFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFindings

`func (o *StudyAISummary) SetKeyFindings(v []string)`

SetKeyFindings sets KeyFindings field to given value.


### GetImportantNumbers

`func (o *StudyAISummary) GetImportantNumbers() []StudyAIImportantNumber`

GetImportantNumbers returns the ImportantNumbers field if non-nil, zero value otherwise.

### GetImportantNumbersOk

`func (o *StudyAISummary) GetImportantNumbersOk() (*[]StudyAIImportantNumber, bool)`

GetImportantNumbersOk returns a tuple with the ImportantNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportantNumbers

`func (o *StudyAISummary) SetImportantNumbers(v []StudyAIImportantNumber)`

SetImportantNumbers sets ImportantNumbers field to given value.


### GetCaveats

`func (o *StudyAISummary) GetCaveats() []string`

GetCaveats returns the Caveats field if non-nil, zero value otherwise.

### GetCaveatsOk

`func (o *StudyAISummary) GetCaveatsOk() (*[]string, bool)`

GetCaveatsOk returns a tuple with the Caveats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaveats

`func (o *StudyAISummary) SetCaveats(v []string)`

SetCaveats sets Caveats field to given value.


### GetGeneratedAt

`func (o *StudyAISummary) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *StudyAISummary) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *StudyAISummary) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *StudyAISummary) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


