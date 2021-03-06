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

// ListIncomingPhoneNumberAssignedAddOnResponse struct for ListIncomingPhoneNumberAssignedAddOnResponse
type ListIncomingPhoneNumberAssignedAddOnResponse struct {
	AssignedAddOns []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn `json:"assigned_add_ons,omitempty"`
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListIncomingPhoneNumberAssignedAddOnResponse instantiates a new ListIncomingPhoneNumberAssignedAddOnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIncomingPhoneNumberAssignedAddOnResponse() *ListIncomingPhoneNumberAssignedAddOnResponse {
	this := ListIncomingPhoneNumberAssignedAddOnResponse{}
	return &this
}

// NewListIncomingPhoneNumberAssignedAddOnResponseWithDefaults instantiates a new ListIncomingPhoneNumberAssignedAddOnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIncomingPhoneNumberAssignedAddOnResponseWithDefaults() *ListIncomingPhoneNumberAssignedAddOnResponse {
	this := ListIncomingPhoneNumberAssignedAddOnResponse{}
	return &this
}

// GetAssignedAddOns returns the AssignedAddOns field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetAssignedAddOns() []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn {
	if o == nil || o.AssignedAddOns == nil {
		var ret []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
		return ret
	}
	return o.AssignedAddOns
}

// GetAssignedAddOnsOk returns a tuple with the AssignedAddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetAssignedAddOnsOk() ([]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn, bool) {
	if o == nil || o.AssignedAddOns == nil {
		return nil, false
	}
	return o.AssignedAddOns, true
}

// HasAssignedAddOns returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasAssignedAddOns() bool {
	if o != nil && o.AssignedAddOns != nil {
		return true
	}

	return false
}

// SetAssignedAddOns gets a reference to the given []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn and assigns it to the AssignedAddOns field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetAssignedAddOns(v []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) {
	o.AssignedAddOns = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListIncomingPhoneNumberAssignedAddOnResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListIncomingPhoneNumberAssignedAddOnResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedAddOns != nil {
		toSerialize["assigned_add_ons"] = o.AssignedAddOns
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

type NullableListIncomingPhoneNumberAssignedAddOnResponse struct {
	value *ListIncomingPhoneNumberAssignedAddOnResponse
	isSet bool
}

func (v NullableListIncomingPhoneNumberAssignedAddOnResponse) Get() *ListIncomingPhoneNumberAssignedAddOnResponse {
	return v.value
}

func (v *NullableListIncomingPhoneNumberAssignedAddOnResponse) Set(val *ListIncomingPhoneNumberAssignedAddOnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListIncomingPhoneNumberAssignedAddOnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListIncomingPhoneNumberAssignedAddOnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIncomingPhoneNumberAssignedAddOnResponse(val *ListIncomingPhoneNumberAssignedAddOnResponse) *NullableListIncomingPhoneNumberAssignedAddOnResponse {
	return &NullableListIncomingPhoneNumberAssignedAddOnResponse{value: val, isSet: true}
}

func (v NullableListIncomingPhoneNumberAssignedAddOnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIncomingPhoneNumberAssignedAddOnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


