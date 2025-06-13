/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uepol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyAssociationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualUEPolicyAssociation(body)

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

func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualUEPolicyAssociation(polAssoId)

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

func OnDeleteIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualUEPolicyAssociation(polAssoId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyAssociationUpdateRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId, body)

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

type Producer interface {
	HandleCreateIndividualUEPolicyAssociation(*models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualUEPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleDeleteIndividualUEPolicyAssociation(string) *models.ProblemDetails

	HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)
}
