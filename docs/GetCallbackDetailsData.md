# GetCallbackDetailsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Medium** | **string** | Medium by which the callback data is sent. Possible values are one of:   * http   * aws  | [optional] 
**Event** | **string** | Event for which the callback has been created. Valid types are:  * delivered -  message delivered * failed - message bounced * complained - complaint received * bounceback - bounce back notification received * received - message received by our system * queued - message in queue (transient) * hard_bounced - hard bounce received * opened - message opened * clicked - URL in message clicked * unsubscribed - unsubscribe received  | [optional] 
**Channel** | **string** | Name of the channel for which the callback has been created | [optional] 
**Address** | **string** | Address to which the callback data is sent. This will be either a URL for http-based callbacks or the Amazon SQS queue name for SQS-based callbacks | [optional] 
**AwsData** | **string** | Amazon SQS settings | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


