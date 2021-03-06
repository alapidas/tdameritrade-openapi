# \UserInfoPreferencesApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreferences**](UserInfoPreferencesApi.md#GetPreferences) | **Get** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences
[**GetStreamerSubscriptionKeys**](UserInfoPreferencesApi.md#GetStreamerSubscriptionKeys) | **Get** /userprincipals/streamersubscriptionkeys | APIs to access user-authorized accounts and their preferences
[**GetUserPrincipals**](UserInfoPreferencesApi.md#GetUserPrincipals) | **Get** /userprincipals | APIs to access user-authorized accounts and their preferences
[**UpdatePreferences**](UserInfoPreferencesApi.md#UpdatePreferences) | **Put** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences



## GetPreferences

> Preferences GetPreferences(ctx, accountId).Execute()

APIs to access user-authorized accounts and their preferences



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
    resp, r, err := api_client.UserInfoPreferencesApi.GetPreferences(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoPreferencesApi.GetPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreferences`: Preferences
    fmt.Fprintf(os.Stdout, "Response from `UserInfoPreferencesApi.GetPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Preferences**](Preferences.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamerSubscriptionKeys

> SubscriptionKey GetStreamerSubscriptionKeys(ctx).AccountIds(accountIds).Execute()

APIs to access user-authorized accounts and their preferences



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
    accountIds := []string{"Inner_example"} // []string | Account IDs

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInfoPreferencesApi.GetStreamerSubscriptionKeys(context.Background()).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoPreferencesApi.GetStreamerSubscriptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamerSubscriptionKeys`: SubscriptionKey
    fmt.Fprintf(os.Stdout, "Response from `UserInfoPreferencesApi.GetStreamerSubscriptionKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamerSubscriptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIds** | **[]string** | Account IDs | 

### Return type

[**SubscriptionKey**](SubscriptionKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPrincipals

> UserPrincipal GetUserPrincipals(ctx).Fields(fields).Execute()

APIs to access user-authorized accounts and their preferences



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
    fields := []string{"fields=streamerSubscriptionKeys,streamerConnectionInfo"} // []string | A comma separated String which allows one to specify additional fields to return. None of these fields are returned by default. Possible values in this String can be: streamerSubscriptionKeys, streamerConnectionInfo, preferences, surrogateIds

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInfoPreferencesApi.GetUserPrincipals(context.Background()).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoPreferencesApi.GetUserPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPrincipals`: UserPrincipal
    fmt.Fprintf(os.Stdout, "Response from `UserInfoPreferencesApi.GetUserPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma separated String which allows one to specify additional fields to return. None of these fields are returned by default. Possible values in this String can be: streamerSubscriptionKeys, streamerConnectionInfo, preferences, surrogateIds | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreferences

> Preferences UpdatePreferences(ctx, accountId).UpdatePreferences(updatePreferences).Execute()

APIs to access user-authorized accounts and their preferences



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
    updatePreferences := *openapiclient.NewUpdatePreferences() // UpdatePreferences |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInfoPreferencesApi.UpdatePreferences(context.Background(), accountId).UpdatePreferences(updatePreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoPreferencesApi.UpdatePreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePreferences`: Preferences
    fmt.Fprintf(os.Stdout, "Response from `UserInfoPreferencesApi.UpdatePreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePreferences** | [**UpdatePreferences**](UpdatePreferences.md) |  | 

### Return type

[**Preferences**](Preferences.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

