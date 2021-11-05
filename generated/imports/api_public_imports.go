/*
CRM Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// PublicImportsApiService PublicImportsApi service
type PublicImportsApiService service

type ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest struct {
	ctx        _context.Context
	ApiService *PublicImportsApiService
	importId   int64
	after      *string
	limit      *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest) After(after string) ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest) Limit(limit int32) ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest) Execute() (CollectionResponsePublicImportErrorForwardPaging, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ImportsImportIdErrorsGetErrorsExecute(r)
}

/*
GetCrmV3ImportsImportIdErrorsGetErrors Method for GetCrmV3ImportsImportIdErrorsGetErrors

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param importId
 @return ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest
*/
func (a *PublicImportsApiService) GetCrmV3ImportsImportIdErrorsGetErrors(ctx _context.Context, importId int64) ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest {
	return ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest{
		ApiService: a,
		ctx:        ctx,
		importId:   importId,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicImportErrorForwardPaging
func (a *PublicImportsApiService) GetCrmV3ImportsImportIdErrorsGetErrorsExecute(r ApiGetCrmV3ImportsImportIdErrorsGetErrorsRequest) (CollectionResponsePublicImportErrorForwardPaging, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponsePublicImportErrorForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicImportsApiService.GetCrmV3ImportsImportIdErrorsGetErrors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/imports/{importId}/errors"
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", _neturl.PathEscape(parameterToString(r.importId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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