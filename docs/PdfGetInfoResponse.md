# PdfGetInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**Error**](Error.md) |  | [optional] 
**RemainingTokens** | **int64** | Specifies the number of remaining tokens. | [optional] 
**PageCount** | **int32** | Specifies the number of pages of the PDF. | [optional] [readonly] 
**Version** | Pointer to **string** | Specifies the version of the PDF. | [optional] [readonly] 
**Author** | Pointer to **string** | Specifies the author name specified within the PDF, if any. | [optional] [readonly] 
**Title** | Pointer to **string** | Specifies the document title specified within the PDF, if any. | [optional] [readonly] 
**Subject** | Pointer to **string** | Specifies the document subject specified within the PDF, if any. | [optional] [readonly] 
**Producer** | Pointer to **string** | Specifies the producer name specified within the PDF, if any. | [optional] [readonly] 
**Metadata** | Pointer to **string** | Specifies the metadata contained within the PDF, if any. | [optional] [readonly] 
**Keywords** | Pointer to **string** | Specifies the keywords associated with the PDF, if any. | [optional] [readonly] 
**HasUserPassword** | **bool** | Specifies if the PDF is protected with a user password. | [optional] [readonly] 
**HasOwnerPassword** | **bool** | Specifies if the PDF is protected with a owner password. | [optional] [readonly] 
**Encryption** | [**EncryptionAlgorithm**](EncryptionAlgorithm.md) |  | [optional] 
**CanPrint** | **bool** | Specifies if the user is allowed to print the document, but possibly not at the highest quality level. | [optional] [readonly] 
**CanCopy** | **bool** | Specifies if the user is allowed to copy or extract text and graphics from the document. | [optional] [readonly] 
**CanModify** | **bool** | Specifies if the user is allowed to modify the document. | [optional] [readonly] 
**CanAddNotes** | **bool** | Specifies if the user is allowed to add annotations. | [optional] [readonly] 
**CanFillFields** | **bool** | Specifies if the user is allowed to fill-in form fields. | [optional] [readonly] 
**CanCopyAccess** | **bool** | Specifies if the user is allowed for copying or extracting for use with accessibility features. | [optional] [readonly] 
**CanAssemble** | **bool** | Specifies if the user is allowed to assemble the document. | [optional] [readonly] 
**CanPrintFull** | **bool** | Specifies if the user is allowed to print the document with high resolution. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


