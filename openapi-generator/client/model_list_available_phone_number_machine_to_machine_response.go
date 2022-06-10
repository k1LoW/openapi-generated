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

// ListAvailablePhoneNumberMachineToMachineResponse struct for ListAvailablePhoneNumberMachineToMachineResponse
type ListAvailablePhoneNumberMachineToMachineResponse struct {
	AvailablePhoneNumbers []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine `json:"available_phone_numbers,omitempty"`
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListAvailablePhoneNumberMachineToMachineResponse instantiates a new ListAvailablePhoneNumberMachineToMachineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAvailablePhoneNumberMachineToMachineResponse() *ListAvailablePhoneNumberMachineToMachineResponse {
	this := ListAvailablePhoneNumberMachineToMachineResponse{}
	return &this
}

// NewListAvailablePhoneNumberMachineToMachineResponseWithDefaults instantiates a new ListAvailablePhoneNumberMachineToMachineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAvailablePhoneNumberMachineToMachineResponseWithDefaults() *ListAvailablePhoneNumberMachineToMachineResponse {
	this := ListAvailablePhoneNumberMachineToMachineResponse{}
	return &this
}

// GetAvailablePhoneNumbers returns the AvailablePhoneNumbers field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetAvailablePhoneNumbers() []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine {
	if o == nil || o.AvailablePhoneNumbers == nil {
		var ret []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine
		return ret
	}
	return o.AvailablePhoneNumbers
}

// GetAvailablePhoneNumbersOk returns a tuple with the AvailablePhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetAvailablePhoneNumbersOk() ([]ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine, bool) {
	if o == nil || o.AvailablePhoneNumbers == nil {
		return nil, false
	}
	return o.AvailablePhoneNumbers, true
}

// HasAvailablePhoneNumbers returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasAvailablePhoneNumbers() bool {
	if o != nil && o.AvailablePhoneNumbers != nil {
		return true
	}

	return false
}

// SetAvailablePhoneNumbers gets a reference to the given []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine and assigns it to the AvailablePhoneNumbers field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetAvailablePhoneNumbers(v []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine) {
	o.AvailablePhoneNumbers = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListAvailablePhoneNumberMachineToMachineResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListAvailablePhoneNumberMachineToMachineResponse) MarshalJSON() ([]byte, error) {
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

type NullableListAvailablePhoneNumberMachineToMachineResponse struct {
	value *ListAvailablePhoneNumberMachineToMachineResponse
	isSet bool
}

func (v NullableListAvailablePhoneNumberMachineToMachineResponse) Get() *ListAvailablePhoneNumberMachineToMachineResponse {
	return v.value
}

func (v *NullableListAvailablePhoneNumberMachineToMachineResponse) Set(val *ListAvailablePhoneNumberMachineToMachineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAvailablePhoneNumberMachineToMachineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAvailablePhoneNumberMachineToMachineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAvailablePhoneNumberMachineToMachineResponse(val *ListAvailablePhoneNumberMachineToMachineResponse) *NullableListAvailablePhoneNumberMachineToMachineResponse {
	return &NullableListAvailablePhoneNumberMachineToMachineResponse{value: val, isSet: true}
}

func (v NullableListAvailablePhoneNumberMachineToMachineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAvailablePhoneNumberMachineToMachineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


