/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package comm

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnNonUeN2InfoSubscribe(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NonUeN2InfoSubscriptionCreateData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleNonUeN2InfoSubscribe(ctx.Context(), body)

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

func OnReleaseUEContext(ctx sbi.RequestContext, prod Producer) {
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
	prob := prod.HandleReleaseUEContext(ctx.Context(), ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnAMFStatusChangeSubscribe(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleAMFStatusChangeSubscribe(ctx.Context(), body)

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

func OnAMFStatusChangeSubscribeModfy(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleAMFStatusChangeSubscribeModfy(ctx.Context(), subscriptionId, body)

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

func OnAMFStatusChangeUnSubscribe(ctx sbi.RequestContext, prod Producer) {

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleAMFStatusChangeUnSubscribe(ctx.Context(), subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateUEContext(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, ersp, prob := prod.HandleCreateUEContext(ctx.Context(), ueContextId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
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

func OnUEContextTransfer(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleUEContextTransfer(ctx.Context(), ueContextId, body)

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

func OnRelocateUEContext(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, prob := prod.HandleRelocateUEContext(ctx.Context(), ueContextId, body)

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

func OnN1N2MessageTransfer(ctx sbi.RequestContext, prod Producer) {
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
	rsp, ersp, prob := prod.HandleN1N2MessageTransfer(ctx.Context(), ueContextId, body)

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

func OnN1N2MessageSubscribe(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, prob := prod.HandleN1N2MessageSubscribe(ctx.Context(), ueContextId, body)

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

func OnNonUeN2MessageTransfer(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NonUeN2MessageTransferRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleNonUeN2MessageTransfer(ctx.Context(), body)

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

func OnRegistrationStatusUpdate(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleRegistrationStatusUpdate(ctx.Context(), ueContextId, body)

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

func OnNonUeN2InfoUnSubscribe(ctx sbi.RequestContext, prod Producer) {

	// read 'n2NotifySubscriptionId'
	var n2NotifySubscriptionId string
	n2NotifySubscriptionId = ctx.Param("n2NotifySubscriptionId")
	if len(n2NotifySubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "n2NotifySubscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleNonUeN2InfoUnSubscribe(ctx.Context(), n2NotifySubscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnEBIAssignment(ctx sbi.RequestContext, prod Producer) {
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
	rsp, ersp, prob := prod.HandleEBIAssignment(ctx.Context(), ueContextId, body)

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

func OnCancelRelocateUEContext(ctx sbi.RequestContext, prod Producer) {
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
	prob := prod.HandleCancelRelocateUEContext(ctx.Context(), ueContextId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnN1N2MessageUnSubscribe(ctx sbi.RequestContext, prod Producer) {
	var params N1N2MessageUnSubscribeParams

	// read 'ueContextId'
	params.UeContextId = ctx.Param("ueContextId")
	if len(params.UeContextId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueContextId is required"), nil)
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleN1N2MessageUnSubscribe(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleNonUeN2InfoSubscribe(context.Context, *models.NonUeN2InfoSubscriptionCreateData) (map[string]string, *models.NonUeN2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleReleaseUEContext(context.Context, string, *models.UEContextRelease) *models.ProblemDetails

	HandleAMFStatusChangeSubscribe(context.Context, *models.SubscriptionData) (map[string]string, *models.SubscriptionData, *models.ProblemDetails)

	HandleAMFStatusChangeSubscribeModfy(context.Context, string, *models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails)

	HandleAMFStatusChangeUnSubscribe(context.Context, string) *models.ProblemDetails

	HandleCreateUEContext(context.Context, string, *models.CreateUEContextRequest) (map[string]string, *models.CreateUEContextResponse, *models.UeContextCreateError, *models.ProblemDetails)

	HandleUEContextTransfer(context.Context, string, *models.UEContextTransferRequest) (*models.UEContextTransferResponse, *models.ProblemDetails)

	HandleRelocateUEContext(context.Context, string, *models.RelocateUEContextRequest) (map[string]string, *models.UeContextRelocatedData, *models.ProblemDetails)

	HandleN1N2MessageTransfer(context.Context, string, *models.N1N2MessageTransferRequest) (*models.N1N2MessageTransferRspData, *models.N1N2MessageTransferError, *models.ProblemDetails)

	HandleN1N2MessageSubscribe(context.Context, string, *models.UeN1N2InfoSubscriptionCreateData) (map[string]string, *models.UeN1N2InfoSubscriptionCreatedData, *models.ProblemDetails)

	HandleNonUeN2MessageTransfer(context.Context, *models.NonUeN2MessageTransferRequest) (*models.N2InformationTransferRspData, *models.N2InformationTransferError, *models.ProblemDetails)

	HandleRegistrationStatusUpdate(context.Context, string, *models.UeRegStatusUpdateReqData) (*models.UeRegStatusUpdateRspData, *models.ProblemDetails)

	HandleNonUeN2InfoUnSubscribe(context.Context, string) *models.ProblemDetails

	HandleEBIAssignment(context.Context, string, *models.AssignEbiData) (*models.AssignedEbiData, *models.AssignEbiError, *models.ProblemDetails)

	HandleCancelRelocateUEContext(context.Context, string, *models.CancelRelocateUEContextRequest) *models.ProblemDetails

	HandleN1N2MessageUnSubscribe(context.Context, *N1N2MessageUnSubscribeParams) *models.ProblemDetails
}
