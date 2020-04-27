# PdfDigiSignParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**CertificateData** | **string** | Specifies the data of the digital PKCS#12 certificate file. | 
**CertificatePassword** | **string** | Specifies the certificate password. | 
**SignatureMode** | [**SignatureMode**](SignatureMode.md) |  | [optional] 
**SignatureCertificationLevel** | [**SignatureCertificationLevel**](SignatureCertificationLevel.md) |  | [optional] 
**SignatureHashAlgorithm** | [**SignatureHash**](SignatureHash.md) |  | [optional] 
**SignerName** | Pointer to **string** | Specifies the name of the signer. | [optional] [default to PassportPDF]
**Reason** | Pointer to **string** | Specifies the reason of the signature. | [optional] [default to ]
**Location** | Pointer to **string** | Specifies the location where the signature is applied. | [optional] [default to ]
**ContactInfo** | Pointer to **string** | Specifies contact information about the signer. | [optional] [default to ]
**TimeStampURL** | Pointer to **string** | Specifies the URL of the server responsible of providing a time stamp. | [optional] [default to ]
**TimeStampUserName** | Pointer to **string** | Specifies the optional user name associated with the time stamp server. | [optional] [default to ]
**TimeStampPassword** | Pointer to **string** | Specifies the optional password associated with the time stamp server. | [optional] [default to ]
**Linearize** | **bool** | Specifies whether the signed PDF shall be linearized. | [optional] [default to false]
**DrawSignature** | **bool** | Specifies whether the signature shall be drawn on the document. | [optional] [default to false]
**PageNumber** | **int32** | Specifies the number of the page on which the signature shall be drawn. | [optional] [default to 1]
**ShowValidationMark** | **bool** | Specifies whether a validation mark shall be drawn with the signature. | [optional] [default to false]
**ImageData** | Pointer to **string** | Specifies the data of the image to be drawn at the signature location. | [optional] 
**SignatureLayout** | [**DrawableContentLayoutParameters**](DrawableContentLayoutParameters.md) |  | [optional] 
**SignatureText** | [**PdfAlignedTextParameters**](PdfAlignedTextParameters.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


