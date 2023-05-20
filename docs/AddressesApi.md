# \AddressesApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressUtxos**](AddressesApi.md#AddressUtxos) | **Get** /addresses/{address}/utxos | UTxOs at an address
[**DecodeAddress**](AddressesApi.md#DecodeAddress) | **Get** /addresses/{address}/decode | Decoded receiving address information
[**TxCountByAddress**](AddressesApi.md#TxCountByAddress) | **Get** /addresses/{address}/transactions/count | Transaction count of an address
[**UtxoRefsAtAddress**](AddressesApi.md#UtxoRefsAtAddress) | **Get** /addresses/{address}/utxo_refs | UTxO references at an address
[**UtxosByAddresses**](AddressesApi.md#UtxosByAddresses) | **Post** /addresses/utxos | UTxOs at multiple addresses



## AddressUtxos

> []Utxo AddressUtxos(ctx, address).ResolveDatums(resolveDatums).WithCbor(withCbor).Count(count).Page(page).Execute()

UTxOs at an address



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
    address := "addr1qyl9ez2532trgydh5cny3ww72ygw8vugkvu98ln5x2zfgsxyf0yy00uzuxy7aldn75u7r8p99f4s0m93vgr4j7zz7pyqdg24wq" // string | Bech32 encoded address
    resolveDatums := true // bool | Try find and include the corresponding datums for datum hashes (optional) (default to false)
    withCbor := true // bool | Include the CBOR encodings of the transaction outputs in the response (optional) (default to false)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.AddressUtxos(context.Background(), address).ResolveDatums(resolveDatums).WithCbor(withCbor).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.AddressUtxos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressUtxos`: []Utxo
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.AddressUtxos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Bech32 encoded address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressUtxosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolveDatums** | **bool** | Try find and include the corresponding datums for datum hashes | [default to false]
 **withCbor** | **bool** | Include the CBOR encodings of the transaction outputs in the response | [default to false]
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]Utxo**](Utxo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecodeAddress

> AddressInfo DecodeAddress(ctx, address).Execute()

Decoded receiving address information



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
    address := "address_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.DecodeAddress(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.DecodeAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeAddress`: AddressInfo
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.DecodeAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressInfo**](AddressInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxCountByAddress

> AddressTxCount TxCountByAddress(ctx, address).Execute()

Transaction count of an address



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
    address := "addr1qyl9ez2532trgydh5cny3ww72ygw8vugkvu98ln5x2zfgsxyf0yy00uzuxy7aldn75u7r8p99f4s0m93vgr4j7zz7pyqdg24wq" // string | Bech32 encoded address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.TxCountByAddress(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.TxCountByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxCountByAddress`: AddressTxCount
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.TxCountByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Bech32 encoded address | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxCountByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressTxCount**](AddressTxCount.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UtxoRefsAtAddress

> []UtxoRef UtxoRefsAtAddress(ctx, address).Count(count).Page(page).Execute()

UTxO references at an address



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
    address := "addr1qyl9ez2532trgydh5cny3ww72ygw8vugkvu98ln5x2zfgsxyf0yy00uzuxy7aldn75u7r8p99f4s0m93vgr4j7zz7pyqdg24wq" // string | Bech32 encoded address
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.UtxoRefsAtAddress(context.Background(), address).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.UtxoRefsAtAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtxoRefsAtAddress`: []UtxoRef
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.UtxoRefsAtAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Bech32 encoded address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUtxoRefsAtAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]UtxoRef**](UtxoRef.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UtxosByAddresses

> []Utxo UtxosByAddresses(ctx).RequestBody(requestBody).ResolveDatums(resolveDatums).WithCbor(withCbor).Count(count).Page(page).Execute()

UTxOs at multiple addresses



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
    requestBody := []string{"Property_example"} // []string | List of addresses
    resolveDatums := true // bool | Try find and include the corresponding datums for datum hashes (optional) (default to false)
    withCbor := true // bool | Include the CBOR encodings of the transaction outputs in the response (optional) (default to false)
    count := int32(56) // int32 | The max number of results per pagination page (optional) (default to 100)
    page := int32(56) // int32 | Pagination page number to show results for (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.UtxosByAddresses(context.Background()).RequestBody(requestBody).ResolveDatums(resolveDatums).WithCbor(withCbor).Count(count).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.UtxosByAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtxosByAddresses`: []Utxo
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.UtxosByAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUtxosByAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | List of addresses | 
 **resolveDatums** | **bool** | Try find and include the corresponding datums for datum hashes | [default to false]
 **withCbor** | **bool** | Include the CBOR encodings of the transaction outputs in the response | [default to false]
 **count** | **int32** | The max number of results per pagination page | [default to 100]
 **page** | **int32** | Pagination page number to show results for | [default to 1]

### Return type

[**[]Utxo**](Utxo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

