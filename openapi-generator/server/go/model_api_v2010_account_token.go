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

type ApiV2010AccountToken struct {

	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`

	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`

	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`

	// An array representing the ephemeral credentials
	IceServers *[]ApiV2010AccountTokenIceServersInner `json:"ice_servers,omitempty"`

	// The temporary password used for authenticating
	Password *string `json:"password,omitempty"`

	// The duration in seconds the credentials are valid
	Ttl *string `json:"ttl,omitempty"`

	// The temporary username that uniquely identifies a Token
	Username *string `json:"username,omitempty"`
}

// AssertApiV2010AccountTokenRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountTokenRequired(obj ApiV2010AccountToken) error {
	if obj.IceServers != nil {
		for _, el := range *obj.IceServers {
			if err := AssertApiV2010AccountTokenIceServersInnerRequired(el); err != nil {
				return err
			}
		}
	}
	return nil
}

// AssertRecurseApiV2010AccountTokenRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountToken (e.g. [][]ApiV2010AccountToken), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountTokenRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountToken, ok := obj.(ApiV2010AccountToken)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountTokenRequired(aApiV2010AccountToken)
	})
}
