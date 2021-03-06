/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfExtractPageParameters Represents the parameters for an extract page action.
type PdfExtractPageParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages to be extracted.
	PageRange string `json:"PageRange"`
	// Specifies whether each extracted page shall form a separate document.
	ExtractAsSeparate bool `json:"ExtractAsSeparate,omitempty"`
	// Specifies whether the file(s) created as a result of the action shall be available for immediate download.
	ImmediateDownload bool `json:"ImmediateDownload,omitempty"`
}
