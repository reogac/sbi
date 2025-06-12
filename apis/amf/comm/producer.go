/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package comm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnAMFStatusChangeUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleAMFStatusChangeUnSubscribe(subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnReleaseUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.UEContextRelease)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleReleaseUEContext(ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnRelocateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.RelocateUEContextRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRelocateUEContext(ueContextId, body)

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

func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.N1N2MessageTransferRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleN1N2MessageTransfer(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromN1N2MessageTransferError(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnN1N2MessageUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params N1N2MessageUnSubscribeParams

	// read 'ueContextId'
	params.UeContextId = ctx.Param("ueContextId")
	if len(params.UeContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleN1N2MessageUnSubscribe(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnNonUeN2MessageTransfer(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.NonUeN2MessageTransferRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleNonUeN2MessageTransfer(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromN2InformationTransferError(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnEBIAssignment(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.AssignEbiData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleEBIAssignment(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromAssignEbiError(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.UEContextTransferRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUEContextTransfer(ueContextId, body)

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

func OnNonUeN2InfoSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.NonUeN2InfoSubscriptionCreateData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleNonUeN2InfoSubscribe(body)

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

func OnAMFStatusChangeSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.SubscriptionData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAMFStatusChangeSubscribe(body)

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

func OnAMFStatusChangeSubscribeModfy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// decode request body
	body := new(models.SubscriptionData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAMFStatusChangeSubscribeModfy(subscriptionId, body)

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

func OnCreateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.CreateUEContextRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleCreateUEContext(ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUeContextCreateError(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.UeRegStatusUpdateReqData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRegistrationStatusUpdate(ueContextId, body)

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

func OnCancelRelocateUEContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.CancelRelocateUEContextRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleCancelRelocateUEContext(ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnN1N2MessageSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueContextId'
	var ueContextId string
	ueContextId = ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"))
		return
	}

	// decode request body
	body := new(models.UeN1N2InfoSubscriptionCreateData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleN1N2MessageSubscribe(ueContextId, body)

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

func OnNonUeN2InfoUnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'n2NotifySubscriptionId'
	var n2NotifySubscriptionId string
	n2NotifySubscriptionId = ctx.Param("n2NotifySubscriptionId")
	if len(n2NotifySubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "n2NotifySubscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleNonUeN2InfoUnSubscribe(n2NotifySubscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleAMFStatusChangeUnSubscribe(string) *models.ProblemDetails

	HandleReleaseUEContext(string, *models.UEContextRelease) *models.ProblemDetails

	HandleRelocateUEContext(string, *models.RelocateUEContextRequest) (*models.UeContextRelocatedData, *models.ProblemDetails)

	HandleN1N2MessageTransfer(string, *models.N1N2MessageTransferRequest) (*models.N1N2MessageTransferRspData, *models.N1N2MessageTransferError, *models.ProblemDetails)

	HandleN1N2MessageUnSubscribe(*N1N2MessageUnSubscribeParams) *models.ProblemDetails

	HandleNonUeN2MessageTransfer(*models.NonUeN2MessageTransferRequest) (*models.N2InformationTransferRspData, *models.N2InformationTransferError, *models.ProblemDetails)

	HandleEBIAssignment(string, *models.AssignEbiData) (*models.AssignedEbiData, *models.AssignEbiError, *models.ProblemDetails)

	HandleUEContextTransfer(string, *models.UEContextTransferRequest) (*models.UEContextTransferResponse, *models.ProblemDetails)

	HandleNonUeN2InfoSubscribe(*models.NonUeN2InfoSubscriptionCreateData) (*models.NonUeN2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleAMFStatusChangeSubscribe(*models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails)

	HandleAMFStatusChangeSubscribeModfy(string, *models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails)

	HandleCreateUEContext(string, *models.CreateUEContextRequest) (*models.CreateUEContextResponse, *models.UeContextCreateError, *models.ProblemDetails)

	HandleRegistrationStatusUpdate(string, *models.UeRegStatusUpdateReqData) (*models.UeRegStatusUpdateRspData, *models.ProblemDetails)

	HandleCancelRelocateUEContext(string, *models.CancelRelocateUEContextRequest) *models.ProblemDetails

	HandleN1N2MessageSubscribe(string, *models.UeN1N2InfoSubscriptionCreateData) (*models.UeN1N2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleNonUeN2InfoUnSubscribe(string) *models.ProblemDetails
}
