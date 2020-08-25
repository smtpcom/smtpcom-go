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

// GetCallbackDetailsData data object
type GetCallbackDetailsData struct {
	// Medium by which the callback data is sent. Possible values are one of:   * http   * aws
	Medium string `json:"medium,omitempty"`
	// Event for which the callback has been created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received
	Event string `json:"event,omitempty"`
	// Name of the channel for which the callback has been created
	Channel string `json:"channel,omitempty"`
	// Address to which the callback data is sent. This will be either a URL for http-based callbacks or the Amazon SQS queue name for SQS-based callbacks
	Address string `json:"address,omitempty"`
	// Amazon SQS settings
	AwsData string `json:"aws_data,omitempty"`
}