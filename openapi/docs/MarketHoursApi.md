# \MarketHoursApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHours**](MarketHoursApi.md#GetHours) | **Get** /marketdata/hours | Operating hours of markets
[**GetMarketHours**](MarketHoursApi.md#GetMarketHours) | **Get** /marketdata/{market}/hours | Operating hours of markets



## GetHours

> Hours GetHours(ctx).Apikey(apikey).Markets(markets).Date(date).Execute()

Operating hours of markets



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
    markets := "markets_example" // string | The markets for which you're requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX.
    date := "date_example" // string | The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketHoursApi.GetHours(context.Background()).Apikey(apikey).Markets(markets).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketHoursApi.GetHours``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHours`: Hours
    fmt.Fprintf(os.Stdout, "Response from `MarketHoursApi.GetHours`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | API Key | 
 **markets** | **string** | The markets for which you&#39;re requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX. | 
 **date** | **string** | The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. | 

### Return type

[**Hours**](Hours.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHours

> Hours GetMarketHours(ctx, market).Apikey(apikey).Date(date).Execute()

Operating hours of markets



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
    market := "market_example" // string | The markets for which you're requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX.
    date := "date_example" // string | The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketHoursApi.GetMarketHours(context.Background(), market).Apikey(apikey).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketHoursApi.GetMarketHours``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketHours`: Hours
    fmt.Fprintf(os.Stdout, "Response from `MarketHoursApi.GetMarketHours`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**market** | **string** | The markets for which you&#39;re requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | API Key | 

 **date** | **string** | The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. | 

### Return type

[**Hours**](Hours.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

