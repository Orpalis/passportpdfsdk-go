# ImageSaveAsPdfmrcParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | Pointer to **string** | Specifies the number of the page, or the range of pages to be saved as PDF MRC. | [optional] [default to *]
**Conformance** | [**PdfConformance**](PdfConformance.md) |  | [optional] 
**ColorImageCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**BitonalImageCompression** | [**PdfImageCompressionScheme**](PdfImageCompressionScheme.md) |  | [optional] 
**ImageQuality** | **int32** | Specifies the quality to be used for the compression of the images from the PDF.  Must be in the range [0 (best compression - worst quality) - 100 (worst quality - best compression)]. | [optional] [default to 60]
**DownscaleResolution** | **int32** | Specifies the resolution for downscaling the background layer, if any. | [optional] [default to 100]
**PreserveSmoothing** | **bool** | Specifies whether the MRC engine should try to preserve smoothing between different layers.   Enabling this option should globally enhance the text quality but also reduce the compression rate. | [optional] [default to false]
**FastWebView** | **bool** | Specifies whether the PDF shall be optimized for online distribution. | [optional] [default to false]
**JBIG2PMSThreshold** | **float32** | Specifies the threshold value for the JBIG2 encoder pattern matching and substitution between 0 and 1. Any number lower than 1 may lead to lossy compression. | [optional] [default to 0.85]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


