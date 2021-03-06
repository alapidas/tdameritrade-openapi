# \MoversApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovers**](MoversApi.md#GetMovers) | **Get** /marketdata/{index}/movers | Retrieve mover information by index symbol, direction type and change



## GetMovers

> Mover GetMovers(ctx, index).Apikey(apikey).Direction(direction).Change(change).Execute()

Retrieve mover information by index symbol, direction type and change



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apikey := "apikey_example" // string | API Key
    index := "index_example" // string | The index symbol to get movers from.
    direction := "'up' or 'down'" // string | To return movers with the specified directions of up or down.
    change := "'percent' or 'value'" // string | To return movers with the specified change types of percent or value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MoversApi.GetMovers(context.Background(), index).Apikey(apikey).Direction(direction).Change(change).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoversApi.GetMovers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMovers`: Mover
    fmt.Fprintf(os.Stdout, "Response from `MoversApi.GetMovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The index symbol to get movers from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | API Key | 

 **direction** | **string** | To return movers with the specified directions of up or down. | 
 **change** | **string** | To return movers with the specified change types of percent or value. | 

### Return type

[**Mover**](Mover.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

