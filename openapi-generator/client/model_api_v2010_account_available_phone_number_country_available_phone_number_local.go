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

// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct for ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct {
	// The type of Address resource the phone number requires
	AddressRequirements NullableString `json:"address_requirements,omitempty"`
	// Whether the phone number is new to the Twilio platform
	Beta NullableBool `json:"beta,omitempty"`
	Capabilities NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities `json:"capabilities,omitempty"`
	// A formatted version of the phone number
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	// The ISO country code of this phone number
	IsoCountry NullableString `json:"iso_country,omitempty"`
	// The LATA of this phone number
	Lata NullableString `json:"lata,omitempty"`
	// The latitude of this phone number's location
	Latitude NullableString `json:"latitude,omitempty"`
	// The locality or city of this phone number's location
	Locality NullableString `json:"locality,omitempty"`
	// The longitude of this phone number's location
	Longitude NullableString `json:"longitude,omitempty"`
	// The phone number in E.164 format
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The postal or ZIP code of this phone number's location
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The rate center of this phone number
	RateCenter NullableString `json:"rate_center,omitempty"`
	// The two-letter state or province abbreviation of this phone number's location
	Region NullableString `json:"region,omitempty"`
}

// NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal instantiates a new ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal() *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal {
	this := ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal{}
	return &this
}

// NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalWithDefaults instantiates a new ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalWithDefaults() *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal {
	this := ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal{}
	return &this
}

// GetAddressRequirements returns the AddressRequirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetAddressRequirements() string {
	if o == nil || o.AddressRequirements.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressRequirements.Get()
}

// GetAddressRequirementsOk returns a tuple with the AddressRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetAddressRequirementsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressRequirements.Get(), o.AddressRequirements.IsSet()
}

// HasAddressRequirements returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasAddressRequirements() bool {
	if o != nil && o.AddressRequirements.IsSet() {
		return true
	}

	return false
}

// SetAddressRequirements gets a reference to the given NullableString and assigns it to the AddressRequirements field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetAddressRequirements(v string) {
	o.AddressRequirements.Set(&v)
}
// SetAddressRequirementsNil sets the value for AddressRequirements to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetAddressRequirementsNil() {
	o.AddressRequirements.Set(nil)
}

// UnsetAddressRequirements ensures that no value is present for AddressRequirements, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetAddressRequirements() {
	o.AddressRequirements.Unset()
}

// GetBeta returns the Beta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetBeta() bool {
	if o == nil || o.Beta.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Beta.Get()
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetBetaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Beta.Get(), o.Beta.IsSet()
}

// HasBeta returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasBeta() bool {
	if o != nil && o.Beta.IsSet() {
		return true
	}

	return false
}

// SetBeta gets a reference to the given NullableBool and assigns it to the Beta field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetBeta(v bool) {
	o.Beta.Set(&v)
}
// SetBetaNil sets the value for Beta to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetBetaNil() {
	o.Beta.Set(nil)
}

// UnsetBeta ensures that no value is present for Beta, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetBeta() {
	o.Beta.Unset()
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetCapabilities() ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities {
	if o == nil || o.Capabilities.Get() == nil {
		var ret ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities
		return ret
	}
	return *o.Capabilities.Get()
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetCapabilitiesOk() (*ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities.Get(), o.Capabilities.IsSet()
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasCapabilities() bool {
	if o != nil && o.Capabilities.IsSet() {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities and assigns it to the Capabilities field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetCapabilities(v ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities) {
	o.Capabilities.Set(&v)
}
// SetCapabilitiesNil sets the value for Capabilities to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetCapabilitiesNil() {
	o.Capabilities.Set(nil)
}

// UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetCapabilities() {
	o.Capabilities.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetIsoCountry returns the IsoCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetIsoCountry() string {
	if o == nil || o.IsoCountry.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCountry.Get()
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetIsoCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCountry.Get(), o.IsoCountry.IsSet()
}

// HasIsoCountry returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasIsoCountry() bool {
	if o != nil && o.IsoCountry.IsSet() {
		return true
	}

	return false
}

// SetIsoCountry gets a reference to the given NullableString and assigns it to the IsoCountry field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetIsoCountry(v string) {
	o.IsoCountry.Set(&v)
}
// SetIsoCountryNil sets the value for IsoCountry to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetIsoCountryNil() {
	o.IsoCountry.Set(nil)
}

// UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetIsoCountry() {
	o.IsoCountry.Unset()
}

// GetLata returns the Lata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLata() string {
	if o == nil || o.Lata.Get() == nil {
		var ret string
		return ret
	}
	return *o.Lata.Get()
}

// GetLataOk returns a tuple with the Lata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lata.Get(), o.Lata.IsSet()
}

// HasLata returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasLata() bool {
	if o != nil && o.Lata.IsSet() {
		return true
	}

	return false
}

// SetLata gets a reference to the given NullableString and assigns it to the Lata field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLata(v string) {
	o.Lata.Set(&v)
}
// SetLataNil sets the value for Lata to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLataNil() {
	o.Lata.Set(nil)
}

// UnsetLata ensures that no value is present for Lata, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetLata() {
	o.Lata.Unset()
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLatitude() string {
	if o == nil || o.Latitude.Get() == nil {
		var ret string
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLatitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableString and assigns it to the Latitude field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLatitude(v string) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLocality returns the Locality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLocality() string {
	if o == nil || o.Locality.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locality.Get()
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLocalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locality.Get(), o.Locality.IsSet()
}

// HasLocality returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasLocality() bool {
	if o != nil && o.Locality.IsSet() {
		return true
	}

	return false
}

// SetLocality gets a reference to the given NullableString and assigns it to the Locality field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLocality(v string) {
	o.Locality.Set(&v)
}
// SetLocalityNil sets the value for Locality to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLocalityNil() {
	o.Locality.Set(nil)
}

// UnsetLocality ensures that no value is present for Locality, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetLocality() {
	o.Locality.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLongitude() string {
	if o == nil || o.Longitude.Get() == nil {
		var ret string
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetLongitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableString and assigns it to the Longitude field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLongitude(v string) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetRateCenter returns the RateCenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetRateCenter() string {
	if o == nil || o.RateCenter.Get() == nil {
		var ret string
		return ret
	}
	return *o.RateCenter.Get()
}

// GetRateCenterOk returns a tuple with the RateCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetRateCenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateCenter.Get(), o.RateCenter.IsSet()
}

// HasRateCenter returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasRateCenter() bool {
	if o != nil && o.RateCenter.IsSet() {
		return true
	}

	return false
}

// SetRateCenter gets a reference to the given NullableString and assigns it to the RateCenter field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetRateCenter(v string) {
	o.RateCenter.Set(&v)
}
// SetRateCenterNil sets the value for RateCenter to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetRateCenterNil() {
	o.RateCenter.Set(nil)
}

// UnsetRateCenter ensures that no value is present for RateCenter, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetRateCenter() {
	o.RateCenter.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnsetRegion() {
	o.Region.Unset()
}

func (o ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressRequirements.IsSet() {
		toSerialize["address_requirements"] = o.AddressRequirements.Get()
	}
	if o.Beta.IsSet() {
		toSerialize["beta"] = o.Beta.Get()
	}
	if o.Capabilities.IsSet() {
		toSerialize["capabilities"] = o.Capabilities.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if o.IsoCountry.IsSet() {
		toSerialize["iso_country"] = o.IsoCountry.Get()
	}
	if o.Lata.IsSet() {
		toSerialize["lata"] = o.Lata.Get()
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Locality.IsSet() {
		toSerialize["locality"] = o.Locality.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.RateCenter.IsSet() {
		toSerialize["rate_center"] = o.RateCenter.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct {
	value *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal
	isSet bool
}

func (v NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) Get() *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal {
	return v.value
}

func (v *NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) Set(val *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal(val *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) *NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal {
	return &NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal{value: val, isSet: true}
}

func (v NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


