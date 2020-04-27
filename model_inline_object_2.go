/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"os"
)
// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// The data of the document.
	FileData *os.File `json:"fileData"`
	LoadDocumentParameters PdfLoadDocumentParameters `json:"loadDocumentParameters,omitempty"`
}
