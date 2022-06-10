/*
Twilio - Api

This is the public Twilio REST API.

API version: 1.29.1
Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListRecordingAddOnResultPayloadResponse struct for ListRecordingAddOnResultPayloadResponse
type ListRecordingAddOnResultPayloadResponse struct {
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	Payloads []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload `json:"payloads,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListRecordingAddOnResultPayloadResponse instantiates a new ListRecordingAddOnResultPayloadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRecordingAddOnResultPayloadResponse() *ListRecordingAddOnResultPayloadResponse {
	this := ListRecordingAddOnResultPayloadResponse{}
	return &this
}

// NewListRecordingAddOnResultPayloadResponseWithDefaults instantiates a new ListRecordingAddOnResultPayloadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRecordingAddOnResultPayloadResponseWithDefaults() *ListRecordingAddOnResultPayloadResponse {
	this := ListRecordingAddOnResultPayloadResponse{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListRecordingAddOnResultPayloadResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListRecordingAddOnResultPayloadResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListRecordingAddOnResultPayloadResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListRecordingAddOnResultPayloadResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListRecordingAddOnResultPayloadResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPayloads returns the Payloads field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetPayloads() []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload {
	if o == nil || o.Payloads == nil {
		var ret []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload
		return ret
	}
	return o.Payloads
}

// GetPayloadsOk returns a tuple with the Payloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetPayloadsOk() ([]ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload, bool) {
	if o == nil || o.Payloads == nil {
		return nil, false
	}
	return o.Payloads, true
}

// HasPayloads returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasPayloads() bool {
	if o != nil && o.Payloads != nil {
		return true
	}

	return false
}

// SetPayloads gets a reference to the given []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload and assigns it to the Payloads field.
func (o *ListRecordingAddOnResultPayloadResponse) SetPayloads(v []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) {
	o.Payloads = v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListRecordingAddOnResultPayloadResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListRecordingAddOnResultPayloadResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListRecordingAddOnResultPayloadResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordingAddOnResultPayloadResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListRecordingAddOnResultPayloadResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListRecordingAddOnResultPayloadResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListRecordingAddOnResultPayloadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.FirstPageUri != nil {
		toSerialize["first_page_uri"] = o.FirstPageUri
	}
	if o.NextPageUri != nil {
		toSerialize["next_page_uri"] = o.NextPageUri
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Payloads != nil {
		toSerialize["payloads"] = o.Payloads
	}
	if o.PreviousPageUri != nil {
		toSerialize["previous_page_uri"] = o.PreviousPageUri
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableListRecordingAddOnResultPayloadResponse struct {
	value *ListRecordingAddOnResultPayloadResponse
	isSet bool
}

func (v NullableListRecordingAddOnResultPayloadResponse) Get() *ListRecordingAddOnResultPayloadResponse {
	return v.value
}

func (v *NullableListRecordingAddOnResultPayloadResponse) Set(val *ListRecordingAddOnResultPayloadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordingAddOnResultPayloadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordingAddOnResultPayloadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordingAddOnResultPayloadResponse(val *ListRecordingAddOnResultPayloadResponse) *NullableListRecordingAddOnResultPayloadResponse {
	return &NullableListRecordingAddOnResultPayloadResponse{value: val, isSet: true}
}

func (v NullableListRecordingAddOnResultPayloadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordingAddOnResultPayloadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


