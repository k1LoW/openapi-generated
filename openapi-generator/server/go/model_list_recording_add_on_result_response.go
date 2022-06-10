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

type ListRecordingAddOnResultResponse struct {

	AddOnResults []ApiV2010AccountRecordingRecordingAddOnResult `json:"add_on_results,omitempty"`

	End int32 `json:"end,omitempty"`

	FirstPageUri string `json:"first_page_uri,omitempty"`

	NextPageUri string `json:"next_page_uri,omitempty"`

	Page int32 `json:"page,omitempty"`

	PageSize int32 `json:"page_size,omitempty"`

	PreviousPageUri string `json:"previous_page_uri,omitempty"`

	Start int32 `json:"start,omitempty"`

	Uri string `json:"uri,omitempty"`
}

// AssertListRecordingAddOnResultResponseRequired checks if the required fields are not zero-ed
func AssertListRecordingAddOnResultResponseRequired(obj ListRecordingAddOnResultResponse) error {
	for _, el := range obj.AddOnResults {
		if err := AssertApiV2010AccountRecordingRecordingAddOnResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseListRecordingAddOnResultResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ListRecordingAddOnResultResponse (e.g. [][]ListRecordingAddOnResultResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseListRecordingAddOnResultResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aListRecordingAddOnResultResponse, ok := obj.(ListRecordingAddOnResultResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertListRecordingAddOnResultResponseRequired(aListRecordingAddOnResultResponse)
	})
}