/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct {

	// The type of Address resource the phone number requires
	AddressRequirements *string `json:"address_requirements,omitempty"`

	// Whether the phone number is new to the Twilio platform
	Beta *bool `json:"beta,omitempty"`

	Capabilities *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities `json:"capabilities,omitempty"`

	// A formatted version of the phone number
	FriendlyName *string `json:"friendly_name,omitempty"`

	// The ISO country code of this phone number
	IsoCountry *string `json:"iso_country,omitempty"`

	// The LATA of this phone number
	Lata *string `json:"lata,omitempty"`

	// The latitude of this phone number's location
	Latitude *string `json:"latitude,omitempty"`

	// The locality or city of this phone number's location
	Locality *string `json:"locality,omitempty"`

	// The longitude of this phone number's location
	Longitude *string `json:"longitude,omitempty"`

	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The postal or ZIP code of this phone number's location
	PostalCode *string `json:"postal_code,omitempty"`

	// The rate center of this phone number
	RateCenter *string `json:"rate_center,omitempty"`

	// The two-letter state or province abbreviation of this phone number's location
	Region *string `json:"region,omitempty"`
}

// AssertApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalRequired(obj ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal) error {
	if obj.Capabilities != nil {
		if err := AssertApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilitiesRequired(*obj.Capabilities); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal (e.g. [][]ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal, ok := obj.(ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalRequired(aApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal)
	})
}