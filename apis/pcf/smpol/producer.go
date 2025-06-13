/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateSMPolicy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmPolicyContextData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateSMPolicy(body)

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

func OnGetSMPolicy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSMPolicy(smPolicyId)

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

func OnUpdateSMPolicy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	rsp, prob := prod.HandleUpdateSMPolicy(smPolicyId, body)

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

func OnDeleteSMPolicy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
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
	prob := prod.HandleDeleteSMPolicy(smPolicyId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleCreateSMPolicy(*models.SmPolicyContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleGetSMPolicy(string) (*models.SmPolicyControl, *models.ProblemDetails)

	HandleUpdateSMPolicy(string, *models.SmPolicyUpdateContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleDeleteSMPolicy(string, *models.SmPolicyDeleteData) *models.ProblemDetails
}
