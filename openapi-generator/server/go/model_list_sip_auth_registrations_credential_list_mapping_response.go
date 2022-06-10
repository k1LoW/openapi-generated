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

type ListSipAuthRegistrationsCredentialListMappingResponse struct {

	Contents []ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping `json:"contents,omitempty"`

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListSipAuthRegistrationsCredentialListMappingResponseRequired checks if the required fields are not zero-ed
func AssertListSipAuthRegistrationsCredentialListMappingResponseRequired(obj ListSipAuthRegistrationsCredentialListMappingResponse) error {
	for _, el := range obj.Contents {
		if err := AssertApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMappingRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListSipAuthRegistrationsCredentialListMappingResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListSipAuthRegistrationsCredentialListMappingResponse (e.g. [][]ListSipAuthRegistrationsCredentialListMappingResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListSipAuthRegistrationsCredentialListMappingResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListSipAuthRegistrationsCredentialListMappingResponse, ok := obj.(ListSipAuthRegistrationsCredentialListMappingResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListSipAuthRegistrationsCredentialListMappingResponseRequired(aListSipAuthRegistrationsCredentialListMappingResponse)
	})
}
