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

// ApiV2010AccountShortCode struct for ApiV2010AccountShortCode
type ApiV2010AccountShortCode struct {
	// The SID of the Account that created this resource
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The API version used to start a new TwiML session
	ApiVersion NullableString `json:"api_version,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// A string that you assigned to describe this resource
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	// The short code. e.g., 894546.
	ShortCode NullableString `json:"short_code,omitempty"`
	// The unique string that identifies this resource
	Sid NullableString `json:"sid,omitempty"`
	// HTTP method we use to call the sms_fallback_url
	SmsFallbackMethod NullableString `json:"sms_fallback_method,omitempty"`
	// URL Twilio will request if an error occurs in executing TwiML
	SmsFallbackUrl NullableString `json:"sms_fallback_url,omitempty"`
	// HTTP method to use when requesting the sms url
	SmsMethod NullableString `json:"sms_method,omitempty"`
	// URL we call when receiving an incoming SMS message to this short code
	SmsUrl NullableString `json:"sms_url,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri NullableString `json:"uri,omitempty"`
}

// NewApiV2010AccountShortCode instantiates a new ApiV2010AccountShortCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountShortCode() *ApiV2010AccountShortCode {
	this := ApiV2010AccountShortCode{}
	return &this
}

// NewApiV2010AccountShortCodeWithDefaults instantiates a new ApiV2010AccountShortCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountShortCodeWithDefaults() *ApiV2010AccountShortCode {
	this := ApiV2010AccountShortCode{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountShortCode) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountShortCode) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetApiVersion() string {
	if o == nil || o.ApiVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion.Get()
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiVersion.Get(), o.ApiVersion.IsSet()
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasApiVersion() bool {
	if o != nil && o.ApiVersion.IsSet() {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given NullableString and assigns it to the ApiVersion field.
func (o *ApiV2010AccountShortCode) SetApiVersion(v string) {
	o.ApiVersion.Set(&v)
}
// SetApiVersionNil sets the value for ApiVersion to be an explicit nil
func (o *ApiV2010AccountShortCode) SetApiVersionNil() {
	o.ApiVersion.Set(nil)
}

// UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetApiVersion() {
	o.ApiVersion.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountShortCode) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountShortCode) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountShortCode) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountShortCode) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *ApiV2010AccountShortCode) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *ApiV2010AccountShortCode) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetShortCode returns the ShortCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetShortCode() string {
	if o == nil || o.ShortCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortCode.Get()
}

// GetShortCodeOk returns a tuple with the ShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetShortCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortCode.Get(), o.ShortCode.IsSet()
}

// HasShortCode returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasShortCode() bool {
	if o != nil && o.ShortCode.IsSet() {
		return true
	}

	return false
}

// SetShortCode gets a reference to the given NullableString and assigns it to the ShortCode field.
func (o *ApiV2010AccountShortCode) SetShortCode(v string) {
	o.ShortCode.Set(&v)
}
// SetShortCodeNil sets the value for ShortCode to be an explicit nil
func (o *ApiV2010AccountShortCode) SetShortCodeNil() {
	o.ShortCode.Set(nil)
}

// UnsetShortCode ensures that no value is present for ShortCode, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetShortCode() {
	o.ShortCode.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountShortCode) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountShortCode) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetSid() {
	o.Sid.Unset()
}

// GetSmsFallbackMethod returns the SmsFallbackMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetSmsFallbackMethod() string {
	if o == nil || o.SmsFallbackMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.SmsFallbackMethod.Get()
}

// GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetSmsFallbackMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsFallbackMethod.Get(), o.SmsFallbackMethod.IsSet()
}

// HasSmsFallbackMethod returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasSmsFallbackMethod() bool {
	if o != nil && o.SmsFallbackMethod.IsSet() {
		return true
	}

	return false
}

// SetSmsFallbackMethod gets a reference to the given NullableString and assigns it to the SmsFallbackMethod field.
func (o *ApiV2010AccountShortCode) SetSmsFallbackMethod(v string) {
	o.SmsFallbackMethod.Set(&v)
}
// SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil
func (o *ApiV2010AccountShortCode) SetSmsFallbackMethodNil() {
	o.SmsFallbackMethod.Set(nil)
}

// UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetSmsFallbackMethod() {
	o.SmsFallbackMethod.Unset()
}

// GetSmsFallbackUrl returns the SmsFallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetSmsFallbackUrl() string {
	if o == nil || o.SmsFallbackUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SmsFallbackUrl.Get()
}

// GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetSmsFallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsFallbackUrl.Get(), o.SmsFallbackUrl.IsSet()
}

// HasSmsFallbackUrl returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasSmsFallbackUrl() bool {
	if o != nil && o.SmsFallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetSmsFallbackUrl gets a reference to the given NullableString and assigns it to the SmsFallbackUrl field.
func (o *ApiV2010AccountShortCode) SetSmsFallbackUrl(v string) {
	o.SmsFallbackUrl.Set(&v)
}
// SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil
func (o *ApiV2010AccountShortCode) SetSmsFallbackUrlNil() {
	o.SmsFallbackUrl.Set(nil)
}

// UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetSmsFallbackUrl() {
	o.SmsFallbackUrl.Unset()
}

// GetSmsMethod returns the SmsMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetSmsMethod() string {
	if o == nil || o.SmsMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.SmsMethod.Get()
}

// GetSmsMethodOk returns a tuple with the SmsMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetSmsMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsMethod.Get(), o.SmsMethod.IsSet()
}

// HasSmsMethod returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasSmsMethod() bool {
	if o != nil && o.SmsMethod.IsSet() {
		return true
	}

	return false
}

// SetSmsMethod gets a reference to the given NullableString and assigns it to the SmsMethod field.
func (o *ApiV2010AccountShortCode) SetSmsMethod(v string) {
	o.SmsMethod.Set(&v)
}
// SetSmsMethodNil sets the value for SmsMethod to be an explicit nil
func (o *ApiV2010AccountShortCode) SetSmsMethodNil() {
	o.SmsMethod.Set(nil)
}

// UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetSmsMethod() {
	o.SmsMethod.Unset()
}

// GetSmsUrl returns the SmsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetSmsUrl() string {
	if o == nil || o.SmsUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SmsUrl.Get()
}

// GetSmsUrlOk returns a tuple with the SmsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetSmsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsUrl.Get(), o.SmsUrl.IsSet()
}

// HasSmsUrl returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasSmsUrl() bool {
	if o != nil && o.SmsUrl.IsSet() {
		return true
	}

	return false
}

// SetSmsUrl gets a reference to the given NullableString and assigns it to the SmsUrl field.
func (o *ApiV2010AccountShortCode) SetSmsUrl(v string) {
	o.SmsUrl.Set(&v)
}
// SetSmsUrlNil sets the value for SmsUrl to be an explicit nil
func (o *ApiV2010AccountShortCode) SetSmsUrlNil() {
	o.SmsUrl.Set(nil)
}

// UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetSmsUrl() {
	o.SmsUrl.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountShortCode) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountShortCode) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountShortCode) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountShortCode) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountShortCode) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountShortCode) UnsetUri() {
	o.Uri.Unset()
}

func (o ApiV2010AccountShortCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.ApiVersion.IsSet() {
		toSerialize["api_version"] = o.ApiVersion.Get()
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
	if o.ShortCode.IsSet() {
		toSerialize["short_code"] = o.ShortCode.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.SmsFallbackMethod.IsSet() {
		toSerialize["sms_fallback_method"] = o.SmsFallbackMethod.Get()
	}
	if o.SmsFallbackUrl.IsSet() {
		toSerialize["sms_fallback_url"] = o.SmsFallbackUrl.Get()
	}
	if o.SmsMethod.IsSet() {
		toSerialize["sms_method"] = o.SmsMethod.Get()
	}
	if o.SmsUrl.IsSet() {
		toSerialize["sms_url"] = o.SmsUrl.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountShortCode struct {
	value *ApiV2010AccountShortCode
	isSet bool
}

func (v NullableApiV2010AccountShortCode) Get() *ApiV2010AccountShortCode {
	return v.value
}

func (v *NullableApiV2010AccountShortCode) Set(val *ApiV2010AccountShortCode) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountShortCode) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountShortCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountShortCode(val *ApiV2010AccountShortCode) *NullableApiV2010AccountShortCode {
	return &NullableApiV2010AccountShortCode{value: val, isSet: true}
}

func (v NullableApiV2010AccountShortCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountShortCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


