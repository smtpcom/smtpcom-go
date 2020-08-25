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

// GetDomainDetails struct for GetDomainDetails
type GetDomainDetails struct {
	// Request status.
	Status string               `json:"status,omitempty"`
	Data   GetDomainDetailsData `json:"data,omitempty"`
}
