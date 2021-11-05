/*
Transactional Email

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SingleSendApiService SingleSendApi service
type SingleSendApiService service

type ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest struct {
	ctx                        _context.Context
	ApiService                 *SingleSendApiService
	publicSingleSendRequestEgg *PublicSingleSendRequestEgg
}

// A request object describing the email to send.
func (r ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest) PublicSingleSendRequestEgg(publicSingleSendRequestEgg PublicSingleSendRequestEgg) ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest {
	r.publicSingleSendRequestEgg = &publicSingleSendRequestEgg
	return r
}

func (r ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest) Execute() (EmailSendStatusView, *_nethttp.Response, error) {
	return r.ApiService.PostMarketingV3TransactionalSingleEmailSendSendEmailExecute(r)
}

/*
PostMarketingV3TransactionalSingleEmailSendSendEmail Send a single transactional email asynchronously.

Asynchronously send a transactional email. Returns the status of the email send with a statusId that can be used to continuously query for the status using the Email Send Status API.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest
*/
func (a *SingleSendApiService) PostMarketingV3TransactionalSingleEmailSendSendEmail(ctx _context.Context) ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest {
	return ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return EmailSendStatusView
func (a *SingleSendApiService) PostMarketingV3TransactionalSingleEmailSendSendEmailExecute(r ApiPostMarketingV3TransactionalSingleEmailSendSendEmailRequest) (EmailSendStatusView, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmailSendStatusView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SingleSendApiService.PostMarketingV3TransactionalSingleEmailSendSendEmail")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/transactional/single-email/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.publicSingleSendRequestEgg
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