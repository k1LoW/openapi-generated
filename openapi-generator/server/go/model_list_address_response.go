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

type ListAddressResponse struct {

	Addresses []ApiV2010AccountAddress `json:"addresses,omitempty"`

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListAddressResponseRequired checks if the required fields are not zero-ed
func AssertListAddressResponseRequired(obj ListAddressResponse) error {
	for _, el := range obj.Addresses {
		if err := AssertApiV2010AccountAddressRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListAddressResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListAddressResponse (e.g. [][]ListAddressResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListAddressResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListAddressResponse, ok := obj.(ListAddressResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListAddressResponseRequired(aListAddressResponse)
	})
}
