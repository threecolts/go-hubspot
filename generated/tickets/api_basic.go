/*
Tickets

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tickets

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
	"reflect"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// BasicApiService BasicApi service
type BasicApiService service

type ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest struct {
	ctx        _context.Context
	ApiService *BasicApiService
	ticketId   string
}

func (r ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsTicketsTicketIdArchiveExecute(r)
}

/*
DeleteCrmV3ObjectsTicketsTicketIdArchive Archive

Move an Object identified by `{ticketId}` to the recycling bin.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ticketId
 @return ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest
*/
func (a *BasicApiService) DeleteCrmV3ObjectsTicketsTicketIdArchive(ctx _context.Context, ticketId string) ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest {
	return ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		ticketId:   ticketId,
	}
}

// Execute executes the request
func (a *BasicApiService) DeleteCrmV3ObjectsTicketsTicketIdArchiveExecute(r ApiDeleteCrmV3ObjectsTicketsTicketIdArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.DeleteCrmV3ObjectsTicketsTicketIdArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/{ticketId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticketId"+"}", _neturl.PathEscape(parameterToString(r.ticketId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCrmV3ObjectsTicketsGetPageRequest struct {
	ctx          _context.Context
	ApiService   *BasicApiService
	limit        *int32
	after        *string
	properties   *[]string
	associations *[]string
	archived     *bool
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) Limit(limit int32) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) After(after string) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	r.after = &after
	return r
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) Properties(properties []string) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	r.properties = &properties
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) Associations(associations []string) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) Archived(archived bool) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	r.archived = &archived
	return r
}

func (r ApiGetCrmV3ObjectsTicketsGetPageRequest) Execute() (CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ObjectsTicketsGetPageExecute(r)
}

/*
GetCrmV3ObjectsTicketsGetPage List

Read a page of tickets. Control what is returned via the `properties` query param.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCrmV3ObjectsTicketsGetPageRequest
*/
func (a *BasicApiService) GetCrmV3ObjectsTicketsGetPage(ctx _context.Context) ApiGetCrmV3ObjectsTicketsGetPageRequest {
	return ApiGetCrmV3ObjectsTicketsGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
func (a *BasicApiService) GetCrmV3ObjectsTicketsGetPageExecute(r ApiGetCrmV3ObjectsTicketsGetPageRequest) (CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetCrmV3ObjectsTicketsGetPage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associations", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associations", parameterToString(t, "multi"))
		}
	}
	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
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

type ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest struct {
	ctx          _context.Context
	ApiService   *BasicApiService
	ticketId     string
	properties   *[]string
	associations *[]string
	archived     *bool
	idProperty   *string
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) Properties(properties []string) ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest {
	r.properties = &properties
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) Associations(associations []string) ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) Archived(archived bool) ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest {
	r.archived = &archived
	return r
}

// The name of a property whose values are unique for this object type
func (r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) IdProperty(idProperty string) ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) Execute() (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ObjectsTicketsTicketIdGetByIdExecute(r)
}

/*
GetCrmV3ObjectsTicketsTicketIdGetById Read

Read an Object identified by `{ticketId}`. `{ticketId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ticketId
 @return ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest
*/
func (a *BasicApiService) GetCrmV3ObjectsTicketsTicketIdGetById(ctx _context.Context, ticketId string) ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest {
	return ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest{
		ApiService: a,
		ctx:        ctx,
		ticketId:   ticketId,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *BasicApiService) GetCrmV3ObjectsTicketsTicketIdGetByIdExecute(r ApiGetCrmV3ObjectsTicketsTicketIdGetByIdRequest) (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetCrmV3ObjectsTicketsTicketIdGetById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/{ticketId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticketId"+"}", _neturl.PathEscape(parameterToString(r.ticketId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associations", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associations", parameterToString(t, "multi"))
		}
	}
	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
	}
	if r.idProperty != nil {
		localVarQueryParams.Add("idProperty", parameterToString(*r.idProperty, ""))
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

type ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest struct {
	ctx                     _context.Context
	ApiService              *BasicApiService
	ticketId                string
	simplePublicObjectInput *SimplePublicObjectInput
	idProperty              *string
}

func (r ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

// The name of a property whose values are unique for this object type
func (r ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest) IdProperty(idProperty string) ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest) Execute() (SimplePublicObject, *_nethttp.Response, error) {
	return r.ApiService.PatchCrmV3ObjectsTicketsTicketIdUpdateExecute(r)
}

/*
PatchCrmV3ObjectsTicketsTicketIdUpdate Update

Perform a partial update of an Object identified by `{ticketId}`. `{ticketId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ticketId
 @return ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest
*/
func (a *BasicApiService) PatchCrmV3ObjectsTicketsTicketIdUpdate(ctx _context.Context, ticketId string) ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest {
	return ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		ticketId:   ticketId,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicApiService) PatchCrmV3ObjectsTicketsTicketIdUpdateExecute(r ApiPatchCrmV3ObjectsTicketsTicketIdUpdateRequest) (SimplePublicObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.PatchCrmV3ObjectsTicketsTicketIdUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/{ticketId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticketId"+"}", _neturl.PathEscape(parameterToString(r.ticketId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
	}

	if r.idProperty != nil {
		localVarQueryParams.Add("idProperty", parameterToString(*r.idProperty, ""))
	}
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
	localVarPostBody = r.simplePublicObjectInput
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

type ApiPostCrmV3ObjectsTicketsCreateRequest struct {
	ctx                     _context.Context
	ApiService              *BasicApiService
	simplePublicObjectInput *SimplePublicObjectInput
}

func (r ApiPostCrmV3ObjectsTicketsCreateRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiPostCrmV3ObjectsTicketsCreateRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

func (r ApiPostCrmV3ObjectsTicketsCreateRequest) Execute() (SimplePublicObject, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsCreateExecute(r)
}

/*
PostCrmV3ObjectsTicketsCreate Create

Create a ticket with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard tickets is provided.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsCreateRequest
*/
func (a *BasicApiService) PostCrmV3ObjectsTicketsCreate(ctx _context.Context) ApiPostCrmV3ObjectsTicketsCreateRequest {
	return ApiPostCrmV3ObjectsTicketsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicApiService) PostCrmV3ObjectsTicketsCreateExecute(r ApiPostCrmV3ObjectsTicketsCreateRequest) (SimplePublicObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.PostCrmV3ObjectsTicketsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
	}

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
	localVarPostBody = r.simplePublicObjectInput
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