# \DocumentApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentClose**](DocumentApi.md#DocumentClose) | **Post** /api/document/DocumentClose | Closes a previously uploaded document.
[**DocumentGetPreview**](DocumentApi.md#DocumentGetPreview) | **Post** /api/document/DocumentGetPreview | Gets the format, the page count and a thumbnail of a previously uploaded document.
[**DocumentLoad**](DocumentApi.md#DocumentLoad) | **Post** /api/document/DocumentLoad | Loads the provided document file.
[**DocumentLoadFromURI**](DocumentApi.md#DocumentLoadFromURI) | **Post** /api/document/DocumentLoadFromURI | Loads the provided document file from an URI.
[**DocumentLoadMultipart**](DocumentApi.md#DocumentLoadMultipart) | **Post** /api/document/DocumentLoadMultipart | Loads the provided document file using Multipart Upload.



## DocumentClose

> DocumentCloseResponse DocumentClose(ctx, documentCloseParameters)

Closes a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentCloseParameters** | [**DocumentCloseParameters**](DocumentCloseParameters.md)| A DocumentCloseParameters object specifying the parameters of the action. | 

### Return type

[**DocumentCloseResponse**](DocumentCloseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentGetPreview

> DocumentGetPreviewResponse DocumentGetPreview(ctx, getDocumentPreviewParameters)

Gets the format, the page count and a thumbnail of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getDocumentPreviewParameters** | [**GetDocumentPreviewParameters**](GetDocumentPreviewParameters.md)| A GetDocumentPreviewParameters object specifying the parameters of the action. | 

### Return type

[**DocumentGetPreviewResponse**](DocumentGetPreviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentLoad

> DocumentLoadResponse DocumentLoad(ctx, loadDocumentFromByteArrayParameters)

Loads the provided document file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadDocumentFromByteArrayParameters** | [**LoadDocumentFromByteArrayParameters**](LoadDocumentFromByteArrayParameters.md)| A LoadDocumentFromByteArrayParameters object specifying the parameters of the action. | 

### Return type

[**DocumentLoadResponse**](DocumentLoadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentLoadFromURI

> DocumentLoadResponse DocumentLoadFromURI(ctx, loadDocumentFromUriParameters)

Loads the provided document file from an URI.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadDocumentFromUriParameters** | [**LoadDocumentFromUriParameters**](LoadDocumentFromUriParameters.md)| A LoadDocumentFromURIParameters object specifying the parameters of the action. | 

### Return type

[**DocumentLoadResponse**](DocumentLoadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentLoadMultipart

> DocumentLoadResponse DocumentLoadMultipart(ctx, fileData, optional)

Loads the provided document file using Multipart Upload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileData** | ***os.File*****os.File**| The data of the document. | 
 **optional** | ***DocumentLoadMultipartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DocumentLoadMultipartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadDocumentParameters** | [**optional.Interface of LoadDocumentParameters**](LoadDocumentParameters.md)|  | 

### Return type

[**DocumentLoadResponse**](DocumentLoadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

