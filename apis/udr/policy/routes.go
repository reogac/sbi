/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package policy

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "ReadIndividualBdtData",
		Method:  http.MethodGet,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnReadIndividualBdtData,
	},
	{
		Label:   "DeleteIndividualPolicyDataSubscription",
		Method:  http.MethodDelete,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnDeleteIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReplaceOperatorSpecificData",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReplaceOperatorSpecificData,
	},
	{
		Label:   "ReadUsageMonitoringInformation",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnReadUsageMonitoringInformation,
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
		Label:   "CreateUsageMonitoringResource",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnCreateUsageMonitoringResource,
	},
	{
		Label:   "DeleteIndividualBdtData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnDeleteIndividualBdtData,
	},
	{
		Label:   "DeleteOperatorSpecificData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnDeleteOperatorSpecificData,
	},
	{
		Label:   "UpdateSlicePolicyControlData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnUpdateSlicePolicyControlData,
	},
	{
		Label:   "GetMBSSessPolCtrlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/mbs-session-pol-data/:polSessionId",
		Handler: OnGetMBSSessPolCtrlData,
	},
	{
		Label:   "ReadPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId",
		Handler: OnReadPolicyData,
	},
	{
		Label:   "ReadSessionManagementPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnReadSessionManagementPolicyData,
	},
	{
		Label:   "UpdateIndividualBdtData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnUpdateIndividualBdtData,
	},
	{
		Label:   "ReadPlmnUePolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/plmns/:plmnId/ue-policy-set",
		Handler: OnReadPlmnUePolicySet,
	},
	{
		Label:   "ReadSlicePolicyControlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnReadSlicePolicyControlData,
	},
	{
		Label:   "ReadAccessAndMobilityPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/am-data",
		Handler: OnReadAccessAndMobilityPolicyData,
	},
	{
		Label:   "ReplaceIndividualPolicyDataSubscription",
		Method:  http.MethodPut,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnReplaceIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReadUEPolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnReadUEPolicySet,
	},
	{
		Label:   "ReadBdtData",
		Method:  http.MethodGet,
		Path:    "/policy-data/bdt-data",
		Handler: OnReadBdtData,
	},
	{
		Label:   "CreateIndividualBdtData",
		Method:  http.MethodPut,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnCreateIndividualBdtData,
	},
	{
		Label:   "UpdateSessionManagementPolicyData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnUpdateSessionManagementPolicyData,
	},
	{
		Label:   "ReadOperatorSpecificData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReadOperatorSpecificData,
	},
	{
		Label:   "ReadSponsorConnectivityData",
		Method:  http.MethodGet,
		Path:    "/policy-data/sponsor-connectivity-data/:sponsorId",
		Handler: OnReadSponsorConnectivityData,
	},
	{
		Label:   "DeleteUsageMonitoringInformation",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnDeleteUsageMonitoringInformation,
	},
	{
		Label:   "CreateIndividualPolicyDataSubscription",
		Method:  http.MethodPost,
		Path:    "/policy-data/subs-to-notify",
		Handler: OnCreateIndividualPolicyDataSubscription,
	},
	{
		Label:   "UpdateOperatorSpecificData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnUpdateOperatorSpecificData,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
