/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfReduceResponse Represents the response to a reduce action request.
type PdfReduceResponse struct {
	Error Error `json:"Error,omitempty"`
	// Specifies the number of remaining tokens.
	RemainingTokens int64 `json:"RemainingTokens,omitempty"`
	ErrorInfo ReduceErrorInfo `json:"ErrorInfo,omitempty"`
	// Specifies the different warnings that occured during the process of the reduce action.
	WarningsInfo *[]ReduceWarningInfo `json:"WarningsInfo,omitempty"`
	// Specifies whether content has been removed from the PDF during the process of the reduce action.
	ContentRemoved bool `json:"ContentRemoved,omitempty"`
	// Specifies whether the version of the PDF has changed during the process of the reduce action.
	VersionChanged bool `json:"VersionChanged,omitempty"`
	// Specify the size of the new created document.
	NewFileSize int64 `json:"NewFileSize,omitempty"`
}
