/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ImageDetectPageOrientationParameters Represents the parameters for a detect page orientation action.
type ImageDetectPageOrientationParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages to be processed.
	PageRange string `json:"PageRange"`
	// Specifies the language(s) to be used for the OCR.
	Language *string `json:"Language,omitempty"`
	// Specifies whether a rotation shall be automatically applied in order to correct the page orientation when needed.
	AutomaticallyApplyRotation bool `json:"AutomaticallyApplyRotation,omitempty"`
}
