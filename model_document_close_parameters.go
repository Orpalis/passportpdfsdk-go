/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DocumentCloseParameters Represents the parameters for a close document action.
type DocumentCloseParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
}
