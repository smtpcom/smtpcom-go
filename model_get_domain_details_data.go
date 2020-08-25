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

// GetDomainDetailsData Domains
type GetDomainDetailsData struct {
	// Registered domain name
	DomainName string `json:"domain_name,omitempty"`
	// Name of the DKIM selector for this domain
	Selector string `json:"selector,omitempty"`
	// Private key of the DKIM record
	PrivateKey string `json:"private_key,omitempty"`
	// Boolean flag indicating whether this DKIM is valid – i.e. private key matches the public key, identified by a selector for this domain
	IsValid bool `json:"is_valid,omitempty"`
}