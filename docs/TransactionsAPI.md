# \TransactionsAPI

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressByTxo**](TransactionsAPI.md#AddressByTxo) | **Get** /transactions/{tx_hash}/outputs/{index}/address | Address of a transaction output reference
[**IogTxSubmit**](TransactionsAPI.md#IogTxSubmit) | **Post** /submit/tx | Submit transaction
[**MaestroTxSubmit**](TransactionsAPI.md#MaestroTxSubmit) | **Post** /transactions | Submit transaction
[**TxCborByTxHash**](TransactionsAPI.md#TxCborByTxHash) | **Get** /transactions/{tx_hash}/cbor | CBOR bytes of a transaction
[**TxInfo**](TransactionsAPI.md#TxInfo) | **Get** /transactions/{tx_hash} | Transaction information
[**UtxoByTxoRef**](TransactionsAPI.md#UtxoByTxoRef) | **Get** /transactions/{tx_hash}/outputs/{index}/utxo | Transaction output of an output reference



## AddressByTxo

> UtxoAddress AddressByTxo(ctx, txHash, index).Execute()

Address of a transaction output reference



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
    txHash := "9907c1bcab96889368d975ec1964e2fedfef22ce4a0e367bf9cb621b9f0dcb4a" // string | Hex encoded transaction hash
    index := int32(56) // int32 | Transaction output index

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.AddressByTxo(context.Background(), txHash, index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.AddressByTxo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressByTxo`: UtxoAddress
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.AddressByTxo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txHash** | **string** | Hex encoded transaction hash | 
**index** | **int32** | Transaction output index | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressByTxoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UtxoAddress**](UtxoAddress.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IogTxSubmit

> string IogTxSubmit(ctx).Body(body).Execute()

Submit transaction



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
    body := "body_example" // string | CBOR encoded transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.IogTxSubmit(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.IogTxSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IogTxSubmit`: string
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.IogTxSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIogTxSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | CBOR encoded transaction | 

### Return type

**string**

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: application/cbor
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaestroTxSubmit

> string MaestroTxSubmit(ctx).Body(body).Execute()

Submit transaction



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
    body := "body_example" // string | CBOR encoded transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.MaestroTxSubmit(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.MaestroTxSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaestroTxSubmit`: string
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.MaestroTxSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaestroTxSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | CBOR encoded transaction | 

### Return type

**string**

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: application/cbor
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxCborByTxHash

> TxCbor TxCborByTxHash(ctx, txHash).Execute()

CBOR bytes of a transaction



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
    txHash := "9907c1bcab96889368d975ec1964e2fedfef22ce4a0e367bf9cb621b9f0dcb4a" // string | Hex encoded transaction hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.TxCborByTxHash(context.Background(), txHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TxCborByTxHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxCborByTxHash`: TxCbor
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TxCborByTxHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txHash** | **string** | Hex encoded transaction hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxCborByTxHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TxCbor**](TxCbor.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxInfo

> TransactionInfo TxInfo(ctx, txHash).Execute()

Transaction information



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
    txHash := "9907c1bcab96889368d975ec1964e2fedfef22ce4a0e367bf9cb621b9f0dcb4a" // string | Hex encoded transaction hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.TxInfo(context.Background(), txHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TxInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxInfo`: TransactionInfo
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TxInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txHash** | **string** | Hex encoded transaction hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionInfo**](TransactionInfo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UtxoByTxoRef

> Utxo UtxoByTxoRef(ctx, txHash, index).ResolveDatums(resolveDatums).WithCbor(withCbor).Execute()

Transaction output of an output reference



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
    txHash := "9907c1bcab96889368d975ec1964e2fedfef22ce4a0e367bf9cb621b9f0dcb4a" // string | Hex encoded transaction hash
    index := int32(56) // int32 | Transaction output index
    resolveDatums := true // bool | Try find and include the corresponding datums for datum hashes (optional) (default to false)
    withCbor := true // bool | Include the CBOR encodings of the transaction outputs in the response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.UtxoByTxoRef(context.Background(), txHash, index).ResolveDatums(resolveDatums).WithCbor(withCbor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.UtxoByTxoRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtxoByTxoRef`: Utxo
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.UtxoByTxoRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txHash** | **string** | Hex encoded transaction hash | 
**index** | **int32** | Transaction output index | 

### Other Parameters

Other parameters are passed through a pointer to a apiUtxoByTxoRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveDatums** | **bool** | Try find and include the corresponding datums for datum hashes | [default to false]
 **withCbor** | **bool** | Include the CBOR encodings of the transaction outputs in the response | [default to false]

### Return type

[**Utxo**](Utxo.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

