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

type ListConferenceRecordingResponse struct {

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Recordings []ApiV2010AccountConferenceConferenceRecording `json:"recordings,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListConferenceRecordingResponseRequired checks if the required fields are not zero-ed
func AssertListConferenceRecordingResponseRequired(obj ListConferenceRecordingResponse) error {
	for _, el := range obj.Recordings {
		if err := AssertApiV2010AccountConferenceConferenceRecordingRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListConferenceRecordingResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListConferenceRecordingResponse (e.g. [][]ListConferenceRecordingResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListConferenceRecordingResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListConferenceRecordingResponse, ok := obj.(ListConferenceRecordingResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListConferenceRecordingResponseRequired(aListConferenceRecordingResponse)
	})
}