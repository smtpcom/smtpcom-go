# Go API client for smtp

SMTP.com Public API v4

## Installation

Install the client:

```bash
go get github.com/smtpcom/smtpcom-go
```

Import it locally
```golang
import smtp "github.com/smtpcom/smtpcom-go"
```

See https://github.com/smtpcom/smtpcom-go/blob/master/examples/mail/mail.go for quickstart.


## Documentation for API Endpoints

All URIs are relative to *https://api.smtp.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIKeysApi* | [**GetAPIKey**](docs/APIKeysApi.md#getapikey) | **Get** /v4/api_keys/{api_key} | Get API Key Details
*APIKeysApi* | [**V4ApiKeysApiKeyDelete**](docs/APIKeysApi.md#v4apikeysapikeydelete) | **Delete** /v4/api_keys/{api_key} | Delete an API Key
*APIKeysApi* | [**V4ApiKeysApiKeyPatch**](docs/APIKeysApi.md#v4apikeysapikeypatch) | **Patch** /v4/api_keys/{api_key} | Update API Key
*APIKeysApi* | [**V4ApiKeysGet**](docs/APIKeysApi.md#v4apikeysget) | **Get** /v4/api_keys/ | List All API Keys
*APIKeysApi* | [**V4ApiKeysPost**](docs/APIKeysApi.md#v4apikeyspost) | **Post** /v4/api_keys/ | Create a New API Key
*AccountApi* | [**V4AccountContactPatch**](docs/AccountApi.md#v4accountcontactpatch) | **Patch** /v4/account/contact | Update Account Details
*AccountApi* | [**V4AccountGet**](docs/AccountApi.md#v4accountget) | **Get** /v4/account/ | Get Account Details
*AlertsApi* | [**V4AlertsAlertIdDelete**](docs/AlertsApi.md#v4alertsalertiddelete) | **Delete** /v4/alerts/{alert_id} | Delete Alert
*AlertsApi* | [**V4AlertsAlertIdGet**](docs/AlertsApi.md#v4alertsalertidget) | **Get** /v4/alerts/{alert_id} | Get Alert Details
*AlertsApi* | [**V4AlertsAlertIdPatch**](docs/AlertsApi.md#v4alertsalertidpatch) | **Patch** /v4/alerts/{alert_id} | Update Alert Details
*AlertsApi* | [**V4AlertsGet**](docs/AlertsApi.md#v4alertsget) | **Get** /v4/alerts/ | List All Allerts
*AlertsApi* | [**V4AlertsPost**](docs/AlertsApi.md#v4alertspost) | **Post** /v4/alerts/ | Create New Alert
*CallbacksApi* | [**V4CallbacksDelete**](docs/CallbacksApi.md#v4callbacksdelete) | **Delete** /v4/callbacks/ | Delete All Callbacks
*CallbacksApi* | [**V4CallbacksEventDelete**](docs/CallbacksApi.md#v4callbackseventdelete) | **Delete** /v4/callbacks/{event} | Delete Callback
*CallbacksApi* | [**V4CallbacksEventGet**](docs/CallbacksApi.md#v4callbackseventget) | **Get** /v4/callbacks/{event} | Get Callback Details
*CallbacksApi* | [**V4CallbacksEventPatch**](docs/CallbacksApi.md#v4callbackseventpatch) | **Patch** /v4/callbacks/{event} | Update Callback Details
*CallbacksApi* | [**V4CallbacksEventPost**](docs/CallbacksApi.md#v4callbackseventpost) | **Post** /v4/callbacks/{event} | Create Callback
*CallbacksApi* | [**V4CallbacksGet**](docs/CallbacksApi.md#v4callbacksget) | **Get** /v4/callbacks/ | List All Callbacks
*CallbacksApi* | [**V4CallbacksLogGet**](docs/CallbacksApi.md#v4callbackslogget) | **Get** /v4/callbacks/log | View Callback Logs
*ChannelsApi* | [**GetSender**](docs/ChannelsApi.md#getsender) | **Get** /v4/channels/{name} | Get Channel Details
*ChannelsApi* | [**V4ChannelsGet**](docs/ChannelsApi.md#v4channelsget) | **Get** /v4/channels/ | Get All Channels
*ChannelsApi* | [**V4ChannelsNameDelete**](docs/ChannelsApi.md#v4channelsnamedelete) | **Delete** /v4/channels/{name} | Delete a Channel
*ChannelsApi* | [**V4ChannelsNamePatch**](docs/ChannelsApi.md#v4channelsnamepatch) | **Patch** /v4/channels/{name} | Update Channel Details
*ChannelsApi* | [**V4ChannelsPost**](docs/ChannelsApi.md#v4channelspost) | **Post** /v4/channels/ | Create a New Channel
*DKIMsApi* | [**V4DomainsDomainNameDelete**](docs/DKIMsApi.md#v4domainsdomainnamedelete) | **Delete** /v4/domains/{domain_name} | Delete Domain
*DKIMsApi* | [**V4DomainsDomainNameDkimKeysDelete**](docs/DKIMsApi.md#v4domainsdomainnamedkimkeysdelete) | **Delete** /v4/domains/{domain_name}/dkim_keys | Delete DKIM for Domain
*DKIMsApi* | [**V4DomainsDomainNameDkimKeysGet**](docs/DKIMsApi.md#v4domainsdomainnamedkimkeysget) | **Get** /v4/domains/{domain_name}/dkim_keys | Get DKIM for Domain
*DKIMsApi* | [**V4DomainsDomainNameDkimKeysPatch**](docs/DKIMsApi.md#v4domainsdomainnamedkimkeyspatch) | **Patch** /v4/domains/{domain_name}/dkim_keys | Update DKIM Key Details
*DKIMsApi* | [**V4DomainsDomainNameDkimKeysPost**](docs/DKIMsApi.md#v4domainsdomainnamedkimkeyspost) | **Post** /v4/domains/{domain_name}/dkim_keys | Add DKIM for Domain
*DKIMsApi* | [**V4DomainsDomainNameGet**](docs/DKIMsApi.md#v4domainsdomainnameget) | **Get** /v4/domains/{domain_name} | Get Domain Details
*DKIMsApi* | [**V4DomainsDomainNamePatch**](docs/DKIMsApi.md#v4domainsdomainnamepatch) | **Patch** /v4/domains/{domain_name} | Update Domain Details
*DKIMsApi* | [**V4DomainsGet**](docs/DKIMsApi.md#v4domainsget) | **Get** /v4/domains/ | Get All Registered Domains
*DKIMsApi* | [**V4DomainsPost**](docs/DKIMsApi.md#v4domainspost) | **Post** /v4/domains/ | Register a Domain
*MessagesApi* | [**V4MessagesGet**](docs/MessagesApi.md#v4messagesget) | **Get** /v4/messages | Get Detailed Messages Info
*MessagesApi* | [**V4MessagesMimePost**](docs/MessagesApi.md#v4messagesmimepost) | **Post** /v4/messages/mime | Send MIME Message
*MessagesApi* | [**V4MessagesPost**](docs/MessagesApi.md#v4messagespost) | **Post** /v4/messages | Send a Message
*ReportsApi* | [**V4ReportsGet**](docs/ReportsApi.md#v4reportsget) | **Get** /v4/reports/ | Get All Reports
*ReportsApi* | [**V4ReportsOndemandPost**](docs/ReportsApi.md#v4reportsondemandpost) | **Post** /v4/reports/ondemand | Create On-Demand Report
*ReportsApi* | [**V4ReportsPeriodicPost**](docs/ReportsApi.md#v4reportsperiodicpost) | **Post** /v4/reports/periodic | Create Periodic Report
*ReportsApi* | [**V4ReportsPeriodicReportIdDelete**](docs/ReportsApi.md#v4reportsperiodicreportiddelete) | **Delete** /v4/reports/periodic/{report_id} | Delete a Periodic Report
*ReportsApi* | [**V4ReportsPeriodicReportIdPatch**](docs/ReportsApi.md#v4reportsperiodicreportidpatch) | **Patch** /v4/reports/periodic/{report_id} | Update Periodic Report
*ReportsApi* | [**V4ReportsReportIdGet**](docs/ReportsApi.md#v4reportsreportidget) | **Get** /v4/reports/{report_id} | Get Report Details
*StatisticsApi* | [**V4StatsDurationSliceSliceSpecifierGroupByGet**](docs/StatisticsApi.md#v4statsdurationsliceslicespecifiergroupbyget) | **Get** /v4/stats/{duration}/{slice}/{slice_specifier}/{group_by} | Return event aggregates for the specified slices and duration. Slices can be chained.


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountData](docs/AccountData.md)
 - [AccountDataAddress](docs/AccountDataAddress.md)
 - [ApiKey](docs/ApiKey.md)
 - [BadRequestSchema](docs/BadRequestSchema.md)
 - [BadRequestSchemaData](docs/BadRequestSchemaData.md)
 - [Channel](docs/Channel.md)
 - [ChannelData](docs/ChannelData.md)
 - [Channels](docs/Channels.md)
 - [ChannelsData](docs/ChannelsData.md)
 - [ChannelsDataItems](docs/ChannelsDataItems.md)
 - [CreateCallbackResponse](docs/CreateCallbackResponse.md)
 - [CreateDkimKey](docs/CreateDkimKey.md)
 - [CreateDkimKeyData](docs/CreateDkimKeyData.md)
 - [CreateDomainResponse](docs/CreateDomainResponse.md)
 - [CreateDomainResponseData](docs/CreateDomainResponseData.md)
 - [DurationValue](docs/DurationValue.md)
 - [GetAlertDetails](docs/GetAlertDetails.md)
 - [GetAlertDetailsData](docs/GetAlertDetailsData.md)
 - [GetAlertResponse](docs/GetAlertResponse.md)
 - [GetAlertResponseData](docs/GetAlertResponseData.md)
 - [GetAlertResponseDataItems](docs/GetAlertResponseDataItems.md)
 - [GetApiKeys](docs/GetApiKeys.md)
 - [GetApiKeysData](docs/GetApiKeysData.md)
 - [GetApiKeysDataItems](docs/GetApiKeysDataItems.md)
 - [GetCallbackDetails](docs/GetCallbackDetails.md)
 - [GetCallbackDetailsData](docs/GetCallbackDetailsData.md)
 - [GetCallbackLogs](docs/GetCallbackLogs.md)
 - [GetCallbackLogsData](docs/GetCallbackLogsData.md)
 - [GetCallbackLogsDataItems](docs/GetCallbackLogsDataItems.md)
 - [GetCallbackResponse](docs/GetCallbackResponse.md)
 - [GetCallbackResponseData](docs/GetCallbackResponseData.md)
 - [GetCallbackResponseDataItems](docs/GetCallbackResponseDataItems.md)
 - [GetDomainDetails](docs/GetDomainDetails.md)
 - [GetDomainDetailsData](docs/GetDomainDetailsData.md)
 - [GetDomainDetailsResponse](docs/GetDomainDetailsResponse.md)
 - [GetDomainDetailsResponseData](docs/GetDomainDetailsResponseData.md)
 - [GetDomainsResponse](docs/GetDomainsResponse.md)
 - [GetDomainsResponseData](docs/GetDomainsResponseData.md)
 - [GetDomainsResponseDataItems](docs/GetDomainsResponseDataItems.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InvalidKey](docs/InvalidKey.md)
 - [InvalidKeyData](docs/InvalidKeyData.md)
 - [InvalidKeyDataErrors](docs/InvalidKeyDataErrors.md)
 - [MessagesResponse](docs/MessagesResponse.md)
 - [MessagesResponseData](docs/MessagesResponseData.md)
 - [MessagesResponseDataAbuse](docs/MessagesResponseDataAbuse.md)
 - [MessagesResponseDataAbuseComplaints](docs/MessagesResponseDataAbuseComplaints.md)
 - [MessagesResponseDataClicks](docs/MessagesResponseDataClicks.md)
 - [MessagesResponseDataClicksItems](docs/MessagesResponseDataClicksItems.md)
 - [MessagesResponseDataDetails](docs/MessagesResponseDataDetails.md)
 - [MessagesResponseDataDetailsDelivery](docs/MessagesResponseDataDetailsDelivery.md)
 - [MessagesResponseDataItems](docs/MessagesResponseDataItems.md)
 - [MessagesResponseDataMsgData](docs/MessagesResponseDataMsgData.md)
 - [MessagesResponseDataOpens](docs/MessagesResponseDataOpens.md)
 - [MessagesResponseDataOpensItems](docs/MessagesResponseDataOpensItems.md)
 - [MessagesResponseDataUnsubs](docs/MessagesResponseDataUnsubs.md)
 - [MessagesResponseDataUnsubsItems](docs/MessagesResponseDataUnsubsItems.md)
 - [MimeResponse](docs/MimeResponse.md)
 - [PostMessageResponse](docs/PostMessageResponse.md)
 - [PostMessageResponseData](docs/PostMessageResponseData.md)
 - [Report](docs/Report.md)
 - [Reports](docs/Reports.md)
 - [ReportsData](docs/ReportsData.md)
 - [ReportsDataOndemand](docs/ReportsDataOndemand.md)
 - [ReportsDataPeriodic](docs/ReportsDataPeriodic.md)
 - [SliceValue](docs/SliceValue.md)
 - [StatsResponse](docs/StatsResponse.md)
 - [StatsResponseData](docs/StatsResponseData.md)
 - [StatsResponseDataItems](docs/StatsResponseDataItems.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [UpdateAccountResponse](docs/UpdateAccountResponse.md)
 - [UpdateAccountResponseData](docs/UpdateAccountResponseData.md)
 - [V4MessagesBody](docs/V4MessagesBody.md)
 - [V4MessagesBodyAttachments](docs/V4MessagesBodyAttachments.md)
 - [V4MessagesBodyParts](docs/V4MessagesBodyParts.md)
 - [V4MessagesMimeRecipients](docs/V4MessagesMimeRecipients.md)
 - [V4MessagesOriginator](docs/V4MessagesOriginator.md)
 - [V4MessagesRecipients](docs/V4MessagesRecipients.md)
 - [V4MessagesRecipientsTo](docs/V4MessagesRecipientsTo.md)


## Documentation For Authorization



## apiID

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## apiKey

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author

support@smtp.com

