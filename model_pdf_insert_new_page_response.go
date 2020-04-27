/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfInsertNewPageResponse Represents the response to an insert new page action request.
type PdfInsertNewPageResponse struct {
	Error Error `json:"Error,omitempty"`
	// Specifies the number of remaining tokens.
	RemainingTokens int64 `json:"RemainingTokens,omitempty"`
}
