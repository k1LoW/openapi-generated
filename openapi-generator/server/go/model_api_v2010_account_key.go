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

type ApiV2010AccountKey struct {

	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`

	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`

	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
}

// AssertApiV2010AccountKeyRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountKeyRequired(obj ApiV2010AccountKey) error {
	return nil
}

// AssertRecurseApiV2010AccountKeyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountKey (e.g. [][]ApiV2010AccountKey), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountKeyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountKey, ok := obj.(ApiV2010AccountKey)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountKeyRequired(aApiV2010AccountKey)
	})
}
