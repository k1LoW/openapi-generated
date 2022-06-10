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

// ListAvailablePhoneNumberLocalResponse struct for ListAvailablePhoneNumberLocalResponse
type ListAvailablePhoneNumberLocalResponse struct {
	AvailablePhoneNumbers []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal `json:"available_phone_numbers,omitempty"`
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListAvailablePhoneNumberLocalResponse instantiates a new ListAvailablePhoneNumberLocalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAvailablePhoneNumberLocalResponse() *ListAvailablePhoneNumberLocalResponse {
	this := ListAvailablePhoneNumberLocalResponse{}
	return &this
}

// NewListAvailablePhoneNumberLocalResponseWithDefaults instantiates a new ListAvailablePhoneNumberLocalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAvailablePhoneNumberLocalResponseWithDefaults() *ListAvailablePhoneNumberLocalResponse {
	this := ListAvailablePhoneNumberLocalResponse{}
	return &this
}

// GetAvailablePhoneNumbers returns the AvailablePhoneNumbers field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetAvailablePhoneNumbers() []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal {
	if o == nil || o.AvailablePhoneNumbers == nil {
		var ret []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal
		return ret
	}
	return o.AvailablePhoneNumbers
}

// GetAvailablePhoneNumbersOk returns a tuple with the AvailablePhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetAvailablePhoneNumbersOk() ([]ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal, bool) {
	if o == nil || o.AvailablePhoneNumbers == nil {
		return nil, false
	}
	return o.AvailablePhoneNumbers, true
}

// HasAvailablePhoneNumbers returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasAvailablePhoneNumbers() bool {
	if o != nil && o.AvailablePhoneNumbers != nil {
		return true
	}

	return false
}

// SetAvailablePhoneNumbers gets a reference to the given []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal and assigns it to the AvailablePhoneNumbers field.
func (o *ListAvailablePhoneNumberLocalResponse) SetAvailablePhoneNumbers(v []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) {
	o.AvailablePhoneNumbers = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListAvailablePhoneNumberLocalResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListAvailablePhoneNumberLocalResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListAvailablePhoneNumberLocalResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListAvailablePhoneNumberLocalResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListAvailablePhoneNumberLocalResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListAvailablePhoneNumberLocalResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListAvailablePhoneNumberLocalResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberLocalResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberLocalResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberLocalResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListAvailablePhoneNumberLocalResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListAvailablePhoneNumberLocalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailablePhoneNumbers != nil {
		toSerialize["available_phone_numbers"] = o.AvailablePhoneNumbers
	}
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

type NullableListAvailablePhoneNumberLocalResponse struct {
	value *ListAvailablePhoneNumberLocalResponse
	isSet bool
}

func (v NullableListAvailablePhoneNumberLocalResponse) Get() *ListAvailablePhoneNumberLocalResponse {
	return v.value
}

func (v *NullableListAvailablePhoneNumberLocalResponse) Set(val *ListAvailablePhoneNumberLocalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAvailablePhoneNumberLocalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAvailablePhoneNumberLocalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAvailablePhoneNumberLocalResponse(val *ListAvailablePhoneNumberLocalResponse) *NullableListAvailablePhoneNumberLocalResponse {
	return &NullableListAvailablePhoneNumberLocalResponse{value: val, isSet: true}
}

func (v NullableListAvailablePhoneNumberLocalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAvailablePhoneNumberLocalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


