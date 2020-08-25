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

// GetCallbackLogs struct for GetCallbackLogs
type GetCallbackLogs struct {
	// Request status.
	Status string              `json:"status,omitempty"`
	Data   GetCallbackLogsData `json:"data,omitempty"`
}
