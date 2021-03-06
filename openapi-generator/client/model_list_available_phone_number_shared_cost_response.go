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

// ListAvailablePhoneNumberSharedCostResponse struct for ListAvailablePhoneNumberSharedCostResponse
type ListAvailablePhoneNumberSharedCostResponse struct {
	AvailablePhoneNumbers []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost `json:"available_phone_numbers,omitempty"`
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListAvailablePhoneNumberSharedCostResponse instantiates a new ListAvailablePhoneNumberSharedCostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAvailablePhoneNumberSharedCostResponse() *ListAvailablePhoneNumberSharedCostResponse {
	this := ListAvailablePhoneNumberSharedCostResponse{}
	return &this
}

// NewListAvailablePhoneNumberSharedCostResponseWithDefaults instantiates a new ListAvailablePhoneNumberSharedCostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAvailablePhoneNumberSharedCostResponseWithDefaults() *ListAvailablePhoneNumberSharedCostResponse {
	this := ListAvailablePhoneNumberSharedCostResponse{}
	return &this
}

// GetAvailablePhoneNumbers returns the AvailablePhoneNumbers field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetAvailablePhoneNumbers() []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost {
	if o == nil || o.AvailablePhoneNumbers == nil {
		var ret []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost
		return ret
	}
	return o.AvailablePhoneNumbers
}

// GetAvailablePhoneNumbersOk returns a tuple with the AvailablePhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetAvailablePhoneNumbersOk() ([]ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost, bool) {
	if o == nil || o.AvailablePhoneNumbers == nil {
		return nil, false
	}
	return o.AvailablePhoneNumbers, true
}

// HasAvailablePhoneNumbers returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasAvailablePhoneNumbers() bool {
	if o != nil && o.AvailablePhoneNumbers != nil {
		return true
	}

	return false
}

// SetAvailablePhoneNumbers gets a reference to the given []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost and assigns it to the AvailablePhoneNumbers field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetAvailablePhoneNumbers(v []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberSharedCost) {
	o.AvailablePhoneNumbers = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberSharedCostResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListAvailablePhoneNumberSharedCostResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListAvailablePhoneNumberSharedCostResponse) MarshalJSON() ([]byte, error) {
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

type NullableListAvailablePhoneNumberSharedCostResponse struct {
	value *ListAvailablePhoneNumberSharedCostResponse
	isSet bool
}

func (v NullableListAvailablePhoneNumberSharedCostResponse) Get() *ListAvailablePhoneNumberSharedCostResponse {
	return v.value
}

func (v *NullableListAvailablePhoneNumberSharedCostResponse) Set(val *ListAvailablePhoneNumberSharedCostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAvailablePhoneNumberSharedCostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAvailablePhoneNumberSharedCostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAvailablePhoneNumberSharedCostResponse(val *ListAvailablePhoneNumberSharedCostResponse) *NullableListAvailablePhoneNumberSharedCostResponse {
	return &NullableListAvailablePhoneNumberSharedCostResponse{value: val, isSet: true}
}

func (v NullableListAvailablePhoneNumberSharedCostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAvailablePhoneNumberSharedCostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


