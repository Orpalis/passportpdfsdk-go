# \PassportPDFApplicationManagerApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PassportPDFApplicationManagerGetApplicationDownloadLink**](PassportPDFApplicationManagerApi.md#PassportPDFApplicationManagerGetApplicationDownloadLink) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationDownloadLink | 
[**PassportPDFApplicationManagerGetApplicationLatestVersion**](PassportPDFApplicationManagerApi.md#PassportPDFApplicationManagerGetApplicationLatestVersion) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationLatestVersion | 
[**PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion**](PassportPDFApplicationManagerApi.md#PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion | 
[**PassportPDFApplicationManagerGetMaxClientThreads**](PassportPDFApplicationManagerApi.md#PassportPDFApplicationManagerGetMaxClientThreads) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetMaxClientThreads | Gets the maximum number of threads to be used simultaneously by a client application.



## PassportPDFApplicationManagerGetApplicationDownloadLink

> StringResponse PassportPDFApplicationManagerGetApplicationDownloadLink(ctx, applicationId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PassportPDFApplicationManagerGetApplicationLatestVersion

> StringResponse PassportPDFApplicationManagerGetApplicationLatestVersion(ctx, applicationId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion

> StringResponse PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion(ctx, applicationId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PassportPDFApplicationManagerGetMaxClientThreads

> IntegerResponse PassportPDFApplicationManagerGetMaxClientThreads(ctx, applicationId)

Gets the maximum number of threads to be used simultaneously by a client application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 

### Return type

[**IntegerResponse**](IntegerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

