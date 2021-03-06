/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfInsertTextParameters Represents the parameters for an insert text action.
type PdfInsertTextParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages on which the text shall be inserted.
	PageRange string `json:"PageRange"`
	TextParameters PdfAlignedTextParameters `json:"TextParameters"`
	TextBoundingBoxLayout DrawableContentLayoutParameters `json:"TextBoundingBoxLayout,omitempty"`
}
