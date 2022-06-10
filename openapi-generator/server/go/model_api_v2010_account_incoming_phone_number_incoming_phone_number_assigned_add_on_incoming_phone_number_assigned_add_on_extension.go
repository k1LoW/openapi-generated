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

type ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension struct {

	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`

	// The SID that uniquely identifies the assigned Add-on installation
	AssignedAddOnSid *string `json:"assigned_add_on_sid,omitempty"`

	// Whether the Extension will be invoked
	Enabled *bool `json:"enabled,omitempty"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`

	// A string that you assigned to describe the Product this Extension is used within
	ProductName *string `json:"product_name,omitempty"`

	// The SID of the Phone Number to which the Add-on is assigned
	ResourceSid *string `json:"resource_sid,omitempty"`

	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`

	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}

// AssertApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtensionRequired checks if the required fields are not zero-ed
func AssertApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtensionRequired(obj ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension) error {
	return nil
}

// AssertRecurseApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtensionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension (e.g. [][]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtensionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension, ok := obj.(ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtensionRequired(aApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension)
	})
}
