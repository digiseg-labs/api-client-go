# \StudiesAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStudy**](StudiesAPI.md#CreateStudy) | **Post** /studies | Create study
[**DeleteStudyBannerImage**](StudiesAPI.md#DeleteStudyBannerImage) | **Delete** /studies/{study_id}/assets/banner_image | Delete study banner image logo
[**DeleteStudyById**](StudiesAPI.md#DeleteStudyById) | **Delete** /studies/{study_id} | Delete study
[**GetStudyBannerImage**](StudiesAPI.md#GetStudyBannerImage) | **Get** /studies/{study_id}/assets/banner_image | Get study banner image
[**GetStudyById**](StudiesAPI.md#GetStudyById) | **Get** /studies/{study_id} | Get study
[**ListStudies**](StudiesAPI.md#ListStudies) | **Get** /studies | List studies
[**QueryStudyAudienceStats**](StudiesAPI.md#QueryStudyAudienceStats) | **Get** /studies/{study_id}/stats/audiences | Audience statistics for study
[**QueryStudyCountryStats**](StudiesAPI.md#QueryStudyCountryStats) | **Get** /studies/{study_id}/stats/countries | Country statistics for study
[**QueryStudyFrequencyStats**](StudiesAPI.md#QueryStudyFrequencyStats) | **Get** /studies/{study_id}/stats/frequencies | Frequency statistics for study
[**QueryStudyTimingStats**](StudiesAPI.md#QueryStudyTimingStats) | **Get** /studies/{study_id}/stats/timing | Timing statistics for study
[**UpdateStudyById**](StudiesAPI.md#UpdateStudyById) | **Put** /studies/{study_id} | Update study
[**UploadStudyBannerImage**](StudiesAPI.md#UploadStudyBannerImage) | **Put** /studies/{study_id}/assets/banner_image | Upload study banner image



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


## ListStudies

> ListStudies200Response ListStudies(ctx).Sort(sort).FilterIsExample(filterIsExample).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()

List studies



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
	sort := "-created_at" // string | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  (optional) (default to "created_at")
	filterIsExample := true // bool | Optional parameter used to filter for example studies (optional) (default to false)
	filterLabel := "CTV" // string | Optional parameter used to filter by study label (optional)
	filterNameContains := "acme" // string | Optional parameter used to search for studies where the name contains a substring (case insensitive) (optional)
	filterAccountId := "4k3jKJ9D12kj0S4c" // string | Optional parameter used to query studies by specific account IDs (only available to super admins). The value `*` is synonymous for \"all accounts\".  (optional) (default to "The user's account ID")
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.ListStudies(context.Background()).Sort(sort).FilterIsExample(filterIsExample).FilterLabel(filterLabel).FilterNameContains(filterNameContains).FilterAccountId(filterAccountId).PageSize(pageSize).PageAfter(pageAfter).Execute()
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
 **filterIsExample** | **bool** | Optional parameter used to filter for example studies | [default to false]
 **filterLabel** | **string** | Optional parameter used to filter by study label | 
 **filterNameContains** | **string** | Optional parameter used to search for studies where the name contains a substring (case insensitive) | 
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

