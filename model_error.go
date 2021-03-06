/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Error Contains all the information related to an error which has occured.
type Error struct {
	ResultCode PassportPdfStatus `json:"ResultCode,omitempty"`
	// Specifies a result code related to an error which occured in an external component.
	ExtResultStatus *string `json:"ExtResultStatus,omitempty"`
	// Specifies a message which further describes the error.
	ExtResultMessage *string `json:"ExtResultMessage,omitempty"`
	// Specifies a unique identifier, allowing to easily assess the error.
	InternalErrorId *string `json:"InternalErrorId,omitempty"`
}
