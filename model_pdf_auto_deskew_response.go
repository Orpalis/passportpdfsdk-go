/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfAutoDeskewResponse Represents the response to an auto deskew action request.
type PdfAutoDeskewResponse struct {
	Error Error `json:"Error,omitempty"`
	// Specifies the number of remaining tokens.
	RemainingTokens int64 `json:"RemainingTokens,omitempty"`
}
