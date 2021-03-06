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

// ApiV2010AccountAddress struct for ApiV2010AccountAddress
type ApiV2010AccountAddress struct {
	// The SID of the Account that is responsible for the resource
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The city in which the address is located
	City NullableString `json:"city,omitempty"`
	// The name associated with the address
	CustomerName NullableString `json:"customer_name,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// Whether emergency calling has been enabled on this number
	EmergencyEnabled NullableBool `json:"emergency_enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	// The ISO country code of the address
	IsoCountry NullableString `json:"iso_country,omitempty"`
	// The postal code of the address
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The state or region of the address
	Region NullableString `json:"region,omitempty"`
	// The unique string that identifies the resource
	Sid NullableString `json:"sid,omitempty"`
	// The number and street address of the address
	Street NullableString `json:"street,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri NullableString `json:"uri,omitempty"`
	// Whether the address has been validated to comply with local regulation
	Validated NullableBool `json:"validated,omitempty"`
	// Whether the address has been verified to comply with regulation
	Verified NullableBool `json:"verified,omitempty"`
}

// NewApiV2010AccountAddress instantiates a new ApiV2010AccountAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountAddress() *ApiV2010AccountAddress {
	this := ApiV2010AccountAddress{}
	return &this
}

// NewApiV2010AccountAddressWithDefaults instantiates a new ApiV2010AccountAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountAddressWithDefaults() *ApiV2010AccountAddress {
	this := ApiV2010AccountAddress{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountAddress) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountAddress) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ApiV2010AccountAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ApiV2010AccountAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetCity() {
	o.City.Unset()
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetCustomerName() string {
	if o == nil || o.CustomerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerName.Get()
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerName.Get(), o.CustomerName.IsSet()
}

// HasCustomerName returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasCustomerName() bool {
	if o != nil && o.CustomerName.IsSet() {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given NullableString and assigns it to the CustomerName field.
func (o *ApiV2010AccountAddress) SetCustomerName(v string) {
	o.CustomerName.Set(&v)
}
// SetCustomerNameNil sets the value for CustomerName to be an explicit nil
func (o *ApiV2010AccountAddress) SetCustomerNameNil() {
	o.CustomerName.Set(nil)
}

// UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetCustomerName() {
	o.CustomerName.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountAddress) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountAddress) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountAddress) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountAddress) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetEmergencyEnabled returns the EmergencyEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetEmergencyEnabled() bool {
	if o == nil || o.EmergencyEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EmergencyEnabled.Get()
}

// GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetEmergencyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmergencyEnabled.Get(), o.EmergencyEnabled.IsSet()
}

// HasEmergencyEnabled returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasEmergencyEnabled() bool {
	if o != nil && o.EmergencyEnabled.IsSet() {
		return true
	}

	return false
}

// SetEmergencyEnabled gets a reference to the given NullableBool and assigns it to the EmergencyEnabled field.
func (o *ApiV2010AccountAddress) SetEmergencyEnabled(v bool) {
	o.EmergencyEnabled.Set(&v)
}
// SetEmergencyEnabledNil sets the value for EmergencyEnabled to be an explicit nil
func (o *ApiV2010AccountAddress) SetEmergencyEnabledNil() {
	o.EmergencyEnabled.Set(nil)
}

// UnsetEmergencyEnabled ensures that no value is present for EmergencyEnabled, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetEmergencyEnabled() {
	o.EmergencyEnabled.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *ApiV2010AccountAddress) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *ApiV2010AccountAddress) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetIsoCountry returns the IsoCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetIsoCountry() string {
	if o == nil || o.IsoCountry.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCountry.Get()
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetIsoCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCountry.Get(), o.IsoCountry.IsSet()
}

// HasIsoCountry returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasIsoCountry() bool {
	if o != nil && o.IsoCountry.IsSet() {
		return true
	}

	return false
}

// SetIsoCountry gets a reference to the given NullableString and assigns it to the IsoCountry field.
func (o *ApiV2010AccountAddress) SetIsoCountry(v string) {
	o.IsoCountry.Set(&v)
}
// SetIsoCountryNil sets the value for IsoCountry to be an explicit nil
func (o *ApiV2010AccountAddress) SetIsoCountryNil() {
	o.IsoCountry.Set(nil)
}

// UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetIsoCountry() {
	o.IsoCountry.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *ApiV2010AccountAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *ApiV2010AccountAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *ApiV2010AccountAddress) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *ApiV2010AccountAddress) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetRegion() {
	o.Region.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountAddress) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountAddress) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetSid() {
	o.Sid.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *ApiV2010AccountAddress) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *ApiV2010AccountAddress) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetStreet() {
	o.Street.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountAddress) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountAddress) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetUri() {
	o.Uri.Unset()
}

// GetValidated returns the Validated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetValidated() bool {
	if o == nil || o.Validated.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Validated.Get()
}

// GetValidatedOk returns a tuple with the Validated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetValidatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Validated.Get(), o.Validated.IsSet()
}

// HasValidated returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasValidated() bool {
	if o != nil && o.Validated.IsSet() {
		return true
	}

	return false
}

// SetValidated gets a reference to the given NullableBool and assigns it to the Validated field.
func (o *ApiV2010AccountAddress) SetValidated(v bool) {
	o.Validated.Set(&v)
}
// SetValidatedNil sets the value for Validated to be an explicit nil
func (o *ApiV2010AccountAddress) SetValidatedNil() {
	o.Validated.Set(nil)
}

// UnsetValidated ensures that no value is present for Validated, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetValidated() {
	o.Validated.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAddress) GetVerified() bool {
	if o == nil || o.Verified.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAddress) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *ApiV2010AccountAddress) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *ApiV2010AccountAddress) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *ApiV2010AccountAddress) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *ApiV2010AccountAddress) UnsetVerified() {
	o.Verified.Unset()
}

func (o ApiV2010AccountAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.CustomerName.IsSet() {
		toSerialize["customer_name"] = o.CustomerName.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["date_created"] = o.DateCreated.Get()
	}
	if o.DateUpdated.IsSet() {
		toSerialize["date_updated"] = o.DateUpdated.Get()
	}
	if o.EmergencyEnabled.IsSet() {
		toSerialize["emergency_enabled"] = o.EmergencyEnabled.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if o.IsoCountry.IsSet() {
		toSerialize["iso_country"] = o.IsoCountry.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if o.Validated.IsSet() {
		toSerialize["validated"] = o.Validated.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountAddress struct {
	value *ApiV2010AccountAddress
	isSet bool
}

func (v NullableApiV2010AccountAddress) Get() *ApiV2010AccountAddress {
	return v.value
}

func (v *NullableApiV2010AccountAddress) Set(val *ApiV2010AccountAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountAddress(val *ApiV2010AccountAddress) *NullableApiV2010AccountAddress {
	return &NullableApiV2010AccountAddress{value: val, isSet: true}
}

func (v NullableApiV2010AccountAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


