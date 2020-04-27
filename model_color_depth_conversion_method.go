/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ColorDepthConversionMethod Specifies the modes when changing the color depth of an image.
type ColorDepthConversionMethod string

// List of ColorDepthConversionMethod
const (
	DEPTH1_BPP ColorDepthConversionMethod = "Depth1Bpp"
	DEPTH1_BPP_OTSU ColorDepthConversionMethod = "Depth1BppOtsu"
	DEPTH1_BPP_BRADLEY ColorDepthConversionMethod = "Depth1BppBradley"
	DEPTH1_BPP_SAUVOLA ColorDepthConversionMethod = "Depth1BppSauvola"
	DEPTH1_BPP_FAST ColorDepthConversionMethod = "Depth1BppFast"
	DEPTH4_BPP16 ColorDepthConversionMethod = "Depth4Bpp16"
	DEPTH4_BPP_Q ColorDepthConversionMethod = "Depth4BppQ"
	DEPTH8_BPP_GRAY_SCALE ColorDepthConversionMethod = "Depth8BppGrayScale"
	DEPTH8_BPP_GRAY_SCALE_ADV ColorDepthConversionMethod = "Depth8BppGrayScaleAdv"
	DEPTH8_BPP216 ColorDepthConversionMethod = "Depth8Bpp216"
	DEPTH8_BPP_Q ColorDepthConversionMethod = "Depth8BppQ"
	DEPTH16_BPP_RGB555 ColorDepthConversionMethod = "Depth16BppRGB555"
	DEPTH16_BPP_RGB565 ColorDepthConversionMethod = "Depth16BppRGB565"
	DEPTH24_BPP_RGB ColorDepthConversionMethod = "Depth24BppRGB"
	DEPTH32_BPP_ARGB ColorDepthConversionMethod = "Depth32BppARGB"
	DEPTH32_BPP_RGB ColorDepthConversionMethod = "Depth32BppRGB"
	DEPTH32_BPP_PARGB ColorDepthConversionMethod = "Depth32BppPARGB"
	DEPTH48_BPP_RGB ColorDepthConversionMethod = "Depth48BppRGB"
	DEPTH64_BPP_ARGB ColorDepthConversionMethod = "Depth64BppARGB"
	DEPTH64_BPP_PARGB ColorDepthConversionMethod = "Depth64BppPARGB"
)
