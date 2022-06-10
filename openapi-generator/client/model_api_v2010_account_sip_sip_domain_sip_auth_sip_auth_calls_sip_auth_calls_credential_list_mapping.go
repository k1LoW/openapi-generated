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

// ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping struct for ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
type ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping struct {
	// The SID of the Account that created the resource
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	// The unique string that identifies the resource
	Sid NullableString `json:"sid,omitempty"`
}

// NewApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping instantiates a new ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping() *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping {
	this := ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping{}
	return &this
}

// NewApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMappingWithDefaults instantiates a new ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMappingWithDefaults() *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping {
	this := ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnsetSid() {
	o.Sid.Unset()
}

func (o ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping struct {
	value *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
	isSet bool
}

func (v NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) Get() *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping {
	return v.value
}

func (v *NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) Set(val *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping(val *ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) *NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping {
	return &NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping{value: val, isSet: true}
}

func (v NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


