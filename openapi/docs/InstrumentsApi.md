# \InstrumentsApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstrument**](InstrumentsApi.md#GetInstrument) | **Get** /instruments/{accountId} | Search for instrument and fundamental data
[**SearchInstruments**](InstrumentsApi.md#SearchInstruments) | **Get** /instruments | Search for instrument and fundamental data



## GetInstrument

> []map[string]interface{} GetInstrument(ctx, accountId).Execute()

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
    accountId := "accountId_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstrumentsApi.GetInstrument(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsApi.GetInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstrument`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstrumentsApi.GetInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SearchInstruments

> []map[string]interface{} SearchInstruments(ctx).Apikey(apikey).Symbol(symbol).Projection(projection).Execute()

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
    apikey := "apikey_example" // string | API Key
    symbol := "symbol_example" // string | Value to pass to the search. See projection description for more information.
    projection := "projection_example" // string | The type of request:symbol-search: Retrieve instrument data of a specific symbol or cusip symbol-regex: Retrieve instrument data for all symbols matching regex. Example: symbol=XYZ.* will return all symbols beginning with XYZ desc-search: Retrieve instrument data for instruments whose description contains the word supplied. Example: symbol=FakeCompany will return all instruments with FakeCompany in the description. desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.  fundamental: Returns fundamental data for a single instrument specified by exact symbol.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstrumentsApi.SearchInstruments(context.Background()).Apikey(apikey).Symbol(symbol).Projection(projection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsApi.SearchInstruments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchInstruments`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstrumentsApi.SearchInstruments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | API Key | 
 **symbol** | **string** | Value to pass to the search. See projection description for more information. | 
 **projection** | **string** | The type of request:symbol-search: Retrieve instrument data of a specific symbol or cusip symbol-regex: Retrieve instrument data for all symbols matching regex. Example: symbol&#x3D;XYZ.* will return all symbols beginning with XYZ desc-search: Retrieve instrument data for instruments whose description contains the word supplied. Example: symbol&#x3D;FakeCompany will return all instruments with FakeCompany in the description. desc-regex: Search description with full regex support. Example: symbol&#x3D;XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.  fundamental: Returns fundamental data for a single instrument specified by exact symbol. | 

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

