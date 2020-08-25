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

// GetDomainsResponse struct for GetDomainsResponse
type GetDomainsResponse struct {
	// Request status.
	Status string                 `json:"status,omitempty"`
	Data   GetDomainsResponseData `json:"data,omitempty"`
}
