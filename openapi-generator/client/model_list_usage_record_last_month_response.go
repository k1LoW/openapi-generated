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

// ListUsageRecordLastMonthResponse struct for ListUsageRecordLastMonthResponse
type ListUsageRecordLastMonthResponse struct {
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
	UsageRecords []ApiV2010AccountUsageUsageRecordUsageRecordLastMonth `json:"usage_records,omitempty"`
}

// NewListUsageRecordLastMonthResponse instantiates a new ListUsageRecordLastMonthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsageRecordLastMonthResponse() *ListUsageRecordLastMonthResponse {
	this := ListUsageRecordLastMonthResponse{}
	return &this
}

// NewListUsageRecordLastMonthResponseWithDefaults instantiates a new ListUsageRecordLastMonthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsageRecordLastMonthResponseWithDefaults() *ListUsageRecordLastMonthResponse {
	this := ListUsageRecordLastMonthResponse{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListUsageRecordLastMonthResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListUsageRecordLastMonthResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListUsageRecordLastMonthResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListUsageRecordLastMonthResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListUsageRecordLastMonthResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListUsageRecordLastMonthResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListUsageRecordLastMonthResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListUsageRecordLastMonthResponse) SetUri(v string) {
	o.Uri = &v
}

// GetUsageRecords returns the UsageRecords field value if set, zero value otherwise.
func (o *ListUsageRecordLastMonthResponse) GetUsageRecords() []ApiV2010AccountUsageUsageRecordUsageRecordLastMonth {
	if o == nil || o.UsageRecords == nil {
		var ret []ApiV2010AccountUsageUsageRecordUsageRecordLastMonth
		return ret
	}
	return o.UsageRecords
}

// GetUsageRecordsOk returns a tuple with the UsageRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordLastMonthResponse) GetUsageRecordsOk() ([]ApiV2010AccountUsageUsageRecordUsageRecordLastMonth, bool) {
	if o == nil || o.UsageRecords == nil {
		return nil, false
	}
	return o.UsageRecords, true
}

// HasUsageRecords returns a boolean if a field has been set.
func (o *ListUsageRecordLastMonthResponse) HasUsageRecords() bool {
	if o != nil && o.UsageRecords != nil {
		return true
	}

	return false
}

// SetUsageRecords gets a reference to the given []ApiV2010AccountUsageUsageRecordUsageRecordLastMonth and assigns it to the UsageRecords field.
func (o *ListUsageRecordLastMonthResponse) SetUsageRecords(v []ApiV2010AccountUsageUsageRecordUsageRecordLastMonth) {
	o.UsageRecords = v
}

func (o ListUsageRecordLastMonthResponse) MarshalJSON() ([]byte, error) {
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
	if o.PreviousPageUri != nil {
		toSerialize["previous_page_uri"] = o.PreviousPageUri
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.UsageRecords != nil {
		toSerialize["usage_records"] = o.UsageRecords
	}
	return json.Marshal(toSerialize)
}

type NullableListUsageRecordLastMonthResponse struct {
	value *ListUsageRecordLastMonthResponse
	isSet bool
}

func (v NullableListUsageRecordLastMonthResponse) Get() *ListUsageRecordLastMonthResponse {
	return v.value
}

func (v *NullableListUsageRecordLastMonthResponse) Set(val *ListUsageRecordLastMonthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsageRecordLastMonthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsageRecordLastMonthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsageRecordLastMonthResponse(val *ListUsageRecordLastMonthResponse) *NullableListUsageRecordLastMonthResponse {
	return &NullableListUsageRecordLastMonthResponse{value: val, isSet: true}
}

func (v NullableListUsageRecordLastMonthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsageRecordLastMonthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


