# \MessagesApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4MessagesGet**](MessagesApi.md#V4MessagesGet) | **Get** /v4/messages | Get Detailed Messages Info
[**V4MessagesMimePost**](MessagesApi.md#V4MessagesMimePost) | **Post** /v4/messages/mime | Send MIME Message
[**V4MessagesPost**](MessagesApi.md#V4MessagesPost) | **Post** /v4/messages | Send a Message



## V4MessagesGet

> MessagesResponse V4MessagesGet(ctx, start, limit, offset, channel, optional)

Get Detailed Messages Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **string**| The starting time. RFC 2822 or UNIX epoch format | 
**limit** | **int32**| Maximum number of items to return. | 
**offset** | **int32**| Number of items to skip before returning the results. | 
**channel** | **string**| Name of  the channel used to send messages | 
 **optional** | ***V4MessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4MessagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **end** | **optional.String**| The ending time. If not specified, defaults to now. RFC 2822 or UNIX epoch format. | 
 **event** | **optional.String**| Array of any event names for which stats has been requested (&#39;accepted&#39;, &#39;failed&#39;, &#39;delivered&#39;). | 
 **msgId** | **optional.String**| Unique message ID | 

### Return type

[**MessagesResponse**](MessagesResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4MessagesMimePost

> MimeResponse V4MessagesMimePost(ctx, inlineObject1)

Send MIME Message

This method will send a prepared RFC compliant MIME message via a specified channel **Note:** All restrictions from the section entitled `Send a Message` are implied here. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**MimeResponse**](MimeResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4MessagesPost

> PostMessageResponse V4MessagesPost(ctx, optional)

Send a Message

Send a message over HTTP/HTTPS protocol using a specified channel. The request can generate only an HTTP 400 error and will return validation error data. The actual suppression, delivery attempt or SMTP error can be received via callbacks. See the section on Callbacks for more details.  <br> **Limitations:** * The number of recipients in a single email request is limited to `1000` (including to, cc and bcc) * The total payload of the generated MIME message cannot exceed `10Mb` in size  <br> **Notes:** * When JUST two parts provided with `text/plain` and `text/html` types, they are automatically wrapped into a `mime/alternative` container * Adding any other part or attachment to either a single part or attachment or `mime/alternative` automatically wraps all mime containers into a `mime/mixed` container * Content type message/rfc822 is prohibited 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V4MessagesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4MessagesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**PostMessageResponse**](PostMessageResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

