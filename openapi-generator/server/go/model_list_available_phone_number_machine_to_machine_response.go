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

type ListAvailablePhoneNumberMachineToMachineResponse struct {

	AvailablePhoneNumbers []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine `json:"available_phone_numbers,omitempty"`

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListAvailablePhoneNumberMachineToMachineResponseRequired checks if the required fields are not zero-ed
func AssertListAvailablePhoneNumberMachineToMachineResponseRequired(obj ListAvailablePhoneNumberMachineToMachineResponse) error {
	for _, el := range obj.AvailablePhoneNumbers {
		if err := AssertApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachineRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListAvailablePhoneNumberMachineToMachineResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListAvailablePhoneNumberMachineToMachineResponse (e.g. [][]ListAvailablePhoneNumberMachineToMachineResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListAvailablePhoneNumberMachineToMachineResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListAvailablePhoneNumberMachineToMachineResponse, ok := obj.(ListAvailablePhoneNumberMachineToMachineResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListAvailablePhoneNumberMachineToMachineResponseRequired(aListAvailablePhoneNumberMachineToMachineResponse)
	})
}