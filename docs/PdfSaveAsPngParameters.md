# PdfSaveAsPngParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | Pointer to **string** | Specifies the number of the page, or the range of pages to be saved as PNG. | [optional] [default to *]
**Compression** | **int32** | Specifies the level of compression to be used for the PNG output, between 0 (no compression - faster encoding) to 9(max compression - slower encoding). | [optional] [default to 6]
**Interlaced** | **bool** | Specifies if the produced PNG image must be interlaced. | [optional] [default to false]
**Resolution** | **float32** | Specifies the resolution to be used for the rendition process. | [optional] [default to 200]
**RenderFormFields** | **bool** | Specifies whether the form fields of the PDF shall be rendered. | [optional] [default to false]
**KeepRasterPDFResolution** | **bool** | Specifies if the initial image resolution must be kept in case of raster-pdf processing. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


