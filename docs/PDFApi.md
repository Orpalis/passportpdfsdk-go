# \PDFApi

All URIs are relative to *https://passportpdfapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Annotate**](PDFApi.md#Annotate) | **Post** /api/pdf/Annotate | Annotates a previously uploaded document.
[**AutoDeskew**](PDFApi.md#AutoDeskew) | **Post** /api/pdf/AutoDeskew | Performs auto deskew on a page range of a previously uploaded document.
[**ClearPage**](PDFApi.md#ClearPage) | **Post** /api/pdf/ClearPage | Clears a page range from a previously uploaded document.
[**ClonePage**](PDFApi.md#ClonePage) | **Post** /api/pdf/ClonePage | Clones specific pages from a previously uploaded document to another previously uploaded document.
[**ClosePDF**](PDFApi.md#ClosePDF) | **Post** /api/pdf/ClosePDF | Closes a previously uploaded document.
[**ConvertToPDFA**](PDFApi.md#ConvertToPDFA) | **Post** /api/pdf/ConvertToPDFA | Converts a previously uploaded document to PDF/A.
[**DeletePage**](PDFApi.md#DeletePage) | **Post** /api/pdf/DeletePage | Deletes a page range from a previously uploaded document.
[**DetectPageOrientation**](PDFApi.md#DetectPageOrientation) | **Post** /api/pdf/DetectPageOrientation | Detects the orientation of the page(s) of a previously uploaded document and offers to automatically rotate them.
[**DigiSign**](PDFApi.md#DigiSign) | **Post** /api/pdf/DigiSign | Signs a previously uploaded document digitally.
[**DrawImage**](PDFApi.md#DrawImage) | **Post** /api/pdf/DrawImage | Draws an image on a page range of a previously uploaded document.
[**ExtractPage**](PDFApi.md#ExtractPage) | **Post** /api/pdf/ExtractPage | Extracts a page range from a previously uploaded document into one or several new documents.
[**ExtractText**](PDFApi.md#ExtractText) | **Post** /api/pdf/ExtractText | Extracts text from the document pages.
[**Flatten**](PDFApi.md#Flatten) | **Post** /api/pdf/Flatten | Flattens the form-fields, annotations, and/or the layers of a previously uploaded document.
[**GetInfo**](PDFApi.md#GetInfo) | **Post** /api/pdf/GetInfo | Gets information about a previously uploaded document.
[**GetPDFImportSupportedFileExtensions**](PDFApi.md#GetPDFImportSupportedFileExtensions) | **Get** /api/pdf/GetPDFImportSupportedFileExtensions | Gets the supported file extensions by the LoadDocumentAsPDF action.
[**GetPageThumbnail**](PDFApi.md#GetPageThumbnail) | **Post** /api/pdf/GetPageThumbnail | Gets a thumbnail of each page within the provided page range from a previously uploaded document.
[**InsertImage**](PDFApi.md#InsertImage) | **Post** /api/pdf/InsertImage | Inserts an image on a new page of a previously uploaded document.
[**InsertNewPage**](PDFApi.md#InsertNewPage) | **Post** /api/pdf/InsertNewPage | Inserts one or more new blank pages to a specific position in a previously uploaded document.
[**InsertPageNumber**](PDFApi.md#InsertPageNumber) | **Post** /api/pdf/InsertPageNumber | Inserts page number(s) on a document.
[**InsertText**](PDFApi.md#InsertText) | **Post** /api/pdf/InsertText | Inserts text on a document.
[**Linearize**](PDFApi.md#Linearize) | **Post** /api/pdf/Linearize | Linearizes a previously uploaded document.
[**LoadDocumentAsPDF**](PDFApi.md#LoadDocumentAsPDF) | **Post** /api/pdf/LoadDocument | Imports the provided document as PDF.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.
[**LoadDocumentAsPDFMultipart**](PDFApi.md#LoadDocumentAsPDFMultipart) | **Post** /api/pdf/LoadDocumentAsPDFMultipart | Imports the provided document as PDF using Multipart Upload.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.
[**Merge**](PDFApi.md#Merge) | **Post** /api/pdf/Merge | Merges several previously uploaded documents into a new PDF.
[**MergePages**](PDFApi.md#MergePages) | **Post** /api/pdf/MergePages | Merges multiple pages, vertically, within a previously uploaded document into one single page.
[**MovePage**](PDFApi.md#MovePage) | **Post** /api/pdf/MovePage | Moves a page range from a previously uploaded document.
[**OCR**](PDFApi.md#OCR) | **Post** /api/pdf/OCR | Performs optical character recognition on a page range of a previously uploaded document.  The recognized text is added as invisible text on each processed page.  No token is charged for blank pages.
[**Protect**](PDFApi.md#Protect) | **Post** /api/pdf/Protect | Protects a previously uploaded document.
[**ReadBarcodes**](PDFApi.md#ReadBarcodes) | **Post** /api/pdf/ReadBarcodes | Reads barcodes from a previously uploaded document.
[**Reduce**](PDFApi.md#Reduce) | **Post** /api/pdf/Reduce | Reduces the size of a previously uploaded document.
[**RemovePageFormFields**](PDFApi.md#RemovePageFormFields) | **Post** /api/pdf/RemovePageFormFields | Removes the form fields from a page range of a previously uploaded document.
[**RemoveText**](PDFApi.md#RemoveText) | **Post** /api/pdf/RemoveText | Removes text (all text or only invisible one) from a previously uploaded PDF.
[**ReorderPages**](PDFApi.md#ReorderPages) | **Post** /api/pdf/ReorderPages | Reorders pages of a previously uploaded document.
[**RepairDocument**](PDFApi.md#RepairDocument) | **Post** /api/pdf/RepairDocument | Repairs a previously uploaded PDF document.
[**RotatePageStandard**](PDFApi.md#RotatePageStandard) | **Post** /api/pdf/RotatePageStandard | Rotates (standardly) a page range from a previously uploaded document.
[**SaveAsJPEG**](PDFApi.md#SaveAsJPEG) | **Post** /api/pdf/SaveAsJPEG | Saves a previously uploaded document as JPEG, and sends the file data in a JSON-serialized object.
[**SaveAsJPEGFile**](PDFApi.md#SaveAsJPEGFile) | **Post** /api/pdf/SaveAsJPEGFile | Saves a previously uploaded document as JPEG, and streams the file binary data to the response (this is the most efficient download method).
[**SaveAsPNG**](PDFApi.md#SaveAsPNG) | **Post** /api/pdf/SaveAsPNG | Saves a previously uploaded document as PNG, and sends the file data in a JSON-serialized object.
[**SaveAsPNGFile**](PDFApi.md#SaveAsPNGFile) | **Post** /api/pdf/SaveAsPNGFile | Saves a previously uploaded document as PNG, and streams the file binary data to the response (this is the most efficient download method).
[**SaveAsTIFF**](PDFApi.md#SaveAsTIFF) | **Post** /api/pdf/SaveAsTIFF | Saves a previously uploaded document as TIFF, and sends the file data in a JSON-serialized object.
[**SaveAsTIFFFile**](PDFApi.md#SaveAsTIFFFile) | **Post** /api/pdf/SaveAsTIFFFile | Saves a previously uploaded document as TIFF, and streams the file binary data to the response (this is the most efficient download method).
[**SaveAsTIFFMultipage**](PDFApi.md#SaveAsTIFFMultipage) | **Post** /api/pdf/SaveAsTIFFMultipage | Saves a previously uploaded document as multipage TIFF, and sends the file data in a JSON-serialized object.
[**SaveAsTIFFMultipageFile**](PDFApi.md#SaveAsTIFFMultipageFile) | **Post** /api/pdf/SaveAsTIFFMultipageFile | Saves a previously uploaded document as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).
[**SaveDocument**](PDFApi.md#SaveDocument) | **Post** /api/pdf/SaveDocument | Saves a previously uploaded document as PDF, and sends the file data in a JSON-serialized object.
[**SaveDocumentToFile**](PDFApi.md#SaveDocumentToFile) | **Post** /api/pdf/SaveDocumentToFile | Saves a previously uploaded document as PDF, and streams the file binary data to the response (this is the most efficient download method).
[**ScalePage**](PDFApi.md#ScalePage) | **Post** /api/pdf/ScalePage | Scales a page range from a previously uploaded document.
[**SetInfo**](PDFApi.md#SetInfo) | **Post** /api/pdf/SetInfo | Sets information to a previously uploaded document.
[**SetInitialView**](PDFApi.md#SetInitialView) | **Post** /api/pdf/SetInitialView | Sets various document level initial view options to a previously uploaded document.
[**SetPageBox**](PDFApi.md#SetPageBox) | **Post** /api/pdf/SetPageBox | Sets pagebox to a page range from previously uploaded document.
[**SetPassword**](PDFApi.md#SetPassword) | **Post** /api/pdf/SetPassword | Unprotects a previously uploaded document.
[**Split**](PDFApi.md#Split) | **Post** /api/pdf/Split | Splits a previously uploaded document into new ones.
[**SwapPages**](PDFApi.md#SwapPages) | **Post** /api/pdf/SwapPages | Swaps two pages from a previously uploaded document.
[**Unprotect**](PDFApi.md#Unprotect) | **Post** /api/pdf/Unprotect | Unprotects a previously uploaded document.



## Annotate

> PdfAnnotateResponse Annotate(ctx, pdfAnnotateParameters)

Annotates a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfAnnotateParameters** | [**PdfAnnotateParameters**](PdfAnnotateParameters.md)| A PdfAnnotateParameters object specifying the parameters of the action. | 

### Return type

[**PdfAnnotateResponse**](PdfAnnotateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoDeskew

> PdfAutoDeskewResponse AutoDeskew(ctx, pdfAutoDeskewParameters)

Performs auto deskew on a page range of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfAutoDeskewParameters** | [**PdfAutoDeskewParameters**](PdfAutoDeskewParameters.md)| A PdfAutoDeskewParameters object specifying the parameters of the action. | 

### Return type

[**PdfAutoDeskewResponse**](PdfAutoDeskewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearPage

> PdfClearPageResponse ClearPage(ctx, pdfClearPageParameters)

Clears a page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfClearPageParameters** | [**PdfClearPageParameters**](PdfClearPageParameters.md)| A PdfClearPageParameters object specifying the parameters of the action. | 

### Return type

[**PdfClearPageResponse**](PdfClearPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClonePage

> PdfClonePageResponse ClonePage(ctx, pdfClonePageParameters)

Clones specific pages from a previously uploaded document to another previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfClonePageParameters** | [**PdfClonePageParameters**](PdfClonePageParameters.md)| A PdfClonePageParameters object specifying the parameters of the action. | 

### Return type

[**PdfClonePageResponse**](PdfClonePageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClosePDF

> PdfCloseDocumentResponse ClosePDF(ctx, pdfCloseDocumentParameters)

Closes a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfCloseDocumentParameters** | [**PdfCloseDocumentParameters**](PdfCloseDocumentParameters.md)| A PdfCloseDocumentParameters object specifying the parameters of the action. | 

### Return type

[**PdfCloseDocumentResponse**](PdfCloseDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertToPDFA

> PdfConvertToPdfaResponse ConvertToPDFA(ctx, pdfConvertToPdfaParameters)

Converts a previously uploaded document to PDF/A.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfConvertToPdfaParameters** | [**PdfConvertToPdfaParameters**](PdfConvertToPdfaParameters.md)| A PdfConvertToPDFAParameters object specifying the parameters of the action. | 

### Return type

[**PdfConvertToPdfaResponse**](PdfConvertToPDFAResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePage

> PdfDeletePageResponse DeletePage(ctx, pdfDeletePageParameters)

Deletes a page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfDeletePageParameters** | [**PdfDeletePageParameters**](PdfDeletePageParameters.md)| A PdfDeletePageParameters object specifying the parameters of the action. | 

### Return type

[**PdfDeletePageResponse**](PdfDeletePageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetectPageOrientation

> PdfDetectPageOrientationResponse DetectPageOrientation(ctx, pdfDetectPageOrientationParameters)

Detects the orientation of the page(s) of a previously uploaded document and offers to automatically rotate them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfDetectPageOrientationParameters** | [**PdfDetectPageOrientationParameters**](PdfDetectPageOrientationParameters.md)| A PdfDetectPageOrientationParameters object specifying the parameters of the action. | 

### Return type

[**PdfDetectPageOrientationResponse**](PdfDetectPageOrientationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DigiSign

> PdfDigiSignResponse DigiSign(ctx, pdfDigiSignParameters)

Signs a previously uploaded document digitally.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfDigiSignParameters** | [**PdfDigiSignParameters**](PdfDigiSignParameters.md)| A PdfDigiSignParameters object specifying the parameters of the action. | 

### Return type

[**PdfDigiSignResponse**](PdfDigiSignResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrawImage

> PdfDrawImageResponse DrawImage(ctx, pdfDrawImageParameters)

Draws an image on a page range of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfDrawImageParameters** | [**PdfDrawImageParameters**](PdfDrawImageParameters.md)| A PdfDrawImageParameters object specifying the parameters of the action. | 

### Return type

[**PdfDrawImageResponse**](PdfDrawImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractPage

> PdfExtractPageResponse ExtractPage(ctx, pdfExtractPageParameters)

Extracts a page range from a previously uploaded document into one or several new documents.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfExtractPageParameters** | [**PdfExtractPageParameters**](PdfExtractPageParameters.md)| A PdfExtractPageParameters object specifying the parameters of the action. | 

### Return type

[**PdfExtractPageResponse**](PdfExtractPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractText

> PdfExtractTextResponse ExtractText(ctx, pdfExtractTextParameters)

Extracts text from the document pages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfExtractTextParameters** | [**PdfExtractTextParameters**](PdfExtractTextParameters.md)| A PdfExtractTextParameters object specifying the parameters of the action. | 

### Return type

[**PdfExtractTextResponse**](PdfExtractTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Flatten

> PdfFlattenResponse Flatten(ctx, pdfFlattenParameters)

Flattens the form-fields, annotations, and/or the layers of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfFlattenParameters** | [**PdfFlattenParameters**](PdfFlattenParameters.md)| A PdfFlatten object specifying the parameters of the action. | 

### Return type

[**PdfFlattenResponse**](PdfFlattenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> PdfGetInfoResponse GetInfo(ctx, pdfGetInfoParameters)

Gets information about a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfGetInfoParameters** | [**PdfGetInfoParameters**](PdfGetInfoParameters.md)| A PdfGetInfoParameters object specifying the parameters of the action. | 

### Return type

[**PdfGetInfoResponse**](PdfGetInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPDFImportSupportedFileExtensions

> StringArrayResponse GetPDFImportSupportedFileExtensions(ctx, )

Gets the supported file extensions by the LoadDocumentAsPDF action.

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


## GetPageThumbnail

> PdfGetPageThumbnailResponse GetPageThumbnail(ctx, pdfGetPageThumbnailParameters)

Gets a thumbnail of each page within the provided page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfGetPageThumbnailParameters** | [**PdfGetPageThumbnailParameters**](PdfGetPageThumbnailParameters.md)| A PdfGetPageThumbnailParameters object specifying the parameters of the action. | 

### Return type

[**PdfGetPageThumbnailResponse**](PdfGetPageThumbnailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertImage

> PdfInsertImageResponse InsertImage(ctx, pdfInsertImageParameters)

Inserts an image on a new page of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfInsertImageParameters** | [**PdfInsertImageParameters**](PdfInsertImageParameters.md)| A PdfInsertImageParameters object specifying the parameters of the action. | 

### Return type

[**PdfInsertImageResponse**](PdfInsertImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertNewPage

> PdfInsertNewPageResponse InsertNewPage(ctx, pdfInsertNewPageParameters)

Inserts one or more new blank pages to a specific position in a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfInsertNewPageParameters** | [**PdfInsertNewPageParameters**](PdfInsertNewPageParameters.md)| A PdfInsertNewPageParameters object specifying the parameters of the action. | 

### Return type

[**PdfInsertNewPageResponse**](PdfInsertNewPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertPageNumber

> PdfInsertPageNumberResponse InsertPageNumber(ctx, pdfInsertPageNumberParameters)

Inserts page number(s) on a document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfInsertPageNumberParameters** | [**PdfInsertPageNumberParameters**](PdfInsertPageNumberParameters.md)| A PdfInsertPageNumberParameters object specifying the parameters of the action. | 

### Return type

[**PdfInsertPageNumberResponse**](PdfInsertPageNumberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertText

> PdfInsertTextResponse InsertText(ctx, pdfInsertTextParameters)

Inserts text on a document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfInsertTextParameters** | [**PdfInsertTextParameters**](PdfInsertTextParameters.md)| A PdfInsertTextParameters object specifying the parameters of the action. | 

### Return type

[**PdfInsertTextResponse**](PdfInsertTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Linearize

> PdfLinearizeResponse Linearize(ctx, pdfLinearizeParameters)

Linearizes a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfLinearizeParameters** | [**PdfLinearizeParameters**](PdfLinearizeParameters.md)| A PdfLinearizeParameters object specifying the parameters of the action. | 

### Return type

[**PdfLinearizeResponse**](PdfLinearizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadDocumentAsPDF

> PdfLoadDocumentResponse LoadDocumentAsPDF(ctx, pdfLoadDocumentFromByteArrayParameters)

Imports the provided document as PDF.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfLoadDocumentFromByteArrayParameters** | [**PdfLoadDocumentFromByteArrayParameters**](PdfLoadDocumentFromByteArrayParameters.md)| A PdfLoadDocumentFromByteArrayParameters object specifying the parameters of the action. | 

### Return type

[**PdfLoadDocumentResponse**](PdfLoadDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadDocumentAsPDFMultipart

> PdfLoadDocumentResponse LoadDocumentAsPDFMultipart(ctx, fileData, optional)

Imports the provided document as PDF using Multipart Upload.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileData** | ***os.File*****os.File**| The data of the document. | 
 **optional** | ***LoadDocumentAsPDFMultipartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoadDocumentAsPDFMultipartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadDocumentParameters** | [**optional.Interface of PdfLoadDocumentParameters**](PdfLoadDocumentParameters.md)|  | 

### Return type

[**PdfLoadDocumentResponse**](PdfLoadDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Merge

> PdfMergeResponse Merge(ctx, pdfMergeParameters)

Merges several previously uploaded documents into a new PDF.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfMergeParameters** | [**PdfMergeParameters**](PdfMergeParameters.md)| A PdfMergeParameters object specifying the parameters of the action. | 

### Return type

[**PdfMergeResponse**](PdfMergeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergePages

> PdfMergePagesResponse MergePages(ctx, pdfMergePagesParameters)

Merges multiple pages, vertically, within a previously uploaded document into one single page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfMergePagesParameters** | [**PdfMergePagesParameters**](PdfMergePagesParameters.md)| A PdfMergePages object specifying the parameters of the action. | 

### Return type

[**PdfMergePagesResponse**](PdfMergePagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovePage

> PdfMovePageResponse MovePage(ctx, pdfMovePageParameters)

Moves a page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfMovePageParameters** | [**PdfMovePageParameters**](PdfMovePageParameters.md)| A PdfMovePageParameters object specifying the parameters of the action. | 

### Return type

[**PdfMovePageResponse**](PdfMovePageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OCR

> PdfOcrResponse OCR(ctx, pdfOcrParameters)

Performs optical character recognition on a page range of a previously uploaded document.  The recognized text is added as invisible text on each processed page.  No token is charged for blank pages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfOcrParameters** | [**PdfOcrParameters**](PdfOcrParameters.md)| A PdfOCRParameters object specifying the parameters of the action. | 

### Return type

[**PdfOcrResponse**](PdfOCRResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Protect

> PdfProtectResponse Protect(ctx, pdfProtectParameters)

Protects a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfProtectParameters** | [**PdfProtectParameters**](PdfProtectParameters.md)| A PdfProtectParameters object specifying the parameters of the action. | 

### Return type

[**PdfProtectResponse**](PdfProtectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBarcodes

> ReadBarcodesResponse ReadBarcodes(ctx, pdfReadBarcodesParameters)

Reads barcodes from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfReadBarcodesParameters** | [**PdfReadBarcodesParameters**](PdfReadBarcodesParameters.md)| A PdfReadBarcodesParameters object specifying the parameters of the action. | 

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


## Reduce

> PdfReduceResponse Reduce(ctx, pdfReduceParameters)

Reduces the size of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfReduceParameters** | [**PdfReduceParameters**](PdfReduceParameters.md)| A PdfReduceParameters object specifying the parameters of the action. | 

### Return type

[**PdfReduceResponse**](PdfReduceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePageFormFields

> PdfRemovePageFormFieldsResponse RemovePageFormFields(ctx, pdfRemovePageFormFieldsParameters)

Removes the form fields from a page range of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfRemovePageFormFieldsParameters** | [**PdfRemovePageFormFieldsParameters**](PdfRemovePageFormFieldsParameters.md)| A PdfRemovePageFormFieldsParameters object specifying the parameters of the action. | 

### Return type

[**PdfRemovePageFormFieldsResponse**](PdfRemovePageFormFieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveText

> PdfRemoveTextResponse RemoveText(ctx, pdfRemoveTextParameters)

Removes text (all text or only invisible one) from a previously uploaded PDF.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfRemoveTextParameters** | [**PdfRemoveTextParameters**](PdfRemoveTextParameters.md)| A PdfRemoveTextParameters object specifying the parameters of the action. | 

### Return type

[**PdfRemoveTextResponse**](PdfRemoveTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderPages

> PdfReorderPagesResponse ReorderPages(ctx, pdfReorderPagesParameters)

Reorders pages of a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfReorderPagesParameters** | [**PdfReorderPagesParameters**](PdfReorderPagesParameters.md)| A PdfReorderPagesParameters object specifying the parameters of the action. | 

### Return type

[**PdfReorderPagesResponse**](PdfReorderPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepairDocument

> PdfRepairDocumentResponse RepairDocument(ctx, pdfRepairDocumentParameters)

Repairs a previously uploaded PDF document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfRepairDocumentParameters** | [**PdfRepairDocumentParameters**](PdfRepairDocumentParameters.md)| A PdfRepairDocumentParameters object specifying the parameters of the action. | 

### Return type

[**PdfRepairDocumentResponse**](PdfRepairDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotatePageStandard

> PdfRotatePageStandardResponse RotatePageStandard(ctx, pdfRotatePageStandardParameters)

Rotates (standardly) a page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfRotatePageStandardParameters** | [**PdfRotatePageStandardParameters**](PdfRotatePageStandardParameters.md)| A PdfRotatePageStandardParameters object specifying the parameters of the action. | 

### Return type

[**PdfRotatePageStandardResponse**](PdfRotatePageStandardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAsJPEG

> PdfSaveAsJpegResponse SaveAsJPEG(ctx, pdfSaveAsJpegParameters)

Saves a previously uploaded document as JPEG, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsJpegParameters** | [**PdfSaveAsJpegParameters**](PdfSaveAsJpegParameters.md)| A PdfSaveAsJPEGParameters object specifying the parameters of the action. | 

### Return type

[**PdfSaveAsJpegResponse**](PdfSaveAsJPEGResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAsJPEGFile

> *os.File SaveAsJPEGFile(ctx, pdfSaveAsJpegParameters)

Saves a previously uploaded document as JPEG, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsJpegParameters** | [**PdfSaveAsJpegParameters**](PdfSaveAsJpegParameters.md)| A PdfSaveAsJPEGParameters object specifying the parameters of the action. | 

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


## SaveAsPNG

> PdfSaveAsPngResponse SaveAsPNG(ctx, pdfSaveAsPngParameters)

Saves a previously uploaded document as PNG, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsPngParameters** | [**PdfSaveAsPngParameters**](PdfSaveAsPngParameters.md)| A PdfSaveAsPNGParameters object specifying the parameters of the action. | 

### Return type

[**PdfSaveAsPngResponse**](PdfSaveAsPNGResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAsPNGFile

> *os.File SaveAsPNGFile(ctx, pdfSaveAsPngParameters)

Saves a previously uploaded document as PNG, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsPngParameters** | [**PdfSaveAsPngParameters**](PdfSaveAsPngParameters.md)| A PdfSaveAsPNGParameters object specifying the parameters of the action. | 

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


## SaveAsTIFF

> PdfSaveAsTiffResponse SaveAsTIFF(ctx, pdfSaveAsTiffParameters)

Saves a previously uploaded document as TIFF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsTiffParameters** | [**PdfSaveAsTiffParameters**](PdfSaveAsTiffParameters.md)| A PdfSaveAsTIFFParameters object specifying the parameters of the action. | 

### Return type

[**PdfSaveAsTiffResponse**](PdfSaveAsTIFFResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAsTIFFFile

> *os.File SaveAsTIFFFile(ctx, pdfSaveAsTiffParameters)

Saves a previously uploaded document as TIFF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsTiffParameters** | [**PdfSaveAsTiffParameters**](PdfSaveAsTiffParameters.md)| A PdfSaveAsTIFFParameters object specifying the parameters of the action. | 

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


## SaveAsTIFFMultipage

> PdfSaveAsTiffMultipageResponse SaveAsTIFFMultipage(ctx, pdfSaveAsTiffMultipageParameters)

Saves a previously uploaded document as multipage TIFF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsTiffMultipageParameters** | [**PdfSaveAsTiffMultipageParameters**](PdfSaveAsTiffMultipageParameters.md)| A PdfSaveAsTIFFMultipageParameters object specifying the parameters of the action. | 

### Return type

[**PdfSaveAsTiffMultipageResponse**](PdfSaveAsTIFFMultipageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAsTIFFMultipageFile

> *os.File SaveAsTIFFMultipageFile(ctx, pdfSaveAsTiffMultipageParameters)

Saves a previously uploaded document as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveAsTiffMultipageParameters** | [**PdfSaveAsTiffMultipageParameters**](PdfSaveAsTiffMultipageParameters.md)| A PdfSaveAsTIFFMultipageParameters object specifying the parameters of the action. | 

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


## SaveDocument

> PdfSaveDocumentResponse SaveDocument(ctx, pdfSaveDocumentParameters)

Saves a previously uploaded document as PDF, and sends the file data in a JSON-serialized object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveDocumentParameters** | [**PdfSaveDocumentParameters**](PdfSaveDocumentParameters.md)| A PdfSaveDocumentParameters object specifying the parameters of the action. | 

### Return type

[**PdfSaveDocumentResponse**](PdfSaveDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveDocumentToFile

> *os.File SaveDocumentToFile(ctx, pdfSaveDocumentParameters)

Saves a previously uploaded document as PDF, and streams the file binary data to the response (this is the most efficient download method).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSaveDocumentParameters** | [**PdfSaveDocumentParameters**](PdfSaveDocumentParameters.md)| A PdfSaveDocumentParameters object specifying the parameters of the action. | 

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


## ScalePage

> PdfScalePageResponse ScalePage(ctx, pdfScalePageParameters)

Scales a page range from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfScalePageParameters** | [**PdfScalePageParameters**](PdfScalePageParameters.md)| A PdfScalePage object specifying the parameters of the action. | 

### Return type

[**PdfScalePageResponse**](PdfScalePageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInfo

> PdfSetInfoResponse SetInfo(ctx, pdfSetInfoParameters)

Sets information to a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSetInfoParameters** | [**PdfSetInfoParameters**](PdfSetInfoParameters.md)| A PdfSetInfoParameters object specifying the parameters of the action. | 

### Return type

[**PdfSetInfoResponse**](PdfSetInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInitialView

> PdfSetInitialViewResponse SetInitialView(ctx, pdfSetInitialViewParameters)

Sets various document level initial view options to a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSetInitialViewParameters** | [**PdfSetInitialViewParameters**](PdfSetInitialViewParameters.md)| A PdfsetInitialViewParameters object specifying the parameters of the action. | 

### Return type

[**PdfSetInitialViewResponse**](PdfSetInitialViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPageBox

> PdfSetPageBoxResponse SetPageBox(ctx, pdfSetPageBoxParameters)

Sets pagebox to a page range from previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSetPageBoxParameters** | [**PdfSetPageBoxParameters**](PdfSetPageBoxParameters.md)| A PdfSetPageBoxParameters object specifying the parameters of the action. | 

### Return type

[**PdfSetPageBoxResponse**](PdfSetPageBoxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPassword

> PdfSetPasswordResponse SetPassword(ctx, pdfSetPasswordParameters)

Unprotects a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSetPasswordParameters** | [**PdfSetPasswordParameters**](PdfSetPasswordParameters.md)| A PdfSetPasswordParameters object specifying the parameters of the action. | 

### Return type

[**PdfSetPasswordResponse**](PdfSetPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Split

> PdfSplitResponse Split(ctx, pdfSplitParameters)

Splits a previously uploaded document into new ones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSplitParameters** | [**PdfSplitParameters**](PdfSplitParameters.md)| A PdfSplitParameters object specifying the parameters of the action. | 

### Return type

[**PdfSplitResponse**](PdfSplitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapPages

> PdfSwapPagesResponse SwapPages(ctx, pdfSwapPagesParameters)

Swaps two pages from a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfSwapPagesParameters** | [**PdfSwapPagesParameters**](PdfSwapPagesParameters.md)| A PdfSwapPagesParameters object specifying the parameters of the action. | 

### Return type

[**PdfSwapPagesResponse**](PdfSwapPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unprotect

> PdfUnprotectResponse Unprotect(ctx, pdfUnprotectParameters)

Unprotects a previously uploaded document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdfUnprotectParameters** | [**PdfUnprotectParameters**](PdfUnprotectParameters.md)| A PdfUnprotectParameters object specifying the parameters of the action. | 

### Return type

[**PdfUnprotectResponse**](PdfUnprotectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

