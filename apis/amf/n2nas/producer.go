/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnInitialUeMessage(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'callback'
	var callback *models.EndpointInfo
	callbackStr := ctx.Header("callback")
	if len(callbackStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "callback is required"))
		return
	}

	if callback, err = models.EndpointInfoFromString(callbackStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse callback failed: %+v", err)))
		return
	}

	// decode request body
	body := new(models.InitialUeMessage)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleInitialUeMessage(callback, body)

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

func OnNasUl(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId int64
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	if ueId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)))
		return
	}

	// decode request body
	body := new(models.NasUplinkTransport)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleNasUl(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

func OnNasErr(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId int64
	ueIdStr := ctx.Param("ueId")
	if len(ueIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	if ueId, err = models.Int64FromString(ueIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ueId failed: %+v", err)))
		return
	}

	// decode request body
	body := new(models.UplinkNasError)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleNasErr(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

type Producer interface {
	HandleInitialUeMessage(*models.EndpointInfo, *models.InitialUeMessage) (*models.InitialUeMessageResponse, *models.ProblemDetails)

	HandleNasUl(int64, *models.NasUplinkTransport) *models.ProblemDetails

	HandleNasErr(int64, *models.UplinkNasError) *models.ProblemDetails
}
