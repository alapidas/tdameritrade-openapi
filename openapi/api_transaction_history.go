/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TransactionHistoryApiService TransactionHistoryApi service
type TransactionHistoryApiService service

type ApiGetTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionHistoryApiService
	accountId string
	transactionId string
}


func (r ApiGetTransactionRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.GetTransactionExecute(r)
}

/*
 * GetTransaction APIs to access transaction history on the account
 * Transaction for a specific account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @param transactionId Transaction ID
 * @return ApiGetTransactionRequest
 */
func (a *TransactionHistoryApiService) GetTransaction(ctx _context.Context, accountId string, transactionId string) ApiGetTransactionRequest {
	return ApiGetTransactionRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		transactionId: transactionId,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionHistoryApiService) GetTransactionExecute(r ApiGetTransactionRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionHistoryApiService.GetTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{accountId}/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", _neturl.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", _neturl.PathEscape(parameterToString(r.transactionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTransactionsRequest struct {
	ctx _context.Context
	ApiService *TransactionHistoryApiService
	accountId string
	type_ *TransactionType
	symbol *string
	startDate *string
	endDate *string
}

func (r ApiGetTransactionsRequest) Type_(type_ TransactionType) ApiGetTransactionsRequest {
	r.type_ = &type_
	return r
}
func (r ApiGetTransactionsRequest) Symbol(symbol string) ApiGetTransactionsRequest {
	r.symbol = &symbol
	return r
}
func (r ApiGetTransactionsRequest) StartDate(startDate string) ApiGetTransactionsRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGetTransactionsRequest) EndDate(endDate string) ApiGetTransactionsRequest {
	r.endDate = &endDate
	return r
}

func (r ApiGetTransactionsRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.GetTransactionsExecute(r)
}

/*
 * GetTransactions APIs to access transaction history on the account
 * Transactions for a specific account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Account ID
 * @return ApiGetTransactionsRequest
 */
func (a *TransactionHistoryApiService) GetTransactions(ctx _context.Context, accountId string) ApiGetTransactionsRequest {
	return ApiGetTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionHistoryApiService) GetTransactionsExecute(r ApiGetTransactionsRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionHistoryApiService.GetTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{accountId}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", _neturl.PathEscape(parameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, reportError("endDate is required and must be specified")
	}

	localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	localVarQueryParams.Add("symbol", parameterToString(*r.symbol, ""))
	localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
