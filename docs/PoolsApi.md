# \PoolsApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPools**](PoolsApi.md#ListPools) | **Get** /pools | Registered stake pools
[**PoolBlocks**](PoolsApi.md#PoolBlocks) | **Get** /pools/{pool_id}/blocks | Stake pool blocks
[**PoolDelegators**](PoolsApi.md#PoolDelegators) | **Get** /pools/{pool_id}/delegators | Stake pool delegators
[**PoolHistory**](PoolsApi.md#PoolHistory) | **Get** /pools/{pool_id}/history | Stake pool history
[**PoolInfo**](PoolsApi.md#PoolInfo) | **Get** /pools/{pool_id}/info | Stake pool details
[**PoolMetadata**](PoolsApi.md#PoolMetadata) | **Get** /pools/{pool_id}/metadata | Stake pool metadata
[**PoolRelays**](PoolsApi.md#PoolRelays) | **Get** /pools/{pool_id}/relays | Stake pool relays
[**PoolUpdates**](PoolsApi.md#PoolUpdates) | **Get** /pools/{pool_id}/updates | Stake pool updates



## ListPools

> []PoolListInfo ListPools(ctx).Count(count).Page(page).Execute()

Registered stake pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.ListPools(context.Background()).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.ListPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPools`: []PoolListInfo
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.ListPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]PoolListInfo**](PoolListInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolBlocks

> []PoolBlock PoolBlocks(ctx, poolId).EpochNo(epochNo).Count(count).Page(page).Order(order).Execute()

Stake pool blocks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID
    epochNo := int32(56) // int32 | Epoch number to fetch results for (optional)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)
    order := "order_example" // string | The order in which the results are sorted (by absolute slot) (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolBlocks(context.Background(), poolId).EpochNo(epochNo).Count(count).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolBlocks`: []PoolBlock
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **epochNo** | **int32** | Epoch number to fetch results for | 
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]
 **order** | **string** | The order in which the results are sorted (by absolute slot) | [default to &quot;asc&quot;]

### Return type

[**[]PoolBlock**](PoolBlock.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolDelegators

> []DelegatorInfo PoolDelegators(ctx, poolId).Count(count).Page(page).Execute()

Stake pool delegators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolDelegators(context.Background(), poolId).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolDelegators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolDelegators`: []DelegatorInfo
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolDelegators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolDelegatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]DelegatorInfo**](DelegatorInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolHistory

> []PoolHistory PoolHistory(ctx, poolId).EpochNo(epochNo).Count(count).Page(page).Order(order).Execute()

Stake pool history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID
    epochNo := int32(56) // int32 | Epoch number to fetch results for (optional)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)
    order := "order_example" // string | The order in which the results are sorted (by epoch number) (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolHistory(context.Background(), poolId).EpochNo(epochNo).Count(count).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolHistory`: []PoolHistory
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **epochNo** | **int32** | Epoch number to fetch results for | 
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]
 **order** | **string** | The order in which the results are sorted (by epoch number) | [default to &quot;asc&quot;]

### Return type

[**[]PoolHistory**](PoolHistory.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolInfo

> PoolInfo PoolInfo(ctx, poolId).Execute()

Stake pool details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolInfo(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolInfo`: PoolInfo
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolInfo**](PoolInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolMetadata

> PoolMetadata PoolMetadata(ctx, poolId).Execute()

Stake pool metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolMetadata(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolMetadata`: PoolMetadata
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolMetadata**](PoolMetadata.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolRelays

> []PoolRelay PoolRelays(ctx, poolId).Execute()

Stake pool relays



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolRelays(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolRelays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolRelays`: []PoolRelay
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolRelays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolRelaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PoolRelay**](PoolRelay.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolUpdates

> []PoolUpdate PoolUpdates(ctx, poolId).Execute()

Stake pool updates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolId := "pool125kh7e0y9lwya4sz5etmsk7hvga9jtfpuhw00vz9zvk6sh8xh5r" // string | Bech32 encoded pool ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoolsApi.PoolUpdates(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolsApi.PoolUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoolUpdates`: []PoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `PoolsApi.PoolUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Bech32 encoded pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PoolUpdate**](PoolUpdate.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

