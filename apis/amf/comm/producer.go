/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package comm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnNonUeN2InfoSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NonUeN2InfoSubscriptionCreateData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleNonUeN2InfoSubscribe(body)

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

func OnAMFStatusChangeSubscribeModfy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAMFStatusChangeSubscribeModfy(subscriptionId, body)

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

func OnNonUeN2MessageTransfer(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NonUeN2MessageTransferRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleNonUeN2MessageTransfer(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromN2InformationTransferError(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCancelRelocateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.CancelRelocateUEContextRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleCancelRelocateUEContext(ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.N1N2MessageTransferRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleN1N2MessageTransfer(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromN1N2MessageTransferError(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnN1N2MessageSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UeN1N2InfoSubscriptionCreateData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleN1N2MessageSubscribe(ueContextId, body)

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

func OnNonUeN2InfoUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'n2NotifySubscriptionId'
	var n2NotifySubscriptionId string
	n2NotifySubscriptionId = ctx.Param("n2NotifySubscriptionId")
	if len(n2NotifySubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "n2NotifySubscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleNonUeN2InfoUnSubscribe(n2NotifySubscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRelocateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RelocateUEContextRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleRelocateUEContext(ueContextId, body)

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

func OnReleaseUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UEContextRelease)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleReleaseUEContext(ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnEBIAssignment(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AssignEbiData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleEBIAssignment(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromAssignEbiError(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUEContextTransfer(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UEContextTransferRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUEContextTransfer(ueContextId, body)

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

func OnRegistrationStatusUpdate(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UeRegStatusUpdateReqData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleRegistrationStatusUpdate(ueContextId, body)

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

func OnN1N2MessageUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params N1N2MessageUnSubscribeParams

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// read 'ueContextId'
	params.UeContextId = ctx.Param("ueContextId")
	if len(params.UeContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleN1N2MessageUnSubscribe(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnAMFStatusChangeSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAMFStatusChangeSubscribe(body)

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

func OnAMFStatusChangeUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleAMFStatusChangeUnSubscribe(subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.CreateUEContextRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleCreateUEContext(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUeContextCreateError(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleNonUeN2InfoSubscribe(*models.NonUeN2InfoSubscriptionCreateData) (*models.NonUeN2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleAMFStatusChangeSubscribeModfy(string, *models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails)

	HandleNonUeN2MessageTransfer(*models.NonUeN2MessageTransferRequest) (*models.N2InformationTransferRspData, *models.N2InformationTransferError, *models.ProblemDetails)

	HandleCancelRelocateUEContext(string, *models.CancelRelocateUEContextRequest) *models.ProblemDetails

	HandleN1N2MessageTransfer(string, *models.N1N2MessageTransferRequest) (*models.N1N2MessageTransferRspData, *models.N1N2MessageTransferError, *models.ProblemDetails)

	HandleN1N2MessageSubscribe(string, *models.UeN1N2InfoSubscriptionCreateData) (*models.UeN1N2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleNonUeN2InfoUnSubscribe(string) *models.ProblemDetails

	HandleRelocateUEContext(string, *models.RelocateUEContextRequest) (*models.UeContextRelocatedData, *models.ProblemDetails)

	HandleReleaseUEContext(string, *models.UEContextRelease) *models.ProblemDetails

	HandleEBIAssignment(string, *models.AssignEbiData) (*models.AssignedEbiData, *models.AssignEbiError, *models.ProblemDetails)

	HandleUEContextTransfer(string, *models.UEContextTransferRequest) (*models.UEContextTransferResponse, *models.ProblemDetails)

	HandleRegistrationStatusUpdate(string, *models.UeRegStatusUpdateReqData) (*models.UeRegStatusUpdateRspData, *models.ProblemDetails)

	HandleN1N2MessageUnSubscribe(*N1N2MessageUnSubscribeParams) *models.ProblemDetails

	HandleAMFStatusChangeSubscribe(*models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails)

	HandleAMFStatusChangeUnSubscribe(string) *models.ProblemDetails

	HandleCreateUEContext(string, *models.CreateUEContextRequest) (*models.CreateUEContextResponse, *models.UeContextCreateError, *models.ProblemDetails)
}
