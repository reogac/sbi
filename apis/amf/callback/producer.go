/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:18 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package callback

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnSmContextStatusNotify(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params SmContextStatusNotifyParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'sessionId'
	sessionIdStr := ctx.Param("sessionId")
	if len(sessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sessionId is required"))
		return
	}

	if params.SessionId, err = models.Int16FromString(sessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse sessionId failed: %+v", err)))
		return
	}

	// decode request body
	body := new(models.SmContextStatusNotification)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleSmContextStatusNotify(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

type Producer interface {
	HandleSmContextStatusNotify(*SmContextStatusNotifyParams, *models.SmContextStatusNotification) *models.ProblemDetails
}
