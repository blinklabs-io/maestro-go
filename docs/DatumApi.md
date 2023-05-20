# \DatumApi

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupDatum**](DatumApi.md#LookupDatum) | **Get** /datum/{datum_hash} | Datum of a datum hash



## LookupDatum

> Datum LookupDatum(ctx, datumHash).Execute()

Datum of a datum hash



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
    datumHash := "432cb73420839fb517533c365d7ec125c457ea4ba5c0349be81be6796d52ef3b" // string | Hex encoded datum hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatumApi.LookupDatum(context.Background(), datumHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatumApi.LookupDatum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupDatum`: Datum
    fmt.Fprintf(os.Stdout, "Response from `DatumApi.LookupDatum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datumHash** | **string** | Hex encoded datum hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupDatumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Datum**](Datum.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

