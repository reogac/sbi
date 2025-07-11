/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "CreateAMFSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnCreateAMFSubscriptions,
	},
	{
		Label:   "Modify5GVnGroup",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Handler: OnModify5GVnGroup,
	},
	{
		Label:   "Query5mbsData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/5mbs-data",
		Handler: OnQuery5mbsData,
	},
	{
		Label:   "ModifySmfGroupSubscriptions",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnModifySmfGroupSubscriptions,
	},
	{
		Label:   "ModifyPpData",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/pp-data",
		Handler: OnModifyPpData,
	},
	{
		Label:   "ModifyAmfSubscriptionInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnModifyAmfSubscriptionInfo,
	},
	{
		Label:   "QueryPeiInformation",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/pei-info",
		Handler: OnQueryPeiInformation,
	},
	{
		Label:   "CreatePPDataEntry",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/pp-data-store/:afInstanceId",
		Handler: OnCreatePPDataEntry,
	},
	{
		Label:   "GetppData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/pp-data",
		Handler: OnGetppData,
	},
	{
		Label:   "QueryIndividualAuthenticationStatus",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		Handler: OnQueryIndividualAuthenticationStatus,
	},
	{
		Label:   "CreateOrUpdateSmfRegistration",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		Handler: OnCreateOrUpdateSmfRegistration,
	},
	{
		Label:   "CreateOperSpecData",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/operator-specific-data",
		Handler: OnCreateOperSpecData,
	},
	{
		Label:   "CreateSMFSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnCreateSMFSubscriptions,
	},
	{
		Label:   "RemoveHssSubscriptionsInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		Handler: OnRemoveHssSubscriptionsInfo,
	},
	{
		Label:   "ModifyAmfGroupSubscriptions",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnModifyAmfGroupSubscriptions,
	},
	{
		Label:   "Modify5GmbsGroup",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/group-data/mbs-group-membership/:externalGroupId",
		Handler: OnModify5GmbsGroup,
	},
	{
		Label:   "Query5GMbsGroupPPData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/mbs-group-membership/pp-profile-data",
		Handler: OnQuery5GMbsGroupPPData,
	},
	{
		Label:   "QueryNssaiAck",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais",
		Handler: OnQueryNssaiAck,
	},
	{
		Label:   "QuerySmsfContext3gpp",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/smsf-3gpp-access",
		Handler: OnQuerySmsfContext3gpp,
	},
	{
		Label:   "GetHssSDMSubscriptionInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		Handler: OnGetHssSDMSubscriptionInfo,
	},
	{
		Label:   "QueryAmfContext3gpp",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/amf-3gpp-access",
		Handler: OnQueryAmfContext3gpp,
	},
	{
		Label:   "ModifyEesubscription",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		Handler: OnModifyEesubscription,
	},
	{
		Label:   "QueryGroupEEData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-profile-data",
		Handler: OnQueryGroupEEData,
	},
	{
		Label:   "QueryPorseData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/prose-data",
		Handler: OnQueryPorseData,
	},
	{
		Label:   "GetMulticastMbsGroupMemb",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/mbs-group-membership/:externalGroupId",
		Handler: OnGetMulticastMbsGroupMemb,
	},
	{
		Label:   "QueryAuthUPU",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/upu-data",
		Handler: OnQueryAuthUPU,
	},
	{
		Label:   "CreateAmfGroupSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnCreateAmfGroupSubscriptions,
	},
	{
		Label:   "DeleteSmsfContext3gpp",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/smsf-3gpp-access",
		Handler: OnDeleteSmsfContext3gpp,
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/group-identifiers",
		Handler: OnGetGroupIdentifiers,
	},
	{
		Label:   "Create5GVnGroup",
		Method:  http.MethodPut,
		Path:    "/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Handler: OnCreate5GVnGroup,
	},
	{
		Label:   "RemoveAmfGroupSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnRemoveAmfGroupSubscriptions,
	},
	{
		Label:   "QueryPPData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/pp-profile-data",
		Handler: OnQueryPPData,
	},
	{
		Label:   "AmfContext3gpp",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/amf-3gpp-access",
		Handler: OnAmfContext3gpp,
	},
	{
		Label:   "CreateMessageWaitingData",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/mwd",
		Handler: OnCreateMessageWaitingData,
	},
	{
		Label:   "QuerySmsMngData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-mng-data",
		Handler: OnQuerySmsMngData,
	},
	{
		Label:   "Get5GVnGroupConfiguration",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Handler: OnGet5GVnGroupConfiguration,
	},
	{
		Label:   "ModifyAuthenticationSubscription",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-subscription",
		Handler: OnModifyAuthenticationSubscription,
	},
	{
		Label:   "QueryCagAck",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag",
		Handler: OnQueryCagAck,
	},
	{
		Label:   "QuerySmfRegList",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/smf-registrations",
		Handler: OnQuerySmfRegList,
	},
	{
		Label:   "CreateIpSmGwContext",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/ip-sm-gw",
		Handler: OnCreateIpSmGwContext,
	},
	{
		Label:   "RemoveMultipleSubscriptionDataSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/subs-to-notify",
		Handler: OnRemoveMultipleSubscriptionDataSubscriptions,
	},
	{
		Label:   "QueryLcsMoData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/lcs-mo-data",
		Handler: OnQueryLcsMoData,
	},
	{
		Label:   "QueryUeSubscribedData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId",
		Handler: OnQueryUeSubscribedData,
	},
	{
		Label:   "ModifyServiceSpecificAuthorizationInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		Handler: OnModifyServiceSpecificAuthorizationInfo,
	},
	{
		Label:   "CreateAuthenticationStatus",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status",
		Handler: OnCreateAuthenticationStatus,
	},
	{
		Label:   "QueryEeGroupSubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions",
		Handler: OnQueryEeGroupSubscriptions,
	},
	{
		Label:   "QueryEeGroupSubscription",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		Handler: OnQueryEeGroupSubscription,
	},
	{
		Label:   "ModifysubscriptionDataSubscription",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/subs-to-notify/:subsId",
		Handler: OnModifysubscriptionDataSubscription,
	},
	{
		Label:   "CreateServiceSpecificAuthorizationInfo",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		Handler: OnCreateServiceSpecificAuthorizationInfo,
	},
	{
		Label:   "UpdateRoamingInformation",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/roaming-information",
		Handler: OnUpdateRoamingInformation,
	},
	{
		Label:   "UpdateEesubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		Handler: OnUpdateEesubscriptions,
	},
	{
		Label:   "RemoveSmfSubscriptionsInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnRemoveSmfSubscriptionsInfo,
	},
	{
		Label:   "RemoveEeGroupSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		Handler: OnRemoveEeGroupSubscriptions,
	},
	{
		Label:   "CreateHSSSDMSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		Handler: OnCreateHSSSDMSubscriptions,
	},
	{
		Label:   "ModifyHssSDMSubscriptionInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		Handler: OnModifyHssSDMSubscriptionInfo,
	},
	{
		Label:   "QueryV2xData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/v2x-data",
		Handler: OnQueryV2xData,
	},
	{
		Label:   "QueryLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/lcs-bca-data",
		Handler: OnQueryLcsBcaData,
	},
	{
		Label:   "Query5GMbsGroupInternal",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/mbs-group-membership/internal",
		Handler: OnQuery5GMbsGroupInternal,
	},
	{
		Label:   "QueryMessageWaitingData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/mwd",
		Handler: OnQueryMessageWaitingData,
	},
	{
		Label:   "CreateEeGroupSubscriptions",
		Method:  http.MethodPost,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions",
		Handler: OnCreateEeGroupSubscriptions,
	},
	{
		Label:   "QueryCoverageRestrictionData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/coverage-restriction-data",
		Handler: OnQueryCoverageRestrictionData,
	},
	{
		Label:   "ModifyMessageWaitingData",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/mwd",
		Handler: OnModifyMessageWaitingData,
	},
	{
		Label:   "CreateAmfContext3gpp",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/amf-3gpp-access",
		Handler: OnCreateAmfContext3gpp,
	},
	{
		Label:   "ModifyIpSmGwContext",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/ip-sm-gw",
		Handler: OnModifyIpSmGwContext,
	},
	{
		Label:   "Queryeesubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions",
		Handler: OnQueryeesubscriptions,
	},
	{
		Label:   "QuerysdmSubscription",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		Handler: OnQuerysdmSubscription,
	},
	{
		Label:   "GetNiddAuData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/nidd-authorization-data",
		Handler: OnGetNiddAuData,
	},
	{
		Label:   "CreateOrUpdatePeiInformation",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/pei-info",
		Handler: OnCreateOrUpdatePeiInformation,
	},
	{
		Label:   "QuerySmData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/sm-data",
		Handler: OnQuerySmData,
	},
	{
		Label:   "UpdateEeGroupSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		Handler: OnUpdateEeGroupSubscriptions,
	},
	{
		Label:   "RemoveNiddAuthorizationInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/nidd-authorizations",
		Handler: OnRemoveNiddAuthorizationInfo,
	},
	{
		Label:   "CreateSmfGroupSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnCreateSmfGroupSubscriptions,
	},
	{
		Label:   "QuerySmfSelectData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/smf-selection-subscription-data",
		Handler: OnQuerySmfSelectData,
	},
	{
		Label:   "QueryRoamingInformation",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/roaming-information",
		Handler: OnQueryRoamingInformation,
	},
	{
		Label:   "RemoveSmfGroupSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnRemoveSmfGroupSubscriptions,
	},
	{
		Label:   "Query5GmbsGroup",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/mbs-group-membership",
		Handler: OnQuery5GmbsGroup,
	},
	{
		Label:   "GetServiceSpecificAuthorizationInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		Handler: OnGetServiceSpecificAuthorizationInfo,
	},
	{
		Label:   "CreateEeSubscriptions",
		Method:  http.MethodPost,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions",
		Handler: OnCreateEeSubscriptions,
	},
	{
		Label:   "QueryeeSubscription",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		Handler: OnQueryeeSubscription,
	},
	{
		Label:   "Querysdmsubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions",
		Handler: OnQuerysdmsubscriptions,
	},
	{
		Label:   "Updatesdmsubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		Handler: OnUpdatesdmsubscriptions,
	},
	{
		Label:   "RemovesubscriptionDataSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/subs-to-notify/:subsId",
		Handler: OnRemovesubscriptionDataSubscriptions,
	},
	{
		Label:   "Query5GVnGroupInternal",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/5g-vn-groups/internal",
		Handler: OnQuery5GVnGroupInternal,
	},
	{
		Label:   "GetPPDataEntry",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/pp-data-store/:afInstanceId",
		Handler: OnGetPPDataEntry,
	},
	{
		Label:   "QueryAmfContextNon3gpp",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		Handler: OnQueryAmfContextNon3gpp,
	},
	{
		Label:   "QueryOperSpecData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/operator-specific-data",
		Handler: OnQueryOperSpecData,
	},
	{
		Label:   "CreateSmsfContext3gpp",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/smsf-3gpp-access",
		Handler: OnCreateSmsfContext3gpp,
	},
	{
		Label:   "RemoveeeSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		Handler: OnRemoveeeSubscriptions,
	},
	{
		Label:   "GetHssSubscriptionInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		Handler: OnGetHssSubscriptionInfo,
	},
	{
		Label:   "RemovesdmSubscriptions",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		Handler: OnRemovesdmSubscriptions,
	},
	{
		Label:   "QuerySubscriptionDataSubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/subs-to-notify/:subsId",
		Handler: OnQuerySubscriptionDataSubscriptions,
	},
	{
		Label:   "CreateOrUpdateNssaiAck",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais",
		Handler: OnCreateOrUpdateNssaiAck,
	},
	{
		Label:   "Create5GmbsGroup",
		Method:  http.MethodPut,
		Path:    "/subscription-data/group-data/mbs-group-membership/:externalGroupId",
		Handler: OnCreate5GmbsGroup,
	},
	{
		Label:   "QueryIpSmGwContext",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ip-sm-gw",
		Handler: OnQueryIpSmGwContext,
	},
	{
		Label:   "ModifysdmSubscription",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		Handler: OnModifysdmSubscription,
	},
	{
		Label:   "QueryUserConsentData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/uc-data",
		Handler: OnQueryUserConsentData,
	},
	{
		Label:   "CreateCagUpdateAck",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag",
		Handler: OnCreateCagUpdateAck,
	},
	{
		Label:   "UpdateAuthenticationSoR",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		Handler: OnUpdateAuthenticationSoR,
	},
	{
		Label:   "UpdateSmfContext",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		Handler: OnUpdateSmfContext,
	},
	{
		Label:   "QuerySmsData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-data",
		Handler: OnQuerySmsData,
	},
	{
		Label:   "CreateNIDDAuthorizationInfo",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/nidd-authorizations",
		Handler: OnCreateNIDDAuthorizationInfo,
	},
	{
		Label:   "CreateAuthenticationSoR",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		Handler: OnCreateAuthenticationSoR,
	},
	{
		Label:   "CreateIndividualAuthenticationStatus",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		Handler: OnCreateIndividualAuthenticationStatus,
	},
	{
		Label:   "QueryAuthSoR",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		Handler: OnQueryAuthSoR,
	},
	{
		Label:   "QueryProvisionedData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data",
		Handler: OnQueryProvisionedData,
	},
	{
		Label:   "CreateAmfContextNon3gpp",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		Handler: OnCreateAmfContextNon3gpp,
	},
	{
		Label:   "ModifySmfSubscriptionInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnModifySmfSubscriptionInfo,
	},
	{
		Label:   "QueryContextData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data",
		Handler: OnQueryContextData,
	},
	{
		Label:   "Query5GVNGroupPPData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/5g-vn-groups/pp-profile-data",
		Handler: OnQuery5GVNGroupPPData,
	},
	{
		Label:   "DeleteAuthenticationStatus",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status",
		Handler: OnDeleteAuthenticationStatus,
	},
	{
		Label:   "GetNiddAuthorizationInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/nidd-authorizations",
		Handler: OnGetNiddAuthorizationInfo,
	},
	{
		Label:   "CreateSmsfContextNon3gpp",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		Handler: OnCreateSmsfContextNon3gpp,
	},
	{
		Label:   "RemoveServiceSpecificAuthorizationInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		Handler: OnRemoveServiceSpecificAuthorizationInfo,
	},
	{
		Label:   "CreateHSSSubscriptions",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		Handler: OnCreateHSSSubscriptions,
	},
	{
		Label:   "SubscriptionDataSubscriptions",
		Method:  http.MethodPost,
		Path:    "/subscription-data/subs-to-notify",
		Handler: OnSubscriptionDataSubscriptions,
	},
	{
		Label:   "QueryTraceData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/trace-data",
		Handler: OnQueryTraceData,
	},
	{
		Label:   "GetAmfGroupSubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnGetAmfGroupSubscriptions,
	},
	{
		Label:   "DeleteSmfRegistration",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		Handler: OnDeleteSmfRegistration,
	},
	{
		Label:   "CreateSdmSubscriptions",
		Method:  http.MethodPost,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions",
		Handler: OnCreateSdmSubscriptions,
	},
	{
		Label:   "GetIdentityData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/identity-data",
		Handler: OnGetIdentityData,
	},
	{
		Label:   "QueryAuthSubsData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-subscription",
		Handler: OnQueryAuthSubsData,
	},
	{
		Label:   "QueryAmData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/:servingPlmnId/provisioned-data/am-data",
		Handler: OnQueryAmData,
	},
	{
		Label:   "ModifyOperSpecData",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/operator-specific-data",
		Handler: OnModifyOperSpecData,
	},
	{
		Label:   "DeleteIpSmGwContext",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/ip-sm-gw",
		Handler: OnDeleteIpSmGwContext,
	},
	{
		Label:   "DeletePPDataEntry",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/pp-data-store/:afInstanceId",
		Handler: OnDeletePPDataEntry,
	},
	{
		Label:   "RemoveAmfSubscriptionsInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnRemoveAmfSubscriptionsInfo,
	},
	{
		Label:   "QueryAuthenticationStatus",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status",
		Handler: OnQueryAuthenticationStatus,
	},
	{
		Label:   "ModifyEeGroupSubscription",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		Handler: OnModifyEeGroupSubscription,
	},
	{
		Label:   "QuerySubsToNotify",
		Method:  http.MethodGet,
		Path:    "/subscription-data/subs-to-notify",
		Handler: OnQuerySubsToNotify,
	},
	{
		Label:   "QueryUeLocation",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/location",
		Handler: OnQueryUeLocation,
	},
	{
		Label:   "Delete5GmbsGroup",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/group-data/mbs-group-membership/:externalGroupId",
		Handler: OnDelete5GmbsGroup,
	},
	{
		Label:   "DeleteIndividualAuthenticationStatus",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		Handler: OnDeleteIndividualAuthenticationStatus,
	},
	{
		Label:   "QuerySmfRegistration",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		Handler: OnQuerySmfRegistration,
	},
	{
		Label:   "DeleteOperSpecData",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/operator-specific-data",
		Handler: OnDeleteOperSpecData,
	},
	{
		Label:   "QuerySmsfContextNon3gpp",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		Handler: OnQuerySmsfContextNon3gpp,
	},
	{
		Label:   "DeleteSmsfContextNon3gpp",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		Handler: OnDeleteSmsfContextNon3gpp,
	},
	{
		Label:   "GetSmfSubscriptionInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnGetSmfSubscriptionInfo,
	},
	{
		Label:   "QueryEEData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/ee-profile-data",
		Handler: OnQueryEEData,
	},
	{
		Label:   "QueryLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/lcs-privacy-data",
		Handler: OnQueryLcsPrivacyData,
	},
	{
		Label:   "CreateAuthenticationUPU",
		Method:  http.MethodPut,
		Path:    "/subscription-data/:ueId/ue-update-confirmation-data/upu-data",
		Handler: OnCreateAuthenticationUPU,
	},
	{
		Label:   "RemoveHssSDMSubscriptionsInfo",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		Handler: OnRemoveHssSDMSubscriptionsInfo,
	},
	{
		Label:   "GetOdbData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/operator-determined-barring-data",
		Handler: OnGetOdbData,
	},
	{
		Label:   "GetSSAuData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/service-specific-authorization-data/:serviceType",
		Handler: OnGetSSAuData,
	},
	{
		Label:   "GetSmfGroupSubscriptions",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		Handler: OnGetSmfGroupSubscriptions,
	},
	{
		Label:   "GetMultiplePPDataEntries",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/pp-data-store",
		Handler: OnGetMultiplePPDataEntries,
	},
	{
		Label:   "DeleteMessageWaitingData",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/:ueId/context-data/mwd",
		Handler: OnDeleteMessageWaitingData,
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/shared-data/:sharedDataId",
		Handler: OnGetIndividualSharedData,
	},
	{
		Label:   "ModifyNiddAuthorizationInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/nidd-authorizations",
		Handler: OnModifyNiddAuthorizationInfo,
	},
	{
		Label:   "AmfContextNon3gpp",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		Handler: OnAmfContextNon3gpp,
	},
	{
		Label:   "ModifyHssSubscriptionInfo",
		Method:  http.MethodPatch,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		Handler: OnModifyHssSubscriptionInfo,
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/subscription-data/shared-data",
		Handler: OnGetSharedData,
	},
	{
		Label:   "Query5GVnGroup",
		Method:  http.MethodGet,
		Path:    "/subscription-data/group-data/5g-vn-groups",
		Handler: OnQuery5GVnGroup,
	},
	{
		Label:   "Delete5GVnGroup",
		Method:  http.MethodDelete,
		Path:    "/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Handler: OnDelete5GVnGroup,
	},
	{
		Label:   "GetAmfSubscriptionInfo",
		Method:  http.MethodGet,
		Path:    "/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		Handler: OnGetAmfSubscriptionInfo,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
