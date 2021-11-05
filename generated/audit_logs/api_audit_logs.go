/*
CMS Audit Logs

Use this endpoint to query audit logs of CMS changes that occurred on your HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package audit_logs

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// AuditLogsApiService AuditLogsApi service
type AuditLogsApiService service

type ApiGetCmsV3AuditLogsGetPageRequest struct {
	ctx        _context.Context
	ApiService *AuditLogsApiService
	objectId   *[]string
	userId     *[]string
	after      *string
	before     *string
	sort       *[]string
	eventType  *[]string
	limit      *int32
	objectType *[]string
}

// Comma separated list of object ids to filter by.
func (r ApiGetCmsV3AuditLogsGetPageRequest) ObjectId(objectId []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.objectId = &objectId
	return r
}

// Comma separated list of user ids to filter by.
func (r ApiGetCmsV3AuditLogsGetPageRequest) UserId(userId []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.userId = &userId
	return r
}

// Timestamp after which audit logs will be returned
func (r ApiGetCmsV3AuditLogsGetPageRequest) After(after string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.after = &after
	return r
}

// Timestamp before which audit logs will be returned
func (r ApiGetCmsV3AuditLogsGetPageRequest) Before(before string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.before = &before
	return r
}

// The sort direction for the audit logs. (Can only sort by timestamp).
func (r ApiGetCmsV3AuditLogsGetPageRequest) Sort(sort []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.sort = &sort
	return r
}

// Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED).
func (r ApiGetCmsV3AuditLogsGetPageRequest) EventType(eventType []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.eventType = &eventType
	return r
}

// The number of logs to return.
func (r ApiGetCmsV3AuditLogsGetPageRequest) Limit(limit int32) ApiGetCmsV3AuditLogsGetPageRequest {
	r.limit = &limit
	return r
}

// Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
func (r ApiGetCmsV3AuditLogsGetPageRequest) ObjectType(objectType []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.objectType = &objectType
	return r
}

func (r ApiGetCmsV3AuditLogsGetPageRequest) Execute() (CollectionResponsePublicAuditLog, *_nethttp.Response, error) {
	return r.ApiService.GetCmsV3AuditLogsGetPageExecute(r)
}

/*
GetCmsV3AuditLogsGetPage Query audit logs

Returns audit logs based on filters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCmsV3AuditLogsGetPageRequest
*/
func (a *AuditLogsApiService) GetCmsV3AuditLogsGetPage(ctx _context.Context) ApiGetCmsV3AuditLogsGetPageRequest {
	return ApiGetCmsV3AuditLogsGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicAuditLog
func (a *AuditLogsApiService) GetCmsV3AuditLogsGetPageExecute(r ApiGetCmsV3AuditLogsGetPageRequest) (CollectionResponsePublicAuditLog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponsePublicAuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsApiService.GetCmsV3AuditLogsGetPage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/audit-logs/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.objectId != nil {
		t := *r.objectId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("objectId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("objectId", parameterToString(t, "multi"))
		}
	}
	if r.userId != nil {
		t := *r.userId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("userId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("userId", parameterToString(t, "multi"))
		}
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("eventType", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("eventType", parameterToString(t, "multi"))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.objectType != nil {
		t := *r.objectType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("objectType", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("objectType", parameterToString(t, "multi"))
		}
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