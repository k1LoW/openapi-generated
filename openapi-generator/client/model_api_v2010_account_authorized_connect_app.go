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

// ApiV2010AccountAuthorizedConnectApp struct for ApiV2010AccountAuthorizedConnectApp
type ApiV2010AccountAuthorizedConnectApp struct {
	// The SID of the Account that created the resource
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The company name set for the Connect App
	ConnectAppCompanyName NullableString `json:"connect_app_company_name,omitempty"`
	// A detailed description of the app
	ConnectAppDescription NullableString `json:"connect_app_description,omitempty"`
	// The name of the Connect App
	ConnectAppFriendlyName NullableString `json:"connect_app_friendly_name,omitempty"`
	// The public URL for the Connect App
	ConnectAppHomepageUrl NullableString `json:"connect_app_homepage_url,omitempty"`
	// The SID that we assigned to the Connect App
	ConnectAppSid NullableString `json:"connect_app_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// Permissions authorized to the app
	Permissions []string `json:"permissions,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri NullableString `json:"uri,omitempty"`
}

// NewApiV2010AccountAuthorizedConnectApp instantiates a new ApiV2010AccountAuthorizedConnectApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountAuthorizedConnectApp() *ApiV2010AccountAuthorizedConnectApp {
	this := ApiV2010AccountAuthorizedConnectApp{}
	return &this
}

// NewApiV2010AccountAuthorizedConnectAppWithDefaults instantiates a new ApiV2010AccountAuthorizedConnectApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountAuthorizedConnectAppWithDefaults() *ApiV2010AccountAuthorizedConnectApp {
	this := ApiV2010AccountAuthorizedConnectApp{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetConnectAppCompanyName returns the ConnectAppCompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppCompanyName() string {
	if o == nil || o.ConnectAppCompanyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectAppCompanyName.Get()
}

// GetConnectAppCompanyNameOk returns a tuple with the ConnectAppCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectAppCompanyName.Get(), o.ConnectAppCompanyName.IsSet()
}

// HasConnectAppCompanyName returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppCompanyName() bool {
	if o != nil && o.ConnectAppCompanyName.IsSet() {
		return true
	}

	return false
}

// SetConnectAppCompanyName gets a reference to the given NullableString and assigns it to the ConnectAppCompanyName field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppCompanyName(v string) {
	o.ConnectAppCompanyName.Set(&v)
}
// SetConnectAppCompanyNameNil sets the value for ConnectAppCompanyName to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppCompanyNameNil() {
	o.ConnectAppCompanyName.Set(nil)
}

// UnsetConnectAppCompanyName ensures that no value is present for ConnectAppCompanyName, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppCompanyName() {
	o.ConnectAppCompanyName.Unset()
}

// GetConnectAppDescription returns the ConnectAppDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppDescription() string {
	if o == nil || o.ConnectAppDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectAppDescription.Get()
}

// GetConnectAppDescriptionOk returns a tuple with the ConnectAppDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectAppDescription.Get(), o.ConnectAppDescription.IsSet()
}

// HasConnectAppDescription returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppDescription() bool {
	if o != nil && o.ConnectAppDescription.IsSet() {
		return true
	}

	return false
}

// SetConnectAppDescription gets a reference to the given NullableString and assigns it to the ConnectAppDescription field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppDescription(v string) {
	o.ConnectAppDescription.Set(&v)
}
// SetConnectAppDescriptionNil sets the value for ConnectAppDescription to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppDescriptionNil() {
	o.ConnectAppDescription.Set(nil)
}

// UnsetConnectAppDescription ensures that no value is present for ConnectAppDescription, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppDescription() {
	o.ConnectAppDescription.Unset()
}

// GetConnectAppFriendlyName returns the ConnectAppFriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppFriendlyName() string {
	if o == nil || o.ConnectAppFriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectAppFriendlyName.Get()
}

// GetConnectAppFriendlyNameOk returns a tuple with the ConnectAppFriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectAppFriendlyName.Get(), o.ConnectAppFriendlyName.IsSet()
}

// HasConnectAppFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppFriendlyName() bool {
	if o != nil && o.ConnectAppFriendlyName.IsSet() {
		return true
	}

	return false
}

// SetConnectAppFriendlyName gets a reference to the given NullableString and assigns it to the ConnectAppFriendlyName field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppFriendlyName(v string) {
	o.ConnectAppFriendlyName.Set(&v)
}
// SetConnectAppFriendlyNameNil sets the value for ConnectAppFriendlyName to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppFriendlyNameNil() {
	o.ConnectAppFriendlyName.Set(nil)
}

// UnsetConnectAppFriendlyName ensures that no value is present for ConnectAppFriendlyName, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppFriendlyName() {
	o.ConnectAppFriendlyName.Unset()
}

// GetConnectAppHomepageUrl returns the ConnectAppHomepageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppHomepageUrl() string {
	if o == nil || o.ConnectAppHomepageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectAppHomepageUrl.Get()
}

// GetConnectAppHomepageUrlOk returns a tuple with the ConnectAppHomepageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppHomepageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectAppHomepageUrl.Get(), o.ConnectAppHomepageUrl.IsSet()
}

// HasConnectAppHomepageUrl returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppHomepageUrl() bool {
	if o != nil && o.ConnectAppHomepageUrl.IsSet() {
		return true
	}

	return false
}

// SetConnectAppHomepageUrl gets a reference to the given NullableString and assigns it to the ConnectAppHomepageUrl field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppHomepageUrl(v string) {
	o.ConnectAppHomepageUrl.Set(&v)
}
// SetConnectAppHomepageUrlNil sets the value for ConnectAppHomepageUrl to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppHomepageUrlNil() {
	o.ConnectAppHomepageUrl.Set(nil)
}

// UnsetConnectAppHomepageUrl ensures that no value is present for ConnectAppHomepageUrl, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppHomepageUrl() {
	o.ConnectAppHomepageUrl.Unset()
}

// GetConnectAppSid returns the ConnectAppSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppSid() string {
	if o == nil || o.ConnectAppSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectAppSid.Get()
}

// GetConnectAppSidOk returns a tuple with the ConnectAppSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectAppSid.Get(), o.ConnectAppSid.IsSet()
}

// HasConnectAppSid returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppSid() bool {
	if o != nil && o.ConnectAppSid.IsSet() {
		return true
	}

	return false
}

// SetConnectAppSid gets a reference to the given NullableString and assigns it to the ConnectAppSid field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppSid(v string) {
	o.ConnectAppSid.Set(&v)
}
// SetConnectAppSidNil sets the value for ConnectAppSid to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppSidNil() {
	o.ConnectAppSid.Set(nil)
}

// UnsetConnectAppSid ensures that no value is present for ConnectAppSid, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppSid() {
	o.ConnectAppSid.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetPermissionsOk() ([]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetPermissions(v []string) {
	o.Permissions = v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAuthorizedConnectApp) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAuthorizedConnectApp) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountAuthorizedConnectApp) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountAuthorizedConnectApp) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountAuthorizedConnectApp) UnsetUri() {
	o.Uri.Unset()
}

func (o ApiV2010AccountAuthorizedConnectApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.ConnectAppCompanyName.IsSet() {
		toSerialize["connect_app_company_name"] = o.ConnectAppCompanyName.Get()
	}
	if o.ConnectAppDescription.IsSet() {
		toSerialize["connect_app_description"] = o.ConnectAppDescription.Get()
	}
	if o.ConnectAppFriendlyName.IsSet() {
		toSerialize["connect_app_friendly_name"] = o.ConnectAppFriendlyName.Get()
	}
	if o.ConnectAppHomepageUrl.IsSet() {
		toSerialize["connect_app_homepage_url"] = o.ConnectAppHomepageUrl.Get()
	}
	if o.ConnectAppSid.IsSet() {
		toSerialize["connect_app_sid"] = o.ConnectAppSid.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["date_created"] = o.DateCreated.Get()
	}
	if o.DateUpdated.IsSet() {
		toSerialize["date_updated"] = o.DateUpdated.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountAuthorizedConnectApp struct {
	value *ApiV2010AccountAuthorizedConnectApp
	isSet bool
}

func (v NullableApiV2010AccountAuthorizedConnectApp) Get() *ApiV2010AccountAuthorizedConnectApp {
	return v.value
}

func (v *NullableApiV2010AccountAuthorizedConnectApp) Set(val *ApiV2010AccountAuthorizedConnectApp) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountAuthorizedConnectApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountAuthorizedConnectApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountAuthorizedConnectApp(val *ApiV2010AccountAuthorizedConnectApp) *NullableApiV2010AccountAuthorizedConnectApp {
	return &NullableApiV2010AccountAuthorizedConnectApp{value: val, isSet: true}
}

func (v NullableApiV2010AccountAuthorizedConnectApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountAuthorizedConnectApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

