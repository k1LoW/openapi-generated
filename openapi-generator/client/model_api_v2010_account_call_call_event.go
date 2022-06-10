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

// ApiV2010AccountCallCallEvent struct for ApiV2010AccountCallCallEvent
type ApiV2010AccountCallCallEvent struct {
	// Call Request.
	Request interface{} `json:"request,omitempty"`
	// Call Response with Events.
	Response interface{} `json:"response,omitempty"`
}

// NewApiV2010AccountCallCallEvent instantiates a new ApiV2010AccountCallCallEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountCallCallEvent() *ApiV2010AccountCallCallEvent {
	this := ApiV2010AccountCallCallEvent{}
	return &this
}

// NewApiV2010AccountCallCallEventWithDefaults instantiates a new ApiV2010AccountCallCallEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountCallCallEventWithDefaults() *ApiV2010AccountCallCallEvent {
	this := ApiV2010AccountCallCallEvent{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallEvent) GetRequest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallEvent) GetRequestOk() (*interface{}, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return &o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallEvent) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given interface{} and assigns it to the Request field.
func (o *ApiV2010AccountCallCallEvent) SetRequest(v interface{}) {
	o.Request = v
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallEvent) GetResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallEvent) GetResponseOk() (*interface{}, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return &o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallEvent) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given interface{} and assigns it to the Response field.
func (o *ApiV2010AccountCallCallEvent) SetResponse(v interface{}) {
	o.Response = v
}

func (o ApiV2010AccountCallCallEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountCallCallEvent struct {
	value *ApiV2010AccountCallCallEvent
	isSet bool
}

func (v NullableApiV2010AccountCallCallEvent) Get() *ApiV2010AccountCallCallEvent {
	return v.value
}

func (v *NullableApiV2010AccountCallCallEvent) Set(val *ApiV2010AccountCallCallEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountCallCallEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountCallCallEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountCallCallEvent(val *ApiV2010AccountCallCallEvent) *NullableApiV2010AccountCallCallEvent {
	return &NullableApiV2010AccountCallCallEvent{value: val, isSet: true}
}

func (v NullableApiV2010AccountCallCallEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountCallCallEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

