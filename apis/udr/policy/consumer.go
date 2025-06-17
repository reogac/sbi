/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package policy

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = ""
)

// Summary: Retrieves the session management policy data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/sm-data
// Path Params: ueId
type ReadSessionManagementPolicyDataParams struct {
	UeId     string
	Snssai   *models.Snssai
	Dnn      string
	Fields   []string
	SuppFeat string
}

func ReadSessionManagementPolicyData(cli sbi.ConsumerClient, params ReadSessionManagementPolicyDataParams) (rsp *models.SmPolicyData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/sm-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	if params.Snssai != nil {
		request.AddParam("snssai", models.SnssaiToString(*params.Snssai))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmPolicyData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 414, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the sponsored connectivity information for a given sponsorId
// Description:
// Path: /policy-data/sponsor-connectivity-data/:sponsorId
// Path Params: sponsorId
func ReadSponsorConnectivityData(cli sbi.ConsumerClient, sponsorId string) (rsp *models.SponsorConnectivityData, err error) {

	if len(sponsorId) == 0 {
		err = fmt.Errorf("sponsorId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/sponsor-connectivity-data/%s", PATH_ROOT, sponsorId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SponsorConnectivityData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create or modify the operator specific policy data of a UE
// Description:
// Path: /policy-data/ues/:ueId/operator-specific-data
// Path Params: ueId
func ReplaceOperatorSpecificData(cli sbi.ConsumerClient, ueId string, body *map[string]models.OperatorSpecificDataContainer) (rsp *map[string]models.OperatorSpecificDataContainer, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/operator-specific-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(map[string]models.OperatorSpecificDataContainer)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify the operator specific policy data of a UE
// Description:
// Path: /policy-data/ues/:ueId/operator-specific-data
// Path Params: ueId
func UpdateOperatorSpecificData(cli sbi.ConsumerClient, ueId string, body *[]models.PatchItem) (rsp *models.PatchResult, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/operator-specific-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PatchResult)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves a network Slice specific policy control data resource
// Description:
// Path: /policy-data/slice-control-data/:snssai
// Path Params: snssai
type ReadSlicePolicyControlDataParams struct {
	SuppFeat string
	Snssai   *models.Snssai
}

func ReadSlicePolicyControlData(cli sbi.ConsumerClient, params ReadSlicePolicyControlDataParams) (rsp *models.SlicePolicyData, err error) {

	if params.Snssai == nil {
		err = fmt.Errorf("snssai is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/slice-control-data/%s", PATH_ROOT, models.SnssaiToString(*params.Snssai))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SlicePolicyData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the BDT data information associated with a BDT reference Id
// Description:
// Path: /policy-data/bdt-data/:bdtReferenceId
// Path Params: bdtReferenceId
type ReadIndividualBdtDataParams struct {
	BdtReferenceId string
	SuppFeat       string
}

func ReadIndividualBdtData(cli sbi.ConsumerClient, params ReadIndividualBdtDataParams) (rsp *models.BdtData, err error) {

	if len(params.BdtReferenceId) == 0 {
		err = fmt.Errorf("bdtReferenceId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/bdt-data/%s", PATH_ROOT, params.BdtReferenceId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.BdtData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create a subscription to receive notification of policy data changes
// Description:
// Path: /policy-data/subs-to-notify
// Path Params:
func CreateIndividualPolicyDataSubscription(cli sbi.ConsumerClient, body *models.PolicyDataSubscription) (rsp *models.PolicyDataSubscription, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/subs-to-notify", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.PolicyDataSubscription)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve the UE policy set data for an H-PLMN
// Description:
// Path: /policy-data/plmns/:plmnId/ue-policy-set
// Path Params: plmnId
func ReadPlmnUePolicySet(cli sbi.ConsumerClient, plmnId string) (rsp *models.UePolicySet, err error) {

	if len(plmnId) == 0 {
		err = fmt.Errorf("plmnId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/plmns/%s/ue-policy-set", PATH_ROOT, plmnId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UePolicySet)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 412, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve the policy data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId
// Path Params: ueId
type ReadPolicyDataParams struct {
	UeId            string
	SuppFeat        string
	DataSubsetNames []string
}

func ReadPolicyData(cli sbi.ConsumerClient, params ReadPolicyDataParams) (rsp *models.PolicyDataForIndividualUe, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	if len(params.DataSubsetNames) > 0 {
		request.AddParam("data-subset-names", models.ArrayOfStringToString(params.DataSubsetNames))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PolicyDataForIndividualUe)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete a usage monitoring resource
// Description:
// Path: /policy-data/ues/:ueId/sm-data/:usageMonId
// Path Params: ueId, usageMonId
type DeleteUsageMonitoringInformationParams struct {
	UeId       string
	UsageMonId string
}

func DeleteUsageMonitoringInformation(cli sbi.ConsumerClient, params DeleteUsageMonitoringInformationParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.UsageMonId) == 0 {
		err = fmt.Errorf("usageMonId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/sm-data/%s", PATH_ROOT, params.UeId, params.UsageMonId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes an BDT data resource associated with an BDT reference Id
// Description:
// Path: /policy-data/bdt-data/:bdtReferenceId
// Path Params: bdtReferenceId
func DeleteIndividualBdtData(cli sbi.ConsumerClient, bdtReferenceId string) (err error) {

	if len(bdtReferenceId) == 0 {
		err = fmt.Errorf("bdtReferenceId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/bdt-data/%s", PATH_ROOT, bdtReferenceId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify the UE policy set data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/ue-policy-set
// Path Params: ueId
func UpdateUEPolicySet(cli sbi.ConsumerClient, ueId string, body *models.UePolicySetPatch) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/ue-policy-set", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify the session management policy data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/sm-data
// Path Params: ueId
func UpdateSessionManagementPolicyData(cli sbi.ConsumerClient, ueId string, body *models.SmPolicyDataPatch) (rsp *models.SmPolicyData, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/sm-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmPolicyData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: When the feature OSDResource_Create_Delete is supported, delete OperatorSpecificData resource
// Description:
// Path: /policy-data/ues/:ueId/operator-specific-data
// Path Params: ueId
func DeleteOperatorSpecificData(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/operator-specific-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the access and mobility policy data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/am-data
// Path Params: ueId
func ReadAccessAndMobilityPolicyData(cli sbi.ConsumerClient, ueId string) (rsp *models.AmPolicyData, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/am-data", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.AmPolicyData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the UE policy set data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/ue-policy-set
// Path Params: ueId
type ReadUEPolicySetParams struct {
	UeId     string
	SuppFeat string
}

func ReadUEPolicySet(cli sbi.ConsumerClient, params ReadUEPolicySetParams) (rsp *models.UePolicySet, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/ue-policy-set", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UePolicySet)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve a usage monitoring resource
// Description:
// Path: /policy-data/ues/:ueId/sm-data/:usageMonId
// Path Params: ueId, usageMonId
type ReadUsageMonitoringInformationParams struct {
	UeId       string
	UsageMonId string
	SuppFeat   string
}

func ReadUsageMonitoringInformation(cli sbi.ConsumerClient, params ReadUsageMonitoringInformationParams) (rsp *models.UsageMonData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.UsageMonId) == 0 {
		err = fmt.Errorf("usageMonId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/sm-data/%s", PATH_ROOT, params.UeId, params.UsageMonId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UsageMonData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 414, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieves the BDT data collection
// Description:
// Path: /policy-data/bdt-data
// Path Params:
type ReadBdtDataParams struct {
	BdtRefIds []string
	SuppFeat  string
}

func ReadBdtData(cli sbi.ConsumerClient, params ReadBdtDataParams) (rsp *[]models.BdtData, err error) {

	path := fmt.Sprintf("%s/policy-data/bdt-data", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.BdtRefIds) > 0 {
		request.AddParam("bdt-ref-ids", models.ArrayOfStringToString(params.BdtRefIds))
	}
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new([]models.BdtData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modifies an BDT data resource associated with an BDT reference Id
// Description:
// Path: /policy-data/bdt-data/:bdtReferenceId
// Path Params: bdtReferenceId
func UpdateIndividualBdtData(cli sbi.ConsumerClient, bdtReferenceId string, body *models.BdtDataPatch) (rsp *models.BdtData, err error) {

	if len(bdtReferenceId) == 0 {
		err = fmt.Errorf("bdtReferenceId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/bdt-data/%s", PATH_ROOT, bdtReferenceId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.BdtData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify a network Slice specific policy control data resource
// Description:
// Path: /policy-data/slice-control-data/:snssai
// Path Params: snssai
func UpdateSlicePolicyControlData(cli sbi.ConsumerClient, snssai *models.Snssai, body *models.SlicePolicyDataPatch) (rsp *models.SlicePolicyData, err error) {

	if snssai == nil {
		err = fmt.Errorf("snssai is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/slice-control-data/%s", PATH_ROOT, models.SnssaiToString(*snssai))
	request := sbi.NewRequest(path, http.MethodPatch, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SlicePolicyData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve MBS Session Policy Control Data for an MBS Session.
// Description:
// Path: /policy-data/mbs-session-pol-data/:polSessionId
// Path Params: polSessionId
func GetMBSSessPolCtrlData(cli sbi.ConsumerClient, polSessionId *models.MbsSessPolDataId) (rsp *models.MbsSessPolCtrlData, err error) {

	if polSessionId == nil {
		err = fmt.Errorf("polSessionId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/mbs-session-pol-data/%s", PATH_ROOT, models.MbsSessPolDataIdToString(*polSessionId))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.MbsSessPolCtrlData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 414, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create or modify the UE policy set data for a subscriber
// Description:
// Path: /policy-data/ues/:ueId/ue-policy-set
// Path Params: ueId
func CreateOrReplaceUEPolicySet(cli sbi.ConsumerClient, ueId string, body *models.UePolicySet) (rsp *models.UePolicySet, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/ue-policy-set", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UePolicySet)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create a usage monitoring resource
// Description:
// Path: /policy-data/ues/:ueId/sm-data/:usageMonId
// Path Params: ueId, usageMonId
type CreateUsageMonitoringResourceParams struct {
	UeId       string
	UsageMonId string
}

func CreateUsageMonitoringResource(cli sbi.ConsumerClient, params CreateUsageMonitoringResourceParams, body *models.UsageMonData) (rsp *models.UsageMonData, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.UsageMonId) == 0 {
		err = fmt.Errorf("usageMonId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/sm-data/%s", PATH_ROOT, params.UeId, params.UsageMonId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.UsageMonData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 411, 413, 414, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Creates an BDT data resource associated with an BDT reference Id
// Description:
// Path: /policy-data/bdt-data/:bdtReferenceId
// Path Params: bdtReferenceId
func CreateIndividualBdtData(cli sbi.ConsumerClient, bdtReferenceId string, body *models.BdtData) (rsp *models.BdtData, err error) {

	if len(bdtReferenceId) == 0 {
		err = fmt.Errorf("bdtReferenceId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/bdt-data/%s", PATH_ROOT, bdtReferenceId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.BdtData)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 411, 413, 414, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete the individual Policy Data subscription
// Description:
// Path: /policy-data/subs-to-notify/:subsId
// Path Params: subsId
func DeleteIndividualPolicyDataSubscription(cli sbi.ConsumerClient, subsId string) (err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/subs-to-notify/%s", PATH_ROOT, subsId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve the operator specific policy data of an UE
// Description:
// Path: /policy-data/ues/:ueId/operator-specific-data
// Path Params: ueId
type ReadOperatorSpecificDataParams struct {
	UeId     string
	Fields   []string
	SuppFeat string
}

func ReadOperatorSpecificData(cli sbi.ConsumerClient, params ReadOperatorSpecificDataParams) (rsp *map[string]models.OperatorSpecificDataContainer, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/ues/%s/operator-specific-data", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.Fields) > 0 {
		request.AddParam("fields", models.ArrayOfStringToString(params.Fields))
	}
	if len(params.SuppFeat) > 0 {
		request.AddParam("supp-feat", params.SuppFeat)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(map[string]models.OperatorSpecificDataContainer)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 414, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Modify a subscription to receive notification of policy data changes
// Description:
// Path: /policy-data/subs-to-notify/:subsId
// Path Params: subsId
func ReplaceIndividualPolicyDataSubscription(cli sbi.ConsumerClient, subsId string, body *models.PolicyDataSubscription) (rsp *models.PolicyDataSubscription, err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policy-data/subs-to-notify/%s", PATH_ROOT, subsId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PolicyDataSubscription)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
