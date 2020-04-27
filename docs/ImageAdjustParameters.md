# ImageAdjustParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the number of the page, or the range of pages to adjust. | 
**RoiLeft** | **int32** | Specifies the left coordinate, in pixel, of the region to process. | [optional] [default to 0]
**RoiTop** | **int32** | Specifies the top coordinate, in pixel, of the region to process. | [optional] [default to 0]
**RoiWidth** | **int32** | Specifies the width, in pixel, of the region to process. 0 causes the entire image to be processed. | [optional] [default to 0]
**RoiHeight** | **int32** | Specifies the height, in pixel, of the region to process. 0 causes the entire image to be processed. | [optional] [default to 0]
**GammaCorrection** | **int32** | Specifies the gamma correction parameter. | [optional] [default to 0]
**Brightness** | **int32** | Specifies the brightness parameter. | [optional] [default to 0]
**Contrast** | **int32** | Specifies the contrast parameter. | [optional] [default to 0]
**Saturation** | **int32** | Specifies the saturation parameter. | [optional] [default to 0]
**AutoContrastEnhancement** | **bool** | Specifies whether the contrast shall be automatically enhanced. | [optional] [default to false]
**ContrastHistogramStretch** | **bool** | Specifies whether a contrast histogram stretch filter shall be performed. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


