# ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRequirements** | Pointer to **NullableString** | The type of Address resource the phone number requires | [optional] 
**Beta** | Pointer to **NullableBool** | Whether the phone number is new to the Twilio platform | [optional] 
**Capabilities** | Pointer to [**NullableApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities**](ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities.md) |  | [optional] 
**FriendlyName** | Pointer to **NullableString** | A formatted version of the phone number | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code of this phone number | [optional] 
**Lata** | Pointer to **NullableString** | The LATA of this phone number | [optional] 
**Latitude** | Pointer to **NullableString** | The latitude of this phone number&#39;s location | [optional] 
**Locality** | Pointer to **NullableString** | The locality or city of this phone number&#39;s location | [optional] 
**Longitude** | Pointer to **NullableString** | The longitude of this phone number&#39;s location | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number in E.164 format | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal or ZIP code of this phone number&#39;s location | [optional] 
**RateCenter** | Pointer to **NullableString** | The rate center of this phone number | [optional] 
**Region** | Pointer to **NullableString** | The two-letter state or province abbreviation of this phone number&#39;s location | [optional] 

## Methods

### NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational

`func NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational() *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational`

NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational instantiates a new ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNationalWithDefaults

`func NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNationalWithDefaults() *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational`

NewApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNationalWithDefaults instantiates a new ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRequirements

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetAddressRequirements() string`

GetAddressRequirements returns the AddressRequirements field if non-nil, zero value otherwise.

### GetAddressRequirementsOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetAddressRequirementsOk() (*string, bool)`

GetAddressRequirementsOk returns a tuple with the AddressRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRequirements

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetAddressRequirements(v string)`

SetAddressRequirements sets AddressRequirements field to given value.

### HasAddressRequirements

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasAddressRequirements() bool`

HasAddressRequirements returns a boolean if a field has been set.

### SetAddressRequirementsNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetAddressRequirementsNil(b bool)`

 SetAddressRequirementsNil sets the value for AddressRequirements to be an explicit nil

### UnsetAddressRequirements
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetAddressRequirements()`

UnsetAddressRequirements ensures that no value is present for AddressRequirements, not even an explicit nil
### GetBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### SetBetaNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetBetaNil(b bool)`

 SetBetaNil sets the value for Beta to be an explicit nil

### UnsetBeta
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetBeta()`

UnsetBeta ensures that no value is present for Beta, not even an explicit nil
### GetCapabilities

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetCapabilities() ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetCapabilitiesOk() (*ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetCapabilities(v ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIsoCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetLata

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLata() string`

GetLata returns the Lata field if non-nil, zero value otherwise.

### GetLataOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLataOk() (*string, bool)`

GetLataOk returns a tuple with the Lata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLata

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLata(v string)`

SetLata sets Lata field to given value.

### HasLata

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasLata() bool`

HasLata returns a boolean if a field has been set.

### SetLataNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLataNil(b bool)`

 SetLataNil sets the value for Lata to be an explicit nil

### UnsetLata
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetLata()`

UnsetLata ensures that no value is present for Lata, not even an explicit nil
### GetLatitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLocality

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetLongitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetPhoneNumber

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPostalCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetRateCenter

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetRateCenter() string`

GetRateCenter returns the RateCenter field if non-nil, zero value otherwise.

### GetRateCenterOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetRateCenterOk() (*string, bool)`

GetRateCenterOk returns a tuple with the RateCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCenter

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetRateCenter(v string)`

SetRateCenter sets RateCenter field to given value.

### HasRateCenter

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasRateCenter() bool`

HasRateCenter returns a boolean if a field has been set.

### SetRateCenterNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetRateCenterNil(b bool)`

 SetRateCenterNil sets the value for RateCenter to be an explicit nil

### UnsetRateCenter
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetRateCenter()`

UnsetRateCenter ensures that no value is present for RateCenter, not even an explicit nil
### GetRegion

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberNational) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


