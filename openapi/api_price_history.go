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

// PriceHistoryApiService PriceHistoryApi service
type PriceHistoryApiService service

type ApiGetPriceHistoryRequest struct {
	ctx _context.Context
	ApiService *PriceHistoryApiService
	apikey *string
	symbol string
	periodType *string
	period *int32
	frequencyType *string
	frequency *int32
	endDate *string
	startDate *string
	needExtendedHoursData *bool
	x *string
}

func (r ApiGetPriceHistoryRequest) Apikey(apikey string) ApiGetPriceHistoryRequest {
	r.apikey = &apikey
	return r
}
func (r ApiGetPriceHistoryRequest) PeriodType(periodType string) ApiGetPriceHistoryRequest {
	r.periodType = &periodType
	return r
}
func (r ApiGetPriceHistoryRequest) Period(period int32) ApiGetPriceHistoryRequest {
	r.period = &period
	return r
}
func (r ApiGetPriceHistoryRequest) FrequencyType(frequencyType string) ApiGetPriceHistoryRequest {
	r.frequencyType = &frequencyType
	return r
}
func (r ApiGetPriceHistoryRequest) Frequency(frequency int32) ApiGetPriceHistoryRequest {
	r.frequency = &frequency
	return r
}
func (r ApiGetPriceHistoryRequest) EndDate(endDate string) ApiGetPriceHistoryRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGetPriceHistoryRequest) StartDate(startDate string) ApiGetPriceHistoryRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGetPriceHistoryRequest) NeedExtendedHoursData(needExtendedHoursData bool) ApiGetPriceHistoryRequest {
	r.needExtendedHoursData = &needExtendedHoursData
	return r
}
func (r ApiGetPriceHistoryRequest) X(x string) ApiGetPriceHistoryRequest {
	r.x = &x
	return r
}

func (r ApiGetPriceHistoryRequest) Execute() (CandleList, *_nethttp.Response, error) {
	return r.ApiService.GetPriceHistoryExecute(r)
}

/*
 * GetPriceHistory Historical price data for charts
 * Get price history for a symbol
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol Enter one symbol
 * @return ApiGetPriceHistoryRequest
 */
func (a *PriceHistoryApiService) GetPriceHistory(ctx _context.Context, symbol string) ApiGetPriceHistoryRequest {
	return ApiGetPriceHistoryRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

/*
 * Execute executes the request
 * @return CandleList
 */
func (a *PriceHistoryApiService) GetPriceHistoryExecute(r ApiGetPriceHistoryRequest) (CandleList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CandleList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceHistoryApiService.GetPriceHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketdata/{symbol}/pricehistory"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", _neturl.PathEscape(parameterToString(r.symbol, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.apikey == nil {
		return localVarReturnValue, nil, reportError("apikey is required and must be specified")
	}
	if r.periodType == nil {
		return localVarReturnValue, nil, reportError("periodType is required and must be specified")
	}
	if r.period == nil {
		return localVarReturnValue, nil, reportError("period is required and must be specified")
	}
	if r.frequencyType == nil {
		return localVarReturnValue, nil, reportError("frequencyType is required and must be specified")
	}
	if r.frequency == nil {
		return localVarReturnValue, nil, reportError("frequency is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, reportError("endDate is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}
	if r.needExtendedHoursData == nil {
		return localVarReturnValue, nil, reportError("needExtendedHoursData is required and must be specified")
	}
	if r.x == nil {
		return localVarReturnValue, nil, reportError("x is required and must be specified")
	}

	localVarQueryParams.Add("apikey", parameterToString(*r.apikey, ""))
	localVarQueryParams.Add("periodType", parameterToString(*r.periodType, ""))
	localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	localVarQueryParams.Add("frequencyType", parameterToString(*r.frequencyType, ""))
	localVarQueryParams.Add("frequency", parameterToString(*r.frequency, ""))
	localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	localVarQueryParams.Add("needExtendedHoursData", parameterToString(*r.needExtendedHoursData, ""))
	localVarQueryParams.Add("x", parameterToString(*r.x, ""))
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
