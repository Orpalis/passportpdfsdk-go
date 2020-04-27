# PdfExtractTextParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the number of the page, or the range of pages to extract text from. | 
**TextExtractionMode** | [**PdfExtractTextMode**](PdfExtractTextMode.md) |  | [optional] 
**TextExtractionAreaLeft** | **float32** | Specifies the left coordinate of the text extraction area for the PageAreaFormated extract mode. | [optional] [default to 0]
**TextExtractionAreaTop** | **float32** | Specifies the top coordinate of the text extraction area for the PageAreaFormated extract mode. | [optional] [default to 0]
**TextExtractionAreaWidth** | **float32** | Specifies the width of the text extraction area for the PageAreaFormated extract mode. | [optional] [default to 0]
**TextExtractionAreaHeight** | **float32** | Specifies the height of the text extraction area for the PageAreaFormated extract mode. | [optional] [default to 0]
**IncludeHeader** | **bool** | Specifies whether the CSV output should include header information. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


