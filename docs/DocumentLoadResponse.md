# DocumentLoadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**Error**](Error.md) |  | [optional] 
**RemainingTokens** | **int64** | Specifies the number of remaining tokens. | [optional] 
**FileId** | Pointer to **string** | Specifies the file identifier of the loaded document. | [optional] [readonly] 
**DocumentFormat** | [**DocumentFormat**](DocumentFormat.md) |  | [optional] 
**PageCount** | **int32** | Specifies the number of pages into the loaded document. | [optional] [readonly] 
**ThumbnailData** | Pointer to **string** | Specifies the data of a thumbnail from the first page of the document, in PNG format. Only applicable if the GetPreview field of the request has been set to true. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


