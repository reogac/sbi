/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package policy

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "UpdateSessionManagementPolicyData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnUpdateSessionManagementPolicyData,
	},
	{
		Label:   "CreateUsageMonitoringResource",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnCreateUsageMonitoringResource,
	},
	{
		Label:   "CreateIndividualPolicyDataSubscription",
		Method:  http.MethodPost,
		Path:    "/policy-data/subs-to-notify",
		Handler: OnCreateIndividualPolicyDataSubscription,
	},
	{
		Label:   "UpdateSlicePolicyControlData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnUpdateSlicePolicyControlData,
	},
	{
		Label:   "ReadAccessAndMobilityPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/am-data",
		Handler: OnReadAccessAndMobilityPolicyData,
	},
	{
		Label:   "CreateOrReplaceUEPolicySet",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnCreateOrReplaceUEPolicySet,
	},
	{
		Label:   "ReadSessionManagementPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data",
		Handler: OnReadSessionManagementPolicyData,
	},
	{
		Label:   "DeleteUsageMonitoringInformation",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnDeleteUsageMonitoringInformation,
	},
	{
		Label:   "ReplaceIndividualPolicyDataSubscription",
		Method:  http.MethodPut,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnReplaceIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReadOperatorSpecificData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReadOperatorSpecificData,
	},
	{
		Label:   "ReadSlicePolicyControlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/slice-control-data/:snssai",
		Handler: OnReadSlicePolicyControlData,
	},
	{
		Label:   "ReadSponsorConnectivityData",
		Method:  http.MethodGet,
		Path:    "/policy-data/sponsor-connectivity-data/:sponsorId",
		Handler: OnReadSponsorConnectivityData,
	},
	{
		Label:   "DeleteOperatorSpecificData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnDeleteOperatorSpecificData,
	},
	{
		Label:   "GetMBSSessPolCtrlData",
		Method:  http.MethodGet,
		Path:    "/policy-data/mbs-session-pol-data/:polSessionId",
		Handler: OnGetMBSSessPolCtrlData,
	},
	{
		Label:   "ReadUEPolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnReadUEPolicySet,
	},
	{
		Label:   "UpdateUEPolicySet",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/ue-policy-set",
		Handler: OnUpdateUEPolicySet,
	},
	{
		Label:   "DeleteIndividualBdtData",
		Method:  http.MethodDelete,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnDeleteIndividualBdtData,
	},
	{
		Label:   "UpdateIndividualBdtData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnUpdateIndividualBdtData,
	},
	{
		Label:   "ReplaceOperatorSpecificData",
		Method:  http.MethodPut,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnReplaceOperatorSpecificData,
	},
	{
		Label:   "UpdateOperatorSpecificData",
		Method:  http.MethodPatch,
		Path:    "/policy-data/ues/:ueId/operator-specific-data",
		Handler: OnUpdateOperatorSpecificData,
	},
	{
		Label:   "DeleteIndividualPolicyDataSubscription",
		Method:  http.MethodDelete,
		Path:    "/policy-data/subs-to-notify/:subsId",
		Handler: OnDeleteIndividualPolicyDataSubscription,
	},
	{
		Label:   "ReadPlmnUePolicySet",
		Method:  http.MethodGet,
		Path:    "/policy-data/plmns/:plmnId/ue-policy-set",
		Handler: OnReadPlmnUePolicySet,
	},
	{
		Label:   "ReadUsageMonitoringInformation",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId/sm-data/:usageMonId",
		Handler: OnReadUsageMonitoringInformation,
	},
	{
		Label:   "ReadIndividualBdtData",
		Method:  http.MethodGet,
		Path:    "/policy-data/bdt-data/:bdtReferenceId",
		Handler: OnReadIndividualBdtData,
	},
	{
		Label:   "ReadPolicyData",
		Method:  http.MethodGet,
		Path:    "/policy-data/ues/:ueId",
		Handler: OnReadPolicyData,
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
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
