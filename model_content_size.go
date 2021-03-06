/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContentSize Defines the different possible sizes for content added on a page.
type ContentSize string

// List of ContentSize
const (
	SMALL ContentSize = "Small"
	MEDIUM ContentSize = "Medium"
	LARGE ContentSize = "Large"
	ABSOLUTE ContentSize = "Absolute"
	AUTO_FIT ContentSize = "AutoFit"
)
