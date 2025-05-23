/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:04:53 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package nasdl

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnNasDl(ctx sbi.RequestContext, handler any) {
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
	body := new(models.NasDownlinkTransport)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleNasDl(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

type Producer interface {
	HandleNasDl(int64, *models.NasDownlinkTransport) *models.ProblemDetails
}
