# \QuotesApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuote**](QuotesApi.md#GetQuote) | **Get** /marketdata/{symbol}/quotes | Search for instrument and fundamental data
[**GetQuotes**](QuotesApi.md#GetQuotes) | **Get** /marketdata/quotes | Search for instrument and fundamental data



## GetQuote

> []map[string]interface{} GetQuote(ctx, symbol).ApiKey(apiKey).Execute()

Search for instrument and fundamental data



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
    symbol := "symbol_example" // string | Account ID
    apiKey := "apiKey_example" // string | API Key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.GetQuote(context.Background(), symbol).ApiKey(apiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.GetQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuote`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.GetQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **string** | API Key | 

### Return type

**[]map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> []map[string]interface{} GetQuotes(ctx).ApiKey(apiKey).Symbol(symbol).Execute()

Search for instrument and fundamental data



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
    apiKey := "apiKey_example" // string | API Key
    symbol := "symbol_example" // string | Symbol to search for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.GetQuotes(context.Background()).ApiKey(apiKey).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.GetQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuotes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API Key | 
 **symbol** | **string** | Symbol to search for | 

### Return type

**[]map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

