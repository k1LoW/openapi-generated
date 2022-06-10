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

type ListSipCredentialListMappingResponse struct {

	CredentialListMappings []ApiV2010AccountSipSipDomainSipCredentialListMapping `json:"credential_list_mappings,omitempty"`

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListSipCredentialListMappingResponseRequired checks if the required fields are not zero-ed
func AssertListSipCredentialListMappingResponseRequired(obj ListSipCredentialListMappingResponse) error {
	for _, el := range obj.CredentialListMappings {
		if err := AssertApiV2010AccountSipSipDomainSipCredentialListMappingRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListSipCredentialListMappingResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListSipCredentialListMappingResponse (e.g. [][]ListSipCredentialListMappingResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListSipCredentialListMappingResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListSipCredentialListMappingResponse, ok := obj.(ListSipCredentialListMappingResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListSipCredentialListMappingResponseRequired(aListSipCredentialListMappingResponse)
	})
}