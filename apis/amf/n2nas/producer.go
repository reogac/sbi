/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnInitialUeMessage(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'callback'
	var callback *models.EndpointInfo
	callbackStr := ctx.Header("callback")
	if len(callbackStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "callback is required"), nil)
		return
	}

	if callback, err = models.EndpointInfoFromString(callbackStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse callback failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.InitialUeMessage)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleInitialUeMessage(callback, body)

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

func OnNasUl(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.NasUplinkTransport)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleNasUl(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnNasErr(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.UplinkNasError)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleNasErr(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

type Producer interface {
	HandleInitialUeMessage(*models.EndpointInfo, *models.InitialUeMessage) (*models.InitialUeMessageResponse, *models.ProblemDetails)

	HandleNasUl(int64, *models.NasUplinkTransport) *models.ProblemDetails

	HandleNasErr(int64, *models.UplinkNasError) *models.ProblemDetails
}
