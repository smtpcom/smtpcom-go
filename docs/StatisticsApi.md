# \StatisticsApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4StatsDurationSliceSliceSpecifierGroupByGet**](StatisticsApi.md#V4StatsDurationSliceSliceSpecifierGroupByGet) | **Get** /v4/stats/{duration}/{slice}/{slice_specifier}/{group_by} | Return event aggregates for the specified slices and duration. Slices can be chained.



## V4StatsDurationSliceSliceSpecifierGroupByGet

> StatsResponse V4StatsDurationSliceSliceSpecifierGroupByGet(ctx, start, duration, slice, sliceSpecifier, groupBy, limit, offset, optional)

Return event aggregates for the specified slices and duration. Slices can be chained.

**Get stats for a period**<br> Request:<br> *_/v4/stats/last_day*<br> *_/v4/stats?start=Tue%2C%2016%20Jan%2015%3A14%3A29%20%2B0000*<br> Response:<br> ``` {\"accepted\": 300, \"delivered\": 100, \"complained\": 0, \"failed\": 50, \"bounced\": 150, \"queued\": 0} ``` <br><br> **Get stats for a period, grouped by channel**<br> Request:<br> *_/v4/stats/last_day/channel*<br> Response:<br> ``` {\"channel1\": {\"accepted\": 30, \"delivered\": 10, \"complained\": 0, \"failed\": 5, \"bounced\": 15, \"queued\": 0}, \"channel2\": {\"accepted\": 0, \"delivered\": 0, \"complained\": 0, \"failed\": 0, \"bounced\": 0, \"queued\": 0}} ``` <br><br> **Get stats for specific sending domain and channel (sender) and period, grouped by ISP**<br> Request:<br> *_/v4/stats/last_day/channel/marketing/domain/smtp.com/rcpt_isp*<br> Response:<br> ``` {\"yahoo\": {\"accepted\": 30, \"delivered\": 10, \"complained\": 0, \"failed\": 5, \"bounced\": 15, \"queued\": 0}, \"google\": {\"accepted\": 0, \"delivered\": 0, \"complained\": 0, \"failed\": 0, \"bounced\": 0, \"queued\": 0}} ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **string**| The starting time. Required if the &#x60;{duration}&#x60; path parameter is not specified. RFC 2822 or UNIX epoch format. | 
**duration** | **string**| A standardized shorthand for a known start/end bracket. Duration automatically supersedes start/end values provided as query string parameters. One of either the &#x60;{duration}&#x60; path parameter or the &#x60;start&#x60; query parameter must be specified.  &lt;table&gt;  &lt;tr&gt;&lt;th&gt;Value&lt;/th&gt;&lt;th&gt;Start&lt;/th&gt;&lt;th&gt;End&lt;/th&gt;&lt;th&gt;Slicable&lt;/th&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_24hrs&lt;/td&gt;&lt;td&gt;84,400 seconds ago&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_30days&lt;/td&gt;&lt;td&gt;18,144,000 seconds ago&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_7days&lt;/td&gt;&lt;td&gt;604,800 seconds ago&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_day&lt;/td&gt;&lt;td&gt;00:00:00 of the previous day&lt;/td&gt;&lt;td&gt;23:59:59 of the previous day&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_hour&lt;/td&gt;&lt;td&gt;00:00 of the previous hour&lt;/td&gt;&lt;td&gt;59:59 of the previous hour&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_month or mtd&lt;/td&gt;&lt;td&gt;1st day 00:00:00 of previous month&lt;/td&gt;&lt;td&gt;23:59:59 last day of previous month&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_week&lt;/td&gt;&lt;td&gt;Monday 00:00:00 of previous week&lt;/td&gt;&lt;td&gt;Sunday 23:59:59 of previous week&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;this_day&lt;/td&gt;&lt;td&gt;00:00:00 of current day&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;this_hour&lt;/td&gt;&lt;td&gt;00:00 of current clock hour&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;this_month&lt;/td&gt;&lt;td&gt;1st day 00:00:00 of current month&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;this_week&lt;/td&gt;&lt;td&gt;Monday 00:00:00 of current week&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;yes&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;last_year&lt;/td&gt;&lt;td&gt; Jan 1st 00:00:00 of previous year&lt;/td&gt;&lt;td&gt;Dec 31st 23:59:59 of previous year&lt;/td&gt;&lt;td&gt;no&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;this_year or ytd&lt;/td&gt;&lt;td&gt;Jan 1st  00:00:0 of current year&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;no&lt;/td&gt;&lt;/tr&gt;  &lt;tr&gt;&lt;td&gt;total&lt;/td&gt;&lt;td&gt;Account creation date&lt;/td&gt;&lt;td&gt;Now&lt;/td&gt;&lt;td&gt;no&lt;/td&gt;&lt;/tr&gt;  &lt;/table&gt;  | 
**slice** | **string**| A reducing method which can be applied to the requested duration. A final slice without an optional slice specifier will define a grouping.  Possible Values:  * &#x60;channel&#x60;: (optional) A given account&#39;s sender  * &#x60;domain&#x60;: (optional) Sending domain  * &#x60;rcpt_domain&#x60;: (optional) Recieving domain  * &#x60;rcpt_isp&#x60;: (optional) Receiving ISP     Slices can be chained in a meaningful way – for example:   &#x60;&#x60;&#x60;   /last_month/channel/marketing/domain/smtp.com/rcpt_domain?event&#x3D;complained   &#x60;&#x60;&#x60; would produce an aggregate of complaints for a current account’s channel (sender) called “marketing” which were:   * sent from the registered email domain “smtp.com”, and    * are grouped by receiving domains       The response would look something like:   &#x60;&#x60;&#x60;   {“google.com”: {“complained”: 5}, “yahoo.com”: {“complained”:1}, “aol.com”: {“complained”:1}}   &#x60;&#x60;&#x60;  | 
**sliceSpecifier** | **string**| slice value (smtp.com, sender1) | 
**groupBy** | **string**| Define a grouping:  * &#x60;channel&#x60; - optionally to be followed by a channel ID or name specifier  * &#x60;domain&#x60;  - optionally to be followed by a registered domain name  * &#x60;rcpt_domain&#x60; - optionally to be followed by a registered domain name  * &#x60;rcpt_isp&#x60; - optionally to be followed by a registered domain name  | 
**limit** | **int32**| Maximum number of items to return. | 
**offset** | **int32**| Number of items to skip before returning the results. | 
 **optional** | ***V4StatsDurationSliceSliceSpecifierGroupByGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4StatsDurationSliceSliceSpecifierGroupByGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **end** | **optional.String**| The ending time. If not specified, defaults to now. RFC 2822 or UNIX epoch format. | 
 **event** | **optional.String**| Array of any event names for which stats has been requested. | 

### Return type

[**StatsResponse**](StatsResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

