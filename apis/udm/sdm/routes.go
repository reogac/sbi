/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:43 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: OnGetUeCtxInAmfData,
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: OnGetSmsMngtData,
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
		Label:   "Unsubscribe",
		Method:  http.MethodDelete,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnUnsubscribe,
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: OnSNSSAIsAck,
	},
	{
		Label:   "ModifySharedDataSubs",
		Method:  http.MethodPatch,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnModifySharedDataSubs,
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: OnGetSharedData,
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: OnSubscribeToSharedData,
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: OnGetMultipleIdentifiers,
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smf-data",
		Handler: OnGetUeCtxInSmfData,
	},
	{
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: OnGetTraceConfigData,
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/id-translation-result",
		Handler: OnGetSupiOrGpsi,
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/sor-ack",
		Handler: OnSorAckInfo,
	},
	{
		Label:   "Modify",
		Method:  http.MethodPatch,
		Path:    "/id/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: OnModify,
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
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: OnGetNSSAI,
	},
	{
		Label:   "GetEcrData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/ecr-data",
		Handler: OnGetEcrData,
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: OnGetUcData,
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
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodDelete,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: OnUnsubscribeForSharedData,
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: OnGetIndividualSharedData,
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: OnGetDataSets,
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: OnGetAmData,
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
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: OnGetSmsData,
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/id/:ueId/lcs-privacy-data",
		Handler: OnGetLcsPrivacyData,
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: OnGetMbsData,
	},
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: OnGetUeCtxInSmsfData,
	},
	{
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: OnGetV2xData,
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodPut,
		Path:    "/:supi/am-data/cag-ack",
		Handler: OnCAGAck,
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: OnGetGroupIdentifiers,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
