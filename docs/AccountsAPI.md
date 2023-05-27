# \AccountsAPI

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAddresses**](AccountsAPI.md#AccountAddresses) | **Get** /accounts/{stake_address}/addresses | Stake account addresses
[**AccountAssets**](AccountsAPI.md#AccountAssets) | **Get** /accounts/{stake_address}/assets | Stake account assets
[**AccountHistory**](AccountsAPI.md#AccountHistory) | **Get** /accounts/{stake_address}/history | Stake account history
[**AccountInfo**](AccountsAPI.md#AccountInfo) | **Get** /accounts/{stake_address} | Stake account information
[**AccountRewards**](AccountsAPI.md#AccountRewards) | **Get** /accounts/{stake_address}/rewards | Stake account rewards
[**AccountUpdates**](AccountsAPI.md#AccountUpdates) | **Get** /accounts/{stake_address}/updates | Stake account updates



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
    resp, r, err := apiClient.AccountsAPI.AccountAddresses(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAddresses`: []string
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountAddresses`: %v\n", resp)
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
    resp, r, err := apiClient.AccountsAPI.AccountAssets(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAssets`: []Asset
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountAssets`: %v\n", resp)
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
    resp, r, err := apiClient.AccountsAPI.AccountHistory(context.Background(), stakeAddress).EpochNo(epochNo).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountHistory`: []AccountHistory
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountHistory`: %v\n", resp)
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
    resp, r, err := apiClient.AccountsAPI.AccountInfo(context.Background(), stakeAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountInfo`: AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountInfo`: %v\n", resp)
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
    resp, r, err := apiClient.AccountsAPI.AccountRewards(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountRewards`: []AccountReward
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountRewards`: %v\n", resp)
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
    resp, r, err := apiClient.AccountsAPI.AccountUpdates(context.Background(), stakeAddress).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountUpdates`: []AccountUpdate
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountUpdates`: %v\n", resp)
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

