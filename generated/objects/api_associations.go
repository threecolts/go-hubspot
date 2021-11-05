/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

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

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest struct {
	ctx             _context.Context
	ApiService      *AssociationsApiService
	objectType      string
	objectId        string
	toObjectType    string
	toObjectId      string
	associationType string
}

func (r ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r)
}

/*
DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive Remove an association between two objects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param objectId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest
*/
func (a *AssociationsApiService) DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(ctx _context.Context, objectType string, objectId string, toObjectType string, toObjectId string, associationType string) ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest {
	return ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest{
		ApiService:      a,
		ctx:             ctx,
		objectType:      objectType,
		objectId:        objectId,
		toObjectType:    toObjectType,
		toObjectId:      toObjectId,
		associationType: associationType,
	}
}

// Execute executes the request
func (a *AssociationsApiService) DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveExecute(r ApiDeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", _neturl.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectId"+"}", _neturl.PathEscape(parameterToString(r.objectId, "")), -1)
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

type ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest struct {
	ctx          _context.Context
	ApiService   *AssociationsApiService
	objectType   string
	objectId     string
	toObjectType string
	after        *string
	limit        *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest) After(after string) ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest) Limit(limit int32) ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest) Execute() (CollectionResponseAssociatedIdForwardPaging, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllExecute(r)
}

/*
GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAll List associations of an object by type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param objectId
 @param toObjectType
 @return ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest
*/
func (a *AssociationsApiService) GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAll(ctx _context.Context, objectType string, objectId string, toObjectType string) ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest {
	return ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest{
		ApiService:   a,
		ctx:          ctx,
		objectType:   objectType,
		objectId:     objectId,
		toObjectType: toObjectType,
	}
}

// Execute executes the request
//  @return CollectionResponseAssociatedIdForwardPaging
func (a *AssociationsApiService) GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllExecute(r ApiGetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAllRequest) (CollectionResponseAssociatedIdForwardPaging, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponseAssociatedIdForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAll")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", _neturl.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectId"+"}", _neturl.PathEscape(parameterToString(r.objectId, "")), -1)
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

type ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest struct {
	ctx             _context.Context
	ApiService      *AssociationsApiService
	objectType      string
	objectId        string
	toObjectType    string
	toObjectId      string
	associationType string
}

func (r ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) Execute() (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	return r.ApiService.PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r)
}

/*
PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate Associate an object with another object

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param objectId
 @param toObjectType
 @param toObjectId
 @param associationType
 @return ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest
*/
func (a *AssociationsApiService) PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(ctx _context.Context, objectType string, objectId string, toObjectType string, toObjectId string, associationType string) ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest {
	return ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		objectType:      objectType,
		objectId:        objectId,
		toObjectType:    toObjectType,
		toObjectId:      toObjectId,
		associationType: associationType,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *AssociationsApiService) PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateExecute(r ApiPutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreateRequest) (SimplePublicObjectWithAssociations, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", _neturl.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectId"+"}", _neturl.PathEscape(parameterToString(r.objectId, "")), -1)
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