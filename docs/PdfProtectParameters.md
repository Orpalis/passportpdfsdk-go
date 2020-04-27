# PdfProtectParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**OwnerPassword** | Pointer to **string** | Specifies the owner password to be set. | [optional] [default to ]
**UserPassword** | Pointer to **string** | Specifies the user password to be set. | [optional] [default to ]
**Encryption** | [**EncryptionAlgorithm**](EncryptionAlgorithm.md) |  | [optional] 
**CanPrint** | **bool** | Allows the user to print the document, but possibly not at the highest quality level. Ignored if no encryption algorithm is set. | [optional] [default to true]
**CanCopy** | **bool** | Allows the user to copy or extract text and graphics from the document. Ignored if the no encryption scheme is set. | [optional] [default to true]
**CanModify** | **bool** | Allows the user to modify the document. Ignored if the no encryption scheme is set. | [optional] [default to true]
**CanAddNotes** | **bool** | Allows the user to add annotations. Ignored if the no encryption scheme is set. | [optional] [default to true]
**CanFillFields** | **bool** | Allows the user to fill-in form fields. Only works with 128-bit encryption. | [optional] [default to true]
**CanCopyAccess** | **bool** | Enables copying or extracting for use with accessibility features. Only works with 128-bit encryption. | [optional] [default to true]
**CanAssemble** | **bool** | Allows the user to assemble the document. Only works with 128-bit encryption. | [optional] [default to true]
**CanPrintFull** | **bool** | Allows high resolution printing of the document. Only works with 128-bit encryption. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


