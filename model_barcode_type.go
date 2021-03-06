/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BarcodeType Defines the different type of barcodes which can be read.
type BarcodeType string

// List of BarcodeType
const (
	LINEAR BarcodeType = "Linear"
	DATA_MATRIX BarcodeType = "DataMatrix"
	QR BarcodeType = "QR"
	MICRO_QR BarcodeType = "MicroQR"
	PDF417 BarcodeType = "PDF417"
	AZTEC BarcodeType = "Aztec"
)
