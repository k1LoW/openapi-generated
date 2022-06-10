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

type ApiV2010AccountSipSipDomainSipCredentialListMapping struct {

	// The unique id of the Account that is responsible for this resource.
	AccountSid *string `json:"account_sid,omitempty"`

	// The date that this resource was created, given as GMT in RFC 2822 format.
	DateCreated *string `json:"date_created,omitempty"`

	// The date that this resource was last updated, given as GMT in RFC 2822 format.
	DateUpdated *string `json:"date_updated,omitempty"`

	// The unique string that identifies the SipDomain resource.
	DomainSid *string `json:"domain_sid,omitempty"`

	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName *string `json:"friendly_name,omitempty"`

	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`

	// The URI for this resource, relative to https://api.twilio.com
	Uri *string `json:"uri,omitempty"`
}

// AssertApiV2010AccountSipSipDomainSipCredentialListMappingRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountSipSipDomainSipCredentialListMappingRequired(obj ApiV2010AccountSipSipDomainSipCredentialListMapping) error {
	return nil
}

// AssertRecurseApiV2010AccountSipSipDomainSipCredentialListMappingRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountSipSipDomainSipCredentialListMapping (e.g. [][]ApiV2010AccountSipSipDomainSipCredentialListMapping), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountSipSipDomainSipCredentialListMappingRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountSipSipDomainSipCredentialListMapping, ok := obj.(ApiV2010AccountSipSipDomainSipCredentialListMapping)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountSipSipDomainSipCredentialListMappingRequired(aApiV2010AccountSipSipDomainSipCredentialListMapping)
	})
}