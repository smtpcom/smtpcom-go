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

// AccountData Account
type AccountData struct {
	// Status of the account
	Status string `json:"status,omitempty"`
	// First name of account owner
	FirstName string `json:"first_name,omitempty"`
	// Last name of account owner
	LastName string `json:"last_name,omitempty"`
	// Phone number of account owner
	Phone string `json:"phone,omitempty"`
	// Website of account owner
	Website string `json:"website,omitempty"`
	// Email address of account owner
	Email string `json:"email,omitempty"`
	// Account owner’s company name
	CompanyName string             `json:"company_name,omitempty"`
	Address     AccountDataAddress `json:"address,omitempty"`
	// Account's usage
	Usage int32 `json:"usage,omitempty"`
	// Account's creation time (in UTC)
	DateCreated int32 `json:"date created,omitempty"`
}
