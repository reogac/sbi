/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateSMPolicy(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmPolicyContextData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleCreateSMPolicy(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetSMPolicy(ctx sbi.RequestContext, prod Producer) {

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSMPolicy(ctx.Context(), smPolicyId)

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

func OnUpdateSMPolicy(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmPolicyUpdateContextData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSMPolicy(ctx.Context(), smPolicyId, body)

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

func OnDeleteSMPolicy(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmPolicyDeleteData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteSMPolicy(ctx.Context(), smPolicyId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleCreateSMPolicy(context.Context, *models.SmPolicyContextData) (map[string]string, *models.SmPolicyDecision, *models.ProblemDetails)

	HandleGetSMPolicy(context.Context, string) (*models.SmPolicyControl, *models.ProblemDetails)

	HandleUpdateSMPolicy(context.Context, string, *models.SmPolicyUpdateContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleDeleteSMPolicy(context.Context, string, *models.SmPolicyDeleteData) *models.ProblemDetails
}
