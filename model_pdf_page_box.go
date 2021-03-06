/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfPageBox Defines the boundary boxes (page boxes) that relate to the size of the pages in the PDF document.
type PdfPageBox string

// List of PdfPageBox
const (
	MEDIA_BOX PdfPageBox = "MediaBox"
	CROP_BOX PdfPageBox = "CropBox"
	BLEED_BOX PdfPageBox = "BleedBox"
	TRIM_BOX PdfPageBox = "TrimBox"
	ART_BOX PdfPageBox = "ArtBox"
)
