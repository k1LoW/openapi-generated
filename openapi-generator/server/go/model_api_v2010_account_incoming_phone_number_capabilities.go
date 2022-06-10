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

// ApiV2010AccountIncomingPhoneNumberCapabilities - Indicate if a phone can receive calls or messages
type ApiV2010AccountIncomingPhoneNumberCapabilities struct {

	Fax bool `json:"fax,omitempty"`

	Mms bool `json:"mms,omitempty"`

	Sms bool `json:"sms,omitempty"`

	Voice bool `json:"voice,omitempty"`
}

// AssertApiV2010AccountIncomingPhoneNumberCapabilitiesRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountIncomingPhoneNumberCapabilitiesRequired(obj ApiV2010AccountIncomingPhoneNumberCapabilities) error {
	return nil
}

// AssertRecurseApiV2010AccountIncomingPhoneNumberCapabilitiesRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountIncomingPhoneNumberCapabilities (e.g. [][]ApiV2010AccountIncomingPhoneNumberCapabilities), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountIncomingPhoneNumberCapabilitiesRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountIncomingPhoneNumberCapabilities, ok := obj.(ApiV2010AccountIncomingPhoneNumberCapabilities)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountIncomingPhoneNumberCapabilitiesRequired(aApiV2010AccountIncomingPhoneNumberCapabilities)
	})
}
