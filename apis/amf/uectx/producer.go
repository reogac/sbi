/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnUeContextRelease(ctx sbi.RequestContext, handler any) {
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
	body := new(models.UeContextReleaseRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUeContextRelease(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

func OnRrcInactivityStatusReport(ctx sbi.RequestContext, handler any) {
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
	body := new(models.RrcInactivityTransportReport)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleRrcInactivityStatusReport(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

func OnN2SmInfoUplink(ctx sbi.RequestContext, handler any) {
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
	body := new(models.N2SmInfoUplinkTransport)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleN2SmInfoUplink(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

type Producer interface {
	HandleUeContextRelease(int64, *models.UeContextReleaseRequest) *models.ProblemDetails

	HandleRrcInactivityStatusReport(int64, *models.RrcInactivityTransportReport) *models.ProblemDetails

	HandleN2SmInfoUplink(int64, *models.N2SmInfoUplinkTransport) *models.ProblemDetails
}
