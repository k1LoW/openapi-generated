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

// ApiV2010AccountIncomingPhoneNumberCapabilities Indicate if a phone can receive calls or messages
type ApiV2010AccountIncomingPhoneNumberCapabilities struct {
	Fax *bool `json:"fax,omitempty"`
	Mms *bool `json:"mms,omitempty"`
	Sms *bool `json:"sms,omitempty"`
	Voice *bool `json:"voice,omitempty"`
}

// NewApiV2010AccountIncomingPhoneNumberCapabilities instantiates a new ApiV2010AccountIncomingPhoneNumberCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountIncomingPhoneNumberCapabilities() *ApiV2010AccountIncomingPhoneNumberCapabilities {
	this := ApiV2010AccountIncomingPhoneNumberCapabilities{}
	return &this
}

// NewApiV2010AccountIncomingPhoneNumberCapabilitiesWithDefaults instantiates a new ApiV2010AccountIncomingPhoneNumberCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountIncomingPhoneNumberCapabilitiesWithDefaults() *ApiV2010AccountIncomingPhoneNumberCapabilities {
	this := ApiV2010AccountIncomingPhoneNumberCapabilities{}
	return &this
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetFax() bool {
	if o == nil || o.Fax == nil {
		var ret bool
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetFaxOk() (*bool, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given bool and assigns it to the Fax field.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) SetFax(v bool) {
	o.Fax = &v
}

// GetMms returns the Mms field value if set, zero value otherwise.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetMms() bool {
	if o == nil || o.Mms == nil {
		var ret bool
		return ret
	}
	return *o.Mms
}

// GetMmsOk returns a tuple with the Mms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetMmsOk() (*bool, bool) {
	if o == nil || o.Mms == nil {
		return nil, false
	}
	return o.Mms, true
}

// HasMms returns a boolean if a field has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) HasMms() bool {
	if o != nil && o.Mms != nil {
		return true
	}

	return false
}

// SetMms gets a reference to the given bool and assigns it to the Mms field.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) SetMms(v bool) {
	o.Mms = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetSms() bool {
	if o == nil || o.Sms == nil {
		var ret bool
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetSmsOk() (*bool, bool) {
	if o == nil || o.Sms == nil {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) HasSms() bool {
	if o != nil && o.Sms != nil {
		return true
	}

	return false
}

// SetSms gets a reference to the given bool and assigns it to the Sms field.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) SetSms(v bool) {
	o.Sms = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetVoice() bool {
	if o == nil || o.Voice == nil {
		var ret bool
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) GetVoiceOk() (*bool, bool) {
	if o == nil || o.Voice == nil {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) HasVoice() bool {
	if o != nil && o.Voice != nil {
		return true
	}

	return false
}

// SetVoice gets a reference to the given bool and assigns it to the Voice field.
func (o *ApiV2010AccountIncomingPhoneNumberCapabilities) SetVoice(v bool) {
	o.Voice = &v
}

func (o ApiV2010AccountIncomingPhoneNumberCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.Mms != nil {
		toSerialize["mms"] = o.Mms
	}
	if o.Sms != nil {
		toSerialize["sms"] = o.Sms
	}
	if o.Voice != nil {
		toSerialize["voice"] = o.Voice
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountIncomingPhoneNumberCapabilities struct {
	value *ApiV2010AccountIncomingPhoneNumberCapabilities
	isSet bool
}

func (v NullableApiV2010AccountIncomingPhoneNumberCapabilities) Get() *ApiV2010AccountIncomingPhoneNumberCapabilities {
	return v.value
}

func (v *NullableApiV2010AccountIncomingPhoneNumberCapabilities) Set(val *ApiV2010AccountIncomingPhoneNumberCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountIncomingPhoneNumberCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountIncomingPhoneNumberCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountIncomingPhoneNumberCapabilities(val *ApiV2010AccountIncomingPhoneNumberCapabilities) *NullableApiV2010AccountIncomingPhoneNumberCapabilities {
	return &NullableApiV2010AccountIncomingPhoneNumberCapabilities{value: val, isSet: true}
}

func (v NullableApiV2010AccountIncomingPhoneNumberCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountIncomingPhoneNumberCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


