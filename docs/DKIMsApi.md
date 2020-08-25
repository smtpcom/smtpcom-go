# \DKIMsApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4DomainsDomainNameDelete**](DKIMsApi.md#V4DomainsDomainNameDelete) | **Delete** /v4/domains/{domain_name} | Delete Domain
[**V4DomainsDomainNameDkimKeysDelete**](DKIMsApi.md#V4DomainsDomainNameDkimKeysDelete) | **Delete** /v4/domains/{domain_name}/dkim_keys | Delete DKIM for Domain
[**V4DomainsDomainNameDkimKeysGet**](DKIMsApi.md#V4DomainsDomainNameDkimKeysGet) | **Get** /v4/domains/{domain_name}/dkim_keys | Get DKIM for Domain
[**V4DomainsDomainNameDkimKeysPatch**](DKIMsApi.md#V4DomainsDomainNameDkimKeysPatch) | **Patch** /v4/domains/{domain_name}/dkim_keys | Update DKIM Key Details
[**V4DomainsDomainNameDkimKeysPost**](DKIMsApi.md#V4DomainsDomainNameDkimKeysPost) | **Post** /v4/domains/{domain_name}/dkim_keys | Add DKIM for Domain
[**V4DomainsDomainNameGet**](DKIMsApi.md#V4DomainsDomainNameGet) | **Get** /v4/domains/{domain_name} | Get Domain Details
[**V4DomainsDomainNamePatch**](DKIMsApi.md#V4DomainsDomainNamePatch) | **Patch** /v4/domains/{domain_name} | Update Domain Details
[**V4DomainsGet**](DKIMsApi.md#V4DomainsGet) | **Get** /v4/domains/ | Get All Registered Domains
[**V4DomainsPost**](DKIMsApi.md#V4DomainsPost) | **Post** /v4/domains/ | Register a Domain



## V4DomainsDomainNameDelete

> StatusResponse V4DomainsDomainNameDelete(ctx, domainName)

Delete Domain

Deletes a domain that was previously registered by the current account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNameDkimKeysDelete

> StatusResponse V4DomainsDomainNameDkimKeysDelete(ctx, domainName)

Delete DKIM for Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNameDkimKeysGet

> GetDomainDetails V4DomainsDomainNameDkimKeysGet(ctx, domainName)

Get DKIM for Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 

### Return type

[**GetDomainDetails**](GetDomainDetails.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNameDkimKeysPatch

> StatusResponse V4DomainsDomainNameDkimKeysPatch(ctx, domainName, selector, privateKey)

Update DKIM Key Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 
**selector** | **string**| Name of DKIM selector for this domain | 
**privateKey** | **string**| Private key of the DKIM record | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNameDkimKeysPost

> CreateDkimKey V4DomainsDomainNameDkimKeysPost(ctx, domainName, selector, privateKey)

Add DKIM for Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 
**selector** | **string**| Name of DKIM selector for this domain | 
**privateKey** | **string**| Private key of the DKIM record | 

### Return type

[**CreateDkimKey**](CreateDkimKey.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNameGet

> GetDomainDetailsResponse V4DomainsDomainNameGet(ctx, domainName)

Get Domain Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 

### Return type

[**GetDomainDetailsResponse**](GetDomainDetailsResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsDomainNamePatch

> StatusResponse V4DomainsDomainNamePatch(ctx, domainName, enabled)

Update Domain Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name of interest | 
**enabled** | **bool**| Whether the domain is enabled | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsGet

> GetDomainsResponse V4DomainsGet(ctx, )

Get All Registered Domains

Returns all domains registered by this account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetDomainsResponse**](GetDomainsResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4DomainsPost

> CreateDomainResponse V4DomainsPost(ctx, domainName)

Register a Domain

Add a new domain to the list of account's registered domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**| Domain name to add | 

### Return type

[**CreateDomainResponse**](CreateDomainResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

