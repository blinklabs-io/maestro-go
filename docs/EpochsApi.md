# \EpochsApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrentEpoch**](EpochsApi.md#CurrentEpoch) | **Get** /epochs/current | Current epoch of the network
[**EpochInfo**](EpochsApi.md#EpochInfo) | **Get** /epochs/{epoch_no}/info | Epoch details



## CurrentEpoch

> CurrentEpochInfo CurrentEpoch(ctx).Execute()

Current epoch of the network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/blinklabs-io/maestro-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpochsApi.CurrentEpoch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpochsApi.CurrentEpoch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CurrentEpoch`: CurrentEpochInfo
    fmt.Fprintf(os.Stdout, "Response from `EpochsApi.CurrentEpoch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCurrentEpochRequest struct via the builder pattern


### Return type

[**CurrentEpochInfo**](CurrentEpochInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EpochInfo

> EpochInfo EpochInfo(ctx, epochNo).Execute()

Epoch details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/blinklabs-io/maestro-go"
)

func main() {
    epochNo := int32(56) // int32 | Epoch height (number)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpochsApi.EpochInfo(context.Background(), epochNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpochsApi.EpochInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EpochInfo`: EpochInfo
    fmt.Fprintf(os.Stdout, "Response from `EpochsApi.EpochInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epochNo** | **int32** | Epoch height (number) | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpochInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpochInfo**](EpochInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

