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

// CreateDkimKey struct for CreateDkimKey
type CreateDkimKey struct {
	// Request status.
	Status string            `json:"status,omitempty"`
	Data   CreateDkimKeyData `json:"data,omitempty"`
}