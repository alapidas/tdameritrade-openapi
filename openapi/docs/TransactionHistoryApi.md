# \TransactionHistoryApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionHistoryApi.md#GetTransaction) | **Get** /accounts/{accountId}/transactions/{transactionId} | APIs to access transaction history on the account
[**GetTransactions**](TransactionHistoryApi.md#GetTransactions) | **Get** /accounts/{accountId}/transactions | APIs to access transaction history on the account



## GetTransaction

> Transaction GetTransaction(ctx, accountId, transactionId).Execute()

APIs to access transaction history on the account



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
    transactionId := "transactionId_example" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionHistoryApi.GetTransaction(context.Background(), accountId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionHistoryApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionHistoryApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Transaction**](Transaction.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> Transaction GetTransactions(ctx, accountId).Type_(type_).Symbol(symbol).StartDate(startDate).EndDate(endDate).Execute()

APIs to access transaction history on the account



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
    type_ := openapiclient.TransactionType("ALL") // TransactionType | Only transactions with the specified type will be returned.
    symbol := "symbol_example" // string | Only transactions with the specified symbol will be returned.
    startDate := "startDate_example" // string | Only transactions after the Start Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd.
    endDate := "endDate_example" // string | Only transactions before the End Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionHistoryApi.GetTransactions(context.Background(), accountId).Type_(type_).Symbol(symbol).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionHistoryApi.GetTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactions`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionHistoryApi.GetTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**TransactionType**](TransactionType.md) | Only transactions with the specified type will be returned. | 
 **symbol** | **string** | Only transactions with the specified symbol will be returned. | 
 **startDate** | **string** | Only transactions after the Start Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd. | 
 **endDate** | **string** | Only transactions before the End Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

