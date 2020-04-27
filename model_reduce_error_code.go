/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ReduceErrorCode Specifies the different errors which can cause the failure of a reduce action.
type ReduceErrorCode string

// List of ReduceErrorCode
const (
	OK ReduceErrorCode = "OK"
	GET_PAGE_IMAGES_COUNT ReduceErrorCode = "GetPageImagesCount"
	MRC_POST_OPERATIONS_FAILURE ReduceErrorCode = "MRCPostOperationsFailure"
	PAGE_CONVERSION_FAILURE ReduceErrorCode = "PageConversionFailure"
	DOCUMENT_ENCRYPTED ReduceErrorCode = "DocumentEncrypted"
	UNEXPECTED_ERROR ReduceErrorCode = "UnexpectedError"
)
