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
)

func OnReplaceOperatorSpecificData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReplaceOperatorSpecificData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateOperatorSpecificData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateOperatorSpecificData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadSlicePolicyControlData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadSlicePolicyControlDataParams

	// read 'snssai'
	snssaiStr := ctx.Param("snssai")
	if len(snssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "snssai is required"), nil)
		return
	}

	if params.Snssai, err = models.SnssaiFromString(snssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse snssai failed: %+v", err)), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadSlicePolicyControlData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReadSessionManagementPolicyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadSessionManagementPolicyDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'snssai'
	snssaiStr := ctx.Param("snssai")
	if len(snssaiStr) > 0 {
		if params.Snssai, err = models.SnssaiFromString(snssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse snssai failed: %+v", err)), nil)
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadSessionManagementPolicyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReadSponsorConnectivityData(ctx sbi.RequestContext, prod Producer) {

	// read 'sponsorId'
	var sponsorId string
	sponsorId = ctx.Param("sponsorId")
	if len(sponsorId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sponsorId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadSponsorConnectivityData(sponsorId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadPlmnUePolicySet(ctx sbi.RequestContext, prod Producer) {

	// read 'plmnId'
	var plmnId string
	plmnId = ctx.Param("plmnId")
	if len(plmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "plmnId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadPlmnUePolicySet(plmnId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReadIndividualBdtData(ctx sbi.RequestContext, prod Producer) {
	var params ReadIndividualBdtDataParams

	// read 'bdtReferenceId'
	params.BdtReferenceId = ctx.Param("bdtReferenceId")
	if len(params.BdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadIndividualBdtData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateIndividualPolicyDataSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyDataSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualPolicyDataSubscription(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDeleteIndividualBdtData(ctx sbi.RequestContext, prod Producer) {

	// read 'bdtReferenceId'
	var bdtReferenceId string
	bdtReferenceId = ctx.Param("bdtReferenceId")
	if len(bdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualBdtData(bdtReferenceId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadPolicyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadPolicyDataParams

	// read 'data-subset-names'
	dataSubsetNamesStr := ctx.Param("data-subset-names")
	if len(dataSubsetNamesStr) > 0 {
		if params.DataSubsetNames, err = models.ArrayOfStringFromString(dataSubsetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse data-subset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadPolicyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDeleteUsageMonitoringInformation(ctx sbi.RequestContext, prod Producer) {
	var params DeleteUsageMonitoringInformationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'usageMonId'
	params.UsageMonId = ctx.Param("usageMonId")
	if len(params.UsageMonId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "usageMonId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteUsageMonitoringInformation(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteOperatorSpecificData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteOperatorSpecificData(ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateUEPolicySet(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UePolicySetPatch)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateUEPolicySet(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateSessionManagementPolicyData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmPolicyDataPatch)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSessionManagementPolicyData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadUsageMonitoringInformation(ctx sbi.RequestContext, prod Producer) {
	var params ReadUsageMonitoringInformationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'usageMonId'
	params.UsageMonId = ctx.Param("usageMonId")
	if len(params.UsageMonId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "usageMonId is required"), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadUsageMonitoringInformation(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadBdtData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadBdtDataParams

	// read 'bdt-ref-ids'
	bdtRefIdsStr := ctx.Param("bdt-ref-ids")
	if len(bdtRefIdsStr) > 0 {
		if params.BdtRefIds, err = models.ArrayOfStringFromString(bdtRefIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse bdt-ref-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadBdtData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUpdateIndividualBdtData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'bdtReferenceId'
	var bdtReferenceId string
	bdtReferenceId = ctx.Param("bdtReferenceId")
	if len(bdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.BdtDataPatch)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateIndividualBdtData(bdtReferenceId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadAccessAndMobilityPolicyData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadAccessAndMobilityPolicyData(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReadUEPolicySet(ctx sbi.RequestContext, prod Producer) {
	var params ReadUEPolicySetParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadUEPolicySet(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateIndividualBdtData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'bdtReferenceId'
	var bdtReferenceId string
	bdtReferenceId = ctx.Param("bdtReferenceId")
	if len(bdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.BdtData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualBdtData(bdtReferenceId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDeleteIndividualPolicyDataSubscription(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualPolicyDataSubscription(subsId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadOperatorSpecificData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadOperatorSpecificDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// call application handler
	rsp, prob := prod.HandleReadOperatorSpecificData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUpdateSlicePolicyControlData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'snssai'
	var snssai *models.Snssai
	snssaiStr := ctx.Param("snssai")
	if len(snssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "snssai is required"), nil)
		return
	}

	if snssai, err = models.SnssaiFromString(snssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse snssai failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SlicePolicyDataPatch)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSlicePolicyControlData(snssai, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMBSSessPolCtrlData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'polSessionId'
	var polSessionId *models.MbsSessPolDataId
	polSessionIdStr := ctx.Param("polSessionId")
	if len(polSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polSessionId is required"), nil)
		return
	}

	if polSessionId, err = models.MbsSessPolDataIdFromString(polSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse polSessionId failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMBSSessPolCtrlData(polSessionId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateOrReplaceUEPolicySet(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UePolicySet)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateOrReplaceUEPolicySet(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateUsageMonitoringResource(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateUsageMonitoringResourceParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'usageMonId'
	params.UsageMonId = ctx.Param("usageMonId")
	if len(params.UsageMonId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "usageMonId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UsageMonData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateUsageMonitoringResource(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReplaceIndividualPolicyDataSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyDataSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReplaceIndividualPolicyDataSubscription(subsId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleReplaceOperatorSpecificData(string, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleUpdateOperatorSpecificData(string, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleReadSlicePolicyControlData(*ReadSlicePolicyControlDataParams) (*models.SlicePolicyData, *models.ProblemDetails)

	HandleReadSessionManagementPolicyData(*ReadSessionManagementPolicyDataParams) (*models.SmPolicyData, *models.ProblemDetails)

	HandleReadSponsorConnectivityData(string) (*models.SponsorConnectivityData, *models.ProblemDetails)

	HandleReadPlmnUePolicySet(string) (*models.UePolicySet, *models.ProblemDetails)

	HandleReadIndividualBdtData(*ReadIndividualBdtDataParams) (*models.BdtData, *models.ProblemDetails)

	HandleCreateIndividualPolicyDataSubscription(*models.PolicyDataSubscription) (*models.PolicyDataSubscription, *models.ProblemDetails)

	HandleDeleteIndividualBdtData(string) *models.ProblemDetails

	HandleReadPolicyData(*ReadPolicyDataParams) (*models.PolicyDataForIndividualUe, *models.ProblemDetails)

	HandleDeleteUsageMonitoringInformation(*DeleteUsageMonitoringInformationParams) *models.ProblemDetails

	HandleDeleteOperatorSpecificData(string) *models.ProblemDetails

	HandleUpdateUEPolicySet(string, *models.UePolicySetPatch) *models.ProblemDetails

	HandleUpdateSessionManagementPolicyData(string, *models.SmPolicyDataPatch) (*models.SmPolicyData, *models.ProblemDetails)

	HandleReadUsageMonitoringInformation(*ReadUsageMonitoringInformationParams) (*models.UsageMonData, *models.ProblemDetails)

	HandleReadBdtData(*ReadBdtDataParams) (*[]models.BdtData, *models.ProblemDetails)

	HandleUpdateIndividualBdtData(string, *models.BdtDataPatch) (*models.BdtData, *models.ProblemDetails)

	HandleReadAccessAndMobilityPolicyData(string) (*models.AmPolicyData, *models.ProblemDetails)

	HandleReadUEPolicySet(*ReadUEPolicySetParams) (*models.UePolicySet, *models.ProblemDetails)

	HandleCreateIndividualBdtData(string, *models.BdtData) (*models.BdtData, *models.ProblemDetails)

	HandleDeleteIndividualPolicyDataSubscription(string) *models.ProblemDetails

	HandleReadOperatorSpecificData(*ReadOperatorSpecificDataParams) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleUpdateSlicePolicyControlData(*models.Snssai, *models.SlicePolicyDataPatch) (*models.SlicePolicyData, *models.ProblemDetails)

	HandleGetMBSSessPolCtrlData(*models.MbsSessPolDataId) (*models.MbsSessPolCtrlData, *models.ProblemDetails)

	HandleCreateOrReplaceUEPolicySet(string, *models.UePolicySet) (*models.UePolicySet, *models.ProblemDetails)

	HandleCreateUsageMonitoringResource(*CreateUsageMonitoringResourceParams, *models.UsageMonData) (*models.UsageMonData, *models.ProblemDetails)

	HandleReplaceIndividualPolicyDataSubscription(string, *models.PolicyDataSubscription) (*models.PolicyDataSubscription, *models.ProblemDetails)
}
