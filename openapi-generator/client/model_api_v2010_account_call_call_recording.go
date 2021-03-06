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

// ApiV2010AccountCallCallRecording struct for ApiV2010AccountCallCallRecording
type ApiV2010AccountCallCallRecording struct {
	// The SID of the Account that created the resource
	AccountSid NullableString `json:"account_sid,omitempty"`
	// The API version used to make the recording
	ApiVersion NullableString `json:"api_version,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid NullableString `json:"call_sid,omitempty"`
	// The number of channels in the final recording file
	Channels NullableInt32 `json:"channels,omitempty"`
	// The Conference SID that identifies the conference associated with the recording
	ConferenceSid NullableString `json:"conference_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated NullableString `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated NullableString `json:"date_updated,omitempty"`
	// The length of the recording in seconds
	Duration NullableString `json:"duration,omitempty"`
	// How to decrypt the recording.
	EncryptionDetails interface{} `json:"encryption_details,omitempty"`
	// More information about why the recording is missing, if status is `absent`.
	ErrorCode NullableInt32 `json:"error_code,omitempty"`
	// The one-time cost of creating the recording.
	Price NullableFloat32 `json:"price,omitempty"`
	// The currency used in the price property.
	PriceUnit NullableString `json:"price_unit,omitempty"`
	// The unique string that identifies the resource
	Sid NullableString `json:"sid,omitempty"`
	// How the recording was created
	Source NullableString `json:"source,omitempty"`
	// The start time of the recording, given in RFC 2822 format
	StartTime NullableString `json:"start_time,omitempty"`
	// The status of the recording
	Status NullableString `json:"status,omitempty"`
	// The recorded track. Can be: `inbound`, `outbound`, or `both`.
	Track NullableString `json:"track,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri NullableString `json:"uri,omitempty"`
}

// NewApiV2010AccountCallCallRecording instantiates a new ApiV2010AccountCallCallRecording object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2010AccountCallCallRecording() *ApiV2010AccountCallCallRecording {
	this := ApiV2010AccountCallCallRecording{}
	return &this
}

// NewApiV2010AccountCallCallRecordingWithDefaults instantiates a new ApiV2010AccountCallCallRecording object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2010AccountCallCallRecordingWithDefaults() *ApiV2010AccountCallCallRecording {
	this := ApiV2010AccountCallCallRecording{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetAccountSid() string {
	if o == nil || o.AccountSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountSid.Get()
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSid.Get(), o.AccountSid.IsSet()
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasAccountSid() bool {
	if o != nil && o.AccountSid.IsSet() {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given NullableString and assigns it to the AccountSid field.
func (o *ApiV2010AccountCallCallRecording) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
// SetAccountSidNil sets the value for AccountSid to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetAccountSidNil() {
	o.AccountSid.Set(nil)
}

// UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetAccountSid() {
	o.AccountSid.Unset()
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetApiVersion() string {
	if o == nil || o.ApiVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion.Get()
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiVersion.Get(), o.ApiVersion.IsSet()
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasApiVersion() bool {
	if o != nil && o.ApiVersion.IsSet() {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given NullableString and assigns it to the ApiVersion field.
func (o *ApiV2010AccountCallCallRecording) SetApiVersion(v string) {
	o.ApiVersion.Set(&v)
}
// SetApiVersionNil sets the value for ApiVersion to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetApiVersionNil() {
	o.ApiVersion.Set(nil)
}

// UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetApiVersion() {
	o.ApiVersion.Unset()
}

// GetCallSid returns the CallSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetCallSid() string {
	if o == nil || o.CallSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.CallSid.Get()
}

// GetCallSidOk returns a tuple with the CallSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetCallSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallSid.Get(), o.CallSid.IsSet()
}

// HasCallSid returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasCallSid() bool {
	if o != nil && o.CallSid.IsSet() {
		return true
	}

	return false
}

// SetCallSid gets a reference to the given NullableString and assigns it to the CallSid field.
func (o *ApiV2010AccountCallCallRecording) SetCallSid(v string) {
	o.CallSid.Set(&v)
}
// SetCallSidNil sets the value for CallSid to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetCallSidNil() {
	o.CallSid.Set(nil)
}

// UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetCallSid() {
	o.CallSid.Unset()
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetChannels() int32 {
	if o == nil || o.Channels.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Channels.Get()
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels.Get(), o.Channels.IsSet()
}

// HasChannels returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasChannels() bool {
	if o != nil && o.Channels.IsSet() {
		return true
	}

	return false
}

// SetChannels gets a reference to the given NullableInt32 and assigns it to the Channels field.
func (o *ApiV2010AccountCallCallRecording) SetChannels(v int32) {
	o.Channels.Set(&v)
}
// SetChannelsNil sets the value for Channels to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetChannelsNil() {
	o.Channels.Set(nil)
}

// UnsetChannels ensures that no value is present for Channels, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetChannels() {
	o.Channels.Unset()
}

// GetConferenceSid returns the ConferenceSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetConferenceSid() string {
	if o == nil || o.ConferenceSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConferenceSid.Get()
}

// GetConferenceSidOk returns a tuple with the ConferenceSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetConferenceSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceSid.Get(), o.ConferenceSid.IsSet()
}

// HasConferenceSid returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasConferenceSid() bool {
	if o != nil && o.ConferenceSid.IsSet() {
		return true
	}

	return false
}

// SetConferenceSid gets a reference to the given NullableString and assigns it to the ConferenceSid field.
func (o *ApiV2010AccountCallCallRecording) SetConferenceSid(v string) {
	o.ConferenceSid.Set(&v)
}
// SetConferenceSidNil sets the value for ConferenceSid to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetConferenceSidNil() {
	o.ConferenceSid.Set(nil)
}

// UnsetConferenceSid ensures that no value is present for ConferenceSid, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetConferenceSid() {
	o.ConferenceSid.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetDateCreated() string {
	if o == nil || o.DateCreated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableString and assigns it to the DateCreated field.
func (o *ApiV2010AccountCallCallRecording) SetDateCreated(v string) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetDateUpdated() string {
	if o == nil || o.DateUpdated.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated.Get()
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateUpdated.Get(), o.DateUpdated.IsSet()
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasDateUpdated() bool {
	if o != nil && o.DateUpdated.IsSet() {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given NullableString and assigns it to the DateUpdated field.
func (o *ApiV2010AccountCallCallRecording) SetDateUpdated(v string) {
	o.DateUpdated.Set(&v)
}
// SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetDateUpdatedNil() {
	o.DateUpdated.Set(nil)
}

// UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetDateUpdated() {
	o.DateUpdated.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetDuration() string {
	if o == nil || o.Duration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableString and assigns it to the Duration field.
func (o *ApiV2010AccountCallCallRecording) SetDuration(v string) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetDuration() {
	o.Duration.Unset()
}

// GetEncryptionDetails returns the EncryptionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetEncryptionDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EncryptionDetails
}

// GetEncryptionDetailsOk returns a tuple with the EncryptionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetEncryptionDetailsOk() (*interface{}, bool) {
	if o == nil || o.EncryptionDetails == nil {
		return nil, false
	}
	return &o.EncryptionDetails, true
}

// HasEncryptionDetails returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasEncryptionDetails() bool {
	if o != nil && o.EncryptionDetails != nil {
		return true
	}

	return false
}

// SetEncryptionDetails gets a reference to the given interface{} and assigns it to the EncryptionDetails field.
func (o *ApiV2010AccountCallCallRecording) SetEncryptionDetails(v interface{}) {
	o.EncryptionDetails = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *ApiV2010AccountCallCallRecording) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetPrice() float32 {
	if o == nil || o.Price.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *ApiV2010AccountCallCallRecording) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetPrice() {
	o.Price.Unset()
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetPriceUnit() string {
	if o == nil || o.PriceUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.PriceUnit.Get()
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetPriceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceUnit.Get(), o.PriceUnit.IsSet()
}

// HasPriceUnit returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasPriceUnit() bool {
	if o != nil && o.PriceUnit.IsSet() {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given NullableString and assigns it to the PriceUnit field.
func (o *ApiV2010AccountCallCallRecording) SetPriceUnit(v string) {
	o.PriceUnit.Set(&v)
}
// SetPriceUnitNil sets the value for PriceUnit to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetPriceUnitNil() {
	o.PriceUnit.Set(nil)
}

// UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetPriceUnit() {
	o.PriceUnit.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *ApiV2010AccountCallCallRecording) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetSid() {
	o.Sid.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *ApiV2010AccountCallCallRecording) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetSource() {
	o.Source.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetStartTime() string {
	if o == nil || o.StartTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableString and assigns it to the StartTime field.
func (o *ApiV2010AccountCallCallRecording) SetStartTime(v string) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ApiV2010AccountCallCallRecording) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetStatus() {
	o.Status.Unset()
}

// GetTrack returns the Track field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetTrack() string {
	if o == nil || o.Track.Get() == nil {
		var ret string
		return ret
	}
	return *o.Track.Get()
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetTrackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Track.Get(), o.Track.IsSet()
}

// HasTrack returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasTrack() bool {
	if o != nil && o.Track.IsSet() {
		return true
	}

	return false
}

// SetTrack gets a reference to the given NullableString and assigns it to the Track field.
func (o *ApiV2010AccountCallCallRecording) SetTrack(v string) {
	o.Track.Set(&v)
}
// SetTrackNil sets the value for Track to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetTrackNil() {
	o.Track.Set(nil)
}

// UnsetTrack ensures that no value is present for Track, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetTrack() {
	o.Track.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2010AccountCallCallRecording) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2010AccountCallCallRecording) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *ApiV2010AccountCallCallRecording) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *ApiV2010AccountCallCallRecording) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *ApiV2010AccountCallCallRecording) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *ApiV2010AccountCallCallRecording) UnsetUri() {
	o.Uri.Unset()
}

func (o ApiV2010AccountCallCallRecording) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSid.IsSet() {
		toSerialize["account_sid"] = o.AccountSid.Get()
	}
	if o.ApiVersion.IsSet() {
		toSerialize["api_version"] = o.ApiVersion.Get()
	}
	if o.CallSid.IsSet() {
		toSerialize["call_sid"] = o.CallSid.Get()
	}
	if o.Channels.IsSet() {
		toSerialize["channels"] = o.Channels.Get()
	}
	if o.ConferenceSid.IsSet() {
		toSerialize["conference_sid"] = o.ConferenceSid.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["date_created"] = o.DateCreated.Get()
	}
	if o.DateUpdated.IsSet() {
		toSerialize["date_updated"] = o.DateUpdated.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.EncryptionDetails != nil {
		toSerialize["encryption_details"] = o.EncryptionDetails
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.PriceUnit.IsSet() {
		toSerialize["price_unit"] = o.PriceUnit.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Track.IsSet() {
		toSerialize["track"] = o.Track.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiV2010AccountCallCallRecording struct {
	value *ApiV2010AccountCallCallRecording
	isSet bool
}

func (v NullableApiV2010AccountCallCallRecording) Get() *ApiV2010AccountCallCallRecording {
	return v.value
}

func (v *NullableApiV2010AccountCallCallRecording) Set(val *ApiV2010AccountCallCallRecording) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2010AccountCallCallRecording) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2010AccountCallCallRecording) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2010AccountCallCallRecording(val *ApiV2010AccountCallCallRecording) *NullableApiV2010AccountCallCallRecording {
	return &NullableApiV2010AccountCallCallRecording{value: val, isSet: true}
}

func (v NullableApiV2010AccountCallCallRecording) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2010AccountCallCallRecording) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


