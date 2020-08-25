# \AccountApi

All URIs are relative to *https://api.smtp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4AccountContactPatch**](AccountApi.md#V4AccountContactPatch) | **Patch** /v4/account/contact | Update Account Details
[**V4AccountGet**](AccountApi.md#V4AccountGet) | **Get** /v4/account/ | Get Account Details



## V4AccountContactPatch

> UpdateAccountResponse V4AccountContactPatch(ctx, optional)

Update Account Details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V4AccountContactPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V4AccountContactPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstName** | **optional.String**| First name of account owner | 
 **lastName** | **optional.String**| Last name of account owner | 
 **email** | **optional.String**| Email address of account owner | 
 **companyName** | **optional.String**| Account owner’s company name | 
 **phone** | **optional.String**| Phone number of account owner | 
 **website** | **optional.String**| Website of account owner | 
 **addressStreet** | **optional.String**| Account owner’s street address | 
 **addressCity** | **optional.String**| Account owner’s city | 
 **addressState** | **optional.String**| Account owner’s state | 
 **addressCountry** | **optional.String**| Account owner’s country | 

### Return type

[**UpdateAccountResponse**](UpdateAccountResponse.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4AccountGet

> Account V4AccountGet(ctx, )

Get Account Details

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Account**](Account.md)

### Authorization

[apiID](../README.md#apiID), [apiKey](../README.md#apiKey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

