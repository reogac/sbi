/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:31 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) {
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
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId, body)

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

func OnCreateIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PolicyAssociationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualAMPolicyAssociation(body)

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

func OnReadIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualAMPolicyAssociation(polAssoId)

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

func OnDeleteIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualAMPolicyAssociation(polAssoId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)

	HandleCreateIndividualAMPolicyAssociation(*models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualAMPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleDeleteIndividualAMPolicyAssociation(string) *models.ProblemDetails
}
