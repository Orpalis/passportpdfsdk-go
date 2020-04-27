# ImageConvertColorDepthParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the number of the page, or the range of pages to process. | 
**ConvertColorDepthMode** | [**ColorDepthConversionMethod**](ColorDepthConversionMethod.md) |  | 
**Depth1BppOtsuThreshold** | **int32** | Specifies the threshold for a conversion using the Depth1BppOtsu method. | [optional] [default to 50]
**Depth1BppBradleyThreshold** | **int32** | Specifies the threshold for a conversion using the Depth1BppBradley method. | [optional] [default to 38]
**Depth1BppSauvolaFactor** | **float32** | Specifies the factor for a conversion using the Depth1BppSauvola method. | [optional] [default to 0.35]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


