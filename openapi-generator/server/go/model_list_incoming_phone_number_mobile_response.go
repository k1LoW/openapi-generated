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

type ListIncomingPhoneNumberMobileResponse struct {

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	IncomingPhoneNumbers []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile `json:"incoming_phone_numbers,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListIncomingPhoneNumberMobileResponseRequired checks if the required fields are not zero-ed
func AssertListIncomingPhoneNumberMobileResponseRequired(obj ListIncomingPhoneNumberMobileResponse) error {
	for _, el := range obj.IncomingPhoneNumbers {
		if err := AssertApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobileRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListIncomingPhoneNumberMobileResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListIncomingPhoneNumberMobileResponse (e.g. [][]ListIncomingPhoneNumberMobileResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListIncomingPhoneNumberMobileResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListIncomingPhoneNumberMobileResponse, ok := obj.(ListIncomingPhoneNumberMobileResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListIncomingPhoneNumberMobileResponseRequired(aListIncomingPhoneNumberMobileResponse)
	})
}