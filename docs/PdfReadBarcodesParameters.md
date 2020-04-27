# PdfReadBarcodesParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The identifier of the previously uploaded file to be processed. | 
**PageRange** | **string** | Specifies the number of the page, or the range of pages to read the barcodes from. | 
**ScanMode** | [**ScanMode**](ScanMode.md) |  | [optional] 
**ScanBarcode1D** | **bool** | Specifies whether the barcodes of type 1D shall be read. | [optional] [default to true]
**ScanBarcodeQR** | **bool** | Specifies whether the barcodes of type QR shall be read. | [optional] [default to true]
**ScanBarcodeMicroQR** | **bool** | Specifies whether the barcodes of type Micro QR shall be read. | [optional] [default to true]
**ScanBarcodePDF417** | **bool** | Specifies whether the barcodes of type PDF 147 shall be read. | [optional] [default to true]
**ScanBarcodeDataMatrix** | **bool** | Specifies whether the barcodes of type Data Matrix shall be read. | [optional] [default to true]
**ScanBarcodeAztec** | **bool** | Specifies whether the barcodes of type Aztec shall be read. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


