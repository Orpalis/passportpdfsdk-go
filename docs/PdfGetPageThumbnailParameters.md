# PdfGetPageThumbnailParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the page or the range of pages to get a thumbnail from. | 
**ThumbnailWidth** | **int32** | Specifies the width, in points, of the thumbnail(s). | [optional] [default to 140]
**ThumbnailHeight** | **int32** | Specifies the height, in points, of the thumbnail(s). | [optional] [default to 220]
**BackgroundColor** | Pointer to **string** | Specifies the background color of the thumbnail(s), using the color name (ie: \&quot;red\&quot;) or its RGBa code (ie: \&quot;rgba(255,0,0,1)\&quot;). | [optional] [default to rgba(0,0,0,0)]
**ThumbnailFitToPageSize** | **bool** | Specifies if the size of the produced thumbnail is automatically adjusted to don&#39;t have any margin. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


