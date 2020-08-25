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

// InvalidKey struct for InvalidKey
type InvalidKey struct {
	// Request status.
	Status string         `json:"status,omitempty"`
	Data   InvalidKeyData `json:"data,omitempty"`
}