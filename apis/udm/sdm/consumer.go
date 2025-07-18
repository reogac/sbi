/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nudm-sdm/v2"
)

// Summary: retrieve a UE's ProSe Subscription Data
// Description:
// Path: /:supi/prose-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetProseDataParams struct {
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
}

func GetProseData(cli sbi.ConsumerClient, params GetProseDataParams) (headers map[string]string, rsp *models.ProseSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/prose-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.ProseSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ProseSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's User Consent Subscription Data
// Description:
// Path: /:supi/uc-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetUcDataParams struct {
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	UcPurpose         string
	IfNoneMatch       string
}

func GetUcData(cli sbi.ConsumerClient, params GetUcDataParams) (headers map[string]string, rsp *models.UcSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/uc-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.UcPurpose) > 0 {
		request.AddParam("uc-purpose", params.UcPurpose)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.UcSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UcSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the subscription
// Description:
// Path: /id/:ueId/sdm-subscriptions/:subscriptionId
// Path Params: ueId, subscriptionId
type ModifyParams struct {
	UeId              string
	SubscriptionId    string
	SupportedFeatures string
}

func Modify(cli sbi.ConsumerClient, params ModifyParams, body *models.SdmSubsModification) (rsp *models.Schema, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/id/%s/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubscriptionId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.Schema)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Schema: %+v", err)
		}
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Nudm_Sdm Info operation for S-NSSAIs acknowledgement
// Description:
// Path: /:supi/am-data/subscribed-snssais-ack
// Path Params: supi
func SNSSAIsAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/subscribed-snssais-ack", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's subscribed Enhanced Coverage Restriction Data
// Description:
// Path: /:supi/am-data/ecr-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetEcrDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetEcrData(cli sbi.ConsumerClient, params GetEcrDataParams) (headers map[string]string, rsp *models.EnhancedCoverageRestrictionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/ecr-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.EnhancedCoverageRestrictionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EnhancedCoverageRestrictionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's UE Context In SMSF Data
// Description:
// Path: /:supi/ue-context-in-smsf-data
// Path Params: supi
type GetUeCtxInSmsfDataParams struct {
	SupportedFeatures string
	Supi              string
}

func GetUeCtxInSmsfData(cli sbi.ConsumerClient, params GetUeCtxInSmsfDataParams) (rsp *models.UeContextInSmsfData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/ue-context-in-smsf-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UeContextInSmsfData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextInSmsfData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's SUPI or GPSI
// Description:
// Path: /id/:ueId/id-translation-result
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type GetSupiOrGpsiParams struct {
	IfModifiedSince   string
	AppPortId         *models.AppPortId
	MtcProviderInfo   string
	SupportedFeatures string
	AfId              string
	AfServiceId       string
	RequestedGpsiType string
	IfNoneMatch       string
	UeId              string
}

func GetSupiOrGpsi(cli sbi.ConsumerClient, params GetSupiOrGpsiParams) (headers map[string]string, rsp *models.IdTranslationResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/id/%s/id-translation-result", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.AfServiceId) > 0 {
		request.AddParam("af-service-id", params.AfServiceId)
	}
	if len(params.RequestedGpsiType) > 0 {
		request.AddParam("requested-gpsi-type", params.RequestedGpsiType)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if params.AppPortId != nil {
		request.AddParam("app-port-id", models.AppPortIdToString(*params.AppPortId))
	}
	if len(params.MtcProviderInfo) > 0 {
		request.AddParam("mtc-provider-info", params.MtcProviderInfo)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.IdTranslationResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode IdTranslationResult: %+v", err)
		}
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Nudm_Sdm custom operation to trigger SOR info update
// Description:
// Path: /:supi/am-data/update-sor
// Path Params: supi
func UpdateSORInfo(cli sbi.ConsumerClient, supi string, body *models.SorUpdateInfo) (rsp *models.SorInfo, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/update-sor", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SorInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SorInfo: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: subscribe to notifications for shared data
// Description:
// Path: /shared-data-subscriptions
// Path Params:
// Response headers: Location
func SubscribeToSharedData(cli sbi.ConsumerClient, body *models.SdmSubscription) (headers map[string]string, rsp *models.SdmSubscription, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/shared-data-subscriptions", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.SdmSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SdmSubscription: %+v", err)
		}
	case 400, 404:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's 5MBS Subscription Data
// Description:
// Path: /:supi/5mbs-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetMbsDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetMbsData(cli sbi.ConsumerClient, params GetMbsDataParams) (headers map[string]string, rsp *models.MbsSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/5mbs-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.MbsSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MbsSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve shared data
// Description:
// Path: /shared-data
// Path Params:
// Response headers: Cache-Control, ETag, Last-Modified
type GetSharedDataParams struct {
	SharedDataIds     []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSharedData(cli sbi.ConsumerClient, params GetSharedDataParams) (headers map[string]string, rsp *[]models.SharedData, err error) {

	if len(params.SharedDataIds) == 0 {
		err = fmt.Errorf("shared-data-ids is required")
		return
	}

	path := fmt.Sprintf("%s/shared-data", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("shared-data-ids", models.ArrayOfStringToString(params.SharedDataIds))
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new([]models.SharedData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []SharedData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's LCS Privacy Subscription Data
// Description:
// Path: /id/:ueId/lcs-privacy-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type GetLcsPrivacyDataParams struct {
	IfModifiedSince   string
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
}

func GetLcsPrivacyData(cli sbi.ConsumerClient, params GetLcsPrivacyDataParams) (headers map[string]string, rsp *models.LcsPrivacyData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/id/%s/lcs-privacy-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.LcsPrivacyData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsPrivacyData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's LCS Broadcast Assistance Data Types Subscription Data
// Description:
// Path: /:supi/lcs-bca-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetLcsBcaDataParams struct {
	PlmnId            *models.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetLcsBcaData(cli sbi.ConsumerClient, params GetLcsBcaDataParams) (headers map[string]string, rsp *models.LcsBroadcastAssistanceTypesData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/lcs-bca-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.LcsBroadcastAssistanceTypesData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsBroadcastAssistanceTypesData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Mapping of Group Identifiers
// Description:
// Path: /group-data/group-identifiers
// Path Params:
// Response headers: Cache-Control, ETag, Last-Modified
type GetGroupIdentifiersParams struct {
	IfModifiedSince   string
	ExtGroupId        string
	IntGroupId        string
	UeIdInd           *bool
	SupportedFeatures string
	AfId              string
	IfNoneMatch       string
}

func GetGroupIdentifiers(cli sbi.ConsumerClient, params GetGroupIdentifiersParams) (headers map[string]string, rsp *models.GroupIdentifiers, err error) {

	path := fmt.Sprintf("%s/group-data/group-identifiers", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.ExtGroupId) > 0 {
		request.AddParam("ext-group-id", params.ExtGroupId)
	}
	if len(params.IntGroupId) > 0 {
		request.AddParam("int-group-id", params.IntGroupId)
	}
	if params.UeIdInd != nil {
		request.AddParam("ue-id-ind", models.BoolToString(*params.UeIdInd))
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.GroupIdentifiers)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode GroupIdentifiers: %+v", err)
		}
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's subscribed NSSAI
// Description:
// Path: /:supi/nssai
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetNSSAIParams struct {
	DisasterRoamingInd *bool
	IfNoneMatch        string
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *models.PlmnId
}

func GetNSSAI(cli sbi.ConsumerClient, params GetNSSAIParams) (headers map[string]string, rsp *models.Nssai, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/nssai", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.Nssai)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Nssai: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's LCS Mobile Originated Subscription Data
// Description:
// Path: /:supi/lcs-mo-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetLcsMoDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetLcsMoData(cli sbi.ConsumerClient, params GetLcsMoDataParams) (headers map[string]string, rsp *models.LcsMoData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/lcs-mo-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.LcsMoData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsMoData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: unsubscribe from notifications for shared data
// Description:
// Path: /shared-data-subscriptions/:subscriptionId
// Path Params: subscriptionId
func UnsubscribeForSharedData(cli sbi.ConsumerClient, subscriptionId string) (err error) {

	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	path := fmt.Sprintf("%s/shared-data-subscriptions/%s", PATH_ROOT, subscriptionId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Mapping of UE Identifiers
// Description:
// Path: /multiple-identifiers
// Path Params:
// Response headers: Cache-Control, ETag, Last-Modified
type GetMultipleIdentifiersParams struct {
	SupportedFeatures string
	GpsiList          []string
}

func GetMultipleIdentifiers(cli sbi.ConsumerClient, params GetMultipleIdentifiersParams) (headers map[string]string, rsp *map[string]models.SupiInfo, err error) {

	if len(params.GpsiList) == 0 {
		err = fmt.Errorf("gpsi-list is required")
		return
	}

	path := fmt.Sprintf("%s/multiple-identifiers", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.AddParam("gpsi-list", models.ArrayOfStringToString(params.GpsiList))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(map[string]models.SupiInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]SupiInfo: %+v", err)
		}
	case 400, 401, 403, 404, 429, 500, 502, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's Session Management Subscription Data
// Description:
// Path: /:supi/sm-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetSmDataParams struct {
	Dnn               string
	PlmnId            *models.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	SingleNssai       *models.Snssai
}

func GetSmData(cli sbi.ConsumerClient, params GetSmDataParams) (headers map[string]string, rsp *models.SmSubsData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/sm-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.SingleNssai != nil {
		request.AddParam("single-nssai", models.SnssaiToString(*params.SingleNssai))
	}
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SmSubsData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmSubsData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's V2X Subscription Data
// Description:
// Path: /:supi/v2x-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetV2xDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetV2xData(cli sbi.ConsumerClient, params GetV2xDataParams) (headers map[string]string, rsp *models.V2xSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/v2x-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.V2xSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode V2xSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Nudm_Sdm Info operation for CAG acknowledgement
// Description:
// Path: /:supi/am-data/cag-ack
// Path Params: supi
func CAGAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/cag-ack", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the individual shared data
// Description:
// Path: /shared-data/:sharedDataId
// Path Params: sharedDataId
// Response headers: Cache-Control, ETag, Last-Modified
type GetIndividualSharedDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	SharedDataId      []string
	SupportedFeatures string
}

func GetIndividualSharedData(cli sbi.ConsumerClient, params GetIndividualSharedDataParams) (headers map[string]string, rsp *models.SharedData, err error) {

	if len(params.SharedDataId) == 0 {
		err = fmt.Errorf("sharedDataId is required")
		return
	}

	path := fmt.Sprintf("%s/shared-data/%s", PATH_ROOT, models.ArrayOfStringToString(params.SharedDataId))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SharedData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SharedData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Nudm_Sdm Info for UPU service operation
// Description:
// Path: /:supi/am-data/upu-ack
// Path Params: supi
func UpuAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/upu-ack", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve multiple data sets
// Description:
// Path: /:supi
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetDataSetsParams struct {
	IfModifiedSince    string
	Supi               string
	DatasetNames       []string
	PlmnId             *models.PlmnIdNid
	DisasterRoamingInd *bool
	SupportedFeatures  string
	IfNoneMatch        string
}

func GetDataSets(cli sbi.ConsumerClient, params GetDataSetsParams) (headers map[string]string, rsp *models.SubscriptionDataSets, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if len(params.DatasetNames) == 0 {
		err = fmt.Errorf("dataset-names is required")
		return
	}

	path := fmt.Sprintf("%s/%s", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.AddParam("dataset-names", models.ArrayOfStringToString(params.DatasetNames))
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdNidToString(*params.PlmnId))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SubscriptionDataSets)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SubscriptionDataSets: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's SMF Selection Subscription Data
// Description:
// Path: /:supi/smf-select-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetSmfSelDataParams struct {
	IfNoneMatch        string
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *models.PlmnId
	DisasterRoamingInd *bool
}

func GetSmfSelData(cli sbi.ConsumerClient, params GetSmfSelDataParams) (headers map[string]string, rsp *models.SmfSelectionSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/smf-select-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SmfSelectionSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSelectionSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's SMS Subscription Data
// Description:
// Path: /:supi/sms-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetSmsDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	PlmnId            *models.PlmnId
}

func GetSmsData(cli sbi.ConsumerClient, params GetSmsDataParams) (headers map[string]string, rsp *models.SmsSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/sms-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SmsSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: subscribe to notifications
// Description:
// Path: /id/:ueId/sdm-subscriptions
// Path Params: ueId
// Response headers: Location
func Subscribe(cli sbi.ConsumerClient, ueId string, body *models.SdmSubscription) (headers map[string]string, rsp *models.SdmSubscription, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/id/%s/sdm-subscriptions", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.SdmSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SdmSubscription: %+v", err)
		}
	case 400, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: unsubscribe from notifications
// Description:
// Path: /id/:ueId/sdm-subscriptions/:subscriptionId
// Path Params: ueId, subscriptionId
type UnsubscribeParams struct {
	UeId           string
	SubscriptionId string
}

func Unsubscribe(cli sbi.ConsumerClient, params UnsubscribeParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	path := fmt.Sprintf("%s/id/%s/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubscriptionId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Nudm_Sdm Info service operation
// Description:
// Path: /:supi/am-data/sor-ack
// Path Params: supi
func SorAckInfo(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data/sor-ack", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the subscription
// Description:
// Path: /shared-data-subscriptions/:subscriptionId
// Path Params: subscriptionId
type ModifySharedDataSubsParams struct {
	SubscriptionId    string
	SupportedFeatures string
}

func ModifySharedDataSubs(cli sbi.ConsumerClient, params ModifySharedDataSubsParams, body *models.SdmSubsModification) (rsp *models.Schema, err error) {

	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/shared-data-subscriptions/%s", PATH_ROOT, params.SubscriptionId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.Schema)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Schema: %+v", err)
		}
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's UE Context In AMF Data
// Description:
// Path: /:supi/ue-context-in-amf-data
// Path Params: supi
type GetUeCtxInAmfDataParams struct {
	Supi              string
	SupportedFeatures string
}

func GetUeCtxInAmfData(cli sbi.ConsumerClient, params GetUeCtxInAmfDataParams) (rsp *models.UeContextInAmfData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/ue-context-in-amf-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UeContextInAmfData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextInAmfData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's Access and Mobility Subscription Data
// Description:
// Path: /:supi/am-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetAmDataParams struct {
	DisasterRoamingInd *bool
	IfNoneMatch        string
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *models.PlmnIdNid
	AdjacentPlmns      []models.PlmnId
}

func GetAmData(cli sbi.ConsumerClient, params GetAmDataParams) (headers map[string]string, rsp *models.AccessAndMobilitySubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/am-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdNidToString(*params.PlmnId))
	}
	if len(params.AdjacentPlmns) > 0 {
		request.AddParam("adjacent-plmns", models.ArrayOfPlmnIdToString(params.AdjacentPlmns))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.AccessAndMobilitySubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AccessAndMobilitySubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's UE Context In SMF Data
// Description:
// Path: /:supi/ue-context-in-smf-data
// Path Params: supi
type GetUeCtxInSmfDataParams struct {
	SupportedFeatures string
	Supi              string
}

func GetUeCtxInSmfData(cli sbi.ConsumerClient, params GetUeCtxInSmfDataParams) (rsp *models.UeContextInSmfData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/ue-context-in-smf-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UeContextInSmfData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextInSmfData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's Trace Configuration Data
// Description:
// Path: /:supi/trace-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetTraceConfigDataParams struct {
	PlmnId            *models.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetTraceConfigData(cli sbi.ConsumerClient, params GetTraceConfigDataParams) (headers map[string]string, rsp *models.TraceDataResponse, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/trace-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.TraceDataResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode TraceDataResponse: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve a UE's SMS Management Subscription Data
// Description:
// Path: /:supi/sms-mng-data
// Path Params: supi
// Response headers: Cache-Control, ETag, Last-Modified
type GetSmsMngtDataParams struct {
	PlmnId            *models.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetSmsMngtData(cli sbi.ConsumerClient, params GetSmsMngtDataParams) (headers map[string]string, rsp *models.SmsManagementSubscriptionData, err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	path := fmt.Sprintf("%s/%s/sms-mng-data", PATH_ROOT, params.Supi)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		headers = response.GetHeaders()
		rsp = new(models.SmsManagementSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsManagementSubscriptionData: %+v", err)
		}
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
