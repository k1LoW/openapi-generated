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

type ApiV2010AccountRecordingRecordingAddOnResult struct {

	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`

	// The SID of the Add-on configuration
	AddOnConfigurationSid *string `json:"add_on_configuration_sid,omitempty"`

	// The SID of the Add-on to which the result belongs
	AddOnSid *string `json:"add_on_sid,omitempty"`

	// The date and time in GMT that the result was completed
	DateCompleted *string `json:"date_completed,omitempty"`

	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`

	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`

	// The SID of the recording to which the AddOnResult resource belongs
	ReferenceSid *string `json:"reference_sid,omitempty"`

	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`

	// The status of the result
	Status *string `json:"status,omitempty"`

	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
}

// AssertApiV2010AccountRecordingRecordingAddOnResultRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountRecordingRecordingAddOnResultRequired(obj ApiV2010AccountRecordingRecordingAddOnResult) error {
	return nil
}

// AssertRecurseApiV2010AccountRecordingRecordingAddOnResultRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountRecordingRecordingAddOnResult (e.g. [][]ApiV2010AccountRecordingRecordingAddOnResult), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountRecordingRecordingAddOnResultRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountRecordingRecordingAddOnResult, ok := obj.(ApiV2010AccountRecordingRecordingAddOnResult)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountRecordingRecordingAddOnResultRequired(aApiV2010AccountRecordingRecordingAddOnResult)
	})
}