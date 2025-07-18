/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnHandoverRequest(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.HandoverRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleHandoverRequest(callback, body)

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
	HandleHandoverRequest(*models.EndpointInfo, *models.HandoverRequest) (*models.HandoverRequestAcknowledge, *models.HandoverRequestFailure, *models.ProblemDetails)
}
