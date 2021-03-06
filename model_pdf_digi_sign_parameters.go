/*
 * PassportPDF API
 *
 * Another brick in the cloud
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PdfDigiSignParameters Represents the parameters for a digital signature action.
type PdfDigiSignParameters struct {
	// The identifier of the previously uploaded file to be processed.
	FileId string `json:"FileId"`
	// Specifies the data of the digital PKCS#12 certificate file.
	CertificateData string `json:"CertificateData"`
	// Specifies the certificate password.
	CertificatePassword string `json:"CertificatePassword"`
	SignatureMode SignatureMode `json:"SignatureMode,omitempty"`
	SignatureCertificationLevel SignatureCertificationLevel `json:"SignatureCertificationLevel,omitempty"`
	SignatureHashAlgorithm SignatureHash `json:"SignatureHashAlgorithm,omitempty"`
	// Specifies the name of the signer.
	SignerName *string `json:"SignerName,omitempty"`
	// Specifies the reason of the signature.
	Reason *string `json:"Reason,omitempty"`
	// Specifies the location where the signature is applied.
	Location *string `json:"Location,omitempty"`
	// Specifies contact information about the signer.
	ContactInfo *string `json:"ContactInfo,omitempty"`
	// Specifies the URL of the server responsible of providing a time stamp.
	TimeStampURL *string `json:"TimeStampURL,omitempty"`
	// Specifies the optional user name associated with the time stamp server.
	TimeStampUserName *string `json:"TimeStampUserName,omitempty"`
	// Specifies the optional password associated with the time stamp server.
	TimeStampPassword *string `json:"TimeStampPassword,omitempty"`
	// Specifies whether the signed PDF shall be linearized.
	Linearize bool `json:"Linearize,omitempty"`
	// Specifies whether the signature shall be drawn on the document.
	DrawSignature bool `json:"DrawSignature,omitempty"`
	// Specifies the number of the page on which the signature shall be drawn.
	PageNumber int32 `json:"PageNumber,omitempty"`
	// Specifies whether a validation mark shall be drawn with the signature.
	ShowValidationMark bool `json:"ShowValidationMark,omitempty"`
	// Specifies the data of the image to be drawn at the signature location.
	ImageData *string `json:"ImageData,omitempty"`
	SignatureLayout DrawableContentLayoutParameters `json:"SignatureLayout,omitempty"`
	SignatureText PdfAlignedTextParameters `json:"SignatureText,omitempty"`
}
