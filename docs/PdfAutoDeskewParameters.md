# PdfAutoDeskewParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the page or the page range to be auto-deskewed. | 
**MaxAngleOfResearch** | **float32** | Specifies the maximum skew angle to be detected. A value of below 15 is suggested. | [optional] [default to 15]
**Optimistic** | **bool** | Specifies whether the skew detection engine must be optimistic when detecting angles.  If you know the image is skewed, you should set this to true. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


