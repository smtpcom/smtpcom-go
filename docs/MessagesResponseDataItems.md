# MessagesResponseDataItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | **string** | Unique message ID | [optional] 
**MsgTime** | **int32** | Time at which the message was sent | [optional] 
**Channel** | **string** | Name of the channel on which the message was sent | [optional] 
**SmtpVars** | **map[string]interface{}** | Custom parameters and their value echoed back from &#x60;X-SMTPAPI&#x60; header&#39;s &#x60;unique_args&#x60; parameter | [optional] 
**MsgData** | [**MessagesResponseDataMsgData**](MessagesResponse_data_msg_data.md) |  | [optional] 
**Details** | [**MessagesResponseDataDetails**](MessagesResponse_data_details.md) |  | [optional] 
**Opens** | [**MessagesResponseDataOpens**](MessagesResponse_data_opens.md) |  | [optional] 
**Clicks** | [**MessagesResponseDataClicks**](MessagesResponse_data_clicks.md) |  | [optional] 
**Abuse** | [**MessagesResponseDataAbuse**](MessagesResponse_data_abuse.md) |  | [optional] 
**Unsubs** | [**MessagesResponseDataUnsubs**](MessagesResponse_data_unsubs.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


