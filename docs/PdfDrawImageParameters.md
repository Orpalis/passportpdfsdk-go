# PdfDrawImageParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the page or the page range on which the image shall be drawn. | 
**ImageData** | Pointer to **string** | Specifies the data of the image to be drawn. | [optional] 
**ImageFileId** | Pointer to **string** | Specifies the file ID of the image to be drawn. | [optional] [default to ]
**Quality** | **int32** | Specifies the level of quality to be used for the compression, from 1 (poorest) to 100 (greatest). | [optional] [default to 75]
**ColorImageCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**BitonalCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**ImageLayout** | [**DrawableContentLayoutParameters**](DrawableContentLayoutParameters.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


