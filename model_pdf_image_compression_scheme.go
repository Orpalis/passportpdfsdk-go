/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfImageCompressionScheme Specifies the standard scheme to be used to compress image data in PDF documents.
type PdfImageCompressionScheme string

// List of PdfImageCompressionScheme
const (
	NONE PdfImageCompressionScheme = "None"
	FLATE PdfImageCompressionScheme = "Flate"
	CCIT4 PdfImageCompressionScheme = "CCIT4"
	JPEG PdfImageCompressionScheme = "JPEG"
	JBIG2 PdfImageCompressionScheme = "JBIG2"
	JPEG2000 PdfImageCompressionScheme = "JPEG2000"
)
