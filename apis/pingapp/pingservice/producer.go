/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:07 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pingservice

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnPing(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PingRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandlePing(body)

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

func OnForward(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PingFwRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleForward(body)

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
	HandlePing(*models.PingRequest) (*models.PingResponse, *models.ProblemDetails)

	HandleForward(*models.PingFwRequest) (*models.PingResponse, *models.ProblemDetails)
}
