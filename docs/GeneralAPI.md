# \GeneralAPI

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChainTip**](GeneralAPI.md#ChainTip) | **Get** /chain-tip | Chain tip details
[**EraHistory**](GeneralAPI.md#EraHistory) | **Get** /era-history | Era history
[**ProtocolParams**](GeneralAPI.md#ProtocolParams) | **Get** /protocol-params | Network protocol parameters
[**SystemStart**](GeneralAPI.md#SystemStart) | **Get** /system-start | System start time



## ChainTip

> ChainTip ChainTip(ctx).Execute()

Chain tip details



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
    resp, r, err := apiClient.GeneralAPI.ChainTip(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.ChainTip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChainTip`: ChainTip
    fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.ChainTip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChainTipRequest struct via the builder pattern


### Return type

[**ChainTip**](ChainTip.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EraHistory

> []EraSummary EraHistory(ctx).Execute()

Era history



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
    resp, r, err := apiClient.GeneralAPI.EraHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.EraHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EraHistory`: []EraSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.EraHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEraHistoryRequest struct via the builder pattern


### Return type

[**[]EraSummary**](EraSummary.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtocolParams

> ProtocolParameters ProtocolParams(ctx).Execute()

Network protocol parameters



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
    resp, r, err := apiClient.GeneralAPI.ProtocolParams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.ProtocolParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtocolParams`: ProtocolParameters
    fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.ProtocolParams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProtocolParamsRequest struct via the builder pattern


### Return type

[**ProtocolParameters**](ProtocolParameters.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemStart

> SystemStart SystemStart(ctx).Execute()

System start time



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
    resp, r, err := apiClient.GeneralAPI.SystemStart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.SystemStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemStart`: SystemStart
    fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.SystemStart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemStartRequest struct via the builder pattern


### Return type

[**SystemStart**](SystemStart.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

