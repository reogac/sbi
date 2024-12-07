/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnQueryNssaiAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryNssaiAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryOperSpecData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryOperSpecDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryOperSpecData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateEeSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.EeSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateEeSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

}

func OnRemoveMultipleSubscriptionDataSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params RemoveMultipleSubscriptionDataSubscriptionsParams

	// read 'implicit-unsubscribe-indication'
	implicitUnsubscribeIndicationStr := ctx.Param("implicit-unsubscribe-indication")
	if len(implicitUnsubscribeIndicationStr) > 0 {
		var implicitUnsubscribeIndicationTmp bool
		if implicitUnsubscribeIndicationTmp, err = models.BoolFromString(implicitUnsubscribeIndicationStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse implicit-unsubscribe-indication failed: %+v", err)))
			return
		}
		params.ImplicitUnsubscribeIndication = &implicitUnsubscribeIndicationTmp
	}

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"))
		return
	}

	// read 'nf-instance-id'
	params.NfInstanceId = ctx.Param("nf-instance-id")

	// read 'delete-all-nfs'
	deleteAllNfsStr := ctx.Param("delete-all-nfs")
	if len(deleteAllNfsStr) > 0 {
		var deleteAllNfsTmp bool
		if deleteAllNfsTmp, err = models.BoolFromString(deleteAllNfsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse delete-all-nfs failed: %+v", err)))
			return
		}
		params.DeleteAllNfs = &deleteAllNfsTmp
	}

	// call application handler
	prod.HandleRemoveMultipleSubscriptionDataSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.AuthEvent)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationStatus(ueId, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuerySmsData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerySmsDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQuerySmsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySubscriptionDataSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQuerySubscriptionDataSubscriptions(subsId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryProvisionedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryProvisionedDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryProvisionedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemoveSmfSubscriptionsInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveSmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveSmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemoveEeGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveEeGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnAmfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params AmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContextNon3gpp(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateSdmSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateSdmSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

}

func OnQueryLcsBcaData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryLcsBcaDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryLcsBcaData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyAmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnDeleteIndividualAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params DeleteIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"))
		return
	}

	// call application handler
	prod.HandleDeleteIndividualAuthenticationStatus(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryEeGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEeGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnDeletePPDataEntry(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params DeletePPDataEntryParams

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"))
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeletePPDataEntry(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnModifyPpData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyPpDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyPpData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetAmfSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyEeGroupSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEeGroupSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryEeGroupSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleQueryEeGroupSubscription(&params)

	// success
	ctx.WriteResponse(200, nil)

}

func OnQuery5mbsData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params Query5mbsDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	rsp := prod.HandleQuery5mbsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryAuthUPU(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryAuthUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthUPU(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnGetppData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetppDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetppData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetPPDataEntry(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetPPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetPPDataEntry(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnModifysdmSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysdmSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGet5GVnGroupConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGet5GVnGroupConfiguration(externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryLcsPrivacyData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryLcsPrivacyDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp := prod.HandleQueryLcsPrivacyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateServiceSpecificAuthorizationInfoParams

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"))
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.ServiceSpecificAuthorizationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateServiceSpecificAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryRoamingInformation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryRoamingInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryAmfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryAmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAmfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyIpSmGwContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleModifyIpSmGwContext(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryPPData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryPPDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyAmfSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyAmfSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateNIDDAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.NiddAuthorizationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateNIDDAuthorizationInfo(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateOperSpecData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateOperSpecData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetHssSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnGetAmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnAmfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params AmfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContext3gpp(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateSmsfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.SmsfRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryeeSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryeeSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleQueryeeSubscription(&params)

	// success
	ctx.WriteResponse(200, nil)

}

func OnModifysubscriptionDataSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifysubscriptionDataSubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysubscriptionDataSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryContextData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryContextDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'context-dataset-names'
	contextDatasetNamesStr := ctx.Param("context-dataset-names")
	if len(contextDatasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "context-dataset-names is required"))
		return
	}

	if params.ContextDatasetNames, err = models.ArrayOfStringFromString(contextDatasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse context-dataset-names failed: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleQueryContextData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnUpdateRoamingInformation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.RoamingInfoUpdate)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleUpdateRoamingInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateAuthenticationSoR(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	body := new(models.SorData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationSoR(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnDeleteIpSmGwContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleDeleteIpSmGwContext(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryMessageWaitingData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryMessageWaitingDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryMessageWaitingData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySubsToNotify(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerySubsToNotifyParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySubsToNotify(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemovesubscriptionDataSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemovesubscriptionDataSubscriptions(subsId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnModify5GVnGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params Modify5GVnGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GVnGroup(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetMulticastMbsGroupMemb(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMulticastMbsGroupMemb(externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnDeleteOperSpecData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeleteOperSpecData(ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSharedDataParams

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "shared-data-ids is required"))
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQueryLcsMoData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryLcsMoDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp := prod.HandleQueryLcsMoData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuery5GVnGroupInternal(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"))
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroupInternal(internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemoveAmfSubscriptionsInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveAmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveAmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnDeleteMessageWaitingData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleDeleteMessageWaitingData(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemoveHssSubscriptionsInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveHssSubscriptionsInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveHssSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuery5GVnGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroup(gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnGetNiddAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetNiddAuthorizationInfo(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryIndividualAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryIndividualAuthenticationStatusParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryIndividualAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateHSSSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateHSSSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.HssSubscriptionInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateHSSSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdateEeGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.EeSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUpdateEeGroupSubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateSMFSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateSMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.SmfSubscriptionInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateSMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnModifySmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifySmfGroupSubscriptionsParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateIndividualAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateIndividualAuthenticationStatusParams

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"))
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.AuthEvent)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateIndividualAuthenticationStatus(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuerySmfRegList(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerySmfRegListParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegList(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnDeleteSmsfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContextNon3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateHSSSDMSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateHSSSDMSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.HssSubscriptionInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateHSSSDMSubscriptions(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryUeLocation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryUeLocationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryUeLocation(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryV2xData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryV2xDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryV2xData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryAuthSubsData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryAuthSubsDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthSubsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySmsMngData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerySmsMngDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQuerySmsMngData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryUeSubscribedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryUeSubscribedDataParams

	// read 'serving-plmn'
	params.ServingPlmn = ctx.Param("serving-plmn")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryUeSubscribedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateAmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateAmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateAmfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.Amf3GppAccessRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateAmfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.AmfNon3GppAccessRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryEEData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryEEDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyNiddAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyNiddAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyNiddAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreate5GVnGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// decode request body
	body := new(models.FiveGVnGroupConfiguration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GVnGroup(externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQueryPorseData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryPorseDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	rsp := prod.HandleQueryPorseData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnGetSmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemoveSmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveSmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateAMFSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateAMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateAMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetSmfSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetSmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateSmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.SmfSubscriptionInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetMultiplePPDataEntries(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetMultiplePPDataEntriesParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetMultiplePPDataEntries(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnModifyMessageWaitingData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleModifyMessageWaitingData(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreatePPDataEntry(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreatePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"))
		return
	}

	// decode request body
	body := new(models.PpDataEntry)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreatePPDataEntry(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdateEesubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateEesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.EeSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUpdateEesubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemoveeeSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveeeSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveeeSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryPeiInformation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryPeiInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuery5GMbsGroupInternal(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"))
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupInternal(internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQueryAmData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryAmDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryCoverageRestrictionData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryCoverageRestrictionDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryCoverageRestrictionData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModify5GmbsGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params Modify5GmbsGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GmbsGroup(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryeesubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryeesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'event-types'
	eventTypesStr := ctx.Param("event-types")
	if len(eventTypesStr) > 0 {
		if params.EventTypes, err = models.ArrayOfStringFromString(eventTypesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse event-types failed: %+v", err)))
			return
		}
	}

	// read 'nf-identifiers'
	nfIdentifiersStr := ctx.Param("nf-identifiers")
	if len(nfIdentifiersStr) > 0 {
		if params.NfIdentifiers, err = models.ArrayOfNfIdentifierFromString(nfIdentifiersStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse nf-identifiers failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryeesubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySmsfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QuerySmsfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmsfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemoveNiddAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveNiddAuthorizationInfo(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemoveServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"))
		return
	}

	// call application handler
	prod.HandleRemoveServiceSpecificAuthorizationInfo(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnDeleteSmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params DeleteSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"))
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)))
		return
	}

	// call application handler
	prod.HandleDeleteSmfRegistration(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuerysdmsubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerysdmsubscriptionsParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQuerysdmsubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryTraceData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryTraceDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryTraceData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnRemoveAmfGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveAmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryAuthSoR(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryAuthSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthSoR(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryCagAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryCagAckParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryCagAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQueryAmfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryAmfContext3gppParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryAmfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QuerySmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"))
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegistration(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifySmfSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifySmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemovesdmSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemovesdmSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prob := prod.HandleRemovesdmSubscriptions(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetGroupIdentifiersParams

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)))
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQuery5GVNGroupPPData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params Query5GVNGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuery5GVNGroupPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateAuthenticationUPU(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateAuthenticationUPUParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.UpuData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationUPU(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuery5GmbsGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GmbsGroup(gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnCreateOrUpdateNssaiAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateOrUpdateNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	body := new(models.NssaiAckData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateOrUpdateNssaiAck(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdateSmfContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateSmfContextParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"))
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSmfContext(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateMessageWaitingData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.MessageWaitingData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateMessageWaitingData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdatesdmsubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdatesdmsubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUpdatesdmsubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetOdbData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetOdbData(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateOrUpdatePeiInformation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.PeiUpdateInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdatePeiInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnDeleteAuthenticationStatus(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleDeleteAuthenticationStatus(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnModifyEesubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyEesubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEesubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnDelete5GmbsGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDelete5GmbsGroup(externalGroupId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateCagUpdateAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateCagUpdateAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	body := new(models.CagAckData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateCagUpdateAck(&params, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetHssSDMSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSDMSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyHssSDMSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyHssSDMSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSDMSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetIdentityData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetIdentityDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetIdentityData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnModifyServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyServiceSpecificAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryUserConsentData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryUserConsentDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ucPurpose'
	params.UcPurpose = ctx.Param("ucPurpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp := prod.HandleQueryUserConsentData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreate5GmbsGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// decode request body
	body := new(models.MulticastMbsGroupMemb)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GmbsGroup(externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQueryGroupEEData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QueryGroupEEDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// call application handler
	rsp := prod.HandleQueryGroupEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnQuerySmfSelectData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QuerySmfSelectDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp := prod.HandleQuerySmfSelectData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateOrUpdateSmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params CreateOrUpdateSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"))
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)))
		return
	}

	// decode request body
	body := new(models.SmfRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdateSmfRegistration(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnModifyOperSpecData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyOperSpecData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuerySmsfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QuerySmsfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmsfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyHssSubscriptionInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyHssSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnRemoveHssSDMSubscriptionsInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params RemoveHssSDMSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleRemoveHssSDMSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdateAuthenticationSoR(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateAuthenticationSoRParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateAuthenticationSoR(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnCreateEeGroupSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueGroupId'
	var ueGroupId string
	ueGroupId = ctx.Param("ueGroupId")
	if len(ueGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"))
		return
	}

	// decode request body
	body := new(models.EeSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateEeGroupSubscriptions(ueGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetIndividualSharedDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	params.SharedDataId = ctx.Param("sharedDataId")
	if len(params.SharedDataId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetIndividualSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnDelete5GVnGroup(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"))
		return
	}

	// call application handler
	prod.HandleDelete5GVnGroup(externalGroupId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetSSAuData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSSAuDataParams

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"))
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"))
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"))
		return
	}

	// call application handler
	prob := prod.HandleGetSSAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(200, nil)

}

func OnQuery5GMbsGroupPPData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params Query5GMbsGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnQuerySmData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QuerySmDataParams

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"))
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuerySmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnDeleteSmsfContext3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContext3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil)

}

func OnQuerysdmSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params QuerysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"))
		return
	}

	// call application handler
	prod.HandleQuerysdmSubscription(&params)

	// success
	ctx.WriteResponse(200, nil)

}

func OnSubscriptionDataSubscriptions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.SubscriptionDataSubscriptions)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleSubscriptionDataSubscriptions(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

}

func OnGetServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"))
		return
	}

	// call application handler
	rsp := prod.HandleGetServiceSpecificAuthorizationInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnModifyAuthenticationSubscription(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyAuthenticationSubscriptionParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAuthenticationSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnQueryIpSmGwContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params QueryIpSmGwContextParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)))
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryIpSmGwContext(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

}

func OnCreateIpSmGwContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.IpSmGwRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prod.HandleCreateIpSmGwContext(ueId, body)

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetNiddAuData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetNiddAuDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"))
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"))
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")
	if len(params.MtcProviderInformation) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "mtc-provider-information is required"))
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// call application handler
	prob := prod.HandleGetNiddAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(200, nil)

}

func OnCreateSmsfContextNon3gpp(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.SmsfRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleQueryNssaiAck(*QueryNssaiAckParams) *models.NssaiAckData

	HandleQueryOperSpecData(*QueryOperSpecDataParams) *map[string]models.OperatorSpecificDataContainer

	HandleCreateEeSubscriptions(string, *models.EeSubscription) *models.EeSubscription

	HandleRemoveMultipleSubscriptionDataSubscriptions(*RemoveMultipleSubscriptionDataSubscriptionsParams)

	HandleCreateAuthenticationStatus(string, *models.AuthEvent)

	HandleQuerySmsData(*QuerySmsDataParams) *models.SmsSubscriptionData

	HandleQuerySubscriptionDataSubscriptions(string) *models.SubscriptionDataSubscriptions

	HandleQueryProvisionedData(*QueryProvisionedDataParams) *models.ProvisionedDataSets

	HandleRemoveSmfSubscriptionsInfo(*RemoveSmfSubscriptionsInfoParams)

	HandleRemoveEeGroupSubscriptions(*RemoveEeGroupSubscriptionsParams)

	HandleAmfContextNon3gpp(*AmfContextNon3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateSdmSubscriptions(string, *models.SdmSubscription) *models.SdmSubscription

	HandleQueryLcsBcaData(*QueryLcsBcaDataParams) *models.LcsBroadcastAssistanceTypesData

	HandleModifyAmfGroupSubscriptions(*ModifyAmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDeleteIndividualAuthenticationStatus(*DeleteIndividualAuthenticationStatusParams)

	HandleQueryEeGroupSubscriptions(*QueryEeGroupSubscriptionsParams) *[]models.EeSubscription

	HandleDeletePPDataEntry(*DeletePPDataEntryParams) *models.ProblemDetails

	HandleModifyPpData(*ModifyPpDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetAmfSubscriptionInfo(*GetAmfSubscriptionInfoParams) *[]models.AmfSubscriptionInfo

	HandleModifyEeGroupSubscription(*ModifyEeGroupSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryEeGroupSubscription(*QueryEeGroupSubscriptionParams)

	HandleQuery5mbsData(*Query5mbsDataParams) *models.MbsSubscriptionData

	HandleQueryAuthUPU(*QueryAuthUPUParams) *models.UpuData

	HandleGetppData(*GetppDataParams) (*models.PpData, *models.ProblemDetails)

	HandleGetPPDataEntry(*GetPPDataEntryParams) (*models.PpDataEntry, *models.ProblemDetails)

	HandleModifysdmSubscription(*ModifysdmSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGet5GVnGroupConfiguration(string) *models.FiveGVnGroupConfiguration

	HandleQueryLcsPrivacyData(*QueryLcsPrivacyDataParams) *models.LcsPrivacyData

	HandleCreateServiceSpecificAuthorizationInfo(*CreateServiceSpecificAuthorizationInfoParams, *models.ServiceSpecificAuthorizationInfo) *models.ServiceSpecificAuthorizationInfo

	HandleQueryRoamingInformation(string) *models.RoamingInfoUpdate

	HandleQueryAmfContextNon3gpp(*QueryAmfContextNon3gppParams) *models.AmfNon3GppAccessRegistration

	HandleModifyIpSmGwContext(string, *[]models.PatchItem) *models.ProblemDetails

	HandleQueryPPData(*QueryPPDataParams) *models.PpProfileData

	HandleModifyAmfSubscriptionInfo(*ModifyAmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateNIDDAuthorizationInfo(string, *models.NiddAuthorizationInfo) *models.NiddAuthorizationInfo

	HandleCreateOperSpecData(*CreateOperSpecDataParams, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleGetHssSubscriptionInfo(*GetHssSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleGetAmfGroupSubscriptions(*GetAmfGroupSubscriptionsParams) *[]models.AmfSubscriptionInfo

	HandleAmfContext3gpp(*AmfContext3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateSmsfContext3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleQueryeeSubscription(*QueryeeSubscriptionParams)

	HandleModifysubscriptionDataSubscription(*ModifysubscriptionDataSubscriptionParams, *[]models.PatchItem) (*models.Schema, *models.ProblemDetails)

	HandleQueryContextData(*QueryContextDataParams) *models.ContextDataSets

	HandleUpdateRoamingInformation(string, *models.RoamingInfoUpdate) *models.RoamingInfoUpdate

	HandleCreateAuthenticationSoR(*CreateAuthenticationSoRParams, *models.SorData)

	HandleDeleteIpSmGwContext(string)

	HandleQueryMessageWaitingData(*QueryMessageWaitingDataParams) *models.MessageWaitingData

	HandleQuerySubsToNotify(*QuerySubsToNotifyParams) *[]models.SubscriptionDataSubscriptions

	HandleRemovesubscriptionDataSubscriptions(string)

	HandleModify5GVnGroup(*Modify5GVnGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetMulticastMbsGroupMemb(string) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleDeleteOperSpecData(string) *models.ProblemDetails

	HandleGetSharedData(*GetSharedDataParams) (*[]models.SharedData, *models.ProblemDetails)

	HandleQueryLcsMoData(*QueryLcsMoDataParams) *models.LcsMoData

	HandleQuery5GVnGroupInternal([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleRemoveAmfSubscriptionsInfo(*RemoveAmfSubscriptionsInfoParams)

	HandleDeleteMessageWaitingData(string)

	HandleRemoveHssSubscriptionsInfo(*RemoveHssSubscriptionsInfoParams)

	HandleQuery5GVnGroup([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleGetNiddAuthorizationInfo(string) *models.NiddAuthorizationInfo

	HandleQueryIndividualAuthenticationStatus(*QueryIndividualAuthenticationStatusParams) *models.AuthEvent

	HandleCreateHSSSubscriptions(*CreateHSSSubscriptionsParams, *models.HssSubscriptionInfo) *models.HssSubscriptionInfo

	HandleUpdateEeGroupSubscriptions(*UpdateEeGroupSubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleCreateSMFSubscriptions(*CreateSMFSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleModifySmfGroupSubscriptions(*ModifySmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateIndividualAuthenticationStatus(*CreateIndividualAuthenticationStatusParams, *models.AuthEvent)

	HandleQuerySmfRegList(*QuerySmfRegListParams) *[]models.SmfRegistration

	HandleDeleteSmsfContextNon3gpp(string)

	HandleCreateHSSSDMSubscriptions(*CreateHSSSDMSubscriptionsParams, *models.HssSubscriptionInfo)

	HandleQueryUeLocation(*QueryUeLocationParams) *models.LocationInfo

	HandleQueryV2xData(*QueryV2xDataParams) *models.V2xSubscriptionData

	HandleQueryAuthSubsData(*QueryAuthSubsDataParams) *models.AuthenticationSubscription

	HandleQuerySmsMngData(*QuerySmsMngDataParams) *models.SmsManagementSubscriptionData

	HandleQueryUeSubscribedData(*QueryUeSubscribedDataParams) *models.UeSubscribedDataSets

	HandleCreateAmfGroupSubscriptions(*CreateAmfGroupSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleCreateAmfContext3gpp(string, *models.Amf3GppAccessRegistration) *models.Amf3GppAccessRegistration

	HandleCreateAmfContextNon3gpp(string, *models.AmfNon3GppAccessRegistration) *models.Amf3GppAccessRegistration

	HandleQueryEEData(*QueryEEDataParams) *models.EeProfileData

	HandleModifyNiddAuthorizationInfo(*ModifyNiddAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryAuthenticationStatus(*QueryAuthenticationStatusParams) *models.AuthEvent

	HandleCreate5GVnGroup(string, *models.FiveGVnGroupConfiguration) (*models.FiveGVnGroupConfiguration, *models.ProblemDetails)

	HandleQueryPorseData(*QueryPorseDataParams) *models.ProseSubscriptionData

	HandleGetSmfGroupSubscriptions(*GetSmfGroupSubscriptionsParams) *models.SmfSubscriptionInfo

	HandleRemoveSmfGroupSubscriptions(*RemoveSmfGroupSubscriptionsParams)

	HandleCreateAMFSubscriptions(*CreateAMFSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleGetSmfSubscriptionInfo(*GetSmfSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleCreateSmfGroupSubscriptions(*CreateSmfGroupSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleGetMultiplePPDataEntries(*GetMultiplePPDataEntriesParams) (*models.PpDataEntryList, *models.ProblemDetails)

	HandleModifyMessageWaitingData(string, *[]models.PatchItem) *models.ProblemDetails

	HandleCreatePPDataEntry(*CreatePPDataEntryParams, *models.PpDataEntry) (*models.PpDataEntry, *models.ProblemDetails)

	HandleUpdateEesubscriptions(*UpdateEesubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleRemoveeeSubscriptions(*RemoveeeSubscriptionsParams)

	HandleQueryPeiInformation(string) *models.PeiUpdateInfo

	HandleQuery5GMbsGroupInternal([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleQueryAmData(*QueryAmDataParams) *models.AccessAndMobilitySubscriptionData

	HandleQueryCoverageRestrictionData(*QueryCoverageRestrictionDataParams) *models.EnhancedCoverageRestrictionData

	HandleModify5GmbsGroup(*Modify5GmbsGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryeesubscriptions(*QueryeesubscriptionsParams) *[]models.EeSubscriptionExt

	HandleQuerySmsfContext3gpp(*QuerySmsfContext3gppParams) *models.SmsfRegistration

	HandleRemoveNiddAuthorizationInfo(string)

	HandleRemoveServiceSpecificAuthorizationInfo(*RemoveServiceSpecificAuthorizationInfoParams)

	HandleDeleteSmfRegistration(*DeleteSmfRegistrationParams)

	HandleQuerysdmsubscriptions(*QuerysdmsubscriptionsParams) *[]models.SdmSubscription

	HandleQueryTraceData(*QueryTraceDataParams) *models.TraceData

	HandleRemoveAmfGroupSubscriptions(*RemoveAmfGroupSubscriptionsParams)

	HandleQueryAuthSoR(*QueryAuthSoRParams) *models.SorData

	HandleQueryCagAck(*QueryCagAckParams) *models.CagAckData

	HandleQueryAmfContext3gpp(*QueryAmfContext3gppParams) *models.Amf3GppAccessRegistration

	HandleQuerySmfRegistration(*QuerySmfRegistrationParams) *models.SmfRegistration

	HandleModifySmfSubscriptionInfo(*ModifySmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleRemovesdmSubscriptions(*RemovesdmSubscriptionsParams) *models.ProblemDetails

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleQuery5GVNGroupPPData(*Query5GVNGroupPPDataParams) *models.Pp5gVnGroupProfileData

	HandleCreateAuthenticationUPU(*CreateAuthenticationUPUParams, *models.UpuData)

	HandleQuery5GmbsGroup([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleCreateOrUpdateNssaiAck(*CreateOrUpdateNssaiAckParams, *models.NssaiAckData)

	HandleUpdateSmfContext(*UpdateSmfContextParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateMessageWaitingData(string, *models.MessageWaitingData) *models.MessageWaitingData

	HandleUpdatesdmsubscriptions(*UpdatesdmsubscriptionsParams, *models.SdmSubscription) *models.ProblemDetails

	HandleGetOdbData(string) *models.OdbData

	HandleCreateOrUpdatePeiInformation(string, *models.PeiUpdateInfo) *models.PeiUpdateInfo

	HandleDeleteAuthenticationStatus(string)

	HandleModifyEesubscription(*ModifyEesubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDelete5GmbsGroup(string) *models.ProblemDetails

	HandleCreateCagUpdateAck(*CreateCagUpdateAckParams, *models.CagAckData)

	HandleGetHssSDMSubscriptionInfo(*GetHssSDMSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleModifyHssSDMSubscriptionInfo(*ModifyHssSDMSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetIdentityData(*GetIdentityDataParams) (*models.IdentityData, *models.ProblemDetails)

	HandleModifyServiceSpecificAuthorizationInfo(*ModifyServiceSpecificAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryUserConsentData(*QueryUserConsentDataParams) *models.UcSubscriptionData

	HandleCreate5GmbsGroup(string, *models.MulticastMbsGroupMemb) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleQueryGroupEEData(*QueryGroupEEDataParams) *models.EeGroupProfileData

	HandleQuerySmfSelectData(*QuerySmfSelectDataParams) *models.SmfSelectionSubscriptionData

	HandleCreateOrUpdateSmfRegistration(*CreateOrUpdateSmfRegistrationParams, *models.SmfRegistration) *models.SmfRegistration

	HandleModifyOperSpecData(*ModifyOperSpecDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuerySmsfContextNon3gpp(*QuerySmsfContextNon3gppParams) *models.SmsfRegistration

	HandleModifyHssSubscriptionInfo(*ModifyHssSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleRemoveHssSDMSubscriptionsInfo(*RemoveHssSDMSubscriptionsInfoParams)

	HandleUpdateAuthenticationSoR(*UpdateAuthenticationSoRParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateEeGroupSubscriptions(string, *models.EeSubscription) *models.EeSubscription

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (*models.SharedData, *models.ProblemDetails)

	HandleDelete5GVnGroup(string)

	HandleGetSSAuData(*GetSSAuDataParams) *models.ProblemDetails

	HandleQuery5GMbsGroupPPData(*Query5GMbsGroupPPDataParams) (*models.Pp5gMbsGroupProfileData, *models.ProblemDetails)

	HandleQuerySmData(*QuerySmDataParams) *models.SmSubsData

	HandleDeleteSmsfContext3gpp(string)

	HandleQuerysdmSubscription(*QuerysdmSubscriptionParams)

	HandleSubscriptionDataSubscriptions(*models.SubscriptionDataSubscriptions) *models.SubscriptionDataSubscriptions

	HandleGetServiceSpecificAuthorizationInfo(*GetServiceSpecificAuthorizationInfoParams) *models.ServiceSpecificAuthorizationInfo

	HandleModifyAuthenticationSubscription(*ModifyAuthenticationSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryIpSmGwContext(*QueryIpSmGwContextParams) *models.IpSmGwRegistration

	HandleCreateIpSmGwContext(string, *models.IpSmGwRegistration)

	HandleGetNiddAuData(*GetNiddAuDataParams) *models.ProblemDetails

	HandleCreateSmsfContextNon3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration
}
