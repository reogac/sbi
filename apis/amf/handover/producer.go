/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnHandoverRequired(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId int64
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	if ueId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HandoverRequired)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleHandoverRequired(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(400, ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnHandoverNotify(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId int64
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	if ueId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HandoverNotify)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleHandoverNotify(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnHandoverCancel(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId int64
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	if ueId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HandoverCancel)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleHandoverCancel(ueId, body)

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

func OnPathSwitch(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params PathSwitchParams

	// read 'callback'
	callbackStr := ctx.Header("callback")
	if len(callbackStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "callback is required"), nil)
		return
	}

	if params.Callback, err = models.EndpointInfoFromString(callbackStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse callback failed: %+v", err)), nil)
		return
	}

	// read 'ueId'
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	if params.UeId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PathSwitchRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandlePathSwitch(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(400, ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleHandoverRequired(int64, *models.HandoverRequired) (*models.HandoverCommand, *models.HandoverPreparationFailure, *models.ProblemDetails)

	HandleHandoverNotify(int64, *models.HandoverNotify) *models.ProblemDetails

	HandleHandoverCancel(int64, *models.HandoverCancel) (*models.HandoverCancelAcknowledge, *models.ProblemDetails)

	HandlePathSwitch(*PathSwitchParams, *models.PathSwitchRequest) (*models.PathSwitchAcknowledge, *models.PathSwitchFailure, *models.ProblemDetails)
}
