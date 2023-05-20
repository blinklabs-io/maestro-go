# \AccountsApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAddresses**](AccountsApi.md#AccountAddresses) | **Get** /accounts/{stake_address}/addresses | Stake account addresses
[**AccountAssets**](AccountsApi.md#AccountAssets) | **Get** /accounts/{stake_address}/assets | Stake account assets
[**AccountHistory**](AccountsApi.md#AccountHistory) | **Get** /accounts/{stake_address}/history | Stake account history
[**AccountInfo**](AccountsApi.md#AccountInfo) | **Get** /accounts/{stake_address} | Stake account information
[**AccountRewards**](AccountsApi.md#AccountRewards) | **Get** /accounts/{stake_address}/rewards | Stake account rewards
[**AccountUpdates**](AccountsApi.md#AccountUpdates) | **Get** /accounts/{stake_address}/updates | Stake account updates



## AccountAddresses

> []string AccountAddresses(ctx, stakeAddress).Count(count).Page(page).Execute()

Stake account addresses



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountAddresses(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAddresses`: []string
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountAddressesRequest struct via the builder pattern


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
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountAssets

> []Asset AccountAssets(ctx, stakeAddress).Count(count).Page(page).Execute()

Stake account assets



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountAssets(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAssets`: []Asset
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]Asset**](Asset.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountHistory

> []AccountHistory AccountHistory(ctx, stakeAddress).EpochNo(epochNo).Count(count).Page(page).Execute()

Stake account history



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address
    epochNo := int32(56) // int32 | Fetch result for a specific epoch (optional)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountHistory(context.Background(), stakeAddress).EpochNo(epochNo).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountHistory`: []AccountHistory
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **epochNo** | **int32** | Fetch result for a specific epoch | 
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]AccountHistory**](AccountHistory.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountInfo

> AccountInfo AccountInfo(ctx, stakeAddress).Execute()

Stake account information



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountInfo(context.Background(), stakeAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountInfo`: AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountInfo**](AccountInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountRewards

> []AccountReward AccountRewards(ctx, stakeAddress).Count(count).Page(page).Execute()

Stake account rewards



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountRewards(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountRewards`: []AccountReward
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountRewardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]AccountReward**](AccountReward.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountUpdates

> []AccountUpdate AccountUpdates(ctx, stakeAddress).Count(count).Page(page).Execute()

Stake account updates



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
    stakeAddress := "stake1u9w504hwl7sv6ewe72kxx45pzs37zq4qr4tvfqj3n2rqzvseuxpcf" // string | Bech32 encoded stake address
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountUpdates(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountUpdates`: []AccountUpdate
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakeAddress** | **string** | Bech32 encoded stake address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]AccountUpdate**](AccountUpdate.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

