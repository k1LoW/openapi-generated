# ApiV2010AccountCallSiprec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created this resource | [optional] 
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was last updated | [optional] 
**Name** | Pointer to **NullableString** | The name of this resource | [optional] 
**Sid** | Pointer to **NullableString** | The SID of the Siprec resource. | [optional] 
**Status** | Pointer to **NullableString** | The status - one of &#x60;stopped&#x60;, &#x60;in-progress&#x60; | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountCallSiprec

`func NewApiV2010AccountCallSiprec() *ApiV2010AccountCallSiprec`

NewApiV2010AccountCallSiprec instantiates a new ApiV2010AccountCallSiprec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountCallSiprecWithDefaults

`func NewApiV2010AccountCallSiprecWithDefaults() *ApiV2010AccountCallSiprec`

NewApiV2010AccountCallSiprecWithDefaults instantiates a new ApiV2010AccountCallSiprec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountCallSiprec) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountCallSiprec) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountCallSiprec) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountCallSiprec) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountCallSiprec) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountCallSiprec) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallSid

`func (o *ApiV2010AccountCallSiprec) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountCallSiprec) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountCallSiprec) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountCallSiprec) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountCallSiprec) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountCallSiprec) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountCallSiprec) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountCallSiprec) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountCallSiprec) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountCallSiprec) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountCallSiprec) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountCallSiprec) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetName

`func (o *ApiV2010AccountCallSiprec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2010AccountCallSiprec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2010AccountCallSiprec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2010AccountCallSiprec) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2010AccountCallSiprec) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2010AccountCallSiprec) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountCallSiprec) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountCallSiprec) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountCallSiprec) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountCallSiprec) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountCallSiprec) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountCallSiprec) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountCallSiprec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountCallSiprec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountCallSiprec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountCallSiprec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountCallSiprec) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountCallSiprec) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountCallSiprec) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountCallSiprec) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountCallSiprec) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountCallSiprec) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountCallSiprec) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountCallSiprec) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


