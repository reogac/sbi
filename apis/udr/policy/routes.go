/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package policy

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "ReplaceIndividualPolicyDataSubscription",
		Method:  http.MethodPut,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnReplaceIndividualPolicyDataSubscription,
	},
	{
		Label:   "DeleteOperatorSpecificData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnDeleteOperatorSpecificData,
	},
	{
		Label:   "CreateUsageMonitoringResource",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnCreateUsageMonitoringResource,
	},
	{
		Label:   "CreateIndividualBdtData",
		Method:  http.MethodPut,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnCreateIndividualBdtData,
	},
	{
		Label:   "DeleteIndividualPolicyDataSubscription",
		Method:  http.MethodDelete,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnDeleteIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReadOperatorSpecificData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReadOperatorSpecificData,
	},
	{
		Label:   "CreateIndividualPolicyDataSubscription",
		Method:  http.MethodPost,
		Path:    "/policy-data/subs-to-notify",
		Handler: OnCreateIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReplaceOperatorSpecificData",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReplaceOperatorSpecificData,
	},
	{
		Label:   "ReadUEPolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnReadUEPolicySet,
	},
	{
		Label:   "DeleteUsageMonitoringInformation",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnDeleteUsageMonitoringInformation,
	},
	{
		Label:   "ReadBdtData",
		Method:  http.MethodGet,
		Path:    "/policy-data/bdt-data",
		Handler: OnReadBdtData,
	},
	{
		Label:   "ReadIndividualBdtData",
		Method:  http.MethodGet,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnReadIndividualBdtData,
	},
	{
		Label:   "DeleteIndividualBdtData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnDeleteIndividualBdtData,
	},
	{
		Label:   "ReadPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId",
		Handler: OnReadPolicyData,
	},
	{
		Label:   "ReadAccessAndMobilityPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/am-data",
		Handler: OnReadAccessAndMobilityPolicyData,
	},
	{
		Label:   "ReadPlmnUePolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/plmns/:plmnId/ue-policy-set",
		Handler: OnReadPlmnUePolicySet,
	},
	{
		Label:   "UpdateSlicePolicyControlData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnUpdateSlicePolicyControlData,
	},
	{
		Label:   "CreateOrReplaceUEPolicySet",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnCreateOrReplaceUEPolicySet,
	},
	{
		Label:   "UpdateUEPolicySet",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnUpdateUEPolicySet,
	},
	{
		Label:   "ReadUsageMonitoringInformation",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnReadUsageMonitoringInformation,
	},
	{
		Label:   "ReadSponsorConnectivityData",
		Method:  http.MethodGet,
		Path:    "/policy-data/sponsor-connectivity-data/:sponsorId",
		Handler: OnReadSponsorConnectivityData,
	},
	{
		Label:   "ReadSessionManagementPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnReadSessionManagementPolicyData,
	},
	{
		Label:   "UpdateSessionManagementPolicyData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnUpdateSessionManagementPolicyData,
	},
	{
		Label:   "UpdateOperatorSpecificData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnUpdateOperatorSpecificData,
	},
	{
		Label:   "ReadSlicePolicyControlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnReadSlicePolicyControlData,
	},
	{
		Label:   "GetMBSSessPolCtrlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/mbs-session-pol-data/:polSessionId",
		Handler: OnGetMBSSessPolCtrlData,
	},
	{
		Label:   "UpdateIndividualBdtData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnUpdateIndividualBdtData,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
