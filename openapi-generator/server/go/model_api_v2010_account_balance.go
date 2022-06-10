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

type ApiV2010AccountBalance struct {

	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`

	// Account balance
	Balance *string `json:"balance,omitempty"`

	// Currency units
	Currency *string `json:"currency,omitempty"`
}

// AssertApiV2010AccountBalanceRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountBalanceRequired(obj ApiV2010AccountBalance) error {
	return nil
}

// AssertRecurseApiV2010AccountBalanceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountBalance (e.g. [][]ApiV2010AccountBalance), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountBalanceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountBalance, ok := obj.(ApiV2010AccountBalance)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountBalanceRequired(aApiV2010AccountBalance)
	})
}
