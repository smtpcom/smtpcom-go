# \CallbacksApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4CallbacksDelete**](CallbacksApi.md#V4CallbacksDelete) | **Delete** /v4/callbacks/ | Delete All Callbacks
[**V4CallbacksEventDelete**](CallbacksApi.md#V4CallbacksEventDelete) | **Delete** /v4/callbacks/{event} | Delete Callback
[**V4CallbacksEventGet**](CallbacksApi.md#V4CallbacksEventGet) | **Get** /v4/callbacks/{event} | Get Callback Details
[**V4CallbacksEventPatch**](CallbacksApi.md#V4CallbacksEventPatch) | **Patch** /v4/callbacks/{event} | Update Callback Details
[**V4CallbacksEventPost**](CallbacksApi.md#V4CallbacksEventPost) | **Post** /v4/callbacks/{event} | Create Callback
[**V4CallbacksGet**](CallbacksApi.md#V4CallbacksGet) | **Get** /v4/callbacks/ | List All Callbacks
[**V4CallbacksLogGet**](CallbacksApi.md#V4CallbacksLogGet) | **Get** /v4/callbacks/log | View Callback Logs



## V4CallbacksDelete

> StatusResponse V4CallbacksDelete(ctx, )

Delete All Callbacks

### Required Parameters

This endpoint does not need any parameter.

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


## V4CallbacksEventDelete

> StatusResponse V4CallbacksEventDelete(ctx, event, channel)

Delete Callback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**| Event for which the callback has been created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received  | 
**channel** | **string**| Name of the channel for which the callback has been created | 

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


## V4CallbacksEventGet

> GetCallbackDetails V4CallbacksEventGet(ctx, event, channel)

Get Callback Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**| Event for which the callback has been created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received  | 
**channel** | **string**| Name of the channel for which the callback has been created | 

### Return type

[**GetCallbackDetails**](GetCallbackDetails.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4CallbacksEventPatch

> StatusResponse V4CallbacksEventPatch(ctx, event, channel, optional)

Update Callback Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**| Event for which the callback should be created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received  | 
**channel** | **string**| Name of the channel for which the callback has been created | 
 **optional** | ***V4CallbacksEventPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4CallbacksEventPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **medium** | **optional.String**| Medium by which the callback data is sent. Possible values are one of:   * http   * aws  | 
 **address** | **optional.String**| Address to which the callback data is sent. This will be either a URL for http-based callbacks or the Amazon SQS queue name for SQS-based callbacks | 
 **awsData** | **optional.String**| Amazon SQS settings. &#x60;&#x60;&#x60; {key:string, secret:string} &#x60;&#x60;&#x60; must be provided if medium is of type sqs  | 

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


## V4CallbacksEventPost

> CreateCallbackResponse V4CallbacksEventPost(ctx, event, channel, medium, address, optional)

Create Callback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**| Event for which the callback should be created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received  | 
**channel** | **string**| Name of the channel for which the callback has been created. | 
**medium** | **string**| Medium to send callback data. Valid values are one of:   * http   * sqs  | 
**address** | **string**| Address of where the callback data should be sent. This will be either a URL for http-based callbacks or the Amazon SQS queue name for SQS-based callbacks. | 
 **optional** | ***V4CallbacksEventPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4CallbacksEventPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **awsData** | **optional.String**| Amazon SQS settings. &#x60;&#x60;&#x60; {key:string, secret:string} &#x60;&#x60;&#x60; must be provided if medium is of type sqs  | 

### Return type

[**CreateCallbackResponse**](CreateCallbackResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4CallbacksGet

> GetCallbackResponse V4CallbacksGet(ctx, )

List All Callbacks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetCallbackResponse**](GetCallbackResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4CallbacksLogGet

> GetCallbackLogs V4CallbacksLogGet(ctx, channel, optional)

View Callback Logs

Review all callback logs for a specific channel. It may help debug issues related to receiving callbacks on a user's side.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string**| Name of the channel for which the given callback has been created | 
 **optional** | ***V4CallbacksLogGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4CallbacksLogGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of items to return in the response. Default is 20 | 

### Return type

[**GetCallbackLogs**](GetCallbackLogs.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

