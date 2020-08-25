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

// ReportsApiService ReportsApi service
type ReportsApiService service

/*
V4ReportsGet Get All Reports
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Reports
*/
func (a *ReportsApiService) V4ReportsGet(ctx _context.Context) (Reports, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Reports
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/"
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

// V4ReportsOndemandPostOpts Optional parameters for the method 'V4ReportsOndemandPost'
type V4ReportsOndemandPostOpts struct {
	Channel    optional.String
	Type_      optional.String
	End        optional.String
	Domain     optional.String
	RcptDomain optional.String
	Events     optional.String
	Columns    optional.String
}

/*
V4ReportsOndemandPost Create On-Demand Report
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param start Start date/time of the report in RFC 2822 or UNIX epoch format
 * @param optional nil or *V4ReportsOndemandPostOpts - Optional Parameters:
 * @param "Channel" (optional.String) -  Name of the channel for which a given report has been defined
 * @param "Type_" (optional.String) -  Type or report format. If not specified defaults to “csv” - currently the only supported type
 * @param "End" (optional.String) -  End date/time of the report in RFC 2822 or UNIX epoch format (default - now)
 * @param "Domain" (optional.String) -  Filter by the “from” domain of emails
 * @param "RcptDomain" (optional.String) -  Filter by the “to” domain of emails
 * @param "Events" (optional.String) -  Filter by event type. Valid event are:  * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints
 * @param "Columns" (optional.String) -  Array of columns to be specified in the report. These can differ based on any specified event type filter.   Possible column values for all reports are: * `message_id` - Unique message ID * from - From Address * to - Recipient Address * time_rcv - Date Received in RFC 2822 or UNIX epoch format * time_snt - Date Delivered in RFC 2822 or UNIX epoch format * channel - Name of a channel  Additional column values for message reports (hard_bounced, failed, delivered, total) include: * status - Status of delivery * code - SMTP Response Code * message - SMTP Response message * tries - Amount of re-tries (deferred states before final)  Additional column values for spam reports include:  * report_time - Date when an abuse complaint has been reported, RFC 2822 or UNIX epoch format * subject - Email Subject  Additional column values for pending reports include:  * status - Current email status  If not specified all relevant columns are included.
@return StatusResponse
*/
func (a *ReportsApiService) V4ReportsOndemandPost(ctx _context.Context, start int32, localVarOptionals *V4ReportsOndemandPostOpts) (StatusResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StatusResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/ondemand"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarQueryParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	localVarQueryParams.Add("start", parameterToString(start, ""))
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		localVarQueryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RcptDomain.IsSet() {
		localVarQueryParams.Add("rcpt_domain", parameterToString(localVarOptionals.RcptDomain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Events.IsSet() {
		localVarQueryParams.Add("events", parameterToString(localVarOptionals.Events.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Columns.IsSet() {
		localVarQueryParams.Add("columns", parameterToString(localVarOptionals.Columns.Value(), ""))
	}
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

// V4ReportsPeriodicPostOpts Optional parameters for the method 'V4ReportsPeriodicPost'
type V4ReportsPeriodicPostOpts struct {
	Channel      optional.String
	NotifyMethod optional.String
	NotifyDest   optional.String
	Domain       optional.String
	RcptDomain   optional.String
	Events       optional.String
	Columns      optional.String
}

/*
V4ReportsPeriodicPost Create Periodic Report
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param frequency Report frequency – one of:   * daily - every day at a specified hour   * weekly  - Mondays at a specified hour   * monthly - 1st day of the month at a specified hour
 * @param reportTime The hour at which the report should be sent. Value values range from 0 to 23
 * @param optional nil or *V4ReportsPeriodicPostOpts - Optional Parameters:
 * @param "Channel" (optional.String) -  Name of the channel for which a given report has been defined
 * @param "NotifyMethod" (optional.String) -  Notification method to be used when report is completed and can be downloaded
 * @param "NotifyDest" (optional.String) -  A valid URL which will accept the report completion notification. The payload will be ```   {\"message\": \"success\", \"id\": string} ``` where `id` is a Unique report ID
 * @param "Domain" (optional.String) -  Filter by the “From” domain of emails
 * @param "RcptDomain" (optional.String) -  Filter by the “To” domain of emails
 * @param "Events" (optional.String) -  Filter by event type. Valid event are:  * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints
 * @param "Columns" (optional.String) -  Array of columns to be specified in the report. These can differ based on any specified event type filter.   Possible column values are: * `message_id` - Unique message ID * from - From Address * to - Recipient Address * time_rcv - Date Received in RFC 2822 or UNIX epoch format * time_snt - Date Delivered in RFC 2822 or UNIX epoch format * channel - Name of a channel  Additional column values for message reports (hard_bounced, failed, delivered, total) include: * status - Status of delivery * code - SMTP Response Code * message - SMTP Response message * tries - Amount of re-tries (deferred states before final)  Additional column values for spam reports include:  * report_time - Date when an abuse complaint has been reported, RFC 2822 or UNIX epoch format * subject - Email Subject  Additional column values for pending reports include:  * status - Current email status  If not specified all relevant columns are included.
@return StatusResponse
*/
func (a *ReportsApiService) V4ReportsPeriodicPost(ctx _context.Context, frequency string, reportTime int32, localVarOptionals *V4ReportsPeriodicPostOpts) (StatusResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StatusResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/periodic"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("frequency", parameterToString(frequency, ""))
	localVarQueryParams.Add("report_time", parameterToString(reportTime, ""))
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarQueryParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NotifyMethod.IsSet() {
		localVarQueryParams.Add("notify_method", parameterToString(localVarOptionals.NotifyMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NotifyDest.IsSet() {
		localVarQueryParams.Add("notify_dest", parameterToString(localVarOptionals.NotifyDest.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		localVarQueryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RcptDomain.IsSet() {
		localVarQueryParams.Add("rcpt_domain", parameterToString(localVarOptionals.RcptDomain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Events.IsSet() {
		localVarQueryParams.Add("events", parameterToString(localVarOptionals.Events.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Columns.IsSet() {
		localVarQueryParams.Add("columns", parameterToString(localVarOptionals.Columns.Value(), ""))
	}
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

/*
V4ReportsPeriodicReportIdDelete Delete a Periodic Report
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param reportId Id of a given report
@return StatusResponse
*/
func (a *ReportsApiService) V4ReportsPeriodicReportIdDelete(ctx _context.Context, reportId string) (StatusResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StatusResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/periodic/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", _neturl.PathEscape(parameterToString(reportId, "")), -1)

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

// V4ReportsPeriodicReportIdPatchOpts Optional parameters for the method 'V4ReportsPeriodicReportIdPatch'
type V4ReportsPeriodicReportIdPatchOpts struct {
	Channel optional.String
	Events  optional.String
}

/*
V4ReportsPeriodicReportIdPatch Update Periodic Report
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param reportId Id of the report to be updated
 * @param frequency Report frequency – one of:  * daily - every day at a specified hour  * weekly  - Mondays at a specified hour  * monthly - first day of the month at a specified hour.
 * @param reportTime The hour at which the report should be sent. Value values range from 0 to 23
 * @param optional nil or *V4ReportsPeriodicReportIdPatchOpts - Optional Parameters:
 * @param "Channel" (optional.String) -  Name of channel (sender). If not specified all channels will be reported
 * @param "Events" (optional.String) -  Filter by event type. Valid event are: * hard_bounced - just hard bounces * failed - all failed messages, i.e. hard_bounced + the rest of failed * delivered - delivered messages * sent - delivered+failed (default events value) * pending - pending messages * total - all messages, i.e. sent+pending * abuse - spam complaints  If not specified all events are included.
@return StatusResponse
*/
func (a *ReportsApiService) V4ReportsPeriodicReportIdPatch(ctx _context.Context, reportId string, frequency string, reportTime int32, localVarOptionals *V4ReportsPeriodicReportIdPatchOpts) (StatusResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StatusResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/periodic/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", _neturl.PathEscape(parameterToString(reportId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("frequency", parameterToString(frequency, ""))
	localVarQueryParams.Add("report_time", parameterToString(reportTime, ""))
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarQueryParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Events.IsSet() {
		localVarQueryParams.Add("events", parameterToString(localVarOptionals.Events.Value(), ""))
	}
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

/*
V4ReportsReportIdGet Get Report Details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param reportId ID of a given report
@return Report
*/
func (a *ReportsApiService) V4ReportsReportIdGet(ctx _context.Context, reportId string) (Report, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Report
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/reports/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", _neturl.PathEscape(parameterToString(reportId, "")), -1)

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
