# \ImageApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageAdjust**](ImageApi.md#ImageAdjust) | **Post** /api/image/ImageAdjust | Adjusts a previously uploaded image.
[**ImageAutoCrop**](ImageApi.md#ImageAutoCrop) | **Post** /api/image/ImageAutoCrop | Automatically crops a previously uploaded image.
[**ImageCleanupDocument**](ImageApi.md#ImageCleanupDocument) | **Post** /api/image/ImageCleanupDocument | Cleanup a previously uploaded image.
[**ImageCloneRegions**](ImageApi.md#ImageCloneRegions) | **Post** /api/image/ImageCloneRegions | Clones regions from a previously uploaded image into new images.
[**ImageClose**](ImageApi.md#ImageClose) | **Post** /api/image/ImageClose | Closes a previously uploaded image.
[**ImageConvertColorDepth**](ImageApi.md#ImageConvertColorDepth) | **Post** /api/image/ImageConvertColorDepth | Converts the color depth of a previously uploaded image.
[**ImageCrop**](ImageApi.md#ImageCrop) | **Post** /api/image/ImageCrop | Crops a previously uploaded image.
[**ImageDeletePage**](ImageApi.md#ImageDeletePage) | **Post** /api/image/ImageDeletePage | Deletes a page range from a previously uploaded image.
[**ImageDetectBlankPages**](ImageApi.md#ImageDetectBlankPages) | **Post** /api/image/ImageDetectBlankPages | Detects the blank page(s) from a previously uploaded image and offers to remove them.
[**ImageDetectColor**](ImageApi.md#ImageDetectColor) | **Post** /api/image/ImageDetectColor | Performs color detection  on a previously uploaded image.
[**ImageDetectPageOrientation**](ImageApi.md#ImageDetectPageOrientation) | **Post** /api/image/ImageDetectPageOrientation | Detects the orientation of the page(s) of a previously uploaded image and offers to automatically rotate them.
[**ImageFilter**](ImageApi.md#ImageFilter) | **Post** /api/image/ImageFilter | Applies filters to a previously uploaded image.
[**ImageGetPageThumbnail**](ImageApi.md#ImageGetPageThumbnail) | **Post** /api/image/ImageGetPageThumbnail | Gets a thumbnail of each page within the provided page range from a previously uploaded image.
[**ImageGetSupportedFileExtensions**](ImageApi.md#ImageGetSupportedFileExtensions) | **Get** /api/image/ImageGetSupportedFileExtensions | Gets the supported file extensions by the image loading actions.
[**ImageLoad**](ImageApi.md#ImageLoad) | **Post** /api/image/ImageLoad | Loads the provided image file.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.
[**ImageLoadMultipart**](ImageApi.md#ImageLoadMultipart) | **Post** /api/image/ImageLoadMultipart | Loads the provided image file using Multipart Upload.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.
[**ImageMICR**](ImageApi.md#ImageMICR) | **Post** /api/image/ImageMICR | Performs MICR (Magnetic Ink Character Recognition) on a previously uploaded image.
[**ImageReadBarcodes**](ImageApi.md#ImageReadBarcodes) | **Post** /api/image/ImageReadBarcodes | Reads barcodes from a previously uploaded image.
[**ImageResize**](ImageApi.md#ImageResize) | **Post** /api/image/ImageResize | Resizes a previously uploaded image.
[**ImageRotate**](ImageApi.md#ImageRotate) | **Post** /api/image/ImageRotate | Rotates and/or flips a previously uploaded image.
[**ImageSaveAsJPEG**](ImageApi.md#ImageSaveAsJPEG) | **Post** /api/image/ImageSaveAsJPEG | Saves a previously uploaded image as JPEG, and sends the file data in a JSON-serialized object.
[**ImageSaveAsJPEGFile**](ImageApi.md#ImageSaveAsJPEGFile) | **Post** /api/image/ImageSaveAsJPEGFile | Saves a previously uploaded image as JPEG, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSaveAsPDF**](ImageApi.md#ImageSaveAsPDF) | **Post** /api/image/ImageSaveAsPDF | Saves a previously uploaded image as PDF, and sends the file data in a JSON-serialized object.
[**ImageSaveAsPDFFile**](ImageApi.md#ImageSaveAsPDFFile) | **Post** /api/image/ImageSaveAsPDFFile | Saves a previously uploaded image as PDF, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSaveAsPDFMRC**](ImageApi.md#ImageSaveAsPDFMRC) | **Post** /api/image/ImageSaveAsPDFMRC | Saves a previously uploaded image as PDF using MRC compression, and sends the file data in a JSON-serialized object.
[**ImageSaveAsPDFMRCFile**](ImageApi.md#ImageSaveAsPDFMRCFile) | **Post** /api/image/ImageSaveAsPDFMRCFile | Saves a previously uploaded image as PDF using MRC compression, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSaveAsPNG**](ImageApi.md#ImageSaveAsPNG) | **Post** /api/image/ImageSaveAsPNG | Saves a previously uploaded image as PNG, and sends the file data in a JSON-serialized object.
[**ImageSaveAsPNGFile**](ImageApi.md#ImageSaveAsPNGFile) | **Post** /api/image/ImageSaveAsPNGFile | Saves a previously uploaded image as PNG, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSaveAsTIFF**](ImageApi.md#ImageSaveAsTIFF) | **Post** /api/image/ImageSaveAsTIFF | Saves a previously uploaded image as TIFF, and sends the file data in a JSON-serialized object.
[**ImageSaveAsTIFFFile**](ImageApi.md#ImageSaveAsTIFFFile) | **Post** /api/image/ImageSaveAsTIFFFile | Saves a previously uploaded image as TIFF, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSaveAsTIFFMultipage**](ImageApi.md#ImageSaveAsTIFFMultipage) | **Post** /api/image/ImageSaveAsTIFFMultipage | Saves a previously uploaded image as multipage TIFF, and sends the file data in a JSON-serialized object.
[**ImageSaveAsTIFFMultipageFile**](ImageApi.md#ImageSaveAsTIFFMultipageFile) | **Post** /api/image/ImageSaveAsTIFFMultipageFile | Saves a previously uploaded image as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).
[**ImageSwapPages**](ImageApi.md#ImageSwapPages) | **Post** /api/image/ImageSwapPages | Swaps two pages from a previously uploaded image.



## ImageAdjust

> ImageAdjustResponse ImageAdjust(ctx, imageAdjustParameters)

Adjusts a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageAdjustParameters** | [**ImageAdjustParameters**](ImageAdjustParameters.md)| An ImageAdjustParameters object specifying the parameters for the action. | 

### Return type

[**ImageAdjustResponse**](ImageAdjustResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageAutoCrop

> ImageAutoCropResponse ImageAutoCrop(ctx, imageAutoCropParameters)

Automatically crops a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageAutoCropParameters** | [**ImageAutoCropParameters**](ImageAutoCropParameters.md)| An ImageAutoCropParameters object specifying the parameters for the action. | 

### Return type

[**ImageAutoCropResponse**](ImageAutoCropResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageCleanupDocument

> ImageCleanupDocumentResponse ImageCleanupDocument(ctx, imageCleanupDocumentParameters)

Cleanup a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageCleanupDocumentParameters** | [**ImageCleanupDocumentParameters**](ImageCleanupDocumentParameters.md)| An ImageCleanupDocumentParameters object specifying the parameters for the action. | 

### Return type

[**ImageCleanupDocumentResponse**](ImageCleanupDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageCloneRegions

> ImageCloneRegionsResponse ImageCloneRegions(ctx, imageCloneRegionsParameters)

Clones regions from a previously uploaded image into new images.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageCloneRegionsParameters** | [**ImageCloneRegionsParameters**](ImageCloneRegionsParameters.md)| An ImageCloneRegionsParameters object specifying the parameters of the action. | 

### Return type

[**ImageCloneRegionsResponse**](ImageCloneRegionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageClose

> ImageCloseResponse ImageClose(ctx, imageCloseParameters)

Closes a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageCloseParameters** | [**ImageCloseParameters**](ImageCloseParameters.md)| An ImageCloseParameters object specifying the parameters of the action. | 

### Return type

[**ImageCloseResponse**](ImageCloseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageConvertColorDepth

> ImageConvertColorDepthResponse ImageConvertColorDepth(ctx, imageConvertColorDepthParameters)

Converts the color depth of a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageConvertColorDepthParameters** | [**ImageConvertColorDepthParameters**](ImageConvertColorDepthParameters.md)| An ImageConvertColorDepthParameters object specifying the parameters for the action. | 

### Return type

[**ImageConvertColorDepthResponse**](ImageConvertColorDepthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageCrop

> ImageCropResponse ImageCrop(ctx, imageCropParameters)

Crops a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageCropParameters** | [**ImageCropParameters**](ImageCropParameters.md)| An ImageCropParameters object specifying the parameters for the action. | 

### Return type

[**ImageCropResponse**](ImageCropResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDeletePage

> ImageDeletePageResponse ImageDeletePage(ctx, imageDeletePageParameters)

Deletes a page range from a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDeletePageParameters** | [**ImageDeletePageParameters**](ImageDeletePageParameters.md)| An ImageDeletePageParameters object specifying the parameters of the action. | 

### Return type

[**ImageDeletePageResponse**](ImageDeletePageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDetectBlankPages

> ImageDetectBlankPagesResponse ImageDetectBlankPages(ctx, imageDetectBlankPagesParameters)

Detects the blank page(s) from a previously uploaded image and offers to remove them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDetectBlankPagesParameters** | [**ImageDetectBlankPagesParameters**](ImageDetectBlankPagesParameters.md)| An ImageDetectBlankPagesParameters object specifying the parameters of the action. | 

### Return type

[**ImageDetectBlankPagesResponse**](ImageDetectBlankPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDetectColor

> ImageDetectColorResponse ImageDetectColor(ctx, imageDetectColorParameters)

Performs color detection  on a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDetectColorParameters** | [**ImageDetectColorParameters**](ImageDetectColorParameters.md)| An ImageDetectColorParameters object specifying the parameters for the action. | 

### Return type

[**ImageDetectColorResponse**](ImageDetectColorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDetectPageOrientation

> ImageDetectPageOrientationResponse ImageDetectPageOrientation(ctx, imageDetectPageOrientationParameters)

Detects the orientation of the page(s) of a previously uploaded image and offers to automatically rotate them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDetectPageOrientationParameters** | [**ImageDetectPageOrientationParameters**](ImageDetectPageOrientationParameters.md)| An ImagedetectPageOrientationParameters object specifying the parameters of the action. | 

### Return type

[**ImageDetectPageOrientationResponse**](ImageDetectPageOrientationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageFilter

> ImageFilterResponse ImageFilter(ctx, imageFilterParameters)

Applies filters to a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageFilterParameters** | [**ImageFilterParameters**](ImageFilterParameters.md)| An ImageFilterParameters object specifying the parameters for the action. | 

### Return type

[**ImageFilterResponse**](ImageFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageGetPageThumbnail

> ImageGetPageThumbnailResponse ImageGetPageThumbnail(ctx, imageGetPageThumbnailParameters)

Gets a thumbnail of each page within the provided page range from a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageGetPageThumbnailParameters** | [**ImageGetPageThumbnailParameters**](ImageGetPageThumbnailParameters.md)| A PDFGetPageThumbnailParameters object specifying the parameters of the action. | 

### Return type

[**ImageGetPageThumbnailResponse**](ImageGetPageThumbnailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageGetSupportedFileExtensions

> StringArrayResponse ImageGetSupportedFileExtensions(ctx, )

Gets the supported file extensions by the image loading actions.

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


## ImageLoad

> ImageLoadResponse ImageLoad(ctx, loadImageFromByteArrayParameters)

Loads the provided image file.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadImageFromByteArrayParameters** | [**LoadImageFromByteArrayParameters**](LoadImageFromByteArrayParameters.md)| A LoadImageFromByteArrayParameters object specifying the parameters of the action. | 

### Return type

[**ImageLoadResponse**](ImageLoadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageLoadMultipart

> ImageLoadResponse ImageLoadMultipart(ctx, fileData, optional)

Loads the provided image file using Multipart Upload.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileData** | ***os.File*****os.File**| The data of the document. | 
 **optional** | ***ImageLoadMultipartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImageLoadMultipartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadImageParameters** | [**optional.Interface of LoadImageParameters**](LoadImageParameters.md)|  | 

### Return type

[**ImageLoadResponse**](ImageLoadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageMICR

> ImageMicrResponse ImageMICR(ctx, imageMicrParameters)

Performs MICR (Magnetic Ink Character Recognition) on a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageMicrParameters** | [**ImageMicrParameters**](ImageMicrParameters.md)| An ImageMICRParameters object specifying the parameters of the action. | 

### Return type

[**ImageMicrResponse**](ImageMICRResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageReadBarcodes

> ReadBarcodesResponse ImageReadBarcodes(ctx, imageReadBarcodesParameters)

Reads barcodes from a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageReadBarcodesParameters** | [**ImageReadBarcodesParameters**](ImageReadBarcodesParameters.md)| An ImageReadBarcodesParameters object specifying the parameters of the action. | 

### Return type

[**ReadBarcodesResponse**](ReadBarcodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageResize

> ImageResizeResponse ImageResize(ctx, imageResizeParameters)

Resizes a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageResizeParameters** | [**ImageResizeParameters**](ImageResizeParameters.md)| An ImageResizeParameters object specifying the parameters for the action. | 

### Return type

[**ImageResizeResponse**](ImageResizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageRotate

> ImageRotateResponse ImageRotate(ctx, imageRotateParameters)

Rotates and/or flips a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageRotateParameters** | [**ImageRotateParameters**](ImageRotateParameters.md)| An ImageRotateParameters object specifying the parameters for the action. | 

### Return type

[**ImageRotateResponse**](ImageRotateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsJPEG

> ImageSaveAsJpegResponse ImageSaveAsJPEG(ctx, imageSaveAsJpegParameters)

Saves a previously uploaded image as JPEG, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsJpegParameters** | [**ImageSaveAsJpegParameters**](ImageSaveAsJpegParameters.md)| An ImageSaveAsJPEGParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsJpegResponse**](ImageSaveAsJPEGResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsJPEGFile

> *os.File ImageSaveAsJPEGFile(ctx, imageSaveAsJpegParameters)

Saves a previously uploaded image as JPEG, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsJpegParameters** | [**ImageSaveAsJpegParameters**](ImageSaveAsJpegParameters.md)| An ImageSaveAsJPEGParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPDF

> ImageSaveAsPdfResponse ImageSaveAsPDF(ctx, imageSaveAsPdfParameters)

Saves a previously uploaded image as PDF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPdfParameters** | [**ImageSaveAsPdfParameters**](ImageSaveAsPdfParameters.md)| An ImagesaveAsPDFParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsPdfResponse**](ImageSaveAsPDFResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPDFFile

> *os.File ImageSaveAsPDFFile(ctx, imageSaveAsPdfParameters)

Saves a previously uploaded image as PDF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPdfParameters** | [**ImageSaveAsPdfParameters**](ImageSaveAsPdfParameters.md)| An ImagesaveAsPDFParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPDFMRC

> ImageSaveAsPdfmrcResponse ImageSaveAsPDFMRC(ctx, imageSaveAsPdfmrcParameters)

Saves a previously uploaded image as PDF using MRC compression, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPdfmrcParameters** | [**ImageSaveAsPdfmrcParameters**](ImageSaveAsPdfmrcParameters.md)| An ImagesaveAsPDFMRCParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsPdfmrcResponse**](ImageSaveAsPDFMRCResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPDFMRCFile

> *os.File ImageSaveAsPDFMRCFile(ctx, imageSaveAsPdfmrcParameters)

Saves a previously uploaded image as PDF using MRC compression, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPdfmrcParameters** | [**ImageSaveAsPdfmrcParameters**](ImageSaveAsPdfmrcParameters.md)| An ImagesaveAsPDFMRCParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPNG

> ImageSaveAsPngResponse ImageSaveAsPNG(ctx, imageSaveAsPngParameters)

Saves a previously uploaded image as PNG, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPngParameters** | [**ImageSaveAsPngParameters**](ImageSaveAsPngParameters.md)| An ImageSaveAsPNGParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsPngResponse**](ImageSaveAsPNGResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsPNGFile

> *os.File ImageSaveAsPNGFile(ctx, imageSaveAsPngParameters)

Saves a previously uploaded image as PNG, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsPngParameters** | [**ImageSaveAsPngParameters**](ImageSaveAsPngParameters.md)| An ImageSaveAsPNGParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsTIFF

> ImageSaveAsTiffResponse ImageSaveAsTIFF(ctx, imageSaveAsTiffParameters)

Saves a previously uploaded image as TIFF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsTiffParameters** | [**ImageSaveAsTiffParameters**](ImageSaveAsTiffParameters.md)| An ImageSaveAsTIFFParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsTiffResponse**](ImageSaveAsTIFFResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsTIFFFile

> *os.File ImageSaveAsTIFFFile(ctx, imageSaveAsTiffParameters)

Saves a previously uploaded image as TIFF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsTiffParameters** | [**ImageSaveAsTiffParameters**](ImageSaveAsTiffParameters.md)| An ImageSaveAsTIFFParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsTIFFMultipage

> ImageSaveAsTiffMultipageResponse ImageSaveAsTIFFMultipage(ctx, imageSaveAsTiffMultipageParameters)

Saves a previously uploaded image as multipage TIFF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsTiffMultipageParameters** | [**ImageSaveAsTiffMultipageParameters**](ImageSaveAsTiffMultipageParameters.md)| An ImageSaveAsTIFFMultipageParameters object specifying the parameters of the action. | 

### Return type

[**ImageSaveAsTiffMultipageResponse**](ImageSaveAsTIFFMultipageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSaveAsTIFFMultipageFile

> *os.File ImageSaveAsTIFFMultipageFile(ctx, imageSaveAsTiffMultipageParameters)

Saves a previously uploaded image as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSaveAsTiffMultipageParameters** | [**ImageSaveAsTiffMultipageParameters**](ImageSaveAsTiffMultipageParameters.md)| An ImageSaveAsTIFFMultipageParameters object specifying the parameters of the action. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageSwapPages

> ImageSwapPagesResponse ImageSwapPages(ctx, imageSwapPagesParameters)

Swaps two pages from a previously uploaded image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSwapPagesParameters** | [**ImageSwapPagesParameters**](ImageSwapPagesParameters.md)| An ImageSwapPagesParameters object specifying the parameters of the action. | 

### Return type

[**ImageSwapPagesResponse**](ImageSwapPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

