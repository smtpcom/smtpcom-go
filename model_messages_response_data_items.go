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

// MessagesResponseDataItems struct for MessagesResponseDataItems
type MessagesResponseDataItems struct {
	// Unique message ID
	MsgId string `json:"msg_id,omitempty"`
	// Time at which the message was sent
	MsgTime int32 `json:"msg_time,omitempty"`
	// Name of the channel on which the message was sent
	Channel string `json:"channel,omitempty"`
	// Custom parameters and their value echoed back from `X-SMTPAPI` header's `unique_args` parameter
	SmtpVars map[string]interface{}      `json:"smtp_vars,omitempty"`
	MsgData  MessagesResponseDataMsgData `json:"msg_data,omitempty"`
	Details  MessagesResponseDataDetails `json:"details,omitempty"`
	Opens    MessagesResponseDataOpens   `json:"opens,omitempty"`
	Clicks   MessagesResponseDataClicks  `json:"clicks,omitempty"`
	Abuse    MessagesResponseDataAbuse   `json:"abuse,omitempty"`
	Unsubs   MessagesResponseDataUnsubs  `json:"unsubs,omitempty"`
}
