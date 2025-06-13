/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uecm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnUpdate3GppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params Update3GppRegistrationParams

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
	body := new(models.Amf3GppAccessRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdate3GppRegistration(&params, body)

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

func OnDeregAMF(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	prob := prod.HandleDeregAMF(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnIpSmGwRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleIpSmGwRegistration(ueId, body)

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

func OnSendRoutingInfoSm(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleSendRoutingInfoSm(ueId, body)

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

func OnThreeGppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleThreeGppRegistration(ueId, body)

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

func OnThreeGppSmsfDeregistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params ThreeGppSmsfDeregistrationParams

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
	prob := prod.HandleThreeGppSmsfDeregistration(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnNon3GppSmsfDeregistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	prob := prod.HandleNon3GppSmsfDeregistration(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnIpSmGwDeregistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleIpSmGwDeregistration(ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnNwdafDeregistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params NwdafDeregistrationParams

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

	// call application handler
	prob := prod.HandleNwdafDeregistration(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetNon3GppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetNon3GppRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetNon3GppRegistration(&params)

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

func OnUpdateSmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleUpdateSmfRegistration(&params, body)

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

func OnGet3GppSmsfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleGet3GppSmsfRegistration(&params)

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

func OnGetNon3GppSmsfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetNon3GppSmsfRegistrationParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetNon3GppSmsfRegistration(&params)

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

func OnGetLocationInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleGetLocationInfo(&params)

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

func OnNon3GppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleNon3GppRegistration(ueId, body)

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

func OnGetRegistrations(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleGetRegistrations(&params)

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

func OnUpdateNon3GppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateNon3GppRegistrationParams

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
	body := new(models.AmfNon3GppAccessRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateNon3GppRegistration(&params, body)

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

func OnGetSmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSmfRegistrationParams

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

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSmfRegistration(&params)

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

func OnTriggerPCSCFRestoration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.TriggerRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleTriggerPCSCFRestoration(body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateNwdafRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params UpdateNwdafRegistrationParams

	// read 'nwdafRegistrationId'
	params.NwdafRegistrationId = ctx.Param("nwdafRegistrationId")
	if len(params.NwdafRegistrationId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "nwdafRegistrationId is required"), nil)
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

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NwdafRegistrationModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateNwdafRegistration(&params, body)

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

func OnGet3GppRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params Get3GppRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGet3GppRegistration(&params)

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

func OnPeiUpdate(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	prob := prod.HandlePeiUpdate(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnSmfDeregistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params SmfDeregistrationParams

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

	// read 'smf-set-id'
	params.SmfSetId = ctx.Param("smf-set-id")

	// read 'smf-instance-id'
	params.SmfInstanceId = ctx.Param("smf-instance-id")

	// call application handler
	prob := prod.HandleSmfDeregistration(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params RegistrationParams

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

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
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
	rsp, prob := prod.HandleRegistration(&params, body)

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

func OnNon3GppSmsfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleNon3GppSmsfRegistration(ueId, body)

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

func OnUpdateRoamingInformation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleUpdateRoamingInformation(ueId, body)

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

func OnRetrieveSmfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleRetrieveSmfRegistration(&params)

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

func OnThreeGppSmsfRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleThreeGppSmsfRegistration(ueId, body)

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

func OnGetIpSmGwRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetIpSmGwRegistration(ueId)

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

func OnGetNwdafRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleGetNwdafRegistration(&params)

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

func OnNwdafRegistration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params NwdafRegistrationParams

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

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NwdafRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleNwdafRegistration(&params, body)

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
	HandleUpdate3GppRegistration(*Update3GppRegistrationParams, *models.Amf3GppAccessRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleDeregAMF(string, *models.AmfDeregInfo) *models.ProblemDetails

	HandleIpSmGwRegistration(string, *models.IpSmGwRegistration) (*models.IpSmGwRegistration, *models.ProblemDetails)

	HandleSendRoutingInfoSm(string, *models.RoutingInfoSmRequest) (*models.RoutingInfoSmResponse, *models.ProblemDetails)

	HandleThreeGppRegistration(string, *models.Amf3GppAccessRegistration) (*models.Amf3GppAccessRegistration, *models.ProblemDetails)

	HandleThreeGppSmsfDeregistration(*ThreeGppSmsfDeregistrationParams) *models.ProblemDetails

	HandleNon3GppSmsfDeregistration(*Non3GppSmsfDeregistrationParams) *models.ProblemDetails

	HandleIpSmGwDeregistration(string) *models.ProblemDetails

	HandleNwdafDeregistration(*NwdafDeregistrationParams) *models.ProblemDetails

	HandleGetNon3GppRegistration(*GetNon3GppRegistrationParams) (*models.AmfNon3GppAccessRegistration, *models.ProblemDetails)

	HandleUpdateSmfRegistration(*UpdateSmfRegistrationParams, *models.SmfRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleGet3GppSmsfRegistration(*Get3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleGetNon3GppSmsfRegistration(*GetNon3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleGetLocationInfo(*GetLocationInfoParams) (*models.LocationInfo, *models.ProblemDetails)

	HandleNon3GppRegistration(string, *models.AmfNon3GppAccessRegistration) (*models.AmfNon3GppAccessRegistration, *models.ProblemDetails)

	HandleGetRegistrations(*GetRegistrationsParams) (*models.RegistrationDataSets, *models.ProblemDetails)

	HandleUpdateNon3GppRegistration(*UpdateNon3GppRegistrationParams, *models.AmfNon3GppAccessRegistrationModification) (*models.PatchResult, *models.ProblemDetails)

	HandleGetSmfRegistration(*GetSmfRegistrationParams) (*models.SmfRegistrationInfo, *models.ProblemDetails)

	HandleTriggerPCSCFRestoration(*models.TriggerRequest) *models.ProblemDetails

	HandleUpdateNwdafRegistration(*UpdateNwdafRegistrationParams, *models.NwdafRegistrationModification) (*models.Schema, *models.ProblemDetails)

	HandleGet3GppRegistration(*Get3GppRegistrationParams) (*models.Amf3GppAccessRegistration, *models.ProblemDetails)

	HandlePeiUpdate(string, *models.PeiUpdateInfo) *models.ProblemDetails

	HandleSmfDeregistration(*SmfDeregistrationParams) *models.ProblemDetails

	HandleRegistration(*RegistrationParams, *models.SmfRegistration) (*models.SmfRegistration, *models.ProblemDetails)

	HandleNon3GppSmsfRegistration(string, *models.SmsfRegistration) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleUpdateRoamingInformation(string, *models.RoamingInfoUpdate) (*models.RoamingInfoUpdate, *models.ProblemDetails)

	HandleRetrieveSmfRegistration(*RetrieveSmfRegistrationParams) (*models.SmfRegistration, *models.ProblemDetails)

	HandleThreeGppSmsfRegistration(string, *models.SmsfRegistration) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleGetIpSmGwRegistration(string) (*models.IpSmGwRegistration, *models.ProblemDetails)

	HandleGetNwdafRegistration(*GetNwdafRegistrationParams) (*[]models.NwdafRegistration, *models.ProblemDetails)

	HandleNwdafRegistration(*NwdafRegistrationParams, *models.NwdafRegistration) (*models.NwdafRegistration, *models.ProblemDetails)
}
