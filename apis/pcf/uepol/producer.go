/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:32 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uepol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnDeleteIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualUEPolicyAssociation(polAssoId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// decode request body
	body := new(models.PolicyAssociationUpdateRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId, body)

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

func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PolicyAssociationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualUEPolicyAssociation(body)

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

func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualUEPolicyAssociation(polAssoId)

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

type Producer interface {
	HandleDeleteIndividualUEPolicyAssociation(string) *models.ProblemDetails

	HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)

	HandleCreateIndividualUEPolicyAssociation(*models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualUEPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)
}
