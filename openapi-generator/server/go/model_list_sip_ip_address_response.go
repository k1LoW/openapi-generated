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

type ListSipIpAddressResponse struct {

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	IpAddresses []ApiV2010AccountSipSipIpAccessControlListSipIpAddress `json:"ip_addresses,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListSipIpAddressResponseRequired checks if the required fields are not zero-ed
func AssertListSipIpAddressResponseRequired(obj ListSipIpAddressResponse) error {
	for _, el := range obj.IpAddresses {
		if err := AssertApiV2010AccountSipSipIpAccessControlListSipIpAddressRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListSipIpAddressResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListSipIpAddressResponse (e.g. [][]ListSipIpAddressResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListSipIpAddressResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListSipIpAddressResponse, ok := obj.(ListSipIpAddressResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListSipIpAddressResponseRequired(aListSipIpAddressResponse)
	})
}
