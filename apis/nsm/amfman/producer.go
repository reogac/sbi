/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Jul 21 09:39:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
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
	rsp, prob := prod.HandleAmfRegister(body)

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

func OnGetSupportedSlices(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.GetSupportedSlicesRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSupportedSlices(body)

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
	HandleAmfRegister(*models.AmfRegistrationRequest) (*models.AmfRegistrationResponse, *models.ProblemDetails)

	HandleGetSupportedSlices(*models.GetSupportedSlicesRequest) (*models.GetSupportedSlicesResponse, *models.ProblemDetails)
}
