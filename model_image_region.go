/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ImageRegion Represents a region of a given page of an image.
type ImageRegion struct {
	// Specifies the number of the page.
	PageNumber int32 `json:"PageNumber,omitempty"`
	// Specifies, in pixels, the left coordinate of the region.
	Left int32 `json:"Left,omitempty"`
	// Specifies, in pixels, the top coordinate of the region.
	Top int32 `json:"Top,omitempty"`
	// Specifies, in pixels, the width of the region.
	Width int32 `json:"Width,omitempty"`
	// Specifies, in pixels, the height of the region.
	Height int32 `json:"Height,omitempty"`
}
