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

// ListSipCredentialListMappingResponse struct for ListSipCredentialListMappingResponse
type ListSipCredentialListMappingResponse struct {
	CredentialListMappings []ApiV2010AccountSipSipDomainSipCredentialListMapping `json:"credential_list_mappings,omitempty"`
	End *int32 `json:"end,omitempty"`
	FirstPageUri *string `json:"first_page_uri,omitempty"`
	NextPageUri *string `json:"next_page_uri,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PreviousPageUri *string `json:"previous_page_uri,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewListSipCredentialListMappingResponse instantiates a new ListSipCredentialListMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSipCredentialListMappingResponse() *ListSipCredentialListMappingResponse {
	this := ListSipCredentialListMappingResponse{}
	return &this
}

// NewListSipCredentialListMappingResponseWithDefaults instantiates a new ListSipCredentialListMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSipCredentialListMappingResponseWithDefaults() *ListSipCredentialListMappingResponse {
	this := ListSipCredentialListMappingResponse{}
	return &this
}

// GetCredentialListMappings returns the CredentialListMappings field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetCredentialListMappings() []ApiV2010AccountSipSipDomainSipCredentialListMapping {
	if o == nil || o.CredentialListMappings == nil {
		var ret []ApiV2010AccountSipSipDomainSipCredentialListMapping
		return ret
	}
	return o.CredentialListMappings
}

// GetCredentialListMappingsOk returns a tuple with the CredentialListMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetCredentialListMappingsOk() ([]ApiV2010AccountSipSipDomainSipCredentialListMapping, bool) {
	if o == nil || o.CredentialListMappings == nil {
		return nil, false
	}
	return o.CredentialListMappings, true
}

// HasCredentialListMappings returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasCredentialListMappings() bool {
	if o != nil && o.CredentialListMappings != nil {
		return true
	}

	return false
}

// SetCredentialListMappings gets a reference to the given []ApiV2010AccountSipSipDomainSipCredentialListMapping and assigns it to the CredentialListMappings field.
func (o *ListSipCredentialListMappingResponse) SetCredentialListMappings(v []ApiV2010AccountSipSipDomainSipCredentialListMapping) {
	o.CredentialListMappings = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ListSipCredentialListMappingResponse) SetEnd(v int32) {
	o.End = &v
}

// GetFirstPageUri returns the FirstPageUri field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetFirstPageUri() string {
	if o == nil || o.FirstPageUri == nil {
		var ret string
		return ret
	}
	return *o.FirstPageUri
}

// GetFirstPageUriOk returns a tuple with the FirstPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetFirstPageUriOk() (*string, bool) {
	if o == nil || o.FirstPageUri == nil {
		return nil, false
	}
	return o.FirstPageUri, true
}

// HasFirstPageUri returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasFirstPageUri() bool {
	if o != nil && o.FirstPageUri != nil {
		return true
	}

	return false
}

// SetFirstPageUri gets a reference to the given string and assigns it to the FirstPageUri field.
func (o *ListSipCredentialListMappingResponse) SetFirstPageUri(v string) {
	o.FirstPageUri = &v
}

// GetNextPageUri returns the NextPageUri field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetNextPageUri() string {
	if o == nil || o.NextPageUri == nil {
		var ret string
		return ret
	}
	return *o.NextPageUri
}

// GetNextPageUriOk returns a tuple with the NextPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetNextPageUriOk() (*string, bool) {
	if o == nil || o.NextPageUri == nil {
		return nil, false
	}
	return o.NextPageUri, true
}

// HasNextPageUri returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasNextPageUri() bool {
	if o != nil && o.NextPageUri != nil {
		return true
	}

	return false
}

// SetNextPageUri gets a reference to the given string and assigns it to the NextPageUri field.
func (o *ListSipCredentialListMappingResponse) SetNextPageUri(v string) {
	o.NextPageUri = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListSipCredentialListMappingResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListSipCredentialListMappingResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPreviousPageUri returns the PreviousPageUri field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetPreviousPageUri() string {
	if o == nil || o.PreviousPageUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousPageUri
}

// GetPreviousPageUriOk returns a tuple with the PreviousPageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetPreviousPageUriOk() (*string, bool) {
	if o == nil || o.PreviousPageUri == nil {
		return nil, false
	}
	return o.PreviousPageUri, true
}

// HasPreviousPageUri returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasPreviousPageUri() bool {
	if o != nil && o.PreviousPageUri != nil {
		return true
	}

	return false
}

// SetPreviousPageUri gets a reference to the given string and assigns it to the PreviousPageUri field.
func (o *ListSipCredentialListMappingResponse) SetPreviousPageUri(v string) {
	o.PreviousPageUri = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ListSipCredentialListMappingResponse) SetStart(v int32) {
	o.Start = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ListSipCredentialListMappingResponse) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSipCredentialListMappingResponse) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ListSipCredentialListMappingResponse) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ListSipCredentialListMappingResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ListSipCredentialListMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialListMappings != nil {
		toSerialize["credential_list_mappings"] = o.CredentialListMappings
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

type NullableListSipCredentialListMappingResponse struct {
	value *ListSipCredentialListMappingResponse
	isSet bool
}

func (v NullableListSipCredentialListMappingResponse) Get() *ListSipCredentialListMappingResponse {
	return v.value
}

func (v *NullableListSipCredentialListMappingResponse) Set(val *ListSipCredentialListMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSipCredentialListMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSipCredentialListMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSipCredentialListMappingResponse(val *ListSipCredentialListMappingResponse) *NullableListSipCredentialListMappingResponse {
	return &NullableListSipCredentialListMappingResponse{value: val, isSet: true}
}

func (v NullableListSipCredentialListMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSipCredentialListMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


