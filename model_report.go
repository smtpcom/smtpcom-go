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

// Report struct for Report
type Report struct {
	// Report frequency – one of:   * daily - every day at a specified hour   * weekly - Mondays at a specified hour   * monthly - 1st day of the month at a specified hour
	Frequency string `json:"frequency,omitempty"`
	// Unique report ID
	ReportId string `json:"report_id,omitempty"`
	// Preset of events returned in a report:    * hard_bounced   * failed   * delivered   * sent   * pending   * total   * abuse If not specified all events are included.
	Events string `json:"events,omitempty"`
	// Name of channel (sender). If not specified all channels will be reported
	Channel string `json:"channel,omitempty"`
	// The hour at which the report should be sent, values range from 0 to 23
	ReportTime string `json:"report_time,omitempty"`
	// Current status of a given on-demand report
	Status string `json:"status,omitempty"`
	// Human readable name of an on-demand report (auto-generated)
	Name string `json:"name,omitempty"`
	// The unique URL from which to download an on-demand report from
	Url string `json:"url,omitempty"`
	// Time when a given on-demand report has been requestedi. RFC 2822 or UNIX epoch format
	TimeReq int32 `json:"time_req,omitempty"`
	// Percentage of completion for an on-demand report
	Progress int32 `json:"progress,omitempty"`
}
