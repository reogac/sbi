/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package policy

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

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
	prob := prod.HandleUpdateUEPolicySet(ctx.Context(), ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReadUEPolicySet(ctx sbi.RequestContext, prod Producer) {
	var params ReadUEPolicySetParams

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadUEPolicySet(ctx.Context(), &params)

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

	// call application handler
	rsp, prob := prod.HandleReadSessionManagementPolicyData(ctx.Context(), &params)

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

func OnReadUsageMonitoringInformation(ctx sbi.RequestContext, prod Producer) {
	var params ReadUsageMonitoringInformationParams

	// read 'usageMonId'
	params.UsageMonId = ctx.Param("usageMonId")
	if len(params.UsageMonId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "usageMonId is required"), nil)
		return
	}

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadUsageMonitoringInformation(ctx.Context(), &params)

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

func OnReadSponsorConnectivityData(ctx sbi.RequestContext, prod Producer) {

	// read 'sponsorId'
	var sponsorId string
	sponsorId = ctx.Param("sponsorId")
	if len(sponsorId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sponsorId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadSponsorConnectivityData(ctx.Context(), sponsorId)

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
	rsp, prob := prod.HandleReadBdtData(ctx.Context(), &params)

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

func OnDeleteIndividualBdtData(ctx sbi.RequestContext, prod Producer) {

	// read 'bdtReferenceId'
	var bdtReferenceId string
	bdtReferenceId = ctx.Param("bdtReferenceId")
	if len(bdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualBdtData(ctx.Context(), bdtReferenceId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

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
	rsp, prob := prod.HandleUpdateIndividualBdtData(ctx.Context(), bdtReferenceId, body)

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
	headers, rsp, prob := prod.HandleCreateIndividualPolicyDataSubscription(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
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
	prob := prod.HandleDeleteIndividualPolicyDataSubscription(ctx.Context(), subsId)

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
	prob := prod.HandleDeleteOperatorSpecificData(ctx.Context(), ueId)

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
	rsp, prob := prod.HandleReadPlmnUePolicySet(ctx.Context(), plmnId)

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
	rsp, prob := prod.HandleGetMBSSessPolCtrlData(ctx.Context(), polSessionId)

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

func OnReadAccessAndMobilityPolicyData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadAccessAndMobilityPolicyData(ctx.Context(), ueId)

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
	rsp, prob := prod.HandleCreateOrReplaceUEPolicySet(ctx.Context(), ueId, body)

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
	rsp, prob := prod.HandleUpdateSessionManagementPolicyData(ctx.Context(), ueId, body)

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

	// read 'usageMonId'
	params.UsageMonId = ctx.Param("usageMonId")
	if len(params.UsageMonId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "usageMonId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
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
	headers, rsp, prob := prod.HandleCreateUsageMonitoringResource(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
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
	prob := prod.HandleDeleteUsageMonitoringInformation(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

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
	headers, rsp, prob := prod.HandleCreateIndividualBdtData(ctx.Context(), bdtReferenceId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
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
	rsp, prob := prod.HandleReplaceIndividualPolicyDataSubscription(ctx.Context(), subsId, body)

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

func OnReadPolicyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ReadPolicyDataParams

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

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

	// call application handler
	rsp, prob := prod.HandleReadPolicyData(ctx.Context(), &params)

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

	// read 'supp-feat'
	params.SuppFeat = ctx.Param("supp-feat")

	// read 'bdtReferenceId'
	params.BdtReferenceId = ctx.Param("bdtReferenceId")
	if len(params.BdtReferenceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "bdtReferenceId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualBdtData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleReplaceOperatorSpecificData(ctx.Context(), ueId, body)

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
	rsp, prob := prod.HandleUpdateOperatorSpecificData(ctx.Context(), ueId, body)

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
	rsp, prob := prod.HandleReadOperatorSpecificData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleReadSlicePolicyControlData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleUpdateSlicePolicyControlData(ctx.Context(), snssai, body)

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
	HandleUpdateUEPolicySet(context.Context, string, *models.UePolicySetPatch) *models.ProblemDetails

	HandleReadUEPolicySet(context.Context, *ReadUEPolicySetParams) (*models.UePolicySet, *models.ProblemDetails)

	HandleReadSessionManagementPolicyData(context.Context, *ReadSessionManagementPolicyDataParams) (*models.SmPolicyData, *models.ProblemDetails)

	HandleReadUsageMonitoringInformation(context.Context, *ReadUsageMonitoringInformationParams) (*models.UsageMonData, *models.ProblemDetails)

	HandleReadSponsorConnectivityData(context.Context, string) (*models.SponsorConnectivityData, *models.ProblemDetails)

	HandleReadBdtData(context.Context, *ReadBdtDataParams) (*[]models.BdtData, *models.ProblemDetails)

	HandleDeleteIndividualBdtData(context.Context, string) *models.ProblemDetails

	HandleUpdateIndividualBdtData(context.Context, string, *models.BdtDataPatch) (*models.BdtData, *models.ProblemDetails)

	HandleCreateIndividualPolicyDataSubscription(context.Context, *models.PolicyDataSubscription) (map[string]string, *models.PolicyDataSubscription, *models.ProblemDetails)

	HandleDeleteIndividualPolicyDataSubscription(context.Context, string) *models.ProblemDetails

	HandleDeleteOperatorSpecificData(context.Context, string) *models.ProblemDetails

	HandleReadPlmnUePolicySet(context.Context, string) (*models.UePolicySet, *models.ProblemDetails)

	HandleGetMBSSessPolCtrlData(context.Context, *models.MbsSessPolDataId) (*models.MbsSessPolCtrlData, *models.ProblemDetails)

	HandleReadAccessAndMobilityPolicyData(context.Context, string) (*models.AmPolicyData, *models.ProblemDetails)

	HandleCreateOrReplaceUEPolicySet(context.Context, string, *models.UePolicySet) (*models.UePolicySet, *models.ProblemDetails)

	HandleUpdateSessionManagementPolicyData(context.Context, string, *models.SmPolicyDataPatch) (*models.SmPolicyData, *models.ProblemDetails)

	HandleCreateUsageMonitoringResource(context.Context, *CreateUsageMonitoringResourceParams, *models.UsageMonData) (map[string]string, *models.UsageMonData, *models.ProblemDetails)

	HandleDeleteUsageMonitoringInformation(context.Context, *DeleteUsageMonitoringInformationParams) *models.ProblemDetails

	HandleCreateIndividualBdtData(context.Context, string, *models.BdtData) (map[string]string, *models.BdtData, *models.ProblemDetails)

	HandleReplaceIndividualPolicyDataSubscription(context.Context, string, *models.PolicyDataSubscription) (*models.PolicyDataSubscription, *models.ProblemDetails)

	HandleReadPolicyData(context.Context, *ReadPolicyDataParams) (*models.PolicyDataForIndividualUe, *models.ProblemDetails)

	HandleReadIndividualBdtData(context.Context, *ReadIndividualBdtDataParams) (*models.BdtData, *models.ProblemDetails)

	HandleReplaceOperatorSpecificData(context.Context, string, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleUpdateOperatorSpecificData(context.Context, string, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleReadOperatorSpecificData(context.Context, *ReadOperatorSpecificDataParams) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleReadSlicePolicyControlData(context.Context, *ReadSlicePolicyControlDataParams) (*models.SlicePolicyData, *models.ProblemDetails)

	HandleUpdateSlicePolicyControlData(context.Context, *models.Snssai, *models.SlicePolicyDataPatch) (*models.SlicePolicyData, *models.ProblemDetails)
}
