/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ApiV2010AccountCallCallFeedbackSummary struct {

	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`

	// The total number of calls
	CallCount *int32 `json:"call_count,omitempty"`

	// The total number of calls with a feedback entry
	CallFeedbackCount *int32 `json:"call_feedback_count,omitempty"`

	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`

	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`

	// The latest feedback entry date in the summary
	EndDate *string `json:"end_date,omitempty"`

	// Whether the feedback summary includes subaccounts
	IncludeSubaccounts *bool `json:"include_subaccounts,omitempty"`

	// Issues experienced during the call
	Issues *[]interface{} `json:"issues,omitempty"`

	// The average QualityScore of the feedback entries
	QualityScoreAverage *float32 `json:"quality_score_average,omitempty"`

	// The median QualityScore of the feedback entries
	QualityScoreMedian *float32 `json:"quality_score_median,omitempty"`

	// The standard deviation of the quality scores
	QualityScoreStandardDeviation *float32 `json:"quality_score_standard_deviation,omitempty"`

	// A string that uniquely identifies this feedback entry
	Sid *string `json:"sid,omitempty"`

	// The earliest feedback entry date in the summary
	StartDate *string `json:"start_date,omitempty"`

	// The status of the feedback summary
	Status *string `json:"status,omitempty"`
}

// AssertApiV2010AccountCallCallFeedbackSummaryRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountCallCallFeedbackSummaryRequired(obj ApiV2010AccountCallCallFeedbackSummary) error {
	return nil
}

// AssertRecurseApiV2010AccountCallCallFeedbackSummaryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountCallCallFeedbackSummary (e.g. [][]ApiV2010AccountCallCallFeedbackSummary), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountCallCallFeedbackSummaryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountCallCallFeedbackSummary, ok := obj.(ApiV2010AccountCallCallFeedbackSummary)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountCallCallFeedbackSummaryRequired(aApiV2010AccountCallCallFeedbackSummary)
	})
}
