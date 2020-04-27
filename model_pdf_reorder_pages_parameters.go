/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfReorderPagesParameters Represents the parameters for a reorder pages action.
type PdfReorderPagesParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the new pages order. Must contains page numbers separated by ';' or ','.  Example for a 6-pages document: \"1;2;4;3;6;5\".
	PageOrder string `json:"PageOrder"`
}
