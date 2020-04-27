# \ConfigApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigGetAPIVersion**](ConfigApi.md#ConfigGetAPIVersion) | **Get** /api/config/ConfigGetAPIVersion | 
[**ConfigGetMaxAllowedContentLength**](ConfigApi.md#ConfigGetMaxAllowedContentLength) | **Get** /api/config/ConfigGetMaxAllowedContentLength | Gets the maximal length of a request content, in bytes.
[**ConfigGetSuggestedClientTimeout**](ConfigApi.md#ConfigGetSuggestedClientTimeout) | **Get** /api/config/ConfigGetSuggestedClientTimeout | Gets the suggested client API timeout, in milliseconds.
[**ConfigGetSuggestedMaxClientThreads**](ConfigApi.md#ConfigGetSuggestedMaxClientThreads) | **Get** /api/config/ConfigGetSuggestedMaxClientThreads | Gets the suggested maximum number of threads to be used simultaneously by a client application.
[**ConfigGetSupportedFonts**](ConfigApi.md#ConfigGetSupportedFonts) | **Get** /api/config/ConfigGetSupportedFonts | Gets the list of supported fonts for text drawing operations.
[**ConfigGetSupportedOCRLanguages**](ConfigApi.md#ConfigGetSupportedOCRLanguages) | **Get** /api/config/ConfigGetSupportedOCRLanguages | Gets the list of supported languages for OCR.



## ConfigGetAPIVersion

> string ConfigGetAPIVersion(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGetMaxAllowedContentLength

> LongResponse ConfigGetMaxAllowedContentLength(ctx, )

Gets the maximal length of a request content, in bytes.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LongResponse**](LongResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGetSuggestedClientTimeout

> IntegerResponse ConfigGetSuggestedClientTimeout(ctx, )

Gets the suggested client API timeout, in milliseconds.

### Required Parameters

This endpoint does not need any parameter.

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


## ConfigGetSuggestedMaxClientThreads

> IntegerResponse ConfigGetSuggestedMaxClientThreads(ctx, )

Gets the suggested maximum number of threads to be used simultaneously by a client application.

### Required Parameters

This endpoint does not need any parameter.

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


## ConfigGetSupportedFonts

> []Font ConfigGetSupportedFonts(ctx, )

Gets the list of supported fonts for text drawing operations.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Font**](Font.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGetSupportedOCRLanguages

> StringArrayResponse ConfigGetSupportedOCRLanguages(ctx, )

Gets the list of supported languages for OCR.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StringArrayResponse**](StringArrayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

