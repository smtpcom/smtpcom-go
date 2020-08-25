# \AlertsApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4AlertsAlertIdDelete**](AlertsApi.md#V4AlertsAlertIdDelete) | **Delete** /v4/alerts/{alert_id} | Delete Alert
[**V4AlertsAlertIdGet**](AlertsApi.md#V4AlertsAlertIdGet) | **Get** /v4/alerts/{alert_id} | Get Alert Details
[**V4AlertsAlertIdPatch**](AlertsApi.md#V4AlertsAlertIdPatch) | **Patch** /v4/alerts/{alert_id} | Update Alert Details
[**V4AlertsGet**](AlertsApi.md#V4AlertsGet) | **Get** /v4/alerts/ | List All Allerts
[**V4AlertsPost**](AlertsApi.md#V4AlertsPost) | **Post** /v4/alerts/ | Create New Alert



## V4AlertsAlertIdDelete

> StatusResponse V4AlertsAlertIdDelete(ctx, alertId)

Delete Alert

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string**| Id of a given alert | 

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


## V4AlertsAlertIdGet

> GetAlertDetails V4AlertsAlertIdGet(ctx, alertId)

Get Alert Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string**| Id of a given alert | 

### Return type

[**GetAlertDetails**](GetAlertDetails.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4AlertsAlertIdPatch

> StatusResponse V4AlertsAlertIdPatch(ctx, alertId, optional)

Update Alert Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string**| Id of a given alert | 
 **optional** | ***V4AlertsAlertIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4AlertsAlertIdPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **optional.Float32**| A number which represents a percentage of an account’s monthly quota. Must be decimal between 0 and 1 | 

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


## V4AlertsGet

> GetAlertResponse V4AlertsGet(ctx, )

List All Allerts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetAlertResponse**](GetAlertResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4AlertsPost

> StatusResponse V4AlertsPost(ctx, type_, optional)

Create New Alert

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| An alert’s type. Currently only “monthly_quota” is supported | 
 **optional** | ***V4AlertsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4AlertsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **optional.Float32**| A number which represents a percentage of an account’s monthly quota. Must be decimal between 0 and 1 | 

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

