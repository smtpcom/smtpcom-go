# \ChannelsApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSender**](ChannelsApi.md#GetSender) | **Get** /v4/channels/{name} | Get Channel Details
[**V4ChannelsGet**](ChannelsApi.md#V4ChannelsGet) | **Get** /v4/channels/ | Get All Channels
[**V4ChannelsNameDelete**](ChannelsApi.md#V4ChannelsNameDelete) | **Delete** /v4/channels/{name} | Delete a Channel
[**V4ChannelsNamePatch**](ChannelsApi.md#V4ChannelsNamePatch) | **Patch** /v4/channels/{name} | Update Channel Details
[**V4ChannelsPost**](ChannelsApi.md#V4ChannelsPost) | **Post** /v4/channels/ | Create a New Channel



## GetSender

> Channel GetSender(ctx, name)

Get Channel Details

**Note:** This method doesn't return archived channels. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the channel (sender) | 

### Return type

[**Channel**](Channel.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ChannelsGet

> Channels V4ChannelsGet(ctx, )

Get All Channels

**Note:** This method does not return archived channels. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Channels**](Channels.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ChannelsNameDelete

> Channel V4ChannelsNameDelete(ctx, name)

Delete a Channel

**Note:** This method doesn’t permanently delete the channel but rather sets the status to “canceled”. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the channel (sender) | 

### Return type

[**Channel**](Channel.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ChannelsNamePatch

> Channel V4ChannelsNamePatch(ctx, name, optional)

Update Channel Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the channel (sender) | 
 **optional** | ***V4ChannelsNamePatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ChannelsNamePatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smtpUsername** | **optional.String**| Username for the channel | 
 **smtpPassword** | **optional.String**| Password for the channel | 
 **quota** | **optional.Int32**| Quota for the channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ChannelsPost

> Channel V4ChannelsPost(ctx, name, smtpUsername, smtpPassword, quota)

Create a New Channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the channel (sender) | 
**smtpUsername** | **string**| Username for the channel | 
**smtpPassword** | **string**| Password for the channel | 
**quota** | **int32**| Quota for the channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

