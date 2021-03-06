/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TiffSaveCompression Specifies the TIFF compression when saving images in TIFF format.
type TiffSaveCompression string

// List of TiffSaveCompression
const (
	DEFLATE TiffSaveCompression = "Deflate"
	CCITT3 TiffSaveCompression = "CCITT3"
	CCITT4 TiffSaveCompression = "CCITT4"
	LZW TiffSaveCompression = "LZW"
	JPEG TiffSaveCompression = "JPEG"
	RLE TiffSaveCompression = "RLE"
	AUTO TiffSaveCompression = "Auto"
)
