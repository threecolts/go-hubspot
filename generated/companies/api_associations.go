/*
Companies

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package companies

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

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct {
	ctx             _context.Context
	ApiService      *AssociationsApiService
	companyId       string
	toObjectType    string
	toObjectId      string
	associationType string
}

func (r ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r)
}

/*
DeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive Remove an association between two companies

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest
*/
func (a *AssociationsApiService) DeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx _context.Context, companyId string, toObjectType string, toObjectId string, associationType string) ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest {
	return ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest{
		ApiService:      a,
		ctx:             ctx,
		companyId:       companyId,
		toObjectType:    toObjectType,
		toObjectId:      toObjectId,
		associationType: associationType,
	}
}

// Execute executes the request
func (a *AssociationsApiService) DeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r ApiDeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.DeleteCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", _neturl.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", _neturl.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", _neturl.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", _neturl.PathEscape(parameterToString(r.associationType, "")), -1)

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

type ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest struct {
	ctx          _context.Context
	ApiService   *AssociationsApiService
	companyId    string
	toObjectType string
	after        *string
	limit        *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest) After(after string) ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest) Limit(limit int32) ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest) Execute() (CollectionResponseAssociatedIdForwardPaging, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllExecute(r)
}

/*
GetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAll List associations of a company by type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId
 @param toObjectType
 @return ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest
*/
func (a *AssociationsApiService) GetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAll(ctx _context.Context, companyId string, toObjectType string) ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest {
	return ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest{
		ApiService:   a,
		ctx:          ctx,
		companyId:    companyId,
		toObjectType: toObjectType,
	}
}

// Execute executes the request
//  @return CollectionResponseAssociatedIdForwardPaging
func (a *AssociationsApiService) GetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllExecute(r ApiGetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAllRequest) (CollectionResponseAssociatedIdForwardPaging, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponseAssociatedIdForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.GetCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeGetAll")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}/associations/{toObjectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", _neturl.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", _neturl.PathEscape(parameterToString(r.toObjectType, "")), -1)

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

type ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct {
	ctx             _context.Context
	ApiService      *AssociationsApiService
	companyId       string
	toObjectType    string
	toObjectId      string
	associationType string
}

func (r ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) Execute() (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	return r.ApiService.PutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r)
}

/*
PutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate Associate a company with another object

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest
*/
func (a *AssociationsApiService) PutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx _context.Context, companyId string, toObjectType string, toObjectId string, associationType string) ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest {
	return ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		companyId:       companyId,
		toObjectType:    toObjectType,
		toObjectId:      toObjectId,
		associationType: associationType,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *AssociationsApiService) PutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r ApiPutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.PutCrmV3ObjectsCompaniesCompanyIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", _neturl.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", _neturl.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", _neturl.PathEscape(parameterToString(r.toObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationType"+"}", _neturl.PathEscape(parameterToString(r.associationType, "")), -1)

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
