/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateHSSSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateHSSSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetOdbData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetOdbData(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfGroupSubscriptions(&params, body)

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

func OnDeleteIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var params DeleteIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIndividualAuthenticationStatus(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyOperSpecData(&params, body)

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

func OnCreateIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.IpSmGwRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIpSmGwContext(ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryeesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryeesubscriptionsParams

	// read 'event-types'
	eventTypesStr := ctx.Param("event-types")
	if len(eventTypesStr) > 0 {
		if params.EventTypes, err = models.ArrayOfStringFromString(eventTypesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse event-types failed: %+v", err)), nil)
			return
		}
	}

	// read 'nf-identifiers'
	nfIdentifiersStr := ctx.Param("nf-identifiers")
	if len(nfIdentifiersStr) > 0 {
		if params.NfIdentifiers, err = models.ArrayOfNfIdentifierFromString(nfIdentifiersStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse nf-identifiers failed: %+v", err)), nil)
			return
		}
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryeesubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetPPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params GetPPDataEntryParams

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetPPDataEntry(&params)

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

func OnCreateAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContextNon3gppParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAmfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnDeleteOperSpecData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteOperSpecData(ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryMessageWaitingDataParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryMessageWaitingData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyPpData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyPpDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
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
	rsp, prob := prod.HandleModifyPpData(&params, body)

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

func OnQueryIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIndividualAuthenticationStatusParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryIndividualAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryLcsBcaData(ctx sbi.RequestContext, prod Producer) {
	var params QueryLcsBcaDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQueryLcsBcaData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnRemoveNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveNiddAuthorizationInfo(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAuthenticationUPU(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UpuData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationUPU(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSubscriptionInfo(&params, body)

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

func OnModifyHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSDMSubscriptionInfo(&params, body)

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

func OnGetHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSDMSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveMultipleSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params RemoveMultipleSubscriptionDataSubscriptionsParams

	// read 'delete-all-nfs'
	deleteAllNfsStr := ctx.Param("delete-all-nfs")
	if len(deleteAllNfsStr) > 0 {
		var deleteAllNfsTmp bool
		if deleteAllNfsTmp, err = models.BoolFromString(deleteAllNfsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse delete-all-nfs failed: %+v", err)), nil)
			return
		}
		params.DeleteAllNfs = &deleteAllNfsTmp
	}

	// read 'implicit-unsubscribe-indication'
	implicitUnsubscribeIndicationStr := ctx.Param("implicit-unsubscribe-indication")
	if len(implicitUnsubscribeIndicationStr) > 0 {
		var implicitUnsubscribeIndicationTmp bool
		if implicitUnsubscribeIndicationTmp, err = models.BoolFromString(implicitUnsubscribeIndicationStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse implicit-unsubscribe-indication failed: %+v", err)), nil)
			return
		}
		params.ImplicitUnsubscribeIndication = &implicitUnsubscribeIndicationTmp
	}

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'nf-instance-id'
	params.NfInstanceId = ctx.Param("nf-instance-id")

	// call application handler
	prod.HandleRemoveMultipleSubscriptionDataSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGet5GVnGroupConfiguration(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGet5GVnGroupConfiguration(externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryUeLocation(ctx sbi.RequestContext, prod Producer) {
	var params QueryUeLocationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryUeLocation(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmfSelectData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfSelectDataParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmfSelectData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryOperSpecDataParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryOperSpecData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuerysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQuerysdmSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQueryGroupEEData(ctx sbi.RequestContext, prod Producer) {
	var params QueryGroupEEDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryGroupEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdatesdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdatesdmsubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdatesdmsubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, prod Producer) {
	var params GetIndividualSharedDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	params.SharedDataId = ctx.Param("sharedDataId")
	if len(params.SharedDataId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetIndividualSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateNIDDAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.NiddAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateNIDDAuthorizationInfo(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAuthSoR(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSoRParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAuthSoR(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryAmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnDeleteSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContextNon3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEeGroupSubscription(&params, body)

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

func OnUpdateRoamingInformation(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.RoamingInfoUpdate)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleUpdateRoamingInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryContextData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryContextDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'context-dataset-names'
	contextDatasetNamesStr := ctx.Param("context-dataset-names")
	if len(contextDatasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "context-dataset-names is required"), nil)
		return
	}

	if params.ContextDatasetNames, err = models.ArrayOfStringFromString(contextDatasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse context-dataset-names failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryContextData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateEeSubscriptions(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnRemoveSmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfSubscriptionsInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveEeGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveHssSDMSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSDMSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSDMSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryUserConsentData(ctx sbi.RequestContext, prod Producer) {
	var params QueryUserConsentDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ucPurpose'
	params.UcPurpose = ctx.Param("ucPurpose")

	// call application handler
	headers, rsp := prod.HandleQueryUserConsentData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryAuthSubsData(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSubsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAuthSubsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryEeGroupSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnModifyNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyNiddAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyNiddAuthorizationInfo(&params, body)

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

func OnQueryRoamingInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryRoamingInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationStatus(ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySubsToNotify(ctx sbi.RequestContext, prod Producer) {
	var params QuerySubsToNotifyParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySubsToNotify(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContext3gppParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAmfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetMultiplePPDataEntries(ctx sbi.RequestContext, prod Producer) {
	var params GetMultiplePPDataEntriesParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetMultiplePPDataEntries(&params)

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

func OnGetAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuery5GVNGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GVNGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuery5GVNGroupPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateOrUpdatePeiInformation(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.PeiUpdateInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdatePeiInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GmbsGroup(gpsis)

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

func OnQueryAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAuthenticationStatusParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnDeleteSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContext3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteIpSmGwContext(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIpSmGwContext(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateSdmSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnQuerySmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuerySmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuery5mbsData(ctx sbi.RequestContext, prod Producer) {
	var params Query5mbsDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp := prod.HandleQuery5mbsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateOrUpdateNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NssaiAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateOrUpdateNssaiAck(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryLcsMoData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsMoDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp := prod.HandleQueryLcsMoData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuery5GVnGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroupInternal(internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifySmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfGroupSubscriptionsParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfGroupSubscriptions(&params, body)

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

func OnAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContextNon3gpp(&params, body)

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

func OnGetGroupIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetGroupIdentifiersParams

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)), nil)
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(&params)

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

func OnQueryLcsPrivacyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsPrivacyDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryLcsPrivacyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionDataSubscriptions)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleSubscriptionDataSubscriptions(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnCreate5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.FiveGVnGroupConfiguration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GVnGroup(externalGroupId, body)

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

func OnQueryV2xData(ctx sbi.RequestContext, prod Producer) {
	var params QueryV2xDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQueryV2xData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateOperSpecData(&params, body)

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

func OnQueryIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIpSmGwContextParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryIpSmGwContext(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEeGroupSubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerysdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmsubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerysdmsubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateEesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEesubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryUeSubscribedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryUeSubscribedDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'serving-plmn'
	params.ServingPlmn = ctx.Param("serving-plmn")

	// call application handler
	rsp := prod.HandleQueryUeSubscribedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyServiceSpecificAuthorizationInfoParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyServiceSpecificAuthorizationInfo(&params, body)

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

func OnModifysubscriptionDataSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysubscriptionDataSubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysubscriptionDataSubscription(&params, body)

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

func OnCreateServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ServiceSpecificAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateServiceSpecificAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMulticastMbsGroupMemb(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMulticastMbsGroupMemb(externalGroupId)

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

func OnCreateOrUpdateSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdateSmfRegistration(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteMessageWaitingData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteMessageWaitingData(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveServiceSpecificAuthorizationInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryPeiInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryPeiInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.Amf3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegistration(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifySmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfSubscriptionInfo(&params, body)

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

func OnGetNiddAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetNiddAuDataParams

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")
	if len(params.MtcProviderInformation) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "mtc-provider-information is required"), nil)
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleGetNiddAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnCreatePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreatePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PpDataEntry)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreatePPDataEntry(&params, body)

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

func OnQueryEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEeGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAuthenticationSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAuthenticationSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAuthenticationSubscription(&params, body)

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

func OnDeleteAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteAuthenticationStatus(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
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
	prob := prod.HandleModifyMessageWaitingData(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetppData(ctx sbi.RequestContext, prod Producer) {
	var params GetppDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetppData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQuery5GMbsGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GMbsGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupPPData(&params)

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

func OnUpdateSmfContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateSmfContextParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSmfContext(&params, body)

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

func OnCreateSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryeeSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryeeSubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryeeSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQueryCoverageRestrictionData(ctx sbi.RequestContext, prod Producer) {
	var params QueryCoverageRestrictionDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp := prod.HandleQueryCoverageRestrictionData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateCagUpdateAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateCagUpdateAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.CagAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateCagUpdateAck(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryProvisionedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryProvisionedDataParams

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryProvisionedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContext3gpp(&params, body)

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

func OnQuerySmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContext3gppParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySmsfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsMngData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsMngDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmsMngData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnModifysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysdmSubscription(&params, body)

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

func OnCreateIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIndividualAuthenticationStatus(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SorData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationSoR(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateAuthenticationSoR(&params, body)

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

func OnQueryAuthUPU(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthUPU(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetIdentityData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetIdentityDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetIdentityData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDelete5GVnGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDelete5GVnGroup(externalGroupId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreate5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.MulticastMbsGroupMemb)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GmbsGroup(externalGroupId, body)

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

func OnDeleteSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params DeleteSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmfRegistration(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemovesdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemovesdmSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleRemovesdmSubscriptions(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSSAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSSAuDataParams

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleGetSSAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQuerySmfRegList(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmfRegListParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegList(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
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
	prob := prod.HandleModifyIpSmGwContext(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveeeSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveeeSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveeeSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GVnGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GVnGroup(&params, body)

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

func OnQueryCagAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryCagAckParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryCagAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetSmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemovesubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemovesubscriptionDataSubscriptions(subsId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroup(gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySubscriptionDataSubscriptions(subsId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmsData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuerySmsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnDeletePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params DeletePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeletePPDataEntry(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveHssSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfSubscriptionInfo(&params, body)

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

func OnQueryTraceData(ctx sbi.RequestContext, prod Producer) {
	var params QueryTraceDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryTraceData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.AmfNon3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContextNon3gppParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmsfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.MessageWaitingData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateMessageWaitingData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveAmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryPorseData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPorseDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryPorseData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetNiddAuthorizationInfo(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetServiceSpecificAuthorizationInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryPPData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPPDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyEesubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEesubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEesubscription(&params, body)

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

func OnCreateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueGroupId'
	var ueGroupId string
	ueGroupId = ctx.Param("ueGroupId")
	if len(ueGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeGroupSubscriptions(ueGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnGetSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSharedDataParams

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "shared-data-ids is required"), nil)
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSharedData(&params)

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

func OnDelete5GmbsGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDelete5GmbsGroup(externalGroupId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GmbsGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GmbsGroup(&params, body)

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

func OnQueryNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryNssaiAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryEEData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryEEDataParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateHSSSDMSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSDMSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateHSSSDMSubscriptions(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GMbsGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupInternal(internalGroupIds)

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

type Producer interface {
	HandleCreateHSSSubscriptions(*CreateHSSSubscriptionsParams, *models.HssSubscriptionInfo) *models.HssSubscriptionInfo

	HandleGetOdbData(string) *models.OdbData

	HandleModifyAmfGroupSubscriptions(*ModifyAmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDeleteIndividualAuthenticationStatus(*DeleteIndividualAuthenticationStatusParams)

	HandleModifyOperSpecData(*ModifyOperSpecDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateIpSmGwContext(string, *models.IpSmGwRegistration)

	HandleQueryeesubscriptions(*QueryeesubscriptionsParams) *[]models.EeSubscriptionExt

	HandleGetPPDataEntry(*GetPPDataEntryParams) (*models.PpDataEntry, *models.ProblemDetails)

	HandleCreateAmfGroupSubscriptions(*CreateAmfGroupSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleQueryAmfContextNon3gpp(*QueryAmfContextNon3gppParams) *models.AmfNon3GppAccessRegistration

	HandleDeleteOperSpecData(string) *models.ProblemDetails

	HandleQueryMessageWaitingData(*QueryMessageWaitingDataParams) *models.MessageWaitingData

	HandleModifyPpData(*ModifyPpDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryIndividualAuthenticationStatus(*QueryIndividualAuthenticationStatusParams) *models.AuthEvent

	HandleCreateAMFSubscriptions(*CreateAMFSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleQueryLcsBcaData(*QueryLcsBcaDataParams) (map[string]string, *models.LcsBroadcastAssistanceTypesData)

	HandleRemoveNiddAuthorizationInfo(string)

	HandleCreateAuthenticationUPU(*CreateAuthenticationUPUParams, *models.UpuData)

	HandleModifyHssSubscriptionInfo(*ModifyHssSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyHssSDMSubscriptionInfo(*ModifyHssSDMSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetHssSDMSubscriptionInfo(*GetHssSDMSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleRemoveMultipleSubscriptionDataSubscriptions(*RemoveMultipleSubscriptionDataSubscriptionsParams)

	HandleGet5GVnGroupConfiguration(string) *models.FiveGVnGroupConfiguration

	HandleQueryUeLocation(*QueryUeLocationParams) *models.LocationInfo

	HandleQuerySmfSelectData(*QuerySmfSelectDataParams) (map[string]string, *models.SmfSelectionSubscriptionData)

	HandleQueryOperSpecData(*QueryOperSpecDataParams) (map[string]string, *map[string]models.OperatorSpecificDataContainer)

	HandleQuerysdmSubscription(*QuerysdmSubscriptionParams)

	HandleQueryGroupEEData(*QueryGroupEEDataParams) *models.EeGroupProfileData

	HandleUpdatesdmsubscriptions(*UpdatesdmsubscriptionsParams, *models.SdmSubscription) *models.ProblemDetails

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (map[string]string, *models.SharedData, *models.ProblemDetails)

	HandleCreateNIDDAuthorizationInfo(string, *models.NiddAuthorizationInfo) *models.NiddAuthorizationInfo

	HandleQueryAuthSoR(*QueryAuthSoRParams) *models.SorData

	HandleQueryAmData(*QueryAmDataParams) (map[string]string, *models.AccessAndMobilitySubscriptionData)

	HandleDeleteSmsfContextNon3gpp(string)

	HandleModifyEeGroupSubscription(*ModifyEeGroupSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleUpdateRoamingInformation(string, *models.RoamingInfoUpdate) *models.RoamingInfoUpdate

	HandleQueryContextData(*QueryContextDataParams) *models.ContextDataSets

	HandleCreateEeSubscriptions(string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleRemoveSmfSubscriptionsInfo(*RemoveSmfSubscriptionsInfoParams)

	HandleRemoveEeGroupSubscriptions(*RemoveEeGroupSubscriptionsParams)

	HandleRemoveHssSDMSubscriptionsInfo(*RemoveHssSDMSubscriptionsInfoParams)

	HandleQueryUserConsentData(*QueryUserConsentDataParams) (map[string]string, *models.UcSubscriptionData)

	HandleQueryAuthSubsData(*QueryAuthSubsDataParams) *models.AuthenticationSubscription

	HandleQueryEeGroupSubscription(*QueryEeGroupSubscriptionParams)

	HandleModifyNiddAuthorizationInfo(*ModifyNiddAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryRoamingInformation(string) *models.RoamingInfoUpdate

	HandleCreateAuthenticationStatus(string, *models.AuthEvent)

	HandleQuerySubsToNotify(*QuerySubsToNotifyParams) *[]models.SubscriptionDataSubscriptions

	HandleGetSmfGroupSubscriptions(*GetSmfGroupSubscriptionsParams) *models.SmfSubscriptionInfo

	HandleQueryAmfContext3gpp(*QueryAmfContext3gppParams) *models.Amf3GppAccessRegistration

	HandleGetMultiplePPDataEntries(*GetMultiplePPDataEntriesParams) (*models.PpDataEntryList, *models.ProblemDetails)

	HandleGetAmfSubscriptionInfo(*GetAmfSubscriptionInfoParams) *[]models.AmfSubscriptionInfo

	HandleQuery5GVNGroupPPData(*Query5GVNGroupPPDataParams) *models.Pp5gVnGroupProfileData

	HandleCreateOrUpdatePeiInformation(string, *models.PeiUpdateInfo) *models.PeiUpdateInfo

	HandleQuery5GmbsGroup([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleQueryAuthenticationStatus(*QueryAuthenticationStatusParams) *models.AuthEvent

	HandleDeleteSmsfContext3gpp(string)

	HandleDeleteIpSmGwContext(string)

	HandleCreateSdmSubscriptions(string, *models.SdmSubscription) (map[string]string, *models.SdmSubscription)

	HandleQuerySmData(*QuerySmDataParams) (map[string]string, *models.SmSubsData)

	HandleQuery5mbsData(*Query5mbsDataParams) (map[string]string, *models.MbsSubscriptionData)

	HandleCreateOrUpdateNssaiAck(*CreateOrUpdateNssaiAckParams, *models.NssaiAckData)

	HandleQueryLcsMoData(*QueryLcsMoDataParams) (map[string]string, *models.LcsMoData)

	HandleQuery5GVnGroupInternal([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleCreateSmfGroupSubscriptions(*CreateSmfGroupSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleModifySmfGroupSubscriptions(*ModifySmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleAmfContextNon3gpp(*AmfContextNon3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleQueryLcsPrivacyData(*QueryLcsPrivacyDataParams) (map[string]string, *models.LcsPrivacyData)

	HandleGetAmfGroupSubscriptions(*GetAmfGroupSubscriptionsParams) *[]models.AmfSubscriptionInfo

	HandleSubscriptionDataSubscriptions(*models.SubscriptionDataSubscriptions) (map[string]string, *models.SubscriptionDataSubscriptions)

	HandleCreate5GVnGroup(string, *models.FiveGVnGroupConfiguration) (*models.FiveGVnGroupConfiguration, *models.ProblemDetails)

	HandleQueryV2xData(*QueryV2xDataParams) (map[string]string, *models.V2xSubscriptionData)

	HandleCreateOperSpecData(*CreateOperSpecDataParams, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleQueryIpSmGwContext(*QueryIpSmGwContextParams) *models.IpSmGwRegistration

	HandleUpdateEeGroupSubscriptions(*UpdateEeGroupSubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleQuerysdmsubscriptions(*QuerysdmsubscriptionsParams) *[]models.SdmSubscription

	HandleUpdateEesubscriptions(*UpdateEesubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleCreateSMFSubscriptions(*CreateSMFSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleQueryUeSubscribedData(*QueryUeSubscribedDataParams) *models.UeSubscribedDataSets

	HandleModifyServiceSpecificAuthorizationInfo(*ModifyServiceSpecificAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifysubscriptionDataSubscription(*ModifysubscriptionDataSubscriptionParams, *[]models.PatchItem) (*models.Schema, *models.ProblemDetails)

	HandleCreateServiceSpecificAuthorizationInfo(*CreateServiceSpecificAuthorizationInfoParams, *models.ServiceSpecificAuthorizationInfo) *models.ServiceSpecificAuthorizationInfo

	HandleGetMulticastMbsGroupMemb(string) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleCreateOrUpdateSmfRegistration(*CreateOrUpdateSmfRegistrationParams, *models.SmfRegistration) *models.SmfRegistration

	HandleDeleteMessageWaitingData(string)

	HandleRemoveServiceSpecificAuthorizationInfo(*RemoveServiceSpecificAuthorizationInfoParams)

	HandleQueryPeiInformation(string) *models.PeiUpdateInfo

	HandleCreateAmfContext3gpp(string, *models.Amf3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleQuerySmfRegistration(*QuerySmfRegistrationParams) *models.SmfRegistration

	HandleModifySmfSubscriptionInfo(*ModifySmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetNiddAuData(*GetNiddAuDataParams) *models.ProblemDetails

	HandleCreatePPDataEntry(*CreatePPDataEntryParams, *models.PpDataEntry) (*models.PpDataEntry, *models.ProblemDetails)

	HandleQueryEeGroupSubscriptions(*QueryEeGroupSubscriptionsParams) *[]models.EeSubscription

	HandleModifyAuthenticationSubscription(*ModifyAuthenticationSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDeleteAuthenticationStatus(string)

	HandleModifyMessageWaitingData(string, *[]models.PatchItem) *models.ProblemDetails

	HandleGetppData(*GetppDataParams) (map[string]string, *models.PpData, *models.ProblemDetails)

	HandleQuery5GMbsGroupPPData(*Query5GMbsGroupPPDataParams) (*models.Pp5gMbsGroupProfileData, *models.ProblemDetails)

	HandleUpdateSmfContext(*UpdateSmfContextParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateSmsfContextNon3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleQueryeeSubscription(*QueryeeSubscriptionParams)

	HandleQueryCoverageRestrictionData(*QueryCoverageRestrictionDataParams) (map[string]string, *models.EnhancedCoverageRestrictionData)

	HandleCreateCagUpdateAck(*CreateCagUpdateAckParams, *models.CagAckData)

	HandleQueryProvisionedData(*QueryProvisionedDataParams) *models.ProvisionedDataSets

	HandleAmfContext3gpp(*AmfContext3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuerySmsfContext3gpp(*QuerySmsfContext3gppParams) *models.SmsfRegistration

	HandleQuerySmsMngData(*QuerySmsMngDataParams) (map[string]string, *models.SmsManagementSubscriptionData)

	HandleModifysdmSubscription(*ModifysdmSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateIndividualAuthenticationStatus(*CreateIndividualAuthenticationStatusParams, *models.AuthEvent)

	HandleCreateAuthenticationSoR(*CreateAuthenticationSoRParams, *models.SorData)

	HandleUpdateAuthenticationSoR(*UpdateAuthenticationSoRParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryAuthUPU(*QueryAuthUPUParams) *models.UpuData

	HandleGetIdentityData(*GetIdentityDataParams) (map[string]string, *models.IdentityData, *models.ProblemDetails)

	HandleDelete5GVnGroup(string)

	HandleCreate5GmbsGroup(string, *models.MulticastMbsGroupMemb) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleDeleteSmfRegistration(*DeleteSmfRegistrationParams)

	HandleCreateSmsfContext3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleRemovesdmSubscriptions(*RemovesdmSubscriptionsParams) *models.ProblemDetails

	HandleGetSSAuData(*GetSSAuDataParams) *models.ProblemDetails

	HandleQuerySmfRegList(*QuerySmfRegListParams) *[]models.SmfRegistration

	HandleModifyIpSmGwContext(string, *[]models.PatchItem) *models.ProblemDetails

	HandleRemoveeeSubscriptions(*RemoveeeSubscriptionsParams)

	HandleModify5GVnGroup(*Modify5GVnGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryCagAck(*QueryCagAckParams) *models.CagAckData

	HandleGetSmfSubscriptionInfo(*GetSmfSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleRemovesubscriptionDataSubscriptions(string)

	HandleQuery5GVnGroup([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleQuerySubscriptionDataSubscriptions(string) *models.SubscriptionDataSubscriptions

	HandleRemoveSmfGroupSubscriptions(*RemoveSmfGroupSubscriptionsParams)

	HandleQuerySmsData(*QuerySmsDataParams) (map[string]string, *models.SmsSubscriptionData)

	HandleDeletePPDataEntry(*DeletePPDataEntryParams) *models.ProblemDetails

	HandleGetHssSubscriptionInfo(*GetHssSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleRemoveHssSubscriptionsInfo(*RemoveHssSubscriptionsInfoParams)

	HandleModifyAmfSubscriptionInfo(*ModifyAmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryTraceData(*QueryTraceDataParams) (map[string]string, *models.TraceData)

	HandleCreateAmfContextNon3gpp(string, *models.AmfNon3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleQuerySmsfContextNon3gpp(*QuerySmsfContextNon3gppParams) *models.SmsfRegistration

	HandleCreateMessageWaitingData(string, *models.MessageWaitingData) *models.MessageWaitingData

	HandleRemoveAmfSubscriptionsInfo(*RemoveAmfSubscriptionsInfoParams)

	HandleQueryPorseData(*QueryPorseDataParams) (map[string]string, *models.ProseSubscriptionData)

	HandleGetNiddAuthorizationInfo(string) *models.NiddAuthorizationInfo

	HandleGetServiceSpecificAuthorizationInfo(*GetServiceSpecificAuthorizationInfoParams) *models.ServiceSpecificAuthorizationInfo

	HandleRemoveAmfGroupSubscriptions(*RemoveAmfGroupSubscriptionsParams)

	HandleQueryPPData(*QueryPPDataParams) *models.PpProfileData

	HandleModifyEesubscription(*ModifyEesubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateEeGroupSubscriptions(string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleGetSharedData(*GetSharedDataParams) (*[]models.SharedData, *models.ProblemDetails)

	HandleDelete5GmbsGroup(string) *models.ProblemDetails

	HandleModify5GmbsGroup(*Modify5GmbsGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryNssaiAck(*QueryNssaiAckParams) *models.NssaiAckData

	HandleQueryEEData(*QueryEEDataParams) *models.EeProfileData

	HandleCreateHSSSDMSubscriptions(*CreateHSSSDMSubscriptionsParams, *models.HssSubscriptionInfo)

	HandleQuery5GMbsGroupInternal([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)
}
