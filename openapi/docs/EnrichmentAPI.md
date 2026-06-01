# \EnrichmentAPI

All URIs are relative to *https://api.digiseg.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnrichmentJob**](EnrichmentAPI.md#CreateEnrichmentJob) | **Post** /enrichment/jobs | Create enrichment job
[**DeleteEnrichmentJob**](EnrichmentAPI.md#DeleteEnrichmentJob) | **Delete** /enrichment/jobs/{id} | Delete enrichment job
[**DeleteEnrichmentJobResultData**](EnrichmentAPI.md#DeleteEnrichmentJobResultData) | **Delete** /enrichment/jobs/{id}/result | Delete enrichment job result data
[**DeleteEnrichmentJobSourceData**](EnrichmentAPI.md#DeleteEnrichmentJobSourceData) | **Delete** /enrichment/jobs/{id}/source | Delete enrichment job source data
[**DownloadResultFile**](EnrichmentAPI.md#DownloadResultFile) | **Get** /enrichment/jobs/{id}/result/file | Download enriched file
[**GetEnrichmentJobById**](EnrichmentAPI.md#GetEnrichmentJobById) | **Get** /enrichment/jobs/{id} | Get enrichment job by ID
[**ListEnrichmentJobs**](EnrichmentAPI.md#ListEnrichmentJobs) | **Get** /enrichment/jobs | List enrichment jobs
[**PreviewEnrichmentJobSourceData**](EnrichmentAPI.md#PreviewEnrichmentJobSourceData) | **Get** /enrichment/jobs/{id}/source/preview | Preview the source data
[**StartEnrichmentJob**](EnrichmentAPI.md#StartEnrichmentJob) | **Post** /enrichment/jobs/{id}/start | Start enrichment job
[**UpdateEnrichmentJob**](EnrichmentAPI.md#UpdateEnrichmentJob) | **Put** /enrichment/jobs/{id} | Update enrichment job
[**UploadEnrichmentJobSourceData**](EnrichmentAPI.md#UploadEnrichmentJobSourceData) | **Post** /enrichment/jobs/{id}/source | Upload source data to enrichment job



## CreateEnrichmentJob

> EnrichmentJobSingleResponse CreateEnrichmentJob(ctx).EnrichmentJobFull(enrichmentJobFull).Execute()

Create enrichment job

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
	enrichmentJobFull := *openapiclient.NewEnrichmentJobFull("Name_example") // EnrichmentJobFull | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.CreateEnrichmentJob(context.Background()).EnrichmentJobFull(enrichmentJobFull).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.CreateEnrichmentJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnrichmentJob`: EnrichmentJobSingleResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.CreateEnrichmentJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrichmentJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrichmentJobFull** | [**EnrichmentJobFull**](EnrichmentJobFull.md) |  | 

### Return type

[**EnrichmentJobSingleResponse**](EnrichmentJobSingleResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnrichmentJob

> DeleteEnrichmentJob(ctx, id).Execute()

Delete enrichment job

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrichmentAPI.DeleteEnrichmentJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.DeleteEnrichmentJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnrichmentJobRequest struct via the builder pattern


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


## DeleteEnrichmentJobResultData

> DeleteEnrichmentJobResultData(ctx, id).Execute()

Delete enrichment job result data

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrichmentAPI.DeleteEnrichmentJobResultData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.DeleteEnrichmentJobResultData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnrichmentJobResultDataRequest struct via the builder pattern


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


## DeleteEnrichmentJobSourceData

> DeleteEnrichmentJobSourceData(ctx, id).Execute()

Delete enrichment job source data

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrichmentAPI.DeleteEnrichmentJobSourceData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.DeleteEnrichmentJobSourceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnrichmentJobSourceDataRequest struct via the builder pattern


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


## DownloadResultFile

> *os.File DownloadResultFile(ctx, id).Execute()

Download enriched file

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.DownloadResultFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.DownloadResultFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadResultFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.DownloadResultFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadResultFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrichmentJobById

> EnrichmentJobSingleResponse GetEnrichmentJobById(ctx, id).Execute()

Get enrichment job by ID

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.GetEnrichmentJobById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.GetEnrichmentJobById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrichmentJobById`: EnrichmentJobSingleResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.GetEnrichmentJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrichmentJobByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrichmentJobSingleResponse**](EnrichmentJobSingleResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnrichmentJobs

> EnrichmentJobListResponse ListEnrichmentJobs(ctx).PageSize(pageSize).PageAfter(pageAfter).Sort(sort).Execute()

List enrichment jobs

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
	pageSize := int32(56) // int32 | The desired page size (optional) (default to 100)
	pageAfter := "pageAfter_example" // string | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the `meta.page.last_cursor` field.  (optional)
	sort := openapiclient.EnrichmentJobSortOption("created_at") // EnrichmentJobSortOption | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.ListEnrichmentJobs(context.Background()).PageSize(pageSize).PageAfter(pageAfter).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.ListEnrichmentJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnrichmentJobs`: EnrichmentJobListResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.ListEnrichmentJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnrichmentJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The desired page size | [default to 100]
 **pageAfter** | **string** | Optional pagination parameter, indicating the previous cursor value to paginate beyond. The value to provide here is opaque, but can be found in previous requests in the &#x60;meta.page.last_cursor&#x60; field.  | 
 **sort** | [**EnrichmentJobSortOption**](EnrichmentJobSortOption.md) | Defines the field to sort the result items by. Ascending order is applied by default, but the minus character can be used to indicate descending order instead.  | [default to &quot;created_at&quot;]

### Return type

[**EnrichmentJobListResponse**](EnrichmentJobListResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewEnrichmentJobSourceData

> EnrichmentJobSourceDataPreviewResponse PreviewEnrichmentJobSourceData(ctx, id).Execute()

Preview the source data



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.PreviewEnrichmentJobSourceData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.PreviewEnrichmentJobSourceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewEnrichmentJobSourceData`: EnrichmentJobSourceDataPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.PreviewEnrichmentJobSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrichmentJobSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrichmentJobSourceDataPreviewResponse**](EnrichmentJobSourceDataPreviewResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEnrichmentJob

> StartEnrichmentJob(ctx, id).Execute()

Start enrichment job

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrichmentAPI.StartEnrichmentJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.StartEnrichmentJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEnrichmentJobRequest struct via the builder pattern


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


## UpdateEnrichmentJob

> EnrichmentJobSingleResponse UpdateEnrichmentJob(ctx, id).EnrichmentJobUpdate(enrichmentJobUpdate).Execute()

Update enrichment job

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
	id := "id_example" // string | 
	enrichmentJobUpdate := *openapiclient.NewEnrichmentJobUpdate("Name_example") // EnrichmentJobUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrichmentAPI.UpdateEnrichmentJob(context.Background(), id).EnrichmentJobUpdate(enrichmentJobUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.UpdateEnrichmentJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnrichmentJob`: EnrichmentJobSingleResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrichmentAPI.UpdateEnrichmentJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnrichmentJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrichmentJobUpdate** | [**EnrichmentJobUpdate**](EnrichmentJobUpdate.md) |  | 

### Return type

[**EnrichmentJobSingleResponse**](EnrichmentJobSingleResponse.md)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadEnrichmentJobSourceData

> UploadEnrichmentJobSourceData(ctx, id).File(file).Execute()

Upload source data to enrichment job

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
	id := "id_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrichmentAPI.UploadEnrichmentJobSourceData(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentAPI.UploadEnrichmentJobSourceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadEnrichmentJobSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [bearerAuth](../README.md#bearerAuth), [apiKeyHeaderAuth](../README.md#apiKeyHeaderAuth), [apiKeyQueryParamAuth](../README.md#apiKeyQueryParamAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

