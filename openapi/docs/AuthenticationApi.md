# \AuthenticationApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAccessToken**](AuthenticationApi.md#PostAccessToken) | **Post** /oath2/token | Post Access Token



## PostAccessToken

> EASObject PostAccessToken(ctx).EASObject(eASObject).Execute()

Post Access Token



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
    eASObject := *openapiclient.NewEASObject() // EASObject | The access token. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.PostAccessToken(context.Background()).EASObject(eASObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.PostAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAccessToken`: EASObject
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.PostAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eASObject** | [**EASObject**](EASObject.md) | The access token. | 

### Return type

[**EASObject**](EASObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

