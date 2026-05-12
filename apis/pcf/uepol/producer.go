/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uepol

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyAssociationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleCreateIndividualUEPolicyAssociation(ctx.Context(), body)

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

func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, prod Producer) {

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualUEPolicyAssociation(ctx.Context(), polAssoId)

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

func OnDeleteIndividualUEPolicyAssociation(ctx sbi.RequestContext, prod Producer) {

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualUEPolicyAssociation(ctx.Context(), polAssoId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, prod Producer) {
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
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx.Context(), polAssoId, body)

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
	HandleCreateIndividualUEPolicyAssociation(context.Context, *models.PolicyAssociationRequest) (map[string]string, *models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualUEPolicyAssociation(context.Context, string) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleDeleteIndividualUEPolicyAssociation(context.Context, string) *models.ProblemDetails

	HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(context.Context, string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)
}
