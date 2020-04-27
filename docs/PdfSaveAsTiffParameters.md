# PdfSaveAsTiffParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | Pointer to **string** | Specifies the number of the page, or the range of pages to be saved as TIFF. | [optional] [default to *]
**Compression** | [**TiffSaveCompression**](TiffSaveCompression.md) |  | [optional] 
**JpegQuality** | **int32** | Specifies the level of quality to be used if JPEG compression is used, from 1 (poorest) to 100 (greatest). | [optional] [default to 75]
**UseCMYK** | **bool** | Specifies whether the TIFF shall be saved in CMYK color space or not. | [optional] [default to false]
**Resolution** | **float32** | Specifies the resolution to be used for the rendition process. | [optional] [default to 200]
**RenderFormFields** | **bool** | Specifies whether the form fields of the PDF shall be rendered. | [optional] [default to false]
**KeepRasterPDFResolution** | **bool** | Specifies whether the initial image resolution must be kept in case of raster-pdf processing. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


