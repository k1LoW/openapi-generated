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

// ApiV2010AccountQueueMember struct for ApiV2010AccountQueueMember
type ApiV2010AccountQueueMember struct {
	// The SID of the Call the resource is associated with
	CallSid NullableString `json:"call_sid,omitempty"`
	// The date the member was enqueued
	DateEnqueued NullableString `json:"date_enqueued,omitempty"`
	// This member's current position in the queue.
	Position NullableInt32 `json:"position,omitempty"`
	// The SID of the Queue the member is in
	QueueSid NullableString `json:"queue_sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri NullableString `json:"uri,omitempty"`
	// The number of seconds the member has been in the queue.
	WaitTime NullableInt32 `json:"wait_time,omitempty"`
}

// NewApiV2010AccountQueueMember instantiates a new ApiV2010AccountQueueMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountQueueMember() *ApiV2010AccountQueueMember {
	this := ApiV2010AccountQueueMember{}
	return &this
}

// NewApiV2010AccountQueueMemberWithDefaults instantiates a new ApiV2010AccountQueueMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountQueueMemberWithDefaults() *ApiV2010AccountQueueMember {
	this := ApiV2010AccountQueueMember{}
	return &this
}

// GetCallSid returns the CallSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetCallSid() string {
	if o == nil || o.CallSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.CallSid.Get()
}

// GetCallSidOk returns a tuple with the CallSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetCallSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallSid.Get(), o.CallSid.IsSet()
}

// HasCallSid returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasCallSid() bool {
	if o != nil && o.CallSid.IsSet() {
		return true
	}

	return false
}

// SetCallSid gets a reference to the given NullableString and assigns it to the CallSid field.
func (o *ApiV2010AccountQueueMember) SetCallSid(v string) {
	o.CallSid.Set(&v)
}
// SetCallSidNil sets the value for CallSid to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetCallSidNil() {
	o.CallSid.Set(nil)
}

// UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetCallSid() {
	o.CallSid.Unset()
}

// GetDateEnqueued returns the DateEnqueued field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetDateEnqueued() string {
	if o == nil || o.DateEnqueued.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateEnqueued.Get()
}

// GetDateEnqueuedOk returns a tuple with the DateEnqueued field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetDateEnqueuedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateEnqueued.Get(), o.DateEnqueued.IsSet()
}

// HasDateEnqueued returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasDateEnqueued() bool {
	if o != nil && o.DateEnqueued.IsSet() {
		return true
	}

	return false
}

// SetDateEnqueued gets a reference to the given NullableString and assigns it to the DateEnqueued field.
func (o *ApiV2010AccountQueueMember) SetDateEnqueued(v string) {
	o.DateEnqueued.Set(&v)
}
// SetDateEnqueuedNil sets the value for DateEnqueued to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetDateEnqueuedNil() {
	o.DateEnqueued.Set(nil)
}

// UnsetDateEnqueued ensures that no value is present for DateEnqueued, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetDateEnqueued() {
	o.DateEnqueued.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetPosition() int32 {
	if o == nil || o.Position.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableInt32 and assigns it to the Position field.
func (o *ApiV2010AccountQueueMember) SetPosition(v int32) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetPosition() {
	o.Position.Unset()
}

// GetQueueSid returns the QueueSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetQueueSid() string {
	if o == nil || o.QueueSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueueSid.Get()
}

// GetQueueSidOk returns a tuple with the QueueSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetQueueSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueueSid.Get(), o.QueueSid.IsSet()
}

// HasQueueSid returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasQueueSid() bool {
	if o != nil && o.QueueSid.IsSet() {
		return true
	}

	return false
}

// SetQueueSid gets a reference to the given NullableString and assigns it to the QueueSid field.
func (o *ApiV2010AccountQueueMember) SetQueueSid(v string) {
	o.QueueSid.Set(&v)
}
// SetQueueSidNil sets the value for QueueSid to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetQueueSidNil() {
	o.QueueSid.Set(nil)
}

// UnsetQueueSid ensures that no value is present for QueueSid, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetQueueSid() {
	o.QueueSid.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountQueueMember) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetUri() {
	o.Uri.Unset()
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountQueueMember) GetWaitTime() int32 {
	if o == nil || o.WaitTime.Get() == nil {
		var ret int32
		return ret
	}
	return *o.WaitTime.Get()
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountQueueMember) GetWaitTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WaitTime.Get(), o.WaitTime.IsSet()
}

// HasWaitTime returns a boolean if a field has been set.
func (o *ApiV2010AccountQueueMember) HasWaitTime() bool {
	if o != nil && o.WaitTime.IsSet() {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given NullableInt32 and assigns it to the WaitTime field.
func (o *ApiV2010AccountQueueMember) SetWaitTime(v int32) {
	o.WaitTime.Set(&v)
}
// SetWaitTimeNil sets the value for WaitTime to be an explicit nil
func (o *ApiV2010AccountQueueMember) SetWaitTimeNil() {
	o.WaitTime.Set(nil)
}

// UnsetWaitTime ensures that no value is present for WaitTime, not even an explicit nil
func (o *ApiV2010AccountQueueMember) UnsetWaitTime() {
	o.WaitTime.Unset()
}

func (o ApiV2010AccountQueueMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallSid.IsSet() {
		toSerialize["call_sid"] = o.CallSid.Get()
	}
	if o.DateEnqueued.IsSet() {
		toSerialize["date_enqueued"] = o.DateEnqueued.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.QueueSid.IsSet() {
		toSerialize["queue_sid"] = o.QueueSid.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if o.WaitTime.IsSet() {
		toSerialize["wait_time"] = o.WaitTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountQueueMember struct {
	value *ApiV2010AccountQueueMember
	isSet bool
}

func (v NullableApiV2010AccountQueueMember) Get() *ApiV2010AccountQueueMember {
	return v.value
}

func (v *NullableApiV2010AccountQueueMember) Set(val *ApiV2010AccountQueueMember) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountQueueMember) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountQueueMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountQueueMember(val *ApiV2010AccountQueueMember) *NullableApiV2010AccountQueueMember {
	return &NullableApiV2010AccountQueueMember{value: val, isSet: true}
}

func (v NullableApiV2010AccountQueueMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountQueueMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

