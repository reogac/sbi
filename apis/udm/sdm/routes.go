/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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
	{
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodDelete,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnUnsubscribeForSharedData,
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: OnGetGroupIdentifiers,
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
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: OnGetIndividualSharedData,
	},
	{
		Label:   "ModifySharedDataSubs",
		Method:  http.MethodPatch,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnModifySharedDataSubs,
	},
	{
		Label:   "GetEcrData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/ecr-data",
		Handler: OnGetEcrData,
	},
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: OnGetLcsMoData,
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: OnGetLcsBcaData,
	},
	{
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: OnGetV2xData,
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/id/:ueId/sdm-subscriptions",
		Handler: OnSubscribe,
	},
	{
		Label:   "Modify",
		Method:  http.MethodPatch,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnModify,
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: OnGetDataSets,
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/lcs-privacy-data",
		Handler: OnGetLcsPrivacyData,
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: OnGetUcData,
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: OnGetAmData,
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smf-data",
		Handler: OnGetUeCtxInSmfData,
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodDelete,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnUnsubscribe,
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/sor-ack",
		Handler: OnSorAckInfo,
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: OnSubscribeToSharedData,
	},
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: OnGetUeCtxInAmfData,
	},
	{
		Label:   "UpuAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/upu-ack",
		Handler: OnUpuAck,
	},
	{
		Label:   "UpdateSORInfo",
		Method:  http.MethodPost,
		Path:    "/:supi/am-data/update-sor",
		Handler: OnUpdateSORInfo,
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: OnGetMultipleIdentifiers,
	},
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: OnGetUeCtxInSmsfData,
	},
	{
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: OnGetSmsData,
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: OnGetProseData,
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: OnGetMbsData,
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: OnSNSSAIsAck,
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/cag-ack",
		Handler: OnCAGAck,
	},
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: OnGetNSSAI,
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/id-translation-result",
		Handler: OnGetSupiOrGpsi,
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: OnGetSharedData,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
