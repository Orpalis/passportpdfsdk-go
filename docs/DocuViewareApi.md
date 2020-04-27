# \DocuViewareApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocuViewareGetControl**](DocuViewareApi.md#DocuViewareGetControl) | **Post** /api/docuvieware/DocuViewareGetControl | Gets the HTML dom of a DocuVieware control.
[**DocuViewareGetVersion**](DocuViewareApi.md#DocuViewareGetVersion) | **Get** /api/docuvieware/DocuViewareGetVersion | Get the DocuVieware engine version.
[**DocuViewareSave**](DocuViewareApi.md#DocuViewareSave) | **Post** /api/docuvieware/DocuViewareSave | Saves the document being handled by a specific DocuVieware control, in its current state.



## DocuViewareGetControl

> DocuViewareGetControlResponse DocuViewareGetControl(ctx, docuViewareGetControlParameters)

Gets the HTML dom of a DocuVieware control.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**docuViewareGetControlParameters** | [**DocuViewareGetControlParameters**](DocuViewareGetControlParameters.md)| A DocuViewareGetControlParameters object specifying the parameters of the action. | 

### Return type

[**DocuViewareGetControlResponse**](DocuViewareGetControlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocuViewareGetVersion

> StringResponse DocuViewareGetVersion(ctx, )

Get the DocuVieware engine version.

### Required Parameters

This endpoint does not need any parameter.

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


## DocuViewareSave

> DocuViewareSaveResponse DocuViewareSave(ctx, docuViewareSaveParameters)

Saves the document being handled by a specific DocuVieware control, in its current state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**docuViewareSaveParameters** | [**DocuViewareSaveParameters**](DocuViewareSaveParameters.md)| A DocuViewareSaveParameters object specifying the parameters of the action. | 

### Return type

[**DocuViewareSaveResponse**](DocuViewareSaveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

