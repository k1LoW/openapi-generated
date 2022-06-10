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

// ApiV2010AccountSipSipCredentialListSipCredential struct for ApiV2010AccountSipSipCredentialListSipCredential
type ApiV2010AccountSipSipCredentialListSipCredential struct {
	// The unique id of the Account that is responsible for this resource.
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The unique id that identifies the credential list that includes this credential
	CredentialListSid NullableString `json:"credential_list_sid,omitempty"`
	// The date that this resource was created, given as GMT in RFC 2822 format.
	DateCreated NullableString `json:"date_created,omitempty"`
	// The date that this resource was last updated, given as GMT in RFC 2822 format.
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid NullableString `json:"sid,omitempty"`
	// The URI for this resource, relative to https://api.twilio.com
	Uri NullableString `json:"uri,omitempty"`
	// The username for this credential.
	Username NullableString `json:"username,omitempty"`
}

// NewApiV2010AccountSipSipCredentialListSipCredential instantiates a new ApiV2010AccountSipSipCredentialListSipCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountSipSipCredentialListSipCredential() *ApiV2010AccountSipSipCredentialListSipCredential {
	this := ApiV2010AccountSipSipCredentialListSipCredential{}
	return &this
}

// NewApiV2010AccountSipSipCredentialListSipCredentialWithDefaults instantiates a new ApiV2010AccountSipSipCredentialListSipCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountSipSipCredentialListSipCredentialWithDefaults() *ApiV2010AccountSipSipCredentialListSipCredential {
	this := ApiV2010AccountSipSipCredentialListSipCredential{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetCredentialListSid returns the CredentialListSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetCredentialListSid() string {
	if o == nil || o.CredentialListSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.CredentialListSid.Get()
}

// GetCredentialListSidOk returns a tuple with the CredentialListSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetCredentialListSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialListSid.Get(), o.CredentialListSid.IsSet()
}

// HasCredentialListSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasCredentialListSid() bool {
	if o != nil && o.CredentialListSid.IsSet() {
		return true
	}

	return false
}

// SetCredentialListSid gets a reference to the given NullableString and assigns it to the CredentialListSid field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetCredentialListSid(v string) {
	o.CredentialListSid.Set(&v)
}
// SetCredentialListSidNil sets the value for CredentialListSid to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetCredentialListSidNil() {
	o.CredentialListSid.Set(nil)
}

// UnsetCredentialListSid ensures that no value is present for CredentialListSid, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetCredentialListSid() {
	o.CredentialListSid.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetSid() {
	o.Sid.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetUri() {
	o.Uri.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetUsername() {
	o.Username.Unset()
}

func (o ApiV2010AccountSipSipCredentialListSipCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.CredentialListSid.IsSet() {
		toSerialize["credential_list_sid"] = o.CredentialListSid.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["date_created"] = o.DateCreated.Get()
	}
	if o.DateUpdated.IsSet() {
		toSerialize["date_updated"] = o.DateUpdated.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountSipSipCredentialListSipCredential struct {
	value *ApiV2010AccountSipSipCredentialListSipCredential
	isSet bool
}

func (v NullableApiV2010AccountSipSipCredentialListSipCredential) Get() *ApiV2010AccountSipSipCredentialListSipCredential {
	return v.value
}

func (v *NullableApiV2010AccountSipSipCredentialListSipCredential) Set(val *ApiV2010AccountSipSipCredentialListSipCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountSipSipCredentialListSipCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountSipSipCredentialListSipCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountSipSipCredentialListSipCredential(val *ApiV2010AccountSipSipCredentialListSipCredential) *NullableApiV2010AccountSipSipCredentialListSipCredential {
	return &NullableApiV2010AccountSipSipCredentialListSipCredential{value: val, isSet: true}
}

func (v NullableApiV2010AccountSipSipCredentialListSipCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountSipSipCredentialListSipCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

