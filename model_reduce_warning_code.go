/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ReduceWarningCode Specifies the different warnings which can occur during the process of a reduce action.
type ReduceWarningCode string

// List of ReduceWarningCode
const (
	OK ReduceWarningCode = "OK"
	IMAGE_EXTRACTION_FAILURE ReduceWarningCode = "ImageExtractionFailure"
	COLOR_DETECTION_FAILURE ReduceWarningCode = "ColorDetectionFailure"
	IMAGE_RESIZE_FAILURE ReduceWarningCode = "ImageResizeFailure"
	IMAGE_CROP_FAILURE ReduceWarningCode = "ImageCropFailure"
	IMAGE_RESOLUTION_OBTENTION_FAILURE ReduceWarningCode = "ImageResolutionObtentionFailure"
	IMAGE_REPLACEMENT_FAILURE ReduceWarningCode = "ImageReplacementFailure"
	MRC_IMAGE_REPLACEMENT_FAILURE ReduceWarningCode = "MRCImageReplacementFailure"
	PAGE_SELECTION_FAILURE ReduceWarningCode = "PageSelectionFailure"
	IMAGE_OBTENTION_FAILURE ReduceWarningCode = "ImageObtentionFailure"
	FILE_SIZE_REDUCTION_FAILURE ReduceWarningCode = "FileSizeReductionFailure"
	BLANK_PAGE_REMOVAL_FAILURE ReduceWarningCode = "BlankPageRemovalFailure"
)
