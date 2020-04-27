# LoadImageParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Specifies the name of the document. | [optional] 
**ContentEncoding** | [**ContentEncoding**](ContentEncoding.md) |  | [optional] 
**GetPreview** | **bool** | Specifies whether the response must contain a thumbnail of the first page of the document. | [optional] [default to false]
**ThumbnailWidth** | **int32** | Specifies, in pixels, the width of the thumbnail to be retrieved. Only applicable if GetPreview has been set to true. | [optional] [default to 140]
**ThumbnailHeight** | **int32** | Specifies, in pixels, the height of the thumbnail to be retrieved.  Only applicable if GetPreview has been set to true. | [optional] [default to 220]
**ThumbnailBackgroundColor** | Pointer to **string** | Specifies the background color of the thumbnail, using the color name (ie: \&quot;red\&quot;) or its RGBa code (ie: \&quot;rgba(255,0,0,1)\&quot;).   Only applicable if GetPreview has been set to true. | [optional] [default to rgba(0,0,0,0)]
**ThumbnailFitToPageSize** | **bool** | Specifies if the size of the produced thumbnail is automatically adjusted to don&#39;t have any margin.  Only applicable if GetPreview has been set to true. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


