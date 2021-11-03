/*
Calling Extensions API

Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling

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

// SettingsApiService SettingsApi service
type SettingsApiService service

type ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest struct {
	ctx        _context.Context
	ApiService *SettingsApiService
	appId      int32
}

func (r ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmV3ExtensionsCallingAppIdSettingsArchiveExecute(r)
}

/*
DeleteCrmV3ExtensionsCallingAppIdSettingsArchive Delete calling settings

Deletes this calling extension. This will remove your service as an option for all connected accounts.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest
*/
func (a *SettingsApiService) DeleteCrmV3ExtensionsCallingAppIdSettingsArchive(ctx _context.Context, appId int32) ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest {
	return ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
func (a *SettingsApiService) DeleteCrmV3ExtensionsCallingAppIdSettingsArchiveExecute(r ApiDeleteCrmV3ExtensionsCallingAppIdSettingsArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.DeleteCrmV3ExtensionsCallingAppIdSettingsArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

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
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest struct {
	ctx        _context.Context
	ApiService *SettingsApiService
	appId      int32
}

func (r ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest) Execute() (SettingsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsCallingAppIdSettingsGetByIdExecute(r)
}

/*
GetCrmV3ExtensionsCallingAppIdSettingsGetById Get calling settings

Returns the calling extension settings configured for your app.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest
*/
func (a *SettingsApiService) GetCrmV3ExtensionsCallingAppIdSettingsGetById(ctx _context.Context, appId int32) ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest {
	return ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//  @return SettingsResponse
func (a *SettingsApiService) GetCrmV3ExtensionsCallingAppIdSettingsGetByIdExecute(r ApiGetCrmV3ExtensionsCallingAppIdSettingsGetByIdRequest) (SettingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.GetCrmV3ExtensionsCallingAppIdSettingsGetById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

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
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest struct {
	ctx                  _context.Context
	ApiService           *SettingsApiService
	appId                int32
	settingsPatchRequest *SettingsPatchRequest
}

// Updated details for the settings.
func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest) SettingsPatchRequest(settingsPatchRequest SettingsPatchRequest) ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest {
	r.settingsPatchRequest = &settingsPatchRequest
	return r
}

func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest) Execute() (SettingsResponse, *_nethttp.Response, error) {
	return r.ApiService.PatchCrmV3ExtensionsCallingAppIdSettingsUpdateExecute(r)
}

/*
PatchCrmV3ExtensionsCallingAppIdSettingsUpdate Update settings

Updates existing calling extension settings.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest
*/
func (a *SettingsApiService) PatchCrmV3ExtensionsCallingAppIdSettingsUpdate(ctx _context.Context, appId int32) ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest {
	return ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//  @return SettingsResponse
func (a *SettingsApiService) PatchCrmV3ExtensionsCallingAppIdSettingsUpdateExecute(r ApiPatchCrmV3ExtensionsCallingAppIdSettingsUpdateRequest) (SettingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.PatchCrmV3ExtensionsCallingAppIdSettingsUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.settingsPatchRequest == nil {
		return localVarReturnValue, nil, reportError("settingsPatchRequest is required and must be specified")
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
	localVarPostBody = r.settingsPatchRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest struct {
	ctx             _context.Context
	ApiService      *SettingsApiService
	appId           int32
	settingsRequest *SettingsRequest
}

// Settings state to create with.
func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest) SettingsRequest(settingsRequest SettingsRequest) ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest {
	r.settingsRequest = &settingsRequest
	return r
}

func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest) Execute() (SettingsResponse, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ExtensionsCallingAppIdSettingsCreateExecute(r)
}

/*
PostCrmV3ExtensionsCallingAppIdSettingsCreate Configure a calling extension

Used to set the menu label, target iframe URL, and dimensions for your calling extension.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest
*/
func (a *SettingsApiService) PostCrmV3ExtensionsCallingAppIdSettingsCreate(ctx _context.Context, appId int32) ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest {
	return ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//  @return SettingsResponse
func (a *SettingsApiService) PostCrmV3ExtensionsCallingAppIdSettingsCreateExecute(r ApiPostCrmV3ExtensionsCallingAppIdSettingsCreateRequest) (SettingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.PostCrmV3ExtensionsCallingAppIdSettingsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.settingsRequest == nil {
		return localVarReturnValue, nil, reportError("settingsRequest is required and must be specified")
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
	localVarPostBody = r.settingsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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
