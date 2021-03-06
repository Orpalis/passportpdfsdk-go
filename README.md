# Go API client for openapi

Another brick in the cloud

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://passportpdfapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigApi* | [**ConfigGetAPIVersion**](docs/ConfigApi.md#configgetapiversion) | **Get** /api/config/ConfigGetAPIVersion | 
*ConfigApi* | [**ConfigGetMaxAllowedContentLength**](docs/ConfigApi.md#configgetmaxallowedcontentlength) | **Get** /api/config/ConfigGetMaxAllowedContentLength | Gets the maximal length of a request content, in bytes.
*ConfigApi* | [**ConfigGetSuggestedClientTimeout**](docs/ConfigApi.md#configgetsuggestedclienttimeout) | **Get** /api/config/ConfigGetSuggestedClientTimeout | Gets the suggested client API timeout, in milliseconds.
*ConfigApi* | [**ConfigGetSuggestedMaxClientThreads**](docs/ConfigApi.md#configgetsuggestedmaxclientthreads) | **Get** /api/config/ConfigGetSuggestedMaxClientThreads | Gets the suggested maximum number of threads to be used simultaneously by a client application.
*ConfigApi* | [**ConfigGetSupportedFonts**](docs/ConfigApi.md#configgetsupportedfonts) | **Get** /api/config/ConfigGetSupportedFonts | Gets the list of supported fonts for text drawing operations.
*ConfigApi* | [**ConfigGetSupportedOCRLanguages**](docs/ConfigApi.md#configgetsupportedocrlanguages) | **Get** /api/config/ConfigGetSupportedOCRLanguages | Gets the list of supported languages for OCR.
*DocuViewareApi* | [**DocuViewareGetControl**](docs/DocuViewareApi.md#docuviewaregetcontrol) | **Post** /api/docuvieware/DocuViewareGetControl | Gets the HTML dom of a DocuVieware control.
*DocuViewareApi* | [**DocuViewareGetVersion**](docs/DocuViewareApi.md#docuviewaregetversion) | **Get** /api/docuvieware/DocuViewareGetVersion | Get the DocuVieware engine version.
*DocuViewareApi* | [**DocuViewareSave**](docs/DocuViewareApi.md#docuviewaresave) | **Post** /api/docuvieware/DocuViewareSave | Saves the document being handled by a specific DocuVieware control, in its current state.
*DocumentApi* | [**DocumentClose**](docs/DocumentApi.md#documentclose) | **Post** /api/document/DocumentClose | Closes a previously uploaded document.
*DocumentApi* | [**DocumentGetPreview**](docs/DocumentApi.md#documentgetpreview) | **Post** /api/document/DocumentGetPreview | Gets the format, the page count and a thumbnail of a previously uploaded document.
*DocumentApi* | [**DocumentLoad**](docs/DocumentApi.md#documentload) | **Post** /api/document/DocumentLoad | Loads the provided document file.
*DocumentApi* | [**DocumentLoadFromURI**](docs/DocumentApi.md#documentloadfromuri) | **Post** /api/document/DocumentLoadFromURI | Loads the provided document file from an URI.
*DocumentApi* | [**DocumentLoadMultipart**](docs/DocumentApi.md#documentloadmultipart) | **Post** /api/document/DocumentLoadMultipart | Loads the provided document file using Multipart Upload.
*ImageApi* | [**ImageAdjust**](docs/ImageApi.md#imageadjust) | **Post** /api/image/ImageAdjust | Adjusts a previously uploaded image.
*ImageApi* | [**ImageAutoCrop**](docs/ImageApi.md#imageautocrop) | **Post** /api/image/ImageAutoCrop | Automatically crops a previously uploaded image.
*ImageApi* | [**ImageCleanupDocument**](docs/ImageApi.md#imagecleanupdocument) | **Post** /api/image/ImageCleanupDocument | Cleanup a previously uploaded image.
*ImageApi* | [**ImageCloneRegions**](docs/ImageApi.md#imagecloneregions) | **Post** /api/image/ImageCloneRegions | Clones regions from a previously uploaded image into new images.
*ImageApi* | [**ImageClose**](docs/ImageApi.md#imageclose) | **Post** /api/image/ImageClose | Closes a previously uploaded image.
*ImageApi* | [**ImageConvertColorDepth**](docs/ImageApi.md#imageconvertcolordepth) | **Post** /api/image/ImageConvertColorDepth | Converts the color depth of a previously uploaded image.
*ImageApi* | [**ImageCrop**](docs/ImageApi.md#imagecrop) | **Post** /api/image/ImageCrop | Crops a previously uploaded image.
*ImageApi* | [**ImageDeletePage**](docs/ImageApi.md#imagedeletepage) | **Post** /api/image/ImageDeletePage | Deletes a page range from a previously uploaded image.
*ImageApi* | [**ImageDetectBlankPages**](docs/ImageApi.md#imagedetectblankpages) | **Post** /api/image/ImageDetectBlankPages | Detects the blank page(s) from a previously uploaded image and offers to remove them.
*ImageApi* | [**ImageDetectColor**](docs/ImageApi.md#imagedetectcolor) | **Post** /api/image/ImageDetectColor | Performs color detection  on a previously uploaded image.
*ImageApi* | [**ImageDetectPageOrientation**](docs/ImageApi.md#imagedetectpageorientation) | **Post** /api/image/ImageDetectPageOrientation | Detects the orientation of the page(s) of a previously uploaded image and offers to automatically rotate them.
*ImageApi* | [**ImageFilter**](docs/ImageApi.md#imagefilter) | **Post** /api/image/ImageFilter | Applies filters to a previously uploaded image.
*ImageApi* | [**ImageGetPageThumbnail**](docs/ImageApi.md#imagegetpagethumbnail) | **Post** /api/image/ImageGetPageThumbnail | Gets a thumbnail of each page within the provided page range from a previously uploaded image.
*ImageApi* | [**ImageGetSupportedFileExtensions**](docs/ImageApi.md#imagegetsupportedfileextensions) | **Get** /api/image/ImageGetSupportedFileExtensions | Gets the supported file extensions by the image loading actions.
*ImageApi* | [**ImageLoad**](docs/ImageApi.md#imageload) | **Post** /api/image/ImageLoad | Loads the provided image file.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.
*ImageApi* | [**ImageLoadMultipart**](docs/ImageApi.md#imageloadmultipart) | **Post** /api/image/ImageLoadMultipart | Loads the provided image file using Multipart Upload.  Supported image formats can be retrieved by the GetSupportedImageFileExtensions action.
*ImageApi* | [**ImageMICR**](docs/ImageApi.md#imagemicr) | **Post** /api/image/ImageMICR | Performs MICR (Magnetic Ink Character Recognition) on a previously uploaded image.
*ImageApi* | [**ImageReadBarcodes**](docs/ImageApi.md#imagereadbarcodes) | **Post** /api/image/ImageReadBarcodes | Reads barcodes from a previously uploaded image.
*ImageApi* | [**ImageResize**](docs/ImageApi.md#imageresize) | **Post** /api/image/ImageResize | Resizes a previously uploaded image.
*ImageApi* | [**ImageRotate**](docs/ImageApi.md#imagerotate) | **Post** /api/image/ImageRotate | Rotates and/or flips a previously uploaded image.
*ImageApi* | [**ImageSaveAsJPEG**](docs/ImageApi.md#imagesaveasjpeg) | **Post** /api/image/ImageSaveAsJPEG | Saves a previously uploaded image as JPEG, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsJPEGFile**](docs/ImageApi.md#imagesaveasjpegfile) | **Post** /api/image/ImageSaveAsJPEGFile | Saves a previously uploaded image as JPEG, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSaveAsPDF**](docs/ImageApi.md#imagesaveaspdf) | **Post** /api/image/ImageSaveAsPDF | Saves a previously uploaded image as PDF, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsPDFFile**](docs/ImageApi.md#imagesaveaspdffile) | **Post** /api/image/ImageSaveAsPDFFile | Saves a previously uploaded image as PDF, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSaveAsPDFMRC**](docs/ImageApi.md#imagesaveaspdfmrc) | **Post** /api/image/ImageSaveAsPDFMRC | Saves a previously uploaded image as PDF using MRC compression, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsPDFMRCFile**](docs/ImageApi.md#imagesaveaspdfmrcfile) | **Post** /api/image/ImageSaveAsPDFMRCFile | Saves a previously uploaded image as PDF using MRC compression, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSaveAsPNG**](docs/ImageApi.md#imagesaveaspng) | **Post** /api/image/ImageSaveAsPNG | Saves a previously uploaded image as PNG, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsPNGFile**](docs/ImageApi.md#imagesaveaspngfile) | **Post** /api/image/ImageSaveAsPNGFile | Saves a previously uploaded image as PNG, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSaveAsTIFF**](docs/ImageApi.md#imagesaveastiff) | **Post** /api/image/ImageSaveAsTIFF | Saves a previously uploaded image as TIFF, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsTIFFFile**](docs/ImageApi.md#imagesaveastifffile) | **Post** /api/image/ImageSaveAsTIFFFile | Saves a previously uploaded image as TIFF, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSaveAsTIFFMultipage**](docs/ImageApi.md#imagesaveastiffmultipage) | **Post** /api/image/ImageSaveAsTIFFMultipage | Saves a previously uploaded image as multipage TIFF, and sends the file data in a JSON-serialized object.
*ImageApi* | [**ImageSaveAsTIFFMultipageFile**](docs/ImageApi.md#imagesaveastiffmultipagefile) | **Post** /api/image/ImageSaveAsTIFFMultipageFile | Saves a previously uploaded image as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).
*ImageApi* | [**ImageSwapPages**](docs/ImageApi.md#imageswappages) | **Post** /api/image/ImageSwapPages | Swaps two pages from a previously uploaded image.
*PDFApi* | [**Annotate**](docs/PDFApi.md#annotate) | **Post** /api/pdf/Annotate | Annotates a previously uploaded document.
*PDFApi* | [**AutoDeskew**](docs/PDFApi.md#autodeskew) | **Post** /api/pdf/AutoDeskew | Performs auto deskew on a page range of a previously uploaded document.
*PDFApi* | [**ClearPage**](docs/PDFApi.md#clearpage) | **Post** /api/pdf/ClearPage | Clears a page range from a previously uploaded document.
*PDFApi* | [**ClonePage**](docs/PDFApi.md#clonepage) | **Post** /api/pdf/ClonePage | Clones specific pages from a previously uploaded document to another previously uploaded document.
*PDFApi* | [**ClosePDF**](docs/PDFApi.md#closepdf) | **Post** /api/pdf/ClosePDF | Closes a previously uploaded document.
*PDFApi* | [**ConvertToPDFA**](docs/PDFApi.md#converttopdfa) | **Post** /api/pdf/ConvertToPDFA | Converts a previously uploaded document to PDF/A.
*PDFApi* | [**DeletePage**](docs/PDFApi.md#deletepage) | **Post** /api/pdf/DeletePage | Deletes a page range from a previously uploaded document.
*PDFApi* | [**DetectPageOrientation**](docs/PDFApi.md#detectpageorientation) | **Post** /api/pdf/DetectPageOrientation | Detects the orientation of the page(s) of a previously uploaded document and offers to automatically rotate them.
*PDFApi* | [**DigiSign**](docs/PDFApi.md#digisign) | **Post** /api/pdf/DigiSign | Signs a previously uploaded document digitally.
*PDFApi* | [**DrawImage**](docs/PDFApi.md#drawimage) | **Post** /api/pdf/DrawImage | Draws an image on a page range of a previously uploaded document.
*PDFApi* | [**ExtractPage**](docs/PDFApi.md#extractpage) | **Post** /api/pdf/ExtractPage | Extracts a page range from a previously uploaded document into one or several new documents.
*PDFApi* | [**ExtractText**](docs/PDFApi.md#extracttext) | **Post** /api/pdf/ExtractText | Extracts text from the document pages.
*PDFApi* | [**Flatten**](docs/PDFApi.md#flatten) | **Post** /api/pdf/Flatten | Flattens the form-fields, annotations, and/or the layers of a previously uploaded document.
*PDFApi* | [**GetInfo**](docs/PDFApi.md#getinfo) | **Post** /api/pdf/GetInfo | Gets information about a previously uploaded document.
*PDFApi* | [**GetPDFImportSupportedFileExtensions**](docs/PDFApi.md#getpdfimportsupportedfileextensions) | **Get** /api/pdf/GetPDFImportSupportedFileExtensions | Gets the supported file extensions by the LoadDocumentAsPDF action.
*PDFApi* | [**GetPageThumbnail**](docs/PDFApi.md#getpagethumbnail) | **Post** /api/pdf/GetPageThumbnail | Gets a thumbnail of each page within the provided page range from a previously uploaded document.
*PDFApi* | [**InsertImage**](docs/PDFApi.md#insertimage) | **Post** /api/pdf/InsertImage | Inserts an image on a new page of a previously uploaded document.
*PDFApi* | [**InsertNewPage**](docs/PDFApi.md#insertnewpage) | **Post** /api/pdf/InsertNewPage | Inserts one or more new blank pages to a specific position in a previously uploaded document.
*PDFApi* | [**InsertPageNumber**](docs/PDFApi.md#insertpagenumber) | **Post** /api/pdf/InsertPageNumber | Inserts page number(s) on a document.
*PDFApi* | [**InsertText**](docs/PDFApi.md#inserttext) | **Post** /api/pdf/InsertText | Inserts text on a document.
*PDFApi* | [**Linearize**](docs/PDFApi.md#linearize) | **Post** /api/pdf/Linearize | Linearizes a previously uploaded document.
*PDFApi* | [**LoadDocumentAsPDF**](docs/PDFApi.md#loaddocumentaspdf) | **Post** /api/pdf/LoadDocument | Imports the provided document as PDF.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.
*PDFApi* | [**LoadDocumentAsPDFMultipart**](docs/PDFApi.md#loaddocumentaspdfmultipart) | **Post** /api/pdf/LoadDocumentAsPDFMultipart | Imports the provided document as PDF using Multipart Upload.  Supported document formats can be retrieved by the GetPDFImportSupportedFileExtensions action.
*PDFApi* | [**Merge**](docs/PDFApi.md#merge) | **Post** /api/pdf/Merge | Merges several previously uploaded documents into a new PDF.
*PDFApi* | [**MergePages**](docs/PDFApi.md#mergepages) | **Post** /api/pdf/MergePages | Merges multiple pages, vertically, within a previously uploaded document into one single page.
*PDFApi* | [**MovePage**](docs/PDFApi.md#movepage) | **Post** /api/pdf/MovePage | Moves a page range from a previously uploaded document.
*PDFApi* | [**OCR**](docs/PDFApi.md#ocr) | **Post** /api/pdf/OCR | Performs optical character recognition on a page range of a previously uploaded document.  The recognized text is added as invisible text on each processed page.  No token is charged for blank pages.
*PDFApi* | [**Protect**](docs/PDFApi.md#protect) | **Post** /api/pdf/Protect | Protects a previously uploaded document.
*PDFApi* | [**ReadBarcodes**](docs/PDFApi.md#readbarcodes) | **Post** /api/pdf/ReadBarcodes | Reads barcodes from a previously uploaded document.
*PDFApi* | [**Reduce**](docs/PDFApi.md#reduce) | **Post** /api/pdf/Reduce | Reduces the size of a previously uploaded document.
*PDFApi* | [**RemovePageFormFields**](docs/PDFApi.md#removepageformfields) | **Post** /api/pdf/RemovePageFormFields | Removes the form fields from a page range of a previously uploaded document.
*PDFApi* | [**RemoveText**](docs/PDFApi.md#removetext) | **Post** /api/pdf/RemoveText | Removes text (all text or only invisible one) from a previously uploaded PDF.
*PDFApi* | [**ReorderPages**](docs/PDFApi.md#reorderpages) | **Post** /api/pdf/ReorderPages | Reorders pages of a previously uploaded document.
*PDFApi* | [**RepairDocument**](docs/PDFApi.md#repairdocument) | **Post** /api/pdf/RepairDocument | Repairs a previously uploaded PDF document.
*PDFApi* | [**RotatePageStandard**](docs/PDFApi.md#rotatepagestandard) | **Post** /api/pdf/RotatePageStandard | Rotates (standardly) a page range from a previously uploaded document.
*PDFApi* | [**SaveAsJPEG**](docs/PDFApi.md#saveasjpeg) | **Post** /api/pdf/SaveAsJPEG | Saves a previously uploaded document as JPEG, and sends the file data in a JSON-serialized object.
*PDFApi* | [**SaveAsJPEGFile**](docs/PDFApi.md#saveasjpegfile) | **Post** /api/pdf/SaveAsJPEGFile | Saves a previously uploaded document as JPEG, and streams the file binary data to the response (this is the most efficient download method).
*PDFApi* | [**SaveAsPNG**](docs/PDFApi.md#saveaspng) | **Post** /api/pdf/SaveAsPNG | Saves a previously uploaded document as PNG, and sends the file data in a JSON-serialized object.
*PDFApi* | [**SaveAsPNGFile**](docs/PDFApi.md#saveaspngfile) | **Post** /api/pdf/SaveAsPNGFile | Saves a previously uploaded document as PNG, and streams the file binary data to the response (this is the most efficient download method).
*PDFApi* | [**SaveAsTIFF**](docs/PDFApi.md#saveastiff) | **Post** /api/pdf/SaveAsTIFF | Saves a previously uploaded document as TIFF, and sends the file data in a JSON-serialized object.
*PDFApi* | [**SaveAsTIFFFile**](docs/PDFApi.md#saveastifffile) | **Post** /api/pdf/SaveAsTIFFFile | Saves a previously uploaded document as TIFF, and streams the file binary data to the response (this is the most efficient download method).
*PDFApi* | [**SaveAsTIFFMultipage**](docs/PDFApi.md#saveastiffmultipage) | **Post** /api/pdf/SaveAsTIFFMultipage | Saves a previously uploaded document as multipage TIFF, and sends the file data in a JSON-serialized object.
*PDFApi* | [**SaveAsTIFFMultipageFile**](docs/PDFApi.md#saveastiffmultipagefile) | **Post** /api/pdf/SaveAsTIFFMultipageFile | Saves a previously uploaded document as multipage TIFF, and streams the file binary data to the response (this is the most efficient download method).
*PDFApi* | [**SaveDocument**](docs/PDFApi.md#savedocument) | **Post** /api/pdf/SaveDocument | Saves a previously uploaded document as PDF, and sends the file data in a JSON-serialized object.
*PDFApi* | [**SaveDocumentToFile**](docs/PDFApi.md#savedocumenttofile) | **Post** /api/pdf/SaveDocumentToFile | Saves a previously uploaded document as PDF, and streams the file binary data to the response (this is the most efficient download method).
*PDFApi* | [**ScalePage**](docs/PDFApi.md#scalepage) | **Post** /api/pdf/ScalePage | Scales a page range from a previously uploaded document.
*PDFApi* | [**SetInfo**](docs/PDFApi.md#setinfo) | **Post** /api/pdf/SetInfo | Sets information to a previously uploaded document.
*PDFApi* | [**SetInitialView**](docs/PDFApi.md#setinitialview) | **Post** /api/pdf/SetInitialView | Sets various document level initial view options to a previously uploaded document.
*PDFApi* | [**SetPageBox**](docs/PDFApi.md#setpagebox) | **Post** /api/pdf/SetPageBox | Sets pagebox to a page range from previously uploaded document.
*PDFApi* | [**SetPassword**](docs/PDFApi.md#setpassword) | **Post** /api/pdf/SetPassword | Unprotects a previously uploaded document.
*PDFApi* | [**Split**](docs/PDFApi.md#split) | **Post** /api/pdf/Split | Splits a previously uploaded document into new ones.
*PDFApi* | [**SwapPages**](docs/PDFApi.md#swappages) | **Post** /api/pdf/SwapPages | Swaps two pages from a previously uploaded document.
*PDFApi* | [**Unprotect**](docs/PDFApi.md#unprotect) | **Post** /api/pdf/Unprotect | Unprotects a previously uploaded document.
*PassportManagerApi* | [**PassportManagerGetPassportInfo**](docs/PassportManagerApi.md#passportmanagergetpassportinfo) | **Get** /api/passportmanager/PassportManagerGetPassportInfo | 
*PassportPDFApplicationManagerApi* | [**PassportPDFApplicationManagerGetApplicationDownloadLink**](docs/PassportPDFApplicationManagerApi.md#passportpdfapplicationmanagergetapplicationdownloadlink) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationDownloadLink | 
*PassportPDFApplicationManagerApi* | [**PassportPDFApplicationManagerGetApplicationLatestVersion**](docs/PassportPDFApplicationManagerApi.md#passportpdfapplicationmanagergetapplicationlatestversion) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationLatestVersion | 
*PassportPDFApplicationManagerApi* | [**PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion**](docs/PassportPDFApplicationManagerApi.md#passportpdfapplicationmanagergetapplicationminimumsupportedversion) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetApplicationMinimumSupportedVersion | 
*PassportPDFApplicationManagerApi* | [**PassportPDFApplicationManagerGetMaxClientThreads**](docs/PassportPDFApplicationManagerApi.md#passportpdfapplicationmanagergetmaxclientthreads) | **Get** /api/passportpdfapplicationmanager/PassportPDFApplicationManagerGetMaxClientThreads | Gets the maximum number of threads to be used simultaneously by a client application.


## Documentation For Models

 - [AnnotationType](docs/AnnotationType.md)
 - [Barcode1DSymbology](docs/Barcode1DSymbology.md)
 - [BarcodeInfo](docs/BarcodeInfo.md)
 - [BarcodeType](docs/BarcodeType.md)
 - [ClonedImageRegion](docs/ClonedImageRegion.md)
 - [ColorDepthConversionMethod](docs/ColorDepthConversionMethod.md)
 - [ColorDetectionResult](docs/ColorDetectionResult.md)
 - [ContentEncoding](docs/ContentEncoding.md)
 - [ContentHorizontalPosition](docs/ContentHorizontalPosition.md)
 - [ContentSize](docs/ContentSize.md)
 - [ContentVerticalPosition](docs/ContentVerticalPosition.md)
 - [CropUnit](docs/CropUnit.md)
 - [DocuViewareCertificate](docs/DocuViewareCertificate.md)
 - [DocuViewareControlState](docs/DocuViewareControlState.md)
 - [DocuViewareGetControlParameters](docs/DocuViewareGetControlParameters.md)
 - [DocuViewareGetControlResponse](docs/DocuViewareGetControlResponse.md)
 - [DocuViewareLocale](docs/DocuViewareLocale.md)
 - [DocuViewareSaveParameters](docs/DocuViewareSaveParameters.md)
 - [DocuViewareSaveResponse](docs/DocuViewareSaveResponse.md)
 - [DocumentAlignment](docs/DocumentAlignment.md)
 - [DocumentCloseParameters](docs/DocumentCloseParameters.md)
 - [DocumentCloseResponse](docs/DocumentCloseResponse.md)
 - [DocumentFormat](docs/DocumentFormat.md)
 - [DocumentGetPreviewResponse](docs/DocumentGetPreviewResponse.md)
 - [DocumentLoadResponse](docs/DocumentLoadResponse.md)
 - [DocumentPosition](docs/DocumentPosition.md)
 - [DrawableContentLayoutParameters](docs/DrawableContentLayoutParameters.md)
 - [EncryptionAlgorithm](docs/EncryptionAlgorithm.md)
 - [Error](docs/Error.md)
 - [Font](docs/Font.md)
 - [FontStyle](docs/FontStyle.md)
 - [FreeTextAnnotationParameters](docs/FreeTextAnnotationParameters.md)
 - [GetDocumentPreviewParameters](docs/GetDocumentPreviewParameters.md)
 - [ImageAdjustParameters](docs/ImageAdjustParameters.md)
 - [ImageAdjustResponse](docs/ImageAdjustResponse.md)
 - [ImageAutoCropParameters](docs/ImageAutoCropParameters.md)
 - [ImageAutoCropResponse](docs/ImageAutoCropResponse.md)
 - [ImageCleanupDocumentParameters](docs/ImageCleanupDocumentParameters.md)
 - [ImageCleanupDocumentResponse](docs/ImageCleanupDocumentResponse.md)
 - [ImageCloneRegionsParameters](docs/ImageCloneRegionsParameters.md)
 - [ImageCloneRegionsResponse](docs/ImageCloneRegionsResponse.md)
 - [ImageCloseParameters](docs/ImageCloseParameters.md)
 - [ImageCloseResponse](docs/ImageCloseResponse.md)
 - [ImageConvertColorDepthParameters](docs/ImageConvertColorDepthParameters.md)
 - [ImageConvertColorDepthResponse](docs/ImageConvertColorDepthResponse.md)
 - [ImageCropParameters](docs/ImageCropParameters.md)
 - [ImageCropResponse](docs/ImageCropResponse.md)
 - [ImageDeletePageParameters](docs/ImageDeletePageParameters.md)
 - [ImageDeletePageResponse](docs/ImageDeletePageResponse.md)
 - [ImageDetectBlankPagesParameters](docs/ImageDetectBlankPagesParameters.md)
 - [ImageDetectBlankPagesResponse](docs/ImageDetectBlankPagesResponse.md)
 - [ImageDetectColorParameters](docs/ImageDetectColorParameters.md)
 - [ImageDetectColorResponse](docs/ImageDetectColorResponse.md)
 - [ImageDetectPageOrientationParameters](docs/ImageDetectPageOrientationParameters.md)
 - [ImageDetectPageOrientationResponse](docs/ImageDetectPageOrientationResponse.md)
 - [ImageFilterParameters](docs/ImageFilterParameters.md)
 - [ImageFilterResponse](docs/ImageFilterResponse.md)
 - [ImageFilters](docs/ImageFilters.md)
 - [ImageGetPageThumbnailParameters](docs/ImageGetPageThumbnailParameters.md)
 - [ImageGetPageThumbnailResponse](docs/ImageGetPageThumbnailResponse.md)
 - [ImageLoadResponse](docs/ImageLoadResponse.md)
 - [ImageMicrParameters](docs/ImageMicrParameters.md)
 - [ImageMicrResponse](docs/ImageMicrResponse.md)
 - [ImageQuality](docs/ImageQuality.md)
 - [ImageReadBarcodesParameters](docs/ImageReadBarcodesParameters.md)
 - [ImageRegion](docs/ImageRegion.md)
 - [ImageResizeParameters](docs/ImageResizeParameters.md)
 - [ImageResizeResponse](docs/ImageResizeResponse.md)
 - [ImageRotateParameters](docs/ImageRotateParameters.md)
 - [ImageRotateResponse](docs/ImageRotateResponse.md)
 - [ImageSaveAsJpegParameters](docs/ImageSaveAsJpegParameters.md)
 - [ImageSaveAsJpegResponse](docs/ImageSaveAsJpegResponse.md)
 - [ImageSaveAsPdfParameters](docs/ImageSaveAsPdfParameters.md)
 - [ImageSaveAsPdfResponse](docs/ImageSaveAsPdfResponse.md)
 - [ImageSaveAsPdfmrcParameters](docs/ImageSaveAsPdfmrcParameters.md)
 - [ImageSaveAsPdfmrcResponse](docs/ImageSaveAsPdfmrcResponse.md)
 - [ImageSaveAsPngParameters](docs/ImageSaveAsPngParameters.md)
 - [ImageSaveAsPngResponse](docs/ImageSaveAsPngResponse.md)
 - [ImageSaveAsTiffMultipageParameters](docs/ImageSaveAsTiffMultipageParameters.md)
 - [ImageSaveAsTiffMultipageResponse](docs/ImageSaveAsTiffMultipageResponse.md)
 - [ImageSaveAsTiffParameters](docs/ImageSaveAsTiffParameters.md)
 - [ImageSaveAsTiffResponse](docs/ImageSaveAsTiffResponse.md)
 - [ImageSwapPagesParameters](docs/ImageSwapPagesParameters.md)
 - [ImageSwapPagesResponse](docs/ImageSwapPagesResponse.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [IntegerResponse](docs/IntegerResponse.md)
 - [LineAnnotationEndingStyle](docs/LineAnnotationEndingStyle.md)
 - [LineAnnotationParameters](docs/LineAnnotationParameters.md)
 - [LinkAnnotationClickBehaviour](docs/LinkAnnotationClickBehaviour.md)
 - [LinkAnnotationParameters](docs/LinkAnnotationParameters.md)
 - [LinkAnnotationType](docs/LinkAnnotationType.md)
 - [LoadDocumentFromByteArrayParameters](docs/LoadDocumentFromByteArrayParameters.md)
 - [LoadDocumentFromUriParameters](docs/LoadDocumentFromUriParameters.md)
 - [LoadDocumentParameters](docs/LoadDocumentParameters.md)
 - [LoadImageFromByteArrayParameters](docs/LoadImageFromByteArrayParameters.md)
 - [LoadImageParameters](docs/LoadImageParameters.md)
 - [LongResponse](docs/LongResponse.md)
 - [MicrContext](docs/MicrContext.md)
 - [MicrFont](docs/MicrFont.md)
 - [MicrResult](docs/MicrResult.md)
 - [MicrSymbolInfo](docs/MicrSymbolInfo.md)
 - [OutputIntent](docs/OutputIntent.md)
 - [PageImage](docs/PageImage.md)
 - [PageNumberFormat](docs/PageNumberFormat.md)
 - [PageOrientationInfo](docs/PageOrientationInfo.md)
 - [PageText](docs/PageText.md)
 - [PageViewMode](docs/PageViewMode.md)
 - [PassportPdfPassport](docs/PassportPdfPassport.md)
 - [PassportPdfStatus](docs/PassportPdfStatus.md)
 - [PdfAConformance](docs/PdfAConformance.md)
 - [PdfAlignedTextParameters](docs/PdfAlignedTextParameters.md)
 - [PdfAnnotateParameters](docs/PdfAnnotateParameters.md)
 - [PdfAnnotateResponse](docs/PdfAnnotateResponse.md)
 - [PdfAutoDeskewParameters](docs/PdfAutoDeskewParameters.md)
 - [PdfAutoDeskewResponse](docs/PdfAutoDeskewResponse.md)
 - [PdfClearPageParameters](docs/PdfClearPageParameters.md)
 - [PdfClearPageResponse](docs/PdfClearPageResponse.md)
 - [PdfClonePageParameters](docs/PdfClonePageParameters.md)
 - [PdfClonePageResponse](docs/PdfClonePageResponse.md)
 - [PdfCloseDocumentParameters](docs/PdfCloseDocumentParameters.md)
 - [PdfCloseDocumentResponse](docs/PdfCloseDocumentResponse.md)
 - [PdfConformance](docs/PdfConformance.md)
 - [PdfConvertToPdfaParameters](docs/PdfConvertToPdfaParameters.md)
 - [PdfConvertToPdfaResponse](docs/PdfConvertToPdfaResponse.md)
 - [PdfDeletePageParameters](docs/PdfDeletePageParameters.md)
 - [PdfDeletePageResponse](docs/PdfDeletePageResponse.md)
 - [PdfDetectPageOrientationParameters](docs/PdfDetectPageOrientationParameters.md)
 - [PdfDetectPageOrientationResponse](docs/PdfDetectPageOrientationResponse.md)
 - [PdfDigiSignParameters](docs/PdfDigiSignParameters.md)
 - [PdfDigiSignResponse](docs/PdfDigiSignResponse.md)
 - [PdfDrawImageParameters](docs/PdfDrawImageParameters.md)
 - [PdfDrawImageResponse](docs/PdfDrawImageResponse.md)
 - [PdfExtractPageParameters](docs/PdfExtractPageParameters.md)
 - [PdfExtractPageResponse](docs/PdfExtractPageResponse.md)
 - [PdfExtractTextMode](docs/PdfExtractTextMode.md)
 - [PdfExtractTextParameters](docs/PdfExtractTextParameters.md)
 - [PdfExtractTextResponse](docs/PdfExtractTextResponse.md)
 - [PdfFlattenParameters](docs/PdfFlattenParameters.md)
 - [PdfFlattenResponse](docs/PdfFlattenResponse.md)
 - [PdfGetInfoParameters](docs/PdfGetInfoParameters.md)
 - [PdfGetInfoResponse](docs/PdfGetInfoResponse.md)
 - [PdfGetPageThumbnailParameters](docs/PdfGetPageThumbnailParameters.md)
 - [PdfGetPageThumbnailResponse](docs/PdfGetPageThumbnailResponse.md)
 - [PdfImageCompressionScheme](docs/PdfImageCompressionScheme.md)
 - [PdfInitViewLayoutMode](docs/PdfInitViewLayoutMode.md)
 - [PdfInitViewNonFullScreenPageMode](docs/PdfInitViewNonFullScreenPageMode.md)
 - [PdfInitViewPageMode](docs/PdfInitViewPageMode.md)
 - [PdfInsertImageParameters](docs/PdfInsertImageParameters.md)
 - [PdfInsertImageResponse](docs/PdfInsertImageResponse.md)
 - [PdfInsertNewPageParameters](docs/PdfInsertNewPageParameters.md)
 - [PdfInsertNewPageResponse](docs/PdfInsertNewPageResponse.md)
 - [PdfInsertPageNumberParameters](docs/PdfInsertPageNumberParameters.md)
 - [PdfInsertPageNumberResponse](docs/PdfInsertPageNumberResponse.md)
 - [PdfInsertTextParameters](docs/PdfInsertTextParameters.md)
 - [PdfInsertTextResponse](docs/PdfInsertTextResponse.md)
 - [PdfLinearizeParameters](docs/PdfLinearizeParameters.md)
 - [PdfLinearizeResponse](docs/PdfLinearizeResponse.md)
 - [PdfLoadDocumentFromByteArrayParameters](docs/PdfLoadDocumentFromByteArrayParameters.md)
 - [PdfLoadDocumentParameters](docs/PdfLoadDocumentParameters.md)
 - [PdfLoadDocumentResponse](docs/PdfLoadDocumentResponse.md)
 - [PdfMergePagesParameters](docs/PdfMergePagesParameters.md)
 - [PdfMergePagesResponse](docs/PdfMergePagesResponse.md)
 - [PdfMergeParameters](docs/PdfMergeParameters.md)
 - [PdfMergeResponse](docs/PdfMergeResponse.md)
 - [PdfMovePageParameters](docs/PdfMovePageParameters.md)
 - [PdfMovePageResponse](docs/PdfMovePageResponse.md)
 - [PdfOcrParameters](docs/PdfOcrParameters.md)
 - [PdfOcrResponse](docs/PdfOcrResponse.md)
 - [PdfPageBox](docs/PdfPageBox.md)
 - [PdfProtectParameters](docs/PdfProtectParameters.md)
 - [PdfProtectResponse](docs/PdfProtectResponse.md)
 - [PdfReadBarcodesParameters](docs/PdfReadBarcodesParameters.md)
 - [PdfReduceParameters](docs/PdfReduceParameters.md)
 - [PdfReduceResponse](docs/PdfReduceResponse.md)
 - [PdfRemovePageFormFieldsParameters](docs/PdfRemovePageFormFieldsParameters.md)
 - [PdfRemovePageFormFieldsResponse](docs/PdfRemovePageFormFieldsResponse.md)
 - [PdfRemoveTextParameters](docs/PdfRemoveTextParameters.md)
 - [PdfRemoveTextResponse](docs/PdfRemoveTextResponse.md)
 - [PdfReorderPagesParameters](docs/PdfReorderPagesParameters.md)
 - [PdfReorderPagesResponse](docs/PdfReorderPagesResponse.md)
 - [PdfRepairDocumentParameters](docs/PdfRepairDocumentParameters.md)
 - [PdfRepairDocumentResponse](docs/PdfRepairDocumentResponse.md)
 - [PdfRotatePageStandardParameters](docs/PdfRotatePageStandardParameters.md)
 - [PdfRotatePageStandardResponse](docs/PdfRotatePageStandardResponse.md)
 - [PdfSaveAsJpegParameters](docs/PdfSaveAsJpegParameters.md)
 - [PdfSaveAsJpegResponse](docs/PdfSaveAsJpegResponse.md)
 - [PdfSaveAsPngParameters](docs/PdfSaveAsPngParameters.md)
 - [PdfSaveAsPngResponse](docs/PdfSaveAsPngResponse.md)
 - [PdfSaveAsTiffMultipageParameters](docs/PdfSaveAsTiffMultipageParameters.md)
 - [PdfSaveAsTiffMultipageResponse](docs/PdfSaveAsTiffMultipageResponse.md)
 - [PdfSaveAsTiffParameters](docs/PdfSaveAsTiffParameters.md)
 - [PdfSaveAsTiffResponse](docs/PdfSaveAsTiffResponse.md)
 - [PdfSaveDocumentParameters](docs/PdfSaveDocumentParameters.md)
 - [PdfSaveDocumentResponse](docs/PdfSaveDocumentResponse.md)
 - [PdfScalePageParameters](docs/PdfScalePageParameters.md)
 - [PdfScalePageResponse](docs/PdfScalePageResponse.md)
 - [PdfSetInfoParameters](docs/PdfSetInfoParameters.md)
 - [PdfSetInfoResponse](docs/PdfSetInfoResponse.md)
 - [PdfSetInitialViewParameters](docs/PdfSetInitialViewParameters.md)
 - [PdfSetInitialViewResponse](docs/PdfSetInitialViewResponse.md)
 - [PdfSetPageBoxParameters](docs/PdfSetPageBoxParameters.md)
 - [PdfSetPageBoxResponse](docs/PdfSetPageBoxResponse.md)
 - [PdfSetPasswordParameters](docs/PdfSetPasswordParameters.md)
 - [PdfSetPasswordResponse](docs/PdfSetPasswordResponse.md)
 - [PdfSplitMethod](docs/PdfSplitMethod.md)
 - [PdfSplitParameters](docs/PdfSplitParameters.md)
 - [PdfSplitResponse](docs/PdfSplitResponse.md)
 - [PdfSwapPagesParameters](docs/PdfSwapPagesParameters.md)
 - [PdfSwapPagesResponse](docs/PdfSwapPagesResponse.md)
 - [PdfTextParameters](docs/PdfTextParameters.md)
 - [PdfUnprotectParameters](docs/PdfUnprotectParameters.md)
 - [PdfUnprotectResponse](docs/PdfUnprotectResponse.md)
 - [PdfVersion](docs/PdfVersion.md)
 - [PrintQuality](docs/PrintQuality.md)
 - [ReadBarcodesResponse](docs/ReadBarcodesResponse.md)
 - [ReduceErrorCode](docs/ReduceErrorCode.md)
 - [ReduceErrorInfo](docs/ReduceErrorInfo.md)
 - [ReduceWarningCode](docs/ReduceWarningCode.md)
 - [ReduceWarningInfo](docs/ReduceWarningInfo.md)
 - [ResizeMode](docs/ResizeMode.md)
 - [ResizeUnit](docs/ResizeUnit.md)
 - [RotateFlip](docs/RotateFlip.md)
 - [RotateMode](docs/RotateMode.md)
 - [RubberStampAnnotationIcon](docs/RubberStampAnnotationIcon.md)
 - [RubberStampAnnotationParameters](docs/RubberStampAnnotationParameters.md)
 - [ScanMode](docs/ScanMode.md)
 - [SignatureCertificationLevel](docs/SignatureCertificationLevel.md)
 - [SignatureHash](docs/SignatureHash.md)
 - [SignatureMode](docs/SignatureMode.md)
 - [SquareAndCircleAnnotationParameters](docs/SquareAndCircleAnnotationParameters.md)
 - [StandardFontName](docs/StandardFontName.md)
 - [StickyNoteAnnotationIcon](docs/StickyNoteAnnotationIcon.md)
 - [StickyNoteAnnotationParameters](docs/StickyNoteAnnotationParameters.md)
 - [StringArrayResponse](docs/StringArrayResponse.md)
 - [StringResponse](docs/StringResponse.md)
 - [TextAlignment](docs/TextAlignment.md)
 - [TiffSaveCompression](docs/TiffSaveCompression.md)
 - [ToolbarStyle](docs/ToolbarStyle.md)
 - [ViewerZoomMode](docs/ViewerZoomMode.md)


## Documentation For Authorization



## PassportPDFAuthentication

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author



