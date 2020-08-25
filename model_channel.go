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

// Channel struct for Channel
type Channel struct {
	// Request status.
	Status string      `json:"status,omitempty"`
	Data   ChannelData `json:"data,omitempty"`
}
