/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uecm

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetNwdafRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetNwdafRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'analytics-ids'
	analyticsIdsStr := ctx.Param("analytics-ids")
	if len(analyticsIdsStr) > 0 {
		if params.AnalyticsIds, err = models.ArrayOfStringFromString(analyticsIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse analytics-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetNwdafRegistration(ctx.Context(), &params)

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

func OnSendRoutingInfoSm(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.RoutingInfoSmRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleSendRoutingInfoSm(ctx.Context(), ueId, body)

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

func OnGetNon3GppRegistration(ctx sbi.RequestContext, prod Producer) {
	var params GetNon3GppRegistrationParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetNon3GppRegistration(ctx.Context(), &params)

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

func OnGetNon3GppSmsfRegistration(ctx sbi.RequestContext, prod Producer) {
	var params GetNon3GppSmsfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetNon3GppSmsfRegistration(ctx.Context(), &params)

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

func OnNon3GppSmsfRegistration(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, prob := prod.HandleNon3GppSmsfRegistration(ctx.Context(), ueId, body)

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

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnNwdafDeregistration(ctx sbi.RequestContext, prod Producer) {
	var params NwdafDeregistrationParams

	// read 'nwdafRegistrationId'
	params.NwdafRegistrationId = ctx.Param("nwdafRegistrationId")
	if len(params.NwdafRegistrationId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "nwdafRegistrationId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleNwdafDeregistration(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnPeiUpdate(ctx sbi.RequestContext, prod Producer) {
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
	prob := prod.HandlePeiUpdate(ctx.Context(), ueId, body)

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
	rsp, prob := prod.HandleUpdateRoamingInformation(ctx.Context(), ueId, body)

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

func OnNon3GppRegistration(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleNon3GppRegistration(ctx.Context(), ueId, body)

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

func OnRetrieveSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params RetrieveSmfRegistrationParams

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
	rsp, prob := prod.HandleRetrieveSmfRegistration(ctx.Context(), &params)

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

func OnGet3GppSmsfRegistration(ctx sbi.RequestContext, prod Producer) {
	var params Get3GppSmsfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGet3GppSmsfRegistration(ctx.Context(), &params)

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

func OnSmfDeregistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params SmfDeregistrationParams

	// read 'smf-set-id'
	params.SmfSetId = ctx.Param("smf-set-id")

	// read 'smf-instance-id'
	params.SmfInstanceId = ctx.Param("smf-instance-id")

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
	prob := prod.HandleSmfDeregistration(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateSmfRegistrationParams

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
	body := new(models.SmfRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSmfRegistration(ctx.Context(), &params, body)

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

func OnRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params RegistrationParams

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
	rsp, prob := prod.HandleRegistration(ctx.Context(), &params, body)

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

func OnThreeGppSmsfDeregistration(ctx sbi.RequestContext, prod Producer) {
	var params ThreeGppSmsfDeregistrationParams

	// read 'If-Match'
	params.IfMatch = ctx.Header("If-Match")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'smsf-set-id'
	params.SmsfSetId = ctx.Param("smsf-set-id")

	// call application handler
	prob := prod.HandleThreeGppSmsfDeregistration(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnNon3GppSmsfDeregistration(ctx sbi.RequestContext, prod Producer) {
	var params Non3GppSmsfDeregistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'smsf-set-id'
	params.SmsfSetId = ctx.Param("smsf-set-id")

	// read 'If-Match'
	params.IfMatch = ctx.Header("If-Match")

	// call application handler
	prob := prod.HandleNon3GppSmsfDeregistration(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetIpSmGwRegistration(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetIpSmGwRegistration(ctx.Context(), ueId)

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

func OnGetLocationInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetLocationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetLocationInfo(ctx.Context(), &params)

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

func OnGetRegistrations(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetRegistrationsParams

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'registration-dataset-names'
	registrationDatasetNamesStr := ctx.Param("registration-dataset-names")
	if len(registrationDatasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "registration-dataset-names is required"), nil)
		return
	}

	if params.RegistrationDatasetNames, err = models.ArrayOfStringFromString(registrationDatasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse registration-dataset-names failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetRegistrations(ctx.Context(), &params)

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

func OnThreeGppRegistration(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleThreeGppRegistration(ctx.Context(), ueId, body)

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

func OnDeregAMF(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.AmfDeregInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeregAMF(ctx.Context(), ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSmfRegistration(ctx.Context(), &params)

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

func OnThreeGppSmsfRegistration(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, prob := prod.HandleThreeGppSmsfRegistration(ctx.Context(), ueId, body)

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

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnIpSmGwDeregistration(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleIpSmGwDeregistration(ctx.Context(), ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnNwdafRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params NwdafRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'nwdafRegistrationId'
	params.NwdafRegistrationId = ctx.Param("nwdafRegistrationId")
	if len(params.NwdafRegistrationId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "nwdafRegistrationId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NwdafRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleNwdafRegistration(ctx.Context(), &params, body)

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

func OnGet3GppRegistration(ctx sbi.RequestContext, prod Producer) {
	var params Get3GppRegistrationParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGet3GppRegistration(ctx.Context(), &params)

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

func OnUpdateNwdafRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateNwdafRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'nwdafRegistrationId'
	params.NwdafRegistrationId = ctx.Param("nwdafRegistrationId")
	if len(params.NwdafRegistrationId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "nwdafRegistrationId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NwdafRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateNwdafRegistration(ctx.Context(), &params, body)

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

func OnUpdate3GppRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Update3GppRegistrationParams

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
	body := new(models.Amf3GppAccessRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdate3GppRegistration(ctx.Context(), &params, body)

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

func OnUpdateNon3GppRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateNon3GppRegistrationParams

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
	body := new(models.AmfNon3GppAccessRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateNon3GppRegistration(ctx.Context(), &params, body)

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

func OnIpSmGwRegistration(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleIpSmGwRegistration(ctx.Context(), ueId, body)

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

func OnTriggerPCSCFRestoration(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.TriggerRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleTriggerPCSCFRestoration(ctx.Context(), body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleGetNwdafRegistration(context.Context, *GetNwdafRegistrationParams) (*[]models.NwdafRegistration, *models.ProblemDetails)

	HandleSendRoutingInfoSm(context.Context, string, *models.RoutingInfoSmRequest) (*models.RoutingInfoSmResponse, *models.ProblemDetails)

	HandleGetNon3GppRegistration(context.Context, *GetNon3GppRegistrationParams) (*models.AmfNon3GppAccessRegistration, *models.ProblemDetails)

	HandleGetNon3GppSmsfRegistration(context.Context, *GetNon3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleNon3GppSmsfRegistration(context.Context, string, *models.SmsfRegistration) (map[string]string, *models.SmsfRegistration, *models.ProblemDetails)

	HandleNwdafDeregistration(context.Context, *NwdafDeregistrationParams) *models.ProblemDetails

	HandlePeiUpdate(context.Context, string, *models.PeiUpdateInfo) *models.ProblemDetails

	HandleUpdateRoamingInformation(context.Context, string, *models.RoamingInfoUpdate) (*models.RoamingInfoUpdate, *models.ProblemDetails)

	HandleNon3GppRegistration(context.Context, string, *models.AmfNon3GppAccessRegistration) (*models.AmfNon3GppAccessRegistration, *models.ProblemDetails)

	HandleRetrieveSmfRegistration(context.Context, *RetrieveSmfRegistrationParams) (*models.SmfRegistration, *models.ProblemDetails)

	HandleGet3GppSmsfRegistration(context.Context, *Get3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleSmfDeregistration(context.Context, *SmfDeregistrationParams) *models.ProblemDetails

	HandleUpdateSmfRegistration(context.Context, *UpdateSmfRegistrationParams, *models.SmfRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleRegistration(context.Context, *RegistrationParams, *models.SmfRegistration) (*models.SmfRegistration, *models.ProblemDetails)

	HandleThreeGppSmsfDeregistration(context.Context, *ThreeGppSmsfDeregistrationParams) *models.ProblemDetails

	HandleNon3GppSmsfDeregistration(context.Context, *Non3GppSmsfDeregistrationParams) *models.ProblemDetails

	HandleGetIpSmGwRegistration(context.Context, string) (*models.IpSmGwRegistration, *models.ProblemDetails)

	HandleGetLocationInfo(context.Context, *GetLocationInfoParams) (*models.LocationInfo, *models.ProblemDetails)

	HandleGetRegistrations(context.Context, *GetRegistrationsParams) (*models.RegistrationDataSets, *models.ProblemDetails)

	HandleThreeGppRegistration(context.Context, string, *models.Amf3GppAccessRegistration) (*models.Amf3GppAccessRegistration, *models.ProblemDetails)

	HandleDeregAMF(context.Context, string, *models.AmfDeregInfo) *models.ProblemDetails

	HandleGetSmfRegistration(context.Context, *GetSmfRegistrationParams) (*models.SmfRegistrationInfo, *models.ProblemDetails)

	HandleThreeGppSmsfRegistration(context.Context, string, *models.SmsfRegistration) (map[string]string, *models.SmsfRegistration, *models.ProblemDetails)

	HandleIpSmGwDeregistration(context.Context, string) *models.ProblemDetails

	HandleNwdafRegistration(context.Context, *NwdafRegistrationParams, *models.NwdafRegistration) (*models.NwdafRegistration, *models.ProblemDetails)

	HandleGet3GppRegistration(context.Context, *Get3GppRegistrationParams) (*models.Amf3GppAccessRegistration, *models.ProblemDetails)

	HandleUpdateNwdafRegistration(context.Context, *UpdateNwdafRegistrationParams, *models.NwdafRegistrationModification) (*models.Schema, *models.ProblemDetails)

	HandleUpdate3GppRegistration(context.Context, *Update3GppRegistrationParams, *models.Amf3GppAccessRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleUpdateNon3GppRegistration(context.Context, *UpdateNon3GppRegistrationParams, *models.AmfNon3GppAccessRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleIpSmGwRegistration(context.Context, string, *models.IpSmGwRegistration) (*models.IpSmGwRegistration, *models.ProblemDetails)

	HandleTriggerPCSCFRestoration(context.Context, *models.TriggerRequest) *models.ProblemDetails
}
