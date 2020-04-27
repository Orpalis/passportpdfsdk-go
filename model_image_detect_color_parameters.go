/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ImageDetectColorParameters Represents the parameters for an image color detection action.
type ImageDetectColorParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the number of the page, or the range of pages to perform color detection on.
	PageRange string `json:"PageRange"`
	// Specifies whether to automatically convert the image in its best suited/optimized bits-per-pixel encoding.
	AutoConvert bool `json:"AutoConvert,omitempty"`
	// Specifies whether the characters shall be repaired during bitonal conversion, if any, or not.
	AutoRepairCharacters bool `json:"AutoRepairCharacters,omitempty"`
}
