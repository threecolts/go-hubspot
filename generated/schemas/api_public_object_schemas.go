/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

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

// PublicObjectSchemasApiService PublicObjectSchemasApi service
type PublicObjectSchemasApiService service

type ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest struct {
	ctx        _context.Context
	ApiService *PublicObjectSchemasApiService
	objectType string
}

func (r ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeExecute(r)
}

/*
DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge Method for DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest

Deprecated
*/
func (a *PublicObjectSchemasApiService) DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge(ctx _context.Context, objectType string) ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest {
	return ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
// Deprecated
func (a *PublicObjectSchemasApiService) DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeExecute(r ApiDeleteCrmObjectSchemasV3SchemasObjectTypePurgePurgeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicObjectSchemasApiService.DeleteCrmObjectSchemasV3SchemasObjectTypePurgePurge")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/schemas/{objectType}/purge"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", _neturl.PathEscape(parameterToString(r.objectType, "")), -1)

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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
