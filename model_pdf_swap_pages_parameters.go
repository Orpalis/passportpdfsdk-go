/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfSwapPagesParameters Represents the parameters for a swap pages action.
type PdfSwapPagesParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the first page.
	Page1 int32 `json:"Page1"`
	// Specifies the number of the second page.
	Page2 int32 `json:"Page2"`
}
