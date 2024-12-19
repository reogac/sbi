/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:28:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnHandoverRequest(ctx sbi.RequestContext, handler any) {
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
	body := new(models.HandoverRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleHandoverRequest(callback, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(400, ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

type Producer interface {
	HandleHandoverRequest(*models.EndpointInfo, *models.HandoverRequest) (*models.HandoverRequestAcknowledge, *models.HandoverRequestFailure, *models.ProblemDetails)
}
