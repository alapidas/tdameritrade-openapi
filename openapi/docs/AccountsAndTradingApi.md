# \AccountsAndTradingApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](AccountsAndTradingApi.md#CancelOrder) | **Delete** /accounts/{accountId}/orders/{orderId} | Cancel Order
[**DeleteSavedOrder**](AccountsAndTradingApi.md#DeleteSavedOrder) | **Delete** /accounts/{accountId}/savedorders/{savedOrderId} | Delete Saved Order
[**GetAccount**](AccountsAndTradingApi.md#GetAccount) | **Get** /accounts/{accountId} | Get Account
[**GetAccounts**](AccountsAndTradingApi.md#GetAccounts) | **Get** /accounts | Get Accounts
[**GetOrder**](AccountsAndTradingApi.md#GetOrder) | **Get** /accounts/{accountId}/orders/{orderId} | Get Order
[**GetOrdersByPath**](AccountsAndTradingApi.md#GetOrdersByPath) | **Get** /accounts/{accountId}/orders | Get Orders By Path
[**GetOrdersByQuery**](AccountsAndTradingApi.md#GetOrdersByQuery) | **Get** /orders | Get Orders By Query
[**GetSavedOrder**](AccountsAndTradingApi.md#GetSavedOrder) | **Get** /accounts/{accountId}/savedorders/{savedOrderId} | Get Saved Order
[**GetSavedOrdersByPath**](AccountsAndTradingApi.md#GetSavedOrdersByPath) | **Get** /accounts/{accountId}/savedorders | Get Saved Orders by Path
[**PlaceOrder**](AccountsAndTradingApi.md#PlaceOrder) | **Post** /accounts/{accountId}/orders | Place Order
[**ReplaceOrder**](AccountsAndTradingApi.md#ReplaceOrder) | **Put** /accounts/{accountId}/orders/{orderId} | Replace Order
[**ReplaceSavedOrder**](AccountsAndTradingApi.md#ReplaceSavedOrder) | **Put** /accounts/{accountId}/savedorders/{savedOrderId} | Replace Saved Order
[**SaveOrder**](AccountsAndTradingApi.md#SaveOrder) | **Post** /accounts/{accountId}/savedorders | Create Saved Order



## CancelOrder

> CancelOrder(ctx, accountId, orderId).Execute()

Cancel Order



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
    accountId := int32(56) // int32 | Account Number
    orderId := int64(789) // int64 | Order Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.CancelOrder(context.Background(), accountId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.CancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**orderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


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


## DeleteSavedOrder

> DeleteSavedOrder(ctx, accountId, savedOrderId).Execute()

Delete Saved Order



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
    accountId := int32(56) // int32 | Account Number
    savedOrderId := int64(789) // int64 | Order Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.DeleteSavedOrder(context.Background(), accountId, savedOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.DeleteSavedOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**savedOrderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedOrderRequest struct via the builder pattern


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


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get Account



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
    accountId := int32(56) // int32 | Account Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> Account GetAccounts(ctx).Execute()

Get Accounts



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
    resp, r, err := api_client.AccountsAndTradingApi.GetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccounts`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, accountId, orderId).Execute()

Get Order



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
    accountId := int32(56) // int32 | Account Number
    orderId := int64(789) // int64 | Order Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetOrder(context.Background(), accountId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**orderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByPath

> Order GetOrdersByPath(ctx, accountId).MaxResults(maxResults).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).Status(status).Execute()

Get Orders By Path



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
    accountId := int32(56) // int32 | Account Number
    maxResults := int32(56) // int32 | The maximum number of orders to retrieve.
    fromEnteredTime := "fromEnteredTime_example" // string | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz Date must be within 60 days from today's date. 'toEnteredTime' must also be set.
    toEnteredTime := "toEnteredTime_example" // string | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz. 'fromEnteredTime' must also be set.
    status := int64(789) // int64 | The maximum number of orders to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetOrdersByPath(context.Background(), accountId).MaxResults(maxResults).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetOrdersByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrdersByPath`: Order
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetOrdersByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **int32** | The maximum number of orders to retrieve. | 
 **fromEnteredTime** | **string** | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz Date must be within 60 days from today&#39;s date. &#39;toEnteredTime&#39; must also be set. | 
 **toEnteredTime** | **string** | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. &#39;fromEnteredTime&#39; must also be set. | 
 **status** | **int64** | The maximum number of orders to retrieve. | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByQuery

> GetOrdersByQuery(ctx).AccountId(accountId).MaxResults(maxResults).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).Status(status).Execute()

Get Orders By Query



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
    accountId := int32(56) // int32 | Account Number
    maxResults := int32(56) // int32 | The maximum number of orders to retrieve.
    fromEnteredTime := "fromEnteredTime_example" // string | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz Date must be within 60 days from today's date. 'toEnteredTime' must also be set.
    toEnteredTime := "toEnteredTime_example" // string | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz. 'fromEnteredTime' must also be set.
    status := int64(789) // int64 | The maximum number of orders to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetOrdersByQuery(context.Background()).AccountId(accountId).MaxResults(maxResults).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetOrdersByQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Account Number | 
 **maxResults** | **int32** | The maximum number of orders to retrieve. | 
 **fromEnteredTime** | **string** | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz Date must be within 60 days from today&#39;s date. &#39;toEnteredTime&#39; must also be set. | 
 **toEnteredTime** | **string** | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. &#39;fromEnteredTime&#39; must also be set. | 
 **status** | **int64** | The maximum number of orders to retrieve. | 

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


## GetSavedOrder

> Order GetSavedOrder(ctx, accountId, savedOrderId).Execute()

Get Saved Order



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
    accountId := int32(56) // int32 | Account Number
    savedOrderId := int64(789) // int64 | Order Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetSavedOrder(context.Background(), accountId, savedOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetSavedOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavedOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetSavedOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**savedOrderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedOrdersByPath

> []SavedOrder GetSavedOrdersByPath(ctx, accountId).Execute()

Get Saved Orders by Path



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
    accountId := int32(56) // int32 | Account Number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.GetSavedOrdersByPath(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.GetSavedOrdersByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavedOrdersByPath`: []SavedOrder
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.GetSavedOrdersByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedOrdersByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SavedOrder**](SavedOrder.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> Order PlaceOrder(ctx, accountId).Order(order).Execute()

Place Order



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
    accountId := int32(56) // int32 | Account Number
    order := *openapiclient.NewOrder() // Order | The order to place. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.PlaceOrder(context.Background(), accountId).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.PlaceOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `AccountsAndTradingApi.PlaceOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**Order**](Order.md) | The order to place. | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrder

> ReplaceOrder(ctx, accountId, orderId).Order(order).Execute()

Replace Order



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
    accountId := int32(56) // int32 | Account Number
    orderId := int64(789) // int64 | Order Number
    order := *openapiclient.NewOrder() // Order | The order to place. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.ReplaceOrder(context.Background(), accountId, orderId).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.ReplaceOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**orderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | [**Order**](Order.md) | The order to place. | 

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


## ReplaceSavedOrder

> ReplaceSavedOrder(ctx, accountId, savedOrderId).SavedOrder(savedOrder).Execute()

Replace Saved Order



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
    accountId := int32(56) // int32 | Account Number
    savedOrderId := int64(789) // int64 | Order Number
    savedOrder := *openapiclient.NewSavedOrder() // SavedOrder | The order to place. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.ReplaceSavedOrder(context.Background(), accountId, savedOrderId).SavedOrder(savedOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.ReplaceSavedOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**savedOrderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSavedOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **savedOrder** | [**SavedOrder**](SavedOrder.md) | The order to place. | 

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


## SaveOrder

> SaveOrder(ctx, accountId, orderId).SavedOrder(savedOrder).Execute()

Create Saved Order



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
    accountId := int32(56) // int32 | Account Number
    orderId := int64(789) // int64 | Order Number
    savedOrder := *openapiclient.NewSavedOrder() // SavedOrder | The order to save. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsAndTradingApi.SaveOrder(context.Background(), accountId, orderId).SavedOrder(savedOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAndTradingApi.SaveOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | Account Number | 
**orderId** | **int64** | Order Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **savedOrder** | [**SavedOrder**](SavedOrder.md) | The order to save. | 

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

