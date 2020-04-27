/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ImageSaveAsPngParameters Represents the parameters for a save as PNG action.
type ImageSaveAsPngParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages to be saved as PNG.
	PageRange *string `json:"PageRange,omitempty"`
	// Specifies the level of compression to be used for the PNG output, between 0 (no compression - faster encoding) to 9(max compression - slower encoding).
	Compression int32 `json:"Compression,omitempty"`
	// Specifies if the produced PNG image must be interlaced.
	Interlaced bool `json:"Interlaced,omitempty"`
}
