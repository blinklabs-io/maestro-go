# \ScriptsAPI

All URIs are relative to *https://mainnet.gomaestro-api.org/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScriptByHash**](ScriptsAPI.md#ScriptByHash) | **Get** /scripts/{script_hash} | Script of a script hash



## ScriptByHash

> Script ScriptByHash(ctx, scriptHash).Execute()

Script of a script hash



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
    scriptHash := "3a888d65f16790950a72daee1f63aa05add6d268434107cfa5b67712" // string | Hex encoded script hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScriptsAPI.ScriptByHash(context.Background(), scriptHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.ScriptByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScriptByHash`: Script
    fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.ScriptByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptHash** | **string** | Hex encoded script hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiScriptByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Script**](Script.md)

### Authorization

[api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

