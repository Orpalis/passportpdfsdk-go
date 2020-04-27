# PdfLoadDocumentFromByteArrayParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Specifies the data of the document. | 
**FileName** | Pointer to **string** | Specifies the name of the document. | [optional] 
**Password** | Pointer to **string** | Specifies the password of the document. | [optional] 
**Conformance** | [**PdfConformance**](PdfConformance.md) |  | [optional] 
**ContentEncoding** | [**ContentEncoding**](ContentEncoding.md) |  | [optional] 
**EnableColorDetection** | **bool** | Specifies whether color detection must be used while importing a raster format to PDF. | [optional] [default to false]
**GetPreview** | **bool** | Specifies whether the response must contain a thumbnail of the first page of the document. | [optional] [default to false]
**ThumbnailWidth** | **int32** | Specifies, in pixels, the width of the thumbnail to be retrieved. Only applicable if GetPreview has been set to true. | [optional] [default to 140]
**ThumbnailHeight** | **int32** | Specifies, in pixels, the height of the thumbnail to be retrieved.  Only applicable if GetPreview has been set to true. | [optional] [default to 220]
**ThumbnailBackgroundColor** | Pointer to **string** | Specifies the background color of the thumbnail, using the color name (ie: \&quot;red\&quot;) or its RGBa code (ie: \&quot;rgba(255,0,0,1)\&quot;).   Only applicable if GetPreview has been set to true. | [optional] [default to rgba(0,0,0,0)]
**ThumbnailFitToPageSize** | **bool** | Specifies if the size of the produced thumbnail is automatically adjusted to don&#39;t have any margin.  Only applicable if GetPreview has been set to true. | [optional] [default to true]
**TxtPageWidth** | **float32** | Specifies the page width, in points, of produced documents from txt files. | [optional] [default to 595]
**TxtPageHeight** | **float32** | Specifies the page height, in points, of produced documents from txt files. | [optional] [default to 842]
**TxtPageMarginLeft** | **float32** | Specifies the page margin left, in points, of produced documents from txt files. | [optional] [default to 10]
**TxtPageMarginTop** | **float32** | Specifies the page margin top, in points, of produced documents from txt files. | [optional] [default to 10]
**TxtPageMarginRight** | **float32** | Specifies the page margin right, in points, of produced documents from txt files. | [optional] [default to 10]
**TxtPageMarginBottom** | **float32** | Specifies the page margin bottom, in points, of produced documents from txt files. | [optional] [default to 10]
**TxtHorizontalTextAlignment** | [**TextAlignment**](TextAlignment.md) |  | [optional] 
**TxtFontSize** | **float32** | Specifies the text size, in points, to be used for producing documents from txt files. | [optional] [default to 12]
**TxtFontFamily** | Pointer to **string** | Specifies the name of the font to be used for producing documents from txt files. | [optional] [default to Arial Unicode MS]
**TxtFontBold** | **bool** | Specifies whether the font to be used for producing documents from txt files must have a bold style. | [optional] [default to false]
**TxtFontItalic** | **bool** | Specifies whether the font to be used for producing documents from txt files must have an italic style. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


