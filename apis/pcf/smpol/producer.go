/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
	body := new(models.SmPolicyContextData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateSMPolicy(body)

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

func OnGetSMPolicy(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSMPolicy(smPolicyId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// decode request body
	body := new(models.SmPolicyUpdateContextData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSMPolicy(smPolicyId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
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
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// decode request body
	body := new(models.SmPolicyDeleteData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleDeleteSMPolicy(smPolicyId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleCreateSMPolicy(*models.SmPolicyContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleGetSMPolicy(string) (*models.SmPolicyControl, *models.ProblemDetails)

	HandleUpdateSMPolicy(string, *models.SmPolicyUpdateContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleDeleteSMPolicy(string, *models.SmPolicyDeleteData) *models.ProblemDetails
}
