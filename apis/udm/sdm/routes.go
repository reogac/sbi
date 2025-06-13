/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: OnGetLcsMoData,
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: OnGetProseData,
	},
	{
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/:supi/smf-select-data",
		Handler: OnGetSmfSelData,
	},
	{
		Label:   "GetSmData",
		Method:  http.MethodGet,
		Path:    "/:supi/sm-data",
		Handler: OnGetSmData,
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: OnGetMbsData,
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: OnSubscribeToSharedData,
	},
	{
		Label:   "ModifySharedDataSubs",
		Method:  http.MethodPatch,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnModifySharedDataSubs,
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: OnGetGroupIdentifiers,
	},
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: OnGetNSSAI,
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: OnGetLcsBcaData,
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: OnGetUcData,
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/id-translation-result",
		Handler: OnGetSupiOrGpsi,
	},
	{
		Label:   "GetEcrData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/ecr-data",
		Handler: OnGetEcrData,
	},
	{
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: OnGetSmsData,
	},
	{
		Label:   "UpuAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/upu-ack",
		Handler: OnUpuAck,
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: OnSNSSAIsAck,
	},
	{
		Label:   "UpdateSORInfo",
		Method:  http.MethodPost,
		Path:    "/:supi/am-data/update-sor",
		Handler: OnUpdateSORInfo,
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: OnGetAmData,
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/lcs-privacy-data",
		Handler: OnGetLcsPrivacyData,
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/cag-ack",
		Handler: OnCAGAck,
	},
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: OnGetUeCtxInSmsfData,
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/id/:ueId/sdm-subscriptions",
		Handler: OnSubscribe,
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/sor-ack",
		Handler: OnSorAckInfo,
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smf-data",
		Handler: OnGetUeCtxInSmfData,
	},
	{
		Label:   "Modify",
		Method:  http.MethodPatch,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnModify,
	},
	{
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: OnGetV2xData,
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodDelete,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnUnsubscribe,
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: OnGetIndividualSharedData,
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: OnGetMultipleIdentifiers,
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: OnGetDataSets,
	},
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: OnGetUeCtxInAmfData,
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: OnGetSharedData,
	},
	{
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodDelete,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnUnsubscribeForSharedData,
	},
	{
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: OnGetTraceConfigData,
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: OnGetSmsMngtData,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
