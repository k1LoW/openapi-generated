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

// ListUsageRecordMonthlyResponse struct for ListUsageRecordMonthlyResponse
type ListUsageRecordMonthlyResponse struct {
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
	UsageRecords []ApiV2010AccountUsageUsageRecordUsageRecordMonthly `json:"usage_records,omitempty"`
}

// NewListUsageRecordMonthlyResponse instantiates a new ListUsageRecordMonthlyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsageRecordMonthlyResponse() *ListUsageRecordMonthlyResponse {
	this := ListUsageRecordMonthlyResponse{}
	return &this
}

// NewListUsageRecordMonthlyResponseWithDefaults instantiates a new ListUsageRecordMonthlyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsageRecordMonthlyResponseWithDefaults() *ListUsageRecordMonthlyResponse {
	this := ListUsageRecordMonthlyResponse{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListUsageRecordMonthlyResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListUsageRecordMonthlyResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListUsageRecordMonthlyResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListUsageRecordMonthlyResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListUsageRecordMonthlyResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListUsageRecordMonthlyResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListUsageRecordMonthlyResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListUsageRecordMonthlyResponse) SetUri(v string) {
	o.Uri = &v
}

// GetUsageRecords returns the UsageRecords field value if set, zero value otherwise.
func (o *ListUsageRecordMonthlyResponse) GetUsageRecords() []ApiV2010AccountUsageUsageRecordUsageRecordMonthly {
	if o == nil || o.UsageRecords == nil {
		var ret []ApiV2010AccountUsageUsageRecordUsageRecordMonthly
		return ret
	}
	return o.UsageRecords
}

// GetUsageRecordsOk returns a tuple with the UsageRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageRecordMonthlyResponse) GetUsageRecordsOk() ([]ApiV2010AccountUsageUsageRecordUsageRecordMonthly, bool) {
	if o == nil || o.UsageRecords == nil {
		return nil, false
	}
	return o.UsageRecords, true
}

// HasUsageRecords returns a boolean if a field has been set.
func (o *ListUsageRecordMonthlyResponse) HasUsageRecords() bool {
	if o != nil && o.UsageRecords != nil {
		return true
	}

	return false
}

// SetUsageRecords gets a reference to the given []ApiV2010AccountUsageUsageRecordUsageRecordMonthly and assigns it to the UsageRecords field.
func (o *ListUsageRecordMonthlyResponse) SetUsageRecords(v []ApiV2010AccountUsageUsageRecordUsageRecordMonthly) {
	o.UsageRecords = v
}

func (o ListUsageRecordMonthlyResponse) MarshalJSON() ([]byte, error) {
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

type NullableListUsageRecordMonthlyResponse struct {
	value *ListUsageRecordMonthlyResponse
	isSet bool
}

func (v NullableListUsageRecordMonthlyResponse) Get() *ListUsageRecordMonthlyResponse {
	return v.value
}

func (v *NullableListUsageRecordMonthlyResponse) Set(val *ListUsageRecordMonthlyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsageRecordMonthlyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsageRecordMonthlyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsageRecordMonthlyResponse(val *ListUsageRecordMonthlyResponse) *NullableListUsageRecordMonthlyResponse {
	return &NullableListUsageRecordMonthlyResponse{value: val, isSet: true}
}

func (v NullableListUsageRecordMonthlyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsageRecordMonthlyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


