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

// CreateCallbackResponse struct for CreateCallbackResponse
type CreateCallbackResponse struct {
	// Request status.
	Status string `json:"status,omitempty"`
	// Data
	Data map[string]interface{} `json:"data,omitempty"`
}