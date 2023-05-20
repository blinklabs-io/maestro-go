# \AssetsApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetAddresses**](AssetsApi.md#AssetAddresses) | **Get** /assets/{asset}/addresses | Native asset addresses
[**AssetInfo**](AssetsApi.md#AssetInfo) | **Get** /assets/{asset} | Native asset information
[**AssetTxs**](AssetsApi.md#AssetTxs) | **Get** /assets/{asset}/txs | Native asset transaction
[**AssetUpdates**](AssetsApi.md#AssetUpdates) | **Get** /assets/{asset}/updates | Native asset updates
[**PolicyAddresses**](AssetsApi.md#PolicyAddresses) | **Get** /assets/policy/{policy}/addresses | Minting policy addresses
[**PolicyInfo**](AssetsApi.md#PolicyInfo) | **Get** /assets/policy/{policy} | Minting policy information
[**PolicyTxs**](AssetsApi.md#PolicyTxs) | **Get** /assets/policy/{policy}/txs | Minting policy transactions
[**PolicyUtxos**](AssetsApi.md#PolicyUtxos) | **Get** /assets/policy/{policy}/utxos | Minting policy UTxOs



## AssetAddresses

> []string AssetAddresses(ctx, asset).Count(count).Page(page).Execute()

Native asset addresses



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
    asset := "29d222ce763455e3d7a09a665ce554f00ac89d2e99a1a83d267170c64d494e" // string | Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetAddresses(context.Background(), asset).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetAddresses`: []string
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | **string** | Native asset (concatenation of hex encoded policy ID and hex encoded asset name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

**[]string**

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetInfo

> AssetInfo AssetInfo(ctx, asset).Execute()

Native asset information



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
    asset := "29d222ce763455e3d7a09a665ce554f00ac89d2e99a1a83d267170c64d494e" // string | Native asset (concatenation of hex encoded policy ID and hex encoded asset name)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetInfo(context.Background(), asset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetInfo`: AssetInfo
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | **string** | Native asset (concatenation of hex encoded policy ID and hex encoded asset name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetInfo**](AssetInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetTxs

> []AssetTx AssetTxs(ctx, asset).FromHeight(fromHeight).Count(count).Page(page).Order(order).Execute()

Native asset transaction



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
    asset := "29d222ce763455e3d7a09a665ce554f00ac89d2e99a1a83d267170c64d494e" // string | Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
    fromHeight := int32(56) // int32 | Return only transactions on or after a specific block height (optional)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)
    order := "order_example" // string | The order in which the results are sorted (by block height) (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetTxs(context.Background(), asset).FromHeight(fromHeight).Count(count).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetTxs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetTxs`: []AssetTx
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetTxs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | **string** | Native asset (concatenation of hex encoded policy ID and hex encoded asset name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetTxsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromHeight** | **int32** | Return only transactions on or after a specific block height | 
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]
 **order** | **string** | The order in which the results are sorted (by block height) | [default to &quot;asc&quot;]

### Return type

[**[]AssetTx**](AssetTx.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetUpdates

> []MintingTx AssetUpdates(ctx, asset).Count(count).Page(page).Order(order).Execute()

Native asset updates



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
    asset := "29d222ce763455e3d7a09a665ce554f00ac89d2e99a1a83d267170c64d494e" // string | Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)
    order := "order_example" // string | The order in which the results are sorted (by block height) (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetUpdates(context.Background(), asset).Count(count).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetUpdates`: []MintingTx
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | **string** | Native asset (concatenation of hex encoded policy ID and hex encoded asset name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]
 **order** | **string** | The order in which the results are sorted (by block height) | [default to &quot;asc&quot;]

### Return type

[**[]MintingTx**](MintingTx.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyAddresses

> []string PolicyAddresses(ctx, policy).Count(count).Page(page).Execute()

Minting policy addresses



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
    policy := "policy_example" // string | Hex encoded Policy ID
    count := int32(56) // int32 | The max number of results per page (optional) (default to 100)
    page := int32(56) // int32 | The page number for the results (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.PolicyAddresses(context.Background(), policy).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.PolicyAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyAddresses`: []string
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.PolicyAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Hex encoded Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per page | [default to 100]
 **page** | **int32** | The page number for the results | [default to 1]

### Return type

**[]string**

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyInfo

> []AssetInfo PolicyInfo(ctx, policy).Count(count).Page(page).Execute()

Minting policy information



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
    policy := "policy_example" // string | Hex encoded policy ID
    count := int32(56) // int32 | The max number of results per page (optional) (default to 100)
    page := int32(56) // int32 | The page number for the results (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.PolicyInfo(context.Background(), policy).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.PolicyInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyInfo`: []AssetInfo
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.PolicyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Hex encoded policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per page | [default to 100]
 **page** | **int32** | The page number for the results | [default to 1]

### Return type

[**[]AssetInfo**](AssetInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyTxs

> []AssetTx PolicyTxs(ctx, policy).FromHeight(fromHeight).Count(count).Page(page).Order(order).Execute()

Minting policy transactions



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
    policy := "policy_example" // string | Hex encoded policy ID
    fromHeight := int64(789) // int64 | Return only transactions after supplied block height (optional)
    count := int32(56) // int32 | The max number of results per page (optional) (default to 100)
    page := int32(56) // int32 | The page number for the results (optional) (default to 1)
    order := "order_example" // string | The order in which the results are sorted (by block height) (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.PolicyTxs(context.Background(), policy).FromHeight(fromHeight).Count(count).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.PolicyTxs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyTxs`: []AssetTx
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.PolicyTxs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Hex encoded policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyTxsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromHeight** | **int64** | Return only transactions after supplied block height | 
 **count** | **int32** | The max number of results per page | [default to 100]
 **page** | **int32** | The page number for the results | [default to 1]
 **order** | **string** | The order in which the results are sorted (by block height) | [default to &quot;asc&quot;]

### Return type

[**[]AssetTx**](AssetTx.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyUtxos

> []PolicyUtxo PolicyUtxos(ctx, policy).Count(count).Page(page).Execute()

Minting policy UTxOs



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
    policy := "policy_example" // string | Hex encoded policy ID
    count := int32(56) // int32 | The max number of results per page (optional) (default to 100)
    page := int32(56) // int32 | The page number for the results (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.PolicyUtxos(context.Background(), policy).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.PolicyUtxos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyUtxos`: []PolicyUtxo
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.PolicyUtxos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Hex encoded policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyUtxosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per page | [default to 100]
 **page** | **int32** | The page number for the results | [default to 1]

### Return type

[**[]PolicyUtxo**](PolicyUtxo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

