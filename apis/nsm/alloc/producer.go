/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package alloc

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnAmfRegister(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AmfRegistrationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfRegister(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnAllocateIpSegments(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.IpSegmentRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAllocateIpSegments(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReleaseIpSegments(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.IpSegmentReleaseRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleReleaseIpSegments(ctx.Context(), body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleAmfRegister(context.Context, *models.AmfRegistrationRequest) (*models.AmfRegistrationResponse, *models.ProblemDetails)

	HandleAllocateIpSegments(context.Context, *models.IpSegmentRequest) (*models.IpSegmentResponse, *models.ProblemDetails)

	HandleReleaseIpSegments(context.Context, *models.IpSegmentReleaseRequest) *models.ProblemDetails
}
