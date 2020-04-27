# ImageMicrParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the page or the range of page to be processed. | 
**Font** | [**MicrFont**](MicrFont.md) |  | 
**Context** | [**MicrContext**](MicrContext.md) |  | 
**WhiteList** | Pointer to **string** | Specifies the characters to be ignored by the recognition process. | [optional] [default to ]
**RoiLeft** | **int32** | Specifies the left coordinate, in pixel, of the region to process. | [optional] [default to 0]
**RoiTop** | **int32** | Specifies the top coordinate, in pixel, of the region to process. | [optional] [default to 0]
**RoiWidth** | **int32** | Specifies the width, in pixel, of the region to process. 0 causes the entire image to be processed. | [optional] [default to 0]
**RoiHeight** | **int32** | Specifies the height, in pixel, of the region to process. 0 causes the entire image to be processed. | [optional] [default to 0]
**ExpectedSymbolsCount** | **int32** | Specifies the number of symbols expected to be detected, use 0 if unknown. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


