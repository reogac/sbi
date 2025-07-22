/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnSessionResourceNotify(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smCtxRef'
	var smCtxRef string
	smCtxRef = ctx.Param("smCtxRef")
	if len(smCtxRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smCtxRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SessionResourceNotification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleSessionResourceNotify(smCtxRef, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnSessionResourceModifyIndication(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smCtxRef'
	var smCtxRef string
	smCtxRef = ctx.Param("smCtxRef")
	if len(smCtxRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smCtxRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SessionResourceModifyIndication)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionResourceModifyIndication(smCtxRef, body)

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

type Producer interface {
	HandleSessionResourceNotify(string, *models.SessionResourceNotification) *models.ProblemDetails

	HandleSessionResourceModifyIndication(string, *models.SessionResourceModifyIndication) (*models.SessionResourceModifyConfirm, *models.ProblemDetails)
}
