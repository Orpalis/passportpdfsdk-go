# PdfReduceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**OutputVersion** | [**PdfVersion**](PdfVersion.md) |  | [optional] 
**ImageQuality** | [**ImageQuality**](ImageQuality.md) |  | [optional] 
**RecompressImages** | **bool** | Specifies whether the images from the PDF shall be recompressed. | [optional] [default to true]
**EnableColorDetection** | **bool** | Specifies whether color detection must be performed on the images from the PDF. | [optional] [default to true]
**PackDocument** | **bool** | Specifies whether the PDF shall be packed when saved in order to reduce its size. | [optional] [default to true]
**PackFonts** | **bool** | Specifies whether the PDF fonts must be packed in order to reduce their size. | [optional] [default to true]
**DownscaleImages** | **bool** | Specifies whether the images from the PDF shall be downscaled. | [optional] [default to true]
**DownscaleResolution** | **int32** | Specifies the resolution to be used to downscale images. | [optional] [default to 150]
**FastWebView** | **bool** | Specifies whether the PDF shall be optimized for online distribution. | [optional] [default to false]
**RemoveFormFields** | **bool** | Specifies whether the form fields shall be removed from the PDF. | [optional] [default to false]
**RemoveAnnotations** | **bool** | Specifies whether the annotations shall be removed from the PDF. | [optional] [default to false]
**RemoveBookmarks** | **bool** | Specifies whether the bookmarks shall be removed from the PDF. | [optional] [default to false]
**RemoveHyperlinks** | **bool** | Specifies whether the hyperlinks shall be removed from the PDF. | [optional] [default to false]
**RemoveEmbeddedFiles** | **bool** | Specifies whether the embedded files shall be removed from the PDF. | [optional] [default to false]
**RemoveBlankPages** | **bool** | Specifies whether the blank pages shall be removed. | [optional] [default to false]
**RemoveJavaScript** | **bool** | Specifies whether the JavaScript shall be removed. | [optional] [default to false]
**EnableJPEG2000** | **bool** | Specifies whether the JPEG2000 compression scheme shall be used to compress the images of the PDF. | [optional] [default to true]
**EnableJBIG2** | **bool** | Specifies whether the JBIG2 compression scheme shall be used to compress the bitonal images of the PDF. | [optional] [default to true]
**EnableCharRepair** | **bool** | Specifies whether characters repairing shall be performed during bitonal conversion. | [optional] [default to false]
**EnableMRC** | **bool** | Specifies whether MRC shall be used for compressing the PDF contents. | [optional] [default to false]
**PreserveSmoothing** | **bool** | Specifies if the MRC engine shall try to preserve smoothing between different layers. | [optional] [default to false]
**DownscaleResolutionMRC** | **int32** | Specifies the resolution for downscaling the background layer by the MRC engine, if any. | [optional] [default to 100]
**RemoveMetadata** | **bool** | Specifies whether the metadata shall be removed. | [optional] [default to false]
**RemovePageThumbnails** | **bool** | Specifies whether the page thumbnails shall be removed. | [optional] [default to false]
**RemovePagePieceInfo** | **bool** | Specifies whether the page PieceInfo dictionary used to hold private application data shall be removed. | [optional] [default to false]
**JBIG2PMSThreshold** | **float32** | Specifies the threshold value for the JBIG2 encoder pattern matching and substitution between 0 and 1. Any number lower than 1 may lead to lossy compression. | [optional] [default to 0.85]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


