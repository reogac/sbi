/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uecm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "GetRegistrations",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations",
		Handler: OnGetRegistrations,
	},
	{
		Label:   "ThreeGppRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: OnThreeGppRegistration,
	},
	{
		Label:   "SmfDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnSmfDeregistration,
	},
	{
		Label:   "Get3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: OnGet3GppSmsfRegistration,
	},
	{
		Label:   "IpSmGwRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: OnIpSmGwRegistration,
	},
	{
		Label:   "Update3GppRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: OnUpdate3GppRegistration,
	},
	{
		Label:   "Registration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnRegistration,
	},
	{
		Label:   "ThreeGppSmsfRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: OnThreeGppSmsfRegistration,
	},
	{
		Label:   "GetIpSmGwRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: OnGetIpSmGwRegistration,
	},
	{
		Label:   "SendRoutingInfoSm",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/send-routing-info-sm",
		Handler: OnSendRoutingInfoSm,
	},
	{
		Label:   "Get3GppRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: OnGet3GppRegistration,
	},
	{
		Label:   "UpdateRoamingInformation",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/roaming-info-update",
		Handler: OnUpdateRoamingInformation,
	},
	{
		Label:   "ThreeGppSmsfDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: OnThreeGppSmsfDeregistration,
	},
	{
		Label:   "Non3GppRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: OnNon3GppRegistration,
	},
	{
		Label:   "GetNon3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnGetNon3GppSmsfRegistration,
	},
	{
		Label:   "Non3GppSmsfDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnNon3GppSmsfDeregistration,
	},
	{
		Label:   "IpSmGwDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: OnIpSmGwDeregistration,
	},
	{
		Label:   "NwdafRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: OnNwdafRegistration,
	},
	{
		Label:   "GetNon3GppRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: OnGetNon3GppRegistration,
	},
	{
		Label:   "TriggerPCSCFRestoration",
		Method:  http.MethodPost,
		Path:    "/restore-pcscf",
		Handler: OnTriggerPCSCFRestoration,
	},
	{
		Label:   "GetNwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations",
		Handler: OnGetNwdafRegistration,
	},
	{
		Label:   "NwdafDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: OnNwdafDeregistration,
	},
	{
		Label:   "PeiUpdate",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/pei-update",
		Handler: OnPeiUpdate,
	},
	{
		Label:   "UpdateNon3GppRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: OnUpdateNon3GppRegistration,
	},
	{
		Label:   "Non3GppSmsfRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnNon3GppSmsfRegistration,
	},
	{
		Label:   "GetSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations",
		Handler: OnGetSmfRegistration,
	},
	{
		Label:   "RetrieveSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnRetrieveSmfRegistration,
	},
	{
		Label:   "UpdateSmfRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnUpdateSmfRegistration,
	},
	{
		Label:   "GetLocationInfo",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/location",
		Handler: OnGetLocationInfo,
	},
	{
		Label:   "UpdateNwdafRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: OnUpdateNwdafRegistration,
	},
	{
		Label:   "DeregAMF",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/dereg-amf",
		Handler: OnDeregAMF,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
