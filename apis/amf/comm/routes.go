/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package comm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "EBIAssignment",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/assign-ebi",
		Handler: OnEBIAssignment,
	},
	{
		Label:   "RegistrationStatusUpdate",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer-update",
		Handler: OnRegistrationStatusUpdate,
	},
	{
		Label:   "N1N2MessageSubscribe",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
		Handler: OnN1N2MessageSubscribe,
	},
	{
		Label:   "NonUeN2InfoSubscribe",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/subscriptions",
		Handler: OnNonUeN2InfoSubscribe,
	},
	{
		Label:   "AMFStatusChangeSubscribeModfy",
		Method:  http.MethodPut,
		Path:    "/subscriptions/:subscriptionId",
		Handler: OnAMFStatusChangeSubscribeModfy,
	},
	{
		Label:   "AMFStatusChangeUnSubscribe",
		Method:  http.MethodDelete,
		Path:    "/subscriptions/:subscriptionId",
		Handler: OnAMFStatusChangeUnSubscribe,
	},
	{
		Label:   "RelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/relocate",
		Handler: OnRelocateUEContext,
	},
	{
		Label:   "N1N2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages",
		Handler: OnN1N2MessageTransfer,
	},
	{
		Label:   "NonUeN2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/transfer",
		Handler: OnNonUeN2MessageTransfer,
	},
	{
		Label:   "ReleaseUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/release",
		Handler: OnReleaseUEContext,
	},
	{
		Label:   "UEContextTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer",
		Handler: OnUEContextTransfer,
	},
	{
		Label:   "N1N2MessageUnSubscribe",
		Method:  http.MethodDelete,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
		Handler: OnN1N2MessageUnSubscribe,
	},
	{
		Label:   "CreateUEContext",
		Method:  http.MethodPut,
		Path:    "/ue-contexts/:ueContextId",
		Handler: OnCreateUEContext,
	},
	{
		Label:   "CancelRelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/cancel-relocate",
		Handler: OnCancelRelocateUEContext,
	},
	{
		Label:   "NonUeN2InfoUnSubscribe",
		Method:  http.MethodDelete,
		Path:    "/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
		Handler: OnNonUeN2InfoUnSubscribe,
	},
	{
		Label:   "AMFStatusChangeSubscribe",
		Method:  http.MethodPost,
		Path:    "/subscriptions",
		Handler: OnAMFStatusChangeSubscribe,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
