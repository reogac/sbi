/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:33 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

func OnUpdateAmfUeContextInfo(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.AmfUeContextInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateAmfUeContextInfo(ctx.Context(), ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnUeContextModify(ctx sbi.RequestContext, prod Producer) {
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
	var body *models.UeContextModifyRequest
	body = new(models.UeContextModifyRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, ersp, prob := prod.HandleUeContextModify(ctx.Context(), ueId, body)

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

func OnUeContextRelease(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.UeContextReleaseCommand)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUeContextRelease(ctx.Context(), ueId, body)

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

func OnUeContextSetup(ctx sbi.RequestContext, prod Producer) {
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
	var body *models.UeContextSetupRequest
	body = new(models.UeContextSetupRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, ersp, prob := prod.HandleUeContextSetup(ctx.Context(), ueId, body)

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
	HandleUpdateAmfUeContextInfo(context.Context, int64, *models.AmfUeContextInfo) *models.ProblemDetails

	HandleUeContextModify(context.Context, int64, *models.UeContextModifyRequest) (*models.UeContextModifyResponse, *models.UeContextModifyFailure, *models.ProblemDetails)

	HandleUeContextRelease(context.Context, int64, *models.UeContextReleaseCommand) (*models.UeContextReleaseComplete, *models.ProblemDetails)

	HandleUeContextSetup(context.Context, int64, *models.UeContextSetupRequest) (*models.UeContextSetupResponse, *models.UeContextSetupFailure, *models.ProblemDetails)
}
