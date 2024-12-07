/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:20 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnSessionResourceNotify(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smCtxRef'
	var smCtxRef string
	smCtxRef = ctx.Param("smCtxRef")
	if len(smCtxRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smCtxRef is required"))
		return
	}

	// decode request body
	body := new(models.SessionResourceNotification)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleSessionResourceNotify(smCtxRef, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

func OnSessionResourceModifyIndication(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smCtxRef'
	var smCtxRef string
	smCtxRef = ctx.Param("smCtxRef")
	if len(smCtxRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smCtxRef is required"))
		return
	}

	// decode request body
	body := new(models.SessionResourceModifyIndication)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionResourceModifyIndication(smCtxRef, body)

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

type Producer interface {
	HandleSessionResourceNotify(string, *models.SessionResourceNotification) *models.ProblemDetails

	HandleSessionResourceModifyIndication(string, *models.SessionResourceModifyIndication) (*models.SessionResourceModifyConfirm, *models.ProblemDetails)
}
