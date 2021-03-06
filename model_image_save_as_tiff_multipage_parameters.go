/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ImageSaveAsTiffMultipageParameters Represents the parameters for a save as TIFF multipage action.
type ImageSaveAsTiffMultipageParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages to be saved as TIFF.
	PageRange *string `json:"PageRange,omitempty"`
	Compression TiffSaveCompression `json:"Compression,omitempty"`
	// Specifies the level of quality to be used if JPEG compression is used, from 1 (poorest) to 100 (greatest).
	JpegQuality int32 `json:"JpegQuality,omitempty"`
}
