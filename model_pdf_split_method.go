/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfSplitMethod Defines the different available methods for splitting a PDF.
type PdfSplitMethod string

// List of PdfSplitMethod
const (
	SPLIT_BY_NUMBER_OF_PAGES PdfSplitMethod = "SplitByNumberOfPages"
	SPLIT_BY_FILE_SIZE PdfSplitMethod = "SplitByFileSize"
	SPLIT_BY_TOP_LEVEL_BOOKMARKS PdfSplitMethod = "SplitByTopLevelBookmarks"
)
