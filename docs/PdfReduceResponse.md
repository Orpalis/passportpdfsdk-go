# PdfReduceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**Error**](Error.md) |  | [optional] 
**RemainingTokens** | **int64** | Specifies the number of remaining tokens. | [optional] 
**ErrorInfo** | [**ReduceErrorInfo**](ReduceErrorInfo.md) |  | [optional] 
**WarningsInfo** | Pointer to [**[]ReduceWarningInfo**](ReduceWarningInfo.md) | Specifies the different warnings that occured during the process of the reduce action. | [optional] [readonly] 
**ContentRemoved** | **bool** | Specifies whether content has been removed from the PDF during the process of the reduce action. | [optional] [readonly] 
**VersionChanged** | **bool** | Specifies whether the version of the PDF has changed during the process of the reduce action. | [optional] [readonly] 
**NewFileSize** | **int64** | Specify the size of the new created document. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


