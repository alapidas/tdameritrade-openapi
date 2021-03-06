# \OptionChainsApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionChains**](OptionChainsApi.md#GetOptionChains) | **Get** /marketdata/chains | Get Option Chains for optionable symbols



## GetOptionChains

> []map[string]interface{} GetOptionChains(ctx).Apikey(apikey).Symbol(symbol).ContractType(contractType).StrikeCount(strikeCount).IncludeQuotes(includeQuotes).Strategy(strategy).Execute()

Get Option Chains for optionable symbols



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
    symbol := "symbol_example" // string | Enter one symbol
    contractType := openapiclient.PutOrCall("PUT") // PutOrCall | Type of contracts to return in the chain. Can be CALL, PUT, or ALL. Default is ALL.
    strikeCount := "strikeCount_example" // string | The number of strikes to return above and below the at-the-money price.
    includeQuotes := "'TRUE' or 'FALSE'" // string | Include quotes for options in the option chain. Can be TRUE or FALSE. Default is FALSE.
    strategy := openapiclient.Strategy("SINGLE") // Strategy | Passing a value returns a Strategy Chain. Possible values are SINGLE, ANALYTICAL (allows use of the volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values), COVERED, VERTICAL, CALENDAR, STRANGLE, STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, or ROLL. Default is SINGLE.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptionChainsApi.GetOptionChains(context.Background()).Apikey(apikey).Symbol(symbol).ContractType(contractType).StrikeCount(strikeCount).IncludeQuotes(includeQuotes).Strategy(strategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionChainsApi.GetOptionChains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionChains`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionChainsApi.GetOptionChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | API Key | 
 **symbol** | **string** | Enter one symbol | 
 **contractType** | [**PutOrCall**](PutOrCall.md) | Type of contracts to return in the chain. Can be CALL, PUT, or ALL. Default is ALL. | 
 **strikeCount** | **string** | The number of strikes to return above and below the at-the-money price. | 
 **includeQuotes** | **string** | Include quotes for options in the option chain. Can be TRUE or FALSE. Default is FALSE. | 
 **strategy** | [**Strategy**](Strategy.md) | Passing a value returns a Strategy Chain. Possible values are SINGLE, ANALYTICAL (allows use of the volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values), COVERED, VERTICAL, CALENDAR, STRANGLE, STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, or ROLL. Default is SINGLE. | 

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

