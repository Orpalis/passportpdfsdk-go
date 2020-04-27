/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfRemovePageFormFieldsParameters Represents the parameters for a remove page form fields action.
type PdfRemovePageFormFieldsParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the page or the page range whose form fields shall be removed.
	PageRange string `json:"PageRange"`
}
