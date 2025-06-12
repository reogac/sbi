/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uecm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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
		Label:   "Update3GppRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: OnUpdate3GppRegistration,
	},
	{
		Label:   "ThreeGppSmsfRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: OnThreeGppSmsfRegistration,
	},
	{
		Label:   "PeiUpdate",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/pei-update",
		Handler: OnPeiUpdate,
	},
	{
		Label:   "GetSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations",
		Handler: OnGetSmfRegistration,
	},
	{
		Label:   "UpdateSmfRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnUpdateSmfRegistration,
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
		Label:   "GetIpSmGwRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: OnGetIpSmGwRegistration,
	},
	{
		Label:   "Get3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: OnGet3GppSmsfRegistration,
	},
	{
		Label:   "GetNon3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnGetNon3GppSmsfRegistration,
	},
	{
		Label:   "IpSmGwRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: OnIpSmGwRegistration,
	},
	{
		Label:   "GetLocationInfo",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/location",
		Handler: OnGetLocationInfo,
	},
	{
		Label:   "GetNwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations",
		Handler: OnGetNwdafRegistration,
	},
	{
		Label:   "UpdateNwdafRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: OnUpdateNwdafRegistration,
	},
	{
		Label:   "Non3GppRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: OnNon3GppRegistration,
	},
	{
		Label:   "UpdateNon3GppRegistration",
		Method:  http.MethodPatch,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: OnUpdateNon3GppRegistration,
	},
	{
		Label:   "NwdafDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: OnNwdafDeregistration,
	},
	{
		Label:   "ThreeGppRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: OnThreeGppRegistration,
	},
	{
		Label:   "DeregAMF",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/dereg-amf",
		Handler: OnDeregAMF,
	},
	{
		Label:   "RetrieveSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnRetrieveSmfRegistration,
	},
	{
		Label:   "SmfDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnSmfDeregistration,
	},
	{
		Label:   "Non3GppSmsfRegistration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnNon3GppSmsfRegistration,
	},
	{
		Label:   "Non3GppSmsfDeregistration",
		Method:  http.MethodDelete,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: OnNon3GppSmsfDeregistration,
	},
	{
		Label:   "GetRegistrations",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations",
		Handler: OnGetRegistrations,
	},
	{
		Label:   "Registration",
		Method:  http.MethodPut,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: OnRegistration,
	},
	{
		Label:   "TriggerPCSCFRestoration",
		Method:  http.MethodPost,
		Path:    "/restore-pcscf",
		Handler: OnTriggerPCSCFRestoration,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
