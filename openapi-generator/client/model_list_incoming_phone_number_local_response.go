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

// ListIncomingPhoneNumberLocalResponse struct for ListIncomingPhoneNumberLocalResponse
type ListIncomingPhoneNumberLocalResponse struct {
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	IncomingPhoneNumbers []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal `json:"incoming_phone_numbers,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListIncomingPhoneNumberLocalResponse instantiates a new ListIncomingPhoneNumberLocalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIncomingPhoneNumberLocalResponse() *ListIncomingPhoneNumberLocalResponse {
	this := ListIncomingPhoneNumberLocalResponse{}
	return &this
}

// NewListIncomingPhoneNumberLocalResponseWithDefaults instantiates a new ListIncomingPhoneNumberLocalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIncomingPhoneNumberLocalResponseWithDefaults() *ListIncomingPhoneNumberLocalResponse {
	this := ListIncomingPhoneNumberLocalResponse{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListIncomingPhoneNumberLocalResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListIncomingPhoneNumberLocalResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetIncomingPhoneNumbers returns the IncomingPhoneNumbers field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetIncomingPhoneNumbers() []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal {
	if o == nil || o.IncomingPhoneNumbers == nil {
		var ret []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal
		return ret
	}
	return o.IncomingPhoneNumbers
}

// GetIncomingPhoneNumbersOk returns a tuple with the IncomingPhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetIncomingPhoneNumbersOk() ([]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal, bool) {
	if o == nil || o.IncomingPhoneNumbers == nil {
		return nil, false
	}
	return o.IncomingPhoneNumbers, true
}

// HasIncomingPhoneNumbers returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasIncomingPhoneNumbers() bool {
	if o != nil && o.IncomingPhoneNumbers != nil {
		return true
	}

	return false
}

// SetIncomingPhoneNumbers gets a reference to the given []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal and assigns it to the IncomingPhoneNumbers field.
func (o *ListIncomingPhoneNumberLocalResponse) SetIncomingPhoneNumbers(v []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal) {
	o.IncomingPhoneNumbers = v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListIncomingPhoneNumberLocalResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListIncomingPhoneNumberLocalResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListIncomingPhoneNumberLocalResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListIncomingPhoneNumberLocalResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListIncomingPhoneNumberLocalResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberLocalResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberLocalResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberLocalResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListIncomingPhoneNumberLocalResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListIncomingPhoneNumberLocalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.FirstPageUri != nil {
		toSerialize["first_page_uri"] = o.FirstPageUri
	}
	if o.IncomingPhoneNumbers != nil {
		toSerialize["incoming_phone_numbers"] = o.IncomingPhoneNumbers
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

type NullableListIncomingPhoneNumberLocalResponse struct {
	value *ListIncomingPhoneNumberLocalResponse
	isSet bool
}

func (v NullableListIncomingPhoneNumberLocalResponse) Get() *ListIncomingPhoneNumberLocalResponse {
	return v.value
}

func (v *NullableListIncomingPhoneNumberLocalResponse) Set(val *ListIncomingPhoneNumberLocalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListIncomingPhoneNumberLocalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListIncomingPhoneNumberLocalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIncomingPhoneNumberLocalResponse(val *ListIncomingPhoneNumberLocalResponse) *NullableListIncomingPhoneNumberLocalResponse {
	return &NullableListIncomingPhoneNumberLocalResponse{value: val, isSet: true}
}

func (v NullableListIncomingPhoneNumberLocalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIncomingPhoneNumberLocalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

