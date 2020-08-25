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

// GetDomainsResponseDataItems struct for GetDomainsResponseDataItems
type GetDomainsResponseDataItems struct {
	// Registered domain name
	DomainName string `json:"domain_name,omitempty"`
	// Flag indicating whether the domain is enabled for the given account
	Enabled bool `json:"enabled,omitempty"`
}