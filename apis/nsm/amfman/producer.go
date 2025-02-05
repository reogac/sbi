/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Feb  5 09:54:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnAmfRegister(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.AmfRegistrationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfRegister(body)

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
	HandleAmfRegister(*models.AmfRegistrationRequest) (*models.AmfRegistrationResponse, *models.ProblemDetails)
}
