# PdfSetInitialViewParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageMode** | [**PdfInitViewPageMode**](PdfInitViewPageMode.md) |  | [optional] 
**LayoutMode** | [**PdfInitViewLayoutMode**](PdfInitViewLayoutMode.md) |  | [optional] 
**NonFullScreenPageMode** | [**PdfInitViewNonFullScreenPageMode**](PdfInitViewNonFullScreenPageMode.md) |  | [optional] 
**OpenPage** | **int32** | Specifies which page should be displayed when the document is opened. | [optional] [default to 1]
**OpenZoom** | **float32** | Specifies the default zoom factor to be used when the document is opened. Value of 1 to represent the 100% zoom, 2 means 200%, 0,5 means 50%, etc. | [optional] [default to 1]
**HideToolbar** | **bool** | A flag specifying whether to hide the viewer application’s tool bars when the document is active. Default value: false. | [optional] [default to false]
**HideMenubar** | **bool** | (Optional) A flag specifying whether to hide the viewer application’s menu bar when the document is active. Default value: false. | [optional] [default to false]
**HideWindowUI** | **bool** | (Optional) A flag specifying whether to hide user interface elements in the document’s window (such as scroll bars and navigation controls),  leaving only the document’s contents displayed. Default value: false. | [optional] [default to false]
**FitWindow** | **bool** | (Optional) A flag specifying whether to resize the document’s window to fit the size of the first displayed page. Default value: false. | [optional] [default to false]
**CenterWindow** | **bool** | (Optional) A flag specifying whether to position the document’s window in the center of the screen. Default value: false. | [optional] [default to false]
**DisplayDocTitle** | **bool** | (Optional; PDF 1.4) A flag specifying whether the window’s title bar should display the document title taken from the Title entry of the document information dictionary.  If false, the title bar should instead display the name of the PDF file containing the document. Default value: false. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


