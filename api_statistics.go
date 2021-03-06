/*
 * SMTP Public API overview
 *
 * SMTP.com Public API v4
 *
 * API version: 4.0.0
 * Contact: support@smtp.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smtp

import (
	_bytes "bytes"
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// StatisticsApiService StatisticsApi service
type StatisticsApiService service

// V4StatsDurationSliceSliceSpecifierGroupByGetOpts Optional parameters for the method 'V4StatsDurationSliceSliceSpecifierGroupByGet'
type V4StatsDurationSliceSliceSpecifierGroupByGetOpts struct {
	End   optional.String
	Event optional.String
}

/*
V4StatsDurationSliceSliceSpecifierGroupByGet Return event aggregates for the specified slices and duration. Slices can be chained.
**Get stats for a period**&lt;br&gt; Request:&lt;br&gt; *_/v4/stats/last_day*&lt;br&gt; *_/v4/stats?start&#x3D;Tue%2C%2016%20Jan%2015%3A14%3A29%20%2B0000*&lt;br&gt; Response:&lt;br&gt; &#x60;&#x60;&#x60; {\&quot;accepted\&quot;: 300, \&quot;delivered\&quot;: 100, \&quot;complained\&quot;: 0, \&quot;failed\&quot;: 50, \&quot;bounced\&quot;: 150, \&quot;queued\&quot;: 0} &#x60;&#x60;&#x60; &lt;br&gt;&lt;br&gt; **Get stats for a period, grouped by channel**&lt;br&gt; Request:&lt;br&gt; *_/v4/stats/last_day/channel*&lt;br&gt; Response:&lt;br&gt; &#x60;&#x60;&#x60; {\&quot;channel1\&quot;: {\&quot;accepted\&quot;: 30, \&quot;delivered\&quot;: 10, \&quot;complained\&quot;: 0, \&quot;failed\&quot;: 5, \&quot;bounced\&quot;: 15, \&quot;queued\&quot;: 0}, \&quot;channel2\&quot;: {\&quot;accepted\&quot;: 0, \&quot;delivered\&quot;: 0, \&quot;complained\&quot;: 0, \&quot;failed\&quot;: 0, \&quot;bounced\&quot;: 0, \&quot;queued\&quot;: 0}} &#x60;&#x60;&#x60; &lt;br&gt;&lt;br&gt; **Get stats for specific sending domain and channel (sender) and period, grouped by ISP**&lt;br&gt; Request:&lt;br&gt; *_/v4/stats/last_day/channel/marketing/domain/smtp.com/rcpt_isp*&lt;br&gt; Response:&lt;br&gt; &#x60;&#x60;&#x60; {\&quot;yahoo\&quot;: {\&quot;accepted\&quot;: 30, \&quot;delivered\&quot;: 10, \&quot;complained\&quot;: 0, \&quot;failed\&quot;: 5, \&quot;bounced\&quot;: 15, \&quot;queued\&quot;: 0}, \&quot;google\&quot;: {\&quot;accepted\&quot;: 0, \&quot;delivered\&quot;: 0, \&quot;complained\&quot;: 0, \&quot;failed\&quot;: 0, \&quot;bounced\&quot;: 0, \&quot;queued\&quot;: 0}} &#x60;&#x60;&#x60;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param start The starting time. Required if the `{duration}` path parameter is not specified. RFC 2822 or UNIX epoch format.
 * @param duration A standardized shorthand for a known start/end bracket. Duration automatically supersedes start/end values provided as query string parameters. One of either the `{duration}` path parameter or the `start` query parameter must be specified.  <table>  <tr><th>Value</th><th>Start</th><th>End</th><th>Slicable</th></tr>  <tr><td>last_24hrs</td><td>84,400 seconds ago</td><td>Now</td><td>yes</td></tr>  <tr><td>last_30days</td><td>18,144,000 seconds ago</td><td>Now</td><td>yes</td></tr>  <tr><td>last_7days</td><td>604,800 seconds ago</td><td>Now</td><td>yes</td></tr>  <tr><td>last_day</td><td>00:00:00 of the previous day</td><td>23:59:59 of the previous day</td><td>yes</td></tr>  <tr><td>last_hour</td><td>00:00 of the previous hour</td><td>59:59 of the previous hour</td><td>yes</td></tr>  <tr><td>last_month or mtd</td><td>1st day 00:00:00 of previous month</td><td>23:59:59 last day of previous month</td><td>yes</td></tr>  <tr><td>last_week</td><td>Monday 00:00:00 of previous week</td><td>Sunday 23:59:59 of previous week</td><td>yes</td></tr>  <tr><td>this_day</td><td>00:00:00 of current day</td><td>Now</td><td>yes</td></tr>  <tr><td>this_hour</td><td>00:00 of current clock hour</td><td>Now</td><td>yes</td></tr>  <tr><td>this_month</td><td>1st day 00:00:00 of current month</td><td>Now</td><td>yes</td></tr>  <tr><td>this_week</td><td>Monday 00:00:00 of current week</td><td>Now</td><td>yes</td></tr>  <tr><td>last_year</td><td> Jan 1st 00:00:00 of previous year</td><td>Dec 31st 23:59:59 of previous year</td><td>no</td></tr>  <tr><td>this_year or ytd</td><td>Jan 1st  00:00:0 of current year</td><td>Now</td><td>no</td></tr>  <tr><td>total</td><td>Account creation date</td><td>Now</td><td>no</td></tr>  </table>
 * @param slice A reducing method which can be applied to the requested duration. A final slice without an optional slice specifier will define a grouping.  Possible Values:  * `channel`: (optional) A given account's sender  * `domain`: (optional) Sending domain  * `rcpt_domain`: (optional) Recieving domain  * `rcpt_isp`: (optional) Receiving ISP     Slices can be chained in a meaningful way – for example:   ```   /last_month/channel/marketing/domain/smtp.com/rcpt_domain?event=complained   ``` would produce an aggregate of complaints for a current account’s channel (sender) called “marketing” which were:   * sent from the registered email domain “smtp.com”, and    * are grouped by receiving domains       The response would look something like:   ```   {“google.com”: {“complained”: 5}, “yahoo.com”: {“complained”:1}, “aol.com”: {“complained”:1}}   ```
 * @param sliceSpecifier slice value (smtp.com, sender1)
 * @param groupBy Define a grouping:  * `channel` - optionally to be followed by a channel ID or name specifier  * `domain`  - optionally to be followed by a registered domain name  * `rcpt_domain` - optionally to be followed by a registered domain name  * `rcpt_isp` - optionally to be followed by a registered domain name
 * @param limit Maximum number of items to return.
 * @param offset Number of items to skip before returning the results.
 * @param optional nil or *V4StatsDurationSliceSliceSpecifierGroupByGetOpts - Optional Parameters:
 * @param "End" (optional.String) -  The ending time. If not specified, defaults to now. RFC 2822 or UNIX epoch format.
 * @param "Event" (optional.String) -  Array of any event names for which stats has been requested.
@return StatsResponse
*/
func (a *StatisticsApiService) V4StatsDurationSliceSliceSpecifierGroupByGet(ctx _context.Context, start string, duration string, slice string, sliceSpecifier string, groupBy string, limit int32, offset int32, localVarOptionals *V4StatsDurationSliceSliceSpecifierGroupByGetOpts) (StatsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StatsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/stats/{duration}/{slice}/{slice_specifier}/{group_by}"
	localVarPath = strings.Replace(localVarPath, "{"+"duration"+"}", _neturl.PathEscape(parameterToString(duration, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"slice"+"}", _neturl.PathEscape(parameterToString(slice, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"slice_specifier"+"}", _neturl.PathEscape(parameterToString(sliceSpecifier, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"group_by"+"}", _neturl.PathEscape(parameterToString(groupBy, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if limit > 1000 {
		return localVarReturnValue, nil, reportError("limit must be less than 1000")
	}
	if offset < 0 {
		return localVarReturnValue, nil, reportError("offset must be greater than 0")
	}

	localVarQueryParams.Add("start", parameterToString(start, ""))
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Event.IsSet() {
		localVarQueryParams.Add("event", parameterToString(localVarOptionals.Event.Value(), ""))
	}
	localVarQueryParams.Add("limit", parameterToString(limit, ""))
	localVarQueryParams.Add("offset", parameterToString(offset, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-SMTPCOM-API"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InvalidKey
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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
