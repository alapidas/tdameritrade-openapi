# \WatchlistApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWatchlist**](WatchlistApi.md#CreateWatchlist) | **Post** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
[**DeleteWatchlist**](WatchlistApi.md#DeleteWatchlist) | **Delete** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlist**](WatchlistApi.md#GetWatchlist) | **Get** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlistMultipleAccounts**](WatchlistApi.md#GetWatchlistMultipleAccounts) | **Get** /accounts/watchlists | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlistSingleAccount**](WatchlistApi.md#GetWatchlistSingleAccount) | **Get** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
[**ReplaceWatchlist**](WatchlistApi.md#ReplaceWatchlist) | **Put** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**UpdateWatchlist**](WatchlistApi.md#UpdateWatchlist) | **Patch** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist



## CreateWatchlist

> CreateWatchlist(ctx, accountId).CreateWatchlist(createWatchlist).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    createWatchlist := *openapiclient.NewCreateWatchlist() // CreateWatchlist |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.CreateWatchlist(context.Background(), accountId).CreateWatchlist(createWatchlist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.CreateWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWatchlist** | [**CreateWatchlist**](CreateWatchlist.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWatchlist

> DeleteWatchlist(ctx, accountId, watchlistId).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    watchlistId := "watchlistId_example" // string | Watchlist ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.DeleteWatchlist(context.Background(), accountId, watchlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.DeleteWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**watchlistId** | **string** | Watchlist ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlist

> []Watchlist GetWatchlist(ctx, accountId, watchlistId).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    watchlistId := "watchlistId_example" // string | Watchlist ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.GetWatchlist(context.Background(), accountId, watchlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.GetWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlist`: []Watchlist
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.GetWatchlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**watchlistId** | **string** | Watchlist ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistMultipleAccounts

> []Watchlist GetWatchlistMultipleAccounts(ctx).Execute()

APIs to perform CRUD operations on Account Watchlist



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.GetWatchlistMultipleAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.GetWatchlistMultipleAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlistMultipleAccounts`: []Watchlist
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.GetWatchlistMultipleAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistMultipleAccountsRequest struct via the builder pattern


### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistSingleAccount

> []Watchlist GetWatchlistSingleAccount(ctx, accountId).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    resp, r, err := api_client.WatchlistApi.GetWatchlistSingleAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.GetWatchlistSingleAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlistSingleAccount`: []Watchlist
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.GetWatchlistSingleAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistSingleAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceWatchlist

> ReplaceWatchlist(ctx, accountId, watchlistId).UpdateWatchlist(updateWatchlist).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    watchlistId := "watchlistId_example" // string | Watchlist ID
    updateWatchlist := *openapiclient.NewUpdateWatchlist() // UpdateWatchlist |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.ReplaceWatchlist(context.Background(), accountId, watchlistId).UpdateWatchlist(updateWatchlist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.ReplaceWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**watchlistId** | **string** | Watchlist ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWatchlist** | [**UpdateWatchlist**](UpdateWatchlist.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlist

> UpdateWatchlist(ctx, accountId, watchlistId).UpdateWatchlist(updateWatchlist).Execute()

APIs to perform CRUD operations on Account Watchlist



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
    watchlistId := "watchlistId_example" // string | Watchlist ID
    updateWatchlist := *openapiclient.NewUpdateWatchlist() // UpdateWatchlist |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WatchlistApi.UpdateWatchlist(context.Background(), accountId, watchlistId).UpdateWatchlist(updateWatchlist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.UpdateWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**watchlistId** | **string** | Watchlist ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWatchlist** | [**UpdateWatchlist**](UpdateWatchlist.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

