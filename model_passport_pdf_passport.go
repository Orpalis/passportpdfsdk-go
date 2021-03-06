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
	"time"
)
// PassportPdfPassport struct for PassportPdfPassport
type PassportPdfPassport struct {
	PassportId *string `json:"PassportId,omitempty"`
	PlanId int32 `json:"PlanId,omitempty"`
	OptionalLinkedSubscriptionId *string `json:"OptionalLinkedSubscriptionId,omitempty"`
	SubscriptionDate time.Time `json:"SubscriptionDate,omitempty"`
	ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
	CurrentTokensUsed int64 `json:"CurrentTokensUsed,omitempty"`
	ExtraTokens int64 `json:"ExtraTokens,omitempty"`
	ExtraTokensEndDate time.Time `json:"ExtraTokensEndDate,omitempty"`
	IsYearlyPaid bool `json:"IsYearlyPaid,omitempty"`
	IsActive bool `json:"IsActive,omitempty"`
	TokensPercentUsageAlert int32 `json:"TokensPercentUsageAlert,omitempty"`
	TokensPercentUsageAlertSent bool `json:"TokensPercentUsageAlertSent,omitempty"`
	IsManaged bool `json:"IsManaged,omitempty"`
	NextMonthlyTerm time.Time `json:"NextMonthlyTerm,omitempty"`
	RemainingTokens int64 `json:"RemainingTokens,omitempty"`
}
