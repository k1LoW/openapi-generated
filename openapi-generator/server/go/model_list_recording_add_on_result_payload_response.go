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

type ListRecordingAddOnResultPayloadResponse struct {

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	Payloads []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload `json:"payloads,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListRecordingAddOnResultPayloadResponseRequired checks if the required fields are not zero-ed
func AssertListRecordingAddOnResultPayloadResponseRequired(obj ListRecordingAddOnResultPayloadResponse) error {
	for _, el := range obj.Payloads {
		if err := AssertApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayloadRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListRecordingAddOnResultPayloadResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListRecordingAddOnResultPayloadResponse (e.g. [][]ListRecordingAddOnResultPayloadResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListRecordingAddOnResultPayloadResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListRecordingAddOnResultPayloadResponse, ok := obj.(ListRecordingAddOnResultPayloadResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListRecordingAddOnResultPayloadResponseRequired(aListRecordingAddOnResultPayloadResponse)
	})
}