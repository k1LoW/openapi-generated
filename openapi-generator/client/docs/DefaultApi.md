# \DefaultApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](DefaultApi.md#CreateAccount) | **Post** /2010-04-01/Accounts.json | 
[**CreateAddress**](DefaultApi.md#CreateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**CreateApplication**](DefaultApi.md#CreateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**CreateCall**](DefaultApi.md#CreateCall) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls.json | 
[**CreateCallFeedbackSummary**](DefaultApi.md#CreateCallFeedbackSummary) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary.json | 
[**CreateCallRecording**](DefaultApi.md#CreateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**CreateIncomingPhoneNumber**](DefaultApi.md#CreateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**CreateIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#CreateIncomingPhoneNumberAssignedAddOn) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 
[**CreateIncomingPhoneNumberLocal**](DefaultApi.md#CreateIncomingPhoneNumberLocal) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Local.json | 
[**CreateIncomingPhoneNumberMobile**](DefaultApi.md#CreateIncomingPhoneNumberMobile) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 
[**CreateIncomingPhoneNumberTollFree**](DefaultApi.md#CreateIncomingPhoneNumberTollFree) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages.json | 
[**CreateMessageFeedback**](DefaultApi.md#CreateMessageFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Feedback.json | 
[**CreateNewKey**](DefaultApi.md#CreateNewKey) | **Post** /2010-04-01/Accounts/{AccountSid}/Keys.json | 
[**CreateNewSigningKey**](DefaultApi.md#CreateNewSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**CreateParticipant**](DefaultApi.md#CreateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**CreatePayments**](DefaultApi.md#CreatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments.json | 
[**CreateQueue**](DefaultApi.md#CreateQueue) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues.json | 
[**CreateSipAuthCallsCredentialListMapping**](DefaultApi.md#CreateSipAuthCallsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 
[**CreateSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#CreateSipAuthCallsIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 
[**CreateSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#CreateSipAuthRegistrationsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 
[**CreateSipCredential**](DefaultApi.md#CreateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**CreateSipCredentialList**](DefaultApi.md#CreateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**CreateSipCredentialListMapping**](DefaultApi.md#CreateSipCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 
[**CreateSipDomain**](DefaultApi.md#CreateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**CreateSipIpAccessControlList**](DefaultApi.md#CreateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**CreateSipIpAccessControlListMapping**](DefaultApi.md#CreateSipIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 
[**CreateSipIpAddress**](DefaultApi.md#CreateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**CreateSiprec**](DefaultApi.md#CreateSiprec) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Siprec.json | 
[**CreateStream**](DefaultApi.md#CreateStream) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Streams.json | 
[**CreateToken**](DefaultApi.md#CreateToken) | **Post** /2010-04-01/Accounts/{AccountSid}/Tokens.json | 
[**CreateUsageTrigger**](DefaultApi.md#CreateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**CreateValidationRequest**](DefaultApi.md#CreateValidationRequest) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**DeleteAddress**](DefaultApi.md#DeleteAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**DeleteCall**](DefaultApi.md#DeleteCall) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**DeleteCallFeedbackSummary**](DefaultApi.md#DeleteCallFeedbackSummary) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 
[**DeleteCallRecording**](DefaultApi.md#DeleteCallRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**DeleteConferenceRecording**](DefaultApi.md#DeleteConferenceRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**DeleteConnectApp**](DefaultApi.md#DeleteConnectApp) | **Delete** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**DeleteIncomingPhoneNumber**](DefaultApi.md#DeleteIncomingPhoneNumber) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**DeleteIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#DeleteIncomingPhoneNumberAssignedAddOn) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**DeleteKey**](DefaultApi.md#DeleteKey) | **Delete** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**DeleteMedia**](DefaultApi.md#DeleteMedia) | **Delete** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**DeleteOutgoingCallerId**](DefaultApi.md#DeleteOutgoingCallerId) | **Delete** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**DeleteParticipant**](DefaultApi.md#DeleteParticipant) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**DeleteQueue**](DefaultApi.md#DeleteQueue) | **Delete** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**DeleteRecording**](DefaultApi.md#DeleteRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json | 
[**DeleteRecordingAddOnResult**](DefaultApi.md#DeleteRecordingAddOnResult) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**DeleteRecordingAddOnResultPayload**](DefaultApi.md#DeleteRecordingAddOnResultPayload) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**DeleteRecordingTranscription**](DefaultApi.md#DeleteRecordingTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**DeleteSigningKey**](DefaultApi.md#DeleteSigningKey) | **Delete** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**DeleteSipAuthCallsCredentialListMapping**](DefaultApi.md#DeleteSipAuthCallsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**DeleteSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#DeleteSipAuthCallsIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**DeleteSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#DeleteSipAuthRegistrationsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**DeleteSipCredential**](DefaultApi.md#DeleteSipCredential) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**DeleteSipCredentialList**](DefaultApi.md#DeleteSipCredentialList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**DeleteSipCredentialListMapping**](DefaultApi.md#DeleteSipCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**DeleteSipDomain**](DefaultApi.md#DeleteSipDomain) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**DeleteSipIpAccessControlList**](DefaultApi.md#DeleteSipIpAccessControlList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**DeleteSipIpAccessControlListMapping**](DefaultApi.md#DeleteSipIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**DeleteSipIpAddress**](DefaultApi.md#DeleteSipIpAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**DeleteTranscription**](DefaultApi.md#DeleteTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**DeleteUsageTrigger**](DefaultApi.md#DeleteUsageTrigger) | **Delete** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**FetchAccount**](DefaultApi.md#FetchAccount) | **Get** /2010-04-01/Accounts/{Sid}.json | 
[**FetchAddress**](DefaultApi.md#FetchAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**FetchApplication**](DefaultApi.md#FetchApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**FetchAuthorizedConnectApp**](DefaultApi.md#FetchAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json | 
[**FetchAvailablePhoneNumberCountry**](DefaultApi.md#FetchAvailablePhoneNumberCountry) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}.json | 
[**FetchBalance**](DefaultApi.md#FetchBalance) | **Get** /2010-04-01/Accounts/{AccountSid}/Balance.json | 
[**FetchCall**](DefaultApi.md#FetchCall) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**FetchCallFeedback**](DefaultApi.md#FetchCallFeedback) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 
[**FetchCallFeedbackSummary**](DefaultApi.md#FetchCallFeedbackSummary) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 
[**FetchCallNotification**](DefaultApi.md#FetchCallNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications/{Sid}.json | 
[**FetchCallRecording**](DefaultApi.md#FetchCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**FetchConference**](DefaultApi.md#FetchConference) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json | 
[**FetchConferenceRecording**](DefaultApi.md#FetchConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**FetchConnectApp**](DefaultApi.md#FetchConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**FetchIncomingPhoneNumber**](DefaultApi.md#FetchIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**FetchIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#FetchIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**FetchIncomingPhoneNumberAssignedAddOnExtension**](DefaultApi.md#FetchIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions/{Sid}.json | 
[**FetchKey**](DefaultApi.md#FetchKey) | **Get** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**FetchMedia**](DefaultApi.md#FetchMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**FetchNotification**](DefaultApi.md#FetchNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Notifications/{Sid}.json | 
[**FetchOutgoingCallerId**](DefaultApi.md#FetchOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**FetchParticipant**](DefaultApi.md#FetchParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**FetchQueue**](DefaultApi.md#FetchQueue) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**FetchRecording**](DefaultApi.md#FetchRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json | 
[**FetchRecordingAddOnResult**](DefaultApi.md#FetchRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**FetchRecordingAddOnResultPayload**](DefaultApi.md#FetchRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**FetchRecordingTranscription**](DefaultApi.md#FetchRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**FetchShortCode**](DefaultApi.md#FetchShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 
[**FetchSigningKey**](DefaultApi.md#FetchSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**FetchSipAuthCallsCredentialListMapping**](DefaultApi.md#FetchSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**FetchSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#FetchSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#FetchSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**FetchSipCredential**](DefaultApi.md#FetchSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**FetchSipCredentialList**](DefaultApi.md#FetchSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**FetchSipCredentialListMapping**](DefaultApi.md#FetchSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**FetchSipDomain**](DefaultApi.md#FetchSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**FetchSipIpAccessControlList**](DefaultApi.md#FetchSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**FetchSipIpAccessControlListMapping**](DefaultApi.md#FetchSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipIpAddress**](DefaultApi.md#FetchSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**FetchTranscription**](DefaultApi.md#FetchTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**FetchUsageTrigger**](DefaultApi.md#FetchUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**ListAccount**](DefaultApi.md#ListAccount) | **Get** /2010-04-01/Accounts.json | 
[**ListAddress**](DefaultApi.md#ListAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**ListApplication**](DefaultApi.md#ListApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**ListAuthorizedConnectApp**](DefaultApi.md#ListAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json | 
[**ListAvailablePhoneNumberCountry**](DefaultApi.md#ListAvailablePhoneNumberCountry) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers.json | 
[**ListAvailablePhoneNumberLocal**](DefaultApi.md#ListAvailablePhoneNumberLocal) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Local.json | 
[**ListAvailablePhoneNumberMachineToMachine**](DefaultApi.md#ListAvailablePhoneNumberMachineToMachine) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/MachineToMachine.json | 
[**ListAvailablePhoneNumberMobile**](DefaultApi.md#ListAvailablePhoneNumberMobile) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Mobile.json | 
[**ListAvailablePhoneNumberNational**](DefaultApi.md#ListAvailablePhoneNumberNational) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/National.json | 
[**ListAvailablePhoneNumberSharedCost**](DefaultApi.md#ListAvailablePhoneNumberSharedCost) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/SharedCost.json | 
[**ListAvailablePhoneNumberTollFree**](DefaultApi.md#ListAvailablePhoneNumberTollFree) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/TollFree.json | 
[**ListAvailablePhoneNumberVoip**](DefaultApi.md#ListAvailablePhoneNumberVoip) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Voip.json | 
[**ListCall**](DefaultApi.md#ListCall) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls.json | 
[**ListCallEvent**](DefaultApi.md#ListCallEvent) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Events.json | 
[**ListCallNotification**](DefaultApi.md#ListCallNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications.json | 
[**ListCallRecording**](DefaultApi.md#ListCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**ListConference**](DefaultApi.md#ListConference) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences.json | 
[**ListConferenceRecording**](DefaultApi.md#ListConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings.json | 
[**ListConnectApp**](DefaultApi.md#ListConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps.json | 
[**ListDependentPhoneNumber**](DefaultApi.md#ListDependentPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{AddressSid}/DependentPhoneNumbers.json | 
[**ListIncomingPhoneNumber**](DefaultApi.md#ListIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**ListIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#ListIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 
[**ListIncomingPhoneNumberAssignedAddOnExtension**](DefaultApi.md#ListIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions.json | 
[**ListIncomingPhoneNumberLocal**](DefaultApi.md#ListIncomingPhoneNumberLocal) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Local.json | 
[**ListIncomingPhoneNumberMobile**](DefaultApi.md#ListIncomingPhoneNumberMobile) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 
[**ListIncomingPhoneNumberTollFree**](DefaultApi.md#ListIncomingPhoneNumberTollFree) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json | 
[**ListKey**](DefaultApi.md#ListKey) | **Get** /2010-04-01/Accounts/{AccountSid}/Keys.json | 
[**ListMedia**](DefaultApi.md#ListMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media.json | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members.json | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages.json | 
[**ListNotification**](DefaultApi.md#ListNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Notifications.json | 
[**ListOutgoingCallerId**](DefaultApi.md#ListOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**ListParticipant**](DefaultApi.md#ListParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**ListQueue**](DefaultApi.md#ListQueue) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues.json | 
[**ListRecording**](DefaultApi.md#ListRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings.json | 
[**ListRecordingAddOnResult**](DefaultApi.md#ListRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults.json | 
[**ListRecordingAddOnResultPayload**](DefaultApi.md#ListRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json | 
[**ListRecordingTranscription**](DefaultApi.md#ListRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions.json | 
[**ListShortCode**](DefaultApi.md#ListShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes.json | 
[**ListSigningKey**](DefaultApi.md#ListSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**ListSipAuthCallsCredentialListMapping**](DefaultApi.md#ListSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 
[**ListSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#ListSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 
[**ListSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#ListSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 
[**ListSipCredential**](DefaultApi.md#ListSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**ListSipCredentialList**](DefaultApi.md#ListSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**ListSipCredentialListMapping**](DefaultApi.md#ListSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 
[**ListSipDomain**](DefaultApi.md#ListSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**ListSipIpAccessControlList**](DefaultApi.md#ListSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**ListSipIpAccessControlListMapping**](DefaultApi.md#ListSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 
[**ListSipIpAddress**](DefaultApi.md#ListSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**ListTranscription**](DefaultApi.md#ListTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions.json | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records.json | 
[**ListUsageRecordAllTime**](DefaultApi.md#ListUsageRecordAllTime) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/AllTime.json | 
[**ListUsageRecordDaily**](DefaultApi.md#ListUsageRecordDaily) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Daily.json | 
[**ListUsageRecordLastMonth**](DefaultApi.md#ListUsageRecordLastMonth) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/LastMonth.json | 
[**ListUsageRecordMonthly**](DefaultApi.md#ListUsageRecordMonthly) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Monthly.json | 
[**ListUsageRecordThisMonth**](DefaultApi.md#ListUsageRecordThisMonth) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/ThisMonth.json | 
[**ListUsageRecordToday**](DefaultApi.md#ListUsageRecordToday) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Today.json | 
[**ListUsageRecordYearly**](DefaultApi.md#ListUsageRecordYearly) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Yearly.json | 
[**ListUsageRecordYesterday**](DefaultApi.md#ListUsageRecordYesterday) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Yesterday.json | 
[**ListUsageTrigger**](DefaultApi.md#ListUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Post** /2010-04-01/Accounts/{Sid}.json | 
[**UpdateAddress**](DefaultApi.md#UpdateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**UpdateApplication**](DefaultApi.md#UpdateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**UpdateCall**](DefaultApi.md#UpdateCall) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**UpdateCallFeedback**](DefaultApi.md#UpdateCallFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 
[**UpdateCallRecording**](DefaultApi.md#UpdateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**UpdateConference**](DefaultApi.md#UpdateConference) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json | 
[**UpdateConferenceRecording**](DefaultApi.md#UpdateConferenceRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**UpdateConnectApp**](DefaultApi.md#UpdateConnectApp) | **Post** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**UpdateIncomingPhoneNumber**](DefaultApi.md#UpdateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**UpdateKey**](DefaultApi.md#UpdateKey) | **Post** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**UpdateOutgoingCallerId**](DefaultApi.md#UpdateOutgoingCallerId) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**UpdateParticipant**](DefaultApi.md#UpdateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**UpdatePayments**](DefaultApi.md#UpdatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments/{Sid}.json | 
[**UpdateQueue**](DefaultApi.md#UpdateQueue) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**UpdateShortCode**](DefaultApi.md#UpdateShortCode) | **Post** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 
[**UpdateSigningKey**](DefaultApi.md#UpdateSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**UpdateSipCredential**](DefaultApi.md#UpdateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**UpdateSipCredentialList**](DefaultApi.md#UpdateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**UpdateSipDomain**](DefaultApi.md#UpdateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**UpdateSipIpAccessControlList**](DefaultApi.md#UpdateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**UpdateSipIpAddress**](DefaultApi.md#UpdateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**UpdateSiprec**](DefaultApi.md#UpdateSiprec) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Siprec/{Sid}.json | 
[**UpdateStream**](DefaultApi.md#UpdateStream) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Streams/{Sid}.json | 
[**UpdateUsageTrigger**](DefaultApi.md#UpdateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 



## CreateAccount

> ApiV2010Account CreateAccount(ctx).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    friendlyName := "friendlyName_example" // string | A human readable description of the account to create, defaults to `SubAccount Created at {YYYY-MM-DD HH:MM meridian}` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAccount(context.Background()).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **string** | A human readable description of the account to create, defaults to &#x60;SubAccount Created at {YYYY-MM-DD HH:MM meridian}&#x60; | 

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddress

> ApiV2010AccountAddress CreateAddress(ctx, accountSid).City(city).CustomerName(customerName).IsoCountry(isoCountry).PostalCode(postalCode).Region(region).Street(street).AutoCorrectAddress(autoCorrectAddress).EmergencyEnabled(emergencyEnabled).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
    city := "city_example" // string | The city of the new address.
    customerName := "customerName_example" // string | The name to associate with the new address.
    isoCountry := "isoCountry_example" // string | The ISO country code of the new address.
    postalCode := "postalCode_example" // string | The postal code of the new address.
    region := "region_example" // string | The state or region of the new address.
    street := "street_example" // string | The number and street address of the new address.
    autoCorrectAddress := true // bool | Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide. (optional)
    emergencyEnabled := true // bool | Whether to enable emergency calling on the new address. Can be: `true` or `false`. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the new address. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAddress(context.Background(), accountSid).City(city).CustomerName(customerName).IsoCountry(isoCountry).PostalCode(postalCode).Region(region).Street(street).AutoCorrectAddress(autoCorrectAddress).EmergencyEnabled(emergencyEnabled).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **city** | **string** | The city of the new address. | 
 **customerName** | **string** | The name to associate with the new address. | 
 **isoCountry** | **string** | The ISO country code of the new address. | 
 **postalCode** | **string** | The postal code of the new address. | 
 **region** | **string** | The state or region of the new address. | 
 **street** | **string** | The number and street address of the new address. | 
 **autoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide. | 
 **emergencyEnabled** | **bool** | Whether to enable emergency calling on the new address. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the new address. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> ApiV2010AccountApplication CreateApplication(ctx, accountSid).ApiVersion(apiVersion).FriendlyName(friendlyName).MessageStatusCallback(messageStatusCallback).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsStatusCallback(smsStatusCallback).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    apiVersion := "apiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is the account's default API version. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the new application. It can be up to 64 characters long. (optional)
    messageStatusCallback := "messageStatusCallback_example" // string | The URL we should call using a POST method to send message status information to your application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. (optional)
    smsStatusCallback := "smsStatusCallback_example" // string | The URL we should call using a POST method to send status information about SMS messages sent by the application. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`. (optional)
    voiceCallerIdLookup := true // bool | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL we should call when the phone number assigned to this application receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateApplication(context.Background(), accountSid).ApiVersion(apiVersion).FriendlyName(friendlyName).MessageStatusCallback(messageStatusCallback).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsStatusCallback(smsStatusCallback).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is the account&#39;s default API version. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the new application. It can be up to 64 characters long. | 
 **messageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application. | 
 **smsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsStatusCallback** | **string** | The URL we should call using a POST method to send status information about SMS messages sent by the application. | 
 **smsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceCallerIdLookup** | **bool** | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call. | 

### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCall

> ApiV2010AccountCall CreateCall(ctx, accountSid).From(from).To(to).ApplicationSid(applicationSid).AsyncAmd(asyncAmd).AsyncAmdStatusCallback(asyncAmdStatusCallback).AsyncAmdStatusCallbackMethod(asyncAmdStatusCallbackMethod).Byoc(byoc).CallReason(callReason).CallToken(callToken).CallerId(callerId).FallbackMethod(fallbackMethod).FallbackUrl(fallbackUrl).MachineDetection(machineDetection).MachineDetectionSilenceTimeout(machineDetectionSilenceTimeout).MachineDetectionSpeechEndThreshold(machineDetectionSpeechEndThreshold).MachineDetectionSpeechThreshold(machineDetectionSpeechThreshold).MachineDetectionTimeout(machineDetectionTimeout).Method(method).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).SendDigits(sendDigits).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).StatusCallback(statusCallback).StatusCallbackEvent(statusCallbackEvent).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Timeout(timeout).Trim(trim).Twiml(twiml).Url(url).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    from := "from_example" // string | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `From` must also be a phone number.
    to := "to_example" // string | The phone number, SIP address, or client identifier to call.
    applicationSid := "applicationSid_example" // string | The SID of the Application resource that will handle the call, if the call will be handled by an application. (optional)
    asyncAmd := "asyncAmd_example" // string | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: `true` or `false`. (optional)
    asyncAmdStatusCallback := "asyncAmdStatusCallback_example" // string | The URL that we should call using the `async_amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax. (optional)
    asyncAmdStatusCallbackMethod := "asyncAmdStatusCallbackMethod_example" // string | The HTTP method we should use when calling the `async_amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. (optional)
    byoc := "byoc_example" // string | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta) (optional)
    callReason := "callReason_example" // string | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta) (optional)
    callToken := "callToken_example" // string | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. Pass an incoming call's call_token value to a forwarded call via the call_token parameter when creating a new call. A forwarded call should bear the same CallerID of the original incoming call. (optional)
    callerId := "callerId_example" // string | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. (optional)
    fallbackMethod := "fallbackMethod_example" // string | The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    fallbackUrl := "fallbackUrl_example" // string | The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    machineDetection := "machineDetection_example" // string | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection). (optional)
    machineDetectionSilenceTimeout := int32(56) // int32 | The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000. (optional)
    machineDetectionSpeechEndThreshold := int32(56) // int32 | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200. (optional)
    machineDetectionSpeechThreshold := int32(56) // int32 | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400. (optional)
    machineDetectionTimeout := int32(56) // int32 | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds. (optional)
    method := "method_example" // string | The HTTP method we should use when calling the `url` parameter's value. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    record := true // bool | Whether to record the call. Can be `true` to record the phone call, or `false` to not. The default is `false`. The `recording_url` is sent to the `status_callback` URL. (optional)
    recordingChannels := "recordingChannels_example" // string | The number of channels in the final recording. Can be: `mono` or `dual`. The default is `mono`. `mono` records both legs of the call in a single channel of the recording file. `dual` records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call. (optional)
    recordingStatusCallback := "recordingStatusCallback_example" // string | The URL that we call when the recording is available to be accessed. (optional)
    recordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording status events that will trigger calls to the URL specified in `recording_status_callback`. Can be: `in-progress`, `completed` and `absent`. Defaults to `completed`. Separate  multiple values with a space. (optional)
    recordingStatusCallbackMethod := "recordingStatusCallbackMethod_example" // string | The HTTP method we should use when calling the `recording_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. (optional)
    recordingTrack := "recordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio. (optional)
    sendDigits := "sendDigits_example" // string | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (`0`-`9`), '`#`', '`*`' and '`w`', to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be `ww1234#`. Remember to URL-encode this string, since the '`#`' character has special meaning in a URL. If both `SendDigits` and `MachineDetection` parameters are provided, then `MachineDetection` will be ignored. (optional)
    sipAuthPassword := "sipAuthPassword_example" // string | The password required to authenticate the user account specified in `sip_auth_username`. (optional)
    sipAuthUsername := "sipAuthUsername_example" // string | The username used to authenticate the caller making a SIP call. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). (optional)
    statusCallbackEvent := []string{"Inner_example"} // []string | The call progress events that we will send to the `status_callback` URL. Can be: `initiated`, `ringing`, `answered`, and `completed`. If no event is specified, we send the `completed` status. If you want to receive multiple events, specify each one in a separate `status_callback_event` parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample=code-create-a-call-resource-and-specify-a-statuscallbackevent&code-sdk-version=json). If an `application_sid` is present, this parameter is ignored. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use when calling the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    timeLimit := int32(56) // int32 | The maximum duration of the call in seconds. Constraints depend on account and configuration. (optional)
    timeout := int32(56) // int32 | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is `60` seconds and the maximum is `600` seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as `15` seconds, to hang up before reaching an answering machine or voicemail. (optional)
    trim := "trim_example" // string | Whether to trim any leading and trailing silence from the recording. Can be: `trim-silence` or `do-not-trim` and the default is `trim-silence`. (optional)
    twiml := "twiml_example" // string | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both `twiml` and `url` are provided then `twiml` parameter will be ignored. Max 4000 characters. (optional)
    url := "url_example" // string | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCall(context.Background(), accountSid).From(from).To(to).ApplicationSid(applicationSid).AsyncAmd(asyncAmd).AsyncAmdStatusCallback(asyncAmdStatusCallback).AsyncAmdStatusCallbackMethod(asyncAmdStatusCallbackMethod).Byoc(byoc).CallReason(callReason).CallToken(callToken).CallerId(callerId).FallbackMethod(fallbackMethod).FallbackUrl(fallbackUrl).MachineDetection(machineDetection).MachineDetectionSilenceTimeout(machineDetectionSilenceTimeout).MachineDetectionSpeechEndThreshold(machineDetectionSpeechEndThreshold).MachineDetectionSpeechThreshold(machineDetectionSpeechThreshold).MachineDetectionTimeout(machineDetectionTimeout).Method(method).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).SendDigits(sendDigits).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).StatusCallback(statusCallback).StatusCallbackEvent(statusCallbackEvent).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Timeout(timeout).Trim(trim).Twiml(twiml).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;From&#x60; must also be a phone number. | 
 **to** | **string** | The phone number, SIP address, or client identifier to call. | 
 **applicationSid** | **string** | The SID of the Application resource that will handle the call, if the call will be handled by an application. | 
 **asyncAmd** | **string** | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **asyncAmdStatusCallback** | **string** | The URL that we should call using the &#x60;async_amd_status_callback_method&#x60; to notify customer application whether the call was answered by human, machine or fax. | 
 **asyncAmdStatusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;async_amd_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **byoc** | **string** | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta) | 
 **callReason** | **string** | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta) | 
 **callToken** | **string** | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. Pass an incoming call&#39;s call_token value to a forwarded call via the call_token parameter when creating a new call. A forwarded call should bear the same CallerID of the original incoming call. | 
 **callerId** | **string** | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;. | 
 **fallbackMethod** | **string** | The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **fallbackUrl** | **string** | The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **machineDetection** | **string** | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: &#x60;Enable&#x60; or &#x60;DetectMessageEnd&#x60;. Use &#x60;Enable&#x60; if you would like us to return &#x60;AnsweredBy&#x60; as soon as the called party is identified. Use &#x60;DetectMessageEnd&#x60;, if you would like to leave a message on an answering machine. If &#x60;send_digits&#x60; is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection). | 
 **machineDetectionSilenceTimeout** | **int32** | The number of milliseconds of initial silence after which an &#x60;unknown&#x60; AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000. | 
 **machineDetectionSpeechEndThreshold** | **int32** | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200. | 
 **machineDetectionSpeechThreshold** | **int32** | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400. | 
 **machineDetectionTimeout** | **int32** | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with &#x60;AnsweredBy&#x60; of &#x60;unknown&#x60;. The default timeout is 30 seconds. | 
 **method** | **string** | The HTTP method we should use when calling the &#x60;url&#x60; parameter&#39;s value. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **record** | **bool** | Whether to record the call. Can be &#x60;true&#x60; to record the phone call, or &#x60;false&#x60; to not. The default is &#x60;false&#x60;. The &#x60;recording_url&#x60; is sent to the &#x60;status_callback&#x60; URL. | 
 **recordingChannels** | **string** | The number of channels in the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60;. The default is &#x60;mono&#x60;. &#x60;mono&#x60; records both legs of the call in a single channel of the recording file. &#x60;dual&#x60; records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call. | 
 **recordingStatusCallback** | **string** | The URL that we call when the recording is available to be accessed. | 
 **recordingStatusCallbackEvent** | **[]string** | The recording status events that will trigger calls to the URL specified in &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Defaults to &#x60;completed&#x60;. Separate  multiple values with a space. | 
 **recordingStatusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **recordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio. | 
 **sendDigits** | **string** | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (&#x60;0&#x60;-&#x60;9&#x60;), &#39;&#x60;#&#x60;&#39;, &#39;&#x60;*&#x60;&#39; and &#39;&#x60;w&#x60;&#39;, to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be &#x60;ww1234#&#x60;. Remember to URL-encode this string, since the &#39;&#x60;#&#x60;&#39; character has special meaning in a URL. If both &#x60;SendDigits&#x60; and &#x60;MachineDetection&#x60; parameters are provided, then &#x60;MachineDetection&#x60; will be ignored. | 
 **sipAuthPassword** | **string** | The password required to authenticate the user account specified in &#x60;sip_auth_username&#x60;. | 
 **sipAuthUsername** | **string** | The username used to authenticate the caller making a SIP call. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). | 
 **statusCallbackEvent** | **[]string** | The call progress events that we will send to the &#x60;status_callback&#x60; URL. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. If no event is specified, we send the &#x60;completed&#x60; status. If you want to receive multiple events, specify each one in a separate &#x60;status_callback_event&#x60; parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample&#x3D;code-create-a-call-resource-and-specify-a-statuscallbackevent&amp;code-sdk-version&#x3D;json). If an &#x60;application_sid&#x60; is present, this parameter is ignored. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **timeLimit** | **int32** | The maximum duration of the call in seconds. Constraints depend on account and configuration. | 
 **timeout** | **int32** | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is &#x60;60&#x60; seconds and the maximum is &#x60;600&#x60; seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as &#x60;15&#x60; seconds, to hang up before reaching an answering machine or voicemail. | 
 **trim** | **string** | Whether to trim any leading and trailing silence from the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;trim-silence&#x60;. | 
 **twiml** | **string** | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both &#x60;twiml&#x60; and &#x60;url&#x60; are provided then &#x60;twiml&#x60; parameter will be ignored. Max 4000 characters. | 
 **url** | **string** | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). | 

### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary CreateCallFeedbackSummary(ctx, accountSid).EndDate(endDate).StartDate(startDate).IncludeSubaccounts(includeSubaccounts).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    endDate := time.Now() // string | Only include feedback given on or before this date. Format is `YYYY-MM-DD` and specified in UTC.
    startDate := time.Now() // string | Only include feedback given on or after this date. Format is `YYYY-MM-DD` and specified in UTC.
    includeSubaccounts := true // bool | Whether to also include Feedback resources from all subaccounts. `true` includes feedback from all subaccounts and `false`, the default, includes feedback from only the specified account. (optional)
    statusCallback := "statusCallback_example" // string | The URL that we will request when the feedback summary is complete. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method (`GET` or `POST`) we use to make the request to the `StatusCallback` URL. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCallFeedbackSummary(context.Background(), accountSid).EndDate(endDate).StartDate(startDate).IncludeSubaccounts(includeSubaccounts).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCallFeedbackSummary`: ApiV2010AccountCallCallFeedbackSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCallFeedbackSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallFeedbackSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **string** | Only include feedback given on or before this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
 **startDate** | **string** | Only include feedback given on or after this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
 **includeSubaccounts** | **bool** | Whether to also include Feedback resources from all subaccounts. &#x60;true&#x60; includes feedback from all subaccounts and &#x60;false&#x60;, the default, includes feedback from only the specified account. | 
 **statusCallback** | **string** | The URL that we will request when the feedback summary is complete. | 
 **statusCallbackMethod** | **string** | The HTTP method (&#x60;GET&#x60; or &#x60;POST&#x60;) we use to make the request to the &#x60;StatusCallback&#x60; URL. | 

### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](ApiV2010AccountCallCallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallRecording

> ApiV2010AccountCallCallRecording CreateCallRecording(ctx, accountSid, callSid).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).Trim(trim).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    callSid := "callSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.
    recordingChannels := "recordingChannels_example" // string | The number of channels used in the recording. Can be: `mono` or `dual` and the default is `mono`. `mono` records all parties of the call into one channel. `dual` records each party of a 2-party call into separate channels. (optional)
    recordingStatusCallback := "recordingStatusCallback_example" // string | The URL we should call using the `recording_status_callback_method` on each recording event specified in  `recording_status_callback_event`. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback). (optional)
    recordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space. (optional)
    recordingStatusCallbackMethod := "recordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    recordingTrack := "recordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio. (optional)
    trim := "trim_example" // string | Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCallRecording(context.Background(), accountSid, callSid).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).Trim(trim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**callSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordingChannels** | **string** | The number of channels used in the recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. &#x60;mono&#x60; records all parties of the call into one channel. &#x60;dual&#x60; records each party of a 2-party call into separate channels. | 
 **recordingStatusCallback** | **string** | The URL we should call using the &#x60;recording_status_callback_method&#x60; on each recording event specified in  &#x60;recording_status_callback_event&#x60;. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback). | 
 **recordingStatusCallbackEvent** | **[]string** | The recording status events on which we should call the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60; and the default is &#x60;completed&#x60;. Separate multiple event values with a space. | 
 **recordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **recordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio. | 
 **trim** | **string** | Whether to trim any leading and trailing silence in the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;do-not-trim&#x60;. &#x60;trim-silence&#x60; trims the silence from the beginning and end of the recording and &#x60;do-not-trim&#x60; does not. | 

### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber CreateIncomingPhoneNumber(ctx, accountSid).AddressSid(addressSid).ApiVersion(apiVersion).AreaCode(areaCode).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).PhoneNumber(phoneNumber).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    addressSid := "addressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    apiVersion := "apiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    areaCode := "areaCode_example" // string | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an `area_code` or a `phone_number`.** (US and Canada only). (optional)
    bundleSid := "bundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    emergencyAddressSid := "emergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    emergencyStatus := "emergencyStatus_example" // string | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number. (optional)
    identitySid := "identitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    smsApplicationSid := "smsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    trunkSid := "trunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    voiceApplicationSid := "voiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    voiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceReceiveMode := "voiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIncomingPhoneNumber(context.Background(), accountSid).AddressSid(addressSid).ApiVersion(apiVersion).AreaCode(areaCode).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).PhoneNumber(phoneNumber).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncomingPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **apiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **areaCode** | **string** | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an &#x60;area_code&#x60; or a &#x60;phone_number&#x60;.** (US and Canada only). | 
 **bundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **emergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **emergencyStatus** | **string** | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number. | 
 **identitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **phoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **smsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **trunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **voiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **voiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **voiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn CreateIncomingPhoneNumberAssignedAddOn(ctx, accountSid, resourceSid).InstalledAddOnSid(installedAddOnSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to assign the Add-on.
    installedAddOnSid := "installedAddOnSid_example" // string | The SID that identifies the Add-on installation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIncomingPhoneNumberAssignedAddOn(context.Background(), accountSid, resourceSid).InstalledAddOnSid(installedAddOnSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberAssignedAddOn`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**resourceSid** | **string** | The SID of the Phone Number to assign the Add-on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncomingPhoneNumberAssignedAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installedAddOnSid** | **string** | The SID that identifies the Add-on installation. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberLocal

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal CreateIncomingPhoneNumberLocal(ctx, accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    phoneNumber := "phoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
    addressSid := "addressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    apiVersion := "apiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    bundleSid := "bundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    emergencyAddressSid := "emergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    emergencyStatus := "emergencyStatus_example" // string | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    identitySid := "identitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    smsApplicationSid := "smsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    trunkSid := "trunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    voiceApplicationSid := "voiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    voiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceReceiveMode := "voiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIncomingPhoneNumberLocal(context.Background(), accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberLocal`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncomingPhoneNumberLocalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **addressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **apiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **bundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **emergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **emergencyStatus** | **string** | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **identitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **smsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **trunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **voiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **voiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **voiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberMobile

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile CreateIncomingPhoneNumberMobile(ctx, accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    phoneNumber := "phoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
    addressSid := "addressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    apiVersion := "apiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    bundleSid := "bundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    emergencyAddressSid := "emergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    emergencyStatus := "emergencyStatus_example" // string | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number. (optional)
    identitySid := "identitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    smsApplicationSid := "smsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those of the application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    trunkSid := "trunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    voiceApplicationSid := "voiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    voiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceReceiveMode := "voiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIncomingPhoneNumberMobile(context.Background(), accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberMobile`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncomingPhoneNumberMobileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **addressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **apiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **bundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **emergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **emergencyStatus** | **string** | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number. | 
 **identitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **smsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those of the application. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **trunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **voiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **voiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **voiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberTollFree

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree CreateIncomingPhoneNumberTollFree(ctx, accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    phoneNumber := "phoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
    addressSid := "addressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    apiVersion := "apiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    bundleSid := "bundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    emergencyAddressSid := "emergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    emergencyStatus := "emergencyStatus_example" // string | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    identitySid := "identitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations. (optional)
    smsApplicationSid := "smsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    trunkSid := "trunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    voiceApplicationSid := "voiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    voiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceReceiveMode := "voiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIncomingPhoneNumberTollFree(context.Background(), accountSid).PhoneNumber(phoneNumber).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberTollFree`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncomingPhoneNumberTollFreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **addressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **apiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **bundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **emergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **emergencyStatus** | **string** | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **identitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations. | 
 **smsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all &#x60;sms_*_url&#x60; values and use those of the application. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **trunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **voiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **voiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **voiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ApiV2010AccountMessage CreateMessage(ctx, accountSid).To(to).AddressRetention(addressRetention).ApplicationSid(applicationSid).Attempt(attempt).Body(body).ContentRetention(contentRetention).ForceDelivery(forceDelivery).From(from).MaxPrice(maxPrice).MediaUrl(mediaUrl).MessagingServiceSid(messagingServiceSid).PersistentAction(persistentAction).ProvideFeedback(provideFeedback).ScheduleType(scheduleType).SendAsMms(sendAsMms).SendAt(sendAt).SmartEncoded(smartEncoded).StatusCallback(statusCallback).ValidityPeriod(validityPeriod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    to := "to_example" // string | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels.
    addressRetention := "addressRetention_example" // string | Determines if the address can be stored or obfuscated based on privacy settings (optional)
    applicationSid := "applicationSid_example" // string | The SID of the application that should receive message status. We POST a `message_sid` parameter and a `message_status` parameter with a value of `sent` or `failed` to the [application](https://www.twilio.com/docs/usage/api/applications)'s `message_status_callback`. If a `status_callback` parameter is also passed, it will be ignored and the application's `message_status_callback` parameter will be used. (optional)
    attempt := int32(56) // int32 | Total number of attempts made ( including this ) to send out the message regardless of the provider used (optional)
    body := "body_example" // string | The text of the message you want to send. Can be up to 1,600 characters in length. (optional)
    contentRetention := "contentRetention_example" // string | Determines if the message content can be stored or redacted based on privacy settings (optional)
    forceDelivery := true // bool | Reserved (optional)
    from := "from_example" // string | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using `messaging_service_sid`, this parameter must be empty. (optional)
    maxPrice := float32(8.14) // float32 | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds `max_price`, the message will fail and a status of `Failed` is sent to the status callback. If `MaxPrice` is not set, the message cost is not checked. (optional)
    mediaUrl := []string{"Inner_example"} // []string | The URL of the media to send with the message. The media can be of type `gif`, `png`, and `jpeg` and will be formatted correctly on the recipient's device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple `media_url` parameters in the POST request. You can include up to 10 `media_url` parameters per message. You can send images in an SMS message in only the US and Canada. (optional)
    messagingServiceSid := "messagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the `from` parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the `from` phone number for delivery. (optional)
    persistentAction := []string{"Inner_example"} // []string | Rich actions for Channels Messages. (optional)
    provideFeedback := true // bool | Whether to confirm delivery of the message. Set this value to `true` if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is `false` by default. (optional)
    scheduleType := "scheduleType_example" // string | Indicates your intent to schedule a message. Pass the value `fixed` to schedule a message at a fixed time. (optional)
    sendAsMms := true // bool | If set to True, Twilio will deliver the message as a single MMS message, regardless of the presence of media. (optional)
    sendAt := time.Now() // time.Time | The time that Twilio will send the message. Must be in ISO 8601 format. (optional)
    smartEncoded := true // bool | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: `true` or `false`. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If specified, we POST these message status changes to the URL: `queued`, `failed`, `sent`, `delivered`, or `undelivered`. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including `MessageSid`, `MessageStatus`, and `ErrorCode`. If you include this parameter with the `messaging_service_sid`, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed. (optional)
    validityPeriod := int32(56) // int32 | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMessage(context.Background(), accountSid).To(to).AddressRetention(addressRetention).ApplicationSid(applicationSid).Attempt(attempt).Body(body).ContentRetention(contentRetention).ForceDelivery(forceDelivery).From(from).MaxPrice(maxPrice).MediaUrl(mediaUrl).MessagingServiceSid(messagingServiceSid).PersistentAction(persistentAction).ProvideFeedback(provideFeedback).ScheduleType(scheduleType).SendAsMms(sendAsMms).SendAt(sendAt).SmartEncoded(smartEncoded).StatusCallback(statusCallback).ValidityPeriod(validityPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels. | 
 **addressRetention** | **string** | Determines if the address can be stored or obfuscated based on privacy settings | 
 **applicationSid** | **string** | The SID of the application that should receive message status. We POST a &#x60;message_sid&#x60; parameter and a &#x60;message_status&#x60; parameter with a value of &#x60;sent&#x60; or &#x60;failed&#x60; to the [application](https://www.twilio.com/docs/usage/api/applications)&#39;s &#x60;message_status_callback&#x60;. If a &#x60;status_callback&#x60; parameter is also passed, it will be ignored and the application&#39;s &#x60;message_status_callback&#x60; parameter will be used. | 
 **attempt** | **int32** | Total number of attempts made ( including this ) to send out the message regardless of the provider used | 
 **body** | **string** | The text of the message you want to send. Can be up to 1,600 characters in length. | 
 **contentRetention** | **string** | Determines if the message content can be stored or redacted based on privacy settings | 
 **forceDelivery** | **bool** | Reserved | 
 **from** | **string** | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using &#x60;messaging_service_sid&#x60;, this parameter must be empty. | 
 **maxPrice** | **float32** | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds &#x60;max_price&#x60;, the message will fail and a status of &#x60;Failed&#x60; is sent to the status callback. If &#x60;MaxPrice&#x60; is not set, the message cost is not checked. | 
 **mediaUrl** | **[]string** | The URL of the media to send with the message. The media can be of type &#x60;gif&#x60;, &#x60;png&#x60;, and &#x60;jpeg&#x60; and will be formatted correctly on the recipient&#39;s device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple &#x60;media_url&#x60; parameters in the POST request. You can include up to 10 &#x60;media_url&#x60; parameters per message. You can send images in an SMS message in only the US and Canada. | 
 **messagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the &#x60;from&#x60; parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the &#x60;from&#x60; phone number for delivery. | 
 **persistentAction** | **[]string** | Rich actions for Channels Messages. | 
 **provideFeedback** | **bool** | Whether to confirm delivery of the message. Set this value to &#x60;true&#x60; if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is &#x60;false&#x60; by default. | 
 **scheduleType** | **string** | Indicates your intent to schedule a message. Pass the value &#x60;fixed&#x60; to schedule a message at a fixed time. | 
 **sendAsMms** | **bool** | If set to True, Twilio will deliver the message as a single MMS message, regardless of the presence of media. | 
 **sendAt** | **time.Time** | The time that Twilio will send the message. Must be in ISO 8601 format. | 
 **smartEncoded** | **bool** | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If specified, we POST these message status changes to the URL: &#x60;queued&#x60;, &#x60;failed&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, or &#x60;undelivered&#x60;. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including &#x60;MessageSid&#x60;, &#x60;MessageStatus&#x60;, and &#x60;ErrorCode&#x60;. If you include this parameter with the &#x60;messaging_service_sid&#x60;, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed. | 
 **validityPeriod** | **int32** | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds. | 

### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessageFeedback

> ApiV2010AccountMessageMessageFeedback CreateMessageFeedback(ctx, accountSid, messageSid).Outcome(outcome).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    messageSid := "messageSid_example" // string | The SID of the Message resource for which the feedback was provided.
    outcome := "outcome_example" // string | Whether the feedback has arrived. Can be: `unconfirmed` or `confirmed`. If `provide_feedback`=`true` in [the initial HTTP POST](https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource), the initial value of this property is `unconfirmed`. After the message arrives, update the value to `confirmed`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMessageFeedback(context.Background(), accountSid, messageSid).Outcome(outcome).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessageFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessageFeedback`: ApiV2010AccountMessageMessageFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessageFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**messageSid** | **string** | The SID of the Message resource for which the feedback was provided. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **outcome** | **string** | Whether the feedback has arrived. Can be: &#x60;unconfirmed&#x60; or &#x60;confirmed&#x60;. If &#x60;provide_feedback&#x60;&#x3D;&#x60;true&#x60; in [the initial HTTP POST](https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource), the initial value of this property is &#x60;unconfirmed&#x60;. After the message arrives, update the value to &#x60;confirmed&#x60;. | 

### Return type

[**ApiV2010AccountMessageMessageFeedback**](ApiV2010AccountMessageMessageFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewKey

> ApiV2010AccountNewKey CreateNewKey(ctx, accountSid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateNewKey(context.Background(), accountSid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewKey`: ApiV2010AccountNewKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountNewKey**](ApiV2010AccountNewKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewSigningKey

> ApiV2010AccountNewSigningKey CreateNewSigningKey(ctx, accountSid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateNewSigningKey(context.Background(), accountSid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSigningKey`: ApiV2010AccountNewSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountNewSigningKey**](ApiV2010AccountNewSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipant

> ApiV2010AccountConferenceParticipant CreateParticipant(ctx, accountSid, conferenceSid).From(from).To(to).Beep(beep).Byoc(byoc).CallReason(callReason).CallSidToCoach(callSidToCoach).CallerId(callerId).Coaching(coaching).ConferenceRecord(conferenceRecord).ConferenceRecordingStatusCallback(conferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackEvent(conferenceRecordingStatusCallbackEvent).ConferenceRecordingStatusCallbackMethod(conferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(conferenceStatusCallback).ConferenceStatusCallbackEvent(conferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(conferenceStatusCallbackMethod).ConferenceTrim(conferenceTrim).EarlyMedia(earlyMedia).EndConferenceOnExit(endConferenceOnExit).JitterBufferSize(jitterBufferSize).Label(label).MaxParticipants(maxParticipants).Muted(muted).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).Region(region).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).StartConferenceOnEnter(startConferenceOnEnter).StatusCallback(statusCallback).StatusCallbackEvent(statusCallbackEvent).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Timeout(timeout).WaitMethod(waitMethod).WaitUrl(waitUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    conferenceSid := "conferenceSid_example" // string | The SID of the participant's conference.
    from := "from_example" // string | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `from` must also be a phone number. If `to` is sip address, this value of `from` should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
    to := "to_example" // string | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as `sip:name@company.com`. Client identifiers are formatted `client:name`. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
    beep := "beep_example" // string | Whether to play a notification beep to the conference when the participant joins. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`. (optional)
    byoc := "byoc_example" // string | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta) (optional)
    callReason := "callReason_example" // string | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta) (optional)
    callSidToCoach := "callSidToCoach_example" // string | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`. (optional)
    callerId := "callerId_example" // string | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `callerId` must also be a phone number. If `to` is sip address, this value of `callerId` should be a username portion to be used to populate the From header that is passed to the SIP endpoint. (optional)
    coaching := true // bool | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined. (optional)
    conferenceRecord := "conferenceRecord_example" // string | Whether to record the conference the participant is joining. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`. (optional)
    conferenceRecordingStatusCallback := "conferenceRecordingStatusCallback_example" // string | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available. (optional)
    conferenceRecordingStatusCallbackEvent := []string{"Inner_example"} // []string | The conference recording state changes that generate a call to `conference_recording_status_callback`. Can be: `in-progress`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'` (optional)
    conferenceRecordingStatusCallbackMethod := "conferenceRecordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    conferenceStatusCallback := "conferenceStatusCallback_example" // string | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored. (optional)
    conferenceStatusCallbackEvent := []string{"Inner_example"} // []string | The conference state changes that should generate a call to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `modify`, `speaker`, and `announcement`. Separate multiple values with a space. Defaults to `start end`. (optional)
    conferenceStatusCallbackMethod := "conferenceStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    conferenceTrim := "conferenceTrim_example" // string | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`. (optional)
    earlyMedia := true // bool | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: `true` or `false` and defaults to `true`. (optional)
    endConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`. (optional)
    jitterBufferSize := "jitterBufferSize_example" // string | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant's audio is mixed into the conference. Can be: `off`, `small`, `medium`, and `large`. Default to `large`. (optional)
    label := "label_example" // string | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant. (optional)
    maxParticipants := int32(56) // int32 | The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`. (optional)
    muted := true // bool | Whether the agent is muted in the conference. Can be `true` or `false` and the default is `false`. (optional)
    record := true // bool | Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`. (optional)
    recordingChannels := "recordingChannels_example" // string | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`. (optional)
    recordingStatusCallback := "recordingStatusCallback_example" // string | The URL that we should call using the `recording_status_callback_method` when the recording status changes. (optional)
    recordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording state changes that should generate a call to `recording_status_callback`. Can be: `started`, `in-progress`, `paused`, `resumed`, `stopped`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'`. (optional)
    recordingStatusCallbackMethod := "recordingStatusCallbackMethod_example" // string | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    recordingTrack := "recordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is sent from Twilio. `both` records the audio that is received and sent by Twilio. (optional)
    region := "region_example" // string | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`. (optional)
    sipAuthPassword := "sipAuthPassword_example" // string | The SIP password for authentication. (optional)
    sipAuthUsername := "sipAuthUsername_example" // string | The SIP username used for authentication. (optional)
    startConferenceOnEnter := true // bool | Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackEvent := []string{"Inner_example"} // []string | The conference state changes that should generate a call to `status_callback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` and `POST` and defaults to `POST`. (optional)
    timeLimit := int32(56) // int32 | The maximum duration of the call in seconds. Constraints depend on account and configuration. (optional)
    timeout := int32(56) // int32 | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between `5` and `600`, inclusive. The default value is `60`. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds. (optional)
    waitMethod := "waitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    waitUrl := "waitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateParticipant(context.Background(), accountSid, conferenceSid).From(from).To(to).Beep(beep).Byoc(byoc).CallReason(callReason).CallSidToCoach(callSidToCoach).CallerId(callerId).Coaching(coaching).ConferenceRecord(conferenceRecord).ConferenceRecordingStatusCallback(conferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackEvent(conferenceRecordingStatusCallbackEvent).ConferenceRecordingStatusCallbackMethod(conferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(conferenceStatusCallback).ConferenceStatusCallbackEvent(conferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(conferenceStatusCallbackMethod).ConferenceTrim(conferenceTrim).EarlyMedia(earlyMedia).EndConferenceOnExit(endConferenceOnExit).JitterBufferSize(jitterBufferSize).Label(label).MaxParticipants(maxParticipants).Muted(muted).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingTrack(recordingTrack).Region(region).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).StartConferenceOnEnter(startConferenceOnEnter).StatusCallback(statusCallback).StatusCallbackEvent(statusCallbackEvent).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Timeout(timeout).WaitMethod(waitMethod).WaitUrl(waitUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**conferenceSid** | **string** | The SID of the participant&#39;s conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;from&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;from&#x60; should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint. | 
 **to** | **string** | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as &#x60;sip:name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified. | 
 **beep** | **string** | Whether to play a notification beep to the conference when the participant joins. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;. | 
 **byoc** | **string** | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta) | 
 **callReason** | **string** | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta) | 
 **callSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;. | 
 **callerId** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;callerId&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;callerId&#x60; should be a username portion to be used to populate the From header that is passed to the SIP endpoint. | 
 **coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined. | 
 **conferenceRecord** | **string** | Whether to record the conference the participant is joining. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;. | 
 **conferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available. | 
 **conferenceRecordingStatusCallbackEvent** | **[]string** | The conference recording state changes that generate a call to &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60; | 
 **conferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored. | 
 **conferenceStatusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;modify&#x60;, &#x60;speaker&#x60;, and &#x60;announcement&#x60;. Separate multiple values with a space. Defaults to &#x60;start end&#x60;. | 
 **conferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceTrim** | **string** | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;. | 
 **earlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. | 
 **endConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **jitterBufferSize** | **string** | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant&#39;s audio is mixed into the conference. Can be: &#x60;off&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, and &#x60;large&#x60;. Default to &#x60;large&#x60;. | 
 **label** | **string** | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant. | 
 **maxParticipants** | **int32** | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;. | 
 **muted** | **bool** | Whether the agent is muted in the conference. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **recordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. | 
 **recordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes. | 
 **recordingStatusCallbackEvent** | **[]string** | The recording state changes that should generate a call to &#x60;recording_status_callback&#x60;. Can be: &#x60;started&#x60;, &#x60;in-progress&#x60;, &#x60;paused&#x60;, &#x60;resumed&#x60;, &#x60;stopped&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60;. | 
 **recordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **recordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is sent from Twilio. &#x60;both&#x60; records the audio that is received and sent by Twilio. | 
 **region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;. | 
 **sipAuthPassword** | **string** | The SIP password for authentication. | 
 **sipAuthUsername** | **string** | The SIP username used for authentication. | 
 **startConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; and &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **timeLimit** | **int32** | The maximum duration of the call in seconds. Constraints depend on account and configuration. | 
 **timeout** | **int32** | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between &#x60;5&#x60; and &#x60;600&#x60;, inclusive. The default value is &#x60;60&#x60;. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds. | 
 **waitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **waitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 

### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayments

> ApiV2010AccountCallPayments CreatePayments(ctx, accountSid, callSid).IdempotencyKey(idempotencyKey).StatusCallback(statusCallback).BankAccountType(bankAccountType).ChargeAmount(chargeAmount).Currency(currency).Description(description).Input(input).MinPostalCodeLength(minPostalCodeLength).Parameter(parameter).PaymentConnector(paymentConnector).PaymentMethod(paymentMethod).PostalCode(postalCode).SecurityCode(securityCode).Timeout(timeout).TokenType(tokenType).ValidCardTypes(validCardTypes).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    callSid := "callSid_example" // string | The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF.
    idempotencyKey := "idempotencyKey_example" // string | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
    statusCallback := "statusCallback_example" // string | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback)
    bankAccountType := "bankAccountType_example" // string | Type of bank account if payment source is ACH. One of `consumer-checking`, `consumer-savings`, or `commercial-checking`. The default value is `consumer-checking`. (optional)
    chargeAmount := float32(8.14) // float32 | A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with `currency` field. Leave blank or set to 0 to tokenize. (optional)
    currency := "currency_example" // string | The currency of the `charge_amount`, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is `USD` and all values allowed from the <Pay> Connector are accepted. (optional)
    description := "description_example" // string | The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions. (optional)
    input := "input_example" // string | A list of inputs that should be accepted. Currently only `dtmf` is supported. All digits captured during a pay session are redacted from the logs. (optional)
    minPostalCodeLength := int32(56) // int32 | A positive integer that is used to validate the length of the `PostalCode` inputted by the user. User must enter this many digits. (optional)
    parameter := TODO // interface{} | A single-level JSON object used to pass custom parameters to payment processors. (Required for ACH payments). The information that has to be included here depends on the <Pay> Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors). (optional)
    paymentConnector := "paymentConnector_example" // string | This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [<Pay> Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is `Default`. (optional)
    paymentMethod := "paymentMethod_example" // string | Type of payment being captured. One of `credit-card` or `ach-debit`. The default value is `credit-card`. (optional)
    postalCode := true // bool | Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is `true`. (optional)
    securityCode := true // bool | Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is `true`. (optional)
    timeout := int32(56) // int32 | The number of seconds that <Pay> should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is `5`, maximum is `600`. (optional)
    tokenType := "tokenType_example" // string | Indicates whether the payment method should be tokenized as a `one-time` or `reusable` token. The default value is `reusable`. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized. (optional)
    validCardTypes := "validCardTypes_example" // string | Credit card types separated by space that Pay should accept. The default value is `visa mastercard amex` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreatePayments(context.Background(), accountSid, callSid).IdempotencyKey(idempotencyKey).StatusCallback(statusCallback).BankAccountType(bankAccountType).ChargeAmount(chargeAmount).Currency(currency).Description(description).Input(input).MinPostalCodeLength(minPostalCodeLength).Parameter(parameter).PaymentConnector(paymentConnector).PaymentMethod(paymentMethod).PostalCode(postalCode).SecurityCode(securityCode).Timeout(timeout).TokenType(tokenType).ValidCardTypes(validCardTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayments`: ApiV2010AccountCallPayments
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**callSid** | **string** | The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. | 
 **statusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback) | 
 **bankAccountType** | **string** | Type of bank account if payment source is ACH. One of &#x60;consumer-checking&#x60;, &#x60;consumer-savings&#x60;, or &#x60;commercial-checking&#x60;. The default value is &#x60;consumer-checking&#x60;. | 
 **chargeAmount** | **float32** | A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with &#x60;currency&#x60; field. Leave blank or set to 0 to tokenize. | 
 **currency** | **string** | The currency of the &#x60;charge_amount&#x60;, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is &#x60;USD&#x60; and all values allowed from the &lt;Pay&gt; Connector are accepted. | 
 **description** | **string** | The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions. | 
 **input** | **string** | A list of inputs that should be accepted. Currently only &#x60;dtmf&#x60; is supported. All digits captured during a pay session are redacted from the logs. | 
 **minPostalCodeLength** | **int32** | A positive integer that is used to validate the length of the &#x60;PostalCode&#x60; inputted by the user. User must enter this many digits. | 
 **parameter** | [**interface{}**](interface{}.md) | A single-level JSON object used to pass custom parameters to payment processors. (Required for ACH payments). The information that has to be included here depends on the &lt;Pay&gt; Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors). | 
 **paymentConnector** | **string** | This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [&lt;Pay&gt; Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is &#x60;Default&#x60;. | 
 **paymentMethod** | **string** | Type of payment being captured. One of &#x60;credit-card&#x60; or &#x60;ach-debit&#x60;. The default value is &#x60;credit-card&#x60;. | 
 **postalCode** | **bool** | Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;. | 
 **securityCode** | **bool** | Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;. | 
 **timeout** | **int32** | The number of seconds that &lt;Pay&gt; should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is &#x60;5&#x60;, maximum is &#x60;600&#x60;. | 
 **tokenType** | **string** | Indicates whether the payment method should be tokenized as a &#x60;one-time&#x60; or &#x60;reusable&#x60; token. The default value is &#x60;reusable&#x60;. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized. | 
 **validCardTypes** | **string** | Credit card types separated by space that Pay should accept. The default value is &#x60;visa mastercard amex&#x60; | 

### Return type

[**ApiV2010AccountCallPayments**](ApiV2010AccountCallPayments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueue

> ApiV2010AccountQueue CreateQueue(ctx, accountSid).FriendlyName(friendlyName).MaxSize(maxSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long.
    maxSize := int32(56) // int32 | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateQueue(context.Background(), accountSid).FriendlyName(friendlyName).MaxSize(maxSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. | 
 **maxSize** | **int32** | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. | 

### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping CreateSipAuthCallsCredentialListMapping(ctx, accountSid, domainSid).CredentialListSid(credentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    credentialListSid := "credentialListSid_example" // string | The SID of the CredentialList resource to map to the SIP domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipAuthCallsCredentialListMapping(context.Background(), accountSid, domainSid).CredentialListSid(credentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthCallsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**domainSid** | **string** | The SID of the SIP domain that will contain the new resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipAuthCallsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping CreateSipAuthCallsIpAccessControlListMapping(ctx, accountSid, domainSid).IpAccessControlListSid(ipAccessControlListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The SID of the IpAccessControlList resource to map to the SIP domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipAuthCallsIpAccessControlListMapping(context.Background(), accountSid, domainSid).IpAccessControlListSid(ipAccessControlListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthCallsIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**domainSid** | **string** | The SID of the SIP domain that will contain the new resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipAuthCallsIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipAccessControlListSid** | **string** | The SID of the IpAccessControlList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping CreateSipAuthRegistrationsCredentialListMapping(ctx, accountSid, domainSid).CredentialListSid(credentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    credentialListSid := "credentialListSid_example" // string | The SID of the CredentialList resource to map to the SIP domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipAuthRegistrationsCredentialListMapping(context.Background(), accountSid, domainSid).CredentialListSid(credentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthRegistrationsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**domainSid** | **string** | The SID of the SIP domain that will contain the new resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipAuthRegistrationsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential CreateSipCredential(ctx, accountSid, credentialListSid).Password(password).Username(username).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    credentialListSid := "credentialListSid_example" // string | The unique id that identifies the credential list to include the created credential.
    password := "password_example" // string | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
    username := "username_example" // string | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipCredential(context.Background(), accountSid, credentialListSid).Password(password).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**credentialListSid** | **string** | The unique id that identifies the credential list to include the created credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;) | 
 **username** | **string** | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio&#39;s challenge of the initial INVITE. It can be up to 32 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialList

> ApiV2010AccountSipSipCredentialList CreateSipCredentialList(ctx, accountSid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    friendlyName := "friendlyName_example" // string | A human readable descriptive text that describes the CredentialList, up to 64 characters long.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipCredentialList(context.Background(), accountSid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | A human readable descriptive text that describes the CredentialList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping CreateSipCredentialListMapping(ctx, accountSid, domainSid).CredentialListSid(credentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped.
    credentialListSid := "credentialListSid_example" // string | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipCredentialListMapping(context.Background(), accountSid, domainSid).CredentialListSid(credentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredentialListMapping`: ApiV2010AccountSipSipDomainSipCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialListSid** | **string** | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](ApiV2010AccountSipSipDomainSipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipDomain

> ApiV2010AccountSipSipDomain CreateSipDomain(ctx, accountSid).DomainName(domainName).ByocTrunkSid(byocTrunkSid).EmergencyCallerSid(emergencyCallerSid).EmergencyCallingEnabled(emergencyCallingEnabled).FriendlyName(friendlyName).Secure(secure).SipRegistration(sipRegistration).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceStatusCallbackMethod(voiceStatusCallbackMethod).VoiceStatusCallbackUrl(voiceStatusCallbackUrl).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    domainName := "domainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`.
    byocTrunkSid := "byocTrunkSid_example" // string | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. (optional)
    emergencyCallerSid := "emergencyCallerSid_example" // string | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. (optional)
    emergencyCallingEnabled := true // bool | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the resource. It can be up to 64 characters long. (optional)
    secure := true // bool | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. (optional)
    sipRegistration := true // bool | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    voiceStatusCallbackMethod := "voiceStatusCallbackMethod_example" // string | The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`. (optional)
    voiceStatusCallbackUrl := "voiceStatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL we should when the domain receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipDomain(context.Background(), accountSid).DomainName(domainName).ByocTrunkSid(byocTrunkSid).EmergencyCallerSid(emergencyCallerSid).EmergencyCallingEnabled(emergencyCallingEnabled).FriendlyName(friendlyName).Secure(secure).SipRegistration(sipRegistration).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceStatusCallbackMethod(voiceStatusCallbackMethod).VoiceStatusCallbackUrl(voiceStatusCallbackUrl).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;. | 
 **byocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. | 
 **emergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. | 
 **emergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long. | 
 **secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. | 
 **sipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not. | 
 **voiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **voiceUrl** | **string** | The URL we should when the domain receives a call. | 

### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList CreateSipIpAccessControlList(ctx, accountSid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    friendlyName := "friendlyName_example" // string | A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipIpAccessControlList(context.Background(), accountSid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipIpAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping CreateSipIpAccessControlListMapping(ctx, accountSid, domainSid).IpAccessControlListSid(ipAccessControlListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The unique id of the IP access control list to map to the SIP domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipIpAccessControlListMapping(context.Background(), accountSid, domainSid).IpAccessControlListSid(ipAccessControlListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipAccessControlListSid** | **string** | The unique id of the IP access control list to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress CreateSipIpAddress(ctx, accountSid, ipAccessControlListSid).FriendlyName(friendlyName).IpAddress(ipAddress).CidrPrefixLength(cidrPrefixLength).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The IpAccessControlList Sid with which to associate the created IpAddress resource.
    friendlyName := "friendlyName_example" // string | A human readable descriptive text for this resource, up to 64 characters long.
    ipAddress := "ipAddress_example" // string | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
    cidrPrefixLength := int32(56) // int32 | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSipIpAddress(context.Background(), accountSid, ipAccessControlListSid).FriendlyName(friendlyName).IpAddress(ipAddress).CidrPrefixLength(cidrPrefixLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**ipAccessControlListSid** | **string** | The IpAccessControlList Sid with which to associate the created IpAddress resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A human readable descriptive text for this resource, up to 64 characters long. | 
 **ipAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. | 
 **cidrPrefixLength** | **int32** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiprec

> ApiV2010AccountCallSiprec CreateSiprec(ctx, accountSid, callSid).ConnectorName(connectorName).Name(name).Parameter1Name(parameter1Name).Parameter1Value(parameter1Value).Parameter10Name(parameter10Name).Parameter10Value(parameter10Value).Parameter11Name(parameter11Name).Parameter11Value(parameter11Value).Parameter12Name(parameter12Name).Parameter12Value(parameter12Value).Parameter13Name(parameter13Name).Parameter13Value(parameter13Value).Parameter14Name(parameter14Name).Parameter14Value(parameter14Value).Parameter15Name(parameter15Name).Parameter15Value(parameter15Value).Parameter16Name(parameter16Name).Parameter16Value(parameter16Value).Parameter17Name(parameter17Name).Parameter17Value(parameter17Value).Parameter18Name(parameter18Name).Parameter18Value(parameter18Value).Parameter19Name(parameter19Name).Parameter19Value(parameter19Value).Parameter2Name(parameter2Name).Parameter2Value(parameter2Value).Parameter20Name(parameter20Name).Parameter20Value(parameter20Value).Parameter21Name(parameter21Name).Parameter21Value(parameter21Value).Parameter22Name(parameter22Name).Parameter22Value(parameter22Value).Parameter23Name(parameter23Name).Parameter23Value(parameter23Value).Parameter24Name(parameter24Name).Parameter24Value(parameter24Value).Parameter25Name(parameter25Name).Parameter25Value(parameter25Value).Parameter26Name(parameter26Name).Parameter26Value(parameter26Value).Parameter27Name(parameter27Name).Parameter27Value(parameter27Value).Parameter28Name(parameter28Name).Parameter28Value(parameter28Value).Parameter29Name(parameter29Name).Parameter29Value(parameter29Value).Parameter3Name(parameter3Name).Parameter3Value(parameter3Value).Parameter30Name(parameter30Name).Parameter30Value(parameter30Value).Parameter31Name(parameter31Name).Parameter31Value(parameter31Value).Parameter32Name(parameter32Name).Parameter32Value(parameter32Value).Parameter33Name(parameter33Name).Parameter33Value(parameter33Value).Parameter34Name(parameter34Name).Parameter34Value(parameter34Value).Parameter35Name(parameter35Name).Parameter35Value(parameter35Value).Parameter36Name(parameter36Name).Parameter36Value(parameter36Value).Parameter37Name(parameter37Name).Parameter37Value(parameter37Value).Parameter38Name(parameter38Name).Parameter38Value(parameter38Value).Parameter39Name(parameter39Name).Parameter39Value(parameter39Value).Parameter4Name(parameter4Name).Parameter4Value(parameter4Value).Parameter40Name(parameter40Name).Parameter40Value(parameter40Value).Parameter41Name(parameter41Name).Parameter41Value(parameter41Value).Parameter42Name(parameter42Name).Parameter42Value(parameter42Value).Parameter43Name(parameter43Name).Parameter43Value(parameter43Value).Parameter44Name(parameter44Name).Parameter44Value(parameter44Value).Parameter45Name(parameter45Name).Parameter45Value(parameter45Value).Parameter46Name(parameter46Name).Parameter46Value(parameter46Value).Parameter47Name(parameter47Name).Parameter47Value(parameter47Value).Parameter48Name(parameter48Name).Parameter48Value(parameter48Value).Parameter49Name(parameter49Name).Parameter49Value(parameter49Value).Parameter5Name(parameter5Name).Parameter5Value(parameter5Value).Parameter50Name(parameter50Name).Parameter50Value(parameter50Value).Parameter51Name(parameter51Name).Parameter51Value(parameter51Value).Parameter52Name(parameter52Name).Parameter52Value(parameter52Value).Parameter53Name(parameter53Name).Parameter53Value(parameter53Value).Parameter54Name(parameter54Name).Parameter54Value(parameter54Value).Parameter55Name(parameter55Name).Parameter55Value(parameter55Value).Parameter56Name(parameter56Name).Parameter56Value(parameter56Value).Parameter57Name(parameter57Name).Parameter57Value(parameter57Value).Parameter58Name(parameter58Name).Parameter58Value(parameter58Value).Parameter59Name(parameter59Name).Parameter59Value(parameter59Value).Parameter6Name(parameter6Name).Parameter6Value(parameter6Value).Parameter60Name(parameter60Name).Parameter60Value(parameter60Value).Parameter61Name(parameter61Name).Parameter61Value(parameter61Value).Parameter62Name(parameter62Name).Parameter62Value(parameter62Value).Parameter63Name(parameter63Name).Parameter63Value(parameter63Value).Parameter64Name(parameter64Name).Parameter64Value(parameter64Value).Parameter65Name(parameter65Name).Parameter65Value(parameter65Value).Parameter66Name(parameter66Name).Parameter66Value(parameter66Value).Parameter67Name(parameter67Name).Parameter67Value(parameter67Value).Parameter68Name(parameter68Name).Parameter68Value(parameter68Value).Parameter69Name(parameter69Name).Parameter69Value(parameter69Value).Parameter7Name(parameter7Name).Parameter7Value(parameter7Value).Parameter70Name(parameter70Name).Parameter70Value(parameter70Value).Parameter71Name(parameter71Name).Parameter71Value(parameter71Value).Parameter72Name(parameter72Name).Parameter72Value(parameter72Value).Parameter73Name(parameter73Name).Parameter73Value(parameter73Value).Parameter74Name(parameter74Name).Parameter74Value(parameter74Value).Parameter75Name(parameter75Name).Parameter75Value(parameter75Value).Parameter76Name(parameter76Name).Parameter76Value(parameter76Value).Parameter77Name(parameter77Name).Parameter77Value(parameter77Value).Parameter78Name(parameter78Name).Parameter78Value(parameter78Value).Parameter79Name(parameter79Name).Parameter79Value(parameter79Value).Parameter8Name(parameter8Name).Parameter8Value(parameter8Value).Parameter80Name(parameter80Name).Parameter80Value(parameter80Value).Parameter81Name(parameter81Name).Parameter81Value(parameter81Value).Parameter82Name(parameter82Name).Parameter82Value(parameter82Value).Parameter83Name(parameter83Name).Parameter83Value(parameter83Value).Parameter84Name(parameter84Name).Parameter84Value(parameter84Value).Parameter85Name(parameter85Name).Parameter85Value(parameter85Value).Parameter86Name(parameter86Name).Parameter86Value(parameter86Value).Parameter87Name(parameter87Name).Parameter87Value(parameter87Value).Parameter88Name(parameter88Name).Parameter88Value(parameter88Value).Parameter89Name(parameter89Name).Parameter89Value(parameter89Value).Parameter9Name(parameter9Name).Parameter9Value(parameter9Value).Parameter90Name(parameter90Name).Parameter90Value(parameter90Value).Parameter91Name(parameter91Name).Parameter91Value(parameter91Value).Parameter92Name(parameter92Name).Parameter92Value(parameter92Value).Parameter93Name(parameter93Name).Parameter93Value(parameter93Value).Parameter94Name(parameter94Name).Parameter94Value(parameter94Value).Parameter95Name(parameter95Name).Parameter95Value(parameter95Value).Parameter96Name(parameter96Name).Parameter96Value(parameter96Value).Parameter97Name(parameter97Name).Parameter97Value(parameter97Value).Parameter98Name(parameter98Name).Parameter98Value(parameter98Value).Parameter99Name(parameter99Name).Parameter99Value(parameter99Value).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Siprec resource.
    callSid := "callSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Siprec resource is associated with.
    connectorName := "connectorName_example" // string | Unique name used when configuring the connector via Marketplace Add-on. (optional)
    name := "name_example" // string | The user-specified name of this Siprec, if one was given when the Siprec was created. This may be used to stop the Siprec. (optional)
    parameter1Name := "parameter1Name_example" // string | Parameter name (optional)
    parameter1Value := "parameter1Value_example" // string | Parameter value (optional)
    parameter10Name := "parameter10Name_example" // string | Parameter name (optional)
    parameter10Value := "parameter10Value_example" // string | Parameter value (optional)
    parameter11Name := "parameter11Name_example" // string | Parameter name (optional)
    parameter11Value := "parameter11Value_example" // string | Parameter value (optional)
    parameter12Name := "parameter12Name_example" // string | Parameter name (optional)
    parameter12Value := "parameter12Value_example" // string | Parameter value (optional)
    parameter13Name := "parameter13Name_example" // string | Parameter name (optional)
    parameter13Value := "parameter13Value_example" // string | Parameter value (optional)
    parameter14Name := "parameter14Name_example" // string | Parameter name (optional)
    parameter14Value := "parameter14Value_example" // string | Parameter value (optional)
    parameter15Name := "parameter15Name_example" // string | Parameter name (optional)
    parameter15Value := "parameter15Value_example" // string | Parameter value (optional)
    parameter16Name := "parameter16Name_example" // string | Parameter name (optional)
    parameter16Value := "parameter16Value_example" // string | Parameter value (optional)
    parameter17Name := "parameter17Name_example" // string | Parameter name (optional)
    parameter17Value := "parameter17Value_example" // string | Parameter value (optional)
    parameter18Name := "parameter18Name_example" // string | Parameter name (optional)
    parameter18Value := "parameter18Value_example" // string | Parameter value (optional)
    parameter19Name := "parameter19Name_example" // string | Parameter name (optional)
    parameter19Value := "parameter19Value_example" // string | Parameter value (optional)
    parameter2Name := "parameter2Name_example" // string | Parameter name (optional)
    parameter2Value := "parameter2Value_example" // string | Parameter value (optional)
    parameter20Name := "parameter20Name_example" // string | Parameter name (optional)
    parameter20Value := "parameter20Value_example" // string | Parameter value (optional)
    parameter21Name := "parameter21Name_example" // string | Parameter name (optional)
    parameter21Value := "parameter21Value_example" // string | Parameter value (optional)
    parameter22Name := "parameter22Name_example" // string | Parameter name (optional)
    parameter22Value := "parameter22Value_example" // string | Parameter value (optional)
    parameter23Name := "parameter23Name_example" // string | Parameter name (optional)
    parameter23Value := "parameter23Value_example" // string | Parameter value (optional)
    parameter24Name := "parameter24Name_example" // string | Parameter name (optional)
    parameter24Value := "parameter24Value_example" // string | Parameter value (optional)
    parameter25Name := "parameter25Name_example" // string | Parameter name (optional)
    parameter25Value := "parameter25Value_example" // string | Parameter value (optional)
    parameter26Name := "parameter26Name_example" // string | Parameter name (optional)
    parameter26Value := "parameter26Value_example" // string | Parameter value (optional)
    parameter27Name := "parameter27Name_example" // string | Parameter name (optional)
    parameter27Value := "parameter27Value_example" // string | Parameter value (optional)
    parameter28Name := "parameter28Name_example" // string | Parameter name (optional)
    parameter28Value := "parameter28Value_example" // string | Parameter value (optional)
    parameter29Name := "parameter29Name_example" // string | Parameter name (optional)
    parameter29Value := "parameter29Value_example" // string | Parameter value (optional)
    parameter3Name := "parameter3Name_example" // string | Parameter name (optional)
    parameter3Value := "parameter3Value_example" // string | Parameter value (optional)
    parameter30Name := "parameter30Name_example" // string | Parameter name (optional)
    parameter30Value := "parameter30Value_example" // string | Parameter value (optional)
    parameter31Name := "parameter31Name_example" // string | Parameter name (optional)
    parameter31Value := "parameter31Value_example" // string | Parameter value (optional)
    parameter32Name := "parameter32Name_example" // string | Parameter name (optional)
    parameter32Value := "parameter32Value_example" // string | Parameter value (optional)
    parameter33Name := "parameter33Name_example" // string | Parameter name (optional)
    parameter33Value := "parameter33Value_example" // string | Parameter value (optional)
    parameter34Name := "parameter34Name_example" // string | Parameter name (optional)
    parameter34Value := "parameter34Value_example" // string | Parameter value (optional)
    parameter35Name := "parameter35Name_example" // string | Parameter name (optional)
    parameter35Value := "parameter35Value_example" // string | Parameter value (optional)
    parameter36Name := "parameter36Name_example" // string | Parameter name (optional)
    parameter36Value := "parameter36Value_example" // string | Parameter value (optional)
    parameter37Name := "parameter37Name_example" // string | Parameter name (optional)
    parameter37Value := "parameter37Value_example" // string | Parameter value (optional)
    parameter38Name := "parameter38Name_example" // string | Parameter name (optional)
    parameter38Value := "parameter38Value_example" // string | Parameter value (optional)
    parameter39Name := "parameter39Name_example" // string | Parameter name (optional)
    parameter39Value := "parameter39Value_example" // string | Parameter value (optional)
    parameter4Name := "parameter4Name_example" // string | Parameter name (optional)
    parameter4Value := "parameter4Value_example" // string | Parameter value (optional)
    parameter40Name := "parameter40Name_example" // string | Parameter name (optional)
    parameter40Value := "parameter40Value_example" // string | Parameter value (optional)
    parameter41Name := "parameter41Name_example" // string | Parameter name (optional)
    parameter41Value := "parameter41Value_example" // string | Parameter value (optional)
    parameter42Name := "parameter42Name_example" // string | Parameter name (optional)
    parameter42Value := "parameter42Value_example" // string | Parameter value (optional)
    parameter43Name := "parameter43Name_example" // string | Parameter name (optional)
    parameter43Value := "parameter43Value_example" // string | Parameter value (optional)
    parameter44Name := "parameter44Name_example" // string | Parameter name (optional)
    parameter44Value := "parameter44Value_example" // string | Parameter value (optional)
    parameter45Name := "parameter45Name_example" // string | Parameter name (optional)
    parameter45Value := "parameter45Value_example" // string | Parameter value (optional)
    parameter46Name := "parameter46Name_example" // string | Parameter name (optional)
    parameter46Value := "parameter46Value_example" // string | Parameter value (optional)
    parameter47Name := "parameter47Name_example" // string | Parameter name (optional)
    parameter47Value := "parameter47Value_example" // string | Parameter value (optional)
    parameter48Name := "parameter48Name_example" // string | Parameter name (optional)
    parameter48Value := "parameter48Value_example" // string | Parameter value (optional)
    parameter49Name := "parameter49Name_example" // string | Parameter name (optional)
    parameter49Value := "parameter49Value_example" // string | Parameter value (optional)
    parameter5Name := "parameter5Name_example" // string | Parameter name (optional)
    parameter5Value := "parameter5Value_example" // string | Parameter value (optional)
    parameter50Name := "parameter50Name_example" // string | Parameter name (optional)
    parameter50Value := "parameter50Value_example" // string | Parameter value (optional)
    parameter51Name := "parameter51Name_example" // string | Parameter name (optional)
    parameter51Value := "parameter51Value_example" // string | Parameter value (optional)
    parameter52Name := "parameter52Name_example" // string | Parameter name (optional)
    parameter52Value := "parameter52Value_example" // string | Parameter value (optional)
    parameter53Name := "parameter53Name_example" // string | Parameter name (optional)
    parameter53Value := "parameter53Value_example" // string | Parameter value (optional)
    parameter54Name := "parameter54Name_example" // string | Parameter name (optional)
    parameter54Value := "parameter54Value_example" // string | Parameter value (optional)
    parameter55Name := "parameter55Name_example" // string | Parameter name (optional)
    parameter55Value := "parameter55Value_example" // string | Parameter value (optional)
    parameter56Name := "parameter56Name_example" // string | Parameter name (optional)
    parameter56Value := "parameter56Value_example" // string | Parameter value (optional)
    parameter57Name := "parameter57Name_example" // string | Parameter name (optional)
    parameter57Value := "parameter57Value_example" // string | Parameter value (optional)
    parameter58Name := "parameter58Name_example" // string | Parameter name (optional)
    parameter58Value := "parameter58Value_example" // string | Parameter value (optional)
    parameter59Name := "parameter59Name_example" // string | Parameter name (optional)
    parameter59Value := "parameter59Value_example" // string | Parameter value (optional)
    parameter6Name := "parameter6Name_example" // string | Parameter name (optional)
    parameter6Value := "parameter6Value_example" // string | Parameter value (optional)
    parameter60Name := "parameter60Name_example" // string | Parameter name (optional)
    parameter60Value := "parameter60Value_example" // string | Parameter value (optional)
    parameter61Name := "parameter61Name_example" // string | Parameter name (optional)
    parameter61Value := "parameter61Value_example" // string | Parameter value (optional)
    parameter62Name := "parameter62Name_example" // string | Parameter name (optional)
    parameter62Value := "parameter62Value_example" // string | Parameter value (optional)
    parameter63Name := "parameter63Name_example" // string | Parameter name (optional)
    parameter63Value := "parameter63Value_example" // string | Parameter value (optional)
    parameter64Name := "parameter64Name_example" // string | Parameter name (optional)
    parameter64Value := "parameter64Value_example" // string | Parameter value (optional)
    parameter65Name := "parameter65Name_example" // string | Parameter name (optional)
    parameter65Value := "parameter65Value_example" // string | Parameter value (optional)
    parameter66Name := "parameter66Name_example" // string | Parameter name (optional)
    parameter66Value := "parameter66Value_example" // string | Parameter value (optional)
    parameter67Name := "parameter67Name_example" // string | Parameter name (optional)
    parameter67Value := "parameter67Value_example" // string | Parameter value (optional)
    parameter68Name := "parameter68Name_example" // string | Parameter name (optional)
    parameter68Value := "parameter68Value_example" // string | Parameter value (optional)
    parameter69Name := "parameter69Name_example" // string | Parameter name (optional)
    parameter69Value := "parameter69Value_example" // string | Parameter value (optional)
    parameter7Name := "parameter7Name_example" // string | Parameter name (optional)
    parameter7Value := "parameter7Value_example" // string | Parameter value (optional)
    parameter70Name := "parameter70Name_example" // string | Parameter name (optional)
    parameter70Value := "parameter70Value_example" // string | Parameter value (optional)
    parameter71Name := "parameter71Name_example" // string | Parameter name (optional)
    parameter71Value := "parameter71Value_example" // string | Parameter value (optional)
    parameter72Name := "parameter72Name_example" // string | Parameter name (optional)
    parameter72Value := "parameter72Value_example" // string | Parameter value (optional)
    parameter73Name := "parameter73Name_example" // string | Parameter name (optional)
    parameter73Value := "parameter73Value_example" // string | Parameter value (optional)
    parameter74Name := "parameter74Name_example" // string | Parameter name (optional)
    parameter74Value := "parameter74Value_example" // string | Parameter value (optional)
    parameter75Name := "parameter75Name_example" // string | Parameter name (optional)
    parameter75Value := "parameter75Value_example" // string | Parameter value (optional)
    parameter76Name := "parameter76Name_example" // string | Parameter name (optional)
    parameter76Value := "parameter76Value_example" // string | Parameter value (optional)
    parameter77Name := "parameter77Name_example" // string | Parameter name (optional)
    parameter77Value := "parameter77Value_example" // string | Parameter value (optional)
    parameter78Name := "parameter78Name_example" // string | Parameter name (optional)
    parameter78Value := "parameter78Value_example" // string | Parameter value (optional)
    parameter79Name := "parameter79Name_example" // string | Parameter name (optional)
    parameter79Value := "parameter79Value_example" // string | Parameter value (optional)
    parameter8Name := "parameter8Name_example" // string | Parameter name (optional)
    parameter8Value := "parameter8Value_example" // string | Parameter value (optional)
    parameter80Name := "parameter80Name_example" // string | Parameter name (optional)
    parameter80Value := "parameter80Value_example" // string | Parameter value (optional)
    parameter81Name := "parameter81Name_example" // string | Parameter name (optional)
    parameter81Value := "parameter81Value_example" // string | Parameter value (optional)
    parameter82Name := "parameter82Name_example" // string | Parameter name (optional)
    parameter82Value := "parameter82Value_example" // string | Parameter value (optional)
    parameter83Name := "parameter83Name_example" // string | Parameter name (optional)
    parameter83Value := "parameter83Value_example" // string | Parameter value (optional)
    parameter84Name := "parameter84Name_example" // string | Parameter name (optional)
    parameter84Value := "parameter84Value_example" // string | Parameter value (optional)
    parameter85Name := "parameter85Name_example" // string | Parameter name (optional)
    parameter85Value := "parameter85Value_example" // string | Parameter value (optional)
    parameter86Name := "parameter86Name_example" // string | Parameter name (optional)
    parameter86Value := "parameter86Value_example" // string | Parameter value (optional)
    parameter87Name := "parameter87Name_example" // string | Parameter name (optional)
    parameter87Value := "parameter87Value_example" // string | Parameter value (optional)
    parameter88Name := "parameter88Name_example" // string | Parameter name (optional)
    parameter88Value := "parameter88Value_example" // string | Parameter value (optional)
    parameter89Name := "parameter89Name_example" // string | Parameter name (optional)
    parameter89Value := "parameter89Value_example" // string | Parameter value (optional)
    parameter9Name := "parameter9Name_example" // string | Parameter name (optional)
    parameter9Value := "parameter9Value_example" // string | Parameter value (optional)
    parameter90Name := "parameter90Name_example" // string | Parameter name (optional)
    parameter90Value := "parameter90Value_example" // string | Parameter value (optional)
    parameter91Name := "parameter91Name_example" // string | Parameter name (optional)
    parameter91Value := "parameter91Value_example" // string | Parameter value (optional)
    parameter92Name := "parameter92Name_example" // string | Parameter name (optional)
    parameter92Value := "parameter92Value_example" // string | Parameter value (optional)
    parameter93Name := "parameter93Name_example" // string | Parameter name (optional)
    parameter93Value := "parameter93Value_example" // string | Parameter value (optional)
    parameter94Name := "parameter94Name_example" // string | Parameter name (optional)
    parameter94Value := "parameter94Value_example" // string | Parameter value (optional)
    parameter95Name := "parameter95Name_example" // string | Parameter name (optional)
    parameter95Value := "parameter95Value_example" // string | Parameter value (optional)
    parameter96Name := "parameter96Name_example" // string | Parameter name (optional)
    parameter96Value := "parameter96Value_example" // string | Parameter value (optional)
    parameter97Name := "parameter97Name_example" // string | Parameter name (optional)
    parameter97Value := "parameter97Value_example" // string | Parameter value (optional)
    parameter98Name := "parameter98Name_example" // string | Parameter name (optional)
    parameter98Value := "parameter98Value_example" // string | Parameter value (optional)
    parameter99Name := "parameter99Name_example" // string | Parameter name (optional)
    parameter99Value := "parameter99Value_example" // string | Parameter value (optional)
    statusCallback := "statusCallback_example" // string | Absolute URL of the status callback. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The http method for the status_callback (one of GET, POST). (optional)
    track := "track_example" // string | One of `inbound_track`, `outbound_track`, `both_tracks`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSiprec(context.Background(), accountSid, callSid).ConnectorName(connectorName).Name(name).Parameter1Name(parameter1Name).Parameter1Value(parameter1Value).Parameter10Name(parameter10Name).Parameter10Value(parameter10Value).Parameter11Name(parameter11Name).Parameter11Value(parameter11Value).Parameter12Name(parameter12Name).Parameter12Value(parameter12Value).Parameter13Name(parameter13Name).Parameter13Value(parameter13Value).Parameter14Name(parameter14Name).Parameter14Value(parameter14Value).Parameter15Name(parameter15Name).Parameter15Value(parameter15Value).Parameter16Name(parameter16Name).Parameter16Value(parameter16Value).Parameter17Name(parameter17Name).Parameter17Value(parameter17Value).Parameter18Name(parameter18Name).Parameter18Value(parameter18Value).Parameter19Name(parameter19Name).Parameter19Value(parameter19Value).Parameter2Name(parameter2Name).Parameter2Value(parameter2Value).Parameter20Name(parameter20Name).Parameter20Value(parameter20Value).Parameter21Name(parameter21Name).Parameter21Value(parameter21Value).Parameter22Name(parameter22Name).Parameter22Value(parameter22Value).Parameter23Name(parameter23Name).Parameter23Value(parameter23Value).Parameter24Name(parameter24Name).Parameter24Value(parameter24Value).Parameter25Name(parameter25Name).Parameter25Value(parameter25Value).Parameter26Name(parameter26Name).Parameter26Value(parameter26Value).Parameter27Name(parameter27Name).Parameter27Value(parameter27Value).Parameter28Name(parameter28Name).Parameter28Value(parameter28Value).Parameter29Name(parameter29Name).Parameter29Value(parameter29Value).Parameter3Name(parameter3Name).Parameter3Value(parameter3Value).Parameter30Name(parameter30Name).Parameter30Value(parameter30Value).Parameter31Name(parameter31Name).Parameter31Value(parameter31Value).Parameter32Name(parameter32Name).Parameter32Value(parameter32Value).Parameter33Name(parameter33Name).Parameter33Value(parameter33Value).Parameter34Name(parameter34Name).Parameter34Value(parameter34Value).Parameter35Name(parameter35Name).Parameter35Value(parameter35Value).Parameter36Name(parameter36Name).Parameter36Value(parameter36Value).Parameter37Name(parameter37Name).Parameter37Value(parameter37Value).Parameter38Name(parameter38Name).Parameter38Value(parameter38Value).Parameter39Name(parameter39Name).Parameter39Value(parameter39Value).Parameter4Name(parameter4Name).Parameter4Value(parameter4Value).Parameter40Name(parameter40Name).Parameter40Value(parameter40Value).Parameter41Name(parameter41Name).Parameter41Value(parameter41Value).Parameter42Name(parameter42Name).Parameter42Value(parameter42Value).Parameter43Name(parameter43Name).Parameter43Value(parameter43Value).Parameter44Name(parameter44Name).Parameter44Value(parameter44Value).Parameter45Name(parameter45Name).Parameter45Value(parameter45Value).Parameter46Name(parameter46Name).Parameter46Value(parameter46Value).Parameter47Name(parameter47Name).Parameter47Value(parameter47Value).Parameter48Name(parameter48Name).Parameter48Value(parameter48Value).Parameter49Name(parameter49Name).Parameter49Value(parameter49Value).Parameter5Name(parameter5Name).Parameter5Value(parameter5Value).Parameter50Name(parameter50Name).Parameter50Value(parameter50Value).Parameter51Name(parameter51Name).Parameter51Value(parameter51Value).Parameter52Name(parameter52Name).Parameter52Value(parameter52Value).Parameter53Name(parameter53Name).Parameter53Value(parameter53Value).Parameter54Name(parameter54Name).Parameter54Value(parameter54Value).Parameter55Name(parameter55Name).Parameter55Value(parameter55Value).Parameter56Name(parameter56Name).Parameter56Value(parameter56Value).Parameter57Name(parameter57Name).Parameter57Value(parameter57Value).Parameter58Name(parameter58Name).Parameter58Value(parameter58Value).Parameter59Name(parameter59Name).Parameter59Value(parameter59Value).Parameter6Name(parameter6Name).Parameter6Value(parameter6Value).Parameter60Name(parameter60Name).Parameter60Value(parameter60Value).Parameter61Name(parameter61Name).Parameter61Value(parameter61Value).Parameter62Name(parameter62Name).Parameter62Value(parameter62Value).Parameter63Name(parameter63Name).Parameter63Value(parameter63Value).Parameter64Name(parameter64Name).Parameter64Value(parameter64Value).Parameter65Name(parameter65Name).Parameter65Value(parameter65Value).Parameter66Name(parameter66Name).Parameter66Value(parameter66Value).Parameter67Name(parameter67Name).Parameter67Value(parameter67Value).Parameter68Name(parameter68Name).Parameter68Value(parameter68Value).Parameter69Name(parameter69Name).Parameter69Value(parameter69Value).Parameter7Name(parameter7Name).Parameter7Value(parameter7Value).Parameter70Name(parameter70Name).Parameter70Value(parameter70Value).Parameter71Name(parameter71Name).Parameter71Value(parameter71Value).Parameter72Name(parameter72Name).Parameter72Value(parameter72Value).Parameter73Name(parameter73Name).Parameter73Value(parameter73Value).Parameter74Name(parameter74Name).Parameter74Value(parameter74Value).Parameter75Name(parameter75Name).Parameter75Value(parameter75Value).Parameter76Name(parameter76Name).Parameter76Value(parameter76Value).Parameter77Name(parameter77Name).Parameter77Value(parameter77Value).Parameter78Name(parameter78Name).Parameter78Value(parameter78Value).Parameter79Name(parameter79Name).Parameter79Value(parameter79Value).Parameter8Name(parameter8Name).Parameter8Value(parameter8Value).Parameter80Name(parameter80Name).Parameter80Value(parameter80Value).Parameter81Name(parameter81Name).Parameter81Value(parameter81Value).Parameter82Name(parameter82Name).Parameter82Value(parameter82Value).Parameter83Name(parameter83Name).Parameter83Value(parameter83Value).Parameter84Name(parameter84Name).Parameter84Value(parameter84Value).Parameter85Name(parameter85Name).Parameter85Value(parameter85Value).Parameter86Name(parameter86Name).Parameter86Value(parameter86Value).Parameter87Name(parameter87Name).Parameter87Value(parameter87Value).Parameter88Name(parameter88Name).Parameter88Value(parameter88Value).Parameter89Name(parameter89Name).Parameter89Value(parameter89Value).Parameter9Name(parameter9Name).Parameter9Value(parameter9Value).Parameter90Name(parameter90Name).Parameter90Value(parameter90Value).Parameter91Name(parameter91Name).Parameter91Value(parameter91Value).Parameter92Name(parameter92Name).Parameter92Value(parameter92Value).Parameter93Name(parameter93Name).Parameter93Value(parameter93Value).Parameter94Name(parameter94Name).Parameter94Value(parameter94Value).Parameter95Name(parameter95Name).Parameter95Value(parameter95Value).Parameter96Name(parameter96Name).Parameter96Value(parameter96Value).Parameter97Name(parameter97Name).Parameter97Value(parameter97Value).Parameter98Name(parameter98Name).Parameter98Value(parameter98Value).Parameter99Name(parameter99Name).Parameter99Value(parameter99Value).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSiprec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiprec`: ApiV2010AccountCallSiprec
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSiprec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Siprec resource. | 
**callSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Siprec resource is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiprecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectorName** | **string** | Unique name used when configuring the connector via Marketplace Add-on. | 
 **name** | **string** | The user-specified name of this Siprec, if one was given when the Siprec was created. This may be used to stop the Siprec. | 
 **parameter1Name** | **string** | Parameter name | 
 **parameter1Value** | **string** | Parameter value | 
 **parameter10Name** | **string** | Parameter name | 
 **parameter10Value** | **string** | Parameter value | 
 **parameter11Name** | **string** | Parameter name | 
 **parameter11Value** | **string** | Parameter value | 
 **parameter12Name** | **string** | Parameter name | 
 **parameter12Value** | **string** | Parameter value | 
 **parameter13Name** | **string** | Parameter name | 
 **parameter13Value** | **string** | Parameter value | 
 **parameter14Name** | **string** | Parameter name | 
 **parameter14Value** | **string** | Parameter value | 
 **parameter15Name** | **string** | Parameter name | 
 **parameter15Value** | **string** | Parameter value | 
 **parameter16Name** | **string** | Parameter name | 
 **parameter16Value** | **string** | Parameter value | 
 **parameter17Name** | **string** | Parameter name | 
 **parameter17Value** | **string** | Parameter value | 
 **parameter18Name** | **string** | Parameter name | 
 **parameter18Value** | **string** | Parameter value | 
 **parameter19Name** | **string** | Parameter name | 
 **parameter19Value** | **string** | Parameter value | 
 **parameter2Name** | **string** | Parameter name | 
 **parameter2Value** | **string** | Parameter value | 
 **parameter20Name** | **string** | Parameter name | 
 **parameter20Value** | **string** | Parameter value | 
 **parameter21Name** | **string** | Parameter name | 
 **parameter21Value** | **string** | Parameter value | 
 **parameter22Name** | **string** | Parameter name | 
 **parameter22Value** | **string** | Parameter value | 
 **parameter23Name** | **string** | Parameter name | 
 **parameter23Value** | **string** | Parameter value | 
 **parameter24Name** | **string** | Parameter name | 
 **parameter24Value** | **string** | Parameter value | 
 **parameter25Name** | **string** | Parameter name | 
 **parameter25Value** | **string** | Parameter value | 
 **parameter26Name** | **string** | Parameter name | 
 **parameter26Value** | **string** | Parameter value | 
 **parameter27Name** | **string** | Parameter name | 
 **parameter27Value** | **string** | Parameter value | 
 **parameter28Name** | **string** | Parameter name | 
 **parameter28Value** | **string** | Parameter value | 
 **parameter29Name** | **string** | Parameter name | 
 **parameter29Value** | **string** | Parameter value | 
 **parameter3Name** | **string** | Parameter name | 
 **parameter3Value** | **string** | Parameter value | 
 **parameter30Name** | **string** | Parameter name | 
 **parameter30Value** | **string** | Parameter value | 
 **parameter31Name** | **string** | Parameter name | 
 **parameter31Value** | **string** | Parameter value | 
 **parameter32Name** | **string** | Parameter name | 
 **parameter32Value** | **string** | Parameter value | 
 **parameter33Name** | **string** | Parameter name | 
 **parameter33Value** | **string** | Parameter value | 
 **parameter34Name** | **string** | Parameter name | 
 **parameter34Value** | **string** | Parameter value | 
 **parameter35Name** | **string** | Parameter name | 
 **parameter35Value** | **string** | Parameter value | 
 **parameter36Name** | **string** | Parameter name | 
 **parameter36Value** | **string** | Parameter value | 
 **parameter37Name** | **string** | Parameter name | 
 **parameter37Value** | **string** | Parameter value | 
 **parameter38Name** | **string** | Parameter name | 
 **parameter38Value** | **string** | Parameter value | 
 **parameter39Name** | **string** | Parameter name | 
 **parameter39Value** | **string** | Parameter value | 
 **parameter4Name** | **string** | Parameter name | 
 **parameter4Value** | **string** | Parameter value | 
 **parameter40Name** | **string** | Parameter name | 
 **parameter40Value** | **string** | Parameter value | 
 **parameter41Name** | **string** | Parameter name | 
 **parameter41Value** | **string** | Parameter value | 
 **parameter42Name** | **string** | Parameter name | 
 **parameter42Value** | **string** | Parameter value | 
 **parameter43Name** | **string** | Parameter name | 
 **parameter43Value** | **string** | Parameter value | 
 **parameter44Name** | **string** | Parameter name | 
 **parameter44Value** | **string** | Parameter value | 
 **parameter45Name** | **string** | Parameter name | 
 **parameter45Value** | **string** | Parameter value | 
 **parameter46Name** | **string** | Parameter name | 
 **parameter46Value** | **string** | Parameter value | 
 **parameter47Name** | **string** | Parameter name | 
 **parameter47Value** | **string** | Parameter value | 
 **parameter48Name** | **string** | Parameter name | 
 **parameter48Value** | **string** | Parameter value | 
 **parameter49Name** | **string** | Parameter name | 
 **parameter49Value** | **string** | Parameter value | 
 **parameter5Name** | **string** | Parameter name | 
 **parameter5Value** | **string** | Parameter value | 
 **parameter50Name** | **string** | Parameter name | 
 **parameter50Value** | **string** | Parameter value | 
 **parameter51Name** | **string** | Parameter name | 
 **parameter51Value** | **string** | Parameter value | 
 **parameter52Name** | **string** | Parameter name | 
 **parameter52Value** | **string** | Parameter value | 
 **parameter53Name** | **string** | Parameter name | 
 **parameter53Value** | **string** | Parameter value | 
 **parameter54Name** | **string** | Parameter name | 
 **parameter54Value** | **string** | Parameter value | 
 **parameter55Name** | **string** | Parameter name | 
 **parameter55Value** | **string** | Parameter value | 
 **parameter56Name** | **string** | Parameter name | 
 **parameter56Value** | **string** | Parameter value | 
 **parameter57Name** | **string** | Parameter name | 
 **parameter57Value** | **string** | Parameter value | 
 **parameter58Name** | **string** | Parameter name | 
 **parameter58Value** | **string** | Parameter value | 
 **parameter59Name** | **string** | Parameter name | 
 **parameter59Value** | **string** | Parameter value | 
 **parameter6Name** | **string** | Parameter name | 
 **parameter6Value** | **string** | Parameter value | 
 **parameter60Name** | **string** | Parameter name | 
 **parameter60Value** | **string** | Parameter value | 
 **parameter61Name** | **string** | Parameter name | 
 **parameter61Value** | **string** | Parameter value | 
 **parameter62Name** | **string** | Parameter name | 
 **parameter62Value** | **string** | Parameter value | 
 **parameter63Name** | **string** | Parameter name | 
 **parameter63Value** | **string** | Parameter value | 
 **parameter64Name** | **string** | Parameter name | 
 **parameter64Value** | **string** | Parameter value | 
 **parameter65Name** | **string** | Parameter name | 
 **parameter65Value** | **string** | Parameter value | 
 **parameter66Name** | **string** | Parameter name | 
 **parameter66Value** | **string** | Parameter value | 
 **parameter67Name** | **string** | Parameter name | 
 **parameter67Value** | **string** | Parameter value | 
 **parameter68Name** | **string** | Parameter name | 
 **parameter68Value** | **string** | Parameter value | 
 **parameter69Name** | **string** | Parameter name | 
 **parameter69Value** | **string** | Parameter value | 
 **parameter7Name** | **string** | Parameter name | 
 **parameter7Value** | **string** | Parameter value | 
 **parameter70Name** | **string** | Parameter name | 
 **parameter70Value** | **string** | Parameter value | 
 **parameter71Name** | **string** | Parameter name | 
 **parameter71Value** | **string** | Parameter value | 
 **parameter72Name** | **string** | Parameter name | 
 **parameter72Value** | **string** | Parameter value | 
 **parameter73Name** | **string** | Parameter name | 
 **parameter73Value** | **string** | Parameter value | 
 **parameter74Name** | **string** | Parameter name | 
 **parameter74Value** | **string** | Parameter value | 
 **parameter75Name** | **string** | Parameter name | 
 **parameter75Value** | **string** | Parameter value | 
 **parameter76Name** | **string** | Parameter name | 
 **parameter76Value** | **string** | Parameter value | 
 **parameter77Name** | **string** | Parameter name | 
 **parameter77Value** | **string** | Parameter value | 
 **parameter78Name** | **string** | Parameter name | 
 **parameter78Value** | **string** | Parameter value | 
 **parameter79Name** | **string** | Parameter name | 
 **parameter79Value** | **string** | Parameter value | 
 **parameter8Name** | **string** | Parameter name | 
 **parameter8Value** | **string** | Parameter value | 
 **parameter80Name** | **string** | Parameter name | 
 **parameter80Value** | **string** | Parameter value | 
 **parameter81Name** | **string** | Parameter name | 
 **parameter81Value** | **string** | Parameter value | 
 **parameter82Name** | **string** | Parameter name | 
 **parameter82Value** | **string** | Parameter value | 
 **parameter83Name** | **string** | Parameter name | 
 **parameter83Value** | **string** | Parameter value | 
 **parameter84Name** | **string** | Parameter name | 
 **parameter84Value** | **string** | Parameter value | 
 **parameter85Name** | **string** | Parameter name | 
 **parameter85Value** | **string** | Parameter value | 
 **parameter86Name** | **string** | Parameter name | 
 **parameter86Value** | **string** | Parameter value | 
 **parameter87Name** | **string** | Parameter name | 
 **parameter87Value** | **string** | Parameter value | 
 **parameter88Name** | **string** | Parameter name | 
 **parameter88Value** | **string** | Parameter value | 
 **parameter89Name** | **string** | Parameter name | 
 **parameter89Value** | **string** | Parameter value | 
 **parameter9Name** | **string** | Parameter name | 
 **parameter9Value** | **string** | Parameter value | 
 **parameter90Name** | **string** | Parameter name | 
 **parameter90Value** | **string** | Parameter value | 
 **parameter91Name** | **string** | Parameter name | 
 **parameter91Value** | **string** | Parameter value | 
 **parameter92Name** | **string** | Parameter name | 
 **parameter92Value** | **string** | Parameter value | 
 **parameter93Name** | **string** | Parameter name | 
 **parameter93Value** | **string** | Parameter value | 
 **parameter94Name** | **string** | Parameter name | 
 **parameter94Value** | **string** | Parameter value | 
 **parameter95Name** | **string** | Parameter name | 
 **parameter95Value** | **string** | Parameter value | 
 **parameter96Name** | **string** | Parameter name | 
 **parameter96Value** | **string** | Parameter value | 
 **parameter97Name** | **string** | Parameter name | 
 **parameter97Value** | **string** | Parameter value | 
 **parameter98Name** | **string** | Parameter name | 
 **parameter98Value** | **string** | Parameter value | 
 **parameter99Name** | **string** | Parameter name | 
 **parameter99Value** | **string** | Parameter value | 
 **statusCallback** | **string** | Absolute URL of the status callback. | 
 **statusCallbackMethod** | **string** | The http method for the status_callback (one of GET, POST). | 
 **track** | **string** | One of &#x60;inbound_track&#x60;, &#x60;outbound_track&#x60;, &#x60;both_tracks&#x60;. | 

### Return type

[**ApiV2010AccountCallSiprec**](ApiV2010AccountCallSiprec.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStream

> ApiV2010AccountCallStream CreateStream(ctx, accountSid, callSid).Url(url).Name(name).Parameter1Name(parameter1Name).Parameter1Value(parameter1Value).Parameter10Name(parameter10Name).Parameter10Value(parameter10Value).Parameter11Name(parameter11Name).Parameter11Value(parameter11Value).Parameter12Name(parameter12Name).Parameter12Value(parameter12Value).Parameter13Name(parameter13Name).Parameter13Value(parameter13Value).Parameter14Name(parameter14Name).Parameter14Value(parameter14Value).Parameter15Name(parameter15Name).Parameter15Value(parameter15Value).Parameter16Name(parameter16Name).Parameter16Value(parameter16Value).Parameter17Name(parameter17Name).Parameter17Value(parameter17Value).Parameter18Name(parameter18Name).Parameter18Value(parameter18Value).Parameter19Name(parameter19Name).Parameter19Value(parameter19Value).Parameter2Name(parameter2Name).Parameter2Value(parameter2Value).Parameter20Name(parameter20Name).Parameter20Value(parameter20Value).Parameter21Name(parameter21Name).Parameter21Value(parameter21Value).Parameter22Name(parameter22Name).Parameter22Value(parameter22Value).Parameter23Name(parameter23Name).Parameter23Value(parameter23Value).Parameter24Name(parameter24Name).Parameter24Value(parameter24Value).Parameter25Name(parameter25Name).Parameter25Value(parameter25Value).Parameter26Name(parameter26Name).Parameter26Value(parameter26Value).Parameter27Name(parameter27Name).Parameter27Value(parameter27Value).Parameter28Name(parameter28Name).Parameter28Value(parameter28Value).Parameter29Name(parameter29Name).Parameter29Value(parameter29Value).Parameter3Name(parameter3Name).Parameter3Value(parameter3Value).Parameter30Name(parameter30Name).Parameter30Value(parameter30Value).Parameter31Name(parameter31Name).Parameter31Value(parameter31Value).Parameter32Name(parameter32Name).Parameter32Value(parameter32Value).Parameter33Name(parameter33Name).Parameter33Value(parameter33Value).Parameter34Name(parameter34Name).Parameter34Value(parameter34Value).Parameter35Name(parameter35Name).Parameter35Value(parameter35Value).Parameter36Name(parameter36Name).Parameter36Value(parameter36Value).Parameter37Name(parameter37Name).Parameter37Value(parameter37Value).Parameter38Name(parameter38Name).Parameter38Value(parameter38Value).Parameter39Name(parameter39Name).Parameter39Value(parameter39Value).Parameter4Name(parameter4Name).Parameter4Value(parameter4Value).Parameter40Name(parameter40Name).Parameter40Value(parameter40Value).Parameter41Name(parameter41Name).Parameter41Value(parameter41Value).Parameter42Name(parameter42Name).Parameter42Value(parameter42Value).Parameter43Name(parameter43Name).Parameter43Value(parameter43Value).Parameter44Name(parameter44Name).Parameter44Value(parameter44Value).Parameter45Name(parameter45Name).Parameter45Value(parameter45Value).Parameter46Name(parameter46Name).Parameter46Value(parameter46Value).Parameter47Name(parameter47Name).Parameter47Value(parameter47Value).Parameter48Name(parameter48Name).Parameter48Value(parameter48Value).Parameter49Name(parameter49Name).Parameter49Value(parameter49Value).Parameter5Name(parameter5Name).Parameter5Value(parameter5Value).Parameter50Name(parameter50Name).Parameter50Value(parameter50Value).Parameter51Name(parameter51Name).Parameter51Value(parameter51Value).Parameter52Name(parameter52Name).Parameter52Value(parameter52Value).Parameter53Name(parameter53Name).Parameter53Value(parameter53Value).Parameter54Name(parameter54Name).Parameter54Value(parameter54Value).Parameter55Name(parameter55Name).Parameter55Value(parameter55Value).Parameter56Name(parameter56Name).Parameter56Value(parameter56Value).Parameter57Name(parameter57Name).Parameter57Value(parameter57Value).Parameter58Name(parameter58Name).Parameter58Value(parameter58Value).Parameter59Name(parameter59Name).Parameter59Value(parameter59Value).Parameter6Name(parameter6Name).Parameter6Value(parameter6Value).Parameter60Name(parameter60Name).Parameter60Value(parameter60Value).Parameter61Name(parameter61Name).Parameter61Value(parameter61Value).Parameter62Name(parameter62Name).Parameter62Value(parameter62Value).Parameter63Name(parameter63Name).Parameter63Value(parameter63Value).Parameter64Name(parameter64Name).Parameter64Value(parameter64Value).Parameter65Name(parameter65Name).Parameter65Value(parameter65Value).Parameter66Name(parameter66Name).Parameter66Value(parameter66Value).Parameter67Name(parameter67Name).Parameter67Value(parameter67Value).Parameter68Name(parameter68Name).Parameter68Value(parameter68Value).Parameter69Name(parameter69Name).Parameter69Value(parameter69Value).Parameter7Name(parameter7Name).Parameter7Value(parameter7Value).Parameter70Name(parameter70Name).Parameter70Value(parameter70Value).Parameter71Name(parameter71Name).Parameter71Value(parameter71Value).Parameter72Name(parameter72Name).Parameter72Value(parameter72Value).Parameter73Name(parameter73Name).Parameter73Value(parameter73Value).Parameter74Name(parameter74Name).Parameter74Value(parameter74Value).Parameter75Name(parameter75Name).Parameter75Value(parameter75Value).Parameter76Name(parameter76Name).Parameter76Value(parameter76Value).Parameter77Name(parameter77Name).Parameter77Value(parameter77Value).Parameter78Name(parameter78Name).Parameter78Value(parameter78Value).Parameter79Name(parameter79Name).Parameter79Value(parameter79Value).Parameter8Name(parameter8Name).Parameter8Value(parameter8Value).Parameter80Name(parameter80Name).Parameter80Value(parameter80Value).Parameter81Name(parameter81Name).Parameter81Value(parameter81Value).Parameter82Name(parameter82Name).Parameter82Value(parameter82Value).Parameter83Name(parameter83Name).Parameter83Value(parameter83Value).Parameter84Name(parameter84Name).Parameter84Value(parameter84Value).Parameter85Name(parameter85Name).Parameter85Value(parameter85Value).Parameter86Name(parameter86Name).Parameter86Value(parameter86Value).Parameter87Name(parameter87Name).Parameter87Value(parameter87Value).Parameter88Name(parameter88Name).Parameter88Value(parameter88Value).Parameter89Name(parameter89Name).Parameter89Value(parameter89Value).Parameter9Name(parameter9Name).Parameter9Value(parameter9Value).Parameter90Name(parameter90Name).Parameter90Value(parameter90Value).Parameter91Name(parameter91Name).Parameter91Value(parameter91Value).Parameter92Name(parameter92Name).Parameter92Value(parameter92Value).Parameter93Name(parameter93Name).Parameter93Value(parameter93Value).Parameter94Name(parameter94Name).Parameter94Value(parameter94Value).Parameter95Name(parameter95Name).Parameter95Value(parameter95Value).Parameter96Name(parameter96Name).Parameter96Value(parameter96Value).Parameter97Name(parameter97Name).Parameter97Value(parameter97Value).Parameter98Name(parameter98Name).Parameter98Value(parameter98Value).Parameter99Name(parameter99Name).Parameter99Value(parameter99Value).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Stream resource.
    callSid := "callSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Stream resource is associated with.
    url := "url_example" // string | Relative or absolute url where WebSocket connection will be established.
    name := "name_example" // string | The user-specified name of this Stream, if one was given when the Stream was created. This may be used to stop the Stream. (optional)
    parameter1Name := "parameter1Name_example" // string | Parameter name (optional)
    parameter1Value := "parameter1Value_example" // string | Parameter value (optional)
    parameter10Name := "parameter10Name_example" // string | Parameter name (optional)
    parameter10Value := "parameter10Value_example" // string | Parameter value (optional)
    parameter11Name := "parameter11Name_example" // string | Parameter name (optional)
    parameter11Value := "parameter11Value_example" // string | Parameter value (optional)
    parameter12Name := "parameter12Name_example" // string | Parameter name (optional)
    parameter12Value := "parameter12Value_example" // string | Parameter value (optional)
    parameter13Name := "parameter13Name_example" // string | Parameter name (optional)
    parameter13Value := "parameter13Value_example" // string | Parameter value (optional)
    parameter14Name := "parameter14Name_example" // string | Parameter name (optional)
    parameter14Value := "parameter14Value_example" // string | Parameter value (optional)
    parameter15Name := "parameter15Name_example" // string | Parameter name (optional)
    parameter15Value := "parameter15Value_example" // string | Parameter value (optional)
    parameter16Name := "parameter16Name_example" // string | Parameter name (optional)
    parameter16Value := "parameter16Value_example" // string | Parameter value (optional)
    parameter17Name := "parameter17Name_example" // string | Parameter name (optional)
    parameter17Value := "parameter17Value_example" // string | Parameter value (optional)
    parameter18Name := "parameter18Name_example" // string | Parameter name (optional)
    parameter18Value := "parameter18Value_example" // string | Parameter value (optional)
    parameter19Name := "parameter19Name_example" // string | Parameter name (optional)
    parameter19Value := "parameter19Value_example" // string | Parameter value (optional)
    parameter2Name := "parameter2Name_example" // string | Parameter name (optional)
    parameter2Value := "parameter2Value_example" // string | Parameter value (optional)
    parameter20Name := "parameter20Name_example" // string | Parameter name (optional)
    parameter20Value := "parameter20Value_example" // string | Parameter value (optional)
    parameter21Name := "parameter21Name_example" // string | Parameter name (optional)
    parameter21Value := "parameter21Value_example" // string | Parameter value (optional)
    parameter22Name := "parameter22Name_example" // string | Parameter name (optional)
    parameter22Value := "parameter22Value_example" // string | Parameter value (optional)
    parameter23Name := "parameter23Name_example" // string | Parameter name (optional)
    parameter23Value := "parameter23Value_example" // string | Parameter value (optional)
    parameter24Name := "parameter24Name_example" // string | Parameter name (optional)
    parameter24Value := "parameter24Value_example" // string | Parameter value (optional)
    parameter25Name := "parameter25Name_example" // string | Parameter name (optional)
    parameter25Value := "parameter25Value_example" // string | Parameter value (optional)
    parameter26Name := "parameter26Name_example" // string | Parameter name (optional)
    parameter26Value := "parameter26Value_example" // string | Parameter value (optional)
    parameter27Name := "parameter27Name_example" // string | Parameter name (optional)
    parameter27Value := "parameter27Value_example" // string | Parameter value (optional)
    parameter28Name := "parameter28Name_example" // string | Parameter name (optional)
    parameter28Value := "parameter28Value_example" // string | Parameter value (optional)
    parameter29Name := "parameter29Name_example" // string | Parameter name (optional)
    parameter29Value := "parameter29Value_example" // string | Parameter value (optional)
    parameter3Name := "parameter3Name_example" // string | Parameter name (optional)
    parameter3Value := "parameter3Value_example" // string | Parameter value (optional)
    parameter30Name := "parameter30Name_example" // string | Parameter name (optional)
    parameter30Value := "parameter30Value_example" // string | Parameter value (optional)
    parameter31Name := "parameter31Name_example" // string | Parameter name (optional)
    parameter31Value := "parameter31Value_example" // string | Parameter value (optional)
    parameter32Name := "parameter32Name_example" // string | Parameter name (optional)
    parameter32Value := "parameter32Value_example" // string | Parameter value (optional)
    parameter33Name := "parameter33Name_example" // string | Parameter name (optional)
    parameter33Value := "parameter33Value_example" // string | Parameter value (optional)
    parameter34Name := "parameter34Name_example" // string | Parameter name (optional)
    parameter34Value := "parameter34Value_example" // string | Parameter value (optional)
    parameter35Name := "parameter35Name_example" // string | Parameter name (optional)
    parameter35Value := "parameter35Value_example" // string | Parameter value (optional)
    parameter36Name := "parameter36Name_example" // string | Parameter name (optional)
    parameter36Value := "parameter36Value_example" // string | Parameter value (optional)
    parameter37Name := "parameter37Name_example" // string | Parameter name (optional)
    parameter37Value := "parameter37Value_example" // string | Parameter value (optional)
    parameter38Name := "parameter38Name_example" // string | Parameter name (optional)
    parameter38Value := "parameter38Value_example" // string | Parameter value (optional)
    parameter39Name := "parameter39Name_example" // string | Parameter name (optional)
    parameter39Value := "parameter39Value_example" // string | Parameter value (optional)
    parameter4Name := "parameter4Name_example" // string | Parameter name (optional)
    parameter4Value := "parameter4Value_example" // string | Parameter value (optional)
    parameter40Name := "parameter40Name_example" // string | Parameter name (optional)
    parameter40Value := "parameter40Value_example" // string | Parameter value (optional)
    parameter41Name := "parameter41Name_example" // string | Parameter name (optional)
    parameter41Value := "parameter41Value_example" // string | Parameter value (optional)
    parameter42Name := "parameter42Name_example" // string | Parameter name (optional)
    parameter42Value := "parameter42Value_example" // string | Parameter value (optional)
    parameter43Name := "parameter43Name_example" // string | Parameter name (optional)
    parameter43Value := "parameter43Value_example" // string | Parameter value (optional)
    parameter44Name := "parameter44Name_example" // string | Parameter name (optional)
    parameter44Value := "parameter44Value_example" // string | Parameter value (optional)
    parameter45Name := "parameter45Name_example" // string | Parameter name (optional)
    parameter45Value := "parameter45Value_example" // string | Parameter value (optional)
    parameter46Name := "parameter46Name_example" // string | Parameter name (optional)
    parameter46Value := "parameter46Value_example" // string | Parameter value (optional)
    parameter47Name := "parameter47Name_example" // string | Parameter name (optional)
    parameter47Value := "parameter47Value_example" // string | Parameter value (optional)
    parameter48Name := "parameter48Name_example" // string | Parameter name (optional)
    parameter48Value := "parameter48Value_example" // string | Parameter value (optional)
    parameter49Name := "parameter49Name_example" // string | Parameter name (optional)
    parameter49Value := "parameter49Value_example" // string | Parameter value (optional)
    parameter5Name := "parameter5Name_example" // string | Parameter name (optional)
    parameter5Value := "parameter5Value_example" // string | Parameter value (optional)
    parameter50Name := "parameter50Name_example" // string | Parameter name (optional)
    parameter50Value := "parameter50Value_example" // string | Parameter value (optional)
    parameter51Name := "parameter51Name_example" // string | Parameter name (optional)
    parameter51Value := "parameter51Value_example" // string | Parameter value (optional)
    parameter52Name := "parameter52Name_example" // string | Parameter name (optional)
    parameter52Value := "parameter52Value_example" // string | Parameter value (optional)
    parameter53Name := "parameter53Name_example" // string | Parameter name (optional)
    parameter53Value := "parameter53Value_example" // string | Parameter value (optional)
    parameter54Name := "parameter54Name_example" // string | Parameter name (optional)
    parameter54Value := "parameter54Value_example" // string | Parameter value (optional)
    parameter55Name := "parameter55Name_example" // string | Parameter name (optional)
    parameter55Value := "parameter55Value_example" // string | Parameter value (optional)
    parameter56Name := "parameter56Name_example" // string | Parameter name (optional)
    parameter56Value := "parameter56Value_example" // string | Parameter value (optional)
    parameter57Name := "parameter57Name_example" // string | Parameter name (optional)
    parameter57Value := "parameter57Value_example" // string | Parameter value (optional)
    parameter58Name := "parameter58Name_example" // string | Parameter name (optional)
    parameter58Value := "parameter58Value_example" // string | Parameter value (optional)
    parameter59Name := "parameter59Name_example" // string | Parameter name (optional)
    parameter59Value := "parameter59Value_example" // string | Parameter value (optional)
    parameter6Name := "parameter6Name_example" // string | Parameter name (optional)
    parameter6Value := "parameter6Value_example" // string | Parameter value (optional)
    parameter60Name := "parameter60Name_example" // string | Parameter name (optional)
    parameter60Value := "parameter60Value_example" // string | Parameter value (optional)
    parameter61Name := "parameter61Name_example" // string | Parameter name (optional)
    parameter61Value := "parameter61Value_example" // string | Parameter value (optional)
    parameter62Name := "parameter62Name_example" // string | Parameter name (optional)
    parameter62Value := "parameter62Value_example" // string | Parameter value (optional)
    parameter63Name := "parameter63Name_example" // string | Parameter name (optional)
    parameter63Value := "parameter63Value_example" // string | Parameter value (optional)
    parameter64Name := "parameter64Name_example" // string | Parameter name (optional)
    parameter64Value := "parameter64Value_example" // string | Parameter value (optional)
    parameter65Name := "parameter65Name_example" // string | Parameter name (optional)
    parameter65Value := "parameter65Value_example" // string | Parameter value (optional)
    parameter66Name := "parameter66Name_example" // string | Parameter name (optional)
    parameter66Value := "parameter66Value_example" // string | Parameter value (optional)
    parameter67Name := "parameter67Name_example" // string | Parameter name (optional)
    parameter67Value := "parameter67Value_example" // string | Parameter value (optional)
    parameter68Name := "parameter68Name_example" // string | Parameter name (optional)
    parameter68Value := "parameter68Value_example" // string | Parameter value (optional)
    parameter69Name := "parameter69Name_example" // string | Parameter name (optional)
    parameter69Value := "parameter69Value_example" // string | Parameter value (optional)
    parameter7Name := "parameter7Name_example" // string | Parameter name (optional)
    parameter7Value := "parameter7Value_example" // string | Parameter value (optional)
    parameter70Name := "parameter70Name_example" // string | Parameter name (optional)
    parameter70Value := "parameter70Value_example" // string | Parameter value (optional)
    parameter71Name := "parameter71Name_example" // string | Parameter name (optional)
    parameter71Value := "parameter71Value_example" // string | Parameter value (optional)
    parameter72Name := "parameter72Name_example" // string | Parameter name (optional)
    parameter72Value := "parameter72Value_example" // string | Parameter value (optional)
    parameter73Name := "parameter73Name_example" // string | Parameter name (optional)
    parameter73Value := "parameter73Value_example" // string | Parameter value (optional)
    parameter74Name := "parameter74Name_example" // string | Parameter name (optional)
    parameter74Value := "parameter74Value_example" // string | Parameter value (optional)
    parameter75Name := "parameter75Name_example" // string | Parameter name (optional)
    parameter75Value := "parameter75Value_example" // string | Parameter value (optional)
    parameter76Name := "parameter76Name_example" // string | Parameter name (optional)
    parameter76Value := "parameter76Value_example" // string | Parameter value (optional)
    parameter77Name := "parameter77Name_example" // string | Parameter name (optional)
    parameter77Value := "parameter77Value_example" // string | Parameter value (optional)
    parameter78Name := "parameter78Name_example" // string | Parameter name (optional)
    parameter78Value := "parameter78Value_example" // string | Parameter value (optional)
    parameter79Name := "parameter79Name_example" // string | Parameter name (optional)
    parameter79Value := "parameter79Value_example" // string | Parameter value (optional)
    parameter8Name := "parameter8Name_example" // string | Parameter name (optional)
    parameter8Value := "parameter8Value_example" // string | Parameter value (optional)
    parameter80Name := "parameter80Name_example" // string | Parameter name (optional)
    parameter80Value := "parameter80Value_example" // string | Parameter value (optional)
    parameter81Name := "parameter81Name_example" // string | Parameter name (optional)
    parameter81Value := "parameter81Value_example" // string | Parameter value (optional)
    parameter82Name := "parameter82Name_example" // string | Parameter name (optional)
    parameter82Value := "parameter82Value_example" // string | Parameter value (optional)
    parameter83Name := "parameter83Name_example" // string | Parameter name (optional)
    parameter83Value := "parameter83Value_example" // string | Parameter value (optional)
    parameter84Name := "parameter84Name_example" // string | Parameter name (optional)
    parameter84Value := "parameter84Value_example" // string | Parameter value (optional)
    parameter85Name := "parameter85Name_example" // string | Parameter name (optional)
    parameter85Value := "parameter85Value_example" // string | Parameter value (optional)
    parameter86Name := "parameter86Name_example" // string | Parameter name (optional)
    parameter86Value := "parameter86Value_example" // string | Parameter value (optional)
    parameter87Name := "parameter87Name_example" // string | Parameter name (optional)
    parameter87Value := "parameter87Value_example" // string | Parameter value (optional)
    parameter88Name := "parameter88Name_example" // string | Parameter name (optional)
    parameter88Value := "parameter88Value_example" // string | Parameter value (optional)
    parameter89Name := "parameter89Name_example" // string | Parameter name (optional)
    parameter89Value := "parameter89Value_example" // string | Parameter value (optional)
    parameter9Name := "parameter9Name_example" // string | Parameter name (optional)
    parameter9Value := "parameter9Value_example" // string | Parameter value (optional)
    parameter90Name := "parameter90Name_example" // string | Parameter name (optional)
    parameter90Value := "parameter90Value_example" // string | Parameter value (optional)
    parameter91Name := "parameter91Name_example" // string | Parameter name (optional)
    parameter91Value := "parameter91Value_example" // string | Parameter value (optional)
    parameter92Name := "parameter92Name_example" // string | Parameter name (optional)
    parameter92Value := "parameter92Value_example" // string | Parameter value (optional)
    parameter93Name := "parameter93Name_example" // string | Parameter name (optional)
    parameter93Value := "parameter93Value_example" // string | Parameter value (optional)
    parameter94Name := "parameter94Name_example" // string | Parameter name (optional)
    parameter94Value := "parameter94Value_example" // string | Parameter value (optional)
    parameter95Name := "parameter95Name_example" // string | Parameter name (optional)
    parameter95Value := "parameter95Value_example" // string | Parameter value (optional)
    parameter96Name := "parameter96Name_example" // string | Parameter name (optional)
    parameter96Value := "parameter96Value_example" // string | Parameter value (optional)
    parameter97Name := "parameter97Name_example" // string | Parameter name (optional)
    parameter97Value := "parameter97Value_example" // string | Parameter value (optional)
    parameter98Name := "parameter98Name_example" // string | Parameter name (optional)
    parameter98Value := "parameter98Value_example" // string | Parameter value (optional)
    parameter99Name := "parameter99Name_example" // string | Parameter name (optional)
    parameter99Value := "parameter99Value_example" // string | Parameter value (optional)
    statusCallback := "statusCallback_example" // string | Absolute URL of the status callback. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The http method for the status_callback (one of GET, POST). (optional)
    track := "track_example" // string | One of `inbound_track`, `outbound_track`, `both_tracks`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStream(context.Background(), accountSid, callSid).Url(url).Name(name).Parameter1Name(parameter1Name).Parameter1Value(parameter1Value).Parameter10Name(parameter10Name).Parameter10Value(parameter10Value).Parameter11Name(parameter11Name).Parameter11Value(parameter11Value).Parameter12Name(parameter12Name).Parameter12Value(parameter12Value).Parameter13Name(parameter13Name).Parameter13Value(parameter13Value).Parameter14Name(parameter14Name).Parameter14Value(parameter14Value).Parameter15Name(parameter15Name).Parameter15Value(parameter15Value).Parameter16Name(parameter16Name).Parameter16Value(parameter16Value).Parameter17Name(parameter17Name).Parameter17Value(parameter17Value).Parameter18Name(parameter18Name).Parameter18Value(parameter18Value).Parameter19Name(parameter19Name).Parameter19Value(parameter19Value).Parameter2Name(parameter2Name).Parameter2Value(parameter2Value).Parameter20Name(parameter20Name).Parameter20Value(parameter20Value).Parameter21Name(parameter21Name).Parameter21Value(parameter21Value).Parameter22Name(parameter22Name).Parameter22Value(parameter22Value).Parameter23Name(parameter23Name).Parameter23Value(parameter23Value).Parameter24Name(parameter24Name).Parameter24Value(parameter24Value).Parameter25Name(parameter25Name).Parameter25Value(parameter25Value).Parameter26Name(parameter26Name).Parameter26Value(parameter26Value).Parameter27Name(parameter27Name).Parameter27Value(parameter27Value).Parameter28Name(parameter28Name).Parameter28Value(parameter28Value).Parameter29Name(parameter29Name).Parameter29Value(parameter29Value).Parameter3Name(parameter3Name).Parameter3Value(parameter3Value).Parameter30Name(parameter30Name).Parameter30Value(parameter30Value).Parameter31Name(parameter31Name).Parameter31Value(parameter31Value).Parameter32Name(parameter32Name).Parameter32Value(parameter32Value).Parameter33Name(parameter33Name).Parameter33Value(parameter33Value).Parameter34Name(parameter34Name).Parameter34Value(parameter34Value).Parameter35Name(parameter35Name).Parameter35Value(parameter35Value).Parameter36Name(parameter36Name).Parameter36Value(parameter36Value).Parameter37Name(parameter37Name).Parameter37Value(parameter37Value).Parameter38Name(parameter38Name).Parameter38Value(parameter38Value).Parameter39Name(parameter39Name).Parameter39Value(parameter39Value).Parameter4Name(parameter4Name).Parameter4Value(parameter4Value).Parameter40Name(parameter40Name).Parameter40Value(parameter40Value).Parameter41Name(parameter41Name).Parameter41Value(parameter41Value).Parameter42Name(parameter42Name).Parameter42Value(parameter42Value).Parameter43Name(parameter43Name).Parameter43Value(parameter43Value).Parameter44Name(parameter44Name).Parameter44Value(parameter44Value).Parameter45Name(parameter45Name).Parameter45Value(parameter45Value).Parameter46Name(parameter46Name).Parameter46Value(parameter46Value).Parameter47Name(parameter47Name).Parameter47Value(parameter47Value).Parameter48Name(parameter48Name).Parameter48Value(parameter48Value).Parameter49Name(parameter49Name).Parameter49Value(parameter49Value).Parameter5Name(parameter5Name).Parameter5Value(parameter5Value).Parameter50Name(parameter50Name).Parameter50Value(parameter50Value).Parameter51Name(parameter51Name).Parameter51Value(parameter51Value).Parameter52Name(parameter52Name).Parameter52Value(parameter52Value).Parameter53Name(parameter53Name).Parameter53Value(parameter53Value).Parameter54Name(parameter54Name).Parameter54Value(parameter54Value).Parameter55Name(parameter55Name).Parameter55Value(parameter55Value).Parameter56Name(parameter56Name).Parameter56Value(parameter56Value).Parameter57Name(parameter57Name).Parameter57Value(parameter57Value).Parameter58Name(parameter58Name).Parameter58Value(parameter58Value).Parameter59Name(parameter59Name).Parameter59Value(parameter59Value).Parameter6Name(parameter6Name).Parameter6Value(parameter6Value).Parameter60Name(parameter60Name).Parameter60Value(parameter60Value).Parameter61Name(parameter61Name).Parameter61Value(parameter61Value).Parameter62Name(parameter62Name).Parameter62Value(parameter62Value).Parameter63Name(parameter63Name).Parameter63Value(parameter63Value).Parameter64Name(parameter64Name).Parameter64Value(parameter64Value).Parameter65Name(parameter65Name).Parameter65Value(parameter65Value).Parameter66Name(parameter66Name).Parameter66Value(parameter66Value).Parameter67Name(parameter67Name).Parameter67Value(parameter67Value).Parameter68Name(parameter68Name).Parameter68Value(parameter68Value).Parameter69Name(parameter69Name).Parameter69Value(parameter69Value).Parameter7Name(parameter7Name).Parameter7Value(parameter7Value).Parameter70Name(parameter70Name).Parameter70Value(parameter70Value).Parameter71Name(parameter71Name).Parameter71Value(parameter71Value).Parameter72Name(parameter72Name).Parameter72Value(parameter72Value).Parameter73Name(parameter73Name).Parameter73Value(parameter73Value).Parameter74Name(parameter74Name).Parameter74Value(parameter74Value).Parameter75Name(parameter75Name).Parameter75Value(parameter75Value).Parameter76Name(parameter76Name).Parameter76Value(parameter76Value).Parameter77Name(parameter77Name).Parameter77Value(parameter77Value).Parameter78Name(parameter78Name).Parameter78Value(parameter78Value).Parameter79Name(parameter79Name).Parameter79Value(parameter79Value).Parameter8Name(parameter8Name).Parameter8Value(parameter8Value).Parameter80Name(parameter80Name).Parameter80Value(parameter80Value).Parameter81Name(parameter81Name).Parameter81Value(parameter81Value).Parameter82Name(parameter82Name).Parameter82Value(parameter82Value).Parameter83Name(parameter83Name).Parameter83Value(parameter83Value).Parameter84Name(parameter84Name).Parameter84Value(parameter84Value).Parameter85Name(parameter85Name).Parameter85Value(parameter85Value).Parameter86Name(parameter86Name).Parameter86Value(parameter86Value).Parameter87Name(parameter87Name).Parameter87Value(parameter87Value).Parameter88Name(parameter88Name).Parameter88Value(parameter88Value).Parameter89Name(parameter89Name).Parameter89Value(parameter89Value).Parameter9Name(parameter9Name).Parameter9Value(parameter9Value).Parameter90Name(parameter90Name).Parameter90Value(parameter90Value).Parameter91Name(parameter91Name).Parameter91Value(parameter91Value).Parameter92Name(parameter92Name).Parameter92Value(parameter92Value).Parameter93Name(parameter93Name).Parameter93Value(parameter93Value).Parameter94Name(parameter94Name).Parameter94Value(parameter94Value).Parameter95Name(parameter95Name).Parameter95Value(parameter95Value).Parameter96Name(parameter96Name).Parameter96Value(parameter96Value).Parameter97Name(parameter97Name).Parameter97Value(parameter97Value).Parameter98Name(parameter98Name).Parameter98Value(parameter98Value).Parameter99Name(parameter99Name).Parameter99Value(parameter99Value).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStream`: ApiV2010AccountCallStream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Stream resource. | 
**callSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Stream resource is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **url** | **string** | Relative or absolute url where WebSocket connection will be established. | 
 **name** | **string** | The user-specified name of this Stream, if one was given when the Stream was created. This may be used to stop the Stream. | 
 **parameter1Name** | **string** | Parameter name | 
 **parameter1Value** | **string** | Parameter value | 
 **parameter10Name** | **string** | Parameter name | 
 **parameter10Value** | **string** | Parameter value | 
 **parameter11Name** | **string** | Parameter name | 
 **parameter11Value** | **string** | Parameter value | 
 **parameter12Name** | **string** | Parameter name | 
 **parameter12Value** | **string** | Parameter value | 
 **parameter13Name** | **string** | Parameter name | 
 **parameter13Value** | **string** | Parameter value | 
 **parameter14Name** | **string** | Parameter name | 
 **parameter14Value** | **string** | Parameter value | 
 **parameter15Name** | **string** | Parameter name | 
 **parameter15Value** | **string** | Parameter value | 
 **parameter16Name** | **string** | Parameter name | 
 **parameter16Value** | **string** | Parameter value | 
 **parameter17Name** | **string** | Parameter name | 
 **parameter17Value** | **string** | Parameter value | 
 **parameter18Name** | **string** | Parameter name | 
 **parameter18Value** | **string** | Parameter value | 
 **parameter19Name** | **string** | Parameter name | 
 **parameter19Value** | **string** | Parameter value | 
 **parameter2Name** | **string** | Parameter name | 
 **parameter2Value** | **string** | Parameter value | 
 **parameter20Name** | **string** | Parameter name | 
 **parameter20Value** | **string** | Parameter value | 
 **parameter21Name** | **string** | Parameter name | 
 **parameter21Value** | **string** | Parameter value | 
 **parameter22Name** | **string** | Parameter name | 
 **parameter22Value** | **string** | Parameter value | 
 **parameter23Name** | **string** | Parameter name | 
 **parameter23Value** | **string** | Parameter value | 
 **parameter24Name** | **string** | Parameter name | 
 **parameter24Value** | **string** | Parameter value | 
 **parameter25Name** | **string** | Parameter name | 
 **parameter25Value** | **string** | Parameter value | 
 **parameter26Name** | **string** | Parameter name | 
 **parameter26Value** | **string** | Parameter value | 
 **parameter27Name** | **string** | Parameter name | 
 **parameter27Value** | **string** | Parameter value | 
 **parameter28Name** | **string** | Parameter name | 
 **parameter28Value** | **string** | Parameter value | 
 **parameter29Name** | **string** | Parameter name | 
 **parameter29Value** | **string** | Parameter value | 
 **parameter3Name** | **string** | Parameter name | 
 **parameter3Value** | **string** | Parameter value | 
 **parameter30Name** | **string** | Parameter name | 
 **parameter30Value** | **string** | Parameter value | 
 **parameter31Name** | **string** | Parameter name | 
 **parameter31Value** | **string** | Parameter value | 
 **parameter32Name** | **string** | Parameter name | 
 **parameter32Value** | **string** | Parameter value | 
 **parameter33Name** | **string** | Parameter name | 
 **parameter33Value** | **string** | Parameter value | 
 **parameter34Name** | **string** | Parameter name | 
 **parameter34Value** | **string** | Parameter value | 
 **parameter35Name** | **string** | Parameter name | 
 **parameter35Value** | **string** | Parameter value | 
 **parameter36Name** | **string** | Parameter name | 
 **parameter36Value** | **string** | Parameter value | 
 **parameter37Name** | **string** | Parameter name | 
 **parameter37Value** | **string** | Parameter value | 
 **parameter38Name** | **string** | Parameter name | 
 **parameter38Value** | **string** | Parameter value | 
 **parameter39Name** | **string** | Parameter name | 
 **parameter39Value** | **string** | Parameter value | 
 **parameter4Name** | **string** | Parameter name | 
 **parameter4Value** | **string** | Parameter value | 
 **parameter40Name** | **string** | Parameter name | 
 **parameter40Value** | **string** | Parameter value | 
 **parameter41Name** | **string** | Parameter name | 
 **parameter41Value** | **string** | Parameter value | 
 **parameter42Name** | **string** | Parameter name | 
 **parameter42Value** | **string** | Parameter value | 
 **parameter43Name** | **string** | Parameter name | 
 **parameter43Value** | **string** | Parameter value | 
 **parameter44Name** | **string** | Parameter name | 
 **parameter44Value** | **string** | Parameter value | 
 **parameter45Name** | **string** | Parameter name | 
 **parameter45Value** | **string** | Parameter value | 
 **parameter46Name** | **string** | Parameter name | 
 **parameter46Value** | **string** | Parameter value | 
 **parameter47Name** | **string** | Parameter name | 
 **parameter47Value** | **string** | Parameter value | 
 **parameter48Name** | **string** | Parameter name | 
 **parameter48Value** | **string** | Parameter value | 
 **parameter49Name** | **string** | Parameter name | 
 **parameter49Value** | **string** | Parameter value | 
 **parameter5Name** | **string** | Parameter name | 
 **parameter5Value** | **string** | Parameter value | 
 **parameter50Name** | **string** | Parameter name | 
 **parameter50Value** | **string** | Parameter value | 
 **parameter51Name** | **string** | Parameter name | 
 **parameter51Value** | **string** | Parameter value | 
 **parameter52Name** | **string** | Parameter name | 
 **parameter52Value** | **string** | Parameter value | 
 **parameter53Name** | **string** | Parameter name | 
 **parameter53Value** | **string** | Parameter value | 
 **parameter54Name** | **string** | Parameter name | 
 **parameter54Value** | **string** | Parameter value | 
 **parameter55Name** | **string** | Parameter name | 
 **parameter55Value** | **string** | Parameter value | 
 **parameter56Name** | **string** | Parameter name | 
 **parameter56Value** | **string** | Parameter value | 
 **parameter57Name** | **string** | Parameter name | 
 **parameter57Value** | **string** | Parameter value | 
 **parameter58Name** | **string** | Parameter name | 
 **parameter58Value** | **string** | Parameter value | 
 **parameter59Name** | **string** | Parameter name | 
 **parameter59Value** | **string** | Parameter value | 
 **parameter6Name** | **string** | Parameter name | 
 **parameter6Value** | **string** | Parameter value | 
 **parameter60Name** | **string** | Parameter name | 
 **parameter60Value** | **string** | Parameter value | 
 **parameter61Name** | **string** | Parameter name | 
 **parameter61Value** | **string** | Parameter value | 
 **parameter62Name** | **string** | Parameter name | 
 **parameter62Value** | **string** | Parameter value | 
 **parameter63Name** | **string** | Parameter name | 
 **parameter63Value** | **string** | Parameter value | 
 **parameter64Name** | **string** | Parameter name | 
 **parameter64Value** | **string** | Parameter value | 
 **parameter65Name** | **string** | Parameter name | 
 **parameter65Value** | **string** | Parameter value | 
 **parameter66Name** | **string** | Parameter name | 
 **parameter66Value** | **string** | Parameter value | 
 **parameter67Name** | **string** | Parameter name | 
 **parameter67Value** | **string** | Parameter value | 
 **parameter68Name** | **string** | Parameter name | 
 **parameter68Value** | **string** | Parameter value | 
 **parameter69Name** | **string** | Parameter name | 
 **parameter69Value** | **string** | Parameter value | 
 **parameter7Name** | **string** | Parameter name | 
 **parameter7Value** | **string** | Parameter value | 
 **parameter70Name** | **string** | Parameter name | 
 **parameter70Value** | **string** | Parameter value | 
 **parameter71Name** | **string** | Parameter name | 
 **parameter71Value** | **string** | Parameter value | 
 **parameter72Name** | **string** | Parameter name | 
 **parameter72Value** | **string** | Parameter value | 
 **parameter73Name** | **string** | Parameter name | 
 **parameter73Value** | **string** | Parameter value | 
 **parameter74Name** | **string** | Parameter name | 
 **parameter74Value** | **string** | Parameter value | 
 **parameter75Name** | **string** | Parameter name | 
 **parameter75Value** | **string** | Parameter value | 
 **parameter76Name** | **string** | Parameter name | 
 **parameter76Value** | **string** | Parameter value | 
 **parameter77Name** | **string** | Parameter name | 
 **parameter77Value** | **string** | Parameter value | 
 **parameter78Name** | **string** | Parameter name | 
 **parameter78Value** | **string** | Parameter value | 
 **parameter79Name** | **string** | Parameter name | 
 **parameter79Value** | **string** | Parameter value | 
 **parameter8Name** | **string** | Parameter name | 
 **parameter8Value** | **string** | Parameter value | 
 **parameter80Name** | **string** | Parameter name | 
 **parameter80Value** | **string** | Parameter value | 
 **parameter81Name** | **string** | Parameter name | 
 **parameter81Value** | **string** | Parameter value | 
 **parameter82Name** | **string** | Parameter name | 
 **parameter82Value** | **string** | Parameter value | 
 **parameter83Name** | **string** | Parameter name | 
 **parameter83Value** | **string** | Parameter value | 
 **parameter84Name** | **string** | Parameter name | 
 **parameter84Value** | **string** | Parameter value | 
 **parameter85Name** | **string** | Parameter name | 
 **parameter85Value** | **string** | Parameter value | 
 **parameter86Name** | **string** | Parameter name | 
 **parameter86Value** | **string** | Parameter value | 
 **parameter87Name** | **string** | Parameter name | 
 **parameter87Value** | **string** | Parameter value | 
 **parameter88Name** | **string** | Parameter name | 
 **parameter88Value** | **string** | Parameter value | 
 **parameter89Name** | **string** | Parameter name | 
 **parameter89Value** | **string** | Parameter value | 
 **parameter9Name** | **string** | Parameter name | 
 **parameter9Value** | **string** | Parameter value | 
 **parameter90Name** | **string** | Parameter name | 
 **parameter90Value** | **string** | Parameter value | 
 **parameter91Name** | **string** | Parameter name | 
 **parameter91Value** | **string** | Parameter value | 
 **parameter92Name** | **string** | Parameter name | 
 **parameter92Value** | **string** | Parameter value | 
 **parameter93Name** | **string** | Parameter name | 
 **parameter93Value** | **string** | Parameter value | 
 **parameter94Name** | **string** | Parameter name | 
 **parameter94Value** | **string** | Parameter value | 
 **parameter95Name** | **string** | Parameter name | 
 **parameter95Value** | **string** | Parameter value | 
 **parameter96Name** | **string** | Parameter name | 
 **parameter96Value** | **string** | Parameter value | 
 **parameter97Name** | **string** | Parameter name | 
 **parameter97Value** | **string** | Parameter value | 
 **parameter98Name** | **string** | Parameter name | 
 **parameter98Value** | **string** | Parameter value | 
 **parameter99Name** | **string** | Parameter name | 
 **parameter99Value** | **string** | Parameter value | 
 **statusCallback** | **string** | Absolute URL of the status callback. | 
 **statusCallbackMethod** | **string** | The http method for the status_callback (one of GET, POST). | 
 **track** | **string** | One of &#x60;inbound_track&#x60;, &#x60;outbound_track&#x60;, &#x60;both_tracks&#x60;. | 

### Return type

[**ApiV2010AccountCallStream**](ApiV2010AccountCallStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> ApiV2010AccountToken CreateToken(ctx, accountSid).Ttl(ttl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ttl := int32(56) // int32 | The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateToken(context.Background(), accountSid).Ttl(ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: ApiV2010AccountToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ttl** | **int32** | The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours). | 

### Return type

[**ApiV2010AccountToken**](ApiV2010AccountToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsageTrigger

> ApiV2010AccountUsageUsageTrigger CreateUsageTrigger(ctx, accountSid).CallbackUrl(callbackUrl).TriggerValue(triggerValue).UsageCategory(usageCategory).CallbackMethod(callbackMethod).FriendlyName(friendlyName).Recurring(recurring).TriggerBy(triggerBy).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    callbackUrl := "callbackUrl_example" // string | The URL we should call using `callback_method` when the trigger fires.
    triggerValue := "triggerValue_example" // string | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`.
    usageCategory := "usageCategory_example" // string | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value.
    callbackMethod := "callbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    recurring := "recurring_example" // string | The frequency of a recurring UsageTrigger.  Can be: `daily`, `monthly`, or `yearly` for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT. (optional)
    triggerBy := "triggerBy_example" // string | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is `usage`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateUsageTrigger(context.Background(), accountSid).CallbackUrl(callbackUrl).TriggerValue(triggerValue).UsageCategory(usageCategory).CallbackMethod(callbackMethod).FriendlyName(friendlyName).Recurring(recurring).TriggerBy(triggerBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsageTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackUrl** | **string** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires. | 
 **triggerValue** | **string** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as &#x60;+30&#x60; to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a &#x60;+&#x60; as &#x60;%2B&#x60;. | 
 **usageCategory** | **string** | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value. | 
 **callbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **recurring** | **string** | The frequency of a recurring UsageTrigger.  Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT. | 
 **triggerBy** | **string** | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is &#x60;usage&#x60;. | 

### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateValidationRequest

> ApiV2010AccountValidationRequest CreateValidationRequest(ctx, accountSid).PhoneNumber(phoneNumber).CallDelay(callDelay).Extension(extension).FriendlyName(friendlyName).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource.
    phoneNumber := "phoneNumber_example" // string | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
    callDelay := int32(56) // int32 | The number of seconds to delay before initiating the verification call. Can be an integer between `0` and `60`, inclusive. The default is `0`. (optional)
    extension := "extension_example" // string | The digits to dial after connecting the verification call. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information about the verification process to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`, and the default is `POST`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateValidationRequest(context.Background(), accountSid).PhoneNumber(phoneNumber).CallDelay(callDelay).Extension(extension).FriendlyName(friendlyName).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateValidationRequest`: ApiV2010AccountValidationRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateValidationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateValidationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumber** | **string** | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. | 
 **callDelay** | **int32** | The number of seconds to delay before initiating the verification call. Can be an integer between &#x60;0&#x60; and &#x60;60&#x60;, inclusive. The default is &#x60;0&#x60;. | 
 **extension** | **string** | The digits to dial after connecting the verification call. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information about the verification process to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;, and the default is &#x60;POST&#x60;. | 

### Return type

[**ApiV2010AccountValidationRequest**](ApiV2010AccountValidationRequest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAddress(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteApplication(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCall

> DeleteCall(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete.
    sid := "sid_example" // string | The Twilio-provided Call SID that uniquely identifies the Call resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCall(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete. | 
**sid** | **string** | The Twilio-provided Call SID that uniquely identifies the Call resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallFeedbackSummary

> DeleteCallFeedbackSummary(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    sid := "sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCallFeedbackSummary(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallFeedbackSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallRecording

> DeleteCallRecording(ctx, accountSid, callSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCallRecording(context.Background(), accountSid, callSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConferenceRecording

> DeleteConferenceRecording(ctx, accountSid, conferenceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.
    conferenceSid := "conferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteConferenceRecording(context.Background(), accountSid, conferenceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete. | 
**conferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectApp

> DeleteConnectApp(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteConnectApp(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumber

> DeleteIncomingPhoneNumber(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteIncomingPhoneNumber(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncomingPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumberAssignedAddOn

> DeleteIncomingPhoneNumberAssignedAddOn(ctx, accountSid, resourceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteIncomingPhoneNumberAssignedAddOn(context.Background(), accountSid, resourceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete. | 
**resourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncomingPhoneNumberAssignedAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteKey(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMedia

> DeleteMedia(ctx, accountSid, messageSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete.
    messageSid := "messageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Media resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteMedia(context.Background(), accountSid, messageSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete. | 
**messageSid** | **string** | The SID of the Message resource that this Media resource belongs to. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Media resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessage

> DeleteMessage(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteMessage(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutgoingCallerId

> DeleteOutgoingCallerId(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteOutgoingCallerId(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutgoingCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParticipant

> DeleteParticipant(ctx, accountSid, conferenceSid, callSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete.
    conferenceSid := "conferenceSid_example" // string | The SID of the conference with the participants to delete.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteParticipant(context.Background(), accountSid, conferenceSid, callSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete. | 
**conferenceSid** | **string** | The SID of the conference with the participants to delete. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueue

> DeleteQueue(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteQueue(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecording

> DeleteRecording(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRecording(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingAddOnResult

> DeleteRecordingAddOnResult(ctx, accountSid, referenceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the result to delete belongs.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRecordingAddOnResult(context.Background(), accountSid, referenceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete. | 
**referenceSid** | **string** | The SID of the recording to which the result to delete belongs. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingAddOnResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingAddOnResultPayload

> DeleteRecordingAddOnResultPayload(ctx, accountSid, referenceSid, addOnResultSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs.
    addOnResultSid := "addOnResultSid_example" // string | The SID of the AddOnResult to which the payloads to delete belongs.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRecordingAddOnResultPayload(context.Background(), accountSid, referenceSid, addOnResultSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete. | 
**referenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs. | 
**addOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to delete belongs. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingAddOnResultPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingTranscription

> DeleteRecordingTranscription(ctx, accountSid, recordingSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
    recordingSid := "recordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRecordingTranscription(context.Background(), accountSid, recordingSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete. | 
**recordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSigningKey

> DeleteSigningKey(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | 
    sid := "sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSigningKey(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsCredentialListMapping

> DeleteSipAuthCallsCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resource to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipAuthCallsCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resource to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipAuthCallsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsIpAccessControlListMapping

> DeleteSipAuthCallsIpAccessControlListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipAuthCallsIpAccessControlListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipAuthCallsIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthRegistrationsCredentialListMapping

> DeleteSipAuthRegistrationsCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipAuthRegistrationsCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipAuthRegistrationsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredential

> DeleteSipCredential(ctx, accountSid, credentialListSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    credentialListSid := "credentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credentials.
    sid := "sid_example" // string | The unique id that identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipCredential(context.Background(), accountSid, credentialListSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**credentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials. | 
**sid** | **string** | The unique id that identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialList

> DeleteSipCredentialList(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    sid := "sid_example" // string | The credential list Sid that uniquely identifies this resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipCredentialList(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**sid** | **string** | The credential list Sid that uniquely identifies this resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialListMapping

> DeleteSipCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipDomain

> DeleteSipDomain(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipDomain(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlList

> DeleteSipIpAccessControlList(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipIpAccessControlList(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipIpAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlListMapping

> DeleteSipIpAccessControlListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipIpAccessControlListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP domain. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAddress

> DeleteSipIpAddress(ctx, accountSid, ipAccessControlListSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to delete.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSipIpAddress(context.Background(), accountSid, ipAccessControlListSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**ipAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to delete. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranscription

> DeleteTranscription(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteTranscription(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsageTrigger

> DeleteUsageTrigger(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteUsageTrigger(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsageTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccount

> ApiV2010Account FetchAccount(ctx, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sid := "sid_example" // string | The Account Sid that uniquely identifies the account to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchAccount(context.Background(), sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | The Account Sid that uniquely identifies the account to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAddress

> ApiV2010AccountAddress FetchAddress(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchAddress(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchApplication

> ApiV2010AccountApplication FetchApplication(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchApplication(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizedConnectApp

> ApiV2010AccountAuthorizedConnectApp FetchAuthorizedConnectApp(ctx, accountSid, connectAppSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
    connectAppSid := "connectAppSid_example" // string | The SID of the Connect App to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchAuthorizedConnectApp(context.Background(), accountSid, connectAppSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAuthorizedConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAuthorizedConnectApp`: ApiV2010AccountAuthorizedConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAuthorizedConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch. | 
**connectAppSid** | **string** | The SID of the Connect App to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAuthorizedConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAuthorizedConnectApp**](ApiV2010AccountAuthorizedConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailablePhoneNumberCountry

> ApiV2010AccountAvailablePhoneNumberCountry FetchAvailablePhoneNumberCountry(ctx, accountSid, countryCode).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country to fetch available phone number information about.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchAvailablePhoneNumberCountry(context.Background(), accountSid, countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAvailablePhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAvailablePhoneNumberCountry`: ApiV2010AccountAvailablePhoneNumberCountry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAvailablePhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country to fetch available phone number information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAvailablePhoneNumberCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAvailablePhoneNumberCountry**](ApiV2010AccountAvailablePhoneNumberCountry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBalance

> ApiV2010AccountBalance FetchBalance(ctx, accountSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique SID identifier of the Account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchBalance(context.Background(), accountSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBalance`: ApiV2010AccountBalance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique SID identifier of the Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV2010AccountBalance**](ApiV2010AccountBalance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCall

> ApiV2010AccountCall FetchCall(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch.
    sid := "sid_example" // string | The SID of the Call resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchCall(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch. | 
**sid** | **string** | The SID of the Call resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedback

> ApiV2010AccountCallCallFeedback FetchCallFeedback(ctx, accountSid, callSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    callSid := "callSid_example" // string | The call sid that uniquely identifies the call

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchCallFeedback(context.Background(), accountSid, callSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallFeedback`: ApiV2010AccountCallCallFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**callSid** | **string** | The call sid that uniquely identifies the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCallFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCallCallFeedback**](ApiV2010AccountCallCallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary FetchCallFeedbackSummary(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    sid := "sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchCallFeedbackSummary(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallFeedbackSummary`: ApiV2010AccountCallCallFeedbackSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallFeedbackSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCallFeedbackSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](ApiV2010AccountCallCallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallNotification

> ApiV2010AccountCallCallNotificationInstance FetchCallNotification(ctx, accountSid, callSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Call Notification resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchCallNotification(context.Background(), accountSid, callSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallNotification`: ApiV2010AccountCallCallNotificationInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Call Notification resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCallNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountCallCallNotificationInstance**](ApiV2010AccountCallCallNotificationInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallRecording

> ApiV2010AccountCallCallRecording FetchCallRecording(ctx, accountSid, callSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchCallRecording(context.Background(), accountSid, callSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConference

> ApiV2010AccountConference FetchConference(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Conference resource to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchConference(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConference`: ApiV2010AccountConference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Conference resource to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountConference**](ApiV2010AccountConference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConferenceRecording

> ApiV2010AccountConferenceConferenceRecording FetchConferenceRecording(ctx, accountSid, conferenceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.
    conferenceSid := "conferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchConferenceRecording(context.Background(), accountSid, conferenceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConferenceRecording`: ApiV2010AccountConferenceConferenceRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch. | 
**conferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectApp

> ApiV2010AccountConnectApp FetchConnectApp(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchConnectApp(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConnectApp`: ApiV2010AccountConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountConnectApp**](ApiV2010AccountConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber FetchIncomingPhoneNumber(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchIncomingPhoneNumber(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIncomingPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn FetchIncomingPhoneNumberAssignedAddOn(ctx, accountSid, resourceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchIncomingPhoneNumberAssignedAddOn(context.Background(), accountSid, resourceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumberAssignedAddOn`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch. | 
**resourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIncomingPhoneNumberAssignedAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOnExtension

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension FetchIncomingPhoneNumberAssignedAddOnExtension(ctx, accountSid, resourceSid, assignedAddOnSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    assignedAddOnSid := "assignedAddOnSid_example" // string | The SID that uniquely identifies the assigned Add-on installation.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension(context.Background(), accountSid, resourceSid, assignedAddOnSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumberAssignedAddOnExtension`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch. | 
**resourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned. | 
**assignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIncomingPhoneNumberAssignedAddOnExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKey

> ApiV2010AccountKey FetchKey(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchKey(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchKey`: ApiV2010AccountKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountKey**](ApiV2010AccountKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMedia

> ApiV2010AccountMessageMedia FetchMedia(ctx, accountSid, messageSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch.
    messageSid := "messageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Media resource to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchMedia(context.Background(), accountSid, messageSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMedia`: ApiV2010AccountMessageMedia
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch. | 
**messageSid** | **string** | The SID of the Message resource that this Media resource belongs to. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Media resource to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountMessageMedia**](ApiV2010AccountMessageMedia.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ApiV2010AccountQueueMember FetchMember(ctx, accountSid, queueSid, callSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch.
    queueSid := "queueSid_example" // string | The SID of the Queue in which to find the members to fetch.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchMember(context.Background(), accountSid, queueSid, callSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMember`: ApiV2010AccountQueueMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch. | 
**queueSid** | **string** | The SID of the Queue in which to find the members to fetch. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountQueueMember**](ApiV2010AccountQueueMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ApiV2010AccountMessage FetchMessage(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchMessage(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNotification

> ApiV2010AccountNotificationInstance FetchNotification(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Notification resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchNotification(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNotification`: ApiV2010AccountNotificationInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Notification resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountNotificationInstance**](ApiV2010AccountNotificationInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOutgoingCallerId

> ApiV2010AccountOutgoingCallerId FetchOutgoingCallerId(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchOutgoingCallerId(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchOutgoingCallerId`: ApiV2010AccountOutgoingCallerId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchOutgoingCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountOutgoingCallerId**](ApiV2010AccountOutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ApiV2010AccountConferenceParticipant FetchParticipant(ctx, accountSid, conferenceSid, callSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch.
    conferenceSid := "conferenceSid_example" // string | The SID of the conference with the participant to fetch.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchParticipant(context.Background(), accountSid, conferenceSid, callSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch. | 
**conferenceSid** | **string** | The SID of the conference with the participant to fetch. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQueue

> ApiV2010AccountQueue FetchQueue(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchQueue(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> ApiV2010AccountRecording FetchRecording(ctx, accountSid, sid).IncludeSoftDeleted(includeSoftDeleted).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to fetch.
    includeSoftDeleted := true // bool | A boolean parameter indicating whether to retrieve soft deleted recordings or not. Recordings metadata are kept after deletion for a retention period of 40 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchRecording(context.Background(), accountSid, sid).IncludeSoftDeleted(includeSoftDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecording`: ApiV2010AccountRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSoftDeleted** | **bool** | A boolean parameter indicating whether to retrieve soft deleted recordings or not. Recordings metadata are kept after deletion for a retention period of 40 days. | 

### Return type

[**ApiV2010AccountRecording**](ApiV2010AccountRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResult

> ApiV2010AccountRecordingRecordingAddOnResult FetchRecordingAddOnResult(ctx, accountSid, referenceSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the result to fetch belongs.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchRecordingAddOnResult(context.Background(), accountSid, referenceSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingAddOnResult`: ApiV2010AccountRecordingRecordingAddOnResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingAddOnResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch. | 
**referenceSid** | **string** | The SID of the recording to which the result to fetch belongs. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRecordingAddOnResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountRecordingRecordingAddOnResult**](ApiV2010AccountRecordingRecordingAddOnResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResultPayload

> ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload FetchRecordingAddOnResultPayload(ctx, accountSid, referenceSid, addOnResultSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs.
    addOnResultSid := "addOnResultSid_example" // string | The SID of the AddOnResult to which the payload to fetch belongs.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchRecordingAddOnResultPayload(context.Background(), accountSid, referenceSid, addOnResultSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingAddOnResultPayload`: ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingAddOnResultPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch. | 
**referenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs. | 
**addOnResultSid** | **string** | The SID of the AddOnResult to which the payload to fetch belongs. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRecordingAddOnResultPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload**](ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingTranscription

> ApiV2010AccountRecordingRecordingTranscription FetchRecordingTranscription(ctx, accountSid, recordingSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
    recordingSid := "recordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchRecordingTranscription(context.Background(), accountSid, recordingSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingTranscription`: ApiV2010AccountRecordingRecordingTranscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch. | 
**recordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountRecordingRecordingTranscription**](ApiV2010AccountRecordingRecordingTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> ApiV2010AccountShortCode FetchShortCode(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchShortCode(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchShortCode`: ApiV2010AccountShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountShortCode**](ApiV2010AccountShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSigningKey

> ApiV2010AccountSigningKey FetchSigningKey(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | 
    sid := "sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSigningKey(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSigningKey`: ApiV2010AccountSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSigningKey**](ApiV2010AccountSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping FetchSipAuthCallsCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipAuthCallsCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthCallsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipAuthCallsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping FetchSipAuthCallsIpAccessControlListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipAuthCallsIpAccessControlListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthCallsIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipAuthCallsIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping FetchSipAuthRegistrationsCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipAuthRegistrationsCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthRegistrationsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipAuthRegistrationsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential FetchSipCredential(ctx, accountSid, credentialListSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    credentialListSid := "credentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credential.
    sid := "sid_example" // string | The unique id that identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipCredential(context.Background(), accountSid, credentialListSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**credentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credential. | 
**sid** | **string** | The unique id that identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialList

> ApiV2010AccountSipSipCredentialList FetchSipCredentialList(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    sid := "sid_example" // string | The credential list Sid that uniquely identifies this resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipCredentialList(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**sid** | **string** | The credential list Sid that uniquely identifies this resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping FetchSipCredentialListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipCredentialListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredentialListMapping`: ApiV2010AccountSipSipDomainSipCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](ApiV2010AccountSipSipDomainSipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipDomain

> ApiV2010AccountSipSipDomain FetchSipDomain(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipDomain(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList FetchSipIpAccessControlList(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipIpAccessControlList(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipIpAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping FetchSipIpAccessControlListMapping(ctx, accountSid, domainSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipIpAccessControlListMapping(context.Background(), accountSid, domainSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP domain. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress FetchSipIpAddress(ctx, accountSid, ipAccessControlListSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to fetch.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the IpAddress resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchSipIpAddress(context.Background(), accountSid, ipAccessControlListSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**ipAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to fetch. | 
**sid** | **string** | A 34 character string that uniquely identifies the IpAddress resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSipIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTranscription

> ApiV2010AccountTranscription FetchTranscription(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchTranscription(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTranscription`: ApiV2010AccountTranscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountTranscription**](ApiV2010AccountTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsageTrigger

> ApiV2010AccountUsageUsageTrigger FetchUsageTrigger(ctx, accountSid, sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FetchUsageTrigger(context.Background(), accountSid, sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsageTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccount

> ListAccountResponse ListAccount(ctx).FriendlyName(friendlyName).Status(status).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    friendlyName := "friendlyName_example" // string | Only return the Account resources with friendly names that exactly match this name. (optional)
    status := "status_example" // string | Only return Account resources with the given status. Can be `closed`, `suspended` or `active`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAccount(context.Background()).FriendlyName(friendlyName).Status(status).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccount`: ListAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **string** | Only return the Account resources with friendly names that exactly match this name. | 
 **status** | **string** | Only return Account resources with the given status. Can be &#x60;closed&#x60;, &#x60;suspended&#x60; or &#x60;active&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAccountResponse**](ListAccountResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddress

> ListAddressResponse ListAddress(ctx, accountSid).CustomerName(customerName).FriendlyName(friendlyName).IsoCountry(isoCountry).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read.
    customerName := "customerName_example" // string | The `customer_name` of the Address resources to read. (optional)
    friendlyName := "friendlyName_example" // string | The string that identifies the Address resources to read. (optional)
    isoCountry := "isoCountry_example" // string | The ISO country code of the Address resources to read. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAddress(context.Background(), accountSid).CustomerName(customerName).FriendlyName(friendlyName).IsoCountry(isoCountry).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAddress`: ListAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerName** | **string** | The &#x60;customer_name&#x60; of the Address resources to read. | 
 **friendlyName** | **string** | The string that identifies the Address resources to read. | 
 **isoCountry** | **string** | The ISO country code of the Address resources to read. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAddressResponse**](ListAddressResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplication

> ListApplicationResponse ListApplication(ctx, accountSid).FriendlyName(friendlyName).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read.
    friendlyName := "friendlyName_example" // string | The string that identifies the Application resources to read. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListApplication(context.Background(), accountSid).FriendlyName(friendlyName).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplication`: ListApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | The string that identifies the Application resources to read. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListApplicationResponse**](ListApplicationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizedConnectApp

> ListAuthorizedConnectAppResponse ListAuthorizedConnectApp(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAuthorizedConnectApp(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAuthorizedConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizedConnectApp`: ListAuthorizedConnectAppResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAuthorizedConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizedConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAuthorizedConnectAppResponse**](ListAuthorizedConnectAppResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberCountry

> ListAvailablePhoneNumberCountryResponse ListAvailablePhoneNumberCountry(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberCountry(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberCountry`: ListAvailablePhoneNumberCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberCountryResponse**](ListAvailablePhoneNumberCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberLocal

> ListAvailablePhoneNumberLocalResponse ListAvailablePhoneNumberLocal(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample=code-find-phone-numbers-by-number-pattern) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample=code-find-phone-numbers-by-character-pattern). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberLocal(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberLocal`: ListAvailablePhoneNumberLocalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberLocalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-number-pattern) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-character-pattern). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberLocalResponse**](ListAvailablePhoneNumberLocalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberMachineToMachine

> ListAvailablePhoneNumberMachineToMachineResponse ListAvailablePhoneNumberMachineToMachine(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberMachineToMachine(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberMachineToMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberMachineToMachine`: ListAvailablePhoneNumberMachineToMachineResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberMachineToMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberMachineToMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberMachineToMachineResponse**](ListAvailablePhoneNumberMachineToMachineResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberMobile

> ListAvailablePhoneNumberMobileResponse ListAvailablePhoneNumberMobile(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberMobile(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberMobile`: ListAvailablePhoneNumberMobileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberMobileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberMobileResponse**](ListAvailablePhoneNumberMobileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberNational

> ListAvailablePhoneNumberNationalResponse ListAvailablePhoneNumberNational(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberNational(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberNational``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberNational`: ListAvailablePhoneNumberNationalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberNational`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberNationalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberNationalResponse**](ListAvailablePhoneNumberNationalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberSharedCost

> ListAvailablePhoneNumberSharedCostResponse ListAvailablePhoneNumberSharedCost(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberSharedCost(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberSharedCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberSharedCost`: ListAvailablePhoneNumberSharedCostResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberSharedCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberSharedCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberSharedCostResponse**](ListAvailablePhoneNumberSharedCostResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberTollFree

> ListAvailablePhoneNumberTollFreeResponse ListAvailablePhoneNumberTollFree(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberTollFree(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberTollFree`: ListAvailablePhoneNumberTollFreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberTollFreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberTollFreeResponse**](ListAvailablePhoneNumberTollFreeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberVoip

> ListAvailablePhoneNumberVoipResponse ListAvailablePhoneNumberVoip(ctx, accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    countryCode := "countryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    areaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    contains := "contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    smsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    mmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    voiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    excludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    excludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    nearNumber := "nearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    nearLatLong := "nearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    inPostalCode := "inPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRegion := "inRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    inRateCenter := "inRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    inLata := "inLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    inLocality := "inLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    faxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAvailablePhoneNumberVoip(context.Background(), accountSid, countryCode).AreaCode(areaCode).Contains(contains).SmsEnabled(smsEnabled).MmsEnabled(mmsEnabled).VoiceEnabled(voiceEnabled).ExcludeAllAddressRequired(excludeAllAddressRequired).ExcludeLocalAddressRequired(excludeLocalAddressRequired).ExcludeForeignAddressRequired(excludeForeignAddressRequired).Beta(beta).NearNumber(nearNumber).NearLatLong(nearLatLong).Distance(distance).InPostalCode(inPostalCode).InRegion(inRegion).InRateCenter(inRateCenter).InLata(inLata).InLocality(inLocality).FaxEnabled(faxEnabled).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberVoip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberVoip`: ListAvailablePhoneNumberVoipResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**countryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberVoipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **areaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **smsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **mmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **excludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **excludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **nearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **nearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **inPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **inRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **inRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **inLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **inLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **faxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailablePhoneNumberVoipResponse**](ListAvailablePhoneNumberVoipResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCall

> ListCallResponse ListCall(ctx, accountSid).To(to).From(from).ParentCallSid(parentCallSid).Status(status).StartTime(startTime).StartTime2(startTime2).StartTime3(startTime3).EndTime(endTime).EndTime2(endTime2).EndTime3(endTime3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read.
    to := "to_example" // string | Only show calls made to this phone number, SIP address, Client identifier or SIM SID. (optional)
    from := "from_example" // string | Only include calls from this phone number, SIP address, Client identifier or SIM SID. (optional)
    parentCallSid := "parentCallSid_example" // string | Only include calls spawned by calls with this SID. (optional)
    status := "status_example" // string | The status of the calls to include. Can be: `queued`, `ringing`, `in-progress`, `canceled`, `completed`, `failed`, `busy`, or `no-answer`. (optional)
    startTime := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    startTime2 := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    startTime3 := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    endTime := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    endTime2 := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    endTime3 := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCall(context.Background(), accountSid).To(to).From(from).ParentCallSid(parentCallSid).Status(status).StartTime(startTime).StartTime2(startTime2).StartTime3(startTime3).EndTime(endTime).EndTime2(endTime2).EndTime3(endTime3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCall`: ListCallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | Only show calls made to this phone number, SIP address, Client identifier or SIM SID. | 
 **from** | **string** | Only include calls from this phone number, SIP address, Client identifier or SIM SID. | 
 **parentCallSid** | **string** | Only include calls spawned by calls with this SID. | 
 **status** | **string** | The status of the calls to include. Can be: &#x60;queued&#x60;, &#x60;ringing&#x60;, &#x60;in-progress&#x60;, &#x60;canceled&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, &#x60;busy&#x60;, or &#x60;no-answer&#x60;. | 
 **startTime** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **startTime2** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **startTime3** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **endTime** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **endTime2** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **endTime3** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCallResponse**](ListCallResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallEvent

> ListCallEventResponse ListCallEvent(ctx, accountSid, callSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique SID identifier of the Account.
    callSid := "callSid_example" // string | The unique SID identifier of the Call.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCallEvent(context.Background(), accountSid, callSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallEvent`: ListCallEventResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique SID identifier of the Account. | 
**callSid** | **string** | The unique SID identifier of the Call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCallEventResponse**](ListCallEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallNotification

> ListCallNotificationResponse ListCallNotification(ctx, accountSid, callSid).Log(log).MessageDate(messageDate).MessageDate2(messageDate2).MessageDate3(messageDate3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resources to read.
    log := int32(56) // int32 | Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read. (optional)
    messageDate := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    messageDate2 := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    messageDate3 := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCallNotification(context.Background(), accountSid, callSid).Log(log).MessageDate(messageDate).MessageDate2(messageDate2).MessageDate3(messageDate3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallNotification`: ListCallNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **log** | **int32** | Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read. | 
 **messageDate** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **messageDate2** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **messageDate3** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCallNotificationResponse**](ListCallNotificationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecording

> ListCallRecordingResponse ListCallRecording(ctx, accountSid, callSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.
    dateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    dateCreated2 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    dateCreated3 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCallRecording(context.Background(), accountSid, callSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallRecording`: ListCallRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **dateCreated2** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **dateCreated3** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCallRecordingResponse**](ListCallRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConference

> ListConferenceResponse ListConference(ctx, accountSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).DateUpdated(dateUpdated).DateUpdated2(dateUpdated2).DateUpdated3(dateUpdated3).FriendlyName(friendlyName).Status(status).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read.
    dateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    dateCreated2 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    dateCreated3 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    dateUpdated := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    dateUpdated2 := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    dateUpdated3 := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    friendlyName := "friendlyName_example" // string | The string that identifies the Conference resources to read. (optional)
    status := "status_example" // string | The status of the resources to read. Can be: `init`, `in-progress`, or `completed`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConference(context.Background(), accountSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).DateUpdated(dateUpdated).DateUpdated2(dateUpdated2).DateUpdated3(dateUpdated3).FriendlyName(friendlyName).Status(status).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConference`: ListConferenceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **dateCreated2** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **dateCreated3** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **dateUpdated** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **dateUpdated2** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **dateUpdated3** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **friendlyName** | **string** | The string that identifies the Conference resources to read. | 
 **status** | **string** | The status of the resources to read. Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListConferenceResponse**](ListConferenceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceRecording

> ListConferenceRecordingResponse ListConferenceRecording(ctx, accountSid, conferenceSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
    conferenceSid := "conferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to read.
    dateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    dateCreated2 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    dateCreated3 := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConferenceRecording(context.Background(), accountSid, conferenceSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConferenceRecording`: ListConferenceRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read. | 
**conferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **dateCreated2** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **dateCreated3** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListConferenceRecordingResponse**](ListConferenceRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectApp

> ListConnectAppResponse ListConnectApp(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConnectApp(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectApp`: ListConnectAppResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListConnectAppResponse**](ListConnectAppResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDependentPhoneNumber

> ListDependentPhoneNumberResponse ListDependentPhoneNumber(ctx, accountSid, addressSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
    addressSid := "addressSid_example" // string | The SID of the Address resource associated with the phone number.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListDependentPhoneNumber(context.Background(), accountSid, addressSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDependentPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentPhoneNumber`: ListDependentPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDependentPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read. | 
**addressSid** | **string** | The SID of the Address resource associated with the phone number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDependentPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDependentPhoneNumberResponse**](ListDependentPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumber

> ListIncomingPhoneNumberResponse ListIncomingPhoneNumber(ctx, accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read.
    beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    friendlyName := "friendlyName_example" // string | A string that identifies the IncomingPhoneNumber resources to read. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    origin := "origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumber(context.Background(), accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumber`: ListIncomingPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **friendlyName** | **string** | A string that identifies the IncomingPhoneNumber resources to read. | 
 **phoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberResponse**](ListIncomingPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOn

> ListIncomingPhoneNumberAssignedAddOnResponse ListIncomingPhoneNumberAssignedAddOn(ctx, accountSid, resourceSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumberAssignedAddOn(context.Background(), accountSid, resourceSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberAssignedAddOn`: ListIncomingPhoneNumberAssignedAddOnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
**resourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberAssignedAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberAssignedAddOnResponse**](ListIncomingPhoneNumberAssignedAddOnResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOnExtension

> ListIncomingPhoneNumberAssignedAddOnExtensionResponse ListIncomingPhoneNumberAssignedAddOnExtension(ctx, accountSid, resourceSid, assignedAddOnSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    resourceSid := "resourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    assignedAddOnSid := "assignedAddOnSid_example" // string | The SID that uniquely identifies the assigned Add-on installation.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension(context.Background(), accountSid, resourceSid, assignedAddOnSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberAssignedAddOnExtension`: ListIncomingPhoneNumberAssignedAddOnExtensionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
**resourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned. | 
**assignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberAssignedAddOnExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberAssignedAddOnExtensionResponse**](ListIncomingPhoneNumberAssignedAddOnExtensionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberLocal

> ListIncomingPhoneNumberLocalResponse ListIncomingPhoneNumberLocal(ctx, accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    friendlyName := "friendlyName_example" // string | A string that identifies the resources to read. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    origin := "origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumberLocal(context.Background(), accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberLocal`: ListIncomingPhoneNumberLocalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberLocalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **friendlyName** | **string** | A string that identifies the resources to read. | 
 **phoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberLocalResponse**](ListIncomingPhoneNumberLocalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberMobile

> ListIncomingPhoneNumberMobileResponse ListIncomingPhoneNumberMobile(ctx, accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    friendlyName := "friendlyName_example" // string | A string that identifies the resources to read. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    origin := "origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumberMobile(context.Background(), accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberMobile`: ListIncomingPhoneNumberMobileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberMobileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **friendlyName** | **string** | A string that identifies the resources to read. | 
 **phoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberMobileResponse**](ListIncomingPhoneNumberMobileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberTollFree

> ListIncomingPhoneNumberTollFreeResponse ListIncomingPhoneNumberTollFree(ctx, accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    friendlyName := "friendlyName_example" // string | A string that identifies the resources to read. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    origin := "origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIncomingPhoneNumberTollFree(context.Background(), accountSid).Beta(beta).FriendlyName(friendlyName).PhoneNumber(phoneNumber).Origin(origin).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberTollFree`: ListIncomingPhoneNumberTollFreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingPhoneNumberTollFreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **friendlyName** | **string** | A string that identifies the resources to read. | 
 **phoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIncomingPhoneNumberTollFreeResponse**](ListIncomingPhoneNumberTollFreeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKey

> ListKeyResponse ListKey(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListKey(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKey`: ListKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListKeyResponse**](ListKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMedia

> ListMediaResponse ListMedia(ctx, accountSid, messageSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read.
    messageSid := "messageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    dateCreated := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    dateCreated2 := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    dateCreated3 := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMedia(context.Background(), accountSid, messageSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMedia`: ListMediaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read. | 
**messageSid** | **string** | The SID of the Message resource that this Media resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateCreated** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **dateCreated2** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **dateCreated3** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMediaResponse**](ListMediaResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ListMemberResponse ListMember(ctx, accountSid, queueSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read.
    queueSid := "queueSid_example" // string | The SID of the Queue in which to find the members
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMember(context.Background(), accountSid, queueSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMember`: ListMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read. | 
**queueSid** | **string** | The SID of the Queue in which to find the members | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ListMessageResponse ListMessage(ctx, accountSid).To(to).From(from).DateSent(dateSent).DateSent2(dateSent2).DateSent3(dateSent3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read.
    to := "to_example" // string | Read messages sent to only this phone number. (optional)
    from := "from_example" // string | Read messages sent from only this phone number or alphanumeric sender ID. (optional)
    dateSent := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    dateSent2 := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    dateSent3 := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMessage(context.Background(), accountSid).To(to).From(from).DateSent(dateSent).DateSent2(dateSent2).DateSent3(dateSent3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessage`: ListMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | Read messages sent to only this phone number. | 
 **from** | **string** | Read messages sent from only this phone number or alphanumeric sender ID. | 
 **dateSent** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **dateSent2** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **dateSent3** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMessageResponse**](ListMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotification

> ListNotificationResponse ListNotification(ctx, accountSid).Log(log).MessageDate(messageDate).MessageDate2(messageDate2).MessageDate3(messageDate3).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read.
    log := int32(56) // int32 | Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read. (optional)
    messageDate := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    messageDate2 := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    messageDate3 := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListNotification(context.Background(), accountSid).Log(log).MessageDate(messageDate).MessageDate2(messageDate2).MessageDate3(messageDate3).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotification`: ListNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **log** | **int32** | Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read. | 
 **messageDate** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **messageDate2** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **messageDate3** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListNotificationResponse**](ListNotificationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutgoingCallerId

> ListOutgoingCallerIdResponse ListOutgoingCallerId(ctx, accountSid).PhoneNumber(phoneNumber).FriendlyName(friendlyName).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read.
    phoneNumber := "phoneNumber_example" // string | The phone number of the OutgoingCallerId resources to read. (optional)
    friendlyName := "friendlyName_example" // string | The string that identifies the OutgoingCallerId resources to read. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListOutgoingCallerId(context.Background(), accountSid).PhoneNumber(phoneNumber).FriendlyName(friendlyName).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOutgoingCallerId`: ListOutgoingCallerIdResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOutgoingCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumber** | **string** | The phone number of the OutgoingCallerId resources to read. | 
 **friendlyName** | **string** | The string that identifies the OutgoingCallerId resources to read. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListOutgoingCallerIdResponse**](ListOutgoingCallerIdResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> ListParticipantResponse ListParticipant(ctx, accountSid, conferenceSid).Muted(muted).Hold(hold).Coaching(coaching).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read.
    conferenceSid := "conferenceSid_example" // string | The SID of the conference with the participants to read.
    muted := true // bool | Whether to return only participants that are muted. Can be: `true` or `false`. (optional)
    hold := true // bool | Whether to return only participants that are on hold. Can be: `true` or `false`. (optional)
    coaching := true // bool | Whether to return only participants who are coaching another call. Can be: `true` or `false`. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListParticipant(context.Background(), accountSid, conferenceSid).Muted(muted).Hold(hold).Coaching(coaching).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParticipant`: ListParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read. | 
**conferenceSid** | **string** | The SID of the conference with the participants to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **muted** | **bool** | Whether to return only participants that are muted. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **hold** | **bool** | Whether to return only participants that are on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **coaching** | **bool** | Whether to return only participants who are coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListParticipantResponse**](ListParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueue

> ListQueueResponse ListQueue(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListQueue(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueue`: ListQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListQueueResponse**](ListQueueResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecording

> ListRecordingResponse ListRecording(ctx, accountSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).CallSid(callSid).ConferenceSid(conferenceSid).IncludeSoftDeleted(includeSoftDeleted).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
    dateCreated := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    dateCreated2 := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    dateCreated3 := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. (optional)
    conferenceSid := "conferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to read. (optional)
    includeSoftDeleted := true // bool | A boolean parameter indicating whether to retrieve soft deleted recordings or not. Recordings metadata are kept after deletion for a retention period of 40 days. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRecording(context.Background(), accountSid).DateCreated(dateCreated).DateCreated2(dateCreated2).DateCreated3(dateCreated3).CallSid(callSid).ConferenceSid(conferenceSid).IncludeSoftDeleted(includeSoftDeleted).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecording`: ListRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateCreated** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **dateCreated2** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **dateCreated3** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. | 
 **conferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to read. | 
 **includeSoftDeleted** | **bool** | A boolean parameter indicating whether to retrieve soft deleted recordings or not. Recordings metadata are kept after deletion for a retention period of 40 days. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRecordingResponse**](ListRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResult

> ListRecordingAddOnResultResponse ListRecordingAddOnResult(ctx, accountSid, referenceSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the result to read belongs.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRecordingAddOnResult(context.Background(), accountSid, referenceSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingAddOnResult`: ListRecordingAddOnResultResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingAddOnResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read. | 
**referenceSid** | **string** | The SID of the recording to which the result to read belongs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordingAddOnResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRecordingAddOnResultResponse**](ListRecordingAddOnResultResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResultPayload

> ListRecordingAddOnResultPayloadResponse ListRecordingAddOnResultPayload(ctx, accountSid, referenceSid, addOnResultSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
    referenceSid := "referenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs.
    addOnResultSid := "addOnResultSid_example" // string | The SID of the AddOnResult to which the payloads to read belongs.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRecordingAddOnResultPayload(context.Background(), accountSid, referenceSid, addOnResultSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingAddOnResultPayload`: ListRecordingAddOnResultPayloadResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingAddOnResultPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read. | 
**referenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs. | 
**addOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to read belongs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordingAddOnResultPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRecordingAddOnResultPayloadResponse**](ListRecordingAddOnResultPayloadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingTranscription

> ListRecordingTranscriptionResponse ListRecordingTranscription(ctx, accountSid, recordingSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
    recordingSid := "recordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRecordingTranscription(context.Background(), accountSid, recordingSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingTranscription`: ListRecordingTranscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read. | 
**recordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRecordingTranscriptionResponse**](ListRecordingTranscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> ListShortCodeResponse ListShortCode(ctx, accountSid).FriendlyName(friendlyName).ShortCode(shortCode).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.
    friendlyName := "friendlyName_example" // string | The string that identifies the ShortCode resources to read. (optional)
    shortCode := "shortCode_example" // string | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListShortCode(context.Background(), accountSid).FriendlyName(friendlyName).ShortCode(shortCode).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortCode`: ListShortCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | The string that identifies the ShortCode resources to read. | 
 **shortCode** | **string** | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListShortCodeResponse**](ListShortCodeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSigningKey

> ListSigningKeyResponse ListSigningKey(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | 
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSigningKey(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSigningKey`: ListSigningKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSigningKeyResponse**](ListSigningKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsCredentialListMapping

> ListSipAuthCallsCredentialListMappingResponse ListSipAuthCallsCredentialListMapping(ctx, accountSid, domainSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipAuthCallsCredentialListMapping(context.Background(), accountSid, domainSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthCallsCredentialListMapping`: ListSipAuthCallsCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipAuthCallsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipAuthCallsCredentialListMappingResponse**](ListSipAuthCallsCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsIpAccessControlListMapping

> ListSipAuthCallsIpAccessControlListMappingResponse ListSipAuthCallsIpAccessControlListMapping(ctx, accountSid, domainSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipAuthCallsIpAccessControlListMapping(context.Background(), accountSid, domainSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthCallsIpAccessControlListMapping`: ListSipAuthCallsIpAccessControlListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipAuthCallsIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipAuthCallsIpAccessControlListMappingResponse**](ListSipAuthCallsIpAccessControlListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthRegistrationsCredentialListMapping

> ListSipAuthRegistrationsCredentialListMappingResponse ListSipAuthRegistrationsCredentialListMapping(ctx, accountSid, domainSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
    domainSid := "domainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipAuthRegistrationsCredentialListMapping(context.Background(), accountSid, domainSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthRegistrationsCredentialListMapping`: ListSipAuthRegistrationsCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read. | 
**domainSid** | **string** | The SID of the SIP domain that contains the resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipAuthRegistrationsCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipAuthRegistrationsCredentialListMappingResponse**](ListSipAuthRegistrationsCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredential

> ListSipCredentialResponse ListSipCredential(ctx, accountSid, credentialListSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    credentialListSid := "credentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credentials.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipCredential(context.Background(), accountSid, credentialListSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredential`: ListSipCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**credentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipCredentialResponse**](ListSipCredentialResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialList

> ListSipCredentialListResponse ListSipCredentialList(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipCredentialList(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredentialList`: ListSipCredentialListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipCredentialListResponse**](ListSipCredentialListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialListMapping

> ListSipCredentialListMappingResponse ListSipCredentialListMapping(ctx, accountSid, domainSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipCredentialListMapping(context.Background(), accountSid, domainSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredentialListMapping`: ListSipCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipCredentialListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipCredentialListMappingResponse**](ListSipCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipDomain

> ListSipDomainResponse ListSipDomain(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipDomain(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipDomain`: ListSipDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipDomainResponse**](ListSipDomainResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlList

> ListSipIpAccessControlListResponse ListSipIpAccessControlList(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipIpAccessControlList(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAccessControlList`: ListSipIpAccessControlListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipIpAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipIpAccessControlListResponse**](ListSipIpAccessControlListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlListMapping

> ListSipIpAccessControlListMappingResponse ListSipIpAccessControlListMapping(ctx, accountSid, domainSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    domainSid := "domainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipIpAccessControlListMapping(context.Background(), accountSid, domainSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAccessControlListMapping`: ListSipIpAccessControlListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**domainSid** | **string** | A 34 character string that uniquely identifies the SIP domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipIpAccessControlListMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipIpAccessControlListMappingResponse**](ListSipIpAccessControlListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAddress

> ListSipIpAddressResponse ListSipIpAddress(ctx, accountSid, ipAccessControlListSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSipIpAddress(context.Background(), accountSid, ipAccessControlListSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAddress`: ListSipIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**ipAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSipIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSipIpAddressResponse**](ListSipIpAddressResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscription

> ListTranscriptionResponse ListTranscription(ctx, accountSid).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListTranscription(context.Background(), accountSid).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTranscription`: ListTranscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListTranscriptionResponse**](ListTranscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> ListUsageRecordResponse ListUsageRecord(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecord(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecord`: ListUsageRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordResponse**](ListUsageRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordAllTime

> ListUsageRecordAllTimeResponse ListUsageRecordAllTime(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordAllTime(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordAllTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordAllTime`: ListUsageRecordAllTimeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordAllTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordAllTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordAllTimeResponse**](ListUsageRecordAllTimeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordDaily

> ListUsageRecordDailyResponse ListUsageRecordDaily(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordDaily(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordDaily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordDaily`: ListUsageRecordDailyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordDaily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordDailyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordDailyResponse**](ListUsageRecordDailyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordLastMonth

> ListUsageRecordLastMonthResponse ListUsageRecordLastMonth(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordLastMonth(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordLastMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordLastMonth`: ListUsageRecordLastMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordLastMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordLastMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordLastMonthResponse**](ListUsageRecordLastMonthResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordMonthly

> ListUsageRecordMonthlyResponse ListUsageRecordMonthly(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordMonthly(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordMonthly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordMonthly`: ListUsageRecordMonthlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordMonthly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordMonthlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordMonthlyResponse**](ListUsageRecordMonthlyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordThisMonth

> ListUsageRecordThisMonthResponse ListUsageRecordThisMonth(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordThisMonth(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordThisMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordThisMonth`: ListUsageRecordThisMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordThisMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordThisMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordThisMonthResponse**](ListUsageRecordThisMonthResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordToday

> ListUsageRecordTodayResponse ListUsageRecordToday(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordToday(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordToday``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordToday`: ListUsageRecordTodayResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordToday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordTodayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordTodayResponse**](ListUsageRecordTodayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordYearly

> ListUsageRecordYearlyResponse ListUsageRecordYearly(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordYearly(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordYearly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordYearly`: ListUsageRecordYearlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordYearly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordYearlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordYearlyResponse**](ListUsageRecordYearlyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordYesterday

> ListUsageRecordYesterdayResponse ListUsageRecordYesterday(ctx, accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    category := "category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    startDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    endDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    includeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageRecordYesterday(context.Background(), accountSid).Category(category).StartDate(startDate).EndDate(endDate).IncludeSubaccounts(includeSubaccounts).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordYesterday``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordYesterday`: ListUsageRecordYesterdayResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordYesterday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRecordYesterdayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **startDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **endDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **includeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordYesterdayResponse**](ListUsageRecordYesterdayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageTrigger

> ListUsageTriggerResponse ListUsageTrigger(ctx, accountSid).Recurring(recurring).TriggerBy(triggerBy).UsageCategory(usageCategory).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read.
    recurring := "recurring_example" // string | The frequency of recurring UsageTriggers to read. Can be: `daily`, `monthly`, or `yearly` to read recurring UsageTriggers. An empty value or a value of `alltime` reads non-recurring UsageTriggers. (optional)
    triggerBy := "triggerBy_example" // string | The trigger field of the UsageTriggers to read.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price). (optional)
    usageCategory := "usageCategory_example" // string | The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories). (optional)
    pageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsageTrigger(context.Background(), accountSid).Recurring(recurring).TriggerBy(triggerBy).UsageCategory(usageCategory).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageTrigger`: ListUsageTriggerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsageTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurring** | **string** | The frequency of recurring UsageTriggers to read. Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; to read recurring UsageTriggers. An empty value or a value of &#x60;alltime&#x60; reads non-recurring UsageTriggers. | 
 **triggerBy** | **string** | The trigger field of the UsageTriggers to read.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price). | 
 **usageCategory** | **string** | The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories). | 
 **pageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageTriggerResponse**](ListUsageTriggerResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> ApiV2010Account UpdateAccount(ctx, sid).FriendlyName(friendlyName).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sid := "sid_example" // string | The Account Sid that uniquely identifies the account to update
    friendlyName := "friendlyName_example" // string | Update the human-readable description of this Account (optional)
    status := "status_example" // string | Alter the status of this account: use `closed` to irreversibly close this account, `suspended` to temporarily suspend it, or `active` to reactivate it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAccount(context.Background(), sid).FriendlyName(friendlyName).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | The Account Sid that uniquely identifies the account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **string** | Update the human-readable description of this Account | 
 **status** | **string** | Alter the status of this account: use &#x60;closed&#x60; to irreversibly close this account, &#x60;suspended&#x60; to temporarily suspend it, or &#x60;active&#x60; to reactivate it. | 

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> ApiV2010AccountAddress UpdateAddress(ctx, accountSid, sid).AutoCorrectAddress(autoCorrectAddress).City(city).CustomerName(customerName).EmergencyEnabled(emergencyEnabled).FriendlyName(friendlyName).PostalCode(postalCode).Region(region).Street(street).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to update.
    autoCorrectAddress := true // bool | Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide. (optional)
    city := "city_example" // string | The city of the address. (optional)
    customerName := "customerName_example" // string | The name to associate with the address. (optional)
    emergencyEnabled := true // bool | Whether to enable emergency calling on the address. Can be: `true` or `false`. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the address. It can be up to 64 characters long. (optional)
    postalCode := "postalCode_example" // string | The postal code of the address. (optional)
    region := "region_example" // string | The state or region of the address. (optional)
    street := "street_example" // string | The number and street address of the address. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAddress(context.Background(), accountSid, sid).AutoCorrectAddress(autoCorrectAddress).City(city).CustomerName(customerName).EmergencyEnabled(emergencyEnabled).FriendlyName(friendlyName).PostalCode(postalCode).Region(region).Street(street).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide. | 
 **city** | **string** | The city of the address. | 
 **customerName** | **string** | The name to associate with the address. | 
 **emergencyEnabled** | **bool** | Whether to enable emergency calling on the address. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the address. It can be up to 64 characters long. | 
 **postalCode** | **string** | The postal code of the address. | 
 **region** | **string** | The state or region of the address. | 
 **street** | **string** | The number and street address of the address. | 

### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiV2010AccountApplication UpdateApplication(ctx, accountSid, sid).ApiVersion(apiVersion).FriendlyName(friendlyName).MessageStatusCallback(messageStatusCallback).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsStatusCallback(smsStatusCallback).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to update.
    apiVersion := "apiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is your account's default API version. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    messageStatusCallback := "messageStatusCallback_example" // string | The URL we should call using a POST method to send message status information to your application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. (optional)
    smsStatusCallback := "smsStatusCallback_example" // string | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`. (optional)
    voiceCallerIdLookup := true // bool | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL we should call when the phone number assigned to this application receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateApplication(context.Background(), accountSid, sid).ApiVersion(apiVersion).FriendlyName(friendlyName).MessageStatusCallback(messageStatusCallback).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsStatusCallback(smsStatusCallback).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is your account&#39;s default API version. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **messageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application. | 
 **smsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsStatusCallback** | **string** | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility. | 
 **smsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceCallerIdLookup** | **bool** | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call. | 

### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCall

> ApiV2010AccountCall UpdateCall(ctx, accountSid, sid).FallbackMethod(fallbackMethod).FallbackUrl(fallbackUrl).Method(method).Status(status).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Twiml(twiml).Url(url).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Call resource to update
    fallbackMethod := "fallbackMethod_example" // string | The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    fallbackUrl := "fallbackUrl_example" // string | The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    method := "method_example" // string | The HTTP method we should use when calling the `url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    status := "status_example" // string | The new status of the resource. Can be: `canceled` or `completed`. Specifying `canceled` will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying `completed` will attempt to hang up a call even if it's already in progress. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use when requesting the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    timeLimit := int32(56) // int32 | The maximum duration of the call in seconds. Constraints depend on account and configuration. (optional)
    twiml := "twiml_example" // string | TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive (optional)
    url := "url_example" // string | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCall(context.Background(), accountSid, sid).FallbackMethod(fallbackMethod).FallbackUrl(fallbackUrl).Method(method).Status(status).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TimeLimit(timeLimit).Twiml(twiml).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Call resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fallbackMethod** | **string** | The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **fallbackUrl** | **string** | The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **method** | **string** | The HTTP method we should use when calling the &#x60;url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **status** | **string** | The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;completed&#x60;. Specifying &#x60;canceled&#x60; will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying &#x60;completed&#x60; will attempt to hang up a call even if it&#39;s already in progress. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). | 
 **statusCallbackMethod** | **string** | The HTTP method we should use when requesting the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **timeLimit** | **int32** | The maximum duration of the call in seconds. Constraints depend on account and configuration. | 
 **twiml** | **string** | TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive | 
 **url** | **string** | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). | 

### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallFeedback

> ApiV2010AccountCallCallFeedback UpdateCallFeedback(ctx, accountSid, callSid).Issue(issue).QualityScore(qualityScore).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    callSid := "callSid_example" // string | The call sid that uniquely identifies the call
    issue := []string{"Inner_example"} // []string | One or more issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, `unsolicited-call`, or `one-way-audio`. (optional)
    qualityScore := int32(56) // int32 | The call quality expressed as an integer from `1` to `5` where `1` represents very poor call quality and `5` represents a perfect call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCallFeedback(context.Background(), accountSid, callSid).Issue(issue).QualityScore(qualityScore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCallFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCallFeedback`: ApiV2010AccountCallCallFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCallFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**callSid** | **string** | The call sid that uniquely identifies the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issue** | **[]string** | One or more issues experienced during the call. The issues can be: &#x60;imperfect-audio&#x60;, &#x60;dropped-call&#x60;, &#x60;incorrect-caller-id&#x60;, &#x60;post-dial-delay&#x60;, &#x60;digits-not-captured&#x60;, &#x60;audio-latency&#x60;, &#x60;unsolicited-call&#x60;, or &#x60;one-way-audio&#x60;. | 
 **qualityScore** | **int32** | The call quality expressed as an integer from &#x60;1&#x60; to &#x60;5&#x60; where &#x60;1&#x60; represents very poor call quality and &#x60;5&#x60; represents a perfect call. | 

### Return type

[**ApiV2010AccountCallCallFeedback**](ApiV2010AccountCallCallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallRecording

> ApiV2010AccountCallCallRecording UpdateCallRecording(ctx, accountSid, callSid, sid).Status(status).PauseBehavior(pauseBehavior).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to update.
    status := "status_example" // string | The new status of the recording. Can be: `stopped`, `paused`, `in-progress`.
    pauseBehavior := "pauseBehavior_example" // string | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCallRecording(context.Background(), accountSid, callSid, sid).Status(status).PauseBehavior(pauseBehavior).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;. | 
 **pauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;. | 

### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConference

> ApiV2010AccountConference UpdateConference(ctx, accountSid, sid).AnnounceMethod(announceMethod).AnnounceUrl(announceUrl).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Conference resource to update
    announceMethod := "announceMethod_example" // string | The HTTP method used to call `announce_url`. Can be: `GET` or `POST` and the default is `POST` (optional)
    announceUrl := "announceUrl_example" // string | The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with `<Play>` or `<Say>`. (optional)
    status := "status_example" // string | The new status of the resource. Can be:  Can be: `init`, `in-progress`, or `completed`. Specifying `completed` will end the conference and hang up all participants (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConference(context.Background(), accountSid, sid).AnnounceMethod(announceMethod).AnnounceUrl(announceUrl).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConference`: ApiV2010AccountConference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Conference resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **announceMethod** | **string** | The HTTP method used to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60; | 
 **announceUrl** | **string** | The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60;. | 
 **status** | **string** | The new status of the resource. Can be:  Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;. Specifying &#x60;completed&#x60; will end the conference and hang up all participants | 

### Return type

[**ApiV2010AccountConference**](ApiV2010AccountConference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConferenceRecording

> ApiV2010AccountConferenceConferenceRecording UpdateConferenceRecording(ctx, accountSid, conferenceSid, sid).Status(status).PauseBehavior(pauseBehavior).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
    conferenceSid := "conferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use `Twilio.CURRENT` to reference the current active recording.
    status := "status_example" // string | The new status of the recording. Can be: `stopped`, `paused`, `in-progress`.
    pauseBehavior := "pauseBehavior_example" // string | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConferenceRecording(context.Background(), accountSid, conferenceSid, sid).Status(status).PauseBehavior(pauseBehavior).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConferenceRecording`: ApiV2010AccountConferenceConferenceRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update. | 
**conferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use &#x60;Twilio.CURRENT&#x60; to reference the current active recording. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;. | 
 **pauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;. | 

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectApp

> ApiV2010AccountConnectApp UpdateConnectApp(ctx, accountSid, sid).AuthorizeRedirectUrl(authorizeRedirectUrl).CompanyName(companyName).DeauthorizeCallbackMethod(deauthorizeCallbackMethod).DeauthorizeCallbackUrl(deauthorizeCallbackUrl).Description(description).FriendlyName(friendlyName).HomepageUrl(homepageUrl).Permissions(permissions).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to update.
    authorizeRedirectUrl := "authorizeRedirectUrl_example" // string | The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App. (optional)
    companyName := "companyName_example" // string | The company name to set for the Connect App. (optional)
    deauthorizeCallbackMethod := "deauthorizeCallbackMethod_example" // string | The HTTP method to use when calling `deauthorize_callback_url`. (optional)
    deauthorizeCallbackUrl := "deauthorizeCallbackUrl_example" // string | The URL to call using the `deauthorize_callback_method` to de-authorize the Connect App. (optional)
    description := "description_example" // string | A description of the Connect App. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    homepageUrl := "homepageUrl_example" // string | A public URL where users can obtain more information about this Connect App. (optional)
    permissions := []string{"Inner_example"} // []string | A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: `get-all` and `post-all`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConnectApp(context.Background(), accountSid, sid).AuthorizeRedirectUrl(authorizeRedirectUrl).CompanyName(companyName).DeauthorizeCallbackMethod(deauthorizeCallbackMethod).DeauthorizeCallbackUrl(deauthorizeCallbackUrl).Description(description).FriendlyName(friendlyName).HomepageUrl(homepageUrl).Permissions(permissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectApp`: ApiV2010AccountConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeRedirectUrl** | **string** | The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App. | 
 **companyName** | **string** | The company name to set for the Connect App. | 
 **deauthorizeCallbackMethod** | **string** | The HTTP method to use when calling &#x60;deauthorize_callback_url&#x60;. | 
 **deauthorizeCallbackUrl** | **string** | The URL to call using the &#x60;deauthorize_callback_method&#x60; to de-authorize the Connect App. | 
 **description** | **string** | A description of the Connect App. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **homepageUrl** | **string** | A public URL where users can obtain more information about this Connect App. | 
 **permissions** | **[]string** | A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: &#x60;get-all&#x60; and &#x60;post-all&#x60;. | 

### Return type

[**ApiV2010AccountConnectApp**](ApiV2010AccountConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber UpdateIncomingPhoneNumber(ctx, accountSid, sid).AccountSid2(accountSid2).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.
    accountSid2 := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). (optional)
    addressSid := "addressSid_example" // string | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations. (optional)
    apiVersion := "apiVersion_example" // string | The API version to use for incoming calls made to the phone number. The default is `2010-04-01`. (optional)
    bundleSid := "bundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    emergencyAddressSid := "emergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from this phone number. (optional)
    emergencyStatus := "emergencyStatus_example" // string | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    identitySid := "identitySid_example" // string | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations. (optional)
    smsApplicationSid := "smsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    statusCallback := "statusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    statusCallbackMethod := "statusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    trunkSid := "trunkSid_example" // string | The SID of the Trunk we should use to handle phone calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    voiceApplicationSid := "voiceApplicationSid_example" // string | The SID of the application we should use to handle phone calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    voiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    voiceReceiveMode := "voiceReceiveMode_example" // string | The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL that we should call to answer a call to the phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateIncomingPhoneNumber(context.Background(), accountSid, sid).AccountSid2(accountSid2).AddressSid(addressSid).ApiVersion(apiVersion).BundleSid(bundleSid).EmergencyAddressSid(emergencyAddressSid).EmergencyStatus(emergencyStatus).FriendlyName(friendlyName).IdentitySid(identitySid).SmsApplicationSid(smsApplicationSid).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).TrunkSid(trunkSid).VoiceApplicationSid(voiceApplicationSid).VoiceCallerIdLookup(voiceCallerIdLookup).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceReceiveMode(voiceReceiveMode).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncomingPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountSid2** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). | 
 **addressSid** | **string** | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations. | 
 **apiVersion** | **string** | The API version to use for incoming calls made to the phone number. The default is &#x60;2010-04-01&#x60;. | 
 **bundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **emergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from this phone number. | 
 **emergencyStatus** | **string** | The parameter displays if emergency calling is enabled for this number. Active numbers may place emergency calls by dialing valid emergency numbers for the country. | 
 **friendlyName** | **string** | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **identitySid** | **string** | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations. | 
 **smsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message. | 
 **statusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **trunkSid** | **string** | The SID of the Trunk we should use to handle phone calls to the phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **voiceApplicationSid** | **string** | The SID of the application we should use to handle phone calls to the phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **voiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **voiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **voiceReceiveMode** | **string** | The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **voiceUrl** | **string** | The URL that we should call to answer a call to the phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> ApiV2010AccountKey UpdateKey(ctx, accountSid, sid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to update.
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateKey(context.Background(), accountSid, sid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKey`: ApiV2010AccountKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountKey**](ApiV2010AccountKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ApiV2010AccountQueueMember UpdateMember(ctx, accountSid, queueSid, callSid).Url(url).Method(method).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update.
    queueSid := "queueSid_example" // string | The SID of the Queue in which to find the members to update.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to update.
    url := "url_example" // string | The absolute URL of the Queue resource.
    method := "method_example" // string | How to pass the update request data. Can be `GET` or `POST` and the default is `POST`. `POST` sends the data as encoded form data and `GET` sends the data as query parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMember(context.Background(), accountSid, queueSid, callSid).Url(url).Method(method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: ApiV2010AccountQueueMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update. | 
**queueSid** | **string** | The SID of the Queue in which to find the members to update. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **url** | **string** | The absolute URL of the Queue resource. | 
 **method** | **string** | How to pass the update request data. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. &#x60;POST&#x60; sends the data as encoded form data and &#x60;GET&#x60; sends the data as query parameters. | 

### Return type

[**ApiV2010AccountQueueMember**](ApiV2010AccountQueueMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ApiV2010AccountMessage UpdateMessage(ctx, accountSid, sid).Body(body).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to update.
    body := "body_example" // string | The text of the message you want to send. Can be up to 1,600 characters long. (optional)
    status := "status_example" // string | When set as `canceled`, allows a message cancelation request if a message has not yet been sent. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMessage(context.Background(), accountSid, sid).Body(body).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | The text of the message you want to send. Can be up to 1,600 characters long. | 
 **status** | **string** | When set as &#x60;canceled&#x60;, allows a message cancelation request if a message has not yet been sent. | 

### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutgoingCallerId

> ApiV2010AccountOutgoingCallerId UpdateOutgoingCallerId(ctx, accountSid, sid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update.
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateOutgoingCallerId(context.Background(), accountSid, sid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOutgoingCallerId`: ApiV2010AccountOutgoingCallerId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutgoingCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountOutgoingCallerId**](ApiV2010AccountOutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParticipant

> ApiV2010AccountConferenceParticipant UpdateParticipant(ctx, accountSid, conferenceSid, callSid).AnnounceMethod(announceMethod).AnnounceUrl(announceUrl).BeepOnExit(beepOnExit).CallSidToCoach(callSidToCoach).Coaching(coaching).EndConferenceOnExit(endConferenceOnExit).Hold(hold).HoldMethod(holdMethod).HoldUrl(holdUrl).Muted(muted).WaitMethod(waitMethod).WaitUrl(waitUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update.
    conferenceSid := "conferenceSid_example" // string | The SID of the conference with the participant to update.
    callSid := "callSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.
    announceMethod := "announceMethod_example" // string | The HTTP method we should use to call `announce_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    announceUrl := "announceUrl_example" // string | The URL we call using the `announce_method` for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains `<Play>` or `<Say>` commands. (optional)
    beepOnExit := true // bool | Whether to play a notification beep to the conference when the participant exits. Can be: `true` or `false`. (optional)
    callSidToCoach := "callSidToCoach_example" // string | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`. (optional)
    coaching := true // bool | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined. (optional)
    endConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`. (optional)
    hold := true // bool | Whether the participant should be on hold. Can be: `true` or `false`. `true` puts the participant on hold, and `false` lets them rejoin the conference. (optional)
    holdMethod := "holdMethod_example" // string | The HTTP method we should use to call `hold_url`. Can be: `GET` or `POST` and the default is `GET`. (optional)
    holdUrl := "holdUrl_example" // string | The URL we call using the `hold_method` for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the `<Play>`, `<Say>` or `<Redirect>` commands. (optional)
    muted := true // bool | Whether the participant should be muted. Can be `true` or `false`. `true` will mute the participant, and `false` will un-mute them. Anything value other than `true` or `false` is interpreted as `false`. (optional)
    waitMethod := "waitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    waitUrl := "waitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateParticipant(context.Background(), accountSid, conferenceSid, callSid).AnnounceMethod(announceMethod).AnnounceUrl(announceUrl).BeepOnExit(beepOnExit).CallSidToCoach(callSidToCoach).Coaching(coaching).EndConferenceOnExit(endConferenceOnExit).Hold(hold).HoldMethod(holdMethod).HoldUrl(holdUrl).Muted(muted).WaitMethod(waitMethod).WaitUrl(waitUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update. | 
**conferenceSid** | **string** | The SID of the conference with the participant to update. | 
**callSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **announceMethod** | **string** | The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **announceUrl** | **string** | The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands. | 
 **beepOnExit** | **bool** | Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **callSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;. | 
 **coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined. | 
 **endConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **hold** | **bool** | Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference. | 
 **holdMethod** | **string** | The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;. | 
 **holdUrl** | **string** | The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands. | 
 **muted** | **bool** | Whether the participant should be muted. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; will mute the participant, and &#x60;false&#x60; will un-mute them. Anything value other than &#x60;true&#x60; or &#x60;false&#x60; is interpreted as &#x60;false&#x60;. | 
 **waitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **waitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 

### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayments

> ApiV2010AccountCallPayments UpdatePayments(ctx, accountSid, callSid, sid).IdempotencyKey(idempotencyKey).StatusCallback(statusCallback).Capture(capture).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource.
    callSid := "callSid_example" // string | The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource.
    sid := "sid_example" // string | The SID of Payments session that needs to be updated.
    idempotencyKey := "idempotencyKey_example" // string | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
    statusCallback := "statusCallback_example" // string | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests.
    capture := "capture_example" // string | The piece of payment information that you wish the caller to enter. Must be one of `payment-card-number`, `expiration-date`, `security-code`, `postal-code`, `bank-routing-number`, or `bank-account-number`. (optional)
    status := "status_example" // string | Indicates whether the current payment session should be cancelled or completed. When `cancel` the payment session is cancelled. When `complete`, Twilio sends the payment information to the selected <Pay> connector for processing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdatePayments(context.Background(), accountSid, callSid, sid).IdempotencyKey(idempotencyKey).StatusCallback(statusCallback).Capture(capture).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePayments`: ApiV2010AccountCallPayments
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource. | 
**callSid** | **string** | The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource. | 
**sid** | **string** | The SID of Payments session that needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **idempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. | 
 **statusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests. | 
 **capture** | **string** | The piece of payment information that you wish the caller to enter. Must be one of &#x60;payment-card-number&#x60;, &#x60;expiration-date&#x60;, &#x60;security-code&#x60;, &#x60;postal-code&#x60;, &#x60;bank-routing-number&#x60;, or &#x60;bank-account-number&#x60;. | 
 **status** | **string** | Indicates whether the current payment session should be cancelled or completed. When &#x60;cancel&#x60; the payment session is cancelled. When &#x60;complete&#x60;, Twilio sends the payment information to the selected &lt;Pay&gt; connector for processing. | 

### Return type

[**ApiV2010AccountCallPayments**](ApiV2010AccountCallPayments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQueue

> ApiV2010AccountQueue UpdateQueue(ctx, accountSid, sid).FriendlyName(friendlyName).MaxSize(maxSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to update
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long. (optional)
    maxSize := int32(56) // int32 | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateQueue(context.Background(), accountSid, sid).FriendlyName(friendlyName).MaxSize(maxSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. | 
 **maxSize** | **int32** | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. | 

### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ApiV2010AccountShortCode UpdateShortCode(ctx, accountSid, sid).ApiVersion(apiVersion).FriendlyName(friendlyName).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to update
    apiVersion := "apiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the `FriendlyName` is the short code. (optional)
    smsFallbackMethod := "smsFallbackMethod_example" // string | The HTTP method that we should use to call the `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    smsFallbackUrl := "smsFallbackUrl_example" // string | The URL that we should call if an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    smsMethod := "smsMethod_example" // string | The HTTP method we should use when calling the `sms_url`. Can be: `GET` or `POST`. (optional)
    smsUrl := "smsUrl_example" // string | The URL we should call when receiving an incoming SMS message to this short code. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateShortCode(context.Background(), accountSid, sid).ApiVersion(apiVersion).FriendlyName(friendlyName).SmsFallbackMethod(smsFallbackMethod).SmsFallbackUrl(smsFallbackUrl).SmsMethod(smsMethod).SmsUrl(smsUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShortCode`: ApiV2010AccountShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. | 
 **friendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the &#x60;FriendlyName&#x60; is the short code. | 
 **smsFallbackMethod** | **string** | The HTTP method that we should use to call the &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **string** | The URL that we should call if an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **smsMethod** | **string** | The HTTP method we should use when calling the &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **smsUrl** | **string** | The URL we should call when receiving an incoming SMS message to this short code. | 

### Return type

[**ApiV2010AccountShortCode**](ApiV2010AccountShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSigningKey

> ApiV2010AccountSigningKey UpdateSigningKey(ctx, accountSid, sid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | 
    sid := "sid_example" // string | 
    friendlyName := "friendlyName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSigningKey(context.Background(), accountSid, sid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSigningKey`: ApiV2010AccountSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** |  | 

### Return type

[**ApiV2010AccountSigningKey**](ApiV2010AccountSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential UpdateSipCredential(ctx, accountSid, credentialListSid, sid).Password(password).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    credentialListSid := "credentialListSid_example" // string | The unique id that identifies the credential list that includes this credential.
    sid := "sid_example" // string | The unique id that identifies the resource to update.
    password := "password_example" // string | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSipCredential(context.Background(), accountSid, credentialListSid, sid).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**credentialListSid** | **string** | The unique id that identifies the credential list that includes this credential. | 
**sid** | **string** | The unique id that identifies the resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;) | 

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredentialList

> ApiV2010AccountSipSipCredentialList UpdateSipCredentialList(ctx, accountSid, sid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the Account that is responsible for this resource.
    sid := "sid_example" // string | The credential list Sid that uniquely identifies this resource
    friendlyName := "friendlyName_example" // string | A human readable descriptive text for a CredentialList, up to 64 characters long.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSipCredentialList(context.Background(), accountSid, sid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the Account that is responsible for this resource. | 
**sid** | **string** | The credential list Sid that uniquely identifies this resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A human readable descriptive text for a CredentialList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipDomain

> ApiV2010AccountSipSipDomain UpdateSipDomain(ctx, accountSid, sid).ByocTrunkSid(byocTrunkSid).DomainName(domainName).EmergencyCallerSid(emergencyCallerSid).EmergencyCallingEnabled(emergencyCallingEnabled).FriendlyName(friendlyName).Secure(secure).SipRegistration(sipRegistration).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceStatusCallbackMethod(voiceStatusCallbackMethod).VoiceStatusCallbackUrl(voiceStatusCallbackUrl).VoiceUrl(voiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to update.
    byocTrunkSid := "byocTrunkSid_example" // string | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. (optional)
    domainName := "domainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`. (optional)
    emergencyCallerSid := "emergencyCallerSid_example" // string | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. (optional)
    emergencyCallingEnabled := true // bool | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you created to describe the resource. It can be up to 64 characters long. (optional)
    secure := true // bool | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. (optional)
    sipRegistration := true // bool | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not. (optional)
    voiceFallbackMethod := "voiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    voiceFallbackUrl := "voiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`. (optional)
    voiceMethod := "voiceMethod_example" // string | The HTTP method we should use to call `voice_url` (optional)
    voiceStatusCallbackMethod := "voiceStatusCallbackMethod_example" // string | The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`. (optional)
    voiceStatusCallbackUrl := "voiceStatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    voiceUrl := "voiceUrl_example" // string | The URL we should call when the domain receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSipDomain(context.Background(), accountSid, sid).ByocTrunkSid(byocTrunkSid).DomainName(domainName).EmergencyCallerSid(emergencyCallerSid).EmergencyCallingEnabled(emergencyCallingEnabled).FriendlyName(friendlyName).Secure(secure).SipRegistration(sipRegistration).VoiceFallbackMethod(voiceFallbackMethod).VoiceFallbackUrl(voiceFallbackUrl).VoiceMethod(voiceMethod).VoiceStatusCallbackMethod(voiceStatusCallbackMethod).VoiceStatusCallbackUrl(voiceStatusCallbackUrl).VoiceUrl(voiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **byocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. | 
 **domainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;. | 
 **emergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. | 
 **emergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. | 
 **friendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long. | 
 **secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. | 
 **sipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not. | 
 **voiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | 
 **voiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60; | 
 **voiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **voiceUrl** | **string** | The URL we should call when the domain receives a call. | 

### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList UpdateSipIpAccessControlList(ctx, accountSid, sid).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    sid := "sid_example" // string | A 34 character string that uniquely identifies the resource to udpate.
    friendlyName := "friendlyName_example" // string | A human readable descriptive text, up to 64 characters long.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSipIpAccessControlList(context.Background(), accountSid, sid).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**sid** | **string** | A 34 character string that uniquely identifies the resource to udpate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipIpAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **string** | A human readable descriptive text, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress UpdateSipIpAddress(ctx, accountSid, ipAccessControlListSid, sid).CidrPrefixLength(cidrPrefixLength).FriendlyName(friendlyName).IpAddress(ipAddress).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    ipAccessControlListSid := "ipAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to update.
    sid := "sid_example" // string | A 34 character string that identifies the IpAddress resource to update.
    cidrPrefixLength := int32(56) // int32 | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. (optional)
    friendlyName := "friendlyName_example" // string | A human readable descriptive text for this resource, up to 64 characters long. (optional)
    ipAddress := "ipAddress_example" // string | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSipIpAddress(context.Background(), accountSid, ipAccessControlListSid, sid).CidrPrefixLength(cidrPrefixLength).FriendlyName(friendlyName).IpAddress(ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**ipAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to update. | 
**sid** | **string** | A 34 character string that identifies the IpAddress resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cidrPrefixLength** | **int32** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. | 
 **friendlyName** | **string** | A human readable descriptive text for this resource, up to 64 characters long. | 
 **ipAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiprec

> ApiV2010AccountCallSiprec UpdateSiprec(ctx, accountSid, callSid, sid).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Siprec resource.
    callSid := "callSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Siprec resource is associated with.
    sid := "sid_example" // string | The SID of the Siprec resource, or the `name` used when creating the resource
    status := "status_example" // string | The status. Must have the value `stopped`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSiprec(context.Background(), accountSid, callSid, sid).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSiprec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiprec`: ApiV2010AccountCallSiprec
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSiprec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Siprec resource. | 
**callSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Siprec resource is associated with. | 
**sid** | **string** | The SID of the Siprec resource, or the &#x60;name&#x60; used when creating the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiprecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | The status. Must have the value &#x60;stopped&#x60; | 

### Return type

[**ApiV2010AccountCallSiprec**](ApiV2010AccountCallSiprec.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStream

> ApiV2010AccountCallStream UpdateStream(ctx, accountSid, callSid, sid).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Stream resource.
    callSid := "callSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Stream resource is associated with.
    sid := "sid_example" // string | The SID of the Stream resource, or the `name` used when creating the resource
    status := "status_example" // string | The status. Must have the value `stopped`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateStream(context.Background(), accountSid, callSid, sid).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStream`: ApiV2010AccountCallStream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Stream resource. | 
**callSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Stream resource is associated with. | 
**sid** | **string** | The SID of the Stream resource, or the &#x60;name&#x60; used when creating the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | The status. Must have the value &#x60;stopped&#x60; | 

### Return type

[**ApiV2010AccountCallStream**](ApiV2010AccountCallStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageTrigger

> ApiV2010AccountUsageUsageTrigger UpdateUsageTrigger(ctx, accountSid, sid).CallbackMethod(callbackMethod).CallbackUrl(callbackUrl).FriendlyName(friendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountSid := "accountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update.
    sid := "sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.
    callbackMethod := "callbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    callbackUrl := "callbackUrl_example" // string | The URL we should call using `callback_method` when the trigger fires. (optional)
    friendlyName := "friendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateUsageTrigger(context.Background(), accountSid, sid).CallbackMethod(callbackMethod).CallbackUrl(callbackUrl).FriendlyName(friendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update. | 
**sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsageTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **callbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **callbackUrl** | **string** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires. | 
 **friendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

