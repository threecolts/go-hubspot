/*
OAuthService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

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

// AccessTokensApiService AccessTokensApi service
type AccessTokensApiService service

type ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest struct {
	ctx        _context.Context
	ApiService *AccessTokensApiService
	token      string
}

func (r ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest) Execute() (AccessTokenInfoResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOauthV1AccessTokensTokenGetAccessTokenExecute(r)
}

/*
GetOauthV1AccessTokensTokenGetAccessToken Method for GetOauthV1AccessTokensTokenGetAccessToken

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token
 @return ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest
*/
func (a *AccessTokensApiService) GetOauthV1AccessTokensTokenGetAccessToken(ctx _context.Context, token string) ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest {
	return ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
		token:      token,
	}
}

// Execute executes the request
//  @return AccessTokenInfoResponse
func (a *AccessTokensApiService) GetOauthV1AccessTokensTokenGetAccessTokenExecute(r ApiGetOauthV1AccessTokensTokenGetAccessTokenRequest) (AccessTokenInfoResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccessTokenInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokensApiService.GetOauthV1AccessTokensTokenGetAccessToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/v1/access-tokens/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", _neturl.PathEscape(parameterToString(r.token, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
