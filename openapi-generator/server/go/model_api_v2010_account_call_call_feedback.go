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

type ApiV2010AccountCallCallFeedback struct {

	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`

	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`

	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`

	// Issues experienced during the call
	Issues *[]string `json:"issues,omitempty"`

	// 1 to 5 quality score
	QualityScore *int32 `json:"quality_score,omitempty"`

	// A string that uniquely identifies this feedback resource
	Sid *string `json:"sid,omitempty"`
}

// AssertApiV2010AccountCallCallFeedbackRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountCallCallFeedbackRequired(obj ApiV2010AccountCallCallFeedback) error {
	return nil
}

// AssertRecurseApiV2010AccountCallCallFeedbackRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountCallCallFeedback (e.g. [][]ApiV2010AccountCallCallFeedback), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountCallCallFeedbackRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountCallCallFeedback, ok := obj.(ApiV2010AccountCallCallFeedback)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountCallCallFeedbackRequired(aApiV2010AccountCallCallFeedback)
	})
}
