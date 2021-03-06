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

// ApiV2010AccountSipSipIpAccessControlList struct for ApiV2010AccountSipSipIpAccessControlList
type ApiV2010AccountSipSipIpAccessControlList struct {
	// The unique sid that identifies this account
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The date this resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// A human readable description of this resource
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	// A string that uniquely identifies this resource
	Sid NullableString `json:"sid,omitempty"`
	// The IP addresses associated with this resource.
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI for this resource
	Uri NullableString `json:"uri,omitempty"`
}

// NewApiV2010AccountSipSipIpAccessControlList instantiates a new ApiV2010AccountSipSipIpAccessControlList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountSipSipIpAccessControlList() *ApiV2010AccountSipSipIpAccessControlList {
	this := ApiV2010AccountSipSipIpAccessControlList{}
	return &this
}

// NewApiV2010AccountSipSipIpAccessControlListWithDefaults instantiates a new ApiV2010AccountSipSipIpAccessControlList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountSipSipIpAccessControlListWithDefaults() *ApiV2010AccountSipSipIpAccessControlList {
	this := ApiV2010AccountSipSipIpAccessControlList{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetSid() {
	o.Sid.Unset()
}

// GetSubresourceUris returns the SubresourceUris field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetSubresourceUris() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SubresourceUris
}

// GetSubresourceUrisOk returns a tuple with the SubresourceUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetSubresourceUrisOk() (map[string]interface{}, bool) {
	if o == nil || o.SubresourceUris == nil {
		return nil, false
	}
	return o.SubresourceUris, true
}

// HasSubresourceUris returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasSubresourceUris() bool {
	if o != nil && o.SubresourceUris != nil {
		return true
	}

	return false
}

// SetSubresourceUris gets a reference to the given map[string]interface{} and assigns it to the SubresourceUris field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetSubresourceUris(v map[string]interface{}) {
	o.SubresourceUris = v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipIpAccessControlList) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipIpAccessControlList) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipIpAccessControlList) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountSipSipIpAccessControlList) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountSipSipIpAccessControlList) UnsetUri() {
	o.Uri.Unset()
}

func (o ApiV2010AccountSipSipIpAccessControlList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["date_created"] = o.DateCreated.Get()
	}
	if o.DateUpdated.IsSet() {
		toSerialize["date_updated"] = o.DateUpdated.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.SubresourceUris != nil {
		toSerialize["subresource_uris"] = o.SubresourceUris
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountSipSipIpAccessControlList struct {
	value *ApiV2010AccountSipSipIpAccessControlList
	isSet bool
}

func (v NullableApiV2010AccountSipSipIpAccessControlList) Get() *ApiV2010AccountSipSipIpAccessControlList {
	return v.value
}

func (v *NullableApiV2010AccountSipSipIpAccessControlList) Set(val *ApiV2010AccountSipSipIpAccessControlList) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountSipSipIpAccessControlList) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountSipSipIpAccessControlList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountSipSipIpAccessControlList(val *ApiV2010AccountSipSipIpAccessControlList) *NullableApiV2010AccountSipSipIpAccessControlList {
	return &NullableApiV2010AccountSipSipIpAccessControlList{value: val, isSet: true}
}

func (v NullableApiV2010AccountSipSipIpAccessControlList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountSipSipIpAccessControlList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


