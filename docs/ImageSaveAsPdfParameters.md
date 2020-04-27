# ImageSaveAsPdfParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | Pointer to **string** | Specifies the number of the page, or the range of pages to be saved as PDF. | [optional] [default to *]
**Conformance** | [**PdfConformance**](PdfConformance.md) |  | [optional] 
**ColorImageCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**BitonalImageCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**EnableColorDetection** | **bool** | Specifies is color detection must be used. | [optional] [default to false]
**ImageQuality** | **int32** | Specifies the quality to be used for the compression of the images from the PDF.  Must be in the range [0 (best compression - worst quality) - 100 (worst quality - best compression)]. | [optional] [default to 75]
**DownscaleResolution** | **int32** | Specifies the resolution for downscaling images, if any. | [optional] [default to 0]
**FastWebView** | **bool** | Specifies whether the PDF shall be optimized for online distribution. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


