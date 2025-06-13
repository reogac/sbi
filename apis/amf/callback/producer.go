/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'sessionId'
	sessionIdStr := ctx.Param("sessionId")
	if len(sessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sessionId is required"), nil)
		return
	}

	if params.SessionId, err = models.Int16FromString(sessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse sessionId failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmContextStatusNotification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleSmContextStatusNotify(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnRanInfoUpdate(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ranId'
	var ranId string
	ranId = ctx.Param("ranId")
	if len(ranId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ranId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RanInfoUpdateData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleRanInfoUpdate(ranId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

type Producer interface {
	HandleSmContextStatusNotify(*SmContextStatusNotifyParams, *models.SmContextStatusNotification) *models.ProblemDetails

	HandleRanInfoUpdate(string, *models.RanInfoUpdateData) *models.ProblemDetails
}
