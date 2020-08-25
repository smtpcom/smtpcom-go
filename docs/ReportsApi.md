# \ReportsApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4ReportsGet**](ReportsApi.md#V4ReportsGet) | **Get** /v4/reports/ | Get All Reports
[**V4ReportsOndemandPost**](ReportsApi.md#V4ReportsOndemandPost) | **Post** /v4/reports/ondemand | Create On-Demand Report
[**V4ReportsPeriodicPost**](ReportsApi.md#V4ReportsPeriodicPost) | **Post** /v4/reports/periodic | Create Periodic Report
[**V4ReportsPeriodicReportIdDelete**](ReportsApi.md#V4ReportsPeriodicReportIdDelete) | **Delete** /v4/reports/periodic/{report_id} | Delete a Periodic Report
[**V4ReportsPeriodicReportIdPatch**](ReportsApi.md#V4ReportsPeriodicReportIdPatch) | **Patch** /v4/reports/periodic/{report_id} | Update Periodic Report
[**V4ReportsReportIdGet**](ReportsApi.md#V4ReportsReportIdGet) | **Get** /v4/reports/{report_id} | Get Report Details



## V4ReportsGet

> Reports V4ReportsGet(ctx, )

Get All Reports

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Reports**](Reports.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ReportsOndemandPost

> StatusResponse V4ReportsOndemandPost(ctx, start, optional)

Create On-Demand Report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **int32**| Start date/time of the report in RFC 2822 or UNIX epoch format | 
 **optional** | ***V4ReportsOndemandPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ReportsOndemandPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **optional.String**| Name of the channel for which a given report has been defined | 
 **type_** | **optional.String**| Type or report format. If not specified defaults to “csv” - currently the only supported type | 
 **end** | **optional.String**| End date/time of the report in RFC 2822 or UNIX epoch format (default - now) | 
 **domain** | **optional.String**| Filter by the “from” domain of emails | 
 **rcptDomain** | **optional.String**| Filter by the “to” domain of emails | 
 **events** | **optional.String**| Filter by event type. Valid event are:  * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints  | 
 **columns** | **optional.String**| Array of columns to be specified in the report. These can differ based on any specified event type filter.   Possible column values for all reports are: * &#x60;message_id&#x60; - Unique message ID * from - From Address * to - Recipient Address * time_rcv - Date Received in RFC 2822 or UNIX epoch format * time_snt - Date Delivered in RFC 2822 or UNIX epoch format * channel - Name of a channel  Additional column values for message reports (hard_bounced, failed, delivered, total) include: * status - Status of delivery * code - SMTP Response Code * message - SMTP Response message * tries - Amount of re-tries (deferred states before final)  Additional column values for spam reports include:  * report_time - Date when an abuse complaint has been reported, RFC 2822 or UNIX epoch format * subject - Email Subject  Additional column values for pending reports include:  * status - Current email status  If not specified all relevant columns are included.  | 

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


## V4ReportsPeriodicPost

> StatusResponse V4ReportsPeriodicPost(ctx, frequency, reportTime, optional)

Create Periodic Report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frequency** | **string**| Report frequency – one of:   * daily - every day at a specified hour   * weekly  - Mondays at a specified hour   * monthly - 1st day of the month at a specified hour  | 
**reportTime** | **int32**| The hour at which the report should be sent. Value values range from 0 to 23 | 
 **optional** | ***V4ReportsPeriodicPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ReportsPeriodicPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **channel** | **optional.String**| Name of the channel for which a given report has been defined | 
 **notifyMethod** | **optional.String**| Notification method to be used when report is completed and can be downloaded | 
 **notifyDest** | **optional.String**| A valid URL which will accept the report completion notification. The payload will be &#x60;&#x60;&#x60;   {\&quot;message\&quot;: \&quot;success\&quot;, \&quot;id\&quot;: string} &#x60;&#x60;&#x60; where &#x60;id&#x60; is a Unique report ID  | 
 **domain** | **optional.String**| Filter by the “From” domain of emails | 
 **rcptDomain** | **optional.String**| Filter by the “To” domain of emails | 
 **events** | **optional.String**| Filter by event type. Valid event are:  * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints  | 
 **columns** | **optional.String**| Array of columns to be specified in the report. These can differ based on any specified event type filter.   Possible column values are: * &#x60;message_id&#x60; - Unique message ID * from - From Address * to - Recipient Address * time_rcv - Date Received in RFC 2822 or UNIX epoch format * time_snt - Date Delivered in RFC 2822 or UNIX epoch format * channel - Name of a channel  Additional column values for message reports (hard_bounced, failed, delivered, total) include: * status - Status of delivery * code - SMTP Response Code * message - SMTP Response message * tries - Amount of re-tries (deferred states before final)  Additional column values for spam reports include:  * report_time - Date when an abuse complaint has been reported, RFC 2822 or UNIX epoch format * subject - Email Subject  Additional column values for pending reports include:  * status - Current email status  If not specified all relevant columns are included.  | 

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


## V4ReportsPeriodicReportIdDelete

> StatusResponse V4ReportsPeriodicReportIdDelete(ctx, reportId)

Delete a Periodic Report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string**| Id of a given report | 

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


## V4ReportsPeriodicReportIdPatch

> StatusResponse V4ReportsPeriodicReportIdPatch(ctx, reportId, frequency, reportTime, optional)

Update Periodic Report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string**| Id of the report to be updated | 
**frequency** | **string**| Report frequency – one of:  * daily - every day at a specified hour  * weekly  - Mondays at a specified hour  * monthly - first day of the month at a specified hour.  | 
**reportTime** | **int32**| The hour at which the report should be sent. Value values range from 0 to 23 | 
 **optional** | ***V4ReportsPeriodicReportIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4ReportsPeriodicReportIdPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **channel** | **optional.String**| Name of channel (sender). If not specified all channels will be reported | 
 **events** | **optional.String**| Filter by event type. Valid event are: * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints  If not specified all events are included.  | 

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


## V4ReportsReportIdGet

> Report V4ReportsReportIdGet(ctx, reportId)

Get Report Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string**| ID of a given report | 

### Return type

[**Report**](Report.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

