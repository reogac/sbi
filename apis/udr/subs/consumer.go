/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = ""
)

// Summary: Create individual sdm subscription
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions
// Path Params: ueId
// Response headers: Location
func CreateSdmSubscriptions(cli sbi.ConsumerClient, ueId string, body *models.SdmSubscription) (headers map[string]string, rsp *models.SdmSubscription, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions", PATH_ROOT, ueId)
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
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the subscribed V2X Data of a UE
// Description:
// Path: /subscription-data/:ueId/v2x-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryV2xDataParams struct {
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryV2xData(cli sbi.ConsumerClient, params QueryV2xDataParams) (headers map[string]string, rsp *models.V2xSubscriptionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/v2x-data", PATH_ROOT, params.UeId)
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
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the data of a 5G MBS Group
// Description:
// Path: /subscription-data/group-data/mbs-group-membership
// Path Params:
func Query5GmbsGroup(cli sbi.ConsumerClient, gpsis []string) (rsp *map[string]models.MulticastMbsGroupMemb, err error) {

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(gpsis) > 0 {
		request.AddParam("gpsis", models.ArrayOfStringToString(gpsis))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(map[string]models.MulticastMbsGroupMemb)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]MulticastMbsGroupMemb: %+v", err)
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

// Summary: Retrieves the SMSF context data of a UE using non-3gpp access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-non-3gpp-access
// Path Params: ueId
type QuerySmsfContextNon3gppParams struct {
	Fields            []string
	SupportedFeatures string
	UeId              string
}

func QuerySmsfContextNon3gpp(cli sbi.ConsumerClient, params QuerySmsfContextNon3gppParams) (rsp *models.SmsfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.SmsfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsfRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Subscription data subscriptions
// Description:
// Path: /subscription-data/subs-to-notify
// Path Params:
// Response headers: Location
func SubscriptionDataSubscriptions(cli sbi.ConsumerClient, body *models.SubscriptionDataSubscriptions) (headers map[string]string, rsp *models.SubscriptionDataSubscriptions, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.SubscriptionDataSubscriptions)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SubscriptionDataSubscriptions: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify an individual sdm subscription
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId
// Path Params: ueId, subsId
type ModifysdmSubscriptionParams struct {
	UeId              string
	SubsId            string
	SupportedFeatures string
}

func ModifysdmSubscription(cli sbi.ConsumerClient, params ModifysdmSubscriptionParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403, 404:
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

// Summary: Delete HSS SDM Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions
// Path Params: ueId, subsId
type RemoveHssSDMSubscriptionsInfoParams struct {
	UeId   string
	SubsId string
}

func RemoveHssSDMSubscriptionsInfo(cli sbi.ConsumerClient, params RemoveHssSDMSubscriptionsInfoParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s/hss-sdm-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve multiple context data sets of a UE
// Description:
// Path: /subscription-data/:ueId/context-data
// Path Params: ueId
type QueryContextDataParams struct {
	UeId                string
	ContextDatasetNames []string
}

func QueryContextData(cli sbi.ConsumerClient, params QueryContextDataParams) (rsp *models.ContextDataSets, err error) {

	if len(params.ContextDatasetNames) == 0 {
		err = fmt.Errorf("context-dataset-names is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("context-dataset-names", models.ArrayOfStringToString(params.ContextDatasetNames))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ContextDataSets)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ContextDataSets: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the LCS Privacy subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/lcs-privacy-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryLcsPrivacyDataParams struct {
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	Fields            []string
}

func QueryLcsPrivacyData(cli sbi.ConsumerClient, params QueryLcsPrivacyDataParams) (headers map[string]string, rsp *models.LcsPrivacyData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/lcs-privacy-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
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
		rsp = new(models.LcsPrivacyData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsPrivacyData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the UPU acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais
// Path Params: ueId
type QueryNssaiAckParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryNssaiAck(cli sbi.ConsumerClient, params QueryNssaiAckParams) (rsp *models.NssaiAckData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/subscribed-snssais", PATH_ROOT, params.UeId)
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
		rsp = new(models.NssaiAckData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode NssaiAckData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the parameter provision profile data of a UE
// Description:
// Path: /subscription-data/:ueId/pp-profile-data
// Path Params: ueId
type QueryPPDataParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryPPData(cli sbi.ConsumerClient, params QueryPPDataParams) (rsp *models.PpProfileData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-profile-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.PpProfileData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PpProfileData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the ee subscriptions of a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions
// Path Params: ueGroupId
type QueryEeGroupSubscriptionsParams struct {
	UeGroupId         string
	SupportedFeatures string
}

func QueryEeGroupSubscriptions(cli sbi.ConsumerClient, params QueryEeGroupSubscriptionsParams) (rsp *[]models.EeSubscription, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions", PATH_ROOT, params.UeGroupId)
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
		rsp = new([]models.EeSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []EeSubscription: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the UPU acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/upu-data
// Path Params: ueId
type QueryAuthUPUParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryAuthUPU(cli sbi.ConsumerClient, params QueryAuthUPUParams) (rsp *models.UpuData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/upu-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.UpuData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UpuData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create HSS Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions
// Path Params: ueId, subsId
type CreateHSSSubscriptionsParams struct {
	UeId   string
	SubsId string
}

func CreateHSSSubscriptions(cli sbi.ConsumerClient, params CreateHSSSubscriptionsParams, body *models.HssSubscriptionInfo) (rsp *models.HssSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/hss-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.HssSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode HssSubscriptionInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete Service Specific Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType
// Path Params: ueId, serviceType
type RemoveServiceSpecificAuthorizationInfoParams struct {
	UeId        string
	ServiceType string
}

func RemoveServiceSpecificAuthorizationInfo(cli sbi.ConsumerClient, params RemoveServiceSpecificAuthorizationInfoParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServiceType) == 0 {
		err = fmt.Errorf("serviceType is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/service-specific-authorizations/%s", PATH_ROOT, params.UeId, params.ServiceType)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the Roaming Information of the EPC domain
// Description:
// Path: /subscription-data/:ueId/context-data/roaming-information
// Path Params: ueId
func QueryRoamingInformation(cli sbi.ConsumerClient, ueId string) (rsp *models.RoamingInfoUpdate, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/roaming-information", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.RoamingInfoUpdate)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode RoamingInfoUpdate: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the data of 5G MBS Group
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/internal
// Path Params:
func Query5GMbsGroupInternal(cli sbi.ConsumerClient, internalGroupIds []string) (rsp *map[string]models.MulticastMbsGroupMemb, err error) {

	if len(internalGroupIds) == 0 {
		err = fmt.Errorf("internal-group-ids is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/internal", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("internal-group-ids", models.ArrayOfStringToString(internalGroupIds))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(map[string]models.MulticastMbsGroupMemb)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]MulticastMbsGroupMemb: %+v", err)
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

// Summary: Deletes AMF Subscription Info for an eeSubscription
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueId, subsId
type RemoveAmfSubscriptionsInfoParams struct {
	SubsId string
	UeId   string
}

func RemoveAmfSubscriptionsInfo(cli sbi.ConsumerClient, params RemoveAmfSubscriptionsInfoParams) (err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the access and mobility subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/am-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryAmDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	ServingPlmnId     string
	Fields            []string
	SupportedFeatures string
}

func QueryAmData(cli sbi.ConsumerClient, params QueryAmDataParams) (headers map[string]string, rsp *models.AccessAndMobilitySubscriptionData, err error) {

	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/am-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
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
		rsp = new(models.AccessAndMobilitySubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AccessAndMobilitySubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the AMF context data of a UE using non-3gpp access
// Description:
// Path: /subscription-data/:ueId/context-data/amf-non-3gpp-access
// Path Params: ueId
type QueryAmfContextNon3gppParams struct {
	SupportedFeatures string
	UeId              string
	Fields            []string
}

func QueryAmfContextNon3gpp(cli sbi.ConsumerClient, params QueryAmfContextNon3gppParams) (rsp *models.AmfNon3GppAccessRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.AmfNon3GppAccessRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AmfNon3GppAccessRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the SMS subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QuerySmsDataParams struct {
	IfModifiedSince   string
	UeId              string
	ServingPlmnId     string
	SupportedFeatures string
	IfNoneMatch       string
}

func QuerySmsData(cli sbi.ConsumerClient, params QuerySmsDataParams) (headers map[string]string, rsp *models.SmsSubscriptionData, err error) {

	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/sms-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
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
		rsp = new(models.SmsSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsSubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve AMF subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueId, subsId
type GetAmfSubscriptionInfoParams struct {
	UeId   string
	SubsId string
}

func GetAmfSubscriptionInfo(cli sbi.ConsumerClient, params GetAmfSubscriptionInfoParams) (rsp *[]models.AmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new([]models.AmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []AmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update an individual sdm subscriptions of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId
// Path Params: ueId, subsId
type UpdatesdmsubscriptionsParams struct {
	SubsId string
	UeId   string
}

func Updatesdmsubscriptions(cli sbi.ConsumerClient, params UpdatesdmsubscriptionsParams, body *models.SdmSubscription) (err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 404:
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

// Summary: Deletes a sdmsubscriptions
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId
// Path Params: ueId, subsId
type RemovesdmSubscriptionsParams struct {
	SubsId string
	UeId   string
}

func RemovesdmSubscriptions(cli sbi.ConsumerClient, params RemovesdmSubscriptionsParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 404:
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

// Summary: Update the Roaming Information of the EPC domain
// Description:
// Path: /subscription-data/:ueId/context-data/roaming-information
// Path Params: ueId
func UpdateRoamingInformation(cli sbi.ConsumerClient, ueId string, body *models.RoamingInfoUpdate) (rsp *models.RoamingInfoUpdate, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/roaming-information", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.RoamingInfoUpdate)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode RoamingInfoUpdate: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the CAG acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag
// Path Params: ueId
type QueryCagAckParams struct {
	SupportedFeatures string
	UeId              string
}

func QueryCagAck(cli sbi.ConsumerClient, params QueryCagAckParams) (rsp *models.CagAckData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/subscribed-cag", PATH_ROOT, params.UeId)
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
		rsp = new(models.CagAckData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode CagAckData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes AMF Subscription Info for an eeSubscription for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueGroupId, subsId
type RemoveAmfGroupSubscriptionsParams struct {
	SubsId    string
	UeGroupId string
}

func RemoveAmfGroupSubscriptions(cli sbi.ConsumerClient, params RemoveAmfGroupSubscriptionsParams) (err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove the SMSF context data of a UE via 3GPP access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-3gpp-access
// Path Params: ueId
func DeleteSmsfContext3gpp(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Mapping of Group Identifiers
// Description:
// Path: /subscription-data/group-data/group-identifiers
// Path Params:
type GetGroupIdentifiersParams struct {
	UeIdInd           *bool
	SupportedFeatures string
	ExtGroupId        string
	IntGroupId        string
}

func GetGroupIdentifiers(cli sbi.ConsumerClient, params GetGroupIdentifiersParams) (rsp *models.GroupIdentifiers, err error) {

	path := fmt.Sprintf("%s/subscription-data/group-data/group-identifiers", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
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
		rsp = new(models.GroupIdentifiers)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode GroupIdentifiers: %+v", err)
		}
	case 403:
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

// Summary: Create an individual 5G VN Grouop
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/:externalGroupId
// Path Params: externalGroupId
func Create5GVnGroup(cli sbi.ConsumerClient, externalGroupId string, body *models.FiveGVnGroupConfiguration) (rsp *models.FiveGVnGroupConfiguration, err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.FiveGVnGroupConfiguration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode FiveGVnGroupConfiguration: %+v", err)
		}
	case 403:
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

// Summary: Create NIDD Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/nidd-authorizations
// Path Params: ueId
func CreateNIDDAuthorizationInfo(cli sbi.ConsumerClient, ueId string, body *models.NiddAuthorizationInfo) (rsp *models.NiddAuthorizationInfo, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/nidd-authorizations", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.NiddAuthorizationInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode NiddAuthorizationInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove the Authentication Status of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status
// Path Params: ueId
func DeleteAuthenticationStatus(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To store the AMF context data of a UE using non-3gpp access in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/amf-non-3gpp-access
// Path Params: ueId
// Response headers: Location
func CreateAmfContextNon3gpp(cli sbi.ConsumerClient, ueId string, body *models.AmfNon3GppAccessRegistration) (headers map[string]string, rsp *models.Amf3GppAccessRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-non-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.Amf3GppAccessRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Amf3GppAccessRegistration: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the Message Waiting Data of the UE
// Description:
// Path: /subscription-data/:ueId/context-data/mwd
// Path Params: ueId
type QueryMessageWaitingDataParams struct {
	Fields            []string
	SupportedFeatures string
	UeId              string
}

func QueryMessageWaitingData(cli sbi.ConsumerClient, params QueryMessageWaitingDataParams) (rsp *models.MessageWaitingData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/mwd", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.MessageWaitingData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MessageWaitingData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: get a Parameter Provisioning Data Entry
// Description:
// Path: /subscription-data/:ueId/pp-data-store/:afInstanceId
// Path Params: ueId, afInstanceId
type GetPPDataEntryParams struct {
	UeId              string
	AfInstanceId      string
	SupportedFeatures string
}

func GetPPDataEntry(cli sbi.ConsumerClient, params GetPPDataEntryParams) (rsp *models.PpDataEntry, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.AfInstanceId) == 0 {
		err = fmt.Errorf("afInstanceId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data-store/%s", PATH_ROOT, params.UeId, params.AfInstanceId)
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
		rsp = new(models.PpDataEntry)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PpDataEntry: %+v", err)
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

// Summary: Retrieves the ee subscriptions of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions
// Path Params: ueId
type QueryeesubscriptionsParams struct {
	SupportedFeatures string
	EventTypes        []string
	NfIdentifiers     []models.NfIdentifier
	UeId              string
}

func Queryeesubscriptions(cli sbi.ConsumerClient, params QueryeesubscriptionsParams) (rsp *[]models.EeSubscriptionExt, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.EventTypes) > 0 {
		request.AddParam("event-types", models.ArrayOfStringToString(params.EventTypes))
	}
	if len(params.NfIdentifiers) > 0 {
		request.AddParam("nf-identifiers", models.ArrayOfNfIdentifierToString(params.NfIdentifiers))
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
		rsp = new([]models.EeSubscriptionExt)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []EeSubscriptionExt: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create AmfSubscriptions for an individual ee subscriptions of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueId, subsId
type CreateAMFSubscriptionsParams struct {
	UeId   string
	SubsId string
}

func CreateAMFSubscriptions(cli sbi.ConsumerClient, params CreateAMFSubscriptionsParams, body *[]models.AmfSubscriptionInfo) (rsp *[]models.AmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new([]models.AmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []AmfSubscriptionInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve identity data by SUPI or GPSI
// Description:
// Path: /subscription-data/:ueId/identity-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type GetIdentityDataParams struct {
	IfModifiedSince string
	UeId            string
	AppPortId       *models.AppPortId
	IfNoneMatch     string
}

func GetIdentityData(cli sbi.ConsumerClient, params GetIdentityDataParams) (headers map[string]string, rsp *models.IdentityData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/identity-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.AppPortId != nil {
		request.AddParam("app-port-id", models.AppPortIdToString(*params.AppPortId))
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
		rsp = new(models.IdentityData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode IdentityData: %+v", err)
		}
	case 403:
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

// Summary: Retrieve a 5GVnGroup configuration
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/:externalGroupId
// Path Params: externalGroupId
func Get5GVnGroupConfiguration(cli sbi.ConsumerClient, externalGroupId string) (rsp *models.FiveGVnGroupConfiguration, err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.FiveGVnGroupConfiguration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode FiveGVnGroupConfiguration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the SMF selection subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/smf-selection-subscription-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QuerySmfSelectDataParams struct {
	Fields            []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	ServingPlmnId     string
}

func QuerySmfSelectData(cli sbi.ConsumerClient, params QuerySmfSelectDataParams) (headers map[string]string, rsp *models.SmfSelectionSubscriptionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/smf-selection-subscription-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
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
		rsp = new(models.SmfSelectionSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSelectionSubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the AMF context data of a UE using 3gpp access
// Description:
// Path: /subscription-data/:ueId/context-data/amf-3gpp-access
// Path Params: ueId
type QueryAmfContext3gppParams struct {
	SupportedFeatures string
	UeId              string
	Fields            []string
}

func QueryAmfContext3gpp(cli sbi.ConsumerClient, params QueryAmfContext3gppParams) (rsp *models.Amf3GppAccessRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.Amf3GppAccessRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Amf3GppAccessRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the SMF registration list of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/smf-registrations
// Path Params: ueId
type QuerySmfRegListParams struct {
	UeId              string
	SupportedFeatures string
}

func QuerySmfRegList(cli sbi.ConsumerClient, params QuerySmfRegListParams) (rsp *[]models.SmfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smf-registrations", PATH_ROOT, params.UeId)
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
		rsp = new([]models.SmfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []SmfRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the subscribed User Consent Data of a UE
// Description:
// Path: /subscription-data/:ueId/uc-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryUserConsentDataParams struct {
	UeId              string
	SupportedFeatures string
	UcPurpose         string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryUserConsentData(cli sbi.ConsumerClient, params QueryUserConsentDataParams) (headers map[string]string, rsp *models.UcSubscriptionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/uc-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.UcPurpose) > 0 {
		request.AddParam("ucPurpose", params.UcPurpose)
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
		rsp = new(models.UcSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UcSubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Updates the ME support of SOR CMCI information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/sor-data
// Path Params: ueId
type UpdateAuthenticationSoRParams struct {
	UeId              string
	SupportedFeatures string
}

func UpdateAuthenticationSoR(cli sbi.ConsumerClient, params UpdateAuthenticationSoRParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/sor-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: To remove an operator-specific data resource of a UE
// Description:
// Path: /subscription-data/:ueId/operator-specific-data
// Path Params: ueId
func DeleteOperSpecData(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/operator-specific-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 403:
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

// Summary: To remove the SMSF context data of a UE via non-3GPP access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-non-3gpp-access
// Path Params: ueId
func DeleteSmsfContextNon3gpp(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-non-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify an individual ee subscription of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId
// Path Params: ueId, subsId
type ModifyEesubscriptionParams struct {
	UeId              string
	SubsId            string
	SupportedFeatures string
}

func ModifyEesubscription(cli sbi.ConsumerClient, params ModifyEesubscriptionParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403, 404:
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

// Summary: Modify an individual ee subscription for a group of a UEs
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId
// Path Params: ueGroupId, subsId
type ModifyEeGroupSubscriptionParams struct {
	SupportedFeatures string
	UeGroupId         string
	SubsId            string
}

func ModifyEeGroupSubscription(cli sbi.ConsumerClient, params ModifyEeGroupSubscriptionParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s", PATH_ROOT, params.UeGroupId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403, 404:
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

// Summary: Create HSS SDM Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions
// Path Params: ueId, subsId
type CreateHSSSDMSubscriptionsParams struct {
	UeId   string
	SubsId string
}

func CreateHSSSDMSubscriptions(cli sbi.ConsumerClient, params CreateHSSSDMSubscriptionsParams, body *models.HssSubscriptionInfo) (err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s/hss-sdm-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To create an individual SMF context data of a UE in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type CreateOrUpdateSmfRegistrationParams struct {
	PduSessionId int
	UeId         string
}

func CreateOrUpdateSmfRegistration(cli sbi.ConsumerClient, params CreateOrUpdateSmfRegistrationParams, body *models.SmfRegistration) (rsp *models.SmfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SmfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfRegistration: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve HSS SDM Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions
// Path Params: ueId, subsId
type GetHssSDMSubscriptionInfoParams struct {
	UeId   string
	SubsId string
}

func GetHssSDMSubscriptionInfo(cli sbi.ConsumerClient, params GetHssSDMSubscriptionInfoParams) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s/hss-sdm-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes a subscriptionDataSubscriptions
// Description:
// Path: /subscription-data/subs-to-notify/:subsId
// Path Params: subsId
func RemovesubscriptionDataSubscriptions(cli sbi.ConsumerClient, subsId string) (err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify/%s", PATH_ROOT, subsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create SMF Subscription Info for a group of UEs or any YE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueGroupId, subsId
type CreateSmfGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func CreateSmfGroupSubscriptions(cli sbi.ConsumerClient, params CreateSmfGroupSubscriptionsParams, body *models.SmfSubscriptionInfo) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve a 5GmbsGroup
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/:externalGroupId
// Path Params: externalGroupId
func GetMulticastMbsGroupMemb(cli sbi.ConsumerClient, externalGroupId string) (rsp *models.MulticastMbsGroupMemb, err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.MulticastMbsGroupMemb)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MulticastMbsGroupMemb: %+v", err)
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

// Summary: Delete HSS Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions
// Path Params: ueId, subsId
type RemoveHssSubscriptionsInfoParams struct {
	SubsId string
	UeId   string
}

func RemoveHssSubscriptionsInfo(cli sbi.ConsumerClient, params RemoveHssSubscriptionsInfoParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/hss-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the Individual Authentication Status of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName
// Path Params: ueId, servingNetworkName
type QueryIndividualAuthenticationStatusParams struct {
	Fields             []string
	SupportedFeatures  string
	UeId               string
	ServingNetworkName string
}

func QueryIndividualAuthenticationStatus(cli sbi.ConsumerClient, params QueryIndividualAuthenticationStatusParams) (rsp *models.AuthEvent, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingNetworkName) == 0 {
		err = fmt.Errorf("servingNetworkName is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status/%s", PATH_ROOT, params.UeId, params.ServingNetworkName)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.AuthEvent)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AuthEvent: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To modify the AMF context data of a UE using 3gpp access in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/amf-3gpp-access
// Path Params: ueId
type AmfContext3gppParams struct {
	UeId              string
	SupportedFeatures string
}

func AmfContext3gpp(cli sbi.ConsumerClient, params AmfContext3gppParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-3gpp-access", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: To store the Authentication Status data of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status
// Path Params: ueId
func CreateAuthenticationStatus(cli sbi.ConsumerClient, ueId string, body *models.AuthEvent) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the SMS management subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-mng-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QuerySmsMngDataParams struct {
	ServingPlmnId     string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
}

func QuerySmsMngData(cli sbi.ConsumerClient, params QuerySmsMngDataParams) (headers map[string]string, rsp *models.SmsManagementSubscriptionData, err error) {

	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/sms-mng-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
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
		rsp = new(models.SmsManagementSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsManagementSubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve a individual eeSubscription for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId
// Path Params: ueGroupId, subsId
type QueryEeGroupSubscriptionParams struct {
	UeGroupId string
	SubsId    string
}

func QueryEeGroupSubscription(cli sbi.ConsumerClient, params QueryEeGroupSubscriptionParams) (err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes subscriptions identified by a given ue-id parameter
// Description:
// Path: /subscription-data/subs-to-notify
// Path Params:
type RemoveMultipleSubscriptionDataSubscriptionsParams struct {
	NfInstanceId                  string
	DeleteAllNfs                  *bool
	ImplicitUnsubscribeIndication *bool
	UeId                          string
}

func RemoveMultipleSubscriptionDataSubscriptions(cli sbi.ConsumerClient, params RemoveMultipleSubscriptionDataSubscriptionsParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ue-id is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	request.AddParam("ue-id", params.UeId)
	if len(params.NfInstanceId) > 0 {
		request.AddParam("nf-instance-id", params.NfInstanceId)
	}
	if params.DeleteAllNfs != nil {
		request.AddParam("delete-all-nfs", models.BoolToString(*params.DeleteAllNfs))
	}
	if params.ImplicitUnsubscribeIndication != nil {
		request.AddParam("implicit-unsubscribe-indication", models.BoolToString(*params.ImplicitUnsubscribeIndication))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the data of 5G VN Group
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/internal
// Path Params:
func Query5GVnGroupInternal(cli sbi.ConsumerClient, internalGroupIds []string) (rsp *map[string]models.FiveGVnGroupConfiguration, err error) {

	if len(internalGroupIds) == 0 {
		err = fmt.Errorf("internal-group-ids is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/internal", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("internal-group-ids", models.ArrayOfStringToString(internalGroupIds))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(map[string]models.FiveGVnGroupConfiguration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]FiveGVnGroupConfiguration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create AmfSubscriptions for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueGroupId, subsId
type CreateAmfGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func CreateAmfGroupSubscriptions(cli sbi.ConsumerClient, params CreateAmfGroupSubscriptionsParams, body *[]models.AmfSubscriptionInfo) (rsp *[]models.AmfSubscriptionInfo, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new([]models.AmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []AmfSubscriptionInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes the 5GmbsGroup
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/:externalGroupId
// Path Params: externalGroupId
func Delete5GmbsGroup(cli sbi.ConsumerClient, externalGroupId string) (err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
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

// Summary: Retrieves the operator specific data of a UE
// Description:
// Path: /subscription-data/:ueId/operator-specific-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryOperSpecDataParams struct {
	UeId              string
	Fields            []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryOperSpecData(cli sbi.ConsumerClient, params QueryOperSpecDataParams) (headers map[string]string, rsp *map[string]models.OperatorSpecificDataContainer, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/operator-specific-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
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
		rsp = new(map[string]models.OperatorSpecificDataContainer)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]OperatorSpecificDataContainer: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the 5GmbsGroup
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/:externalGroupId
// Path Params: externalGroupId
type Modify5GmbsGroupParams struct {
	ExternalGroupId   string
	SupportedFeatures string
}

func Modify5GmbsGroup(cli sbi.ConsumerClient, params Modify5GmbsGroupParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.ExternalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/%s", PATH_ROOT, params.ExternalGroupId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 502, 503:
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

// Summary: Delete NIDD Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/nidd-authorizations
// Path Params: ueId
func RemoveNiddAuthorizationInfo(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/nidd-authorizations", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the AMF Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueId, subsId
type ModifyAmfSubscriptionInfoParams struct {
	UeId              string
	SubsId            string
	SupportedFeatures string
}

func ModifyAmfSubscriptionInfo(cli sbi.ConsumerClient, params ModifyAmfSubscriptionInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieves a individual subscriptionDataSubscription identified by subsId
// Description:
// Path: /subscription-data/subs-to-notify/:subsId
// Path Params: subsId
func QuerySubscriptionDataSubscriptions(cli sbi.ConsumerClient, subsId string) (rsp *models.SubscriptionDataSubscriptions, err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify/%s", PATH_ROOT, subsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SubscriptionDataSubscriptions)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SubscriptionDataSubscriptions: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create Service Specific Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType
// Path Params: ueId, serviceType
type CreateServiceSpecificAuthorizationInfoParams struct {
	UeId        string
	ServiceType string
}

func CreateServiceSpecificAuthorizationInfo(cli sbi.ConsumerClient, params CreateServiceSpecificAuthorizationInfoParams, body *models.ServiceSpecificAuthorizationInfo) (rsp *models.ServiceSpecificAuthorizationInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServiceType) == 0 {
		err = fmt.Errorf("serviceType is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/service-specific-authorizations/%s", PATH_ROOT, params.UeId, params.ServiceType)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.ServiceSpecificAuthorizationInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ServiceSpecificAuthorizationInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify Service Specific Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType
// Path Params: ueId, serviceType
type ModifyServiceSpecificAuthorizationInfoParams struct {
	SupportedFeatures string
	UeId              string
	ServiceType       string
}

func ModifyServiceSpecificAuthorizationInfo(cli sbi.ConsumerClient, params ModifyServiceSpecificAuthorizationInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServiceType) == 0 {
		err = fmt.Errorf("serviceType is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/service-specific-authorizations/%s", PATH_ROOT, params.UeId, params.ServiceType)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: To modify the AMF context data of a UE using non 3gpp access in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/amf-non-3gpp-access
// Path Params: ueId
type AmfContextNon3gppParams struct {
	UeId              string
	SupportedFeatures string
}

func AmfContextNon3gpp(cli sbi.ConsumerClient, params AmfContextNon3gppParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-non-3gpp-access", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: To store the SoR acknowledgement information of a UE and ME support of SOR CMCI
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/sor-data
// Path Params: ueId
type CreateAuthenticationSoRParams struct {
	SupportedFeatures string
	UeId              string
}

func CreateAuthenticationSoR(cli sbi.ConsumerClient, params CreateAuthenticationSoRParams, body *models.SorData) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/sor-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create individual EE subscription
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions
// Path Params: ueId
// Response headers: Location
func CreateEeSubscriptions(cli sbi.ConsumerClient, ueId string, body *models.EeSubscription) (headers map[string]string, rsp *models.EeSubscription, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.EeSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EeSubscription: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update an individual ee subscriptions of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId
// Path Params: ueId, subsId
type UpdateEesubscriptionsParams struct {
	UeId   string
	SubsId string
}

func UpdateEesubscriptions(cli sbi.ConsumerClient, params UpdateEesubscriptionsParams, body *models.EeSubscription) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 404:
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

// Summary: Update an individual ee subscription of a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId
// Path Params: ueGroupId, subsId
type UpdateEeGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func UpdateEeGroupSubscriptions(cli sbi.ConsumerClient, params UpdateEeGroupSubscriptionsParams, body *models.EeSubscription) (err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 404:
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

// Summary: Deletes a eeSubscription for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId
// Path Params: ueGroupId, subsId
type RemoveEeGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func RemoveEeGroupSubscriptions(cli sbi.ConsumerClient, params RemoveEeGroupSubscriptionsParams) (err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve ODB Data data by SUPI or GPSI
// Description:
// Path: /subscription-data/:ueId/operator-determined-barring-data
// Path Params: ueId
func GetOdbData(cli sbi.ConsumerClient, ueId string) (rsp *models.OdbData, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/operator-determined-barring-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.OdbData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode OdbData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create an individual 5G MBS Grouop
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/:externalGroupId
// Path Params: externalGroupId
func Create5GmbsGroup(cli sbi.ConsumerClient, externalGroupId string, body *models.MulticastMbsGroupMemb) (rsp *models.MulticastMbsGroupMemb, err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.MulticastMbsGroupMemb)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MulticastMbsGroupMemb: %+v", err)
		}
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 502, 503:
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

// Summary: Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/sor-data
// Path Params: ueId
type QueryAuthSoRParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryAuthSoR(cli sbi.ConsumerClient, params QueryAuthSoRParams) (rsp *models.SorData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/sor-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.SorData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SorData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the IP-SM-GW context data of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ip-sm-gw
// Path Params: ueId
type QueryIpSmGwContextParams struct {
	UeId              string
	Fields            []string
	SupportedFeatures string
}

func QueryIpSmGwContext(cli sbi.ConsumerClient, params QueryIpSmGwContextParams) (rsp *models.IpSmGwRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ip-sm-gw", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.IpSmGwRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode IpSmGwRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove the Message Waiting Data of the UE
// Description:
// Path: /subscription-data/:ueId/context-data/mwd
// Path Params: ueId
func DeleteMessageWaitingData(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/mwd", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: get a list of Parameter Provisioning Data Entries
// Description:
// Path: /subscription-data/:ueId/pp-data-store
// Path Params: ueId
type GetMultiplePPDataEntriesParams struct {
	UeId              string
	SupportedFeatures string
}

func GetMultiplePPDataEntries(cli sbi.ConsumerClient, params GetMultiplePPDataEntriesParams) (rsp *models.PpDataEntryList, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data-store", PATH_ROOT, params.UeId)
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
		rsp = new(models.PpDataEntryList)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PpDataEntryList: %+v", err)
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

// Summary: Delete SMF Subscription Info for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueGroupId, subsId
type RemoveSmfGroupSubscriptionsParams struct {
	SubsId    string
	UeGroupId string
}

func RemoveSmfGroupSubscriptions(cli sbi.ConsumerClient, params RemoveSmfGroupSubscriptionsParams) (err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify SMF Subscription Info for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueGroupId, subsId
type ModifySmfGroupSubscriptionsParams struct {
	UeGroupId         string
	SubsId            string
	SupportedFeatures string
}

func ModifySmfGroupSubscriptions(cli sbi.ConsumerClient, params ModifySmfGroupSubscriptionsParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: To store the NSSAI update acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais
// Path Params: ueId
type CreateOrUpdateNssaiAckParams struct {
	UeId              string
	SupportedFeatures string
}

func CreateOrUpdateNssaiAck(cli sbi.ConsumerClient, params CreateOrUpdateNssaiAckParams, body *models.NssaiAckData) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/subscribed-snssais", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create the SMSF context data of a UE via 3GPP access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-3gpp-access
// Path Params: ueId
func CreateSmsfContext3gpp(cli sbi.ConsumerClient, ueId string, body *models.SmsfRegistration) (rsp *models.SmsfRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsfRegistration: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create the SMSF context data of a UE via non-3GPP access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-non-3gpp-access
// Path Params: ueId
func CreateSmsfContextNon3gpp(cli sbi.ConsumerClient, ueId string, body *models.SmsfRegistration) (rsp *models.SmsfRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-non-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsfRegistration: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove the IP-SM-GW context data of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ip-sm-gw
// Path Params: ueId
func DeleteIpSmGwContext(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Read the profile of a given UE
// Description:
// Path: /subscription-data/:ueId/pp-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type GetppDataParams struct {
	IfModifiedSince   string
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
}

func GetppData(cli sbi.ConsumerClient, params GetppDataParams) (headers map[string]string, rsp *models.PpData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.PpData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PpData: %+v", err)
		}
	case 403:
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

// Summary: Retrieves the ee profile data of a UE
// Description:
// Path: /subscription-data/:ueId/ee-profile-data
// Path Params: ueId
type QueryEEDataParams struct {
	Fields            []string
	SupportedFeatures string
	UeId              string
}

func QueryEEData(cli sbi.ConsumerClient, params QueryEEDataParams) (rsp *models.EeProfileData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ee-profile-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.EeProfileData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EeProfileData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves a individual sdmSubscription identified by subsId
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId
// Path Params: ueId, subsId
type QuerysdmSubscriptionParams struct {
	UeId   string
	SubsId string
}

func QuerysdmSubscription(cli sbi.ConsumerClient, params QuerysdmSubscriptionParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the trace configuration data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/trace-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryTraceDataParams struct {
	UeId            string
	ServingPlmnId   string
	IfNoneMatch     string
	IfModifiedSince string
}

func QueryTraceData(cli sbi.ConsumerClient, params QueryTraceDataParams) (headers map[string]string, rsp *models.TraceData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/trace-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
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
		rsp = new(models.TraceData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode TraceData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the individual SMF registration of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type QuerySmfRegistrationParams struct {
	PduSessionId      int
	Fields            []string
	SupportedFeatures string
	UeId              string
}

func QuerySmfRegistration(cli sbi.ConsumerClient, params QuerySmfRegistrationParams) (rsp *models.SmfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.SmfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To store the CAG update acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag
// Path Params: ueId
type CreateCagUpdateAckParams struct {
	UeId              string
	SupportedFeatures string
}

func CreateCagUpdateAck(cli sbi.ConsumerClient, params CreateCagUpdateAckParams, body *models.CagAckData) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/subscribed-cag", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve multiple provisioned data sets of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data
// Path Params: ueId, servingPlmnId
type QueryProvisionedDataParams struct {
	DatasetNames  []string
	UeId          string
	ServingPlmnId string
}

func QueryProvisionedData(cli sbi.ConsumerClient, params QueryProvisionedDataParams) (rsp *models.ProvisionedDataSets, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.DatasetNames) > 0 {
		request.AddParam("dataset-names", models.ArrayOfStringToString(params.DatasetNames))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ProvisionedDataSets)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ProvisionedDataSets: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create the Message Waiting Data of the UE
// Description:
// Path: /subscription-data/:ueId/context-data/mwd
// Path Params: ueId
func CreateMessageWaitingData(cli sbi.ConsumerClient, ueId string, body *models.MessageWaitingData) (rsp *models.MessageWaitingData, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/mwd", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.MessageWaitingData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MessageWaitingData: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete a Provisioning Parameter Data Entry
// Description:
// Path: /subscription-data/:ueId/pp-data-store/:afInstanceId
// Path Params: ueId, afInstanceId
type DeletePPDataEntryParams struct {
	AfInstanceId string
	UeId         string
}

func DeletePPDataEntry(cli sbi.ConsumerClient, params DeletePPDataEntryParams) (err error) {

	if len(params.AfInstanceId) == 0 {
		err = fmt.Errorf("afInstanceId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data-store/%s", PATH_ROOT, params.UeId, params.AfInstanceId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
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

// Summary: Retrieves the Authentication Status of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status
// Path Params: ueId
type QueryAuthenticationStatusParams struct {
	SupportedFeatures string
	UeId              string
	Fields            []string
}

func QueryAuthenticationStatus(cli sbi.ConsumerClient, params QueryAuthenticationStatusParams) (rsp *models.AuthEvent, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.AuthEvent)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AuthEvent: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To create an operator-specific data resource of a UE
// Description:
// Path: /subscription-data/:ueId/operator-specific-data
// Path Params: ueId
type CreateOperSpecDataParams struct {
	UeId              string
	SupportedFeatures string
}

func CreateOperSpecData(cli sbi.ConsumerClient, params CreateOperSpecDataParams, body *map[string]models.OperatorSpecificDataContainer) (rsp *map[string]models.OperatorSpecificDataContainer, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/operator-specific-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(map[string]models.OperatorSpecificDataContainer)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]OperatorSpecificDataContainer: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieve HSS Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions
// Path Params: ueId, subsId
type GetHssSubscriptionInfoParams struct {
	UeId   string
	SubsId string
}

func GetHssSubscriptionInfo(cli sbi.ConsumerClient, params GetHssSubscriptionInfoParams) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/hss-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create individual EE subscription for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions
// Path Params: ueGroupId
// Response headers: Location
func CreateEeGroupSubscriptions(cli sbi.ConsumerClient, ueGroupId string, body *models.EeSubscription) (headers map[string]string, rsp *models.EeSubscription, err error) {

	if len(ueGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions", PATH_ROOT, ueGroupId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.EeSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EeSubscription: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the UE's Location Information
// Description:
// Path: /subscription-data/:ueId/context-data/location
// Path Params: ueId
type QueryUeLocationParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryUeLocation(cli sbi.ConsumerClient, params QueryUeLocationParams) (rsp *models.LocationInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/location", PATH_ROOT, params.UeId)
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
		rsp = new(models.LocationInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LocationInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To store the individual Authentication Status data of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName
// Path Params: ueId, servingNetworkName
type CreateIndividualAuthenticationStatusParams struct {
	UeId               string
	ServingNetworkName string
}

func CreateIndividualAuthenticationStatus(cli sbi.ConsumerClient, params CreateIndividualAuthenticationStatusParams, body *models.AuthEvent) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingNetworkName) == 0 {
		err = fmt.Errorf("servingNetworkName is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status/%s", PATH_ROOT, params.UeId, params.ServingNetworkName)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete SMF Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueId, subsId
type RemoveSmfSubscriptionsInfoParams struct {
	SubsId string
	UeId   string
}

func RemoveSmfSubscriptionsInfo(cli sbi.ConsumerClient, params RemoveSmfSubscriptionsInfoParams) (err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the sdm subscriptions of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions
// Path Params: ueId
type QuerysdmsubscriptionsParams struct {
	UeId              string
	SupportedFeatures string
}

func Querysdmsubscriptions(cli sbi.ConsumerClient, params QuerysdmsubscriptionsParams) (rsp *[]models.SdmSubscription, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions", PATH_ROOT, params.UeId)
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
		rsp = new([]models.SdmSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []SdmSubscription: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify HSS SDM Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions
// Path Params: ueId, subsId
type ModifyHssSDMSubscriptionInfoParams struct {
	UeId              string
	SubsId            string
	SupportedFeatures string
}

func ModifyHssSDMSubscriptionInfo(cli sbi.ConsumerClient, params ModifyHssSDMSubscriptionInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/sdm-subscriptions/%s/hss-sdm-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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
// Path: /subscription-data/shared-data
// Path Params:
type GetSharedDataParams struct {
	SharedDataIds     []string
	SupportedFeatures string
}

func GetSharedData(cli sbi.ConsumerClient, params GetSharedDataParams) (rsp *[]models.SharedData, err error) {

	if len(params.SharedDataIds) == 0 {
		err = fmt.Errorf("shared-data-ids is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/shared-data", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.AddParam("shared-data-ids", models.ArrayOfStringToString(params.SharedDataIds))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
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

// Summary: Deletes the 5GVnGroup
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/:externalGroupId
// Path Params: externalGroupId
func Delete5GVnGroup(cli sbi.ConsumerClient, externalGroupId string) (err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/%s", PATH_ROOT, externalGroupId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve Service Specific Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType
// Path Params: ueId, serviceType
type GetServiceSpecificAuthorizationInfoParams struct {
	ServiceType string
	UeId        string
}

func GetServiceSpecificAuthorizationInfo(cli sbi.ConsumerClient, params GetServiceSpecificAuthorizationInfoParams) (rsp *models.ServiceSpecificAuthorizationInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServiceType) == 0 {
		err = fmt.Errorf("serviceType is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/service-specific-authorizations/%s", PATH_ROOT, params.UeId, params.ServiceType)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ServiceSpecificAuthorizationInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ServiceSpecificAuthorizationInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create the IP-SM-GW context data of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ip-sm-gw
// Path Params: ueId
func CreateIpSmGwContext(cli sbi.ConsumerClient, ueId string, body *models.IpSmGwRegistration) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes a eeSubscription
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId
// Path Params: ueId, subsId
type RemoveeeSubscriptionsParams struct {
	UeId   string
	SubsId string
}

func RemoveeeSubscriptions(cli sbi.ConsumerClient, params RemoveeeSubscriptionsParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve ServiceSpecific Authorization Data
// Description:
// Path: /subscription-data/:ueId/service-specific-authorization-data/:serviceType
// Path Params: ueId, serviceType
// Response headers: Cache-Control, ETag, Last-Modified
type GetSSAuDataParams struct {
	IfNoneMatch            string
	IfModifiedSince        string
	UeId                   string
	ServiceType            string
	SingleNssai            *models.VarSnssai
	Dnn                    string
	MtcProviderInformation string
	AfId                   string
}

func GetSSAuData(cli sbi.ConsumerClient, params GetSSAuDataParams) (headers map[string]string, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServiceType) == 0 {
		err = fmt.Errorf("serviceType is required")
		return
	}
	if params.SingleNssai == nil {
		err = fmt.Errorf("single-nssai is required")
		return
	}
	if len(params.Dnn) == 0 {
		err = fmt.Errorf("dnn is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/service-specific-authorization-data/%s", PATH_ROOT, params.UeId, params.ServiceType)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.MtcProviderInformation) > 0 {
		request.AddParam("mtc-provider-information", params.MtcProviderInformation)
	}
	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.AddParam("single-nssai", models.VarSnssaiToString(*params.SingleNssai))
	request.AddParam("dnn", params.Dnn)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		return
	case 403, 404:
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

// Summary: Retrieves the parameter provision profile data for 5G MBS Group
// Description:
// Path: /subscription-data/group-data/mbs-group-membership/pp-profile-data
// Path Params:
type Query5GMbsGroupPPDataParams struct {
	ExtGroupIds       []string
	SupportedFeatures string
}

func Query5GMbsGroupPPData(cli sbi.ConsumerClient, params Query5GMbsGroupPPDataParams) (rsp *models.Pp5gMbsGroupProfileData, err error) {

	path := fmt.Sprintf("%s/subscription-data/group-data/mbs-group-membership/pp-profile-data", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.ExtGroupIds) > 0 {
		request.AddParam("ext-group-ids", models.ArrayOfStringToString(params.ExtGroupIds))
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
		rsp = new(models.Pp5gMbsGroupProfileData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Pp5gMbsGroupProfileData: %+v", err)
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

// Summary: To store the UPU acknowledgement information of a UE
// Description:
// Path: /subscription-data/:ueId/ue-update-confirmation-data/upu-data
// Path Params: ueId
type CreateAuthenticationUPUParams struct {
	UeId              string
	SupportedFeatures string
}

func CreateAuthenticationUPU(cli sbi.ConsumerClient, params CreateAuthenticationUPUParams, body *models.UpuData) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/ue-update-confirmation-data/upu-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve NIDD Authorization Data GPSI or External Group identifier
// Description:
// Path: /subscription-data/:ueId/nidd-authorization-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type GetNiddAuDataParams struct {
	IfModifiedSince        string
	UeId                   string
	SingleNssai            *models.VarSnssai
	Dnn                    string
	MtcProviderInformation string
	AfId                   string
	IfNoneMatch            string
}

func GetNiddAuData(cli sbi.ConsumerClient, params GetNiddAuDataParams) (headers map[string]string, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if params.SingleNssai == nil {
		err = fmt.Errorf("single-nssai is required")
		return
	}
	if len(params.Dnn) == 0 {
		err = fmt.Errorf("dnn is required")
		return
	}
	if len(params.MtcProviderInformation) == 0 {
		err = fmt.Errorf("mtc-provider-information is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/nidd-authorization-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("mtc-provider-information", params.MtcProviderInformation)
	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.AddParam("single-nssai", models.VarSnssaiToString(*params.SingleNssai))
	request.AddParam("dnn", params.Dnn)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		return
	case 403, 404:
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

// Summary: Retrieves the LCS Broadcast Assistance subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/lcs-bca-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryLcsBcaDataParams struct {
	UeId              string
	ServingPlmnId     string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryLcsBcaData(cli sbi.ConsumerClient, params QueryLcsBcaDataParams) (headers map[string]string, rsp *models.LcsBroadcastAssistanceTypesData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/lcs-bca-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
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
		rsp = new(models.LcsBroadcastAssistanceTypesData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsBroadcastAssistanceTypesData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the 5mbs subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/5mbs-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type Query5mbsDataParams struct {
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func Query5mbsData(cli sbi.ConsumerClient, params Query5mbsDataParams) (headers map[string]string, rsp *models.MbsSubscriptionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/5mbs-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.MbsSubscriptionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode MbsSubscriptionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove the Individual Authentication Status of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName
// Path Params: ueId, servingNetworkName
type DeleteIndividualAuthenticationStatusParams struct {
	UeId               string
	ServingNetworkName string
}

func DeleteIndividualAuthenticationStatus(cli sbi.ConsumerClient, params DeleteIndividualAuthenticationStatusParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.ServingNetworkName) == 0 {
		err = fmt.Errorf("servingNetworkName is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-status/%s", PATH_ROOT, params.UeId, params.ServingNetworkName)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the 5GVnGroup
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/:externalGroupId
// Path Params: externalGroupId
type Modify5GVnGroupParams struct {
	ExternalGroupId   string
	SupportedFeatures string
}

func Modify5GVnGroup(cli sbi.ConsumerClient, params Modify5GVnGroupParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.ExternalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/%s", PATH_ROOT, params.ExternalGroupId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieves the PEI Information of the 5GC/EPC domains
// Description:
// Path: /subscription-data/:ueId/context-data/pei-info
// Path Params: ueId
func QueryPeiInformation(cli sbi.ConsumerClient, ueId string) (rsp *models.PeiUpdateInfo, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/pei-info", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PeiUpdateInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PeiUpdateInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update the PEI Information of the 5GC/EPC domains
// Description:
// Path: /subscription-data/:ueId/context-data/pei-info
// Path Params: ueId
func CreateOrUpdatePeiInformation(cli sbi.ConsumerClient, ueId string, body *models.PeiUpdateInfo) (rsp *models.PeiUpdateInfo, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/pei-info", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.PeiUpdateInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PeiUpdateInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve SMF Subscription Info for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueGroupId, subsId
type GetSmfGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func GetSmfGroupSubscriptions(cli sbi.ConsumerClient, params GetSmfGroupSubscriptionsParams) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To modify the SMF context data of a UE in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type UpdateSmfContextParams struct {
	UeId              string
	PduSessionId      int
	SupportedFeatures string
}

func UpdateSmfContext(cli sbi.ConsumerClient, params UpdateSmfContextParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieves the subscribed ProSe service Data of a UE
// Description:
// Path: /subscription-data/:ueId/prose-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryPorseDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	SupportedFeatures string
}

func QueryPorseData(cli sbi.ConsumerClient, params QueryPorseDataParams) (headers map[string]string, rsp *models.ProseSubscriptionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/prose-data", PATH_ROOT, params.UeId)
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
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the authentication subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-subscription
// Path Params: ueId
type QueryAuthSubsDataParams struct {
	UeId              string
	SupportedFeatures string
}

func QueryAuthSubsData(cli sbi.ConsumerClient, params QueryAuthSubsDataParams) (rsp *models.AuthenticationSubscription, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-subscription", PATH_ROOT, params.UeId)
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
		rsp = new(models.AuthenticationSubscription)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AuthenticationSubscription: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To remove an individual SMF context data of a UE the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type DeleteSmfRegistrationParams struct {
	UeId         string
	PduSessionId int
}

func DeleteSmfRegistration(cli sbi.ConsumerClient, params DeleteSmfRegistrationParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify the IP-SM-GW context data of a UE
// Description:
// Path: /subscription-data/:ueId/context-data/ip-sm-gw
// Path Params: ueId
func ModifyIpSmGwContext(cli sbi.ConsumerClient, ueId string, body *[]models.PatchItem) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 403, 404:
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

// Summary: Modify the Message Waiting Data of the UE
// Description:
// Path: /subscription-data/:ueId/context-data/mwd
// Path Params: ueId
func ModifyMessageWaitingData(cli sbi.ConsumerClient, ueId string, body *[]models.PatchItem) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/mwd", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 403, 404:
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

// Summary: create a Provisioning Parameter Data Entry
// Description:
// Path: /subscription-data/:ueId/pp-data-store/:afInstanceId
// Path Params: ueId, afInstanceId
type CreatePPDataEntryParams struct {
	UeId         string
	AfInstanceId string
}

func CreatePPDataEntry(cli sbi.ConsumerClient, params CreatePPDataEntryParams, body *models.PpDataEntry) (rsp *models.PpDataEntry, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.AfInstanceId) == 0 {
		err = fmt.Errorf("afInstanceId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data-store/%s", PATH_ROOT, params.UeId, params.AfInstanceId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.PpDataEntry)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PpDataEntry: %+v", err)
		}
	case 204:
		return
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

// Summary: Retrieve SMF Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueId, subsId
type GetSmfSubscriptionInfoParams struct {
	UeId   string
	SubsId string
}

func GetSmfSubscriptionInfo(cli sbi.ConsumerClient, params GetSmfSubscriptionInfoParams) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create SMF Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueId, subsId
type CreateSMFSubscriptionsParams struct {
	UeId   string
	SubsId string
}

func CreateSMFSubscriptions(cli sbi.ConsumerClient, params CreateSMFSubscriptionsParams, body *models.SmfSubscriptionInfo) (rsp *models.SmfSubscriptionInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmfSubscriptionInfo: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the list of subscriptions
// Description:
// Path: /subscription-data/subs-to-notify
// Path Params:
type QuerySubsToNotifyParams struct {
	UeId              string
	SupportedFeatures string
}

func QuerySubsToNotify(cli sbi.ConsumerClient, params QuerySubsToNotifyParams) (rsp *[]models.SubscriptionDataSubscriptions, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ue-id is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	request.AddParam("ue-id", params.UeId)
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
		rsp = new([]models.SubscriptionDataSubscriptions)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []SubscriptionDataSubscriptions: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the Session Management subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/:servingPlmnId/provisioned-data/sm-data
// Path Params: ueId, servingPlmnId
// Response headers: Cache-Control, ETag, Last-Modified
type QuerySmDataParams struct {
	SingleNssai       *models.VarSnssai
	Dnn               string
	Fields            []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	ServingPlmnId     string
}

func QuerySmData(cli sbi.ConsumerClient, params QuerySmDataParams) (headers map[string]string, rsp *models.SmSubsData, err error) {

	if len(params.ServingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/%s/provisioned-data/sm-data", PATH_ROOT, params.UeId, params.ServingPlmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.SingleNssai != nil {
		request.AddParam("single-nssai", models.VarSnssaiToString(*params.SingleNssai))
	}
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
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
		rsp = new(models.SmSubsData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmSubsData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the provisioned parameter data
// Description:
// Path: /subscription-data/:ueId/pp-data
// Path Params: ueId
type ModifyPpDataParams struct {
	SupportedFeatures string
	UeId              string
}

func ModifyPpData(cli sbi.ConsumerClient, params ModifyPpDataParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/pp-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieves the ee profile data profile data of a group or anyUE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-profile-data
// Path Params: ueGroupId
type QueryGroupEEDataParams struct {
	UeGroupId         string
	SupportedFeatures string
}

func QueryGroupEEData(cli sbi.ConsumerClient, params QueryGroupEEDataParams) (rsp *models.EeGroupProfileData, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-profile-data", PATH_ROOT, params.UeGroupId)
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
		rsp = new(models.EeGroupProfileData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EeGroupProfileData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the parameter provision profile data for 5G VN Group
// Description:
// Path: /subscription-data/group-data/5g-vn-groups/pp-profile-data
// Path Params:
type Query5GVNGroupPPDataParams struct {
	ExtGroupIds       []string
	SupportedFeatures string
}

func Query5GVNGroupPPData(cli sbi.ConsumerClient, params Query5GVNGroupPPDataParams) (rsp *models.Pp5gVnGroupProfileData, err error) {

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/pp-profile-data", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.ExtGroupIds) > 0 {
		request.AddParam("ext-group-ids", models.ArrayOfStringToString(params.ExtGroupIds))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.Pp5gVnGroupProfileData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Pp5gVnGroupProfileData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To store the AMF context data of a UE using 3gpp access in the UDR
// Description:
// Path: /subscription-data/:ueId/context-data/amf-3gpp-access
// Path Params: ueId
// Response headers: Location
func CreateAmfContext3gpp(cli sbi.ConsumerClient, ueId string, body *models.Amf3GppAccessRegistration) (headers map[string]string, rsp *models.Amf3GppAccessRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/amf-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.Amf3GppAccessRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode Amf3GppAccessRegistration: %+v", err)
		}
	case 204:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve individual shared data
// Description:
// Path: /subscription-data/shared-data/:sharedDataId
// Path Params: sharedDataId
// Response headers: Cache-Control, ETag, Last-Modified
type GetIndividualSharedDataParams struct {
	SharedDataId    string
	IfNoneMatch     string
	IfModifiedSince string
}

func GetIndividualSharedData(cli sbi.ConsumerClient, params GetIndividualSharedDataParams) (headers map[string]string, rsp *models.SharedData, err error) {

	if len(params.SharedDataId) == 0 {
		err = fmt.Errorf("sharedDataId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/shared-data/%s", PATH_ROOT, params.SharedDataId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
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

// Summary: Modify an individual subscriptionDataSubscription
// Description:
// Path: /subscription-data/subs-to-notify/:subsId
// Path Params: subsId
type ModifysubscriptionDataSubscriptionParams struct {
	SubsId            string
	SupportedFeatures string
}

func ModifysubscriptionDataSubscription(cli sbi.ConsumerClient, params ModifysubscriptionDataSubscriptionParams, body *[]models.PatchItem) (rsp *models.Schema, err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/subs-to-notify/%s", PATH_ROOT, params.SubsId)
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
	case 204:
		return
	case 403, 404:
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

// Summary: Retrieves the LCS Mobile Originated subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/lcs-mo-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryLcsMoDataParams struct {
	UeId              string
	Fields            []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryLcsMoData(cli sbi.ConsumerClient, params QueryLcsMoDataParams) (headers map[string]string, rsp *models.LcsMoData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/lcs-mo-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
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
		rsp = new(models.LcsMoData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode LcsMoData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the SMSF context data of a UE using 3gpp access
// Description:
// Path: /subscription-data/:ueId/context-data/smsf-3gpp-access
// Path Params: ueId
type QuerySmsfContext3gppParams struct {
	SupportedFeatures string
	UeId              string
	Fields            []string
}

func QuerySmsfContext3gpp(cli sbi.ConsumerClient, params QuerySmsfContext3gppParams) (rsp *models.SmsfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/smsf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmsfRegistration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify SMF Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions
// Path Params: ueId, subsId
type ModifySmfSubscriptionInfoParams struct {
	SupportedFeatures string
	UeId              string
	SubsId            string
}

func ModifySmfSubscriptionInfo(cli sbi.ConsumerClient, params ModifySmfSubscriptionInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/smf-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Modify HSS Subscription Info
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions
// Path Params: ueId, subsId
type ModifyHssSubscriptionInfoParams struct {
	UeId              string
	SubsId            string
	SupportedFeatures string
}

func ModifyHssSubscriptionInfo(cli sbi.ConsumerClient, params ModifyHssSubscriptionInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s/hss-subscriptions", PATH_ROOT, params.UeId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieves the data of a 5G VN Group
// Description:
// Path: /subscription-data/group-data/5g-vn-groups
// Path Params:
func Query5GVnGroup(cli sbi.ConsumerClient, gpsis []string) (rsp *map[string]models.FiveGVnGroupConfiguration, err error) {

	path := fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(gpsis) > 0 {
		request.AddParam("gpsis", models.ArrayOfStringToString(gpsis))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(map[string]models.FiveGVnGroupConfiguration)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode map[string]FiveGVnGroupConfiguration: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the subscribed enhanced Coverage Restriction Data of a UE
// Description:
// Path: /subscription-data/:ueId/coverage-restriction-data
// Path Params: ueId
// Response headers: Cache-Control, ETag, Last-Modified
type QueryCoverageRestrictionDataParams struct {
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func QueryCoverageRestrictionData(cli sbi.ConsumerClient, params QueryCoverageRestrictionDataParams) (headers map[string]string, rsp *models.EnhancedCoverageRestrictionData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/coverage-restriction-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.EnhancedCoverageRestrictionData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EnhancedCoverageRestrictionData: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve AMF subscription Info for a group of UEs or any UE
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueGroupId, subsId
type GetAmfGroupSubscriptionsParams struct {
	UeGroupId string
	SubsId    string
}

func GetAmfGroupSubscriptions(cli sbi.ConsumerClient, params GetAmfGroupSubscriptionsParams) (rsp *[]models.AmfSubscriptionInfo, err error) {

	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new([]models.AmfSubscriptionInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode []AmfSubscriptionInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: To modify operator specific data of a UE
// Description:
// Path: /subscription-data/:ueId/operator-specific-data
// Path Params: ueId
type ModifyOperSpecDataParams struct {
	UeId              string
	SupportedFeatures string
}

func ModifyOperSpecData(cli sbi.ConsumerClient, params ModifyOperSpecDataParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/operator-specific-data", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieve NIDD Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/nidd-authorizations
// Path Params: ueId
func GetNiddAuthorizationInfo(cli sbi.ConsumerClient, ueId string) (rsp *models.NiddAuthorizationInfo, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/nidd-authorizations", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.NiddAuthorizationInfo)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode NiddAuthorizationInfo: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify NIDD Authorization Info
// Description:
// Path: /subscription-data/:ueId/context-data/nidd-authorizations
// Path Params: ueId
type ModifyNiddAuthorizationInfoParams struct {
	UeId              string
	SupportedFeatures string
}

func ModifyNiddAuthorizationInfo(cli sbi.ConsumerClient, params ModifyNiddAuthorizationInfoParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/nidd-authorizations", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieve multiple subscribed data sets of a UE
// Description:
// Path: /subscription-data/:ueId
// Path Params: ueId
type QueryUeSubscribedDataParams struct {
	UeId         string
	DatasetNames []string
	ServingPlmn  string
}

func QueryUeSubscribedData(cli sbi.ConsumerClient, params QueryUeSubscribedDataParams) (rsp *models.UeSubscribedDataSets, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.DatasetNames) > 0 {
		request.AddParam("dataset-names", models.ArrayOfStringToString(params.DatasetNames))
	}
	if len(params.ServingPlmn) > 0 {
		request.AddParam("serving-plmn", params.ServingPlmn)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UeSubscribedDataSets)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeSubscribedDataSets: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: modify the AMF Subscription Info
// Description:
// Path: /subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions
// Path Params: ueGroupId, subsId
type ModifyAmfGroupSubscriptionsParams struct {
	UeGroupId         string
	SubsId            string
	SupportedFeatures string
}

func ModifyAmfGroupSubscriptions(cli sbi.ConsumerClient, params ModifyAmfGroupSubscriptionsParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if len(params.UeGroupId) == 0 {
		err = fmt.Errorf("ueGroupId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/group-data/%s/ee-subscriptions/%s/amf-subscriptions", PATH_ROOT, params.UeGroupId, params.SubsId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: modify the authentication subscription data of a UE
// Description:
// Path: /subscription-data/:ueId/authentication-data/authentication-subscription
// Path Params: ueId
type ModifyAuthenticationSubscriptionParams struct {
	SupportedFeatures string
	UeId              string
}

func ModifyAuthenticationSubscription(cli sbi.ConsumerClient, params ModifyAuthenticationSubscriptionParams, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/authentication-data/authentication-subscription", PATH_ROOT, params.UeId)
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
		rsp = new(models.PatchResult)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PatchResult: %+v", err)
		}
	case 204:
		return
	case 403:
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

// Summary: Retrieve a eeSubscription
// Description:
// Path: /subscription-data/:ueId/context-data/ee-subscriptions/:subsId
// Path Params: ueId, subsId
type QueryeeSubscriptionParams struct {
	UeId   string
	SubsId string
}

func QueryeeSubscription(cli sbi.ConsumerClient, params QueryeeSubscriptionParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.SubsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/subscription-data/%s/context-data/ee-subscriptions/%s", PATH_ROOT, params.UeId, params.SubsId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		return
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
