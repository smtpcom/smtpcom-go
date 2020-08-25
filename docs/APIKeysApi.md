# \APIKeysApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIKey**](APIKeysApi.md#GetAPIKey) | **Get** /v4/api_keys/{api_key} | Get API Key Details
[**V4ApiKeysApiKeyDelete**](APIKeysApi.md#V4ApiKeysApiKeyDelete) | **Delete** /v4/api_keys/{api_key} | Delete an API Key
[**V4ApiKeysApiKeyPatch**](APIKeysApi.md#V4ApiKeysApiKeyPatch) | **Patch** /v4/api_keys/{api_key} | Update API Key
[**V4ApiKeysGet**](APIKeysApi.md#V4ApiKeysGet) | **Get** /v4/api_keys/ | List All API Keys
[**V4ApiKeysPost**](APIKeysApi.md#V4ApiKeysPost) | **Post** /v4/api_keys/ | Create a New API Key



## GetAPIKey

> ApiKey GetAPIKey(ctx, apiKey)

Get API Key Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| API Key Identificator. | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ApiKeysApiKeyDelete

> StatusResponse V4ApiKeysApiKeyDelete(ctx, apiKey)

Delete an API Key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| API key value | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ApiKeysApiKeyPatch

> ApiKey V4ApiKeysApiKeyPatch(ctx, apiKey, optional)

Update API Key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| API key value | 
 **optional** | ***V4ApiKeysApiKeyPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ApiKeysApiKeyPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name for API key | 
 **description** | **optional.String**| Description for API Key | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ApiKeysGet

> GetApiKeys V4ApiKeysGet(ctx, )

List All API Keys

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetApiKeys**](GetApiKeys.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ApiKeysPost

> ApiKey V4ApiKeysPost(ctx, name, optional)

Create a New API Key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name for API key | 
 **optional** | ***V4ApiKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ApiKeysPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**| Description for API key | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

