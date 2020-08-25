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

// UpdateAccountResponse struct for UpdateAccountResponse
type UpdateAccountResponse struct {
	// Request status.
	Status string                    `json:"status,omitempty"`
	Data   UpdateAccountResponseData `json:"data,omitempty"`
}
