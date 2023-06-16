/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/threecolts/go-hubspot"
	"net/url"
	"strings"
)

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiAssociationsArchiveRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	toObjectId           string
	associationType      string
}

func (r ApiAssociationsArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.AssociationsArchiveExecute(r)
}

/*
AssociationsArchive Remove an association between two feedback submissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiAssociationsArchiveRequest
*/
func (a *AssociationsApiService) AssociationsArchive(ctx context.Context, feedbackSubmissionId string, toObjectType string, toObjectId string, associationType string) ApiAssociationsArchiveRequest {
	return ApiAssociationsArchiveRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
		associationType:      associationType,
	}
}

// Execute executes the request
func (a *AssociationsApiService) AssociationsArchiveExecute(r ApiAssociationsArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.AssociationsArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", url.PathEscape(parameterToString(r.associationType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAssociationsCreateRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	toObjectId           string
	associationType      string
}

func (r ApiAssociationsCreateRequest) Execute() (*SimplePublicObjectWithAssociations, *http.Response, error) {
	return r.ApiService.AssociationsCreateExecute(r)
}

/*
AssociationsCreate Associate a feedback submission with another object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiAssociationsCreateRequest
*/
func (a *AssociationsApiService) AssociationsCreate(ctx context.Context, feedbackSubmissionId string, toObjectType string, toObjectId string, associationType string) ApiAssociationsCreateRequest {
	return ApiAssociationsCreateRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
		associationType:      associationType,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *AssociationsApiService) AssociationsCreateExecute(r ApiAssociationsCreateRequest) (*SimplePublicObjectWithAssociations, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.AssociationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", url.PathEscape(parameterToString(r.associationType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAssociationsGetPageRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId string
	toObjectType         string
	after                *string
	limit                *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiAssociationsGetPageRequest) After(after string) ApiAssociationsGetPageRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiAssociationsGetPageRequest) Limit(limit int32) ApiAssociationsGetPageRequest {
	r.limit = &limit
	return r
}

func (r ApiAssociationsGetPageRequest) Execute() (*CollectionResponseAssociatedIdForwardPaging, *http.Response, error) {
	return r.ApiService.AssociationsGetPageExecute(r)
}

/*
AssociationsGetPage List associations of a feedback submission by type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feedbackSubmissionId
 @param toObjectType
 @return ApiAssociationsGetPageRequest
*/
func (a *AssociationsApiService) AssociationsGetPage(ctx context.Context, feedbackSubmissionId string, toObjectType string) ApiAssociationsGetPageRequest {
	return ApiAssociationsGetPageRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
	}
}

// Execute executes the request
//  @return CollectionResponseAssociatedIdForwardPaging
func (a *AssociationsApiService) AssociationsGetPageExecute(r ApiAssociationsGetPageRequest) (*CollectionResponseAssociatedIdForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseAssociatedIdForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.AssociationsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
