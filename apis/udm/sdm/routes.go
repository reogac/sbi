/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:53 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: OnGetNSSAI,
	},
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: OnGetUeCtxInAmfData,
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
		Label:   "GetSmData",
		Method:  http.MethodGet,
		Path:    "/:supi/sm-data",
		Handler: OnGetSmData,
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
		Label:   "Modify",
		Method:  http.MethodPatch,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnModify,
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/sor-ack",
		Handler: OnSorAckInfo,
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: OnGetMultipleIdentifiers,
	},
	{
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/:supi/smf-select-data",
		Handler: OnGetSmfSelData,
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smf-data",
		Handler: OnGetUeCtxInSmfData,
	},
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: OnGetLcsMoData,
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodDelete,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnUnsubscribe,
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/lcs-privacy-data",
		Handler: OnGetLcsPrivacyData,
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: OnGetProseData,
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/id/:ueId/sdm-subscriptions",
		Handler: OnSubscribe,
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: OnSubscribeToSharedData,
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: OnGetAmData,
	},
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: OnGetUeCtxInSmsfData,
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: OnGetUcData,
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/cag-ack",
		Handler: OnCAGAck,
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: OnGetSharedData,
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: OnGetIndividualSharedData,
	},
	{
		Label:   "GetEcrData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/ecr-data",
		Handler: OnGetEcrData,
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: OnGetLcsBcaData,
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: OnGetMbsData,
	},
	{
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: OnGetV2xData,
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/id-translation-result",
		Handler: OnGetSupiOrGpsi,
	},
	{
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodDelete,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnUnsubscribeForSharedData,
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: OnGetSmsMngtData,
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: OnSNSSAIsAck,
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: OnGetDataSets,
	},
	{
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: OnGetTraceConfigData,
	},
	{
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: OnGetSmsData,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
