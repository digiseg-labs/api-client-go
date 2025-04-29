# \StudiesAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedReport**](StudiesAPI.md#CreateSharedReport) | **Post** /studies/{study_id}/shared_reports | Create shared report
[**CreateStudy**](StudiesAPI.md#CreateStudy) | **Post** /studies | Create study
[**CreateStudyEvent**](StudiesAPI.md#CreateStudyEvent) | **Post** /studies/{study_id}/events | Create study event
[**DeleteSharedReportById**](StudiesAPI.md#DeleteSharedReportById) | **Delete** /studies/{study_id}/shared_reports/{report_id} | Delete shared report
[**DeleteStudyBannerImage**](StudiesAPI.md#DeleteStudyBannerImage) | **Delete** /studies/{study_id}/assets/banner_image | Delete study banner image logo
[**DeleteStudyById**](StudiesAPI.md#DeleteStudyById) | **Delete** /studies/{study_id} | Delete study
[**GetSharedReportById**](StudiesAPI.md#GetSharedReportById) | **Get** /studies/{study_id}/shared_reports/{report_id} | Get shared report
[**GetStudyBannerImage**](StudiesAPI.md#GetStudyBannerImage) | **Get** /studies/{study_id}/assets/banner_image | Get study banner image
[**GetStudyById**](StudiesAPI.md#GetStudyById) | **Get** /studies/{study_id} | Get study
[**ListSharedReportsByStudyId**](StudiesAPI.md#ListSharedReportsByStudyId) | **Get** /studies/{study_id}/shared_reports | Get shared reports for a study
[**ListStudies**](StudiesAPI.md#ListStudies) | **Get** /studies | List studies
[**QueryStudyAudienceStats**](StudiesAPI.md#QueryStudyAudienceStats) | **Get** /studies/{study_id}/stats/audiences | Audience statistics for study
[**QueryStudyCountryStats**](StudiesAPI.md#QueryStudyCountryStats) | **Get** /studies/{study_id}/stats/countries | Country statistics for study
[**QueryStudyDeviceStats**](StudiesAPI.md#QueryStudyDeviceStats) | **Get** /studies/{study_id}/stats/devices | Device statistics for study
[**QueryStudyFrequencyStats**](StudiesAPI.md#QueryStudyFrequencyStats) | **Get** /studies/{study_id}/stats/frequencies | Frequency statistics for study
[**QueryStudyTimelineStats**](StudiesAPI.md#QueryStudyTimelineStats) | **Get** /studies/{study_id}/stats/timeline | Timeline statistics for study
[**QueryStudyTimingStats**](StudiesAPI.md#QueryStudyTimingStats) | **Get** /studies/{study_id}/stats/timing | Timing statistics for study
[**UpdateStudyById**](StudiesAPI.md#UpdateStudyById) | **Put** /studies/{study_id} | Update study
[**UploadStudyBannerImage**](StudiesAPI.md#UploadStudyBannerImage) | **Put** /studies/{study_id}/assets/banner_image | Upload study banner image



## CreateSharedReport

> CreateSharedReport201Response CreateSharedReport(ctx, studyId).SharedReportCreation(sharedReportCreation).Execute()

Create shared report

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	sharedReportCreation := *openapiclient.NewSharedReportCreation(openapiclient.SharedReportType("audience_evaluation")) // SharedReportCreation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.CreateSharedReport(context.Background(), studyId).SharedReportCreation(sharedReportCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.CreateSharedReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSharedReport`: CreateSharedReport201Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.CreateSharedReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedReportCreation** | [**SharedReportCreation**](SharedReportCreation.md) |  | 

### Return type

[**CreateSharedReport201Response**](CreateSharedReport201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStudy

> CreateStudy201Response CreateStudy(ctx).StudyCreation(studyCreation).Execute()

Create study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyCreation := *openapiclient.NewStudyCreation("Acme Summer 2024 brand awareness", openapiclient.MeasurementEventSet("impressions_only")) // StudyCreation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.CreateStudy(context.Background()).StudyCreation(studyCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.CreateStudy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStudy`: CreateStudy201Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.CreateStudy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStudyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studyCreation** | [**StudyCreation**](StudyCreation.md) |  | 

### Return type

[**CreateStudy201Response**](CreateStudy201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStudyEvent

> CreateStudyEvent(ctx, studyId).StudyEventCreation(studyEventCreation).Execute()

Create study event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	studyEventCreation := *openapiclient.NewStudyEventCreation() // StudyEventCreation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.CreateStudyEvent(context.Background(), studyId).StudyEventCreation(studyEventCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.CreateStudyEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStudyEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studyEventCreation** | [**StudyEventCreation**](StudyEventCreation.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharedReportById

> DeleteSharedReportById(ctx, studyId, reportId).Execute()

Delete shared report

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.DeleteSharedReportById(context.Background(), studyId, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.DeleteSharedReportById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedReportByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStudyBannerImage

> DeleteStudyBannerImage(ctx, studyId).Execute()

Delete study banner image logo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.DeleteStudyBannerImage(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.DeleteStudyBannerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudyBannerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStudyById

> DeleteStudyById(ctx, studyId).Execute()

Delete study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.DeleteStudyById(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.DeleteStudyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedReportById

> CreateSharedReport201Response GetSharedReportById(ctx, studyId, reportId).Execute()

Get shared report

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.GetSharedReportById(context.Background(), studyId, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.GetSharedReportById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedReportById`: CreateSharedReport201Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.GetSharedReportById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedReportByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateSharedReport201Response**](CreateSharedReport201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudyBannerImage

> GetStudyBannerImage(ctx, studyId).Execute()

Get study banner image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.GetStudyBannerImage(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.GetStudyBannerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudyBannerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudyById

> CreateStudy201Response GetStudyById(ctx, studyId).Execute()

Get study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.GetStudyById(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.GetStudyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudyById`: CreateStudy201Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.GetStudyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateStudy201Response**](CreateStudy201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedReportsByStudyId

> ListSharedReportsByStudyId200Response ListSharedReportsByStudyId(ctx, studyId).Execute()

Get shared reports for a study

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.ListSharedReportsByStudyId(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ListSharedReportsByStudyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedReportsByStudyId`: ListSharedReportsByStudyId200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.ListSharedReportsByStudyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedReportsByStudyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSharedReportsByStudyId200Response**](ListSharedReportsByStudyId200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStudies

> ListStudies200Response ListStudies(ctx).Sort(sort).FilterLifeCycleStage(filterLifeCycleStage).FilterIsExample(filterIsExample).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterStartDateAfter(filterStartDateAfter).FilterStartDateBefore(filterStartDateBefore).FilterEndDate(filterEndDate).FilterCreatedAtAfter(filterCreatedAtAfter).FilterCreatedAtBefore(filterCreatedAtBefore).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

List studies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	sort := "-created_at" // string | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  (optional) (default to "created_at")
	filterLifeCycleStage := openapiclient.StudyLifecycleStage("pending") // StudyLifecycleStage | Optional parameter used to filter studies by their life cycle stage (optional)
	filterIsExample := true // bool | Optional parameter used to filter for example studies (optional) (default to false)
	filterLabel := "CTV" // string | Optional parameter used to filter by study label (optional)
	filterNameContains := "acme" // string | Optional parameter used to search for studies where the name contains a substring (case insensitive) (optional)
	filterStartDateAfter := time.Now() // time.Time | Optional parameter used to search for studies that have started after a specific date (optional)
	filterStartDateBefore := time.Now() // time.Time | Optional parameter used to search for studies that have start before a specific date (optional)
	filterEndDate := time.Now() // time.Time | Optional parameter used to search for studies that have a specific end date (optional)
	filterCreatedAtAfter := time.Now() // time.Time | Optional parameter used to search for studies that have been created after a specific date (optional)
	filterCreatedAtBefore := time.Now() // time.Time | Optional parameter used to search for studies that have been created before a specific date (optional)
	filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query studies by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.ListStudies(context.Background()).Sort(sort).FilterLifeCycleStage(filterLifeCycleStage).FilterIsExample(filterIsExample).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterStartDateAfter(filterStartDateAfter).FilterStartDateBefore(filterStartDateBefore).FilterEndDate(filterEndDate).FilterCreatedAtAfter(filterCreatedAtAfter).FilterCreatedAtBefore(filterCreatedAtBefore).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ListStudies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStudies`: ListStudies200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.ListStudies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStudiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  | [default to &quot;created_at&quot;]
 **filterLifeCycleStage** | [**StudyLifecycleStage**](StudyLifecycleStage.md) | Optional parameter used to filter studies by their life cycle stage | 
 **filterIsExample** | **bool** | Optional parameter used to filter for example studies | [default to false]
 **filterLabel** | **string** | Optional parameter used to filter by study label | 
 **filterNameContains** | **string** | Optional parameter used to search for studies where the name contains a substring (case insensitive) | 
 **filterStartDateAfter** | **time.Time** | Optional parameter used to search for studies that have started after a specific date | 
 **filterStartDateBefore** | **time.Time** | Optional parameter used to search for studies that have start before a specific date | 
 **filterEndDate** | **time.Time** | Optional parameter used to search for studies that have a specific end date | 
 **filterCreatedAtAfter** | **time.Time** | Optional parameter used to search for studies that have been created after a specific date | 
 **filterCreatedAtBefore** | **time.Time** | Optional parameter used to search for studies that have been created before a specific date | 
 **filterAccountId** | **string** | Optional parameter used to query studies by specific account IDs (only available to super admins). The value &#x60;*&#x60; is synonymous for \&quot;all accounts\&quot;.  | [default to &quot;The user&#39;s account ID&quot;]
 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 

### Return type

[**ListStudies200Response**](ListStudies200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyAudienceStats

> QueryStudyAudienceStats200Response QueryStudyAudienceStats(ctx, studyId).Execute()

Audience statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyAudienceStats(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyAudienceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyAudienceStats`: QueryStudyAudienceStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyAudienceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyAudienceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryStudyAudienceStats200Response**](QueryStudyAudienceStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyCountryStats

> QueryStudyCountryStats200Response QueryStudyCountryStats(ctx, studyId).Execute()

Country statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyCountryStats(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyCountryStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyCountryStats`: QueryStudyCountryStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyCountryStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyCountryStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryStudyCountryStats200Response**](QueryStudyCountryStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyDeviceStats

> QueryStudyDeviceStats200Response QueryStudyDeviceStats(ctx, studyId).Execute()

Device statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyDeviceStats(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyDeviceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyDeviceStats`: QueryStudyDeviceStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyDeviceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyDeviceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryStudyDeviceStats200Response**](QueryStudyDeviceStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyFrequencyStats

> QueryStudyFrequencyStats200Response QueryStudyFrequencyStats(ctx, studyId).Execute()

Frequency statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyFrequencyStats(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyFrequencyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyFrequencyStats`: QueryStudyFrequencyStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyFrequencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyFrequencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryStudyFrequencyStats200Response**](QueryStudyFrequencyStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyTimelineStats

> QueryStudyTimelineStats200Response QueryStudyTimelineStats(ctx, studyId).FromDate(fromDate).ToDate(toDate).Execute()

Timeline statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	fromDate := time.Now() // string |  (optional)
	toDate := time.Now() // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyTimelineStats(context.Background(), studyId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyTimelineStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyTimelineStats`: QueryStudyTimelineStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyTimelineStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyTimelineStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**QueryStudyTimelineStats200Response**](QueryStudyTimelineStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryStudyTimingStats

> QueryStudyTimingStats200Response QueryStudyTimingStats(ctx, studyId).Execute()

Timing statistics for study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.QueryStudyTimingStats(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.QueryStudyTimingStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryStudyTimingStats`: QueryStudyTimingStats200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.QueryStudyTimingStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryStudyTimingStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryStudyTimingStats200Response**](QueryStudyTimingStats200Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStudyById

> CreateStudy201Response UpdateStudyById(ctx, studyId).StudyMutation(studyMutation).Execute()

Update study



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	studyMutation := *openapiclient.NewStudyMutation() // StudyMutation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.UpdateStudyById(context.Background(), studyId).StudyMutation(studyMutation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.UpdateStudyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStudyById`: CreateStudy201Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.UpdateStudyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStudyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studyMutation** | [**StudyMutation**](StudyMutation.md) |  | 

### Return type

[**CreateStudy201Response**](CreateStudy201Response.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadStudyBannerImage

> UploadStudyBannerImage(ctx, studyId).Body(body).Execute()

Upload study banner image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/digiseg-labs/api-client-go"
)

func main() {
	studyId := "studyId_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.UploadStudyBannerImage(context.Background(), studyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.UploadStudyBannerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadStudyBannerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: image/gif, image/png, image/jpeg, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

