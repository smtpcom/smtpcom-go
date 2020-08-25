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

// ApiKey struct for ApiKey
type ApiKey struct {
	// Current status of the API key
	Status string `json:"status,omitempty"`
	// Description for API key
	Descrption string `json:"descrption,omitempty"`
	// The actual API key value
	Key string `json:"key,omitempty"`
	// Timestamp of when the API key was first created.
	DateCreated string `json:"date_created,omitempty"`
}