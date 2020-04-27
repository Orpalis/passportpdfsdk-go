/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PageNumberFormat Defines the different formats available for the page number insertion.
type PageNumberFormat string

// List of PageNumberFormat
const (
	PAGE_NUMBER PageNumberFormat = "PageNumber"
	PAGE_PAGE_NUMBER PageNumberFormat = "PagePageNumber"
	PAGE_PAGE_NUMBER_OF_PAGE_COUNT PageNumberFormat = "PagePageNumberOfPageCount"
	CUSTOM_FORMATTING PageNumberFormat = "CustomFormatting"
)
