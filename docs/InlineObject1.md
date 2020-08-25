# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mime** | **string** | A completely prepared full MIME container of the email, compliant with RFC 2045, RFC 2046, RFC 2047, RFC 4288, RFC 4289 and RFC 2049. No validation will be performed during API submission and it will be attempted to be delivered as is. Any errors while processing and delivering this email will be available only via callbacks.   | [optional] 
**Channel** | **string** | Name of the channel through which the email will be sent. | [optional] 
**Recipients** | [**V4MessagesMimeRecipients**](_v4_messages_mime_recipients.md) |  | [optional] 
**Originator** | [**V4MessagesOriginator**](_v4_messages_originator.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


