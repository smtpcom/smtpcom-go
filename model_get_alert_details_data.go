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

// GetAlertDetailsData Alert
type GetAlertDetailsData struct {
	// An alert’s type. Currently only “monthly_quota” is supported
	Type string `json:"type,omitempty"`
	// A number which represents a percentage of an account’s monthly quota. Must be decimal between 0 and 1
	Threshold string `json:"threshold,omitempty"`
	// alert ID
	AlertId string `json:"alert_id,omitempty"`
}