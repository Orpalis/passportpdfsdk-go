/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfSetInfoParameters Represents the parameters for a set info action.
type PdfSetInfoParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the author name specified within the PDF, if any.
	Author *string `json:"Author,omitempty"`
	// Specifies the document title specified within the PDF, if any.
	Title *string `json:"Title,omitempty"`
	// Specifies the document subject specified within the PDF, if any.
	Subject *string `json:"Subject,omitempty"`
	// Specifies the producer name specified within the PDF, if any.
	Producer *string `json:"Producer,omitempty"`
	// Specifies the metadata contained within the PDF, if any.
	Metadata *string `json:"Metadata,omitempty"`
	// Specifies the keywords associated with the PDF, if any.
	Keywords *string `json:"Keywords,omitempty"`
	// Specifies whether the empty info values shall be cleared.
	ClearEmptyValues bool `json:"ClearEmptyValues,omitempty"`
}
